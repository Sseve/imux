package main

import (
	"net/http"

	"github.com/Sseve/imux"
)

// Logger 中间件
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("Request:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Auth 认证中间件
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
                // 认证 token 逻辑
                // ...
		next.ServeHTTP(w, r)
	})
}

type UserForm struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {
	mux := imux.NewRouter()
	// 添加全局中间件
	mux.Use(Logger)

	mux.Get("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		imux.Success(w, imux.Map{"code": 200, "message": "pong"})
	}))

        // 获取 < /pong?name=zhangsan&password=123456 > 查询参数
	mux.Get("/pong", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		password := r.URL.Query().Get("password")
		imux.Success(w, imux.Map{"code": 200, "message": "pong", "name": name, "password": password})
	}))

	// 路由分组
	api := mux.Group("/api", Auth)
	api.Get("/user/:id", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := imux.Param(r, "id")
		imux.Success(w, imux.Map{"code": 200, "message": "Get user id: " + id})
	}))

	api.Post("/user", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user UserForm
		if err := imux.Bind(r, &user); err != nil {
			imux.Failure(w, imux.Map{"code": 500, "message": "bind user error"})
			return
		}
		imux.Success(w, imux.Map{"code": 200, "message": "Create user success"})
	}))

	// 启动服务
	server := http.Server{
		Addr:    ":9990",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Sseve/imux"
	mapi "github.com/Sseve/imux/_example/api"
	"github.com/Sseve/imux/env"
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

// -ldflags="-X main.version=0.1.1"
var version = "0.1.0"

func main() {
	fmt.Println("版本: ", version)
	// 加载配置
	env.LoadEnv(".env")
	mux := imux.NewRouter()
	// 添加全局中间件
	mux.Use(Logger)

	mux.Get("/ping", mapi.Ping)
	// 获取 < /pong?foo=FOO&bar=BAR > 查询参数
	mux.Get("/pong", mapi.Pong)

	// 路由分组
	api := mux.Group("/api", Auth)
	api.Get("/foo/:id", mapi.FooId)
	api.Post("/foo", mapi.Foo)

	// 路由分组: restful api 接口
	v1 := mux.Group("/v1")
	v1.Get("/hello", mapi.HelloGet)
	v1.Post("/hello", mapi.HelloPost)
	v1.Delete("/hello", mapi.HelloDelete)
	v1.Put("/hello", mapi.HelloPut)

	// 启动服务
	address := os.Getenv("app.address")
	server := http.Server{Addr: address, Handler: mux}
	fmt.Printf("App Run [%s]\n", address)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

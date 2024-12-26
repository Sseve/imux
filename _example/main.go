package main

import (
	"fmt"
	"github.com/Sseve/imux/env"
	"net/http"
	"os"

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

var (
	version string = "0.1.0"
)

func main() {
	fmt.Println("当前版本: ", version)
	// 加载配置
	env.LoadEnv(".env")
	mux := imux.NewRouter()
	// 添加全局中间件
	mux.Use(Logger)

	mux.Get("/ping", ping)
	// 获取 < /pong?foo=FOO&bar=BAR > 查询参数
	mux.Get("/pong", pong)

	// 路由分组
	api := mux.Group("/api", Auth)
	api.Get("/foo/:id", fooId)
	api.Post("/foo", foo)

	// 路由分组: RESTFul API
	v1 := mux.Group("/v1")
	v1.Get("/hello", helloGet)
	v1.Post("/hello", helloPost)
	v1.Delete("/hello", helloDelete)
	v1.Put("/hello", helloPut)

	// 启动服务
	address := os.Getenv("app.address")
	server := http.Server{Addr: address, Handler: mux}
	fmt.Printf("App Run [%s]\n", address)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// 全局路由控制器示例
func ping(w http.ResponseWriter, r *http.Request) {
	imux.Success(w, imux.Map{"code": 200, "message": "pong"})
}

// 查询参数接口示例 < /pong?foo=FOO&bar=BAR >
func pong(w http.ResponseWriter, r *http.Request) {
	foo := r.URL.Query().Get("foo")
	bar := r.URL.Query().Get("bar")
	imux.Success(w, imux.Map{"code": 200, "foo": foo, "bar": bar})
}

// 路径参数接口示例 </api/hello/:id>
func fooId(w http.ResponseWriter, r *http.Request) {
	id := imux.Param(r, "id")
	imux.Success(w, imux.Map{"code": 200, "message": "Get foo id: " + id})
}

// 请求参数绑定接口示例
func foo(w http.ResponseWriter, r *http.Request) {
	// 请求参数schema
	type FooForm struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}
	var fooForm FooForm
	if err := imux.Bind(r, &fooForm); err != nil {
		imux.Failure(w, imux.Map{"code": 500, "message": "bind foo error"})
		return
	}
	imux.Success(w, imux.Map{"code": 200, "fooForm": fooForm})
}

// RESTFul API 接口示例
func helloGet(w http.ResponseWriter, r *http.Request) {
	imux.Success(w, imux.Map{"code": 200, "message": "Hello Get"})
}

func helloPost(w http.ResponseWriter, r *http.Request) {
	imux.Success(w, imux.Map{"code": 200, "message": "Hello Post"})
}

func helloDelete(w http.ResponseWriter, r *http.Request) {
	imux.Success(w, imux.Map{"code": 200, "message": "Hello Delete"})
}

func helloPut(w http.ResponseWriter, r *http.Request) {
	imux.Success(w, imux.Map{"code": 200, "message": "Hello Put"})
}

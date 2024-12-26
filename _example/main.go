package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Sseve/imux"
	"github.com/Sseve/imux/_example/api"
	"github.com/Sseve/imux/_example/mws"
	"github.com/Sseve/imux/env"
)

// -ldflags="-X main.version=0.1.1"
var version = "0.1.0"

func main() {
	fmt.Println("版本: ", version)
	// 加载配置
	env.LoadEnv(".env")
	mux := imux.NewRouter()
	// 添加全局中间件
	mux.Use(mws.Logger)

	mux.Get("/ping", api.Ping)
	// 获取 < /pong?foo=FOO&bar=BAR > 查询参数
	mux.Get("/pong", api.Pong)

	// 路由分组
	v1 := mux.Group("/v1", mws.Auth)
	v1.Get("/foo/:id", api.FooId)
	v1.Post("/foo", api.Foo)

	// 路由分组: restful api 接口
	v2 := mux.Group("/v2")
	v2.Get("/hello", api.HelloGet)
	v2.Post("/hello", api.HelloPost)
	v2.Delete("/hello", api.HelloDelete)
	v2.Put("/hello", api.HelloPut)

	// 启动服务
	address := os.Getenv("app.address")
	server := http.Server{Addr: address, Handler: mux}
	fmt.Printf("App Run [%s]\n", address)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

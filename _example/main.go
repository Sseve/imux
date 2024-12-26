package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Sseve/imux"
	mapi "github.com/Sseve/imux/_example/api"
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

	mux.Get("/ping", mapi.Ping)
	// 获取 < /pong?foo=FOO&bar=BAR > 查询参数
	mux.Get("/pong", mapi.Pong)

	// 路由分组
	api := mux.Group("/api", mws.Auth)
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

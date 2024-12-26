package api

import (
	"net/http"

	"github.com/Sseve/imux"
)

// Ping 全局路由控制器示例
func Ping(w http.ResponseWriter, _ *http.Request) {
	imux.Success(w, imux.Map{"code": 200, "message": "pong"})
}

// Pong 查询参数接口示例 < /pong?foo=FOO&bar=BAR >
func Pong(w http.ResponseWriter, r *http.Request) {
	foo := r.URL.Query().Get("foo")
	bar := r.URL.Query().Get("bar")
	imux.Success(w, imux.Map{"code": 200, "foo": foo, "bar": bar})
}

// FooId 路径参数接口示例 </api/hello/:id>
func FooId(w http.ResponseWriter, r *http.Request) {
	id := imux.Param(r, "id")
	imux.Success(w, imux.Map{"code": 200, "message": "Get foo id: " + id})
}

// Foo 请求参数绑定接口示例
func Foo(w http.ResponseWriter, r *http.Request) {
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

// HelloGet GET restful api 接口示例
func HelloGet(w http.ResponseWriter, r *http.Request) {
	imux.Success(w, imux.Map{"code": 200, "message": "Hello " + r.Method})
}

// HelloPost POST restful api 接口示例
func HelloPost(w http.ResponseWriter, r *http.Request) {
	imux.Success(w, imux.Map{"code": 200, "message": "Hello " + r.Method})
}

// HelloDelete DELETE restful api 接口示例
func HelloDelete(w http.ResponseWriter, r *http.Request) {
	imux.Success(w, imux.Map{"code": 200, "message": "Hello " + r.Method})
}

// HelloPut PUT restful api 接口示例
func HelloPut(w http.ResponseWriter, r *http.Request) {
	imux.Success(w, imux.Map{"code": 200, "message": "Hello " + r.Method})
}

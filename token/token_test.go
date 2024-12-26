package token

import (
	"github.com/Sseve/imux/env"
	"log"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	// 切换工作目录到项目根目录
	err := os.Chdir("../_example")
	if err != nil {
		log.Fatalf("Failed to change directory: %v", err)
	}
	env.LoadEnv(".env")
	// t.Skip()
	token, err := Create("zhangsan")
	if err != nil {
		t.Errorf("create token error: %s", err.Error())
	} else {
		t.Logf("create token: %s", token)
	}
}

func TestParse(t *testing.T) {
	payLoad, err := Parse("eyJOYW1lIjoiemhhbmdzYW4iLCJFeHBpcmVzQXQiOi01NjIwMjgwNzgyfQ==.gFEfm2F9rgQfu67rePhoKW_xqvx4-JRUHXafQiCnIzM=")
	if err != nil {
		t.Errorf("Parse token error: %s", err)
	} else {
		t.Log("parse token: ", payLoad.Name)
	}
}

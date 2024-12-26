package pwd

import "testing"

func TestHash(t *testing.T) {
	pwd := Hash("123456")
	t.Log("Hash Pwd: ", pwd)
}

func TestVerify(t *testing.T) {
	ok, err := Verify("$jZae727K08KaOmKSgOaGzww_XVqGr_PKEgIMkjrcbJI=", "123456")
	if err != nil {
		t.Errorf("Verify Error: %s", err.Error())
	}
	if ok == true {
		t.Log("Verify pass success")
	} else {
		t.Log("Verify pass failure")
	}

}

// 运行指定的测试函数:
// go test -v -run "Test(Hash|Verify)"

// 运行所有测试(根目录下执行):
// go test -v ./...

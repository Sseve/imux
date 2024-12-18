package imux

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func Bind(r *http.Request, v any) error {
	if r.Body == nil {
		return errors.New("request body is empty")
	}
	defer r.Body.Close()

	// 解码 JSON 数据
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	if err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}
	return nil
}

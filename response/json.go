package response

import (
	"encoding/json"
	"net/http"
)

const (
	Success = iota // 成功
	PartSuccess // 部分成功
	Fail // 失败
)

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Index   int `json:"index"`
	Code    int `json:"code"`
}

func RenderJson(w http.ResponseWriter, statusCode int, msg string, errors []Error, datas ...interface{}) {
	jsonHeader(w)
	resp := map[string]interface{}{
		"status":  statusCode,
	}
	if errors != nil && len(errors) > 0 {
		resp["errors"] = errors
	}
	if msg != nil && len(msg) > 0 {
		resp["reason"] = msg
	}
	// 是否有其他参数
	if len(datas) > 0 && len(datas) % 2 == 0 {
		for i := 0; i < len(datas); i += 2 {
			k, v := datas[i], datas[i + 1]
			if s, ok := k.(string); ok {
				resp[s] = v
			}
		}
	}
	// 将map转成json字节输出
	b, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte(`{"status":500,"message":"error parsing json with extra datas"}`))
	} else {
		w.Write(b)
	}
}

// 成功响应
func JsonOk(w http.ResponseWriter, datas ...interface{}) {
	RenderJson(w, Success, nil, datas...)
}

// 错误响应
func JsonErr(w http.ResponseWriter, errors []Error, datas ...interface{}) {
	RenderJson(w, Fail, errors, datas...)
}

// 错误响应
func JsonPartSuccess(w http.ResponseWriter, errors []Error, datas ...interface{}) {
	RenderJson(w, PartSuccess, errors, datas...)
}

// json响应头设置
func jsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

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


func RenderJson(w http.ResponseWriter, statusCode int, errors interface{}, datas ...interface{}) {
	jsonHeader(w)
	resp := map[string]interface{}{
		"status":  statusCode,
	}
	if errors != nil {
		resp["errors"] = errors
	}
	// 是否有其他参数
	if len(datas) > 0 && len(datas)%2 == 0 {
		for i := 0; i < len(datas); i += 2 {
			k, v := datas[i], datas[i+1]
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
func JsonErr(w http.ResponseWriter, errors interface{}, datas ...interface{}) {
	RenderJson(w, Fail, errors, datas...)
}

// 错误响应
func JsonPartSuccess(w http.ResponseWriter, errors interface{}, datas ...interface{}) {
	RenderJson(w, PartSuccess, errors, datas...)
}

//// 500响应，带错误提示
//func JsonErrWithMsg(w http.ResponseWriter, msg string, datas ...interface{}) {
//	RenderJson(w, http.StatusInternalServerError, msg, datas...)
//}
//
//// 400响应
//func JsonBadRequest(w http.ResponseWriter, datas ...interface{}) {
//	RenderJson(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), datas...)
//}
//
//// 400响应，带错误提示
//func JsonBadRequestWithMsg(w http.ResponseWriter, msg string, datas ...interface{}) {
//	RenderJson(w, http.StatusBadRequest, msg, datas...)
//}

// json响应头设置
func jsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

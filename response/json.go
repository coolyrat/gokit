package response

import (
	"encoding/json"
	"net/http"
)

func RenderJson(w http.ResponseWriter, statusCode int, msg string, datas ...interface{}) {
	jsonHeader(w)
	resp := map[string]interface{}{
		"status":  statusCode,
		"message": msg,
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

// 200响应
func JsonOk(w http.ResponseWriter, datas ...interface{}) {
	RenderJson(w, http.StatusOK, http.StatusText(http.StatusOK), datas...)
}

// 500响应
func JsonErr(w http.ResponseWriter, datas ...interface{}) {
	RenderJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), datas...)
}

// 500响应，带错误提示
func JsonErrWithMsg(w http.ResponseWriter, msg string, datas ...interface{}) {
	RenderJson(w, http.StatusInternalServerError, msg, datas...)
}

// 400响应
func JsonBadRequest(w http.ResponseWriter, datas ...interface{}) {
	RenderJson(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), datas...)
}

// 400响应，带错误提示
func JsonBadRequestWithMsg(w http.ResponseWriter, msg string, datas ...interface{}) {
	RenderJson(w, http.StatusBadRequest, msg, datas...)
}

// json响应头设置
func jsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// 测试

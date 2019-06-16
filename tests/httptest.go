package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http/httptest"
	"strings"
)

type Response struct {
	Retcode int    `json:"retcode"`
	Msg     string `json:"msg"`
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func Get(uri string, engine *gin.Engine) []byte {
	// 构造GET请求
	r := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, r)

	// 获取响应
	result := w.Result()
	defer result.Body.Close()
	resp, _ := ioutil.ReadAll(result.Body)
	return resp
}

func PostJson(uri string, params map[string]interface{}, engine *gin.Engine) []byte {
	// 构造POST请求
	body, _ := json.Marshal(params)
	r := httptest.NewRequest("POST", uri, bytes.NewReader(body))
	r.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, r)

	// 获取响应
	result := w.Result()
	defer result.Body.Close()
	resp, _ := ioutil.ReadAll(result.Body)
	return resp
}

func PostForm(uri string, params map[string]string, engine *gin.Engine) []byte {
	// 构造post form的参数
	var s []string
	for k, v := range params {
		s = append(s, k+"="+v)
	}
	body := strings.Join(s, "&")
	// 构造请求
	r := httptest.NewRequest("POST", uri, bytes.NewReader([]byte(body)))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, r)

	// 获取响应
	result := w.Result()
	defer result.Body.Close()
	resp, _ := ioutil.ReadAll(result.Body)
	return resp
}

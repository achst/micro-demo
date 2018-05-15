package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/hopehook/micro-demo/api-gateway/g"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/julienschmidt/httprouter"
)

// Raw handler
func Raw(h http.HandlerFunc) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		h(w, r)
		return
	})
}

// WriteString write raw string
func WriteString(w http.ResponseWriter, r *http.Request, result string) {
	Log(r, result)
	w.Write([]byte(result))
}

// WriteBytes write raw bytes
func WriteBytes(w http.ResponseWriter, r *http.Request, result []byte) {
	Log(r, string(result))
	w.Write(result)
}

// CommonWrite write common data response
func CommonWrite(w http.ResponseWriter, r *http.Request, code int, msg string, data interface{}) {
	resultMap := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	result, _ := json.Marshal(resultMap)
	WriteBytes(w, r, result)
}

// CommonWriteSuccess write common data response
func CommonWriteSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
	resultMap := map[string]interface{}{
		"code": 0,
		"msg":  "",
		"data": data,
	}
	result, _ := json.Marshal(resultMap)
	WriteBytes(w, r, result)
}

// Log prints handler info of Request and Response
func Log(r *http.Request, result string) {
	optMap := map[string]interface{}{
		"ip":       r.RemoteAddr,
		"method":   r.Method,
		"path":     r.URL.Path,
		"datetime": time.Now().Format("2006-01-02 15:04:05"),
		"req_data": "",
		"res_data": result,
	}
	g.Logger.Debug(result)
	// req_data
	r.ParseForm()
	optMap["req_data"] = r.Form
	// res_data
	js, err := simplejson.NewJson([]byte(result))
	if err == nil {
		optMap["res_data"] = js.Interface()
	}
	optData, _ := json.Marshal(optMap)
	g.Logger.Info(string(optData))
}

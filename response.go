//  Copyright(C) 2024 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2024-11-17

package webr

import (
	"encoding/json"
	"net/http"

	"github.com/xanygo/anygo/xerror"
)

type Response struct {
	Code   int
	Data   any `json:"Data,omitempty"`
	Msg    string
	Jump   string `json:"Jump,omitempty"`
	Reload bool   `json:"Reload,omitempty"`
}

func (r Response) WriteJSON(w http.ResponseWriter) {
	r.WriteJSONStatus(w, http.StatusOK)
}

func (r Response) WriteJSONStatus(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	bf, _ := json.Marshal(r)
	_, _ = w.Write(bf)
}

// WriteError 以 HTTP Status 200 输出错误
func (r Response) WriteError(w http.ResponseWriter, err error) {
	code := xerror.ErrCode(err, 500)
	r.Code = int(code)
	r.Msg = err.Error()
	r.WriteJSON(w)
}

// WriteErrorStatus 以指定的 HTTP Status 输出错误
func (r Response) WriteErrorStatus(w http.ResponseWriter, err error, status int) {
	code := xerror.ErrCode(err, 500)
	r.Code = int(code)
	r.Msg = err.Error()
	r.WriteJSONStatus(w, status)
}

func (r Response) WriteAutoError(w http.ResponseWriter, err error) {
	if err != nil {
		r.WriteError(w, err)
		return
	}
	r.WriteJSON(w)
}
func (r Response) WriteAutoErrorStatus(w http.ResponseWriter, err error, failStatus int) {
	if err != nil {
		r.WriteErrorStatus(w, err, failStatus)
		return
	}
	r.WriteJSON(w)
}

func WriteJSON(w http.ResponseWriter, code int, msg string, data any) {
	resp := Response{Code: code, Data: data, Msg: msg}
	resp.WriteJSON(w)
}

// WriteJSONError 以 HTTP Status 200 输出错误
func WriteJSONError(w http.ResponseWriter, err error) {
	resp := Response{}
	resp.WriteError(w, err)
}

// WriteJSONErrorStatus 以指定的 HTTP Status 输出错误
func WriteJSONErrorStatus(w http.ResponseWriter, err error, status int) {
	resp := Response{}
	resp.WriteErrorStatus(w, err, status)
}

// WriteJSONAuto 以 HTTP Status 200 输出结果
func WriteJSONAuto(w http.ResponseWriter, err error) {
	resp := Response{}
	resp.WriteAutoError(w, err)
}

// WriteJSONAutoStatus 若 err 为 nil，以 HTTP Status 200 输出 JSON 响应，否则以 failStatus 输出 JSON
func WriteJSONAutoStatus(w http.ResponseWriter, err error, failStatus int) {
	resp := Response{}
	resp.WriteAutoErrorStatus(w, err, failStatus)
}

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

func (r Response) WriteError(w http.ResponseWriter, err error) {
	code := xerror.ErrCode2(err, 500)
	r.Code = int(code)
	r.Msg = "失败：" + err.Error()
	r.WriteJSON(w)
}

func (r Response) WriteErrorAuto(w http.ResponseWriter, err error) {
	if err != nil {
		r.WriteError(w, err)
		return
	}
	r.WriteJSON(w)
}

func WriteJSON(w http.ResponseWriter, code int, msg string, data any) {
	resp := Response{Code: code, Data: data, Msg: msg}
	resp.WriteJSON(w)
}

func WriteJSONError(w http.ResponseWriter, err error) {
	resp := Response{}
	resp.WriteError(w, err)
}

func WriteJSONAuto(w http.ResponseWriter, err error) {
	resp := Response{}
	resp.WriteErrorAuto(w, err)
}

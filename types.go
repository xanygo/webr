//  Copyright(C) 2024 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2024-11-17

package webr

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code int
	Data any
	Msg  string
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

func WriteJSON(w http.ResponseWriter, code int, msg string, data any) {
	resp := Response{Code: code, Data: data, Msg: msg}
	resp.WriteJSON(w)
}

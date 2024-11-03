//  Copyright(C) 2024 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2024-11-03

package webr

import (
	"net/http/httptest"
	"testing"
)

func TestHandler_ServeHTTP(t *testing.T) {
	ts := httptest.NewServer(&Handler{})
	resp, err := ts.Client().Get(ts.URL + "/jquery.min.js")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("want status 200, got %d", resp.StatusCode)
	}
}

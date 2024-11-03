//  Copyright(C) 2024 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2024-11-03

package webr

import (
	"net/http/httptest"
	"path"
	"testing"
)

func TestHandler_ServeHTTP(t *testing.T) {
	ts := httptest.NewServer(&Handler{})
	check := func(t *testing.T, name string) {
		t.Run(name, func(t *testing.T) {
			name = path.Clean(name)
			resp, err := ts.Client().Get(ts.URL + name)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				t.Fatalf("want status 200, got %d", resp.StatusCode)
			}
		})
	}
	check(t, "/jquery.min.js")
	check(t, "/jquery/jquery-3.7.1.min.js")
	for k := range alias {
		check(t, "/"+k)
	}
}

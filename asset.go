//  Copyright(C) 2024 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2024-11-03

package webr

import (
	"crypto/md5"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

//go:embed asset
var resource embed.FS

var _ http.Handler = (*Handler)(nil)

type Handler struct{}

var fsSystem = httpFS{
	FS: http.FS(resource),
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fileName := fileRealPath(req.URL.Path)
	file, err := fsSystem.Open(fileName)
	if err != nil {
		http.NotFound(w, req)
		return
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		http.NotFound(w, req)
		return
	}

	if etag := fileMd5ETag(fileName); etag != "" {
		w.Header().Set("ETag", etag)
		if match := req.Header.Get("If-None-Match"); match == etag {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	http.ServeContent(w, req, info.Name(), info.ModTime(), file)
}

var _ http.FileSystem = (*httpFS)(nil)

type httpFS struct {
	FS http.FileSystem
}

func (h httpFS) Open(name string) (http.File, error) {
	if strings.HasPrefix(name, "/") {
		name = "/asset" + name
	} else {
		name = "/asset/" + name
	}

	return h.FS.Open(name)
}

var md5etag sync.Map

func fileMd5ETag(fp string) string {
	if val, ok := md5etag.Load(fp); ok {
		return val.(string)
	}
	file, err := fsSystem.Open(fp)
	if err != nil {
		return ""
	}
	defer file.Close()
	m := md5.New()
	if _, err = io.Copy(m, file); err != nil {
		return ""
	}
	str := fmt.Sprintf("%x", m.Sum(nil))
	result := fmt.Sprintf("%q", str)
	md5etag.Store(fp, result)
	return result
}

var alias = map[string]string{}

func init() {
	bf, err := resource.ReadFile("asset/alias.json")
	if err == nil {
		err = json.Unmarshal(bf, &alias)
	}
	if err != nil {
		panic(err)
	}
}

func fileRealPath(fp string) string {
	if name, ok := alias[fp]; ok {
		return name
	}
	return fp
}

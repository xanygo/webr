//  Copyright(C) 2024 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2024-11-04

package webr

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"sync"
)

//go:embed asset/asset.zip
var zipFile []byte

func Handler() http.Handler {
	return http.HandlerFunc(zipHandler)
}

func zipHandler(w http.ResponseWriter, req *http.Request) {
	fileName := req.URL.Path
	rd, err := zip.NewReader(bytes.NewReader(zipFile), int64(len(zipFile)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if etag := fileMd5ETag(rd, fileName); etag != "" {
		w.Header().Set("ETag", etag)
		if match := req.Header.Get("If-None-Match"); match == etag {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	http.ServeFileFS(w, req, rd, fileName)
}

var md5etag sync.Map

func fileMd5ETag(fsSystem fs.FS, fp string) string {
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
	str := hex.EncodeToString(m.Sum(nil))
	result := fmt.Sprintf("%q", str)
	md5etag.Store(fp, result)
	return result
}

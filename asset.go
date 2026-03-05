//  Copyright(C) 2026 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2026-03-04

package webr

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"fmt"
	"io/fs"

	"github.com/xanygo/anygo/ds/xsync"
)

func loadZipFile(name string, bf []byte) *xsync.OnceInit[fs.FS] {
	return &xsync.OnceInit[fs.FS]{
		New: func() fs.FS {
			rd, err := zip.NewReader(bytes.NewReader(bf), int64(len(bf)))
			if err != nil {
				panic(fmt.Errorf("unzip %s: %w", name, err))
			}
			return rd
		},
	}
}

//go:embed asset/jquery.zip
var jqueryFile []byte

var jqueryFileOnce = loadZipFile("jquery.zip", jqueryFile)

func JQuery() fs.FS {
	return jqueryFileOnce.Load()
}

//go:embed asset/axios.zip
var axiosFile []byte
var axiosFileOnce = loadZipFile("axios.zip", axiosFile)

func Axios() fs.FS {
	return axiosFileOnce.Load()
}

//go:embed asset/bootstrap.zip
var bootstrapFile []byte
var bootstrapFileOnce = loadZipFile("bootstrap.zip", bootstrapFile)

func Bootstrap() fs.FS {
	return bootstrapFileOnce.Load()
}

//go:embed asset/clipboard.zip
var clipboardFile []byte
var clipboardFileOnce = loadZipFile("clipboard.zip", clipboardFile)

func Clipboard() fs.FS {
	return clipboardFileOnce.Load()
}

//go:embed asset/icons.zip
var iconsFile []byte
var iconsFileOnce = loadZipFile("icons.zip", iconsFile)

func Icons() fs.FS {
	return iconsFileOnce.Load()
}

//go:embed asset/sortable.zip
var sortableFile []byte
var sortableFileOnce = loadZipFile("sortable.zip", sortableFile)

func Sortable() fs.FS {
	return sortableFileOnce.Load()
}

//go:embed asset/wangeditor.zip
var wangEditorFile []byte
var wangEditorFileOnce = loadZipFile("wangeditor.zip", wangEditorFile)

func Wangeditor() fs.FS {
	return wangEditorFileOnce.Load()
}

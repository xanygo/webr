//  Copyright(C) 2024 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2024-11-04

package webr

import (
	_ "embed"

	"github.com/xanygo/anygo/xhttp"
)

//go:embed asset/asset.zip
var zipFile []byte

func Handler() *xhttp.ZipFile {
	hd := &xhttp.ZipFile{
		Content: zipFile,
	}
	if err := hd.Init(); err != nil {
		panic(err)
	}
	return hd
}

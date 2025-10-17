//  Copyright(C) 2024 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2024-11-04

//go:generate anygo-encrypt-zip -m js,css -token 08cc63c24f77fc1d4141663fcb9a9d26 -o ui.ez -go ui_ez.go -var uiAsset -tags "" ui

package webr

import (
	_ "embed"

	"github.com/xanygo/anygo/ds/xsync"
	"github.com/xanygo/anygo/xhttp"
)

//go:embed asset/asset.zip
var zipFile []byte

var assetHandler = xsync.OnceInit[xhttp.FSHandler]{
	New: func() xhttp.FSHandler {
		hd := &xhttp.ZipFile{
			Content: zipFile,
		}
		if err := hd.Init(); err != nil {
			panic(err)
		}
		return hd
	},
}

func Handler() xhttp.FSHandler {
	return xhttp.FSHandlers{
		assetHandler.Load(),
		&xhttp.FS{
			FS: uiAsset,
		},
	}
}

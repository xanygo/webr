//  Copyright(C) 2024 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2024-11-04

//go:generate anygo-encrypt-zip -token 08cc63c24f77fc1d4141663fcb9a9d26 -o ui.ez ui

package webr

import (
	_ "embed"
	"io/fs"

	"github.com/xanygo/anygo/ds/xsync"
	"github.com/xanygo/anygo/ds/xzip"
	"github.com/xanygo/anygo/xcodec"
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

//go:embed ui.ez
var uiAsset []byte

var uiFS = xsync.OnceInit[fs.FS]{
	New: func() fs.FS {
		dz := &xcodec.AesOFB{
			Key: "08cc63c24f77fc1d4141663fcb9a9d26",
		}
		return xzip.MustDecrypt(uiAsset, dz)
	},
}

func Handler() xhttp.FSHandler {
	return xhttp.FSHandlers{
		assetHandler.Load(),
		&xhttp.FS{
			FS: uiFS.Load(),
		},
	}
}

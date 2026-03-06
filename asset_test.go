//  Copyright(C) 2026 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2026-03-04

package webr_test

import (
	"fmt"
	"io/fs"
	"testing"

	"github.com/xanygo/anygo/xt"

	"github.com/xanygo/webr"
)

func TestLoadAll(t *testing.T) {
	var list = []fs.FS{
		webr.Axios(),
		webr.Bootstrap(),
		webr.Axios(),
		webr.Clipboard(),
		webr.JQuery(),
		webr.Sortable(),
		webr.Wangeditor(),
		webr.UI(),
	}
	for idx, v := range list {
		t.Run(fmt.Sprintf("fs_%d", idx), func(t *testing.T) {
			var num int
			err := fs.WalkDir(v, ".", func(path string, d fs.DirEntry, err error) error {
				num++
				xt.NoError(t, err)
				return nil
			})
			xt.NoError(t, err)
			xt.Greater(t, num, 0)
		})
	}
}

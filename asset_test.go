//  Copyright(C) 2026 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2026-03-04

package webr_test

import (
	"testing"

	"github.com/xanygo/webr"
)

func TestLoadAll(t *testing.T) {
	webr.Axios()
	webr.Bootstrap()
	webr.Axios()
	webr.Clipboard()
	webr.JQuery()
	webr.Sortable()
	webr.Wangeditor()
	webr.UI()
}

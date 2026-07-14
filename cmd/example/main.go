//  Copyright(C) 2024 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2024-11-03

package main

import (
	"github.com/xanygo/anygo/xhttp"
	"github.com/xanygo/anygo/xio/xfs"
	"log"
	"net/http"

	"github.com/xanygo/webr"
)

const html = `<html>
<head>
<script type="text/javascript" src="/asset/bootstrap/bootstrap.js" defer ></script>
<link rel="stylesheet" href="/asset/bootstrap/bootstrap.css" />
</head
<body>
 <h1>Hello</h1>
<div>
	<button type="button" class="btn btn-primary">Primary</button>
</div>
</body>
</html>`

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(html))
	})

	ah := xhttp.FSHandlers{
		&xhttp.FS{
			FS: xfs.OverlayFS{
				webr.Bootstrap(),
				webr.Icons(),
				webr.JQuery(),
				webr.Clipboard(),
				webr.Sortable(),
				webr.UI(),
			},
		},
	}
	//fs.WalkDir(webr.UI(), ".", func(path string, d fs.DirEntry, err error) error {
	//	log.Println("path:", path)
	//	return nil
	//})
	xh := http.FileServerFS(ah)
	http.Handle("/asset/", http.StripPrefix("/asset/", xh))
	log.Println("running...")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	log.Println("server exit:", err)
}

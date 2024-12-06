//  Copyright(C) 2024 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2024-11-03

package main

import (
	"log"
	"net/http"

	"github.com/xanygo/webr"
)

const html = `<html>
<head>
<script type="text/javascript" src="/asset/bootstrap/bootstrap.bundle.min.js" defer ></script>
<link rel="stylesheet" href="/asset/bootstrap/bootstrap.min.css" />
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
	http.Handle("/asset/", http.StripPrefix("/asset/", webr.Handler()))
	log.Println("running...")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	log.Println("server exit:", err)
}

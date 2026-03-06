# Web Resource
内嵌打包部分常用 Web 资源，资源以 fs.FS 类型返回，可直接读取，或者用于 Web Server。

## 1. Include Files
```
.
axios
└── axios.js
bootstrap
├── bootstrap.css
└── bootstrap.js
clipboard
└── clipboard.js
jquery
├── form2json.js
└── jquery.js
sortable/
└── sortable.js
wangeditor
├── index.js
└── style.css
icons
├── font
│   ├── bootstrap-icons.css
│   ├── bootstrap-icons.json
│   └── fonts
│       ├── bootstrap-icons.woff
│       └── bootstrap-icons.woff2

```

## 2. Usage
```go
import "net/http"
import "github.com/xanygo/webr"

http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServerFS(webr.Bootstrap())))
```

```html
<script type="text/javascript" src="/asset/bootstrap/bootstrap.js"></script>
<link rel="stylesheet" href="/asset/bootstrap/bootstrap.css" />
```
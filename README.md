# Web Resource


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
├── form.js
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

```go
<script type="text/javascript" src="/asset/bootstrap/bootstrap.js"></script>
<link rel="stylesheet" href="/asset/bootstrap/bootstrap.css" />
```
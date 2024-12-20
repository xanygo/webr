# Web Resource


## 1. Include Files
```
.
├── axios
│   └── axios.min.js
├── bootstrap
│   ├── bootstrap.bundle.min.js
│   └── bootstrap.min.css
├── icons
│   ├── font
│   │   ├── bootstrap-icons.css
│   │   ├── bootstrap-icons.json
│   │   ├── bootstrap-icons.min.css
│   │   └── fonts
│   │       ├── bootstrap-icons.woff
│   │       └── bootstrap-icons.woff2
├── jquery
│   ├── jquery.form.min.js
│   ├── jquery.serializejson.min.js
│   └── jquery.min.js
├── vue3
│   └── vue.global.prod.js
└── wangeditor
    ├── index.min.js
    └── style.css
```

## 2. Usage
```go
import "github.com/xanygo/webr"

http.Handle("/asset/", http.StripPrefix("/asset/", webr.Handler()))
```

```go
<script type="text/javascript" src="/asset/bootstrap/bootstrap.bundle.min.js"></script>
<link rel="stylesheet" href="/asset/bootstrap/bootstrap.min.css" />
```
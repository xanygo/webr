# Web Resource


## 1. Include Files
```
.
├── bootstrap
│   ├── bootstrap-grid.min.css
│   ├── bootstrap-reboot.min.css
│   ├── bootstrap.bundle.min.js
│   └── bootstrap.min.css
├── bootstrap-icons
│   ├── font
│   │   ├── bootstrap-icons.css
│   │   ├── bootstrap-icons.json
│   │   ├── bootstrap-icons.min.css
│   │   ├── bootstrap-icons.scss
│   │   └── fonts
│   │       ├── bootstrap-icons.woff
│   │       └── bootstrap-icons.woff2
├── jquery
│   ├── jquery.form.min.js
│   └── jquery.min.js
└── vue3
    └── vue.global.min.js
```

## 2. Usage
```go
import "github.com/xanygo/webr"

http.Handle("/asset/", http.StripPrefix("/asset/", webr.Handler()))
```

```go
<script type="text/javascript" src="/asset/bootstrap/bootstrap.min.js"></script>
<link rel="stylesheet" href="/asset/bootstrap/bootstrap.min.css" />
```
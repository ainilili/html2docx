# Introduction
Convert [HTML](https://zh.wikipedia.org/wiki/HTML) to [DOCX](https://zh.wikipedia.org/wiki/Office_Open_XML)

# Install
```shell
go get github.com/ainilili/html2docx
```

# Usage
```go
package main

import (
	"github.com/ainilili/html2docx"
	"log"
	"os"
)

func main() {
	source := `<!DOCTYPE html>
<html>
<head>
	<title>我的网页</title>
</head>
<body>
	<h1>欢迎来到我的网页</h1>
	<p>这是一个简单的HTML文档。</p>
</body>
</html>`
	dist, err := html2docx.Convert([]byte(source))
	if err != nil {
		log.Fatalln(err)
	}
	err = os.WriteFile("test.docx", dist, 0666)
	if err != nil {
		panic(err)
	}
}
```
package main

import (
	"github.com/ainilili/html2docx/builder"
	"os"
)

func main() {
	htmlSource := []byte("<!DOCTYPE html><html><head><meta http-equiv=3D\"Content-Type\" content=3D\"text/html; charset=3DUTF-8\"><link rel=3D\"stylesheet\" type=3D\"text/css\" id=3D\"mce-u0\" crossorigin=3D\"anonymous\" referrerpolicy=3D\"origin\" href=3D\"https://cdn.tiny.cloud/1/no-api-key/tinymce/6.4.1-16/skins/ui/oxide/content.min.css\"><link rel=3D\"stylesheet\" type=3D\"text/css\" id=3D\"mce-u1\" crossorigin=3D\"anonymous\" referrerpolicy=3D\"origin\" href=3D\"https://cdn.tiny.cloud/1/no-api-key/tinymce/6.4.1-16/skins/content/default/content.min.css\"></head><body id=3D\"tinymce\" class=3D\"mce-content-body \" data-id=3D\"content\" aria-label=3D\"Rich Text Area. Press ALT-0 for help.\" contenteditable=3D\"true\" spellcheck=3D\"false\"><p style=3D\"text-align: center;\" data-mce-style=3D\"text-align: center;\"><b>fda</b></p><p style=3D\"text-align: right;\" data-mce-style=3D\"text-align: right;\"><b>fdaf</b></p></body></html>")
	bs, err := builder.Build(htmlSource)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("test.docx", bs, 0666)
	if err != nil {
		panic(err)
	}

	//reader.Print()
}

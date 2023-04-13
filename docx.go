package html2docx

import (
	"archive/zip"
	"bytes"
	"github.com/ainilili/html2docx/template"
	"unsafe"
)

type DocxFile struct {
	Name   string
	Source []byte
}

func string2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Convert(htmlSource []byte) ([]byte, error) {
	w := &bytes.Buffer{}
	zw := zip.NewWriter(w)

	data := map[string]interface{}{
		"htmlSource": bytes2string(htmlSource),
		"width":      12240,
		"height":     15840,
		"orient":     "portrait",
		"margins": map[string]interface{}{
			"top":    1440,
			"right":  1440,
			"bottom": 1440,
			"left":   1440,
			"header": 720,
			"footer": 720,
			"gutter": 0,
		},
	}
	documentXmlBytes, err := template.Parse(template.DocumentXmlBytes, data)
	if err != nil {
		return nil, err
	}
	documentMhtBytes, err := template.Parse(template.DocumentMhtBytes, data)
	if err != nil {
		return nil, err
	}

	files := make([]DocxFile, 0)
	files = append(files, DocxFile{Name: "[Content_Types].xml", Source: template.ContentTypesXmlBytes})
	files = append(files, DocxFile{Name: "_rels/.rels", Source: template.RelsXmlBytes})
	files = append(files, DocxFile{Name: "word/document.xml", Source: documentXmlBytes})
	files = append(files, DocxFile{Name: "word/afchunk.mht", Source: documentMhtBytes})
	files = append(files, DocxFile{Name: "word/_rels/document.xml.rels", Source: template.DocumentRelsXmlBytes})

	for _, file := range files {
		hw, err := zw.CreateHeader(&zip.FileHeader{
			Name:   file.Name,
			Method: zip.Store,
		})
		if err != nil {
			return nil, err
		}
		_, err = hw.Write(file.Source)
		if err != nil {
			return nil, err
		}
	}
	err = zw.Close()
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

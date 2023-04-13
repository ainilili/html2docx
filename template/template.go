package template

import (
	"bytes"
	_ "embed"
	"text/template"
)

var (
	//go:embed document.tpl
	DocumentXmlBytes []byte
	//go:embed mht_document.tpl
	DocumentMhtBytes []byte
	//go:embed content_types.xml
	ContentTypesXmlBytes []byte
	//go:embed document.xml.rels
	DocumentRelsXmlBytes []byte
	//go:embed rels.xml
	RelsXmlBytes []byte
)

func Parse(source []byte, data map[string]interface{}) ([]byte, error) {
	tpl, err := template.New("test").Parse(string(source))
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	err = tpl.Execute(buf, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

package tplutil

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/leonzhouwei/llsn/go/conf"
)

func RenderHtml(w http.ResponseWriter, fileNameNoSuffix string) error {
	baseName := fileNameNoSuffix + conf.ViewFileSuffix()
	filePath := filepath.Join(conf.PublicDirHome(), baseName)
	t, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}
	t.Execute(w, nil)
	return nil
}
package tplutil

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/leonzhouwei/gobrth/go/conf"
)

func RenderHtml(w http.ResponseWriter, config conf.Config, fileNameNoSuffix string) error {
	baseName := fileNameNoSuffix + config.ViewFileSuffix()
	filePath := filepath.Join(config.ViewFileHome(), baseName)
	t, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}
	t.Execute(w, nil)
	return nil
}
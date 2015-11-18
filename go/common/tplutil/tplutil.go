package tplutil

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/leonzhouwei/gobrth/go/conf"
)

func RenderHtml(w http.ResponseWriter, config conf.Config, fileNameNoSuffix string) error {
	return RawRenderHtml(w, config.ViewFileHome(), fileNameNoSuffix, config.ViewFileSuffix())
}

func RawRenderHtml(w http.ResponseWriter, viewFileHome string, fileNameNoSuffix string, fileNameSuffix string) error {
	baseName := fileNameNoSuffix + fileNameSuffix
	filePath := filepath.Join(viewFileHome, baseName)
	t, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}
	t.Execute(w, nil)
	return nil
}
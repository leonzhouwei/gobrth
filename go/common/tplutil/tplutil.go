package tplutil

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderHtml(w http.ResponseWriter, viewFileHome string, fileNameNoSuffix string, fileNameSuffix string) error {
	baseName := fileNameNoSuffix + fileNameSuffix
	filePath := filepath.Join(viewFileHome, baseName)
	t, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}
	t.Execute(w, nil)
	return nil
}
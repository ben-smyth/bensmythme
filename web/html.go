package web

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

var Templ = func() *template.Template {
	t := template.New("")
	err := filepath.Walk("web/templates", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			fmt.Println(path)
			_, err = t.ParseFiles(path)
			if err != nil {
				fmt.Println(err)
			}
		}
		return err
	})

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return t
}()

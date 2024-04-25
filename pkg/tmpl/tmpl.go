package tmpl

import (
	. "html/template"
	"io"
	"io/fs"
	"log/slog"
	"path/filepath"
)

var tmpl *Template

func LoadTemplates(path string) error {

    tmpl = New("")
    filepath.WalkDir(path, func(p string, f fs.DirEntry, err error) error {

        if filepath.Ext(f.Name()) != ".html" {
            return nil
        }

        slog.Info("loading template", "file", p)
        tmpl, err = tmpl.Parse(p)
        return nil
    })

    var err error
    tmpl, err = ParseGlob(path + "/*.html")
    if err != nil {
        return err
    }

    return nil
}

type Tmpl struct{}

func TmplHandle() *Tmpl {
    return &Tmpl{}
}

func (t *Tmpl) Execute(w io.Writer, name string, data any) error {
    return tmpl.ExecuteTemplate(w, name, data)
}

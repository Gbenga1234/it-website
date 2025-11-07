package render

import (
    "html/template"
    "os"
    "path/filepath"
)

type Templates struct {
    tpl *template.Template
}

func LoadTemplates(root string) (*Templates, error) {
    // Parse base, partials, and pages
    base := filepath.Join(root, "base.tmpl")
    partialsGlob := filepath.Join(root, "partials", "*.tmpl")
    pagesGlob := filepath.Join(root, "*.tmpl")

    // Parse all templates into one set
    t, err := template.New("").Funcs(template.FuncMap{
        "Year": func() int {
            return 2025
        },
    }).ParseFiles(base)
    if err != nil {
        return nil, err
    }

    // Parse partials
    matches, _ := filepath.Glob(partialsGlob)
    if len(matches) > 0 {
        if _, err := t.ParseFiles(matches...); err != nil {
            return nil, err
        }
    }

    // Parse pages (skip base to avoid duplicate parse)
    files, _ := filepath.Glob(pagesGlob)
    for _, f := range files {
        if filepath.Base(f) == "base.tmpl" {
            continue
        }
        // Parse named templates on the existing set
        b, err := os.ReadFile(f)
        if err != nil {
            return nil, err
        }
        if _, err := t.New(filepath.Base(f)).Parse(string(b)); err != nil {
            return nil, err
        }
    }

    return &Templates{tpl: t}, nil
}

func (t *Templates) Render(wr interface{ Write([]byte) (int, error) }, name string, data any) error {
    // Wrap `ExecuteTemplate` so handlers can call a simple method
    // Expect `name` to be the page template file name (e.g., "home.tmpl")
    return t.tpl.ExecuteTemplate(wr, "base.tmpl", struct {
        Name string
        Data any
    }{
        Name: name,
        Data: data,
    })
}

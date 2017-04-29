package main

import (
	"net/http"
	"log"
	"os"
	"html/template"
	"path/filepath"
	"fmt"
)

type Package struct {
	Name     string
	NumFuncs int
	NumVars  int
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd:%v", err)
	}
	log.Print("Work directory:", wd)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(filepath.Join(wd, "src//go_web_foundation//lecture2//code//main_v2.tmpl"))
		if err != nil {
			fmt.Fprint(w, "Parse:%v", err)
			return
		}
		err = tmpl.Execute(w, &Package{
			Name:     "go-web",
			NumFuncs: 12,
			NumVars:  1200,
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Print("Starting server...v1")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

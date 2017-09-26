package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

type pagedat struct {
	Title string
	Data  map[string]interface{}
	Page  string
}

func main() {
	tmpl, err := template.ParseFiles("site.tmpl")
	if err != nil {
		panic(err)
	}
	pglst := getPgs("pages")
	sort.Strings(pglst)
	for _, p := range getPgs("out") {
		if _, fn := filepath.Split(p); fn != ".gitignore" {
			rp, err := filepath.Rel("out", p)
			if err != nil {
				panic(err)
			}
			ps := filepath.Join("pages", rp)
			i := sort.SearchStrings(pglst, ps)
			e := i < len(pglst)
			if e {
				e = pglst[i] == ps
			}
			if !e {
				err = os.Remove(p)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	for _, p := range pglst {
		pdr, err := ioutil.ReadFile(p)
		if err != nil {
			panic(err)
		}
		pspl := strings.Split(string(pdr), "@---@")
		var pd pagedat
		err = yaml.Unmarshal([]byte(pspl[0]), &pd)
		if err != nil {
			panic(err)
		}
		pd.Page = pspl[1]
		rp, err := filepath.Rel("pages", p)
		if err != nil {
			panic(err)
		}
		outpath := filepath.Join("out", rp)
		dir, filename := filepath.Split(outpath)
		if pd.Title == "" {
			pd.Title = strings.TrimSuffix(filename, filepath.Ext(filename))
		}
		ptmpl := template.New("page")
		_, err = ptmpl.Parse(pd.Page)
		if err != nil {
			panic(err)
		}
		log.Println(pd)
		pb := bytes.NewBuffer(nil)
		ptmpl.Execute(pb, pd)
		pd.Page = pb.String()
		os.MkdirAll(dir, 0600) //dont check errors - we expect errors
		f, err := os.Create(outpath)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		err = tmpl.Execute(f, pd)
		if err != nil {
			panic(err)
		}
		log.Println(pd)
	}
}

func getPgs(dir string) (o []string) {
	o = []string{}
	d, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, f := range d {
		if f.IsDir() {
			o = append(o, getPgs(filepath.Join(dir, f.Name()))...)
		} else {
			o = append(o, filepath.Join(dir, f.Name()))
		}
	}
	return
}

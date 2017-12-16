package main

import (
  "fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
  mux := httprouter.New()
  mux.GET("/", index)
  mux.GET("/about", about)
  mux.GET("/apply", apply)
  mux.POST("/apply", applyProcess)
  mux.GET("/user/:name", user)
  mux.GET("/blog/:category/:article", blogRead)
  mux.POST("/blog/:category/:article", blogWrite)
  http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
  HandleError(w, err)
}

func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
  err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
  HandleError(w, err)
}

func apply(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
  err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
  HandleError(w, err)
}

func applyProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
  HandleError(w, err)
}
func user(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
  fmt.Fprintf(w, "USER, %s!\n", ps.ByName("name"))
}

func blogRead(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, "READ CATEGORY, %s!\n", ps.ByName("category"))
  fmt.Fprintf(w, "READ Article, %s!\n", ps.ByName("article"))
}

func blogWrite(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
  fmt.Fprintf(w, "WRITE CATEGORY, %s!\n", ps.ByName("category"))
  fmt.Fprintf(w, "WRITE Article, %s!\n", ps.ByName("article"))
}

func update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, "Update, %s!\n", ps.ByName("name"))
}

func trailing(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, "Trailing, %s!\n", ps.ByName("name"))
}

func HandleError(w http.ResponseWriter, err error){
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Fatalln(err)
  }
}

package main

import (
	"fmt"
	"log"
	"net/http"
  "html/template"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

var t *template.Template
func init() {
  t = template.Must(template.ParseGlob("html/*.html"))
  
}

func main() {
  http.HandleFunc("/", handleHome)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Test")
}

             
func handleHome(w http.ResponseWriter, r *http.Request) {
  db, err := sql.Open("mysql", "root:password@/aerothon")
  defer db.Close()
  
  results, err := db.Query("SELECT id, name FROM tags")
  if err != nil {
    log.Fatalln(err)
  }
  t.ExecuteTemplate(w, "home.html", nil)
}
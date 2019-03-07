package main

import (
	"fmt"
	"log"
	"net/http"
  "html/template"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type Homefeed struct {
  Headlines string
  Content string
}

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
  
  results, err := db.Query("SELECT headlines, content FROM homefeed")
  if err != nil {
    log.Fatalln(err)
  }
  var homeFeeds []Homefeed
  for results.Next() {
		var homeFeed Homefeed
		// for each row, scan the result into our tag composite object
		err = results.Scan(&homeFeed.Headlines, &homeFeed.Content)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
 
    homeFeeds = append(homeFeeds, homeFeed)
	}
  fmt.Println(homeFeeds)
  t.ExecuteTemplate(w, "home.html", homeFeeds)
}
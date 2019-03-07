package main

import (
	"fmt"
	"log"
	"net/http"
  "html/template"
  "database/sql"
  
  _ "io/ioutil"
  _ "github.com/go-sql-driver/mysql"
)

type Homefeed struct {
  Headlines string
  Content string
}

type FlightData struct {
 Program string
 Msn string
 Harness_len string
 Gross_weight string
 Atms_pre string
 Room_temp string
 Airport string
 Fuel_cap_r string
 Fuel_cap_l string
 Fuel_quant_l string
 Fuel_quant_r string
 Max_att string
 Flight_num string
}

var t *template.Template
func init() {
  t = template.Must(template.ParseGlob("html/*.html"))
  
}

func main() {
  
  http.HandleFunc("/", handleHome) //Handline the webrequest fro home feed
  http.HandleFunc("/flightdata/", handleFlightData) //Handline the webrequest fro home feed
  http.HandleFunc("/flightsearch/", handleFlightDataSearchShow) 
  http.HandleFunc("/flightshow/", handleShowFlight) 
  
	err := http.ListenAndServe(":3000", nil) //Starting the server on port 3000
	if err != nil {   //Error cheking for Webserver
		log.Fatalln(err)
	}
	fmt.Println("Test")
}

/*Function to handle home feed*/
func handleHome(w http.ResponseWriter, r *http.Request) {
  db, err := sql.Open("mysql", "root:password@/aerothon")
  defer db.Close()
  if err != nil {
    log.Fatalln(err)
  }
  
  results, err := db.Query("SELECT headlines, content FROM homefeed")
  if err != nil {
    log.Fatalln(err)
  }
  var homeFeeds []Homefeed
  for results.Next() {
		var homeFeed Homefeed
		err = results.Scan(&homeFeed.Headlines, &homeFeed.Content)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
    homeFeeds = append(homeFeeds, homeFeed)
	}
  fmt.Println(homeFeeds)
  t.ExecuteTemplate(w, "home.html", homeFeeds)
}

/*Function to handle Form screens and screen*/
func handleFlightData(w http.ResponseWriter, r *http.Request) {
  db, err := sql.Open("mysql", "root:password@/aerothon")
  defer db.Close()
  if err != nil {
    log.Fatalln(err)
  }
  
  p := r.FormValue("program")
  fmt.Println(p)
  if r.Method == "GET" {
    t.ExecuteTemplate(w, "flightdata.html", p)
  }
  
  if r.Method == "POST" {   
    err := r.ParseForm()
    if err != nil {
      panic(err)
    }
    qu, err := db.Prepare("INSERT INTO flightdata (program, msn, harness_len, gross_weight, atms_pre, room_temp, airport, fuel_cap_r, fuel_cap_l, fuel_quant_l, fuel_quant_r, max_att, flight_num) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)" )
    defer qu.Close()
    if err != nil {
      log.Fatalln(err)
    }
    
    qu.Exec(r.Form.Get("pname"), r.Form.Get("msn"), r.Form.Get("lengtn"), 
                           r.Form.Get("weight"), r.Form.Get("atmp"), r.Form.Get("roomt"), r.Form.Get("airport"), r.Form.Get("fcl"), r.Form.Get("fcr"), 
                           r.Form.Get("fql"), r.Form.Get("fqr"), r.Form.Get("maxat"), r.Form.Get("flightno"))
    newUrl := "/flightshow?msn=" + r.Form.Get("msn")
    http.Redirect(w, r, newUrl, http.StatusSeeOther)
  }
}


func handleFlightDataSearchShow(w http.ResponseWriter, r *http.Request) {
  db, err := sql.Open("mysql", "root:password@/aerothon")
  defer db.Close()
  if err != nil {
    log.Fatalln(err)
  }
  if r.Method == "GET" {   
  t.ExecuteTemplate(w, "flightsearch.html", nil)
  }
  if r.Method == "POST" {
    var flightdata FlightData
    flightdata.Msn = r.FormValue("msn")
    flightdata.Harness_len = r.FormValue("length")
    flightdata.Gross_weight = r.FormValue("weight")
    flightdata.Atms_pre = r.FormValue("atmp")
    flightdata.Room_temp = r.FormValue("roomt")
    flightdata.Airport = r.FormValue("airport")
    flightdata.Fuel_cap_r = r.FormValue("fcr")
    flightdata.Fuel_cap_l = r.FormValue("fcl")
    
    flightdata.Max_att = r.FormValue("maxat")
    flightdata.Flight_num = r.FormValue("flightno")
    
    if flightdata.Msn != "" {
      results, err := db.Query("SELECT msn FROM flightdata where msn = ?", flightdata.Msn)
      if err != nil {
        log.Fatalln(err)
      }
      results.Next()
      var msn string
      results.Scan(&msn)
      fmt.Println(msn)
      newUrl := "/flightshow?msn=" + msn
      http.Redirect(w, r, newUrl, http.StatusSeeOther)
    }
    if flightdata.Harness_len != "" {
      results, err := db.Query("SELECT msn FROM flightdata where harness_len = ?", flightdata.Harness_len)
      if err != nil {
        log.Fatalln(err)
      }
      results.Next()
      var msn string
      results.Scan(&msn)
      fmt.Println(msn)
    }
    if flightdata.Gross_weight != "" {
      results, err := db.Query("SELECT msn FROM flightdata where gross_weight = ?",flightdata.Gross_weight)
      if err != nil {
        log.Fatalln(err)
      }
      results.Next()
      var msn string
      results.Scan(&msn)
      fmt.Println(msn)
      newUrl := "/flightshow?msn=" + msn
      http.Redirect(w, r, newUrl, http.StatusSeeOther)
    }
    if flightdata.Atms_pre != "" {
      results, err := db.Query("SELECT msn FROM flightdata where Atms_pre = ?", flightdata.Atms_pre)
      if err != nil {
        log.Fatalln(err)
      }
      results.Next()
      var msn string
      results.Scan(&msn)
      fmt.Println(msn)
      newUrl := "/flightshow?msn=" + msn
      http.Redirect(w, r, newUrl, http.StatusSeeOther)
    }
    if flightdata.Room_temp != "" {
      results, err := db.Query("SELECT msn FROM flightdata where Room_temp = ?",flightdata.Room_temp)
      if err != nil {
        log.Fatalln(err)
      }
      results.Next()
      var msn string
      results.Scan(&msn)
      fmt.Println(msn)
      newUrl := "/flightshow?msn=" + msn
      http.Redirect(w, r, newUrl, http.StatusSeeOther)
    }
    if flightdata.Airport != "" {
      results, err := db.Query("SELECT msn FROM flightdata where Airport = ?", flightdata.Airport)
      if err != nil {
        log.Fatalln(err)
      }
      results.Next()
      var msn string
      results.Scan(&msn)
      fmt.Println(msn)
      newUrl := "/flightshow?msn=" + msn
      http.Redirect(w, r, newUrl, http.StatusSeeOther)
    }
    if flightdata.Fuel_cap_r != "" {
      results, err := db.Query("SELECT msn FROM flightdata where Fuel_cap_r = ?", flightdata.Fuel_cap_r)
      if err != nil {
        log.Fatalln(err)
      }
      results.Next()
      var msn string
      results.Scan(&msn)
      fmt.Println(msn)
      newUrl := "/flightshow?msn=" + msn
      http.Redirect(w, r, newUrl, http.StatusSeeOther)
    }
    if flightdata.Fuel_cap_l != "" {
      results, err := db.Query("SELECT msn FROM flightdata where Fuel_cap_l = ?", flightdata.Fuel_cap_l)
      if err != nil {
        log.Fatalln(err)
      }
      results.Next()
      var msn string
      results.Scan(&msn)
      fmt.Println(msn)
      newUrl := "/flightshow?msn=" + msn
      http.Redirect(w, r, newUrl, http.StatusSeeOther)
    }
    
    if flightdata.Max_att != "" {
      results, err := db.Query("SELECT msn FROM flightdata where Max_att = ?",  flightdata.Max_att)
      if err != nil {
        log.Fatalln(err)
      }
      results.Next()
      var msn string
      results.Scan(&msn)
      fmt.Println(msn)
      newUrl := "/flightshow?msn=" + msn
      http.Redirect(w, r, newUrl, http.StatusSeeOther)
    }
    
    if flightdata.Flight_num != "" {
      results, err := db.Query("SELECT msn FROM flightdata where Flight_num = ?", flightdata.Flight_num)
      if err != nil {
        log.Fatalln(err)
      }
      results.Next()
      var msn string
      results.Scan(&msn)
      fmt.Println(msn)
      newUrl := "/flightshow?msn=" + msn
      http.Redirect(w, r, newUrl, http.StatusSeeOther)
    }
  }
}

func handleShowFlight(w http.ResponseWriter, r *http.Request) {
  db, err := sql.Open("mysql", "root:password@/aerothon")
  defer db.Close()
  if err != nil {
    log.Fatalln(err)
  }
  msn := r.FormValue("msn")
  
  results, err := db.Query("SELECT program, msn, harness_len, gross_weight, atms_pre, room_temp, airport, fuel_cap_r, fuel_cap_l, fuel_quant_l, fuel_quant_r, max_att, flight_num FROM flightdata where msn = ?", msn)
  if err != nil {
    log.Fatalln(err)
  }
  results.Next()
  var flightdata FlightData
  err = results.Scan(&flightdata.Program, &flightdata.Msn,   &flightdata.Harness_len, &flightdata.Gross_weight, 
                    &flightdata.Atms_pre, &flightdata.Room_temp, &flightdata.Airport, &flightdata.Fuel_cap_r, &flightdata.Fuel_cap_l,
                    &flightdata.Fuel_quant_l, &flightdata.Fuel_quant_r, &flightdata.Max_att, &flightdata.Flight_num)
  if err != nil {
    panic(err.Error()) // proper error handling instead of panic in your app
  }
  t.ExecuteTemplate(w, "flightshow.html", flightdata)
  
}
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func displayDate(w http.ResponseWriter, req *http.Request) {
	current_time := time.Now()
	fmt.Fprintf(w, current_time.Format("15h04"))
}

func addEntries(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	switch req.Method {
	case http.MethodPost:

		var author string
		var entry string
		var line = ""

		if err := req.ParseForm(); err != nil {
			fmt.Println("Something went bad")
			fmt.Fprintln(w, "Something went bad")
			return
		}
		for key, value := range req.PostForm {

			switch key{
			case "author":
				author = value[0]
			case "entry":
				entry = value[0]
			}
		}
		fmt.Println(w, author, ":", entry)

		line += entry + "\n"

			f, err := os.OpenFile("file.txt",
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(line); err != nil {
				log.Println(err)
			}
		fmt.Fprintf(w, "Information received: %v\n", req.PostForm)}


}
func getEntries (w http.ResponseWriter, req *http.Request){
    body, err := ioutil.ReadFile("file.txt")
    if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }

   fmt.Fprintln(w, string(body) )
}

func main() {
	http.HandleFunc("/", displayDate)
	http.HandleFunc("/entries", getEntries)
	http.HandleFunc("/hello", addEntries)
	http.ListenAndServe(":4567", nil)
   }

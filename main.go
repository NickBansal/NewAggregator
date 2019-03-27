package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const newsAPI = "https://newsapi.org/v2/everything?q=bitcoin&from=2019-02-27&sortBy=publishedAt&apiKey=f454de783a154d599cb6768caf6284fa"

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1><a href='/index'>Home<a></h1>")
}

func main() {
	resp, _ := http.Get(newsAPI)
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)
	resp.Body.Close()
	fmt.Println(stringBody)
}

package main

import (
	"os"
	"bufio"
	"strings"
	"net/http"
	"encoding/json"
	
	"github.com/samuelbalogh/levenshtein"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

var dict = getDictionary()

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getDictionary() []string {
	file, err := os.Open("20k.txt")
	check(err)
	defer file.Close()
        scanner := bufio.NewScanner(file)

	var wordList []string

    	for scanner.Scan() {
        	wordList = append(wordList, scanner.Text())
    	}
	return wordList
}

func findLikelySpellings(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")	

	words := strings.Fields(text)
	
	results := []string{}

	for _, word := range words {
		likelySpelling := levenshtein.GetLikelySpelling(word, dict)
		results = append(results, likelySpelling)
	}

	json.NewEncoder(w).Encode(results)

}

func main() {
        router := mux.NewRouter()
	router.HandleFunc("/check", findLikelySpellings).Methods("POST")
        http.ListenAndServe(":8001", handlers.CORS()(router))
}

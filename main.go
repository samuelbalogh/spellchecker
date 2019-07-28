package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"net/http"
	"encoding/json"
	"spellchecker"
	
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


func getLikelySpelling(word string, dictionary []string) (string) {
	var mostLikelySpelling string
	var distance int
	var mtx  [][]int

	distance = len(word)

	for _, dictWord := range dictionary {
		levMatrix := getLevMatrix(word, dictWord)

		distanceToCurrentWord := getDistance(levMatrix, word, dictWord)
		if distanceToCurrentWord < distance {
			distance = distanceToCurrentWord
			mostLikelySpelling = dictWord
			mtx = levMatrix
		}
		
	} 
	if len(mostLikelySpelling) > 0 {
		printLevMatrix(word, mostLikelySpelling, mtx) 
		fmt.Println(mostLikelySpelling)
	} else {
		mostLikelySpelling = word
	}

	return mostLikelySpelling
}


func findLikelySpellings(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")	

	words := strings.Fields(text)
	
	results := []string{}

	for _, word := range words {
		likelySpelling := getLikelySpelling(word, dict)
		results = append(results, likelySpelling)
	}

	json.NewEncoder(w).Encode(results)

}

func main() {
        router := mux.NewRouter()
	router.HandleFunc("/check", findLikelySpellings).Methods("POST")
        http.ListenAndServe(":8001", handlers.CORS()(router))
}

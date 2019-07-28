package spellchecker

import (
	"os"
	"fmt"
	"bufio"
)


var dict = getDictionary()

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func Min(array [3]int) (int) {
    var min int = array[0]
    for _, value := range array {
        if min > value {
            min = value
        }
    }
    return min
}


func printLevMatrix(source string, target string, a [][]int) {
	fmt.Print("    ")
	for i, _ := range target {
		fmt.Printf(" %v", string(target[i]))
	}
	fmt.Println()
	fmt.Printf("  %v\n", a[0])
	for i, line := range a[1:] {
		fmt.Printf("%+v", string([]rune(source)[i]))
		fmt.Printf(" %v\n", line)
	}
}


func getLevMatrix(source, target string) ([][]int){
	// Levenshtein distance
	// https://en.wikipedia.org/wiki/Levenshtein_distance#Iterative_with_full_matrix
	a := make([][]int, len(source) + 1)
	for i := range a {
    		a[i] = make([]int, len(target) + 1) 
	}

	for i := range a {
		a[i][0] = i
	}

	for j := range a[0] {
		a[0][j] = j
	}
        substitutionCost := 0

        for j := 1; j <= len(target); j++ {
        	for i := 1; i <= len(source); i++ {
          		if source[i - 1] == target[j -1] {
            			substitutionCost = 0
			} else {
            			substitutionCost = 1
			}
          		
			changes := [...]int{
			     a[i-1][j] + 1,  			// deletion
                             a[i][j-1] + 1,                     // insertion
                             a[i-1][j-1] + substitutionCost}    // substitution

			a[i][j] = Min(changes)
		
		}
	}

	return a 

}


func getDistance(levMatrix [][]int, source string, target string) (int) {
	distance := levMatrix[len(source)][len(target)]
	return distance
}


func calculateDistance(source string, target string) (int) {
	mtx := getLevMatrix(source, target)
	distance := getDistance(mtx, source, target)
	return distance
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


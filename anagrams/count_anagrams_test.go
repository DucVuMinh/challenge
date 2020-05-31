package anagrams

import (
	"fmt"
	"testing"
)

var mapChar = map[rune]int64{
	'a' : 89,
	'b' : 3,
	'c' : 5,
	'd' : 7,
	'e' : 11,
	'f' : 13,
	'g' : 17,
	'h' : 23,
	'i' : 29,
	'j': 107,
	'k' : 31,
	'l' : 37,
	'm' : 41,
	'n' : 47,
	'o' : 2,
	'p' : 53,
	'q' : 59,
	'r' : 61,
	's' : 67,
	't' : 71,
	'v' : 73,
	'u' : 97,
	'z' : 79,
	'w' : 83,
	'x' : 101,
	'y': 103,
}

// Complete the sherlockAndAnagrams function below.
func SherlockAndAnagrams(s string) int32 {
	mapCount := map[int64]int64{}
	// find all substring
	// each substring using a hash function to calculate
	// hash function must using:
	//  1.length of string
	//  2.not care about the order of each charater
	// => substrings having same hash have high propability fulfill Anagrams
	slength := len(s)
	for i := 0; i < slength; i++{
		for j := i+1; j <= slength; j ++ {
			var mul int64 = 1
			sub := s[i:j]
			for _,v := range sub {
				mul *=mapChar[v]
			}
			if _,ok := mapCount[mul]; ok {
				mapCount[mul] = mapCount[mul] + 1
			} else {
				mapCount[mul] = 1
			}
		}
	}
	var count int64 = 0
	for _,v := range mapCount {
		count += (v * (v-1))/2
	}
	return int32(count)
}


func TestAnagrams(t *testing.T) {
	c := SherlockAndAnagrams("ifailuhkqqhucpoltgtyovarjsnrbfpvmupwjjjfiwwhrlkpekxxnebfrwibylcvkfealgonjkzwlyfhhkefuvgndgdnbelgruel")
	fmt.Println(c)
}

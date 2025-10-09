package main

//implement WordCount, it should return a map of the coutns of each word in the string 's'.

// Standard library strings.Fields : https://pkg.go.dev/strings#Fields

//func Fields(s string) []string

import "strings"

// "golang.org/x/tour/wc"

func WordCount(s string) map[string]int {
	var fields = strings.Fields(s)
	wordMap := make(map[string]int)
	for _, v := range fields {
		if _, ok := wordMap[v]; ok {
			wordMap[v] = wordMap[v] + 1
		} else {
			wordMap[v] = 1
		}
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}

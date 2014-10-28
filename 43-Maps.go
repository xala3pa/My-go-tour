package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    result := make(map[string]int)
    for _,w:= range strings.Fields(s){
    	result[w] += 1
    }
    return result
}

func main() {
    wc.Test(WordCount)
}
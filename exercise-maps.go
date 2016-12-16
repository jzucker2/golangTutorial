package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) (x map[string]int) {
	x = make(map[string]int)
  	for _, j := range strings.Fields(s) {
  		x[j]++
  	}
  	return
}

func main() {
	wc.Test(WordCount)
}
package main

import (
	//"fmt"
	"strings"
	"golang.org/x/tour/wc"	
)

func WordCount(s string) map[string]int {
	//tokens := []string;
	var tokens = strings.Fields(s);
	//fmt.Printf("Fields are :",tokens)
	wordcount := make(map[string]int)
	for _,v := range tokens{
		if(wordcount[v]<=0){
			wordcount[v] = 1;
		} else {
			wordcount[v] = wordcount[v] + 1;
		}
	}
	return wordcount
}

func main() {
	wc.Test(WordCount)
}

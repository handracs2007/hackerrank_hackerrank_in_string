// https://www.hackerrank.com/challenges/hackerrank-in-a-string/problem

package main

import "fmt"

func hackerrankInString(s string) string {
	var search = "hackerrank"
	var searchIndex = 0

	for _, chr := range s {
		if uint8(chr) == search[searchIndex] {
			searchIndex++

			if searchIndex == len(search) {
				return "YES"
			}
		}
	}

	return "NO"
}

func main() {
	fmt.Println(hackerrankInString("hhaacckkekraraannk"))
	fmt.Println(hackerrankInString("rhbaasdndfsdskgbfefdbrsdfhuyatrjtcrtyytktjjt"))
}

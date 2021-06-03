package basics

import "fmt"

func LearnBasics() []string {
	arr := []string{"hello", "world", "sofam"}
	for key, elem := range arr {
		fmt.Println(key, elem)
	}
	return arr
}

func LearnFunction(index int) []string {
	arr := []string{"hello", "world", "sofam"}
	return append(arr[:index], arr[index+1:]...)
}

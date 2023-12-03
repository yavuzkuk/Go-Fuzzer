package Support

import "fmt"

func Deneme(url string, character string) {
	var totalLenght = 100

	urlLenght := len(url)

	for i := 0; i < totalLenght-urlLenght; i++ {
		fmt.Print(character)
	}

}

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"main.go/Support"
)

var inputUrl string
var inputWordlist string
var currentUrl string
var character string = "-"
var inputStatusCode uint64
var totalTarget int
var successTarget int

func main() {

	defer fmt.Println("<< Tarama bitti >>")

	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		currentArg := args[i]
		switch currentArg {
		case "-u":
			inputUrl = args[i+1]
			break
		case "-w":
			inputWordlist = args[i+1]
			break
		case "-ch":
			character = args[i+1]
			break
		case "-f":
			inputStatusCode, _ = strconv.ParseUint(args[i+1], 10, 0)
			break
		case "-help":
			fmt.Println("-u --> Tarama yapılacak URL\n-W --> Arama yapılacak wordlist\n-ch --> aralarda kullanılıcak karakter\n-f --> Filtreleme için kullanılan parametre")
			break
		}
	}
	inputStatusCode := int(inputStatusCode)

	lastCharacter := inputUrl[len(inputUrl)-1]
	if string(lastCharacter) != "/" {
		inputUrl = inputUrl + "/"
	}

	var info, err = os.Stat(inputWordlist)
	if os.IsNotExist(err) {
		fmt.Println("Böyle bir dosya yok yada dizin hatalı", info)
	} else {
		fileLine, err := os.Open(inputWordlist)
		defer fileLine.Close()

		if err != nil {
			fmt.Println("Wordlist hatası")
		}

		scanner := bufio.NewScanner(fileLine)

		for scanner.Scan() {
			totalTarget = totalTarget + 1
			// fmt.Println(scanner.Text())
			currentUrl := inputUrl + scanner.Text()

			time.Sleep(time.Microsecond * 1)
			response, err3 := http.Get(currentUrl)

			if err3 != nil {
				fmt.Println("İstek atılamadı")
			}
			defer response.Body.Close()
			statusCode := response.StatusCode
			if inputStatusCode == 0 {

				fmt.Print(currentUrl)
				Support.Deneme(currentUrl, character)
				fmt.Print(statusCode, "\n")
				successTarget += 1
			} else if statusCode == inputStatusCode {
				fmt.Print(currentUrl)
				Support.Deneme(currentUrl, character)
				fmt.Print(statusCode, "\n")
				successTarget += 1

			}

		}
		color.Set(color.FgGreen)
		fmt.Println("Toplam ", totalTarget, " hedef tarandı")
		color.Unset()
		if inputStatusCode != 0 {
			color.Set(color.BgGreen)
			fmt.Println("------ Aradığınız kriterlere uygun ", successTarget, " hedef bulundu ------")
			color.Unset()
		}
	}

}

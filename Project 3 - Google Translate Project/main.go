package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/akhil/google-translate/cli"
)

var wg sync.WaitGroup

var sourceLang string
var targetLang string
var sourceText string

func init() {
	//flag can have bool, string etc. in this case we have taken stringVar so store value in sourceLand
	flag.StringVar(&sourceLang, "s", "en", "Source language [en]")
	//t is the paramter for target language, default is id
	flag.StringVar(&targetLang, "t", "fr", "Target language [fr]")
	flag.StringVar(&sourceText, "st", "", "Text to translate")
}

func main() {
	flag.Parse()

	//NFlag just returns the number of flags that have been set
	//so we're checking if it is zero

	if flag.NFlag() == 0 {
		//if zero flags have been set, we will show the usage options
		//os.Args has access to the command line arguments
		//fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		//this is where the magic happens, you are printing out the default values set above
		flag.PrintDefaults()
		os.Exit(1)
	}

	strChan := make(chan string)
	//waitgroup is created so that we can pass it to the function being called as the go-routine as we want to know
	//when it gets done, otherwise the main function doesn't end before the routine has finished
	//we are implementing concurrency here because we whenever we call APIs, and if we want to call APIs multiple
	//times, golang gives us the otion to call them simultaneously and this makes golang a very strong language
	//because when you're calling an API multiple times with different values, you may end up getting errors
	//and if it's a sequential program, not only will it take longer because there could be thousand other clients
	//making requests to the particular API but also, in some cases for some values the API may give error and this
	//may lead to the complete failure of the entire program as opposed to just one routine failing in the case of go
	//imp link - https://blog.logrocket.com/concurrency-patterns-golang-waitgroups-goroutines/

	//wg add, adds a counter, done reduces by 1 and wait waits for it to hit 0

	wg.Add(1)

	reqBody := &cli.RequestBody{
		SourceLang: sourceLang,
		TargetLang: targetLang,
		SourceText: sourceText,
	}

	go cli.RequestTranslate(reqBody, strChan, &wg)

	processedStr := strings.ReplaceAll(<-strChan, " + ", " ")

	fmt.Printf("%s\n", processedStr)

	close(strChan)
	wg.Wait()
}

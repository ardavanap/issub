package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

var WordlistSlice []string

func main() {
	protocolNumb, domain, wordlistpath, outputFile, delay := flags()
	WordlistSlice := givenSubWordlist(wordlistpath)
	if outputFile != "" {
		os.Create(outputFile) // If an output file is chosen by user,creates it.
	}

	switch protocolNumb {
	case 3:
		requestHttp(domain, WordlistSlice, outputFile, delay)
		break
	case 1:
		requesthttps(domain, WordlistSlice, outputFile, delay)
		break
	case 2:
		requestHttp(domain, WordlistSlice, outputFile, delay)
		requesthttps(domain, WordlistSlice, outputFile, delay)
		break
	}

}

func givenSubWordlist(wordlistpath string) (WordlistSlice []string) {

	wordlistFile, err := os.Open(wordlistpath) //Opening and reading wordlist of possible words as subdomain
	errorCheck(err)
	defer wordlistFile.Close()
	fileScanner := bufio.NewScanner(wordlistFile)

	for fileScanner.Scan() {
		WordlistSlice = append(WordlistSlice, fileScanner.Text())
		errorCheck(err)
	}

	return WordlistSlice
}

func flags() (result int, domainx string, wordlsistpath string, outputfileptr string, delay int) {

	wordlistPathptr := flag.String("w", "", "Path of [w]ordlist e.g: ~/wordlist")
	outputFileptr := flag.String("o", "", "[o]utput file destination. e.g: ~/Desktop")
	delayptr := flag.Int("delay", 1, "delay between each request (in seconds)")
	domainptr := flag.String("d", "", "[d]estinition of your choice. e.g:google.com ")
	ishttps := flag.Bool("https", false, "Only request on https protocol (true by default)")
	ishttp := flag.Bool("http", false, "Only request on http protocol")
	httpNhttps := flag.Bool("http-and-https", false, "Request on both http and https protocols")

	flag.Parse()

	if *domainptr == "" { //checks if user provided domain.if they didn't, program terminates
		err := errors.New(`[ERROR] I need a domain in order to find subdomains!`)
		errorCheck(err)
		os.Exit(0)
	}

	if *httpNhttps || (*ishttp && *ishttps) {
		result = 2
	} else if *ishttp {
		result = 3
	} else if *ishttps {
		result = 1
	} else {
		result = 1
	}
	return result, *domainptr, *wordlistPathptr, *outputFileptr, *delayptr
}

func requestHttp(domain string, WordlistSlice []string, outputFile string, delay int) {
	for _, word := range WordlistSlice {

		link := "http://" + word + "." + domain
		_, err := http.Get(link)
		if (outputFile != "") && (err == nil) {
			f, _ := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			defer f.Close()
			writer := bufio.NewWriter(f)    //if an output file is chosen by user,writes probed subdomains in the file
			writer.WriteString(link + "\n") //if not, writes it only as stdout
			err = writer.Flush()
			errorCheck(err)
		}
		if err == nil {
			fmt.Println(link)
		}
		for i := 0; i < delay; i++ {
			time.Sleep(1 * time.Second)
		}

	}

}

func requesthttps(domain string, WordlistSlice []string, outputFile string, delay int) { //works exacly like requesthttp,just in https
	for _, word := range WordlistSlice {

		link := "https://" + word + "." + domain
		_, err := http.Get(link)
		if (outputFile != "") && (err == nil) {
			f, _ := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			defer f.Close()
			writer := bufio.NewWriter(f)
			writer.WriteString(link + "\n")
			err = writer.Flush()
			errorCheck(err)
		}
		if err == nil {
			fmt.Println(link)
		}
		for i := 0; i < delay; i++ {
			time.Sleep(1 * time.Second)
		}

	}
}

func errorCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

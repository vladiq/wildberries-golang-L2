package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

/*
=== Утилита wget ===
Реализовать утилиту wget с возможностью скачивать сайты целиком
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const downloadFolder = "download"

func main() {
	mkdir(downloadFolder)
	for _, link := range os.Args[1:] {
		if numSaved, err := DownloadWebsite(link); err != nil {
			log.Panic(err)
		} else {
			fmt.Println(numSaved)
		}
	}
}

func DownloadWebsite(link string) (int, error) {
	urlSet := make(map[string]struct{})
	numSaved := 0

	link = strings.TrimRight(link, "/")
	URL, err := url.ParseRequestURI(link)
	if err != nil {
		return numSaved, err
	}

	re, err := regexp.Compile("https?://([a-z0-9]+[.])*" + URL.Host)
	if err != nil {
		return numSaved, err
	}

	mkdir(downloadFolder + "/" + URL.Host)

	collector := colly.NewCollector(colly.URLFilters(re))

	collector.OnHTML("a[href]", func(el *colly.HTMLElement) {
		ul := el.Request.AbsoluteURL(el.Attr("href"))
		if _, isExist := urlSet[ul]; !isExist {
			urlSet[ul] = struct{}{}
			collector.Visit(ul)
		}
	})

	collector.OnResponse(func(r *colly.Response) {
		reqUrlPath := r.Request.URL.Path
		fullPath := downloadFolder + "/" + URL.Hostname() + reqUrlPath

		if _, ok := urlSet[fullPath]; ok {
			return
		}

		urlSet[fullPath] = struct{}{}
		if path.Ext(fullPath) == "" {
			mkdir(fullPath)
		} else {
			mkdir(fullPath[:strings.LastIndexByte(fullPath, '/')])
		}

		if path.Ext(reqUrlPath) == "" {
			if fullPath[len(fullPath)-1] != '/' {
				fullPath += "/"
			}
			fullPath += "index.html"
			if _, err := os.Create(fullPath); err != nil {
				fmt.Printf("error creating file: %s\n", err.Error())
			}
		}

		if err = r.Save(fullPath); err != nil {
			panic(err)
		}

		fmt.Println("saved:", URL.Hostname()+reqUrlPath)
		numSaved++
	})

	if err = collector.Visit(URL.String()); err != nil {
		log.Panic("err: visit: " + err.Error())
	}
	collector.Wait()
	return numSaved, nil
}

func mkdir(folderName string) {
	_, err := os.Stat(folderName)
	if os.IsNotExist(err) && os.MkdirAll(folderName, os.ModePerm) != nil {
		log.Panic(err)
	}
}

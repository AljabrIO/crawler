package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sort"

	"github.com/jackdanger/collectlinks"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		urlStr := args[0]
		u, err := url.Parse(urlStr)
		if err != nil {
			fmt.Printf("Invalid URL: %s\n", err)
			os.Exit(1)
		}
		if resp, err := http.Get(urlStr); err != nil {
			fmt.Printf("Getting URL failed: %s\n", err)
			os.Exit(1)
		} else {
			links := collectlinks.All(resp.Body)
			for i, l := range links {
				if x, err := url.Parse(l); err == nil && x.Host != "" {
					// No changes needed
				} else {
					// Probably not a full URL
					lu := *u
					lu.Path = l
					links[i] = lu.String()
				}
			}
			// De-duplicate
			links = distinct(links)
			// Sort links
			sort.Strings(links)
			// Show results
			for _, l := range links {
				fmt.Println(l)
			}
		}
	} else {
		fmt.Println("Please provide URL argument")
		os.Exit(1)
	}
}

func distinct(arg []string) []string {
	tempMap := make(map[string]struct{})

	for _, s := range arg {
		tempMap[s] = struct{}{}
	}

	result := make([]string, 0, len(tempMap))
	for key := range tempMap {
		result = append(result, key)
	}
	return result
}

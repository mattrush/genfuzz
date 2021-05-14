package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func createFuzz(baseUrl string, params string) {
	if len(params) < 0 {
		return
	}

	v, err := url.ParseQuery(params)
	if err != nil {
		log.Fatal(err)
	}

	/* make a copy of the map so we can easily permute one key per pass */
	done := make(map[string][]string, len(v))
	for k, v := range v {
		done[k] = v
	}

	/* we need to generate 2 permutations of every parameter value: FUZZ, and originalFUZZ */
	q := url.Values{}
	for key := range done {
		for i := 0; i < 2; i++ {
			for k, value := range v {
				if key == k && i == 0 {
					q.Set(k, "FUZZ")
				} else if key == k && i == 1 {
					q.Set(k, string(strings.Join(value, ",")+"FUZZ"))
				} else {
					q.Set(k, strings.Join(value, ","))
				}
			}
			fmt.Printf("%s?%s\n", strings.Split(baseUrl, "?")[0], q.Encode())
		}
	}
}

func main() {
        url := "https://www.google.com/test/a"
        params := "a=1&b=a&c=true"

        createFuzz(url, params)
}

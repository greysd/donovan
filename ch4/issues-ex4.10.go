package main

import (
	"fmt"
	"log"
	"os"

	"time"

	"github.com/greysd/gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d тем:\n", result.TotalCount)
	to := time.Now().UTC()
	for _, word := range []int{30, 365, 3650} {
		from := to.AddDate(0, 0, -1*word)
		fmt.Printf("Period from: %v, and to: %v\n", from, to)
		for _, item := range result.Items {
			if item.CreatedAt.After(from) && item.CreatedAt.Before(to) {
				fmt.Printf("# %4d-%02d-%02d:\t%-5d\t%9.9s %.55s\n",
					item.CreatedAt.Year(),
					item.CreatedAt.Month(),
					item.CreatedAt.Day(),
					item.Number, item.User.Login, item.Title)
			}
		}
		to = from
	}
}

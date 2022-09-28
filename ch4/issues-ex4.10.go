package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/greysd/gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d тем:\n", result.TotalCount)
	for _, word := range strings.Fields("month, year, year+") {
		for _, item := range result.Items {
			fmt.Printf("%+v\n", item.CreatedAt)
			fmt.Printf("%v")
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}

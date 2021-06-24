package main

import (
	"0x04/github"
	"0x04/movie"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	//data, err := json.Marshal(movie.Movies)
	data, err := json.MarshalIndent(movie.Movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	fmt.Println(titles)

	fmt.Println("----------------------------------------------------")

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues :\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

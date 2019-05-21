package main

import (
  "fmt"
  "log"
  "flag"

  "github.com/jordify/reddit"
)

func main() {
  subreddit := flag.String("subreddit", "golang", "Subreddit to query")
  flag.Parse()

  items, err := reddit.Get(*subreddit)
  if err != nil {
    log.Fatal(err)
  }
  for _, item := range items {
    fmt.Println(item)
  }
}

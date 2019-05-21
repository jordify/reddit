// Package reddit implements a basic client for the Reddit API.
package reddit

import (
  "fmt"
  "encoding/json"
  "net/http"
  "errors"
)

// Item describes a Reddit API item.
type Item struct {
  Title string
  URL   string
  Comments int `json:"num_comments"`
}

type response struct {
  Data struct {
    Children []struct {
      Data Item
    }
  }
}

func (i Item) String() string {
  com := ""
  switch i.Comments {
  case 0:
    // nothing
  case 1:
    com = " (1 comment)"
  default:
    com = fmt.Sprintf(" (%d comments)", i.Comments)
  }
  return fmt.Sprintf("%s%s\n%s", i.Title, com, i.URL)
}

// Get fetches the most recent Items posted to the specified subreddit.
func Get(subreddit string) ([]Item, error) {
  url := fmt.Sprintf("https://reddit.com/r/%s.json", subreddit)

  client := &http.Client{}
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return nil, err
  }
  req.Header.Set("User-Agent", "jordify-subreddit-reader:v0.1")
  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  if resp.StatusCode != http.StatusOK {
    return nil, errors.New(resp.Status)
  }

  defer resp.Body.Close()
  r := new(response)
  err = json.NewDecoder(resp.Body).Decode(r)
  if err != nil {
    return nil, err
  }
  items := make([]Item, len(r.Data.Children))
  for i, child := range r.Data.Children {
    items[i] = child.Data
  }
  return items, nil
}

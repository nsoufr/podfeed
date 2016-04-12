# podfeed
Podcast Feed Golang Parser

## Usage

```go

package main

import (
  "fmt"
  "log"

  "github.com/nandosousafr/podfeed"
)

func main() {
  podcast, err := podfeed.Fetch("http://feeds.soundcloud.com/users/soundcloud:users:212089450/sounds.rss")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(podcast.Title)
  // => CapyCast

  fmt.Println(podcast.Items[0].Title)
  // => Capycast #4 Solopreneur, Entrepreneur, Intrapreneur

  fmt.Println(podcast.Owner.Name)
  // => Ship It
}
```

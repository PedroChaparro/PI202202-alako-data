package main

import (
  "fmt"
  "github.com/PedroChaparro/PI202202-alako-data/queries"
  "github.com/go-rod/rod"
)

func getLinks(url string, browser *rod.Browser) (links []string) {
  // Create the new tab
  page := browser.MustPage(url).MustWaitLoad()
  defer page.MustClose()

  // Count videos
  elements := page.MustElements("#video-title")

  // Scroll until get at leat 140 videos
  for len(elements) < 140 {
    page.Eval("window.scroll({top: 9999999999})")
    elements = page.MustElements("#video-title")
    fmt.Println(len(elements))
  }

  fmt.Println(len(elements))

  return []string{}
}

func main(){
  // Launch browser instance
  browser := rod.New().MustConnect()
  defer browser.MustClose()

  for _, query := range(queries.Queries) {
    getLinks(query.Url, browser)
  }
}

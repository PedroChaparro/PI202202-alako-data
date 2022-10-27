package main

import (
  "fmt"
  "io/ioutil" // Parse requests
  "net/http" // Requests
  "regexp"
  "github.com/PedroChaparro/PI202202-alako-data/queries"
  "github.com/PedroChaparro/PI202202-alako-data/interfaces"
  "github.com/go-rod/rod" // Crawler
  "github.com/remeh/sizedwaitgroup" // Concurrency
)

// Return resulting links for user query
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
  }

  // Get links
  links = []string{}

  for _, element := range(elements) {
    val := element.MustAttribute("href")
   
    // Prevent nil pointer dereference error 
    if val != nil {
      links = append(links, fmt.Sprintf("https://youtube.com%s", *val))
    }
  }

  return links
}

// Other fucntion
func getData(url string, swg *sizedwaitgroup.SizedWaitGroup) (video interfaces.Video, err error) {
  defer swg.Done() // Finish current "job"
  
  // Get plain html
  reply, err := http.Get(url)
  defer reply.Body.Close()
  
  video = interfaces.Video{}

  if err != nil {
    return video, err
  }

  buffer, err := ioutil.ReadAll(reply.Body)
  html := string(buffer)

  if err != nil {
    return video, err
  }

  titleRegExp := regexp.MustCompile(`<meta name="title"[^>]+content="(.*?)"`)
  tagsRegExp := regexp.MustCompile(`<meta name="keywords"[^>]+content="(.*?)"`)
  thumbnailRegExp := regexp.MustCompile(`<link rel="image_src"[^>]+href="(.*?)"`)

  fmt.Printf("%q\n", titleRegExp.FindString(html))
  video.Title = titleRegExp.FindString(html)
  video.Tags = tagsRegExp.FindString(html)
  video.Thumbnail = thumbnailRegExp.FindString(html)

  fmt.Println(video)

  return video, nil
}

func main(){
  // Launch browser instance
  browser := rod.New().MustConnect()
  defer browser.MustClose()

  // For each query
  for _, query := range(queries.Queries) {
    qLinks := getLinks(query.Url, browser)
    // For each query link
    swg := sizedwaitgroup.New(16)
    for _, link := range(qLinks) {
      swg.Add()
      go getData(link, &swg)
    }
  }
}

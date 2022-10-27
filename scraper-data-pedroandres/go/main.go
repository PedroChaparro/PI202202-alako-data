package main

import (
  "fmt"
  "io/ioutil" // Parse requests
  "net/http" // Requests
  "encoding/json" // Write json file
  "os"
  "time"
  "regexp" // Regular expressions
  "strings"
  "github.com/PedroChaparro/PI202202-alako-data/queries"
  "github.com/PedroChaparro/PI202202-alako-data/interfaces"
  "github.com/go-rod/rod" // Crawler
  "github.com/remeh/sizedwaitgroup" // Concurrency
)

// Global variable
var videos = interfaces.ConcurrentSlice{}

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
  fmt.Printf("Current url: %s\n", url)
 
  // ### ### ###
  // Get plain html
  reply, err := http.Get(url)
  defer reply.Body.Close()
  
  video = interfaces.Video{Url: url}

  if err != nil {
    return video, err
  }

  buffer, err := ioutil.ReadAll(reply.Body)
  html := string(buffer)

  if err != nil {
    return video, err
  }

  // ### ### ###
  // Parse plain html
  titleRegExp := regexp.MustCompile(`<meta name="title"[^>]+content="(.*?)"`)
  descriptionRegExp := regexp.MustCompile(`"shortDescription":"(.*?)"`)
  tagsRegExp := regexp.MustCompile(`<meta name="keywords"[^>]+content="(.*?)"`)
  thumbnailRegExp := regexp.MustCompile(`<link rel="image_src"[^>]+href="(.*?)"`)

  video.Title = strings.Split(titleRegExp.FindString(html), `content="`)[1]
  video.Title = strings.Replace(video.Title, `"`, "", len(video.Title)) // Remove quotes
  video.Description = strings.Split(descriptionRegExp.FindString(html), `"shortDescription":"`) [1]
  video.Description = strings.Replace(video.Description, `"`, "", len(video.Description))
  video.Tags = strings.Split(tagsRegExp.FindString(html), `content="`)[1]
  video.Tags = strings.Replace(video.Tags, `"`, "", len(video.Tags))
  video.Thumbnail = strings.Split(thumbnailRegExp.FindString(html), `href="`)[1]
  video.Thumbnail = strings.Replace(video.Thumbnail, `"`, "", len(video.Thumbnail))

  // ### ### ##3
  // Clear data 
  spacesRegExp := regexp.MustCompile(`\s\s+`)
  lineBreakRegExp := regexp.MustCompile(`\\+n`)
  
  video.Title = spacesRegExp.ReplaceAllString(video.Title, " ")
  video.Title = lineBreakRegExp.ReplaceAllString(video.Title, "")
  video.Title = strings.TrimSpace(video.Title)

  video.Description = spacesRegExp.ReplaceAllString(video.Description, " ")
  video.Description = lineBreakRegExp.ReplaceAllString(video.Description, "")
  video.Description = strings.TrimSpace(video.Description)

  // fmt.Printf("%+v\n", video)

  videos.Lock() // Lock, so it's possible to add whitout lossing data
  defer videos.Unlock()	// Unlock, so other routines can add data
  videos.Items = append(videos.Items, video)

  return video, nil
}

func main(){
  // Launch browser instance
  browser := rod.New().MustConnect()
  defer browser.MustClose()

  // For each query
  for _, query := range(queries.Queries) {
    currLength := len(videos.Items)
    start := time.Now()
    fmt.Printf("🏃 Starting with query: %s\n", query.Query)
    qLinks := getLinks(query.Url, browser)

    // For each query link
    swg := sizedwaitgroup.New(8) // 8 Concurrent routines
    for _, link := range(qLinks) {
      swg.Add()
      go getData(link, &swg)
    }

    fmt.Printf("🏁 %d videos were saved in %s\n", len(videos.Items) - currLength, time.Since(start))

  }

  // Create json file
  jsonString, _ := json.Marshal(videos.Items)
  ioutil.WriteFile("data.json", jsonString, os.ModePerm)

}

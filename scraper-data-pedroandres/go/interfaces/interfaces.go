package interfaces

import (
  "sync"
)

type Query struct {
  Query		string		
  Url		string		`json:Url`
}

type Video struct {
  Title		string
  Description	string
  Tags		string
  Url		string
  Thumbnail 	string
}

type ConcurrentSlice struct {
  sync.RWMutex
  Items []Video
}

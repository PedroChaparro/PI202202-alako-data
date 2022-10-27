package interfaces

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

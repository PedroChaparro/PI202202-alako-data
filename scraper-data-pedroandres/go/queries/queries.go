package queries

type Query struct {
  Query		string		
  Url		string		`json:Url`
}

var Queries = []Query{
  {
    Query: "Colombian touristic places",
    Url: "https://www.youtube.com/results?search_query=Colombian+touristic+places",
  },
}

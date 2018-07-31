package engine

type Request struct {
	Url        string
	ParserFunc func(contents []byte, url string) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

func NilParser() ParseResult {
	return ParseResult{}
}

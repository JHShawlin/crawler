package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re, err := regexp.Compile(cityRe)
	if err != nil {
		panic(err)
	}
	matcher := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matcher {
		//result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        string(m[1]),
				ParserFunc: engine.NilParser,
			})
	}
	return result
}

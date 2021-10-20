package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re, err := regexp.Compile(cityListRe)
	if err != nil {
		panic(err)
	}
	matcher := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matcher {
		//result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}
	return result
}

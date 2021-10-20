package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
)

const profileRe = `<div class="m-btn[^>]+>([^<]+)</div>`

func ParserProfile(name string, contents []byte, url string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	re := regexp.MustCompile(profileRe)
	match := re.FindAllSubmatch(contents, -1)

	for _, m := range match {
		profile.Marital_status = string(m[1])
		profile.Age = string(m[2])
		profile.Constellation = string(m[3])
		profile.Height = string(m[4])
		profile.Workplace = string(m[5])
		profile.Income = string(m[6])
		profile.Education = string(m[7])
	}

	result := engine.ParseResult{
		Requests: []engine.Request{},
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      "",
				Payload: profile,
			},
		},
	}
	return result
}

package persist

import (
	"context"
	"crawler/engine"
	"encoding/json"
	"testing"

	"github.com/olivere/elastic"
)

type peo struct {
	Name string
	Age  int
}

func FromJsonObj(o interface{}) (peo, error) {
	var p peo
	s, err := json.Marshal(o)
	if err != nil {
		return p, err
	}

	err = json.Unmarshal(s, &p)
	return p, err
}

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://www.zhenai.com/zhenghun/anhui",
		Type: "zhenai",
		Id:   "100",
		Payload: peo{
			Name: "Jun",
			Age:  24,
		},
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	err = save(client, expected)

	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index("dating_file").Type("zhenai").Id(expected.Id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)

	if err != nil {
		panic(err)
	}

	actualpeo, _ := FromJsonObj(actual.Payload)
	actual.Payload = actualpeo

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}

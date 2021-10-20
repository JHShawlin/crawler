package fetcher

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

//var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	//<-rateLimiter
	resp, err := JsonGet(url, nil, map[string]string{
		"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_3_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.128 Safari/537.36",
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	body := bufio.NewReader(resp.Body)
	e := determineEncoding(body)
	utf8Reader := transform.NewReader(body, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func JsonGet(requestUrl string, params map[string]string, headers map[string]string) (*http.Response, error) {
	var req *http.Request
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, errors.New("new request is fail: %v ")
	}

	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	for key, val := range headers {
		req.Header.Add(key, val)
	}
	//http client
	c := &http.Client{
		Timeout: time.Second * 4,
	}
	fmt.Println(req.Header)
	return c.Do(req)
}

package ynote

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

type YClient struct {
	Hc        *http.Client
	CookieJar *cookiejar.Jar
}

func New() *YClient {
	CookieJar, _ := cookiejar.New(nil)
	HttpClient := &http.Client{
		CheckRedirect: nil,
		Jar:           CookieJar,
	}
	yclient := &YClient{
		HttpClient,
		CookieJar,
	}
	return yclient
}
func (c *YClient) Login() error {
	response, _ := c.Hc.Get("http://baidu.com")
	fmt.Print(response)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(body))
	return nil
}

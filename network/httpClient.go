package network

import "github.com/Danny-Dasilva/CycleTLS/cycletls"

type Client struct {
	httpClient cycletls.CycleTLS
	options    cycletls.Options
}

func InitClient(headers map[string]string) *Client {
	client := cycletls.Init()
	options := cycletls.Options{
		Body:      "",
		Ja3:       "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-51-57-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0",
		UserAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0",
		Headers:   headers,
	}

	httpClient := Client{
		httpClient: client,
		options:    options,
	}

	return &httpClient
}

func (c Client) GetRequest(url string) (cycletls.Response, error) {
	return c.httpClient.Do(url, c.options, "GET")
}

func (c Client) PostRequest(url string, body string) (cycletls.Response, error) {
	c.options.Body = body
	c.options.Headers["Content-Type"] = "application/json"
	return c.httpClient.Do(url, c.options, "POST")
}
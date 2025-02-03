package main

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
	var res cycletls.Response
	response, error := c.httpClient.Do(url, c.options, "GET")
	if error != nil {
		return res, error
	}
	return response, nil
}

func main() {
	headers := map[string]string{
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Accept-Language": "en-US,en;q=0.5",
	}
	client := InitClient(headers)
	url := "https://www.zillow.com/search/GetSearchPageState.htm?searchQueryState=%7B%22mapBounds%22%3A%7B%22west%22%3A-112.29569230899439%2C%22east%22%3A-77.75467668399439%2C%22south%22%3A36.062333260105035%2C%22north%22%3A49.43591374214795%7D%2C%22isMapVisible%22%3Atrue%2C%22filterState%22%3A%7B%22sortSelection%22%3A%7B%22value%22%3A%22days%22%7D%2C%22isComingSoon%22%3A%7B%22value%22%3Afalse%7D%2C%22isPreMarketForeclosure%22%3A%7B%22value%22%3Atrue%7D%2C%22isPreMarketPreForeclosure%22%3A%7B%22value%22%3Atrue%7D%2C%22isAllHomes%22%3A%7B%22value%22%3Atrue%7D%2C%22isAcceptingBackupOffersSelected%22%3A%7B%22value%22%3Atrue%7D%2C%22isForSaleByAgent%22%3A%7B%22value%22%3Afalse%7D%7D%2C%22isListVisible%22%3Atrue%2C%22mapZoom%22%3A5%2C%22category%22%3A%22cat2%22%2C%22pagination%22%3A%7B%22currentPage%22%3A6%7D%7D&wants=%7B%22cat2%22%3A%5B%22listResults%22%2C%22mapResults%22%5D%2C%22cat1%22%3A%5B%22total%22%5D%7D&requestId=32"
	response, error := client.GetRequest(url)
	if error != nil {
		panic(error)
	}
	println(response.Body)
}

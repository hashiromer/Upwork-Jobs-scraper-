package upwork

import (
	"fmt"
	"scrapers/network"
)

type Upwork struct {
	UpworkHttpClient *network.Client
}

type UrlArgs struct {
	Page     int
	Per_Page int
	Query    string
}

func (u Upwork) ConstructUrl(args UrlArgs) string {
	url := "https://www.upwork.com/search/jobs/url?q=%s&per_page=%d&sort=recency&page=%d"

	return fmt.Sprintf(url, args.Query, args.Per_Page, args.Page)
}

func (u Upwork) SendRequest(url string) (string, error) {
	Upclient := u.UpworkHttpClient
	resp, err := Upclient.GetRequest(url)

	if err != nil {
		return "", err
	}
	return resp.Body, nil

}

func InitUpwork() *Upwork {
	headers := map[string]string{
		"authority":             "www.upwork.com",
		"accept":                "application/json, text/plain",
		"accept-language":       "en",
		"cache-control":         "no-cache",
		"pragma":                "no-cache",
		"referer":               "https://www.upwork.com/search/jobs/url?per_page=10&sort=recency",
		"sec-fetch-site":        "same-origin",
		"sec-gpc":               "1",
		"vnd-eo-parent-span-id": "2724011d-2430-47f5-b5b9-603f2e919685",
		"vnd-eo-span-id":        "9d6e5b36-ace2-402e-a188-01da1d6b84ee",
		"x-odesk-user-agent":    "oDesk LM",
		"x-requested-with":      "XMLHttpRequest",
	}
	client := network.InitClient(headers)
	upwork := Upwork{
		UpworkHttpClient: client,
	}

	return &upwork

}

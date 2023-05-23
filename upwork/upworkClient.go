package upwork

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"scrapers/network"
	"strings"
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
	url := "https://www.upwork.com/ab/jobs/search/url?per_page=%d&sort=recency&payment_verified=1&page=%d&q=%s"
	return fmt.Sprintf(url, args.Per_Page, args.Page, args.Query)
}

func (u Upwork) SendRequest(url string) (string, error) {
	Upclient := u.UpworkHttpClient
	resp, err := Upclient.GetRequest(url)

	if err != nil {
		return "", err
	}
	return resp.Body, nil

}

func mergeMaps(m1, m2 map[string]string) map[string]string {
	// Iterate over m2 and add its key-value pairs to m1
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

func readEnv(filename string) (map[string]string, error) {

	// Open the .env file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Print()
		return nil, err
	}
	defer file.Close()

	// Create a map to store the key-value pairs
	m := make(map[string]string)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line on the "=" character
		parts := strings.SplitN(scanner.Text(), "=", 2)
		if len(parts) == 2 {
			// Trim leading and trailing whitespace from the key and value
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			// Add the key-value pair to the map
			m[key] = value
		}
	}

	return m, nil
}

func InitUpwork() *Upwork {
	headers := map[string]string{
		"authority":       "www.upwork.com",
		"accept":          "application/json, text/plain",
		"accept-language": "en",
		"cache-control":   "no-cache",
		"pragma":          "no-cache",
		"referer":         "https://www.upwork.com/search/jobs/url?per_page=10&sort=recency",
		"sec-fetch-site":  "same-origin",
		"sec-gpc":         "1",

		"vnd-eo-parent-span-id": "b8f3fe84-2aa4-4b9e-b750-6dc0e6e169a9",
		"vnd-eo-span-id":        "0d5bed38-342c-4b44-a885-b9c3de71a32a",
		"x-odesk-user-agent":    "oDesk LM",
		"x-requested-with":      "XMLHttpRequest",
	}
	auth_headers, err := readEnv("upwork/.env")
	if err != nil {
		log.Fatal("Could not read .env file, please add .env file in upwork folder")
	}
	headers = mergeMaps(auth_headers, headers)

	client := network.InitClient(headers)
	upwork := Upwork{
		UpworkHttpClient: client,
	}

	return &upwork

}

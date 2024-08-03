package upwork

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"scrapers/network"
	"strings"
)

type Upwork struct {
	UpworkHttpClient *network.Client
}

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func (u Upwork) SendRequest(query string, variables map[string]interface{}) (string, error) {
	url := "https://www.upwork.com/api/graphql/v1"
	
	requestBody := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}
	
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}
	
	resp, err := u.UpworkHttpClient.PostRequest(url, string(jsonBody))
	if err != nil {
		return "", err
	}
	
	return resp.Body, nil
}

func readGraphQLQuery(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func mergeMaps(m1, m2 map[string]string) map[string]string {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

func readEnv(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	m := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
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
		"content-type":    "application/json",
		"cache-control":   "no-cache",
		"pragma":          "no-cache",
		"referer":         "https://www.upwork.com/search/jobs/",
		"sec-fetch-site":  "same-origin",
		"sec-gpc":         "1",
		"x-requested-with": "XMLHttpRequest",
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
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

func get_data(url string, headers map[string]string) (string, error) {
	client := cycletls.Init()
	response, err := client.Do(
		url,
		cycletls.Options{
			Body:      "",
			Ja3:       "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-51-57-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0",
			UserAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0",
			Headers:   headers,
		},
		"GET",
	)

	if err == nil {
		return string(response.Body), nil
	}
	return response.Body, nil

}

func save_to_file(data string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data)
	if err != nil {
		return err
	}

	writer.Flush()
	return nil
}

// func check_captcha(data map[string]string) bool {
// 	if _, ok := data["blockScript"]; ok {
// 		return true
// 	}
// 	return false

// }

func main() {

	os.RemoveAll("data")

	os.Mkdir("data", 0755)

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
	//Upwork limits pagination to 100 pages
	total_iterations := 10
	//Query to serach for on Upwork, searching for jobs with shopify keyword
	query := "shopify"
	//Number of results per page
	per_page := 100

	start := 1
	for i := start; i <= total_iterations; i++ {
		fmt.Print("Iteration: ", i, "\n")
		upwork_api_url_template := "https://www.upwork.com/search/jobs/url?q=%s&per_page=%d&sort=recency&page=%d"
		url := fmt.Sprintf(upwork_api_url_template, query, per_page, i)

		time.Sleep(2 * time.Second)

		data, err := get_data(url, headers)
		if err != nil {
			fmt.Println(err)
			fmt.Println(i)
			panic(err)

		}

		//Get body of the response
		filename := fmt.Sprintf("data/%d.json", i)
		err = save_to_file(data, filename)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}

	fmt.Println("Scraping done")

	files, err := filepath.Glob("data/*.json")
	if err != nil {
		panic(err)
	}

	var all_jobs []map[string]interface{}
	for _, file := range files {
		fmt.Println(file)
		data, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}

		//Parse data as json without interface
		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		if err != nil {
			panic(err)
		}

		//Get value from key
		key := "searchResults"
		value := result[key]

		//Check for errors
		is_error := value.(map[string]interface{})["jobSearchError"]

		//Skip the file if is_error is True

		if is_error == true {
			fmt.Println("Error")
			continue
		}

		//Get jobs from the json
		jobs := value.(map[string]interface{})["jobs"]

		//Add all jobs to the all_jobs slice
		for _, job := range jobs.([]interface{}) {
			all_jobs = append(all_jobs, job.(map[string]interface{}))
		}

	}

	jobs := map[string]interface{}{
		"jobs": all_jobs,
	}

	json_data, err := json.Marshal(jobs)

	if err != nil {
		panic(err)
	}

	err = save_to_file(string(json_data), "all_jobs.json")
	if err != nil {
		panic(err)
	}

	os.RemoveAll("data")

}

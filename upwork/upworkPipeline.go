package upwork

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type UpworkPipeLine struct {
	upworkClient *Upwork
	// queue := []int{}
	filepaths []string
}

func InitPipeline() *UpworkPipeLine {
	u := UpworkPipeLine{
		upworkClient: InitUpwork(),
	}
	return &u
}

func (u *UpworkPipeLine) CombineFiles() error {

	var all_jobs []interface{}

	for _, file := range u.filepaths {
		data, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}
		var result map[string]interface{}

		//Parse data as json
		err = json.Unmarshal([]byte(data), &result)
		if err != nil {
			panic(err)
		}

		key := "searchResults"
		value := result[key]
		jobs := value.(map[string]interface{})["jobs"].([]interface{})
		for _, job := range jobs {
			all_jobs = append(all_jobs, job.(map[string]interface{}))
		}

		now := time.Now()

		// Format the date and time using the "2006-01-02" layout
		dateString := now.Format("2006-01-02")

		// Create the file name using the date string
		filename := "file_" + dateString + ".json"

		//save to file

		//Convert to json
		json_data, err := json.Marshal(all_jobs)
		if err != nil {
			panic(err)
		}
		err = u.saveToFile(json_data, filename)
		if err != nil {
			panic(err)
		}

	}
	return nil
}

func (u *UpworkPipeLine) saveToFile(data []byte, filename string) error {

	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil

}

func isValidJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

// func print_json(p string) {
// 	b, err := json.MarshalIndent(p, "", "  ")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(string(b))
// }

func isValidApiResponse(data string) bool {

	// Deserialize the JSON string into the Person struct
	var api_error LoggedOutError
	var api_response UpworkApiResponse

	// Deserialize the JSON string into the Person struct
	is_api_error := json.Unmarshal([]byte(data), &api_error)
	api_resp := json.Unmarshal([]byte(data), &api_response)

	//A known error occured
	if is_api_error == nil && api_resp != nil {

		log.Print(api_error.Error.Message)
		return false
		//Unknown response format
	} else if is_api_error != nil && api_resp != nil {

		fmt.Print("Unknown response format")
		fmt.Print(data)
		return false

	} else {

		return !api_response.SearchResults.JobSearchError

	}

}

func (u *UpworkPipeLine) isResponseValid(data string) bool {

	return isValidJSON(data) && isValidApiResponse(data)

}

func (u *UpworkPipeLine) getTotalDocuments(urlArgs UrlArgs) (int, error) {

	var API_Response UpworkApiResponse
	client := u.upworkClient
	url := client.ConstructUrl(urlArgs)
	resp, err := u.upworkClient.SendRequest(url)

	if err != nil {

		log.Fatal(err)
		return 0, err
	}

	// check if response is valid
	if !u.isResponseValid(resp) {
		return 0, fmt.Errorf("invalid response")
	}

	json.Unmarshal([]byte(resp), &API_Response)

	total_docs := API_Response.SearchResults.Paging.Total
	s := fmt.Sprintf("totalzz is %d", total_docs)
	log.Print(s)
	return total_docs, nil
}

func (u *UpworkPipeLine) handleRequest(urlArgs UrlArgs, iteration int) {
	client := u.upworkClient
	url := client.ConstructUrl(urlArgs)
	resp, err := u.upworkClient.SendRequest(url)
	if err != nil {
		log.Fatal(err)

		//check if response is valid
	} else if u.isResponseValid(resp) {
		filename := fmt.Sprintf("data/%d.json", iteration)
		// Convert resp to array of bytes
		err = u.saveToFile([]byte(resp), filename)
		if err != nil {
			log.Println(err)
			panic(errors.New("could not save file"))
		} else {
			u.filepaths = append(u.filepaths, filename)

		}

	} else {
		log.Println("Invalid response returned")

	}

}

func (u *UpworkPipeLine) Run(query string) error {

	err := os.RemoveAll("data")
	if err != nil {
		return err
	}
	err = os.Mkdir("data", 0755)
	if err != nil {
		return err
	}

	var iteration int
	var perPage int
	var total_docs int

	qum := query

	if qum == "" {
		qum = "Empty String"
	}

	info_message := fmt.Sprintf("Finding Total Jobs for %s", qum)
	fmt.Println(info_message)

	urlArgs := UrlArgs{
		Page:     1,
		Per_Page: 10,
		Query:    query,
	}

	//Find total number of iterations
	perPage = 50
	total_docs, err = u.getTotalDocuments(urlArgs)

	if err != nil {
		log.Print("Could not retrive total number of jobs")
		log.Fatal(err)
	}

	info_message = fmt.Sprintf("%s has a total of %d jobs", query, total_docs)
	fmt.Println(info_message)
	iteration = total_docs / perPage

	if iteration >= 100 {
		iteration = 100
	}

	info_message = fmt.Sprintf("A total of %d iterations will be performed", iteration)
	fmt.Println(info_message)

	//Found total iterations
	u.handledataIteration(perPage, iteration, query)
	err = u.CombineFiles()
	os.RemoveAll("data")

	if err != nil {
		panic(err)

	}

	return nil

}

func (u *UpworkPipeLine) handledataIteration(p_per int, iters int, query string) {
	for index := 1; index <= iters; index++ {
		log.Println("Iteration: ", index)
		urlArgs := UrlArgs{
			Page:     index,
			Per_Page: 50,
			Query:    query,
		}

		u.handleRequest(urlArgs, index)

	}
}

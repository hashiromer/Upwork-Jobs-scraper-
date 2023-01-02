package upwork

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
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
	// all_filenames, err := filepath.Glob("data/*.json")
	// if err != nil {
	// 	panic(err)
	// }

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

		//save to file
		filename := "all_jobs.json"

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

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

func isApiError(data string) bool {

	var result map[string]interface{}

	//Parse data as json
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		log.Print("The Api did not return expected response")
		log.Print("The following was the response from API")
		log.Print(data)
		panic(err)
	}

	//Get value from key
	key := "searchResults"
	value := result[key]

	//Check for errors
	is_error := value.(map[string]interface{})["jobSearchError"]

	return is_error == true
}

func (u *UpworkPipeLine) validateResponse(data string) bool {
	return isJSON(data) && !isApiError(data)

}

func (u *UpworkPipeLine) getTotalDocuments(urlArgs UrlArgs) (int, error) {
	client := u.upworkClient
	url := client.ConstructUrl(urlArgs)
	resp, err := u.upworkClient.SendRequest(url)

	if err != nil {

		log.Fatal(err)
		return 0, err
	}

	// check if response is valid
	if !u.validateResponse(resp) {
		return 0, fmt.Errorf("invalid response")
	}

	var API_Response UpworkApiResponse

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
	} else if u.validateResponse(resp) {
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

	info_message := fmt.Sprintf("Finding Total Jobs for %s", query)
	fmt.Println(info_message)

	urlArgs := UrlArgs{
		Page:     1,
		Per_Page: 10,
		Query:    query,
	}

	//Find total number of iterations
	perPage = 50
	total_docs, err = u.getTotalDocuments(urlArgs)

	log.Print(total_docs)
	if err == nil {
		info_message := fmt.Sprintf("%s has a total of %d jobs", query, total_docs)
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

	} else {
		log.Fatal("Could not retrive total number of jobs")
		panic(err)
	}

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

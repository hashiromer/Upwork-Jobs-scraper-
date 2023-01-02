package upwork

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type UpworkPipeLine struct {
	upworkClient *Upwork
	iterations   int
}

func InitPipeline() *UpworkPipeLine {
	u := UpworkPipeLine{
		upworkClient: InitUpwork(),
		iterations:   32,
	}
	return &u
}

func (u *UpworkPipeLine) CombineFiles() error {

	var all_jobs []interface{}
	all_filenames, err := filepath.Glob("data/*.json")
	if err != nil {
		panic(err)
	}

	for _, file := range all_filenames {
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
	for iteration = 1; iteration <= u.iterations; iteration++ {
		log.Println("Iteration: ", iteration)
		urlArgs := UrlArgs{
			Page:     iteration,
			Per_Page: 50,
			Query:    query,
		}
		//It is possible to use a go routine here but be nice to the api or you will be rate limited pretty quickly. It is technically possible to circumvent it using a proxy but it is not recommended.
		// go u.handleRequest(urlArgs, i)
		u.handleRequest(urlArgs, iteration)

	}

	err = u.CombineFiles()
	os.RemoveAll("data")

	if err != nil {
		panic(err)

	}

	return nil

}

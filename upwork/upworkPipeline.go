package upwork

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type UpworkPipeLine struct {
	upworkClient *Upwork
}

func InitPipeline() *UpworkPipeLine {
	return &UpworkPipeLine{
		upworkClient: InitUpwork(),
	}
}

func (u *UpworkPipeLine) appendToJSONL(data []interface{}, filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, item := range data {
		jsonLine, err := json.Marshal(item)
		if err != nil {
			return err
		}
		if _, err := file.Write(append(jsonLine, '\n')); err != nil {
			return err
		}
	}

	return nil
}

func (u *UpworkPipeLine) isResponseValid(data string) bool {
	var resp map[string]interface{}
	err := json.Unmarshal([]byte(data), &resp)
	if err != nil {
		return false
	}
	_, hasErrors := resp["errors"]
	return !hasErrors
}

func (u *UpworkPipeLine) getTotalDocuments(query string, variables map[string]interface{}) (int, error) {
	resp, err := u.upworkClient.SendRequest(query, variables)
	if err != nil {
		return 0, err
	}

	if !u.isResponseValid(resp) {
		return 0, fmt.Errorf("invalid response")
	}

	var graphQLResp map[string]interface{}
	err = json.Unmarshal([]byte(resp), &graphQLResp)
	if err != nil {
		return 0, err
	}

	paging, ok := graphQLResp["data"].(map[string]interface{})["search"].(map[string]interface{})["universalSearchNuxt"].(map[string]interface{})["userJobSearchV1"].(map[string]interface{})["paging"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("unexpected response structure")
	}

	total, ok := paging["total"].(float64)
	if !ok {
		return 0, fmt.Errorf("total is not a number")
	}

	return int(total), nil
}

func (u *UpworkPipeLine) handleRequest(query string, variables map[string]interface{}, filename string) error {
	resp, err := u.upworkClient.SendRequest(query, variables)
	if err != nil {
		return err
	}

	if !u.isResponseValid(resp) {
		//print the response
		fmt.Println(resp)
		return fmt.Errorf("invalid response returned")

	}

	var graphQLResp map[string]interface{}
	err = json.Unmarshal([]byte(resp), &graphQLResp)
	if err != nil {
		return err
	}

	results, ok := graphQLResp["data"].(map[string]interface{})["search"].(map[string]interface{})["universalSearchNuxt"].(map[string]interface{})["userJobSearchV1"].(map[string]interface{})["results"].([]interface{})
	if !ok {
		return fmt.Errorf("unexpected response structure")
	}

	return u.appendToJSONL(results, filename)
}

func (u *UpworkPipeLine) Run(userQuery string) error {
	query, err := readGraphQLQuery("JobSearchQuery.gql")
	if err != nil {
		return err
	}

	variables := map[string]interface{}{
		"requestVariables": map[string]interface{}{
			"userQuery": userQuery,
			"sort":      "recency",
			"highlight": true,
			"paging": map[string]interface{}{
				"offset": 0,
				"count":  50,
			},
		},
	}

	total_docs, err := u.getTotalDocuments(query, variables)
	if err != nil {
		return err
	}

	fmt.Printf("%s has a total of %d jobs\n", userQuery, total_docs)

	iterations := total_docs / 50
	// if iterations > 100 {
	// 	iterations = 100
	// }

	fmt.Printf("A total of %d iterations will be performed\n", iterations)

	now := time.Now()
	dateString := now.Format("2006-01-02")
	filename := fmt.Sprintf("upwork_jobs_%s.jsonl", dateString)

	for i := 0; i < iterations; i++ {
		fmt.Printf("Processing iteration %d of %d\n", i+1, iterations)
		variables["requestVariables"].(map[string]interface{})["paging"].(map[string]interface{})["offset"] = i * 50
		err := u.handleRequest(query, variables, filename)
		if err != nil {
			log.Printf("Error in iteration %d: %v", i+1, err)
		}
	}

	fmt.Printf("Job data has been written to %s\n", filename)
	return nil
}
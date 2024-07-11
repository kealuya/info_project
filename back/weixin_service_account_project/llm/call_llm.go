package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type RequestData struct {
	Inputs struct {
		Query      string `json:"query"`
		WordCount  int    `json:"word_count"`
		ThemeCount int    `json:"theme_count"`
	} `json:"inputs"`
	ResponseMode string `json:"response_mode"`
	User         string `json:"user"`
}

type ResponseData struct {
	WorkflowRunID string `json:"workflow_run_id"`
	TaskID        string `json:"task_id"`
	Data          struct {
		ID         string `json:"id"`
		WorkflowID string `json:"workflow_id"`
		Status     string `json:"status"`
		Outputs    struct {
			Output string `json:"output"`
		} `json:"outputs"`
		Error       interface{} `json:"error"`
		ElapsedTime float64     `json:"elapsed_time"`
		TotalTokens int         `json:"total_tokens"`
		TotalSteps  int         `json:"total_steps"`
		CreatedAt   int64       `json:"created_at"`
		FinishedAt  int64       `json:"finished_at"`
	} `json:"data"`
}

func CallLlm(query string) string {
	url := "http://122.9.41.45:9111/v1/workflows/run"
	requestData := RequestData{
		ResponseMode: "blocking",
		User:         "abc-123",
	}
	requestData.Inputs.Query = query
	//requestData.Inputs.WordCount = 8000
	//requestData.Inputs.ThemeCount = 8

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Printf("Error marshalling request data: %v\n", err)

	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)

	}

	req.Header.Set("Authorization", "Bearer app-ccNX6lmaqAeD6U9VXEcQNwWD")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)

	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		fmt.Printf("Response body: %s\n", body)

	}

	var responseData ResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Printf("Error unmarshalling response data: %v\n", err)

	}

	fmt.Printf("Workflow Run ID: %s\n", responseData.WorkflowRunID)
	fmt.Printf("Task ID: %s\n", responseData.TaskID)
	fmt.Printf("Status: %s\n", responseData.Data.Status)
	fmt.Printf("Output Text: %s\n", responseData.Data.Outputs.Output)
	fmt.Printf("Elapsed Time: %f seconds\n", responseData.Data.ElapsedTime)
	fmt.Printf("Total Tokens: %d\n", responseData.Data.TotalTokens)
	fmt.Printf("Total Steps: %d\n", responseData.Data.TotalSteps)
	fmt.Printf("Created At: %d\n", responseData.Data.CreatedAt)
	fmt.Printf("Finished At: %d\n", responseData.Data.FinishedAt)
	if responseData.Data.Error != nil {
		fmt.Printf("Error: %v\n", responseData.Data.Error)
	}
	return fmt.Sprintf("%s", responseData.Data.Outputs.Output)
}

// decodeUnicode 使用 strconv.Unquote 简化转换
func decodeUnicode(s string) (string, error) {
	// 将字符串中的 Unicode 转移前缀（例如 \uXXXX）表示的内容还原
	out, err := strconv.Unquote(`"` + s + `"`)
	if err != nil {
		return "", err
	}
	return out, nil
}

package llm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"net/http"
	"strconv"
	"strings"
	"weixin_service_account_project/common"
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

func CallLlm(query string) (answer string, resultError error) {
	appKey := "app-ccNX6lmaqAeD6U9VXEcQNwWD"
	url := "http://122.9.41.45:9111/v1/workflows/run"
	requestData := RequestData{
		ResponseMode: "blocking",
		User:         "abc-123",
	}
	requestData.Inputs.Query = query
	//requestData.Inputs.WordCount = 8000
	//requestData.Inputs.ThemeCount = 8

	defer common.RecoverHandler(func(err error) {
		answer = ""
		resultError = err
		return
	})

	jsonData, err := json.Marshal(requestData)
	common.ErrorHandler(err)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	common.ErrorHandler(err)

	req.Header.Set("Authorization", "Bearer "+appKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	common.ErrorHandler(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	common.ErrorHandler(err)

	if resp.StatusCode != http.StatusOK {
		common.ErrorHandler(errors.New(fmt.Sprintf("Error: received status code %d\n . Response body: %s\n", resp.StatusCode, body)))
	}

	var responseData ResponseData
	err = json.Unmarshal(body, &responseData)
	common.ErrorHandler(err)

	// 大模型返回状态记录
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Workflow Run ID: %s\n", responseData.WorkflowRunID))
	sb.WriteString(fmt.Sprintf("Task ID: %s\n", responseData.TaskID))
	sb.WriteString(fmt.Sprintf("Status: %s\n", responseData.Data.Status))
	sb.WriteString(fmt.Sprintf("Output Text: %s\n", responseData.Data.Outputs.Output))
	sb.WriteString(fmt.Sprintf("Elapsed Time: %.2f seconds\n", responseData.Data.ElapsedTime))
	sb.WriteString(fmt.Sprintf("Total Tokens: %d\n", responseData.Data.TotalTokens))
	sb.WriteString(fmt.Sprintf("Total Steps: %d\n", responseData.Data.TotalSteps))
	sb.WriteString(fmt.Sprintf("Created At: %d\n", responseData.Data.CreatedAt))
	sb.WriteString(fmt.Sprintf("Finished At: %d\n", responseData.Data.FinishedAt))
	logs.Info(sb.String())

	if responseData.Data.Error != nil {
		common.ErrorHandler(errors.New(fmt.Sprintf("Error: %v\n", responseData.Data.Error)))
	}
	return fmt.Sprintf("%s", responseData.Data.Outputs.Output), nil
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

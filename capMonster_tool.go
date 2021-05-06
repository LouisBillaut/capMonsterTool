package capMonsterTool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

//CapMonster API URL
const baseURL = "https://api.capmonster.cloud"

type GetBalanceResponse struct {
	//Error identificator.
	//0 - no error, errorCode property missing
	//1 - error, information about it is in the errorCode property
	ErrorId int
	//The number of money available
	Balance float64
	//Error code.
	ErrorCode string
	//Error description.
	ErrorDescription string
}

type CreateTaskResponse struct {
	//Error identificator.
	//0 - no error, errorCode property missing
	//1 - error, information about it is in the errorCode property
	ErrorId int
	//Task ID for future use in getTask method.
	TaskId int
	//Error code.
	ErrorCode string
	//Error description.
	ErrorDescription string
}

type Solution struct {
	//gRecaptcha response
	GRecaptchaResponse string
	//processing - task is not ready yet
	//ready - task complete, solution object can be found in solution property
	Status string
}

type GetTaskResultResponse struct {
	//Error identificator.
	//0 - no error, errorCode property missing
	//1 - error, information about it is in the errorCode property
	ErrorId int
	//processing - task is not ready yet
	//ready - task complete, solution object can be found in solution property
	Status string
	//Captcha solution
	Solution Solution
	//Error code.
	ErrorCode string
	//Error description.
	ErrorDescription string
}

//CapMonsterError struct error
type CapMonsterError struct {
	//errorCode return by CapMonster API
	Err  string
	//errorDescription return by CapMonster API
	Desc string
}

type GetBalancePayload struct {
	//Unique key of your account
	ClientKey string `json:"clientKey"`
}

type GetTaskResultPayload struct {
	//Unique key of your account
	ClientKey string `json:"clientKey"`
	//ID which was obtained in createTask method.
	TaskId int `json:"taskId"`
}

type CreateTaskPayload struct {
	//Unique key of your account
	ClientKey string `json:"clientKey"`
	//Task data
	Task interface{} `json:"task"`
	//not required
	//Optional web address where we will send result of captcha task processing. Contents are sent by POST request and are same to the contents of getTaskResult method. The content of the response is not checked and you must  accept the request in 2 seconds then the connection will be closed.
	CallbackUrl string `json:"callbackUrl"`
}

// close properly the body
func bodyCloser(c io.Closer) {
	if err := c.Close(); err != nil {
		fmt.Errorf("error: %s", err)
	}
}

//do a request with the specified method, url and payload
func doRequest(method string, url string, payload []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer bodyCloser(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//CreateTask creates a task for solving selected captcha type.
//In the parameters you need to pass the client authorization data and typed task data
//Method address: https://api.capmonster.cloud/createTask
//Request format: JSON POST
func CreateTask(clientKey string, task interface{}) (int, error) {
	url := baseURL + "/createTask"
	payload, err := json.Marshal(CreateTaskPayload{ClientKey: clientKey, Task: task})
	if err != nil {
		return 0, err
	}
	body, err := doRequest("POST", url, payload)
	if err != nil {
		return 0, err
	}
	var createTaskResponse CreateTaskResponse
	if err = json.Unmarshal(body, &createTaskResponse); err != nil {
		return 0, err
	}
	if createTaskResponse.ErrorId != 0 {
		return 0, CapMonsterError{Err: createTaskResponse.ErrorCode, Desc: createTaskResponse.ErrorDescription}
	}
	return createTaskResponse.TaskId, nil
}

//CreateTaskWithCallbackUrl creates a task for solving selected captcha type.
//In the parameters you need to pass the client authorization data, typed task data and callback url
//Method address: https://api.capmonster.cloud/createTask
//Request format: JSON POST
func CreateTaskWithCallbackUrl(clientKey string, task interface{}, callbackUrl string) (int, error) {
	url := baseURL + "/createTask"
	payload, err := json.Marshal(CreateTaskPayload{ClientKey: clientKey, Task: task, CallbackUrl: callbackUrl})
	if err != nil {
		return 0, err
	}
	body, err := doRequest("POST", url, payload)
	if err != nil {
		return 0, err
	}
	var createTaskResponse CreateTaskResponse
	if err = json.Unmarshal(body, &createTaskResponse); err != nil {
		return 0, err
	}
	if createTaskResponse.ErrorId != 0 {
		return 0, CapMonsterError{Err: createTaskResponse.ErrorCode, Desc: createTaskResponse.ErrorDescription}
	}
	return createTaskResponse.TaskId, nil
}

//GetTaskResult get the result of the specified task
//In the parameters you need to pass your client key and the ID of your task.
//Method address: https://api.capmonster.cloud/getTaskResult
//Request format: JSON POST
//Request limit: 120 requests per task.
func GetTaskResult(clientKey string, taskId int) (string, error) {
	url := baseURL + "/getTaskResult"
	payload, err := json.Marshal(GetTaskResultPayload{ClientKey: clientKey, TaskId: taskId})
	if err != nil {
		return "", err
	}
	body, err := doRequest("POST", url, payload)
	if err != nil {
		return "", err
	}
	var getTaskResultResponse GetTaskResultResponse
	if err = json.Unmarshal(body, &getTaskResultResponse); err != nil {
		return "", err
	}
	if getTaskResultResponse.ErrorId != 0 {
		return "", CapMonsterError{Err: getTaskResultResponse.ErrorCode, Desc: getTaskResultResponse.ErrorDescription}
	}
	if getTaskResultResponse.Status == "processing" {
		return "processing", CapMonsterProcessing{}
	}
	return getTaskResultResponse.Solution.GRecaptchaResponse, nil
}

//GetBalance retrieve your actual account balance
//In the parameters you need to pass your client key.
//Method address: https://api.capmonster.cloud/getBalance
//Request format: JSON POST
func GetBalance(clientKey string) (float64, error) {
	url := baseURL + "/getBalance"
	payload, err := json.Marshal(GetBalancePayload{ClientKey: clientKey})
	if err != nil {
		return 0., err
	}
	body, err := doRequest("POST", url, payload)
	if err != nil {
		return 0., err
	}
	var getBalanceResponse GetBalanceResponse
	if err = json.Unmarshal(body, &getBalanceResponse); err != nil {
		return 0., err
	}
	if getBalanceResponse.ErrorId != 0 {
		return 0., CapMonsterError{Err: getBalanceResponse.ErrorCode, Desc: getBalanceResponse.ErrorDescription}
	}
	return getBalanceResponse.Balance, nil
}
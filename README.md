# ✨ CapMonsterTool ✨

## About this module
### What is ✨ CapMonsterTool ✨ ?
CapMonsterTool is a set of Go tools designed to simply make requests to the [CapMonster Cloud](https://capmonster.cloud/en/) API.
You need a [CapMonster Cloud](https://capmonster.cloud/en/) account to use this module.

### Developed with
* [Go](https://golang.org/)

### CapMonster Cloud API documentation
[CapMonster Cloud](https://capmonster.cloud/en/) API documentation is available at this [url](https://zennolab.atlassian.net/wiki/spaces/APIS/pages/491575/English+Documentation).   
I strongly advise to read this documentation before using the module

## Getting started
### import the module
The easiest way to use ths module is to import it:
```go
import (
	"github.com/LouisBillaut/capMonsterTool"
)
```

### main API endpoints
There are 3 main API endpoint to CapMonster API:  
* https://api.capmonster.cloud/createTask
  which allows you to create a captcha task
* https://api.capmonster.cloud/getTaskResult
  which allows you to get a request task result
* https://api.capmonster.cloud/getBalance
  which allows you to retrieve account balance
  
### Examples
The `examples/` folder contains some examples of use of this module.  
Please note that before run example, put your own API key in apiKey field.

### Get account Balance
```go
balance, err := capMonsterTool.GetBalance(apiKey)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(balance)
```

### Create a task
Start by create a Task type. All possible tasks type [here](https://zennolab.atlassian.net/wiki/spaces/APIS/pages/557229/Captcha+Task+Types).  
Here is some simple examples:
```go
task1 := NoCaptchaTaskProxyless{Type: "NoCaptchaTaskProxyless", WebsiteURL: myUrl, WebsiteKey: myWebsiteKey}
task2 := ImageToTextTask{Type: "ImageToTextTask", Body: body, CapMonsterModule: "yandex"}
```

Your can now create a Task:
```go
taskId1, err := capMonsterTool.CreateTask(apiKey, task1)
if err != nil {
	fmt.Println(err)
	return
}

taskId2, err := capMonsterTool.CreateTaskWithCallbackUrl(apiKey, task2, myCallbackUrl)
if err != nil {
    fmt.Println(err)
    return
}
```

### Retrieve Task result
⚠️ Please note that calling `GetTaskResult`
does not ensure you have an immediate captcha solution, if the CapMonsterProcessing error is thrown, try again
```go
captchaRes, err := capMonsterTool.GetTaskResult(apiKey, taskId)
for err != nil {
	if !errors.Is(capMonsterTool.CapMonsterProcessing{}, err) {
		fmt.Println(err)
		return
	}
	fmt.Println("task is processing...")
	time.Sleep(2 * time.Second)
	captchaRes, err = capMonsterTool.GetTaskResult(apiKey, taskId)
}
fmt.Println(captchaRes)
```
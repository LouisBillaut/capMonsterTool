package main

import (
	"errors"
	"fmt"
	"github.com/LouisBillaut/capMonsterTool"
	"time"
)

//this example resolve a Google ReCaptcha2 (NoCaptchaTaskProxyless on CapMonster possible captcha tasks)
//example website is https://www.footpatrol.com/product/nike-air-max-95-og/416522_footpatrolcom/

// Your capMonster API key here
const apiKey = "67b6bcbb1a728ea8d563de6d169a2057"
// Delay in seconds (CapMonster time suggested is 2sec)
const delay = 2

func main() {
	//Resolving a Footpatrol captcha as example 😇
	url := "https://www.footpatrol.com/product/nike-air-max-95-og/416522_footpatrolcom/"
	websiteKey := "6LfEwHQUAAAAACTyJsAICi2Lz2oweGni8QOhV-Yl"
	fmt.Println("url is:", url)
	task := capMonsterTool.NoCaptchaTaskProxyless{Type: "NoCaptchaTaskProxyless", WebsiteURL: url, WebsiteKey: websiteKey}
	taskId, err := capMonsterTool.CreateTask(apiKey, task)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("My task id:", taskId)
	fmt.Println("let's resolve the captcha...")
	captchaRes, err := capMonsterTool.GetTaskResult(apiKey, taskId)
	for err != nil {
		if !errors.Is(capMonsterTool.CapMonsterProcessing{}, err) {
			fmt.Println(err)
			return
		}
		fmt.Println("task is processing...")
		time.Sleep(delay * time.Second)
		captchaRes, err = capMonsterTool.GetTaskResult(apiKey, taskId)
	}
	fmt.Println("captcha solution:", captchaRes)
}

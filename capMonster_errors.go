package capMonsterTool

import "fmt"

//this file reference all capMonsterTool module errors

//CapMonsterProcessing struct error
type CapMonsterProcessing struct {
}

func (e CapMonsterError) Error() string {
	return fmt.Sprintf("CapMonster, %s: %s", e.Err, e.Desc)
}

func (e CapMonsterProcessing) Error() string {
	return "CapMonster, captcha is processing..."
}

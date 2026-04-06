package utils

import "fmt"

type DivideError struct {
	Code int
	Message string
	Param int
}


func (e  *DivideError) Error() string{
	return fmt.Sprintf("错误码：%d，描述：%s，出错参数：%d", e.Code, e.Message, e.Param)
}


func Divide2(a, b int) (int, error){
	if b ==0 {
		return 0, &DivideError{
			Code: 1001,
			Message: "除数不能未0",
			Param: b,
		}
	}
	return a / b, nil
}
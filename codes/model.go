package codes

import "fmt"

type Code struct {
	Code    uint
	Message string
}

func (curr Code) String() string {
	return fmt.Sprintf("Code:%v,Message:%v", curr.Code, curr.Message)
}

func (curr Code) Error() string {
	return curr.String()
}

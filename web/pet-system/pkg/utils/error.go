package utils


type error interface{
	Error() string // 仅需实现一个返回错误描述的方法
}
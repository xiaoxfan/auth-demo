package result

type Result struct {
	Result  int         `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const SuccessResult = 0
const ErrorResult = 0
const SuccessMessage = "success"

func Success(data interface{}) interface{} {
	return Result{
		Result:  SuccessResult,
		Message: SuccessMessage,
		Data:    data,
	}
}
func Error(data interface{}, message string) interface{} {
	return Result{
		Result:  ErrorResult,
		Message: message,
		Data:    data,
	}
}

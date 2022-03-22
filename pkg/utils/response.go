/**
 **
 * @struct Result
 **
 * @ReturnErrorResult() Returns an error Result objects
 * @ReturnSuccessResult() Returns a success result with a message
 * @ReturnSuccessMessage() Returns only a success message
 * @ReturnValidateError() Returns a validation error message
 * @ReturnBasicResult() Returns an object only
 * @ReturnAuthResult() Returns an authenticated object with token
**/

package utils

// Result struct for creating response
type Result struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Token   string      `json:"token,omitempty"`
	Count   int64       `json:"count,omitempty"`
}

// ReturnErrorResult return error object
func (r *Result) ReturnErrorResult(error string) Result {
	return Result{
		Success: false,
		Error:   error,
	}
}

// ReturnSuccessResult return success object
func (r *Result) ReturnSuccessResult(data interface{}, message string) Result {
	return Result{
		Success: true,
		Data:    data,
		Message: message,
	}
}

// ReturnSuccessMessage returns just success message
func (r *Result) ReturnSuccessMessage(message string) Result {
	return Result{
		Success: true,
		Message: message,
	}
}

// ReturnBasicResult returns just object without message
func (r *Result) ReturnBasicResult(data interface{}) Result {
	return Result{
		Success: true,
		Data:    data,
	}
}

// ReturnAuthResult returns both object and token
func (r *Result) ReturnAuthResult(data interface{}, token string) Result {
	return Result{
		Success: true,
		Token:   token,
		Data:    data,
	}
}

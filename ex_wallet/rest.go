package ex_wallet

import "errors"

// CommonResponse is a generic response structure.
type CommonResponse[T any] struct {
	Rc     int    `json:"rc"`     // return code, 0=成功; 1=失败
	Mc     string `json:"mc"`     // message code, for internationalization
	Ma     []any  `json:"ma"`     // message args, for internationalization
	Result T      `json:"result"` // returned data
}

// Predefined constants for success and failure codes and messages.
const (
	RcSuccess = 0
	RcFailure = 1
	McSuccess = "SUCCESS"
	McFailure = "FAILURE"
)

// Predefined success and failure responses.
var (
	SuccessResponse = CommonResponse[any]{
		Rc: RcSuccess,
		Mc: McSuccess,
		Ma: []any{},
	}
	FailureResponse = CommonResponse[any]{
		Rc: RcFailure,
		Mc: McFailure,
		Ma: []any{},
	}
)

// NewCommonResponse creates a new CommonResponse instance.
func NewCommonResponse[T any](rc int, mc string, ma []any, result T) CommonResponse[T] {
	return CommonResponse[T]{
		Rc:     rc,
		Mc:     mc,
		Ma:     ma,
		Result: result,
	}
}

// Success creates a success response with custom result.
func Success[T any](result T) CommonResponse[T] {
	return NewCommonResponse(RcSuccess, McSuccess, nil, result)
}

// Failure creates a failure response with an error message.
func Failure(err error) CommonResponse[any] {
	if err == nil {
		err = errors.New("unknown error")
	}
	return NewCommonResponse[any](RcFailure, err.Error(), nil, nil)
}

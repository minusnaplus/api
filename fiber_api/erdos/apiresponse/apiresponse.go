package apiresponse

import (
	"fmt"
	"strconv"
	"math/big"
)

type ApiResponse struct {
	Success bool        `json:"success"`
	Response string     `json:"response"`
	Data     interface{} `json:"data"`
}

func CreateApiResponse(success bool, response string, data interface{}) *ApiResponse {
	return &ApiResponse{Success: success, Response: response, Data: data}
}

func ParseInt64(str string) (int64, error) {
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("conversion error for %s: %v", str, err)
	}
	return val, nil
}

func ParseBigInt(str string) (*big.Int, error) {
	val := big.NewInt(0)
	_, success := val.SetString(str, 10)
	if !success {
		return nil, fmt.Errorf("conversion error for %s", str)
	}
	return val, nil
}
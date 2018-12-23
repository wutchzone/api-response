package response

import (
	"encoding/json"
	"net/http"
)

// Status is an enum for response
type Status int

const (
	// ResponseOK stands for succesfull response
	ResponseOK Status = 1
	// ResponseError stands for unsuccesfull response
	ResponseError Status = 2
)

var responseStatus = map[Status]string{
	ResponseOK:    "OK",
	ResponseError: "Error",
}

var responseMessage = map[Status]string{
	ResponseOK:    "",
	ResponseError: "Something went wrong",
}

// Response HTTP
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// CreateResponse entity
func CreateResponse(w http.ResponseWriter, status Status, dato interface{}) error {
	r := Response{
		Status:  responseStatus[status],
		Message: responseMessage[status],
		Data:    dato,
	}

	d, err := json.Marshal(r)

	if err != nil {
		return err
	}

	w.Write(d)
	return nil
}

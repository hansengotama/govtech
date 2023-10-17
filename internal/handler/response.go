package handler

import (
	"encoding/json"
	"net/http"
)

type HTTPResponse struct {
	Data any              `json:"data"`
	Code int              `json:"code"`
	Err  *HTTPResponseErr `json:"err"`
}

type HTTPResponseErr struct {
	Message string `json:"message"`
}

func writeResponse(w http.ResponseWriter, resp HTTPResponse) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(resp.Code)

	byt, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(byt)
	if err != nil {
		panic(err)
	}
}

func constructSuccessResp(data any, code int) HTTPResponse {
	return HTTPResponse{
		Data: data,
		Code: code,
	}
}

func constructErrResponse(code int, message string) HTTPResponse {
	return HTTPResponse{
		Code: code,
		Err: &HTTPResponseErr{
			Message: message,
		},
	}
}

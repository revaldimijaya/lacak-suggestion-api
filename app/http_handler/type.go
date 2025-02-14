package httphandler

import "github/revaldimijaya/lacak-api/app/usecase"

type HttpResponse struct {
	ResultStatus ResultStatus `json:"result_status"`
	Data         interface{}  `json:"data"`
}

type ResultStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type HttpHandler struct {
	Usecase usecase.Usecase
}

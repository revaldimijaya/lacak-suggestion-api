package httphandler

import "github/revaldimijaya/lacak-api/app/usecase"

func InitHTTPHandler(
	Usecase usecase.Usecase,
) HttpHandler {
	return HttpHandler{
		Usecase: Usecase,
	}
}

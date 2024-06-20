package httphelper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bajra-manandhar17/personal-finance-app/internal/excep"
	"github.com/go-playground/validator/v10"
)

type HttpResponse struct {
	Body              string
	StatusCode        int
	Headers           map[string]string
	MultiValueHeaders map[string][]string
	IsBase64Encoded   bool
}

func convertStrToStruct(str string, placeholderStructObj interface{}) error {
	return json.Unmarshal([]byte(str), placeholderStructObj)
}

func validateStruct(structToValidate interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(structToValidate); err != nil {
		return excep.NewInvalidPayload(err.Error())
	}

	return nil
}

func MapAndValidateBody(body io.Reader, reqBodyStruct interface{}) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("failed to read request body: %w", err)
	}

	if err := convertStrToStruct(string(bodyBytes), reqBodyStruct); err != nil {
		return fmt.Errorf("failed to convert string to struct: %w", err)
	}

	if err := validateStruct(reqBodyStruct); err != nil {
		return fmt.Errorf("failed to validate struct: %w", err)
	}

	return nil
}

func MapErrorToApiResponse(err error) (HttpResponse, error) {
	httpExcep := excep.MapErrorToHttpException(err)
	httpExcepBytes, err := json.Marshal(httpExcep)
	if err != nil {
		return HttpResponse{}, err
	}

	return HttpResponse{
		StatusCode: httpExcep.Status,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(httpExcepBytes),
	}, nil
}

func PrepareApiResponse(resData interface{}) (HttpResponse, error) {
	jsonRes, err := json.Marshal(resData)
	if err != nil {
		return MapErrorToApiResponse(err)
	}

	if resData == nil {
		return HttpResponse{
			StatusCode: http.StatusNoContent,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: string(jsonRes),
		}, nil
	}

	return HttpResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jsonRes),
	}, nil
}

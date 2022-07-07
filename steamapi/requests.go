package steamapi

import (
	"github.com/nlutterman/zmodinigen/errors"
	"net/http"
	"net/url"
)

type APIRequest interface {
	Exec(endpointMap EndpointMap) (*http.Response, *errors.AppError)
}

type FormData interface {
	GetRequestData() url.Values
}

type Request struct {
	endpoint EndpointID
	data     url.Values
}

func (req *Request) exec(method, url string, data url.Values) (*http.Response, *errors.AppError) {
	var response *http.Response
	var appErr *errors.AppError
	var err error

	switch method {
	case http.MethodGet:
		if len(data) > 0 {
			url = url + "?" + data.Encode()
		}
		response, err = http.Get(url)
	default:
		response, err = http.PostForm(url, data)
	}

	if err != nil {
		appErr = errors.NewError(errors.ErrorHTTPRequest, "error performing HTTP request: %v", err)
	} else {
		appErr = CheckStatusCode(response)
	}

	return response, appErr
}

// CheckStatusCode returns an error if the HTTP status code of the response is not within the successful range
func CheckStatusCode(response *http.Response) *errors.AppError {
	if response.StatusCode < 200 || response.StatusCode > 299 {
		return errors.NewError(errors.ErrorHTTPRequest, "non OK HTTP response: %s", response.Status)
	}
	return nil
}

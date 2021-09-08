package helpers

import (
	"context"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// ResponseJSON define the structure of all response
type ResponseJSON struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   *ErrorJSON  `json:"error,omitempty"`
}

// ErrorJSON define the structure of an error
type ErrorJSON struct {
	Message      string `json:"message"`          // for developer
	Reason       string `json:"reason"`           //
	ErrTittleMsg string `json:"error_user_title"` // for user
	ErrBodyMsg   string `json:"error_user_msg"`   // for user
}

// WriteHTTPResponse into the response writer, according to the response code and headers.
// headerMap and data argument are both optional
func WriteHTTPResponse(ctx context.Context, w http.ResponseWriter, httpRespCode int, message string, headerMap map[string]string, data interface{}, errors *ErrorJSON) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpRespCode)
	for k, v := range headerMap {
		w.Header().Add(k, v)
	}

	rJSON := &ResponseJSON{
		Status:  httpRespCode,
		Message: message,
		Data:    data,
		Error:   errors,
	}

	if httpRespCode >= 200 && httpRespCode <= 299 {
		if len(rJSON.Message) == 0 {
			rJSON.Message = "SUCCESS"
		}
	} else {
		if errors == nil {
			rJSON.Error = &ErrorJSON{
				Message:      http.StatusText(httpRespCode),
				Reason:       "",
				ErrTittleMsg: http.StatusText(httpRespCode),
				ErrBodyMsg:   http.StatusText(httpRespCode),
			}
		}

		if len(rJSON.Error.ErrTittleMsg) == 0 {
			rJSON.Error.ErrTittleMsg = http.StatusText(httpRespCode)
		}

		if len(rJSON.Error.ErrBodyMsg) == 0 {
			rJSON.Error.ErrBodyMsg = rJSON.Error.ErrTittleMsg
		}

		if len(rJSON.Message) == 0 {
			rJSON.Message = "FAIL"
		}
	}
	bytes, err := json.Marshal(rJSON)
	if err != nil {
		log.Errorf("Can not marshal. Got %s", err)
	} else {
		i, err := w.Write(bytes)
		if err != nil {
			log.Errorf("Can not write byte stream. Got %s. %d bytes written", err, i)
		}
	}
}

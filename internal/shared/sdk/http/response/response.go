package response

import (
	"encoding/json"
	"maps"
	"net/http"

	"github.com/cristiano-pacheco/goflix/internal/shared/modules/errs"
)

func Error(w http.ResponseWriter, err error) {
	rError, ok := err.(*errs.Error)
	if !ok {
		// If it's not our custom error type, convert it to a generic error
		httpStatus := http.StatusInternalServerError
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpStatus)
		json.NewEncoder(w).Encode(Envelope{
			"error": map[string]string{
				"code":    "internal_server_error",
				"message": "Internal server error",
			},
		})
		return
	}

	if rError.Status == 0 {
		rError.Status = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rError.Status)
	json.NewEncoder(w).Encode(rError)
}

func JSON(w http.ResponseWriter, status int, envelope Envelope, headers http.Header) error {
	js, err := json.MarshalIndent(envelope, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	maps.Copy(w.Header(), headers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

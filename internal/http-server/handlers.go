package http_server

import (
	"encoding/json"

	"net/http"

	usecases "slack-topic-parser/internal/api/application/usecases"
	custom_types "slack-topic-parser/internal/api/domain/custom-types"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	// Get Healthcheck response
	response := usecases.HealthCheck()

	// Response
	SendWithData(w, response)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	// Create data model
	data := custom_types.PostRequest{}

	// Create decoder with request body data
	decoder := json.NewDecoder(r.Body)

	// Decode and save data
	if err := decoder.Decode(&data); err != nil {
		SendUnprocessable(w)
	} else {
		res := usecases.AddUser(data)
		SendWithData(w, res.Data)
	}
}

func addFirstUserHandler(w http.ResponseWriter, r *http.Request) {
	// Create data model
	data := custom_types.PostRequest{}

	// Create decoder with request body data
	decoder := json.NewDecoder(r.Body)

	// Decode and save data
	if err := decoder.Decode(&data); err != nil {
		SendUnprocessable(w)
	} else {
		res := usecases.AddFirstUser(data)
		SendWithData(w, res.Data)
	}
}

func removeUserHandler(w http.ResponseWriter, r *http.Request) {
	// Create data model
	data := custom_types.PostRequest{}

	// Create decoder with request body data
	decoder := json.NewDecoder(r.Body)

	// Decode and save data
	if err := decoder.Decode(&data); err != nil {
		SendUnprocessable(w)
	} else {
		res := usecases.RemoveUser(data)
		SendWithData(w, res.Data)
	}
}

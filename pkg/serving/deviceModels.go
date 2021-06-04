package serving

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iot-for-all/starling/pkg/models"
	"github.com/iot-for-all/starling/pkg/storing"
)

// ListDeviceModels lists all models.
func ListDeviceModels(w http.ResponseWriter, _ *http.Request) {
	items, err := storing.DeviceModels.List()
	if handleError(err, w) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(items)
	handleError(err, w)
}

// GetDeviceModel gets an existing model by its id.
func GetDeviceModel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	model, err := storing.DeviceModels.Get(id)
	if handleError(err, w) {
		return
	}

	if model == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(model)
	handleError(err, w)
}

// UpsertDeviceModel adds a new or updates an existing device model.
func UpsertDeviceModel(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body)
	if handleError(err, w) {
		return
	}

	var model models.DeviceModel
	err = json.Unmarshal(req, &model)
	if handleError(err, w) {
		return
	}

	err = storing.DeviceModels.Set(&model)
	if handleError(err, w) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&model)
	handleError(err, w)
}

// DeleteDeviceModel deletes an existing device model.
func DeleteDeviceModel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := storing.DeviceModels.Delete(id)
	handleError(err, w)
}

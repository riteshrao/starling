package serving

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iot-for-all/starling/pkg/controlling"
	"github.com/rs/zerolog/log"
)

var (
	controller *controlling.Controller
	config     *Config
)

// StartAdmin starts serving administration API requests.
func StartAdmin(cfg *Config, ctrl *controlling.Controller) {
	config = cfg
	controller = ctrl

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("simulation", listSimulations).Methods(http.MethodGet)
	router.HandleFunc("simulation", upsertSimulation).Methods(http.MethodPut)
	router.HandleFunc("simulation/{id}", getSimulation).Methods(http.MethodGet)
	router.HandleFunc("simulation/{id}", deleteSimulation).Methods(http.MethodDelete)
	router.HandleFunc("simulation/{id}/start", startSimulation).Methods(http.MethodPost)
	router.HandleFunc("simulation/{id}/stop", stopSimulation).Methods(http.MethodPost)
	router.HandleFunc("simulation/{id}/provision/{modelId}/{numDevices}", provisionDevices).Methods(http.MethodPost)
	router.HandleFunc("simulation/{id}/provision", deleteAllDevices).Methods(http.MethodDelete)
	router.HandleFunc("simulation/{id}/provision/{modelId}/{numDevices}", deleteDevices).Methods(http.MethodDelete)
	router.HandleFunc("simulation/{id}/deviceConfig", ListDeviceConfigs).Methods(http.MethodGet)
	router.HandleFunc("simulation/{id}/deviceConfig", UpsertDeviceConfig).Methods(http.MethodPut)
	router.HandleFunc("simulation/{id}/deviceConfig/{configId}", GetDeviceConfig).Methods(http.MethodGet)
	router.HandleFunc("simulation/{id}/deviceConfig/{configId}", DeleteDeviceConfig).Methods(http.MethodDelete)

	router.HandleFunc("target", listTargets).Methods(http.MethodGet)
	router.HandleFunc("target", upsertTarget).Methods(http.MethodPut)
	router.HandleFunc("target/{id}", getTarget).Methods(http.MethodGet)
	router.HandleFunc("target/{id}", deleteTarget).Methods(http.MethodDelete)
	router.HandleFunc("target/{id}/device", listTargetDevices).Methods(http.MethodGet)
	router.HandleFunc("target/{id}/device/{deviceId}", getTargetDevice).Methods(http.MethodGet)
	router.HandleFunc("target/{id}/device", upsertTargetDevice).Methods(http.MethodPut)
	router.HandleFunc("target/{id}/device", deleteAllTargetDevices).Methods(http.MethodDelete)
	router.HandleFunc("target/{id}/device/{deviceId}", deleteTargetDevice).Methods(http.MethodDelete)
	router.HandleFunc("target/{id}/models", getTargetModels).Methods(http.MethodGet)
	router.HandleFunc("target/{id}/models", upsertTargetModels).Methods(http.MethodPut)
	router.HandleFunc("target/{id}/models", deleteTargetModels).Methods(http.MethodDelete)

	router.HandleFunc("model", ListDeviceModels).Methods(http.MethodGet)
	router.HandleFunc("model", UpsertDeviceModel).Methods(http.MethodPut)
	router.HandleFunc("model/{id}", GetDeviceModel).Methods(http.MethodGet)
	router.HandleFunc("model/{id}", DeleteDeviceModel).Methods(http.MethodDelete)

	log.Info().Msgf("serving admin requests at http://localhost:%d/api", cfg.AdminPort)
	log.Info().Msgf("service starling ux at http://localhost:%d", cfg.AdminPort)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", cfg.AdminPort), router)
}

// handleError log the error and return http error
func handleError(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

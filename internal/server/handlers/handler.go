package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"

	admission "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

type AdmissionHandler interface {
	Handle(admissionReview *admission.AdmissionReview) (*admission.AdmissionResponse, error)
}

func admissionControlRealHandler(admissionReview *admission.AdmissionReview) *admission.AdmissionResponse {
	handler, err := CreateAdmissionHandlerByRequest(admissionReview)

	if err != nil {
		log.Err(err)
	}

	response, err := handler.Handle(admissionReview)

	if err != nil {
		log.Err(err)
	}

	return response
}

// Handles the raw http requests for an admission webhook.
func AdmissionControlerHandler(w http.ResponseWriter, request *http.Request) {
	runtimeScheme := runtime.NewScheme()
	deserializer := serializer.NewCodecFactory(runtimeScheme).UniversalDeserializer()
	var body []byte

	if request.Body != nil {
		if data, err := ioutil.ReadAll(request.Body); err == nil {
			body = data
		}
	}

	// verify the content type is accurate
	contentType := request.Header.Get("Content-Type")
	if contentType != "application/json" {
		log.Error().Msgf("contentType=%s, expect application/json", contentType)
		return
	}

	log.Info().Msgf("handling request: %s", body)
	var responseObject runtime.Object
	if object, groupVersionKind, err := deserializer.Decode(body, nil, nil); err != nil {
		msg := fmt.Sprintf("Request could not be decoded: %v", err)
		log.Error().Msg(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return

	} else {
		requestedAdmissionReview, ok := object.(*admission.AdmissionReview)
		if !ok {
			log.Error().Msgf("Expected v1.AdmissionReview but got: %T", object)
			return
		}
		responseAdmissionReview := &admission.AdmissionReview{}
		responseAdmissionReview.SetGroupVersionKind(*groupVersionKind)
		responseAdmissionReview.Response = admissionControlRealHandler(requestedAdmissionReview)
		responseAdmissionReview.Response.UID = requestedAdmissionReview.Request.UID
		responseObject = responseAdmissionReview

	}
	log.Info().Msgf("sending response: %v", responseObject)
	responseBytes, err := json.Marshal(responseObject)
	if err != nil {
		log.Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(responseBytes); err != nil {
		log.Err(err)
	}
}

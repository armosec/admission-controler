package handlers

import (
	"errors"

	"github.com/armosec/admission-controller/internal/server/handlers/mutators"
	"github.com/armosec/admission-controller/internal/server/handlers/validators"
	admission "k8s.io/api/admission/v1"
)

var resourceToAdmissionValidators = map[string]AdmissionHandler{
	Pod: new(validators.PodValidator),
}

var resourceToAdmissionMutators = map[string]AdmissionHandler{
	Pod: new(mutators.PodMutator),
}

func createHandlerByResource(resource string, handlerType string) (AdmissionHandler, error) {
	switch handlerType {
	case Validator:
		return resourceToAdmissionValidators[resource], nil
	case Mutator:
		return resourceToAdmissionMutators[resource], nil
	default:
		return nil, errors.New("invalid handler type")
	}
}

func getResourceByRequest(admissionReview *admission.AdmissionReview) string {
	return admissionReview.Request.Resource.Resource
}

func CreateAdmissionHandlerByRequest(admissionRequest *AdmissionRequest) (AdmissionHandler, error) {
	resource := getResourceByRequest(admissionRequest.admissionReview)

	return createHandlerByResource(resource, admissionRequest.requestSource)
}

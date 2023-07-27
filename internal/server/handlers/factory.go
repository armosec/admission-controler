package handlers

import (
	"errors"

	"github.com/armosec/admission-controler/internal/server/handlers/validators"
	admission "k8s.io/api/admission/v1"
)

const (
	Validator = "ValidatingWebhookConfiguration"
	Mutator   = "MutatingWebhookConfiguration"
)

var resourceToAdmissionValidators = map[string]AdmissionHandler{
	Pod: new(validators.PodValidator),
}

func createHandlerByResource(resource string, handlerType string) (AdmissionHandler, error) {
	switch handlerType {
	case Validator:
		return resourceToAdmissionValidators[resource], nil
	case Mutator:
		return nil, nil
	default:
		return nil, errors.New("invalid handler type")
	}
}

func getResourceByRequest(admissionReview *admission.AdmissionReview) string {
	return admissionReview.Request.Resource.Resource
}

func CreateAdmissionHandlerByRequest(admissionReview *admission.AdmissionReview) (AdmissionHandler, error) {
	resource := getResourceByRequest(admissionReview)

	return createHandlerByResource(resource, admissionReview.Kind)
}

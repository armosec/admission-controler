package validators

import (
	"encoding/json"

	"github.com/armosec/armo-admission-controller/internal/alertmanager"
	"github.com/rs/zerolog/log"
	admission "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodValidator struct{}

func getPodFromAdmissionReview(admissionReview *admission.AdmissionReview) (*corev1.Pod, error) {
	rawObject := admissionReview.Request.Object.Raw
	pod := corev1.Pod{}
	if err := json.Unmarshal(rawObject, &pod); err != nil {
		return nil, err
	}

	return &pod, nil
}

func (podValidator *PodValidator) Handle(admissionReview *admission.AdmissionReview) (*admission.AdmissionResponse, error) {
	log.Info().Msgf("Validating pod")

	switch admissionReview.Request.SubResource {
	case "exec":
		return podValidator.handleExec(admissionReview), nil
	default:
		return &admission.AdmissionResponse{Allowed: true}, nil
	}

}

func (podValidator *PodValidator) handleExec(admissionReview *admission.AdmissionReview) *admission.AdmissionResponse {
	pod, err := getPodFromAdmissionReview(admissionReview)

	if err != nil {
		return &admission.AdmissionResponse{
			Result: &metav1.Status{
				Message: err.Error(),
			},
		}
	}

	rawObject := admissionReview.Request.Object.Raw
	execInfo := corev1.PodExecOptions{}
	if err := json.Unmarshal(rawObject, &execInfo); err != nil {
		return &admission.AdmissionResponse{
			Result: &metav1.Status{
				Message: err.Error(),
			},
		}
	}

	alertInfo := alertmanager.AlertInfo{
		Name:        "Pod Exec",
		Severity:    "High",
		Instance:    pod.GetName(),
		Namespace:   pod.GetNamespace(),
		Description: execInfo.Command[0],
	}
	alertmanager.Alert(&alertInfo)

	return &admission.AdmissionResponse{Allowed: true}
}

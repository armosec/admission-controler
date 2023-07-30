package validators

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
	admission "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodValidator struct{}

func (podValidator *PodValidator) Handle(admissionReview *admission.AdmissionReview) (*admission.AdmissionResponse, error) {
	log.Info().Msgf("Validating pod")
	raw := admissionReview.Request.Object.Raw
	pod := corev1.Pod{}
	if err := json.Unmarshal(raw, &pod); err != nil {
		log.Err(err)
		return &admission.AdmissionResponse{
			Result: &metav1.Status{
				Message: err.Error(),
			},
		}, nil
	}

	log.Info().Msgf("Pod's name is: %s\n", pod.GetName())

	return &admission.AdmissionResponse{Allowed: true}, nil
}

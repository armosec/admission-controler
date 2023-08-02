package alertmanager

import (
	"context"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/prometheus/alertmanager/api/v2/client"
	alertapi "github.com/prometheus/alertmanager/api/v2/client/alert"
	"github.com/prometheus/alertmanager/api/v2/models"
	"github.com/rs/zerolog/log"
)

func Alert(alertInfo *AlertInfo) {
	alert := createAlert(alertInfo)

	response, err := sendAlertToAlertmanager(alert)
	if err != nil {
		log.Error().Msgf("Alert manager error: %v", err)
		return
	}

	log.Info().Msgf("Response from alertmanager: %v", response)
}

func createAlert(alertInfo *AlertInfo) *models.PostableAlert {
	alert := &models.PostableAlert{
		Annotations: map[string]string{
			"description": alertInfo.Description,
		},
		Alert: models.Alert{
			Labels: map[string]string{
				"alertname": alertInfo.Name,
				"severity":  alertInfo.Severity,
				"instance":  alertInfo.Instance,
				"namespace": alertInfo.Namespace,
			},
		},
		StartsAt: strfmt.DateTime(time.Now().UTC()),
		EndsAt:   strfmt.DateTime(time.Now().Add(time.Hour).UTC()),
	}

	return alert
}

func sendAlertToAlertmanager(alert *models.PostableAlert) (*alertapi.PostAlertsOK, error) {
	transport := httptransport.New(ALERTMANAGER_HOST, API_PATH, nil)
	alertmanagerClient := client.New(transport, nil)

	postAlertsParams := alertapi.PostAlertsParams{
		Alerts:  []*models.PostableAlert{alert},
		Context: context.Background(),
	}

	response, err := alertmanagerClient.Alert.PostAlerts(&postAlertsParams)
	if err != nil {
		return nil, err
	}

	return response, nil
}

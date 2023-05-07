package usecases

import (
	custom_types "slack-topic-parser/internal/api/domain/custom-types"
)

func HealthCheck() custom_types.HealthCheck {
	r := custom_types.HealthCheck{Status: "OK"}
	return r
}

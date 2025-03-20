package services

import (
	"akatsuki/skeleton-go/src/dto"
)

func HealthService() dto.HealthResponse {

	return dto.HealthResponse{
		Version: "0.1",
		Msg:     "I'm UMAR SAID SEHAT",
	}
}

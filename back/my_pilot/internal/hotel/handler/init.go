package handler

import (
	"my_pilot/internal/hotel/repository"
)

func init() {
	repository.InitDbEngine()
}

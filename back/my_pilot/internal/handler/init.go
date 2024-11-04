package handler

import (
	"my_pilot/internal/repository"
)

func init() {
	repository.InitDbEngine()
}

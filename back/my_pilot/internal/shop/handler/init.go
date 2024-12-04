package handler

import (
	"my_pilot/internal/shop/repository"
)

func init() {
	repository.InitDbEngine()
}

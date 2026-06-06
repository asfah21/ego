package services

import (
	"ego/models"
	"ego/repositories"
)

// GetAllUsers mengambil semua data user untuk dashboard admin
func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

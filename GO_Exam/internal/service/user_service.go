package service

import (
	"errors"

	"project/itStep/internal/repository"
)
func UpdateUserRole(id uint, role string) error{

	allowedRoles := map[string]bool{
		"user" : true,
		"teacher" : true,
		"admin" : true,
	}

	if !allowedRoles[role] {
		return errors.New("Invalid role")
	}
	return repository.UpdateUserRole(id, role)
}
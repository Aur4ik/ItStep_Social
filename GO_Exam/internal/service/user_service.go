package service

import (
	"errors"

	"project/itStep/internal/repository"
)

func UpdateUserRole(id uint, role string) error {
	allowedRoles := map[string]bool{
		"user":    true,
		"teacher": true,
		"admin":   true,
	}

	if !allowedRoles[role] {
		return errors.New("invalid role")
	}

	if role != "admin" {
		currentUser, err := repository.GetUserByID(id)
		if err != nil {
			return errors.New("user not found")
		}
		if currentUser.Role == "admin" {
			count, err := repository.CountAdmins()
			if err != nil {
				return err
			}
			if count <= 1 {
				return errors.New("cannot remove the last admin")
			}
		}
	}

	return repository.UpdateUserRole(id, role)
}

package repository

import "go-scim/pkg/models"

func GetUser(id string) (models.User, error) {

	return models.User{}, nil
}

func UpsertUser(user models.User) error {

	return nil
}

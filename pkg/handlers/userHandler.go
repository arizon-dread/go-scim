package handlers

import (
	"go-scim/internal/repository"
	"go-scim/pkg/errors"
	"go-scim/pkg/models"
	"log"
)

func UpsertUsers(users []models.User) error {

	for _, v := range users {
		if err := repository.UpsertUser(v); err != nil {
			e := errors.ErrorType{
				Type:    1,
				Code:    500,
				Message: "Unable to update user in db",
				Err:     err,
			}
			log.Printf("user not upserted, id: %v, externalId: %v, reason: %w", v.ID, v.ExternalID, e)
			return err
		}
	}
	return nil
}

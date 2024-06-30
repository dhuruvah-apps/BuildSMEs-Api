package utils

import (
	"context"

	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/httpErrors"
	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/logger"
)

// Validate is user from owner of content
func ValidateIsOwner(ctx context.Context, creatorID string, logger logger.Logger) error {
	user, err := GetUserFromCtx(ctx)
	if err != nil {
		return err
	}

	if user.UserID.String() != creatorID {
		logger.Errorf(
			"ValidateIsOwner, userID: %v, creatorID: %v",
			user.UserID.String(),
			creatorID,
		)
		return httpErrors.Forbidden
	}

	return nil
}

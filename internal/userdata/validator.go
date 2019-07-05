package userdata

import (
	"fmt"
	"strings"

	"github.com/stephram/nsm-userdata-service/restapi/models"
)

func validateUser(tokenID string, user *models.User) error {
	if user == nil {
		return fmt.Errorf("invalid UserData")
	}
	if !strings.EqualFold(tokenID, *user.TokenID) {
		return fmt.Errorf("tokenID and UserData.TokenID mismatch")
	}

	return nil
}

package auth

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
)

func BasicAuth(user string, pass string) (interface{}, error) {
	if len(user) == 0 || len(pass) == 0 {
		return nil, fmt.Errorf("invalid credentials")
	}
	return aws.String("authenticated"), nil
}

func ApiKeyAuth(token string) (interface{}, error) {
	if len(token) == 0 {
		return nil, fmt.Errorf("invalid API key")
	}
	return aws.String("ok"), nil
}

package repo

import (
	"os"
	"testing"

	"github.com/artprocessors/nsm-microservice-golang-userdata/internal/utils/ulid"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/artprocessors/nsm-microservice-golang-userdata/restapi/models"

	"github.com/stretchr/testify/assert"
)

var (
	db       *gorm.DB
	userRepo UserRepo
	tokenID  string
)

func TestMain(m *testing.M) {
	db = createDB(true)
	defer db.Close()

	// tokenID = ulid.New()
	// dbUser := DbUser{TokenID: tokenID}
	// db.Create(&dbUser)

	userRepo = New(db)

	os.Exit(m.Run())
}

func TestNew(t *testing.T) {

	t.Run("store and retrieve ok", func(t *testing.T) {
		tokenID := ulid.New()

		var userData *models.User
		var err error

		if userData, err = userRepo.StoreUser(&models.User{
			TokenID: aws.String(tokenID),
		}); err != nil {
			assert.FailNow(t, err.Error())
		}

		assert.NotNil(t, userData)
		assert.Equal(t, tokenID, *userData.TokenID)

		if userData, err = userRepo.FindUser(tokenID); err != nil {
			assert.FailNow(t, err.Error())
		}

		assert.NotNil(t, userData)
		assert.Equal(t, tokenID, *userData.TokenID)

	})

	t.Run("test FindUser", func(t *testing.T) {
		var userData *models.User
		var err error

		if userData, err = userRepo.FindUser(tokenID); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.NotNil(t, userData)
		assert.Equal(t, tokenID, *userData.TokenID)
	})

	t.Run("test StoreConsoleUser", func(t *testing.T) {
		var user *models.User
		var err error

		if user, err = userRepo.StoreUser(&models.User{
			TokenID: aws.String(tokenID),
		}); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.NotNil(t, user)
		assert.Equal(t, tokenID, *user.TokenID)
	})
}

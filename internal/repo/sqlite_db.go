package repo

import (
	"github.com/artprocessors/nsm-microservice-golang-userdata/restapi/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
)

type DbUser struct {
	gorm.Model
	TokenID  string `gorm:"type:varchar(64);unique_index"`
	Name     string `gorm:"type:varchar(128)"`
	GroupID  string
	AvatarID string
	IsChild  bool
}

type DbGameOnResults struct {
	gorm.Model

	TokenID                 string
	BaseballPoints          float64 `gorm:"type:decimal(10, 2)"`
	CycleDistance           float64 `gorm:"type:decimal(10, 2)"`
	HorsePosition           float64 `gorm:"type:decimal(10, 2)"`
	NetballPoints           float64 `gorm:"type:decimal(10, 2)"`
	PressureCookerScore     float64 `gorm:"type:decimal(10, 2)"`
	ReactionTime            float64 `gorm:"type:decimal(10, 2)"`
	RugbyLeagueGoals        float64 `gorm:"type:decimal(10, 2)"`
	RugbyUnionGoals         float64 `gorm:"type:decimal(10, 2)"`
	SoccerGoals             float64 `gorm:"type:decimal(10, 2)"`
	SurfingTime             float64 `gorm:"type:decimal(10, 2)"`
	YouMakeTheRulesVisited  bool    `gorm:"type:decimal(10, 2)"`
	ClassicCatchVideo       string  `gorm:"type:varchar(2048)"`
	YouMakeTheRulesPostcard string  `gorm:"type:varchar(2048)"`
}

type userRepoImpl struct {
	db *gorm.DB
}

func (u *userRepoImpl) FindUser(tokenID string) (*models.User, error) {
	var dbUser DbUser

	query := u.db.Where("token_id = ?", tokenID)
	err := query.Find(&dbUser).Error

	return u.toModelUser(&dbUser), err
}

func (u *userRepoImpl) StoreUser(user *models.User) (*models.User, error) {
	var dbUser DbUser // dbUser := u.toDbUser(user)

	// if u.db.NewRecord(dbUser) {
	// 	err := u.db.FirstOrCreate(dbUser, dbUser).Error
	// 	return u.toModelUser(dbUser), err
	// }
	//
	// err := u.db.Update(dbUser).Error
	// return u.toModelUser(dbUser), err

	err := u.db.Where(DbUser{TokenID: *user.TokenID}).Assign(u.toDbUser(user)).FirstOrCreate(&dbUser).Error
	return u.toModelUser(&dbUser), err
}

func (u *userRepoImpl) FindGameOnResults(tokenID string) (*models.GameOnResults, error) {
	var dbGoRs []DbGameOnResults

	query := u.db.Where("token_id = ?", tokenID)
	err := query.Order("updated_at desc").Limit(255).Find(&dbGoRs).Error

	return u.reduceToLatest(dbGoRs), err
}

func (u *userRepoImpl) StoreGameOnResults(tokenID string, gameOnResults *models.GameOnResults) (*models.GameOnResults, error) {
	dbGoR := u.toDbGameOnResults(gameOnResults)
	dbGoR.TokenID = tokenID

	u.db.NewRecord(dbGoR)
	err := u.db.Create(dbGoR).Error

	return u.toModelGameOnResults(dbGoR), err
}

func (u *userRepoImpl) reduceToLatest(dbGoRs []DbGameOnResults) *models.GameOnResults {
	mGoR := &models.GameOnResults{}

	// Uses latest entry, not the highest
	for i := 0; i < len(dbGoRs); i++ {
		if mGoR.CycleDistance <= 0 {
			mGoR.CycleDistance = dbGoRs[i].CycleDistance
		}
		if mGoR.HorsePosition <= 0 {
			mGoR.HorsePosition = dbGoRs[i].HorsePosition
		}
		if mGoR.NetballPoints <= 0 {
			mGoR.NetballPoints = dbGoRs[i].NetballPoints
		}
		if mGoR.PressureCookerScore <= 0 {
			mGoR.PressureCookerScore = dbGoRs[i].PressureCookerScore
		}
		if mGoR.ReactionTime <= 0 {
			mGoR.ReactionTime = dbGoRs[i].ReactionTime
		}
		if mGoR.RugbyLeagueGoals <= 0 {
			mGoR.RugbyLeagueGoals = dbGoRs[i].RugbyLeagueGoals
		}
		if mGoR.RugbyUnionGoals <= 0 {
			mGoR.RugbyUnionGoals = dbGoRs[i].RugbyUnionGoals
		}
		if mGoR.SoccerGoals <= 0 {
			mGoR.SoccerGoals = dbGoRs[i].SoccerGoals
		}
		if mGoR.SurfingTime <= 0 {
			mGoR.SurfingTime = dbGoRs[i].SurfingTime
		}
		if !mGoR.YouMakeTheRulesVisited {
			mGoR.YouMakeTheRulesVisited = dbGoRs[i].YouMakeTheRulesVisited
		}
		if len(mGoR.YouMakeTheRulesPostcard) == 0 {
			mGoR.YouMakeTheRulesPostcard = dbGoRs[i].YouMakeTheRulesPostcard
		}
		if len(mGoR.ClassicCatchVideo) == 0 {
			mGoR.ClassicCatchVideo = dbGoRs[i].ClassicCatchVideo
		}
		if mGoR.BaseballPoints <= 0 {
			mGoR.BaseballPoints = dbGoRs[i].BaseballPoints
		}
	}
	return mGoR
}

func (u *userRepoImpl) toModelUser(dbUser *DbUser) *models.User {
	return &models.User{
		TokenID:  aws.String(dbUser.TokenID),
		Name:     dbUser.Name,
		AvatarID: dbUser.AvatarID,
		IsChild:  dbUser.IsChild,
		GroupID:  dbUser.GroupID,
	}
}

func (u *userRepoImpl) toDbUser(user *models.User) *DbUser {
	return &DbUser{
		TokenID:  *user.TokenID,
		Name:     user.Name,
		GroupID:  user.GroupID,
		IsChild:  user.IsChild,
		AvatarID: user.AvatarID,
	}
}

func (u *userRepoImpl) toModelGameOnResults(dbGoR *DbGameOnResults) *models.GameOnResults {
	return &models.GameOnResults{
		BaseballPoints:          dbGoR.BaseballPoints,
		ClassicCatchVideo:       dbGoR.ClassicCatchVideo,
		YouMakeTheRulesPostcard: dbGoR.YouMakeTheRulesPostcard,
		YouMakeTheRulesVisited:  dbGoR.YouMakeTheRulesVisited,
		SurfingTime:             dbGoR.SurfingTime,
		SoccerGoals:             dbGoR.SoccerGoals,
		RugbyUnionGoals:         dbGoR.RugbyUnionGoals,
		RugbyLeagueGoals:        dbGoR.RugbyLeagueGoals,
		ReactionTime:            dbGoR.ReactionTime,
		PressureCookerScore:     dbGoR.PressureCookerScore,
		NetballPoints:           dbGoR.NetballPoints,
		HorsePosition:           dbGoR.HorsePosition,
		CycleDistance:           dbGoR.CycleDistance,
	}
}

func (u *userRepoImpl) toDbGameOnResults(gameOnResults *models.GameOnResults) *DbGameOnResults {
	return &DbGameOnResults{
		BaseballPoints:          gameOnResults.BaseballPoints,
		CycleDistance:           gameOnResults.CycleDistance,
		HorsePosition:           gameOnResults.HorsePosition,
		NetballPoints:           gameOnResults.NetballPoints,
		PressureCookerScore:     gameOnResults.PressureCookerScore,
		ReactionTime:            gameOnResults.ReactionTime,
		RugbyLeagueGoals:        gameOnResults.RugbyLeagueGoals,
		RugbyUnionGoals:         gameOnResults.RugbyUnionGoals,
		SoccerGoals:             gameOnResults.SoccerGoals,
		SurfingTime:             gameOnResults.SurfingTime,
		YouMakeTheRulesVisited:  gameOnResults.YouMakeTheRulesVisited,
		YouMakeTheRulesPostcard: gameOnResults.YouMakeTheRulesPostcard,
		ClassicCatchVideo:       gameOnResults.ClassicCatchVideo,
	}
}

func New(db *gorm.DB) UserRepo {
	userRepo := &userRepoImpl{
		db: func() *gorm.DB {
			if db == nil {
				return createDB(true)
			}
			return db
		}(),
	}
	return userRepo
}

func createDB(testMode bool) *gorm.DB {
	return func() *gorm.DB {
		// db, err := gorm.Open("sqlite3", "file:nsmDB?mode=memory&cache=shared")
		db, err := gorm.Open("sqlite3", "./nsmDB.sqlite")
		if err != nil {
			log.WithError(err).Errorf("unable to open database")
			return nil
		}
		db.LogMode(true)
		if testMode {
			db.AutoMigrate(&DbUser{})
			db.AutoMigrate(&DbGameOnResults{})
		}
		db.BlockGlobalUpdate(true)

		log.Infof("opened sqlite database")
		return db
	}()
}

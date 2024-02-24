package repository

import (
	"strconv"

	"github.com/allan7yin/grpc-go/internal/models"
	interfaces "github.com/allan7yin/grpc-go/pkg/v1"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) interfaces.RepoInterface {
	return &Repo{db}
}

// Create
//
// This function creates a new User which was supplied as the argument
func (repo *Repo) Create(user models.User) (models.User, error) {
	err := repo.db.Create(&user).Error

	return user, err
}

// Get
//
// This function returns the user instance which is
// saved on the DB and returns to the usecase
func (repo *Repo) Get(id string) (models.User, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		// Handle error if the string is not a valid integer
		panic(err)
	}
	var user models.User
	er := repo.db.Where("id = ?", intId).First(&user).Error

	return user, er
}

// Update
//
// This function updates the user name to the one specified,
// as we have email, id and name, so name only gets changed
func (repo *Repo) Update(user models.User) error {
	var dbUser models.User
	if err := repo.db.Where("id = ?", user.ID).First(&dbUser).Error; err != nil {
		return err
	}
	dbUser.Name = user.Name
	err := repo.db.Save(dbUser).Error

	return err
}

// Delete
//
// This function creates a the User whose ID was supplied as the argument
func (repo *Repo) Delete(id string) error {
	err := repo.db.Unscoped().Where("id = ?", id).Delete(&models.User{}).Error
	//.Unscoped deletes from db. Simply using .Delete without this provides the row with a `deleted_at`
	// NEXT: Instead of hard deletes, soft delete and consider moving to `deleted users` table

	return err
}

// GetByEmail
//
// This function returns the user instance which is
// saved on the DB and returns to the usecase
func (repo *Repo) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := repo.db.Where("email = ?", email).First(&user).Error

	return user, err
}

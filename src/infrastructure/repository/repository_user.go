package repository

import (
	"app/entity"

	"gorm.io/gorm"
)

type RepositoryUser struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *RepositoryUser {
	return &RepositoryUser{db: db}
}

func (u *RepositoryUser) GetByID(id int) (user *entity.EntityUser, err error) {
	u.db.First(&user, id)

	return user, err
}

func (u *RepositoryUser) GetByMail(email string) (user *entity.EntityUser, err error) {
	err = u.db.Where("email = ?", email).First(&user).Error

	return user, err
}

func (u *RepositoryUser) CreateUser(user *entity.EntityUser) error {
	return u.db.Create(&user).Error
}

func (u *RepositoryUser) UpdateUser(user *entity.EntityUser) error {

	_, err := u.GetByMail(user.Email)

	if err != nil {
		return err
	}

	return u.db.Save(&user).Error
}

func (u *RepositoryUser) DeleteUser(user *entity.EntityUser) error {

	_, err := u.GetByMail(user.Email)

	if err != nil {
		return err
	}

	return u.db.Delete(&user).Error
}

func (u *RepositoryUser) GetUsersFromIDs(ids []int) (users []entity.EntityUser, err error) {
	users = make([]entity.EntityUser, 0)

	err = u.db.Where("id IN ?", ids).Find(&users).Error

	return users, err
}

func (u *RepositoryUser) GetUsers(filters entity.EntityUserFilters) (users []entity.EntityUser, err error) {

	users = make([]entity.EntityUser, 0)

	dbFind := u.db

	if filters.Search != "" {
		dbFind = dbFind.Where("name LIKE ? or email LIKE ?", "%"+filters.Search+"%", "%"+filters.Search+"%")
	}

	if filters.Active != "" {
		dbFind = dbFind.Where("active = ?", filters.Active)
	}

	err = dbFind.Find(&users).Error

	return users, err
}

func (u *RepositoryUser) GetUser(id int) (user *entity.EntityUser, err error) {
	u.db.First(&user, id)

	return user, err
}

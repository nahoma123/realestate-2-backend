package persistence

import (
	"context"
	"fmt"
	"strings"
	"time"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"
	"visitor_management/internal/storage"
	"visitor_management/platform/logger"

	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func InitUserDB(db *gorm.DB) storage.UserStorage {
	return &user{
		db: db,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (p *user) Create(ctx context.Context, user *model.User) (*model.User, error) {
	id, _ := uuid.NewV4()
	user.UserID = id.String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Status = constant.Active

	// Check if the user count is zero or below a certain threshold
	var userCount int64
	err := p.db.Model(&model.User{}).Count(&userCount).Error
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		return nil, errors.ErrInternalServerError.New("unknown error occurred")
	}

	if userCount <= 0 {
		// The registering user is the first user, assign admin privileges
		user.Role = constant.AdminRole
	} else {
		// The registering user is a regular user
		user.Role = constant.RegularUserRole
	}

	hash, err := HashPassword(user.Password)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		return nil, errors.ErrInvalidInput.New(errors.UnknownDbError)
	}

	user.Password = hash
	err = p.db.Create(user).Error
	if err != nil {
		logger.Log().Error(ctx, err.Error())

		if err == gorm.ErrInvalidData {
			return nil, errors.ErrDataExists.Wrap(err, errors.UserIsAlreadyRegistered)
		}

		// Check if the error is a duplicate key error
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			// Handle duplicate key error
			return nil, errors.ErrDataExists.Wrap(fmt.Errorf("%s", fmt.Sprintf(errors.TemplateAlreadyRegistered, "user")), fmt.Sprintf(errors.TemplateAlreadyRegistered, "user"))
		} else if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			// Handle duplicate key error
			return nil, errors.ErrDataExists.Wrap(fmt.Errorf("%s", fmt.Sprintf(errors.TemplateAlreadyRegistered, "user")), fmt.Sprintf(errors.TemplateAlreadyRegistered, "user"))
		}

		return nil, errors.ErrInternalServerError.New("unknown error occurred")
	}

	return user, nil
}
func (p *user) Update(ctx context.Context, user *model.User) (*model.User, error) {
	err := p.db.Model(user).Updates(model.User{ /* fields to update */ }).Error
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		return nil, errors.ErrInternalServerError.New("unknown error occurred")
	}

	return user, nil
}

func (p *user) Get(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := p.db.First(&user, id).Error
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNoRecordFound.New("user not found")
		}
		return nil, errors.ErrInternalServerError.New("unknown error occurred")
	}

	return &user, nil
}

func (p *user) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := p.db.Where("email = ?", email).Preload("Rental").Preload("Properties").First(&user).Error
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNoRecordFound.New("user not found")
		}
		return nil, errors.ErrInternalServerError.New("unknown error occurred")
	}

	return &user, nil
}

func (p *user) GetAll(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.User, error) {
	var users []model.User
	err := p.db.Find(&users).Error
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		return nil, errors.ErrInternalServerError.New("unknown error occurred")
	}

	return users, nil
}

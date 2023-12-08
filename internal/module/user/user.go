package user

import (
	"context"
	"fmt"
	"net/smtp"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"
	"visitor_management/internal/storage"
	"visitor_management/internal/storage/persistence"
	"visitor_management/platform/logger"

	"go.uber.org/zap"
)

func (o *user) RegisterUser(ctx context.Context, user *model.User) (*model.User, error) {
	//

	if err := user.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		o.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	user, err := o.userStorage.Create(ctx, user)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return user, nil
}

func (o *user) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	// if err := user.ValidateUpdateUser(); err != nil {
	// 	err = errors.ErrInvalidInput.Wrap(err, "invalid input")
	// 	o.logger.Info(ctx, "invalid input", zap.Error(err))
	// 	return nil, err
	// }

	err := o.generic.UpdateOne(ctx, "users", user, "user_id", user.UserID)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return nil, nil
}

func (o *user) GetUser(ctx context.Context, id string) (*model.User, error) {
	user, err := o.userStorage.Get(ctx, id)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return user, nil
}

func (o *user) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := o.userStorage.GetUserByEmail(ctx, email)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return user, nil
}

func (o *user) VerifyResetCode(ctx context.Context, userCode int, userId, newPassword string) error {
	user := &model.User{}
	err := o.generic.GetOne(ctx, string(storage.Users), user, "user_id", userId)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return err
	}
	hash, err := persistence.HashPassword(newPassword)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		return errors.ErrInvalidInput.New(errors.UnknownDbError)
	}

	user.Password = hash

	err = o.generic.UpdateOne(ctx, string(storage.Users), user, "user_id", user.UserID)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return err
	}

	return nil
}

func (o *user) CreatePasswordResetRequest(ctx context.Context, userId string) error {
	user := &model.User{}
	user.UserID = userId
	user.ResetCode = constant.RandomSixDigitNumber()
	err := o.generic.UpdateOne(ctx, string(storage.Users), user, "user_id", user.UserID)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return err
	}
	err = o.generic.GetOne(ctx, string(storage.Users), user, "user_id", user.UserID)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return err
	}

	err = o.generic.UpdateOne(ctx, string(storage.Users), user, "user_id", user.UserID)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return err
	}

	// error is igonored to prevent revealing internal server error message
	o.SendEmail(user.Email, "Password Reset Code", fmt.Sprintf("You password reset code is %d", user.ResetCode))
	// send email
	return nil
}

func (o *user) SendEmail(to, subject, body string) error {
	email := constant.GetConfig().ORGANIZATION_EMAIL_EMAIL
	password := constant.GetConfig().ORGANIZATION_EMAIL_PASSWORD

	// Set up authentication information.
	auth := smtp.PlainAuth("", email, password, "smtp.gmail.com")

	// Set up email message headers and body.
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Send email using SMTP.
	err := smtp.SendMail("smtp.gmail.com:587", auth, email, []string{to}, msg)
	if err != nil {
		return err
	}
	return nil
}

func (o *UserModuleWrapper) VerifyForgetPasswordCode(ctx context.Context, userCode int, email, newPassword string) error {
	user := &model.User{}
	err := o.generic.GetOne(ctx, string(storage.Users), user, "email", email)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return err
	}
	hash, err := persistence.HashPassword(newPassword)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		return errors.ErrInvalidInput.New(errors.UnknownDbError)
	}

	user.Password = hash

	if user.ResetCode == userCode && userCode != 0 {
		user.ResetCode = 0
		err = o.generic.UpdateOne(ctx, string(storage.Users), user, "user_id", user.UserID)
		if err != nil {
			o.logger.Warn(ctx, err.Error())
			return err
		}
		return nil
	}

	return errors.ErrResetCodeInvalid.New("invalid invite code entered")
}

func (o *user) ForgotPasswordResetRequest(ctx context.Context, email string) error {
	user := &model.User{}
	user.Email = email
	user.ResetCode = constant.RandomSixDigitNumber()
	err := o.generic.UpdateOne(ctx, string(storage.Users), user, "email", email)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return err
	}
	err = o.generic.GetOne(ctx, string(storage.Users), user, "email", email)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return err
	}

	err = o.generic.UpdateOne(ctx, string(storage.Users), user, "email", email)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return err
	}

	// error is igonored to prevent revealing internal server error message
	o.SendEmail(user.Email, "Password Reset Code", fmt.Sprintf("You password reset code is %d", user.ResetCode))
	// send email
	return nil
}

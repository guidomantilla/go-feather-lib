package security

import (
	"context"
	"errors"
	"strings"

	"github.com/xorcare/pointer"
	"gorm.io/gorm"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/config"
	gorm2 "github.com/guidomantilla/go-feather-lib/pkg/datasource/gorm"
)

type GormPrincipalManager struct {
	transactionHandler gorm2.TransactionHandler
	passwordManager    PasswordManager
}

func NewGormPrincipalManager(transactionHandler gorm2.TransactionHandler, passwordManager PasswordManager) *GormPrincipalManager {
	assert.NotNil(transactionHandler, "starting up - error setting up principalManager: transactionHandler is nil")
	assert.NotNil(passwordManager, "starting up - error setting up principalManager: passwordManager is nil")

	return &GormPrincipalManager{
		transactionHandler: transactionHandler,
		passwordManager:    passwordManager,
	}
}

func (manager *GormPrincipalManager) Create(_ context.Context, _ *Principal) error {
	panic("not implemented. no required for this coding challenge")
}

func (manager *GormPrincipalManager) Update(_ context.Context, _ *Principal) error {
	panic("not implemented. no required for this coding challenge")
}

func (manager *GormPrincipalManager) Delete(_ context.Context, _ string) error {
	panic("not implemented. no required for this coding challenge")
}

func (manager *GormPrincipalManager) Find(ctx context.Context, username string) (*Principal, error) {

	var user *Principal
	err := manager.transactionHandler.HandleTransaction(ctx, func(ctx context.Context, tx *gorm.DB) error {

		var principals []AuthPrincipal
		if err := tx.Find(&principals, "username = ? AND application = ?", username, config.Application).Error; err != nil {
			return err
		}
		if len(principals) == 0 {
			return errors.New("principal does not exists")
		}

		principal := principals[0]
		if principal.Role == nil || *(principal.Role) == "" {
			return ErrAccountEmptyRole
		}

		if principal.Password == nil || *(principal.Password) == "" {
			return ErrAccountEmptyPassword
		}

		if principal.Enabled != nil && !*(principal.Enabled) {
			return ErrAccountDisabled
		}

		resources := make([]string, 0)
		for _, principal := range principals {
			resources = append(resources, strings.Join([]string{*principal.Application, *principal.Permission, *principal.Resource}, " "))
		}

		user = &Principal{
			Username:           principal.Username,
			Role:               principal.Role,
			Password:           principal.Password,
			Passphrase:         principal.Passphrase,
			Enabled:            principal.Enabled,
			NonLocked:          pointer.Bool(true),
			NonExpired:         pointer.Bool(true),
			PasswordNonExpired: pointer.Bool(true),
			SignUpDone:         pointer.Bool(true),
			Resources:          resources,
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (manager *GormPrincipalManager) Exists(_ context.Context, _ string) error {
	panic("not implemented. no required for this coding challenge")
}

func (manager *GormPrincipalManager) ChangePassword(_ context.Context, _ string, _ string) error {
	panic("not implemented. no required for this coding challenge")
}

func (manager *GormPrincipalManager) VerifyResource(ctx context.Context, username string, resource string) error {

	return manager.transactionHandler.HandleTransaction(ctx, func(ctx context.Context, tx *gorm.DB) error {

		var principals []AuthPrincipal
		if err := tx.Find(&principals, "username = ? AND CONCAT(application, ' ', permission, ' ', resource) = ?", username, resource).Error; err != nil {
			return err
		}
		if len(principals) == 0 {
			return errors.New("principal resource undefined")
		}

		return nil
	})
}

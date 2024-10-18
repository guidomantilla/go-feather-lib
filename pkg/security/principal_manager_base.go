package security

import (
	"context"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
)

type BasePrincipalManager struct {
	principalRepo   map[string]*Principal
	resourceRepo    map[string]map[string]string
	passwordManager PasswordManager
}

func NewBasePrincipalManager(passwordManager PasswordManager) *BasePrincipalManager {
	assert.NotNil(passwordManager, "starting up - error setting up principalManager: passwordManager is nil")

	return &BasePrincipalManager{
		passwordManager: passwordManager,
		principalRepo:   make(map[string]*Principal),
		resourceRepo:    make(map[string]map[string]string),
	}
}

func (manager *BasePrincipalManager) Create(ctx context.Context, principal *Principal) error {
	assert.NotNil(ctx, "principal manager - error creating: context is nil")
	assert.NotNil(principal, "principal manager - error creating: principal is nil")

	var err error
	if err = manager.Exists(ctx, *principal.Username); err == nil {
		return ErrAccountExistingUsername
	}

	if err = manager.passwordManager.Validate(*principal.Password); err != nil {
		return err
	}

	if principal.Password, err = manager.passwordManager.Encode(*principal.Password); err != nil {
		return err
	}

	manager.principalRepo[*principal.Username] = principal
	manager.resourceRepo[*principal.Username] = make(map[string]string)

	for _, resource := range principal.Resources {
		manager.resourceRepo[*principal.Username][resource] = resource
	}

	return nil
}

func (manager *BasePrincipalManager) Update(ctx context.Context, principal *Principal) error {
	assert.NotNil(ctx, "principal manager - error updating: context is nil")
	assert.NotNil(principal, "principal manager - error updating: principal is nil")

	return manager.Create(ctx, principal)
}

func (manager *BasePrincipalManager) Delete(ctx context.Context, username string) error {
	assert.NotNil(ctx, "principal manager - error deleting: context is nil")
	assert.NotEmpty(username, "principal manager - error deleting: username is empty")

	delete(manager.principalRepo, username)
	delete(manager.resourceRepo, username)
	return nil
}

func (manager *BasePrincipalManager) Find(ctx context.Context, username string) (*Principal, error) {
	assert.NotNil(ctx, "principal manager - error finding: context is nil")
	assert.NotEmpty(username, "principal manager - error finding: username is empty")

	var ok bool
	var user *Principal
	if user, ok = manager.principalRepo[username]; !ok {
		return nil, ErrAccountInvalidUsername
	}

	if user.Role == nil || *(user.Role) == "" {
		return nil, ErrAccountEmptyRole
	}

	if user.Password == nil || *(user.Password) == "" {
		return nil, ErrAccountEmptyPassword
	}

	if user.Enabled != nil && !*(user.Enabled) {
		return nil, ErrAccountDisabled
	}

	if user.NonLocked != nil && !*(user.NonLocked) {
		return nil, ErrAccountLocked
	}

	if user.NonExpired != nil && !*(user.NonExpired) {
		return nil, ErrAccountExpired
	}

	if user.PasswordNonExpired != nil && !*(user.PasswordNonExpired) {
		return nil, ErrAccountExpiredPassword
	}

	return user, nil
}

func (manager *BasePrincipalManager) Exists(ctx context.Context, username string) error {
	assert.NotNil(ctx, "principal manager - error exists: context is nil")
	assert.NotEmpty(username, "principal manager - error exists: username is empty")

	var ok bool
	if _, ok = manager.principalRepo[username]; !ok {
		return ErrAccountInvalidUsername
	}
	return nil
}

func (manager *BasePrincipalManager) ChangePassword(ctx context.Context, username string, password string) error {
	assert.NotNil(ctx, "principal manager - error changing password: context is nil")
	assert.NotEmpty(username, "principal manager - error changing password: username is empty")
	assert.NotEmpty(password, "principal manager - error changing password: password is empty")

	var err error
	if err = manager.Exists(ctx, username); err != nil {
		return err
	}

	if err = manager.passwordManager.Validate(password); err != nil {
		return err
	}

	user := manager.principalRepo[username]
	if user.Password, err = manager.passwordManager.Encode(password); err != nil {
		return err
	}

	return nil
}

func (manager *BasePrincipalManager) VerifyResource(ctx context.Context, username string, resource string) error {
	assert.NotNil(ctx, "principal manager - error verifying resource: context is nil")
	assert.NotEmpty(username, "principal manager - error verifying resource: username is empty")
	assert.NotEmpty(resource, "principal manager - error verifying resource: resource is empty")

	var err error
	if err = manager.Exists(ctx, username); err != nil {
		return err
	}

	if _, ok := manager.resourceRepo[username][resource]; !ok {
		return ErrAccountInvalidAuthorities
	}

	return nil
}

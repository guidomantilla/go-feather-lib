package security

import (
	"context"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type InMemoryPrincipalManager struct {
	principalRepo   map[string]*Principal
	resourceRepo    map[string]map[string]string
	passwordManager PasswordManager
}

func NewInMemoryPrincipalManager(passwordManager PasswordManager) *InMemoryPrincipalManager {

	if passwordManager == nil {
		log.Fatal("starting up - error setting up principalManager: passwordManager is nil")
	}

	return &InMemoryPrincipalManager{
		passwordManager: passwordManager,
		principalRepo:   make(map[string]*Principal),
		resourceRepo:    make(map[string]map[string]string),
	}
}

func (manager *InMemoryPrincipalManager) Create(ctx context.Context, principal *Principal) error {

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

func (manager *InMemoryPrincipalManager) Update(ctx context.Context, principal *Principal) error {
	return manager.Create(ctx, principal)
}

func (manager *InMemoryPrincipalManager) Delete(_ context.Context, username string) error {
	delete(manager.principalRepo, username)
	delete(manager.resourceRepo, username)
	return nil
}

func (manager *InMemoryPrincipalManager) Find(_ context.Context, username string) (*Principal, error) {

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

func (manager *InMemoryPrincipalManager) Exists(_ context.Context, username string) error {

	var ok bool
	if _, ok = manager.principalRepo[username]; !ok {
		return ErrAccountInvalidUsername
	}
	return nil
}

func (manager *InMemoryPrincipalManager) ChangePassword(ctx context.Context, username string, password string) error {

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

func (manager *InMemoryPrincipalManager) VerifyResource(ctx context.Context, username string, resource string) error {

	var err error
	if err = manager.Exists(ctx, username); err != nil {
		return err
	}

	if _, ok := manager.resourceRepo[username][resource]; !ok {
		return ErrAccountInvalidAuthorities
	}

	return nil
}

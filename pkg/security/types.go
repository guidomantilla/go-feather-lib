package security

type AuthResource struct {
	Name        *string `gorm:"primaryKey" json:"name,omitempty"`
	Application *string `gorm:"primaryKey" json:"application,omitempty"`
	Enabled     *bool   `gorm:"enabled" json:"enabled,omitempty"`
}

type AuthRole struct {
	Name    *string `gorm:"primaryKey" json:"name,omitempty"`
	Enabled *bool   `gorm:"enabled" json:"enabled,omitempty"`
}

type AuthAccessControlList struct {
	Role       *string `gorm:"primaryKey" json:"role,omitempty"`
	Resource   *string `gorm:"primaryKey" json:"resource,omitempty"`
	Permission *string `gorm:"primaryKey" json:"permission,omitempty"`
	Enabled    *bool   `gorm:"enabled" json:"enabled,omitempty"`
}

type AuthUser struct {
	Username   *string `gorm:"primaryKey" json:"username,omitempty"`
	Role       *string `gorm:"role" json:"role,omitempty"`
	Password   *string `gorm:"password" json:"password,omitempty"`
	Passphrase *string `gorm:"passphrase" json:"passphrase,omitempty"`
	Enabled    *bool   `gorm:"enabled" json:"enabled,omitempty"`
}

type AuthPrincipal struct {
	Username    *string `gorm:"username" json:"username,omitempty"`
	Role        *string `gorm:"role" json:"role,omitempty"`
	Application *string `gorm:"application" json:"application,omitempty"`
	Resource    *string `gorm:"resource" json:"resource,omitempty"`
	Permission  *string `gorm:"permission" json:"permission,omitempty"`
	Password    *string `gorm:"password" json:"password,omitempty"`
	Passphrase  *string `gorm:"passphrase" json:"passphrase,omitempty"`
	Enabled     *bool   `gorm:"enabled" json:"enabled,omitempty"`
}

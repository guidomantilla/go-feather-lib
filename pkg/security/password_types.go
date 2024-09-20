package security

var (
	_ PasswordEncoder   = (*Argon2PasswordEncoder)(nil)
	_ PasswordEncoder   = (*BcryptPasswordEncoder)(nil)
	_ PasswordEncoder   = (*Pbkdf2PasswordEncoder)(nil)
	_ PasswordEncoder   = (*ScryptPasswordEncoder)(nil)
	_ PasswordEncoder   = (*DelegatingPasswordEncoder)(nil)
	_ PasswordEncoder   = (*passwordManager)(nil)
	_ PasswordGenerator = (*passwordGenerator)(nil)
	_ PasswordGenerator = (*passwordManager)(nil)
	_ PasswordManager   = (*passwordManager)(nil)
)

const (
	Argon2PrefixKey = "{argon2}"
	BcryptPrefixKey = "{bcrypt}"
	Pbkdf2PrefixKey = "{pbkdf2}"
	ScryptPrefixKey = "{scrypt}"
)

type PasswordEncoder interface {
	Encode(rawPassword string) (*string, error)
	Matches(encodedPassword string, rawPassword string) (*bool, error)
	UpgradeEncoding(encodedPassword string) (*bool, error)
}

type PasswordGenerator interface {
	Generate() string
	Validate(rawPassword string) error
}

//

type PasswordManager interface {
	PasswordEncoder
	PasswordGenerator
}

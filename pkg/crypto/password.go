package crypto

import "golang.org/x/crypto/bcrypt"

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(password, hash string) bool
}

type BcryptHasher struct {
	cost int
}

func NewBcryptHasher(cost int) PasswordHasher {
	return &BcryptHasher{cost: cost}
}

func (b *BcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (b *BcryptHasher) Compare(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

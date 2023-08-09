package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (*string, error) {
	passwordByte := []byte(password)
	h, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	hash := string(h)
	return &hash, nil
}

func VerifyPassword(hash string, original string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(original))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

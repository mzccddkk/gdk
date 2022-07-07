package hashx

import "golang.org/x/crypto/bcrypt"

// Bcrypt Creates a password hash with default cost
func Bcrypt(password string) string {
	return BcryptWithCost(password, bcrypt.DefaultCost)
}

// BcryptWithCost Creates a password hash with cost
//  The higher the cost value, the longer the time
func BcryptWithCost(password string, cost int) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// BcryptVerify Verifies that a password matches a hashed password
func BcryptVerify(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

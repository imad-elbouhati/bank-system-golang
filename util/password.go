package util

func HashPassword(password string) (string, error) {
	hashPassword, error = bcrypt.GenerateFromPassword()
}
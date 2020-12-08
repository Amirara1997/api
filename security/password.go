package security


import (
	"golang.org/x/crypto/bcrypt"

)

func Hash(passwd string) ([]byte, error)  {
	return bcrypt.GenerateFromPassword([]byte(passwd),bcrypt.DefaultCost)

}
func VerfyPassword(hashedPasswd,passwd string) error  {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPasswd),[]byte(passwd))
	return err
}
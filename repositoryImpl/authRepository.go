package mySQL

import "errors"

func (r *MySqlRepository) ForgotPassword(mail string) error {
	if !r.CheckIfEmailExists(mail) {
		return nil
	} else {
		return errors.New("email doesn't exist")
	}
}

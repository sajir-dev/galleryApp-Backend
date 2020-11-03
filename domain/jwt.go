package domain

import "../config"

// BlockJWT keeps the record of logged out tokens
func BlockJWT(tokn string) error {
	type TokenType struct {
		Token string `json:"string"`
	}

	t := TokenType{tokn}
	err := config.BlackList.Insert(t)
	return err
}

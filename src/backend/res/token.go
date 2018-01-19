package res

type Token struct {
	TokenString string
}

func GetToken() (*Token, error) {

	token := Token{
		TokenString: "123456",
	}

	return &token, nil
}

func StoreToken() (error) {
	return nil
}

func GetTokenFromStore() (error) {
	return nil
}

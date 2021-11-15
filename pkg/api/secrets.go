package api

func (client Client) GetSecrets() []string {

	if client.fake {
		return makeFakeList("secrets")
	}

	return []string{""}
}

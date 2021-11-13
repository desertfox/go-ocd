package api

func (client Client) GetSecrets() []string {
	return []string{"secret1", "secret2"}
}

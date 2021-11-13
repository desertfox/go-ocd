package api

func (client Client) GetBuildConfigs() []string {
	return []string{"bc1", "bc2"}
}

func (client Client) GetBuildConfigInstance(namespace, instanceName string) string {
	//return fmt.Sprintf("%s/BuildConfig/%s", namespace, instanceName)
	return fakeYaml
}

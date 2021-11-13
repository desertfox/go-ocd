package api

func (client Client) GetAvailKinds() []string {
	return []string{"BuildConfig", "ImageStream", "DeploymentConfig", "Routes", "Secrets", "Services"}
}

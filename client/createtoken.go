package client

// create a Vault token connected to a policy (that was previously created)
func (c *client) CreateToken(policyName string) (string, error) {
	data := map[string]interface{}{
		"policies": policyName,
		"ttl":      "1h",
	}
	secret, err := c.Write("auth/token/create", data, c.token)
	if err != nil {
		return "", err
	}
	newToken := secret.Auth.ClientToken
	return newToken, nil

}

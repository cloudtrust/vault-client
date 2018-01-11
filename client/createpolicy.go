package client

import "errors"

// create a Vault policy for a given action
func (c *client) CreatePolicy(path string, role string, policyName string) error {
	var data map[string]interface{}
	data = make(map[string]interface{})
	switch role {
	case "readkey":
		data["rules"] = "path " + "\"" + path + "\"" + "{ capabilities = [\"read\", \"list\"] }"
	case "writekey", "createkey", "encrypt", "decrypt":
		data["rules"] = "path " + "\"" + path + "\"" + "{ capabilities = [\"create\", \"update\"] }"
	case "exportkey":
		data["rules"] = "path " + "\"" + path + "\"" + "{ capabilities = [\"read\"] }"
	default:
		return errors.New("Wrong request type")
	}
	var policyPath = "sys/policy/" + policyName
	_, err := c.Write(policyPath, data, c.token)
	if err != nil {
		return err
	}
	return nil
}

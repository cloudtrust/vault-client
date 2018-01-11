package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	vault "github.com/hashicorp/vault/api"
	gentlemen "gopkg.in/h2non/gentleman.v2"
)

// read from Vault on the path, given the access token
func (c *client) Read(path string, token string) (*vault.Secret, error) {
	var req *gentlemen.Request
	req = c.httpclient.Get()
	req.Path("/v1/" + path)
	req.SetHeader("X-Vault-Token", token)

	resp, err := req.Do()

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, resp.Error
	}

	// check if we have an error
	if (resp.StatusCode >= 200 && resp.StatusCode < 400) || resp.StatusCode == 429 {
		// we have no error
		var body vault.Secret
		errJSON := resp.JSON(&body)
		if errJSON != nil {
			return nil, errJSON
		}
		return &body, nil
	}

	if resp.RawResponse.Body != nil && resp.StatusCode == 404 {
		return nil, nil
	}

	// we have an error : store it in the buffer and try to decode it
	var bodyBuf bytes.Buffer
	if _, err := io.Copy(&bodyBuf, resp.RawResponse.Body); err != nil {
		return nil, err
	}

	var errorMsgs []string
	errJSON := json.Unmarshal(bodyBuf.Bytes(), &errorMsgs)
	if errJSON != nil {
		return nil, errors.New(bodyBuf.String())
	}

	// we could not decode : write the errors in a raw format
	var errBody bytes.Buffer
	for _, errMsg := range errorMsgs {
		errBody.WriteString(fmt.Sprintf("* %s", errMsg))
	}
	return nil, fmt.Errorf(errBody.String())
}

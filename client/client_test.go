package client

import (
	"errors"
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
	mock "gopkg.in/h2non/gentleman-mock.v2"
	gentlemen "gopkg.in/h2non/gentleman.v2"
)

var respSecret = map[string]interface{}{
	"auth": map[string]interface{}{
		"client_token": "ABCD",
		"policies": []string{
			"web",
			"stage",
		},
		"metadata": map[string]string{
			"user": "armon",
		},
		"lease_duration": 3600,
		"renewable":      true,
	},
}

func createClient(t *testing.T) Client {

	// Configure the mock via gock
	mock.New("http://127.0.0.1:8200").Reply(204).JSON(respSecret)

	var token = "6aeec159-592b-a807-ac54-3060f204116c"
	var path = "http://127.0.0.1:8200"
	var httpClient = gentlemen.New()
	httpClient.Use(mock.Plugin)
	httpClient.BaseURL(path)
	return &client{
		token:      token,
		httpclient: httpClient,
	}
}

func TestClient_CreatePolicy_ReadKey(t *testing.T) {
	client := createClient(t)
	err := client.CreatePolicy("tenants/rolex/fil1/key1", "readkey", "readkey_rolex_fil1")
	assert.Equal(t, err, nil)
}

func TestClient_CreatePolicy_WrongReq(t *testing.T) {
	client := createClient(t)
	err := client.CreatePolicy("tenants/rolex/fil1/key1", "random", "readkey_rolex_fil1")
	assert.Equal(t, err, errors.New("Wrong request type"))
}

func TestClient_CreatePolicy_CreateKey(t *testing.T) {
	client := createClient(t)
	err := client.CreatePolicy("transit/keys/*", "createkey", "createkey")
	assert.Equal(t, err, nil)
}

func TestClient_CreatePolicy_WriteKey(t *testing.T) {
	client := createClient(t)
	err := client.CreatePolicy("tenants/rolex/fil1/key1", "writekey", "writekey_rolex_fil1")
	assert.Equal(t, err, nil)
}

func TestClient_CreatePolicy_ExportKey(t *testing.T) {
	client := createClient(t)
	err := client.CreatePolicy("transit/export/encryption-key/*", "exportkey", "exportkey")
	assert.Equal(t, err, nil)
}

func TestClient_CreateToken(t *testing.T) {
	client := createClient(t)
	token, err := client.CreateToken("createkey")
	assert.Equal(t, err, nil)
	fmt.Println(token)
}

func TestClient_Read(t *testing.T) {
	client := createClient(t)
	secret, err := client.Read("transit/keys/key10", "6aeec159-592b-a807-ac54-3060f204116c")
	assert.Equal(t, err, nil)
	fmt.Println(secret.Data)
}

func TestClient_Write(t *testing.T) {
	client := createClient(t)
	secret, err := client.Write("transit/keys/key10", map[string]interface{}{
		"type":       "aes256-gcm96",
		"derived":    false,
		"exportable": true}, "6aeec159-592b-a807-ac54-3060f204116c")
	assert.Equal(t, err, nil)
	fmt.Println(secret.Data)
}

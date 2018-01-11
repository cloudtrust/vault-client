package client

import (
	"fmt"
	"testing"
)

func createClient(t *testing.T) Client {
	var token = "6aeec159-592b-a807-ac54-3060f204116c"
	var path = "127.0.0.1:8200"
	client, err := NewClient(token, path)
	if err != nil {
		t.Errorf("%s: I have %s and I should get nil", t.Name(), err)
	}
	return client
}

func TestClient_CreatePolicy(t *testing.T) {
	client := createClient(t)
	err := client.CreatePolicy("transit/keys/*", "createkey", "createkey")
	if err != nil {
		t.Errorf("%s: I have %s and I should get nil", t.Name(), err)
	}
	err = client.CreatePolicy("tenants/rolex/fil1/key1", "writekey", "writekey_rolex_fil1")
	if err != nil {
		t.Errorf("%s: I have %s and I should get nil", t.Name(), err)
	}
	err = client.CreatePolicy("tenants/rolex/fil1/key1", "readkey", "readkey_rolex_fil1")
	if err != nil {
		t.Errorf("%s: I have %s and I should get nil", t.Name(), err)
	}
	err = client.CreatePolicy("transit/export/encryption-key/*", "exportkey", "exportkey")
	if err != nil {
		t.Errorf("%s: I have %s and I should get nil", t.Name(), err)
	}

}

func TestClient_CreateToken(t *testing.T) {
	client := createClient(t)
	token, err := client.CreateToken("createkey")
	if err != nil {
		t.Errorf("%s: I have %s and I should get nil", t.Name(), err)
	}
	fmt.Println(token)

}

func TestClient_Read(t *testing.T) {
	client := createClient(t)
	secret, err := client.Read("transit/keys/key10", "6aeec159-592b-a807-ac54-3060f204116c")
	if err != nil {
		t.Errorf("%s: I have %s and I should get nil", t.Name(), err)
	}
	fmt.Println(secret.Data)
}

func TestClient_Write(t *testing.T) {
	client := createClient(t)
	secret, err := client.Write("transit/keys/key10", map[string]interface{}{
		"type":       "aes256-gcm96",
		"derived":    false,
		"exportable": true}, "6aeec159-592b-a807-ac54-3060f204116c")
	if err != nil {
		t.Errorf("%s: I have %s and I should get nil", t.Name(), err)
	}
	fmt.Println(secret.Data)
}

package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/routes"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"

	uuid "github.com/satori/go.uuid"
)

var defaultSSHKeyType = "ssh-rsa"

type PublicKey struct {
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Key   string `json:"key"`
	User  string `json:"user"`
}

func (c *bepaClient) DeleteDefaultUserPublicKey(publicKeyUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		userUUIDPlaceholder:      c.userUUID,
		publicKeyUUIDPlaceholder: publicKeyUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RoutePublicKeyDelete), replaceDict)

	err := c.Do(http.MethodDelete, apiURL, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *bepaClient) GetOneDefaultUserPublicKey(publicKeyUUID *uuid.UUID) (*PublicKey, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder:      c.userUUID,
		publicKeyUUIDPlaceholder: publicKeyUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RoutePublicKeyGetOne), replaceDict)

	publicKey := &PublicKey{}
	err := c.Do(http.MethodGet, apiURL, nil, &publicKey)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}

func (c *bepaClient) GetAllDefaultUserPublicKeys() ([]*PublicKey, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RoutePublicKeyGetAll), replaceDict)

	publicKeys := []*PublicKey{}
	err := c.Do(http.MethodGet, apiURL, nil, &publicKeys)
	if err != nil {
		return nil, err
	}
	return publicKeys, nil
}

func (c *bepaClient) CreatePublicKeyForDefaultUser(title, keyType, key string) (*PublicKey, error) {
	publicKeyReq := &types.PublicKeyReq{
		Title: title,
		Key:   fmt.Sprintf("%s %s", keyType, key),
	}

	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RoutePublicKeyCreate), replaceDict)

	createdPublicKey := &PublicKey{}
	if err := c.Do(http.MethodPost, apiURL, publicKeyReq, createdPublicKey); err != nil {
		return nil, err
	}
	return createdPublicKey, nil
}

func (c *bepaClient) CreatePublicKeyFromFileForDefaultUser(title, fileAdd string) (*PublicKey, error) {
	if fileAdd == "" {
		fileAdd = os.Getenv("HOME") + "/.ssh/id_rsa.pub"
	}
	key, err := ioutil.ReadFile(fileAdd) // #nosec
	if err != nil {
		return nil, err
	}
	return c.CreatePublicKeyForDefaultUser(title, defaultSSHKeyType, string(key))
}

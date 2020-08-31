package vault

import (
	"net/http"
	"time"

	"github.com/k0kubun/pp"

	"github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
)

type VaultKeyStorage struct {
	keysPath string
	client   *api.Client
}

func New(address, token, keysPath string) (*VaultKeyStorage, error) {
	client, err := api.NewClient(&api.Config{
		Address: address,
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "generating vault client")
	}

	client.SetToken(token)
	println(keysPath + "/richie")
	data, err := client.Logical().Read(keysPath + "/richie")
	if err != nil {
		return nil, errors.Wrap(err, "getting secret")
	}

	pp.Println(data)

	return &VaultKeyStorage{}, nil
}

func (s *VaultKeyStorage) UserKey(username string, servicename string) (string, error) {
	return "", nil
}

func (s *VaultKeyStorage) HasService(servicename string) bool {
	return false
}

func (s *VaultKeyStorage) AuthService(service, user, token string) error {
	return nil
}

func (s *VaultKeyStorage) ServiceID(service, user string) (string, error) {
	return "", nil
}

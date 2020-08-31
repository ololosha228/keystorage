package keystorage

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

type Storage interface {
	UserKey(username string, servicename string) (string, error)
	HasService(servicename string) bool
	AuthService(service, user, token string) error
	ServiceID(service, user string) (string, error)
}

type PrimitiveStorage struct {
	keys    map[string]string
	service string
}

func NewPrimitive(service string) *PrimitiveStorage {
	return &PrimitiveStorage{
		keys:    make(map[string]string),
		service: service,
	}
}

func OpenPrimitive(service, path string) (*PrimitiveStorage, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "can't open storage")
	}

	var keys map[string]string
	err = json.Unmarshal(file, &keys)
	if err != nil {
		return nil, errors.Wrap(err, "can't read json keys")
	}

	return &PrimitiveStorage{
		keys:    keys,
		service: service,
	}, nil
}

func (p *PrimitiveStorage) Dump(path string) error {
	data, err := json.Marshal(p.keys)
	if err != nil {
		return errors.Wrap(err, "can't encode json keys")
	}

	err = ioutil.WriteFile(path, data, 0660)
	if err != nil {
		return errors.Wrap(err, "can't write file with keys")
	}
	return nil
}

func (p *PrimitiveStorage) UserKey(username string, servicename string) (string, error) {
	if servicename != p.service {
		return "", errors.New("can't use not " + p.service + " service")
	}
	key, found := p.keys[username]
	if !found {
		return "", &ErrUserNotFound{User: username}
	}
	return key, nil
}

func (p *PrimitiveStorage) HasService(servicename string) bool {
	return servicename == p.service
}

func (p *PrimitiveStorage) AuthService(service, user, token string) error {
	if service != p.service {
		return errors.New("can't use not " + p.service + " service")
	}

	p.keys[user] = token
	return nil
}

// для удобства, если нужно добавить только 1-2 пользователя
func (p *PrimitiveStorage) Set(user, key string) *PrimitiveStorage {
	p.keys[user] = key
	return p
}

func (p *PrimitiveStorage) ServiceID(service, user string) (string, error) {
	if p.service != service {
		return "", errors.New("can't use not " + p.service + " service")
	}

	return user, nil
}

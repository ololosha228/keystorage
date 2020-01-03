package keystorage

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/ungerik/go-dry"
)

type Keystorage struct {
	filepath string
	Data     *Data
}

var (
	instance *Keystorage
)

type (
	username    = string
	servicename = string
)

type Data map[username]map[servicename]string

func CreateStorage(path string) (*Keystorage, error) {
	if instance != nil {
		return instance, nil
	}

	if dry.FileExists(path) {
		return nil, errors.New("file is exist")
	}

	_, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	res := &Keystorage{
		filepath: path,
		Data:     new(Data),
	}

	return res, nil
}

func OpenStorage(path string) (*Keystorage, error) {
	if instance != nil {
		return instance, nil
	}

	if !dry.FileExists(path) {
		return CreateStorage(path)
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	d := new(Data)
	err = yaml.Unmarshal(b, d)
	if err != nil {
		return nil, err
	}

	res := &Keystorage{
		filepath: path,
		Data:     d,
	}

	return res, nil
}

func (ks *Keystorage) UserKey(username, servicename string) (string, error) {
	d := *ks.Data
	if _, ok := d[username]; !ok {
		return "", errors.New("username '" + username + "' not found")
	}

	key, ok := d[username][servicename]
	if !ok {
		return "", errors.New("username '" + username + "' not registred at '" + servicename + "' service")
	}

	return key, nil
}

func (ks *Keystorage) HasService(servicename string) bool {
	return true
}

func (ks *Keystorage) Write(username, servicename, key string) error {
	if (*ks.Data) == nil {
		(*ks.Data) = make(Data)
	}
	if (*ks.Data)[username] == nil {
		(*ks.Data)[username] = make(map[string]string)
	}

	(*ks.Data)[username][servicename] = key

	return ks.Dump()
}

func (ks *Keystorage) Dump() error {
	b, err := yaml.Marshal(ks.Data)
	if err != nil {
		return err
	}

	fi, _ := os.Stat(ks.filepath)

	return ioutil.WriteFile(ks.filepath, b, fi.Mode())
}

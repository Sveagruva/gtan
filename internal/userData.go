package internal

import (
	"encoding/json"
	"errors"
	"os"
	"os/user"
)

type Server struct {
	Hostname string `json:"hostname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type NamedLocation struct {
	ServerName string `json:"serverName"`
	Path       string `json:"path"`
	// todo: only source and destination option
}

type userData struct {
	Servers        map[string]Server        `json:"hosts"`
	NamedLocations map[string]NamedLocation `json:"namedLocations"`
}

var UserData = userData{
	Servers: map[string]Server{},
	NamedLocations: map[string]NamedLocation{
		"default_destination": {
			ServerName: "local",
			Path:       "~/Downloads",
		},
	},
}

func (u *userData) Save() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}

	filePath := usr.HomeDir + "/.gtan.json"
	jso, err := json.MarshalIndent(u, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(filePath, jso, 0644)
}

func (u *userData) Load() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}

	filePath := usr.HomeDir + "/.gtan.json"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err = u.Save(); err != nil {
			return err
		}
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// todo validate file structure for assumption of unique ip and hostname
	return json.Unmarshal(fileBytes, &u)
}

func (u *userData) AddHost(name, hostname, username, password string) error {
	if _, ok := u.Servers[name]; ok {
		return errors.New("host already exists")
	}

	print(u.Servers[name].Username)

	u.Servers[name] = Server{
		Hostname: hostname,
		Username: username,
		Password: password,
	}

	return u.Save()
}

func (u *userData) RemoveHost(name string) error {
	if _, ok := u.Servers[name]; !ok {
		return errors.New("host doesn't exist")
	}

	// todo: check for named locations
	delete(u.Servers, name)

	return u.Save()
}

func (u *userData) AddNamedLocation(name, server, path string) error {
	if _, ok := u.NamedLocations[name]; ok {
		return errors.New("named location already exists")
	}

	if _, ok := u.Servers[server]; !ok && server != "local" {
		return errors.New("server doesn't exist")
	}

	u.NamedLocations[name] = NamedLocation{
		ServerName: server,
		Path:       path,
	}

	return u.Save()
}

func (u *userData) RemoveNamedLocation(name string) error {
	if _, ok := u.NamedLocations[name]; !ok {
		return errors.New("named location doesn't exist")
	}

	delete(u.NamedLocations, name)

	return u.Save()
}

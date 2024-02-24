package internal

import (
	"errors"
	"os"
	"os/user"
	"strings"
)

type Location struct {
	Path    string
	IsLocal bool
	Server  Server
}

func ParseLocation(location string) (Location, error) {
	if !strings.Contains(location, "@") && !strings.Contains(location, ":") {
		if strings.Contains(location, "/") {
			return Location{location, true, Server{}}, nil
		}

		for name, obj := range UserData.NamedLocations {
			if name == location {
				if obj.ServerName == "local" {
					return Location{obj.Path, true, Server{}}, nil
				}

				for serverName, server := range UserData.Servers {
					if serverName == obj.ServerName {
						return Location{obj.Path, obj.ServerName == "local", server}, nil
					}
				}

				return Location{}, errors.New("server for location not found")
			}
		}
		return Location{}, errors.New("named location not found")
	}

	usernameParts := strings.Split(location, "@")
	if len(usernameParts) > 2 {
		return Location{}, errors.New("@ must appear 0 or 1 times")
	}

	pathParts := strings.Split(usernameParts[len(usernameParts)-1], ":")
	if len(pathParts) != 2 {
		return Location{}, errors.New("no path found, must be of the form <host>:<path>")
	}

	path := pathParts[1]

	// no username, try to find defined server from UserData
	if len(usernameParts) == 1 {
		serverName := pathParts[0]

		if serverName == "local" {
			return Location{path, true, Server{}}, nil
		}

		for name, server := range UserData.Servers {
			if name == serverName {
				return Location{path, false, server}, nil
			}

			if server.Hostname == serverName {
				return Location{path, false, server}, nil
			}
		}

		return Location{}, errors.New("server not found")
	}

	username := usernameParts[0]
	hostname := pathParts[0]

	return Location{path, false, Server{
		Hostname: hostname,
		Username: username,
	}}, nil
}

func HydratePath(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		usr, _ := user.Current()
		path = usr.HomeDir + path[1:]
	}

	if strings.HasPrefix(path, "./") {
		if dir, err := os.Getwd(); err == nil {
			path = dir + path[1:]
		} else {
			return "", err
		}
	}

	return path, nil
}

func (l *Location) RsyncString() string {
	if l.IsLocal {
		return l.Path
	}

	return l.Server.Username + "@" + l.Server.Hostname + ":" + l.Path
}

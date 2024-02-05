package api

import (
	"errors"
	"strings"

	json "github.com/buger/jsonparser"
)

func getUUID(name string) (string, error) {
	resp, e, err := makeRequest("https://api.mojang.com/users/profiles/minecraft/" + name)
	if err != nil || resp.StatusCode != 200 {
		return "", err
	}

	if strings.Contains(string(e), "Couldn't find any profile") {
		return "", errors.New("couldn't find any profile")
	}

	uuid, err := json.GetString(e, "id")
	if err != nil {
		return "", err
	}

	return uuid, nil
}

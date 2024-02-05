package api

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"runtime"
	"strings"

	json "github.com/buger/jsonparser"
)

type BedwarsPlayer struct {
	Nicked  bool
	Name    string
	Rank    string
	Level   int
	Finals  int
	FDeaths int
	Wins    int
	Losses  int
	Beds    int
	Bl      int
	FKDR    float64
	BBLR    float64
	WLR     float64
}

var Key string

type Header struct {
	Key   string
	Value string
}

var (
	// response headers for http requests

	RLLimit     string
	RLRemaining string
	RLReset     string
)

func makeHypixelRequest(url string) (*http.Response, []byte, error) {
	resp, body, err := makeRequest(url, Header{"API-Key", Key})
	if err != nil {
		return resp, body, err
	}

	RLLimit = resp.Header.Get("Ratelimit-Limit")
	RLRemaining = resp.Header.Get("Ratelimit-Remaining")
	RLReset = resp.Header.Get("Ratelimit-Reset")

	return resp, body, nil
}

func makeRequest(url string, headers ...Header) (*http.Response, []byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	for _, header := range headers {
		req.Header.Set(header.Key, header.Value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, nil, err
	}
	return resp, body, nil
}

type Method uint8

const (
	// Method to use for getting stats
	Cubelify Method = iota
	Hypixel
)

/*
	func BatchBWStats(names ...string) []BedwarsPlayer {
		var players []BedwarsPlayer

		for _, name := range names {
			go func(name string) {
				players = append(players, GetBWStats(name))
			}(name)
		}

		for len(players) < len(names) {
			runtime.Gosched()
		}
		return players
	}
*/

func getOr0(data []byte, path ...string) int64 {
	val, err := json.GetInt(data, path...)
	if err != nil {
		return 0
	}
	return val
}

func errlog(err error) {
	_, _, line, ok := runtime.Caller(1)
	fmt.Println(err)
	if ok {
		fmt.Printf("Ran into error on line %d\n", line)
	}
}

func GetBWStats(name string, method Method) BedwarsPlayer {
	player := BedwarsPlayer{
		Nicked:  true,
		Name:    name,
		Rank:    "",
		Level:   0,
		Finals:  0,
		FDeaths: 0,
		Wins:    0,
		Losses:  0,
		Beds:    0,
		Bl:      0,
		FKDR:    0,
		BBLR:    0,
		WLR:     0,
	}

	id, err := getUUID(name)
	if err != nil {
		return player
	}

	var body []byte
	var prefix string
	var response *http.Response

	switch method {
	case Cubelify:
		token, err := GetCubelifyToken()
		if err != nil {
			errlog(err)
			return player
		}
		response, body, err = makeRequest("https://api.cubelify.com/proxy/hypixel/player/"+name, Header{"authorization", "Bearer " + token})
		if err != nil || response.StatusCode != 200 {
			errlog(err)
			return player
		}
		prefix = "data"
	case Hypixel:
		response, body, err = makeHypixelRequest("https://api.hypixel.net/player?uuid=" + id)
		if err != nil || response.StatusCode != 200 {
			errlog(err)
			return player
		}

		success, err := json.GetBoolean(body, "success")
		if !success || err != nil || strings.Contains(string(body), `"player":null}`) {
			errlog(err)
			return player
		}
		prefix = "player"
	}

	// if any of these are null, then the player doesn't exist, is nicked, doesn't have achievements (impossible), or hasn't played bedwars before
	data, _, _, err := json.Get(body, prefix)
	if err != nil && err.Error() != "Key path not found" {
		errlog(err)
		return player
	}
	_, _, _, err = json.Get(data, "stats", "bedwars")
	if err != nil && err.Error() != "Key path not found" {
		errlog(err)
		return player
	}
	_, _, _, err = json.Get(data, "achievements")
	if err != nil && err.Error() != "Key path not found" {
		errlog(err)
		return player
	}

	// stat time
	level := getOr0(data, "achievements", "bedwars_level")
	finals := getOr0(data, "stats", "Bedwars", "final_kills_bedwars")
	fdeaths := getOr0(data, "stats", "Bedwars", "final_deaths_bedwars")
	wins := getOr0(data, "stats", "Bedwars", "wins_bedwars")
	losses := getOr0(data, "stats", "Bedwars", "losses_bedwars")
	beds := getOr0(data, "stats", "Bedwars", "beds_broken_bedwars")
	blost := getOr0(data, "stats", "Bedwars", "beds_lost_bedwars")

	fkdr := math.Round(float64(finals)/float64(fdeaths)*100) / 100
	wlr := math.Round(float64(wins)/float64(losses)*100) / 100
	bblr := math.Round(float64(beds)/float64(blost)*100) / 100

	// todo check if infinite - or + and use 0 or numerator respectively
	if math.IsInf(fkdr, 0) || math.IsNaN(fkdr) {
		fkdr = 0
	}
	if math.IsInf(wlr, 0) || math.IsNaN(wlr) {
		wlr = 0
	}
	if math.IsInf(bblr, 0) || math.IsNaN(bblr) {
		bblr = 0
	}

	player.Nicked = false
	player.Level = int(level)
	player.Rank = getRank(data)
	player.Finals = int(finals)
	player.FDeaths = int(fdeaths)
	player.Wins = int(wins)
	player.Losses = int(losses)
	player.Beds = int(beds)
	player.Bl = int(blost)
	player.FKDR = fkdr
	player.WLR = wlr
	player.BBLR = bblr

	return player
}

func CheckKey() bool {
	resp, KeyJSON, err := makeHypixelRequest("https://api.hypixel.net/player?uuid=8e14fc1a-40d4-4984-8665-0f2f1244060f")
	if err != nil || resp.StatusCode != 200 {
		errlog(err)
		return false
	}
	success, _ := json.GetBoolean(KeyJSON, "success")
	return success
}

func getRank(data []byte) string {
	_, _, _, err := json.Get(data)
	if err == nil {
		e, _ := json.GetString(data, "rank")
		if e == "ADMIN" {
			return "ADMIN"
		} else if e == "GAME_MASTER" {
			return "GM"
		} else if e == "YOUTUBER" {
			return "YOUTUBE"
		}
	}
	_, _, _, err = json.Get(data, "monthlyPackageRank")
	if err == nil {
		e, _ := json.GetString(data, "monthlyPackageRank")
		if e == "NONE" {
			return "MVP+"
		} else if e == "SUPERSTAR" {
			return "MVP++"
		}
	}
	_, _, _, err = json.Get(data, "newPackageRank")
	if err == nil {
		e, _ := json.GetString(data, "newPackageRank")
		if e == "null" {
			return "NON"
		} else if e == "VIP" {
			return "VIP"
		} else if e == "VIP_PLUS" {
			return "VIP+"
		} else if e == "MVP" {
			return "MVP"
		} else if e == "MVP_PLUS" {
			return "MVP+"
		}
	}
	return "NON"
}

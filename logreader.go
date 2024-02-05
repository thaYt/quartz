package main

import (
	"context"
	"os"
	"path/filepath"
	"quartz/api"
	"quartz/app"
	"regexp"
	"runtime"
	"strings"
	"time"
)

var replace string

func splitReplace(s string) []string {
	linefeedReplace := strings.ReplaceAll(s, "\r", "")
	colorReplace := regexp.MustCompile(replace+".").ReplaceAllString(linefeedReplace, "")
	return strings.Split(colorReplace, "\n")
}

func handleLine(newcontent string) {
	lines := strings.Split(newcontent, "\n")
	for i := len(lines); i > 0; i-- {
		line := lines[i-1]
		if strings.HasPrefix(line, "ONLINE: ") {
			app.Emit("nuke")
			for _, name := range strings.Split(line[8:], ", ") {
				go app.Emit("addPlayer", api.GetBWStats(name, Config.Method))
			}
		} else if strings.Contains(line, " has joined (") && strings.HasSuffix(line, ")!") || (strings.HasSuffix(line, "reconnected.") && !(strings.Contains(line, ":"))) {
			stats := api.GetBWStats(strings.Split(line, " ")[0], Config.Method)
			app.Emit("addPlayer", stats)
		} else if strings.Contains(line, "Can't find a player by the name of '") && strings.HasSuffix(line, "@'") {
			split := strings.Split(line, "'")
			stats := api.GetBWStats(strings.ReplaceAll(strings.Split(split[len(split)-2], "@'")[0], "@", ""), Config.Method)
			app.Emit("addPlayer", stats)
		} else if strings.Contains(line, "joined the lobby!") || strings.Contains(line, "Sending you to ") {
			app.Emit("nuke")
		} else if strings.HasSuffix(line, " has quit!") || strings.HasSuffix(line, "FINAL KILL!") || strings.HasSuffix(line, "disconnected.") {
			name := strings.Split(line, " ")[0]
			app.Emit("removePlayer", name)
		}
	}
}

func startLogReader(ctx context.Context) {
	Path := getMinecraftDir()
	oldStat, _ := os.Stat(Path)
	oldSize := oldStat.Size()
	b, _ := os.ReadFile(Path)
	oldLen := len(splitReplace(string(b)))
	Reset := false
	switch runtime.GOOS {
	case "windows":
		replace = "�"
	default:
		replace = "§"
	}
	for {
		if Reset {
			Reset = false
			oldStat, _ = os.Stat(Path)
			oldSize = oldStat.Size()
			b, _ = os.ReadFile(Path)
			oldLen = len(splitReplace(string(b)))
		}
		newStat, _ := os.Stat(Path)
		newSize := newStat.Size()
		if oldSize != newSize {
			c, _ := os.ReadFile(Path)
			begDiff := splitReplace(string(c))
			newLen := len(begDiff)
			for i := oldLen; i < newLen; i++ {
				line := begDiff[i-1]
				if len(line) <= 40 {
					continue
				}
				if strings.Contains(line, "cAPI Key is empty!") {
					go func() {
						stats := api.GetBWStats(strings.Split(begDiff[i-2], "[CHAT] ")[1], Config.Method)
						app.Emit("addPlayer", stats)
					}()
				}
				if line[33:40] != "[CHAT] " {
					continue
				}
				go handleLine(line[40:])
			}
			oldSize = newSize
			oldLen = newLen
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func getLunarDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Cannot find user home directory")
	}

	var path string
	switch runtime.GOOS {
	case "windows", "linux":
		path = filepath.Join(homeDir, ".lunarclient/offline/multiver/logs/latest.log")
	case "darwin":
		path = filepath.Join(homeDir, "Library/Application Support/lunarclient/offline/multiver/logs/latest.log")
	default:
		panic("Unsupported OS")
	}

	return path
}

func getMinecraftDir() string {
	homeDir, err := os.UserHomeDir()
	configDir, _ := os.UserConfigDir()
	if err != nil {
		panic("Cannot find user home directory")
	}

	var path string
	switch runtime.GOOS {
	case "windows":
		path = filepath.Join(configDir, ".minecraft\\logs\\latest.log")
	case "linux":
		path = filepath.Join(homeDir, ".minecraft/logs/latest.log")
	case "darwin":
		path = filepath.Join(homeDir, "Library/Application Support/minecraft/logs/latest.log")
	default:
		panic("Unsupported OS")
	}

	return path
}

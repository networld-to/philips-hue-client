package bridge

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
)

// LoadHueBridgeConfig : Loads the Philips Hue configuration file from file.
func LoadHueBridgeConfig() (hueBridge Bridge) {
	usr, err := user.Current()
	var prefix string
	if err == nil {
		prefix = usr.HomeDir
	}
	content, err := os.ReadFile(prefix + "/.philips-hue.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &hueBridge)
	if err != nil {
		fmt.Print("Error:", err)
	}
	return hueBridge
}

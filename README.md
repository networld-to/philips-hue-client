# Philips Hue API v2 Client

The Philips Hue API is provided by a local bridge device that is connected to
the local network via ethernet cable (not PoE powered). This client implementation
in golang implements the second generation of this local Philips Hue API.

* Official Documentation: https://developers.meethue.com/develop/hue-api-v2/
* API Reference: https://developers.meethue.com/develop/hue-api-v2/api-reference/

## Getting Started

Point the domain `philips-hue.home.arpa` to your Philips Hue Bridge IP (not
necessary, but simplifies access). Make sure you have an appropriate Philips Hue
Bridge configuration under `~/.philips-hue.json`. It should look like this:

```json
{"bridge":"philips-hue.home.arpa","username":"THIS_IS_YOUR_PASSWORD","devicetype":"node-philips-hue"}
```

Instructions and helper code how to setup a new username is coming soon.

To test the library run the client via the following command.

```bash
# Comment or uncomment the appropriate lines in client/main.go and run
$ go run client/main.go
```

Use the `philips-hue-client` in your application. Create a `main.go` or any
other go file that uses the library, see example below.

```golang
package main

import (
	"fmt"
	"log"

	"github.com/networld-to/philips-hue-client/bridge"
)

func main() {
	groups, _, err := bridge.GetGroups()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bridge.ChangeGroupState(groups.Data[0].ID, bridge.BRIGHT_WARM_GREEN_ON))
}
```

Run the following commands. This assumes that there is a file that is using
this library, see example snippet above.

```bash
# Initializes your go module. Change to a name of your liking, e.g. github.com/yourusername/yourmodulename
$ go mod init your/go/mod
$ go mod tidy
$ go run main.go
```

## Functionality

* Lights
    * Get all lights
    * Get single light by ID (uuid)
    * Switch light on and off
    * Change light state (during on switch) with presets. Able to control
        * on/off
        * brightness
        * color
        * color temperature
* Groups, same functionality as for lights, but for a predefined set of lights
* Rooms
    * Get all rooms
    * Get single room by ID (uuid)

## Available Predefined Light States

The following presets can be used to switch on the light and configure it
appropriately, see [presets.go](./bridge/presets.go)

* BRIGHT_NEUTRAL_WHITE_ON
* BRIGHT_WARM_RED_ON
* BRIGHT_WARM_GREEN_ON
* BRIGHT_WARM_BLUE_ON

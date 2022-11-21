package bridge

// Bridge defines the config values to be able to connect to the Philips Hue API
type Bridge struct {
	Bridge     string `json:"bridge"`
	Username   string `json:"username"`
	Devicetype string `json:"devicetype"`
}

type Lights struct {
	Errors []interface{} `json:"errors"`
	Data   []struct {
		ID    string `json:"id"`
		IDV1  string `json:"id_v1"`
		Owner struct {
			Rid   string `json:"rid"`
			Rtype string `json:"rtype"`
		} `json:"owner"`
		Metadata struct {
			Name      string `json:"name"`
			Archetype string `json:"archetype"`
		} `json:"metadata"`
		On On
		Dimming struct {
			Brightness  float64 `json:"brightness"`
			MinDimLevel float64 `json:"min_dim_level"`
		} `json:"dimming"`
		DimmingDelta struct {
		} `json:"dimming_delta"`
		ColorTemperature struct {
			Mirek       int  `json:"mirek"`
			MirekValid  bool `json:"mirek_valid"`
			MirekSchema struct {
				MirekMinimum int `json:"mirek_minimum"`
				MirekMaximum int `json:"mirek_maximum"`
			} `json:"mirek_schema"`
		} `json:"color_temperature"`
		ColorTemperatureDelta struct {
		} `json:"color_temperature_delta"`
		Color struct {
			Xy Xy
			Gamut struct {
				Red struct {
					X float64 `json:"x"`
					Y float64 `json:"y"`
				} `json:"red"`
				Green struct {
					X float64 `json:"x"`
					Y float64 `json:"y"`
				} `json:"green"`
				Blue struct {
					X float64 `json:"x"`
					Y float64 `json:"y"`
				} `json:"blue"`
			} `json:"gamut"`
			GamutType string `json:"gamut_type"`
		} `json:"color"`
		Dynamics struct {
			Status       string   `json:"status"`
			StatusValues []string `json:"status_values"`
			Speed        float64  `json:"speed"`
			SpeedValid   bool     `json:"speed_valid"`
		} `json:"dynamics"`
		Alert struct {
			ActionValues []string `json:"action_values"`
		} `json:"alert"`
		Signaling struct {
		} `json:"signaling"`
		Mode string `json:"mode"`
		Type string `json:"type"`
	} `json:"data"`
}

type OnOff struct {
	// On/Off state of the light on=true, off=false
	On On `json:"on,omitempty"`
}

type LightControl struct {
	// On/Off state of the light on=true, off=false
	On On `json:"on,omitempty"`
	Dimming Dimming `json:"dimming,omitempty"`
	Color Color `json:"color,omitempty"`
	ColorTemperature ColorTemperature `json:"color_temperature,omitempty"`
}

type On struct {
	On bool `json:"on"`
}

type Dimming struct {
	// Brightness percentage. value cannot be 0, writing 0 changes it to lowest possible brightness
	Brightness  float64 `json:"brightness,omitempty"`
}

type Color struct {
	Xy Xy `json:"xy,omitempty"`
}

type Xy struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ColorTemperature struct {
	// Mirek: (integer – minimum: 153 – maximum: 500)
	// color temperature in mirek or null when the light color is not in the ct spectrum
	Mirek int `json:"mirek,omitempty"`
}

type Rooms struct {
	Errors []interface{} `json:"errors"`
	Data   []struct {
		ID       string `json:"id"`
		IDV1     string `json:"id_v1"`
		Children []struct {
			Rid   string `json:"rid"`
			Rtype string `json:"rtype"`
		} `json:"children"`
		Services []struct {
			Rid   string `json:"rid"`
			Rtype string `json:"rtype"`
		} `json:"services"`
		Metadata struct {
			Name      string `json:"name"`
			Archetype string `json:"archetype"`
		} `json:"metadata"`
		Type string `json:"type"`
	} `json:"data"`
}

type Groups struct {
	Errors []interface{} `json:"errors"`
	Data   []struct {
		ID    string `json:"id"`
		IDV1  string `json:"id_v1"`
		Owner struct {
			Rid   string `json:"rid"`
			Rtype string `json:"rtype"`
		} `json:"owner"`
		On On
		Dimming Dimming
		DimmingDelta struct {
		} `json:"dimming_delta"`
		ColorTemperature struct {
		} `json:"color_temperature"`
		ColorTemperatureDelta struct {
		} `json:"color_temperature_delta"`
		Color Color
		Alert struct {
			ActionValues []string `json:"action_values"`
		} `json:"alert"`
		Signaling struct {
		} `json:"signaling"`
		Dynamics struct {
		} `json:"dynamics"`
		Type string `json:"type"`
	} `json:"data"`
}
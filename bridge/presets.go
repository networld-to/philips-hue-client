package bridge

// Xy coordinates generated via https://www.luxalight.eu/en/cie-convertor and
// https://developers.meethue.com/develop/hue-api-v2/core-concepts/

var BRIGHT_NEUTRAL_WHITE_ON = &LightControl{
	On: On{
		On: true,
	},
	// ColorTemperature: ColorTemperature{
	// 	Mirek: 153,
	// },
	Dimming: Dimming{
		Brightness: 100.0,
	},
	Color: Color{
		Xy: Xy{
			X: 0.3333333333333333,
			Y: 0.3333333333333333,
		},
	},
}

var BRIGHT_WARM_RED_ON = &LightControl{
	On: On{
		On: true,
	},
	// ColorTemperature: ColorTemperature{
	// 	Mirek: 500,
	// },
	Dimming: Dimming{
		Brightness: 100.0,
	},
	Color: Color{
		Xy: Xy{
			X: 1,
			Y: 0,
		},
	},
}

var BRIGHT_WARM_GREEN_ON = &LightControl{
	On: On{
		On: true,
	},
	// ColorTemperature: ColorTemperature{
	// 	Mirek: 500,
	// },
	Dimming: Dimming{
		Brightness: 100.0,
	},
	Color: Color{
		Xy: Xy{
			X: 0,
			Y: 1,
		},
	},
}

var BRIGHT_WARM_BLUE_ON = &LightControl{
	On: On{
		On: true,
	},
	// ColorTemperature: ColorTemperature{
	// 	Mirek: 500,
	// },
	Dimming: Dimming{
		Brightness: 100.0,
	},
	Color: Color{
		Xy: Xy{
			X: 0,
			Y: 0,
		},
	},
}
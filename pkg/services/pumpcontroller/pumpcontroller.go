package pumpcontroller

type Message struct {
	PumpGpio  uint8  `json:"pumpGpio"`
	DurationS uint16 `json:"durationS"`
}

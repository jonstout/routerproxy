package device


type Command struct {
	Text string `json:"text"`
	Type string `json:"type"`
	Addr string `json:"address"`
}

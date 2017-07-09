package device


type Command struct {
	RouteTable string `json:"route_table"`
	Selection  string `json:"selection"`
	Text       string `json:"text"`
	Addr string
}

type CommandRequest struct {
	Command Command `json:"command"`
	Devices []int   `json:"devices"`
}

type CommandResponse struct {
	Command Command `json:"command"`
	DeviceID int    `json:"device_id"`
	Output string   `json:"output"`
}

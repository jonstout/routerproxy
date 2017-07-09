package database

import "time"

type Device struct {
	Description string `json:"description"`
	Hostname    string `json:"hostname"`
	ID          int    `json:"id"`
	IPAddress   string `json:"ip_address"`
	Location    string `json:"location"`
	POP         string `json:"pop"`
}


func Devices() []Device {
	t := time.Now().Unix()
	ip := ""
	if t % 2 == 0 {
		ip = "192.168.1.1"
	} else {
		ip = "127.0.0.1"
	}

	result := []Device{
		Device{
			Description: "a little description",
			Hostname: "test.domain.com",
			ID: 1,
			IPAddress: ip,
			Location: "city, ST",
			POP: "POP",
		},
		Device{
			Description: "a little description",
			Hostname: "test.domain.com",
			ID: 2,
			IPAddress: "127.0.0.1",
			Location: "city, ST",
			POP: "POP",
		},
		Device{
			Description: "a little description",
			Hostname: "test.domain.com",
			ID: 3,
			IPAddress: "127.0.0.1",
			Location: "city, ST",
			POP: "POP",
		},
		Device{
			Description: "a little description",
			Hostname: "test.domain.com",
			ID: 4,
			IPAddress: "127.0.0.1",
			Location: "city, ST",
			POP: "POP",
		},
	}

	return result
}

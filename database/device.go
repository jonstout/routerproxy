package database

type Device struct {
	Description string `json:"description"`
	Hostname    string `json:"hostname"`
	ID          int    `json:"id"`
	IPAddress   string `json:"ip_address"`
	Location    string `json:"location"`
	POP         string `json:"pop"`
	Type        string `json:"type"`
}


func GetDevice(ID int) Device {
	result := Device{
		Description: "a little description",
		Hostname: "test.domain.com",
		ID: ID,
		IPAddress: "",
		Location: "city, ST",
		POP: "POP",
		Type: "",
	}

	if ID == 1 {
		result.IPAddress = "156.56.6.115"
		result.Type = "juniper"
	} else if ID == 2 {
		result.IPAddress = "156.56.6.116"
		result.Type = "juniper"
	} else if ID == 3 {
		result.IPAddress = "156.56.6.220"
		result.Type = "cisco"
	} else {
		result.IPAddress = "156.56.6.221"
		result.Type = "cisco"
	}

	return result
}

func GetDevices() []Device {
	result := []Device{
		Device{
			Description: "a little description",
			Hostname: "test.domain.com",
			ID: 1,
			IPAddress: "156.56.6.115",
			Location: "city, ST",
			POP: "POP",
			Type: "juniper",
		},
		Device{
			Description: "a little description",
			Hostname: "test.domain.com",
			ID: 2,
			IPAddress: "156.56.6.116",
			Location: "city, ST",
			POP: "POP",
			Type: "juniper",
		},
		Device{
			Description: "a little description",
			Hostname: "test.domain.com",
			ID: 3,
			IPAddress: "156.56.6.220",
			Location: "city, ST",
			POP: "POP",
			Type: "cisco",
		},
		Device{
			Description: "a little description",
			Hostname: "test.domain.com",
			ID: 4,
			IPAddress: "156.56.6.221",
			Location: "city, ST",
			POP: "POP",
			Type: "cisco",
		},
	}

	return result
}

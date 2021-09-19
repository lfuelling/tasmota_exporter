package main

type Status struct {
	Module       int
	FriendlyName []string
	Topic        string
	ButtonTopic  string
	Power        int
	PowerOnState int
	LedState     int
	LedMask      string
	SaveData     int
	SaveState    int
	SwitchTopic  string
	SwitchMode   []int
	ButtonRetain int
	SwitchRetain int
	SensorRetain int
	PowerRetain  int
}

type DeviceStatus struct {
	Status Status
}

type EnergyStatus struct {
	TotalStartTime string
	Total          float64
	Yesterday      float64
	Today          float64
	Power          int
	ApparentPower  int
	ReactivePower  int
	Factor         float64
	Voltage        int
	Current        float64
}

type StatusSNS struct {
	Time   string
	ENERGY EnergyStatus
}

type DevicePowerStatus struct {
	StatusSNS StatusSNS
}

type DeviceMetrics struct {
	address         string
	label           string
	powerMonitoring bool
	power           DevicePowerStatus
	status          DeviceStatus
}

type ConfigDevice struct {
	Address         string
	Username        string
	Password        string
	Label           string
	PowerMonitoring bool
}

type Config struct {
	Devices       []ConfigDevice
	ListenAddress string
}

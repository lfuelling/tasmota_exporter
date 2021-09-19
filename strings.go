package main

import (
	"strconv"
)

func getHelpString() string {
	return "" +
		// Power metrics
		"# HELP device_power_total Total energy.\n" +
		"# TYPE device_power_total gauge\n" +
		"# HELP device_power_today Energy today.\n" +
		"# TYPE device_power_today gauge\n" +
		"# HELP device_power_yesterday Energy yesterday.\n" +
		"# TYPE device_power_yesterday gauge\n" +
		"# HELP device_power_power Power in Watt.\n" +
		"# TYPE device_power_power gauge\n" +
		"# HELP device_power_apparent Apparent power.\n" +
		"# TYPE device_power_apparent gauge\n" +
		"# HELP device_power_reactive Reactive power.\n" +
		"# TYPE device_power_reactive gauge\n" +
		"# HELP device_power_factor Power factor.\n" +
		"# TYPE device_power_factor gauge\n" +
		"# HELP device_power_voltage Voltage.\n" +
		"# TYPE device_power_voltage gauge\n" +
		"# HELP device_power_current Current in A.\n" +
		"# TYPE device_power_current gauge\n" +

		// Device Metrics
		"# HELP device_status_module Module.\n" +
		"# TYPE device_status_module gauge\n" +
		"# HELP device_status_power Current power state, 1 means on.\n" +
		"# TYPE device_status_power gauge\n" +
		"# HELP device_status_power_on_state Device powerOnState.\n" +
		"# TYPE device_status_power_on_state gauge\n" +
		"# HELP device_status_led_state Device LED state.\n" +
		"# TYPE device_status_led_state gauge\n"
}

func getFetchSuccessString(address string, label string, fetchType string, state int) string {
	return "fetch_success{address=\"" + address + "\",label=\"" + label + "\",type=\"" + fetchType + "\"} " + strconv.Itoa(state) + "\n"
}

func getFetchDurationString(address string, label string, duration int64) string {
	return "fetch_duration{address=\"" + address + "\",label=\"" + label + "\"} " + strconv.FormatInt(duration, 10) + "\n"
}

func getDeviceStatusString(d DeviceMetrics) string {
	return `device_status_module{address="` + d.address + `",label="` + d.label + `"} ` + strconv.Itoa(d.status.Status.Module) + "\n" +
		`device_status_power{address="` + d.address + `",label="` + d.label + `"} ` + strconv.Itoa(d.status.Status.Power) + "\n" +
		`device_status_power_on_state{address="` + d.address + `",label="` + d.label + `"} ` + strconv.Itoa(d.status.Status.PowerOnState) + "\n" +
		`device_status_led_state{address="` + d.address + `",label="` + d.label + `"} ` + strconv.Itoa(d.status.Status.LedState) + "\n"
}

func getDevicePowerString(d DeviceMetrics) string {
	return `device_power_total{address="` + d.address + `",label="` + d.label + `"} ` + strconv.FormatFloat(d.power.StatusSNS.ENERGY.Total, 'E', -1, 32) + "\n" +
		`device_power_today{address="` + d.address + `",label="` + d.label + `"} ` + strconv.FormatFloat(d.power.StatusSNS.ENERGY.Today, 'E', -1, 32) + "\n" +
		`device_power_yesterday{address="` + d.address + `",label="` + d.label + `"} ` + strconv.FormatFloat(d.power.StatusSNS.ENERGY.Yesterday, 'E', -1, 32) + "\n" +
		`device_power_power{address="` + d.address + `",label="` + d.label + `"} ` + strconv.Itoa(d.power.StatusSNS.ENERGY.Power) + "\n" +
		`device_power_apparent{address="` + d.address + `",label="` + d.label + `"} ` + strconv.Itoa(d.power.StatusSNS.ENERGY.ApparentPower) + "\n" +
		`device_power_reactive{address="` + d.address + `",label="` + d.label + `"} ` + strconv.Itoa(d.power.StatusSNS.ENERGY.ReactivePower) + "\n" +
		`device_power_factor{address="` + d.address + `",label="` + d.label + `"} ` + strconv.FormatFloat(d.power.StatusSNS.ENERGY.Factor, 'E', -1, 32) + "\n" +
		`device_power_voltage{address="` + d.address + `",label="` + d.label + `"} ` + strconv.Itoa(d.power.StatusSNS.ENERGY.Voltage) + "\n" +
		`device_power_current{address="` + d.address + `",label="` + d.label + `"} ` + strconv.FormatFloat(d.power.StatusSNS.ENERGY.Current, 'E', -1, 32) + "\n"
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func fetchDeviceStatus(d ConfigDevice) (DeviceStatus, error) {
	var status DeviceStatus

	log.Println("Fetching metrics for '" + d.Address)
	resp, err := http.Get("http://" + d.Address + "/cm?cmnd=Status&user=" + d.Username + "&password=" + d.Password)

	if err != nil {
		return DeviceStatus{}, err
	} else {
		if resp.StatusCode == http.StatusOK {

			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			jsonErr := json.Unmarshal(bodyBytes, &status)
			if jsonErr != nil {
				return DeviceStatus{}, jsonErr
			}
		}
	}

	err1 := resp.Body.Close()
	if err1 != nil {
		return DeviceStatus{}, err1
	}

	return status, nil
}

func fetchDevicePower(d ConfigDevice) (DevicePowerStatus, error) {
	var powerStatus DevicePowerStatus

	log.Println("Fetching power metrics for '" + d.Address)
	resp, err := http.Get("http://" + d.Address + "/cm?cmnd=Status%208&user=" + d.Username + "&password=" + d.Password)

	if err != nil {
		return DevicePowerStatus{}, err
	} else {
		if resp.StatusCode == http.StatusOK {

			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			jsonErr := json.Unmarshal(bodyBytes, &powerStatus)
			if jsonErr != nil {
				return DevicePowerStatus{}, jsonErr
			}
		}
	}

	err1 := resp.Body.Close()
	if err1 != nil {
		return DevicePowerStatus{}, err1
	}

	return powerStatus, nil
}
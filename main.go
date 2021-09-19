package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var config Config

func loadDeviceMetrics() ([]DeviceMetrics, string) {
	var fetchMetrics = ""
	var deviceMetrics []DeviceMetrics

	for _, d := range config.Devices {
		startTime := time.Now().UnixNano()

		var metrics DeviceMetrics
		metrics.address = d.Address
		metrics.label = d.Label
		metrics.powerMonitoring = d.PowerMonitoring

		deviceStatus, statusErr := fetchDeviceStatus(d)

		if statusErr != nil {
			log.Println("Error fetching metrics for '"+d.Address+"'!", statusErr)
			fetchMetrics += getFetchSuccessString(d.Address, d.Label, "status", 0)
		} else {
			metrics.status = deviceStatus
			fetchMetrics += getFetchSuccessString(d.Address, d.Label, "status", 1)
		}

		if d.PowerMonitoring {
			devicePowerStatus, powerErr := fetchDevicePower(d)

			if powerErr != nil {
				log.Println("Error fetching power metrics for '"+d.Address+"'!", powerErr)
				fetchMetrics += getFetchSuccessString(d.Address, d.Label, "power", 0)
			} else {
				metrics.power = devicePowerStatus
				fetchMetrics += getFetchSuccessString(d.Address, d.Label, "power", 1)
			}
		}

		deviceMetrics = append(deviceMetrics, metrics)
		fetchMetrics += getFetchDurationString(d.Address, d.Label, time.Now().UnixNano()-startTime)
	}
	return deviceMetrics, fetchMetrics
}

func renderMetricsResponse() (string, error) {
	devices, metrics := loadDeviceMetrics()

	var res string

	res = getHelpString() + metrics

	for _, device := range devices {
		res += getDeviceStatusString(device)

		if device.powerMonitoring {
			res += getDevicePowerString(device)
		}
	}

	return res, nil
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/metrics" {
		response, err := renderMetricsResponse()
		if err != nil {
			log.Println("Error fetching metrics!", err)
			w.WriteHeader(500)
			_, _ = fmt.Fprint(w, err)
			return
		}
		_, _ = fmt.Fprint(w, response)
	} else {
		log.Println("Not found: '" + r.RequestURI)
		w.WriteHeader(404)
	}
}

func main() {
	var configPath = flag.String("config", "./config.json", "path to the config file")
	flag.Parse()

	log.Println("Loading config from '" + *configPath)
	file, err := os.Open(*configPath)
	defer file.Close()
	if err != nil {
		log.Fatalln("Unable to open config file!", err)
		return
	}
	decoder := json.NewDecoder(file)
	config = Config{}
	err1 := decoder.Decode(&config)
	if err1 != nil {
		log.Fatalln("Unable to read config!", err1)
		return
	}

	server := &http.Server{
		Addr:         config.ListenAddress,
		Handler:      http.HandlerFunc(handleMetrics),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("Starting server on '" + config.ListenAddress)
	err2 := server.ListenAndServe()
	if err2 != nil {
		log.Fatalln("Unable to start server!", err2)
	}
}

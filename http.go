package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func deviceHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Devices)
}

func devicesHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Devices)
}

func switchHandler(w http.ResponseWriter, r *http.Request) {
	var state = r.FormValue("state")
	var mac = r.FormValue(("mac"))
	var bstate = false

	if state == "on" {
		bstate = true
	}

	_, err := SetState(mac, bstate)
	if err != nil {
		fmt.Println(err.Error())
	}
}

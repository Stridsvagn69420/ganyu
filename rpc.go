package main

import (
	"time"

	"github.com/hugolgst/rich-go/client"
)

const RPCID = "963883461905629235"

var Active bool = false

func StartRPC() {
	err := client.Login(RPCID)
	if err == nil {
		Active = true
	}
}

func UpdateRPC(state string, details string) {
	if Active {
		starttime := time.Now()
		timestamp := client.Timestamps{
			Start: &starttime,
		}
		client.SetActivity(client.Activity{
			State:      state,
			Details:    details,
			Timestamps: &timestamp,
		})
	}
}

func StopRPC() {
	if Active {
		client.Logout()
	}
}

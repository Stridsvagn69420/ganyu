package main

import (
	"time"

	"github.com/hugolgst/rich-go/client"
)

const RPCID = "963883461905629235"

var Active bool = false

func StartRPC() {
	err := client.Login(RPCID)
	if err != nil {
		Active = false
	} else {
		Active = true
	}
}

func UpdateRPC(state string, details string) {
	if Active {
		starttime := time.Now()
		client.SetActivity(client.Activity{
			State:   state,
			Details: details,
			Timestamps: &client.Timestamps{
				Start: &starttime,
			},
			Buttons: []*client.Button{
				{
					Label: "GitHub",
					Url:   REPOSITORY,
				},
			},
			LargeImage: "ganyu",
		})
	}
}

func UpdateRPCTime(state string, details string, starttime time.Time) {
	if Active {
		client.SetActivity(client.Activity{
			State:   state,
			Details: details,
			Timestamps: &client.Timestamps{
				Start: &starttime,
			},
			Buttons: []*client.Button{
				{
					Label: "GitHub",
					Url:   REPOSITORY,
				},
			},
			LargeImage: "ganyu",
		})
	}
}

func StopRPC() {
	if Active {
		client.Logout()
	}
}

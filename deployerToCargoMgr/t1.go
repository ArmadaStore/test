package main

import (
	"fmt"
	"os"

	"github.com/ArmadaStore/comms/rpc/taskToCargoMgr"
)

func main() {

	cargoMgrIP := os.Args[1]
	cargoMgrPort := os.Args[2]

	fmt.Println("Cargo Mgr IP: ", cargoMgrIP, "-- Cargo Mgr Port: ", cargoMgrPort)

	requestInfo = taskToCargoMgr.RequesterInfo{
		lat:       44.985024576832416,
		lon:       -93.2273939423279,
		size:      500,
		nReplicas: 2,
		appID:     1234,
	}

}

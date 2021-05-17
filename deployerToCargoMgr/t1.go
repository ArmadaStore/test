package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ArmadaStore/comms/rpc/taskToCargoMgr"
	"google.golang.org/grpc"
)

func main() {

	cargoMgrIP := os.Args[1]
	cargoMgrPort := os.Args[2]

	fmt.Println("Cargo Mgr IP: ", cargoMgrIP, "-- Cargo Mgr Port: ", cargoMgrPort)

	requestInfo := taskToCargoMgr.RequesterInfo{
		Lat:       37.213189,
		Lon:       -82.29759,
		Size:      500,
		NReplicas: 3,
		AppID:     "1234",
	}

	conn, err := grpc.Dial(cargoMgrIP+":"+cargoMgrPort, grpc.WithInsecure())
	if err != nil {
		panic("ERROR: gRPC connection error...")
	}

	fmt.Println("Connection Success with", cargoMgrIP, "and", cargoMgrPort, "\n")

	service := taskToCargoMgr.NewRpcTaskToCargoMgrClient(conn)
	cargoList, err := service.RequestCargo(context.Background(), &requestInfo)
	if err != nil {
		panic("ERROR: Cargo request failed...")
	}

	cargoIPs := cargoList.GetIPs()
	cargoPorts := cargoList.GetPorts()

	if len(cargoIPs) != len(cargoPorts) {
		panic("ERROR: Returned list of Cargo IPs and Ports differ in length")
	}

	for i := 0; i < len(cargoIPs); i++ {
		fmt.Println("Cargo ", i+1, " - ", cargoIPs[i], ":", cargoPorts[i])
	}

	conn.Close()
}

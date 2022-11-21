package main

import (
	"fmt"
	"log"

	"github.com/networld-to/philips-hue-client/bridge"
)


func TestGetLights() {
	fmt.Println("---- BEGIN: GetLights()")
	lights, resp, err := bridge.GetLights()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)
	fmt.Printf("%+v\n", lights.Data)
	for idx, light := range lights.Data {
		fmt.Printf("[%d] %s :: %+v\n", idx, light.ID, light.Metadata)
	}
	fmt.Println("----")

	if len(lights.Data) >= 1 {
		fmt.Println("---- BEGIN: GetLight(:id)")
		lights, _, _ = bridge.GetLight(lights.Data[0].ID)
		fmt.Printf("%+v\n", lights.Data[0])
		fmt.Println("----")
	}
}

func TestGetRooms() {
	fmt.Println("---- BEGIN: GetRooms()")
	rooms, resp, _ := bridge.GetRooms()
	fmt.Println(resp.Status)
	fmt.Printf("%+v\n", rooms)
	for idx, room := range rooms.Data {
		fmt.Printf("[%d] %s :: %+v\n", idx, room.ID, room.Metadata)
	}
	fmt.Println("----")

	fmt.Println("---- BEGIN: GetRoom(:id)")
	rooms, resp, _ = bridge.GetRoom(rooms.Data[0].ID)
	fmt.Println(resp.Status)
	fmt.Printf("%+v\n", rooms)
	for idx, room := range rooms.Data {
		fmt.Printf("[%d] %s :: %+v\n", idx, room.ID, room.Metadata)
	}
	fmt.Println("----")
}

func main() {
	TestGetLights()
	TestGetRooms()
}

package main

import (
	"fmt"
	"slices"
)

type Room struct {
	RoomNumber int
	// Single - Double - Standard
	Type     string
	BedCount int
	IsActive bool
	Price    float32
}

var rooms []Room = *GenerateRooms()

func main() {
	GetUserCommand()
}

func GetUserCommand() {

	var input int

	for i := true; i != false; {

		// fmt.Println(rooms)

		fmt.Println(`
	   1. Add Room
	   2. Reserve Room
	   3. Delete Room 
	  `)

		fmt.Scanln(&input)

		switch input {
		case 1:
			AddRoom()
		case 2:
			var roomNumber int
			fmt.Print("Please Enter Room Number to reserve : ")
			fmt.Scanln(&roomNumber)
			ReserveRoom(roomNumber)
		case 3:
			var roomNumber int
			fmt.Print("Please Enter Room Number to delete : ")
			fmt.Scanln(&roomNumber)
			DeleteRoom(roomNumber)
		default:
			fmt.Print("incorect input !!!")
		}

	}
}

func GenerateRooms() *[]Room {
	var rooms []Room = []Room{}
	rooms = append(rooms, Room{RoomNumber: 1, Type: "Single", BedCount: 2, IsActive: true, Price: 1.260000})
	rooms = append(rooms, Room{RoomNumber: 2, Type: "Double", BedCount: 3, IsActive: true, Price: 1.440000})
	rooms = append(rooms, Room{RoomNumber: 3, Type: "Standard", BedCount: 2, IsActive: true, Price: 1.200000})
	rooms = append(rooms, Room{RoomNumber: 4, Type: "Single", BedCount: 5, IsActive: true, Price: 3.200000})
	rooms = append(rooms, Room{RoomNumber: 5, Type: "Standard", BedCount: 1, IsActive: true, Price: 4.20000})
	rooms = append(rooms, Room{RoomNumber: 6, Type: "Single", BedCount: 4, IsActive: true, Price: 1.104000})
	rooms = append(rooms, Room{RoomNumber: 7, Type: "Double", BedCount: 2, IsActive: true, Price: 1.200000})
	rooms = append(rooms, Room{RoomNumber: 8, Type: "Standard", BedCount: 5, IsActive: true, Price: 121.000})
	rooms = append(rooms, Room{RoomNumber: 9, Type: "Single", BedCount: 2, IsActive: true, Price: 400.000})
	rooms = append(rooms, Room{RoomNumber: 10, Type: "Standard", BedCount: 3, IsActive: true, Price: 220.0000})
	rooms = append(rooms, Room{RoomNumber: 11, Type: "Double", BedCount: 4, IsActive: true, Price: 320.0000})
	rooms = append(rooms, Room{RoomNumber: 12, Type: "Standard", BedCount: 1, IsActive: true, Price: 420.0000})
	return &rooms
}

func ReserveRoom(roomNumber int) {
	room := GetRoom(roomNumber)

	if room == nil {
		fmt.Println("Room not found !!!")
		return
	}

	if room.IsActive {
		fmt.Println("Room is full !!!")
		return
	}

	room.IsActive = true
}

func GetRoom(roomNumber int) *Room {
	for _, room := range rooms {
		if room.RoomNumber == roomNumber {
			return &room
		}
	}
	return nil
}

func DeleteRoom(roomNumber int) {
	for index, room := range rooms {
		if room.RoomNumber == roomNumber {
			fmt.Println(len(rooms))
			rooms = slices.Delete(rooms, index, index+1)
			fmt.Println(rooms)
			return
		}
	}
	fmt.Println("Room not found !!!")
}

func AddRoom() {
	room := GetRoomFromInput()
	rooms = append(rooms, room)
}

func GetRoomFromInput() Room {

	var roomNumber int
	var bedCount int
	var roomType string // Single - Double - Standard
	var price float32

	fmt.Print("Please Enter room number : ")
	fmt.Scanln(&roomNumber)

	fmt.Print("Please Enter room type : ")
	fmt.Scanln(&roomType)

	fmt.Print("Please Enter bed count : ")
	fmt.Scanln(&bedCount)

	fmt.Print("Please Enter price : ")
	fmt.Scanln(&price)

	var room Room = Room{RoomNumber: roomNumber, BedCount: bedCount, Type: roomType, IsActive: true, Price: price}

	return room
}

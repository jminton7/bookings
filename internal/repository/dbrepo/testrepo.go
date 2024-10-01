package dbrepo

import (
	"time"

	"github.com/jminton7/bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

//inserts a reservation into the db
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil
}

//inserts a room restiction into DB
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) (error) {
	return nil
}

//returns true if availability exists for roomId or false if it doesnt
func (m *testDBRepo) SearchAvailabilityByDatesByRoomId(start,end time.Time, roomId int) (bool, error) {
	return false, nil
}

//searches all rooms for availabilities within date range and returns slice of rooms
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

//gets a room by Id
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	return room, nil
}
package repository

import (
	"time"

	"github.com/jminton7/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) (error)
	SearchAvailabilityByDates(start,end time.Time, roomId int) (bool, error) 
}
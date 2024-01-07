package sql

import (
	"log"
)

// Room représente une salle
type Room struct {
	RoomID          int
	RoomName        string
	RoomDescription string
	RoomStatus      bool
}

// GetRooms récupère la liste des salles
func GetRooms() ([]Room, error) {
	var rooms []Room

	db := InitDB()
	defer db.Close()

	stmt, err := db.Prepare("SELECT room_id, room_name, room_description, room_status FROM rooms WHERE room_status = ?")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(true)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		var room Room
		err := rows.Scan(&room.RoomID, &room.RoomName, &room.RoomDescription, &room.RoomStatus)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

/* 	script of the table events :
CREATE TABLE events (
    event_id INT NOT NULL AUTO_INCREMENT,
    host VARCHAR(255) NOT NULL,
    room_name VARCHAR(255) NOT NULL,
    event_certificate INT NOT NULL,
    event_title VARCHAR(255) NOT NULL,
    event_category VARCHAR(255) NOT NULL,
    event_description VARCHAR(255) NOT NULL,
    begin_date DATE NOT NULL, -- format: YYYY-MM-DD
    end_date DATE NOT NULL, -- format: YYYY-MM-DD
    begin_hour TIME NOT NULL, -- format: HH:MM:SS
    end_hour TIME NOT NULL, -- format: HH:MM:SS
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (event_id)
);	*/

// Event représente un événement
type Event struct {
	EventID          int
	Host             string
	RoomName         string
	EventCertificate int
	EventTitle       string
	EventCategory    string
	EventDescription string
	BeginDate        string
	EndDate          string
	BeginHour        string
	EndHour          string
}

// GetEvents récupère la liste des événements
func GetEvents() ([]Event, error) {
	var events []Event

	db := InitDB()
	defer db.Close()

	stmt, err := db.Prepare("SELECT event_id, host, room_name, event_certificate, event_title, event_category, event_description, begin_date, end_date, begin_hour, end_hour FROM events")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.EventID, &event.Host, &event.RoomName, &event.EventCertificate, &event.EventTitle, &event.EventCategory, &event.EventDescription, &event.BeginDate, &event.EndDate, &event.BeginHour, &event.EndHour)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

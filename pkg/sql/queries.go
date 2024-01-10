package sql

import (
	"fmt"
	"log"
	"net/http"
)

// Room représente une salle
type Room struct {
	RoomID          int
	RoomName        string
	RoomDescription string
	RoomStatus      bool
	RoomCapacity    int
	RoomTitle       string
	Message         string
}

// GetRooms récupère la liste des salles
func GetRooms() ([]Room, error) {
	var rooms []Room

	db := InitDB()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM rooms")
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
		var room Room
		err := rows.Scan(&room.RoomID, &room.RoomName, &room.RoomDescription, &room.RoomStatus, &room.RoomCapacity)
		if err != nil {
			fmt.Print("Erreur lors de la récupération des salles : ", err)
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

// Event représente un événement
type Event struct {
	EventID          int
	Host             string
	RoomName         string
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

	stmt, err := db.Prepare("SELECT event_id, host, room_name,  event_title, event_category, event_description, begin_date, end_date, begin_hour, end_hour FROM events")
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
		err := rows.Scan(&event.EventID, &event.Host, &event.RoomName, &event.EventTitle, &event.EventCategory, &event.EventDescription, &event.BeginDate, &event.EndDate, &event.BeginHour, &event.EndHour)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

// DisplayEvent pour afficher les informations de l'évènement sélectionné dont l'id est passé en paramètre de l'URL
func ModifyEvent(r *http.Request) ([]Event, error) {
	var events []Event

	db := InitDB()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM events WHERE event_id = ?")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	// Requête avec l'id dans l'URL
	rows, err := stmt.Query(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.EventID, &event.Host, &event.RoomName, &event.EventTitle, &event.EventCategory, &event.EventDescription, &event.BeginDate, &event.EndDate, &event.BeginHour, &event.EndHour)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

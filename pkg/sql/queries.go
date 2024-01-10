package sql

import (
	"fmt"
	"log"
	"net/http"
)

// Room représente une salle
type Room struct {
	RoomID, RoomCapacity                          int
	RoomName, RoomDescription, RoomTitle, Message string
	RoomStatus                                    bool
}

// Event représente un événement
type Event struct {
	EventID                                                                                                                     int
	Host, RoomName, EventTitle, EventCategory, EventDescription, BeginDate, EndDate, BeginHour, EndHour, Created_at, Updated_at string
}

// GetRooms récupère la liste des salles
func GetRooms() ([]Room, error) {
	// Déclaration d'un tableau de salles
	var rooms []Room

	db := InitDB()

	stmt, err := db.Prepare("SELECT * FROM rooms")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

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

// GetEvents récupère la liste des événements
func GetEvents() ([]Event, error) {
	var events []Event

	db := InitDB()

	stmt, err := db.Prepare("SELECT event_id, host, room_name,  event_title, event_category, event_description, begin_date, end_date, begin_hour, end_hour FROM events")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

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

// GetOneEvent récupère la liste des événements
func ModifyEvent(eventId int) ([]Event, error) {
	var events []Event

	db := InitDB()

	stmt, err := db.Prepare("SELECT event_id, host, room_name,  event_title, event_category, event_description, begin_date, end_date, begin_hour, end_hour FROM events WHERE event_id = ?")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	rows, err := stmt.Query(eventId)
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

// DeleteEvent pour supprimer un évènement dont l'id est passé en paramètre de la fonction
func DeleteEvent(r *http.Request, eventId int) error {
	db := InitDB()

	stmt, err := db.Prepare("DELETE FROM events WHERE event_id = ?")
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = stmt.Exec(eventId)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

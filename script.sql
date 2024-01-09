-- booking database creation with mariadb

CREATE DATABASE booking;

USE booking;

CREATE TABLE users (
    user_id INT NOT NULL AUTO_INCREMENT,
    firstname VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    pwd VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    scope INT NOT NULL DEFAULT 1,
    PRIMARY KEY (user_id)
);

CREATE TABLE rooms (
    room_id INT NOT NULL AUTO_INCREMENT,
    room_name VARCHAR(255) NOT NULL,
    room_description VARCHAR(255) NOT NULL,
    room_status INT NOT NULL,
    room_capacity INT NOT NULL,
    PRIMARY KEY (room_id)
);

CREATE TABLE events (
    event_id INT NOT NULL AUTO_INCREMENT,
    host VARCHAR(255) NOT NULL,
    room_name VARCHAR(255) NOT NULL,
    event_certificate VARCHAR(255) DEFAULT NULL,
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
);


-- create a superadmin user

INSERT INTO users (firstname, lastname, email, pwd, account_status, scope, token) VALUES ('superadmin', 'superadmin', 'lightyagamisere@gmail.com', 'superadmin', 'active', 1, '123456789');

-- create 10 random rooms in the database

INSERT INTO rooms (room_name, room_description, room_status, room_capacity) VALUES ('A07', 'Salle de réunion site:Nation 1', 1, 10);
INSERT INTO rooms (room_name, room_description, room_status, room_capacity) VALUES ('B12', 'Salle de cours site: Nation 1', 1, 10);
INSERT INTO rooms (room_name, room_description, room_status, room_capacity) VALUES ('Salle 15', 'Salle de cours site: Érard', 1, 10);
INSERT INTO rooms (room_name, room_description, room_status, room_capacity) VALUES ('Salle 12', 'Salle de cours site: Érard', 1, 10);


-- create 10 random events in the database

INSERT INTO events (host, room_name, event_certificate, event_title, event_category, event_description, begin_date, end_date, begin_hour, end_hour) VALUES ('superadmin', 'room1', 1, 'event1', 'event1', 'event1', '2021-01-01', '2021-01-01', '00:00:00', '00:00:00');
INSERT INTO events (host, room_name, event_certificate, event_title, event_category, event_description, begin_date, end_date, begin_hour, end_hour) VALUES ('superadmin', 'room2', 1, 'event2', 'event2', 'event2', '2021-01-01', '2021-01-01', '00:00:00', '00:00:00');
INSERT INTO events (host, room_name, event_certificate, event_title, event_category, event_description, begin_date, end_date, begin_hour, end_hour) VALUES ('superadmin', 'room3', 1, 'event3', 'event3', 'event3', '2021-01-01', '2021-01-01', '00:00:00', '00:00:00');
INSERT INTO events (host, room_name, event_certificate, event_title, event_category, event_description, begin_date, end_date, begin_hour, end_hour) VALUES ('superadmin', 'room4', 1, 'event4', 'event4', 'event4', '2021-01-01', '2021-01-01', '00:00:00', '00:00:00');
INSERT INTO events (host, room_name, event_certificate, event_title, event_category, event_description, begin_date, end_date, begin_hour, end_hour) VALUES ('superadmin', 'room5', 1, 'event5', 'event5', 'event5', '2021-01-01', '2021-01-01', '00:00:00', '00:00:00');
INSERT INTO events (host, room_name, event_certificate, event_title, event_category, event_description, begin_date, end_date, begin_hour, end_hour) VALUES ('superadmin', 'room6', 1, 'event6', 'event6', 'event6', '2021-01-01', '2021-01-01', '00:00:00', '00:00:00');
INSERT INTO events (host, room_name, event_certificate, event_title, event_category, event_description, begin_date, end_date, begin_hour, end_hour) VALUES ('superadmin', 'room7', 1, 'event7', 'event7', 'event7', '2021-01-01', '2021-01-01', '00:00:00', '00:00:00');
INSERT INTO events (host, room_name, event_certificate, event_title, event_category, event_description, begin_date, end_date, begin_hour, end_hour) VALUES ('superadmin', 'room8', 1, 'event8', 'event8', 'event8', '2021-01-01', '2021-01-01', '00:00:00', '00:00:00');
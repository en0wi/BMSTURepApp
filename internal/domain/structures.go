package domain

import "time"

type User struct {
	Id            int
	Given_name    string
	Family_name   string
	Middle_name   string
	Student_group string
	Phone_number  string
	Description   string
	TelegramTag   string
	VkLink        string
	Is_admin      bool
	Is_banned     bool
}

type Reservation struct {
	Id             int
	Status         string
	Datetime_start time.Time
	Datetime_end   time.Time
	Color          string
	Place          int
	Is_repeatable  bool
}

type Group struct {
	Id          int
	Group_name  string
	Description string
	Is_academy  bool
}

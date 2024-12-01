package domain

import "time"

type User struct {
	Id           int
	First_name   string
	Middle_name  string
	Last_name    string
	Num_of_group string
	Phone_number string
	Description  string
	Tg_tag       string
	Vk_tag       string
	Is_admin     bool
	Is_banned    bool
}

type Reservation struct {
	Id            int
	Status        string
	Group_id      int
	Start_time    time.Time
	End_time      time.Time
	Color         string
	Place         int
	Is_repeatable bool
}

type Group struct {
	Id          int
	Group_name  string
	Description string
	Is_academy  bool
}

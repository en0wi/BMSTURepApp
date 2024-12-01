package db

import (
	structures "BMSTURepApp/internal/domain"
	_ "github.com/lib/pq"
)

// Create_Reserv Create добавляет нового пользователя в базу данных
func (db *DB) Create_Reserv(reserv structures.Reservation) (int, error) {
	var id int
	err := db.conn.QueryRow("INSERT INTO reservation(id, status,group_id, start_time, end_time, color,place, is_repeatable) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", reserv.Id, reserv.Status, reserv.Group_id, reserv.Start_time, reserv.End_time, reserv.Color, reserv.Place, reserv.Is_repeatable).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db *DB) Read_Reservinfo(id int) (structures.Reservation, error) {
	var reserv structures.Reservation
	err := db.conn.QueryRow("SELECT id, status,group_id, start_time, end_time, color,place, is_repeatable FROM reservation WHERE id=$1", id).Scan(&reserv.Id, &reserv.Status, &reserv.Group_id, &reserv.Start_time, &reserv.End_time, &reserv.Color, &reserv.Place, &reserv.Is_repeatable)
	if err != nil {
		return structures.Reservation{}, err
	}
	return reserv, nil
}

// Update_Reservinfo Update обновляет информацию о пользователе
func (db *DB) Update_Reservinfo(reserv structures.Reservation) error {
	_, err := db.conn.Exec("UPDATE reservation SET status=$2, start_time=$3, end_time=$4, color=$5,place=$6, is_repeatable=$7 WHERE id=$1", reserv.Id, reserv.Status, reserv.Start_time, reserv.End_time, reserv.Color, reserv.Place, reserv.Is_repeatable)
	return err
}

// Delete_Reserv Delete удаляет пользователя по ID
func (db *DB) Delete_Reserv(id int) error {
	_, err := db.conn.Exec("DELETE FROM reservation WHERE id=$1", id)
	return err
}

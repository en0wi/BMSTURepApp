package storage

import (
	structures "BMSTURepApp/internal/domain"
	_ "github.com/lib/pq"
)

// CreateReserv Create добавляет нового пользователя в базу данных
func (db *DB) CreateReserv(reserv structures.Reservation) (int, error) {
	var id int
	err := db.conn.QueryRow("INSERT INTO reservation(id, status, datetime_start, datetime_end, color,place, is_repeatable) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id", reserv.Id, reserv.Status, reserv.Datetime_start, reserv.Datetime_end, reserv.Color, reserv.Place, reserv.Is_repeatable).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db *DB) ReadReservInfo(id int) (structures.Reservation, error) {
	var reserv structures.Reservation
	err := db.conn.QueryRow("SELECT id, status, datetime_start, datetime_end, color,place, is_repeatable FROM reservation WHERE id=$1", id).Scan(&reserv.Id, &reserv.Status, &reserv.Datetime_start, &reserv.Datetime_end, &reserv.Color, &reserv.Place, &reserv.Is_repeatable)
	if err != nil {
		return structures.Reservation{}, err
	}
	return reserv, nil
}

// Update обновляет информацию о пользователе
func (db *DB) UpdateReservInfo(reserv structures.Reservation) error {
	_, err := db.conn.Exec("UPDATE reservation SET status=$2, datetime_start=$3, datetime_end=$4, color=$5,place=$6, is_repeatable=$7 WHERE id=$1", reserv.Id, reserv.Status, reserv.Datetime_start, reserv.Datetime_end, reserv.Color, reserv.Place, reserv.Is_repeatable)
	return err
}

// Delete удаляет пользователя по ID
func (db *DB) DeleteReserv(id int) error {
	_, err := db.conn.Exec("DELETE FROM reservation WHERE id=$1", id)
	return err
}

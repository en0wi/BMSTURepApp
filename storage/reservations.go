package db

import (
  "database/sql"
   _"github.com/lib/pq"
   "time"
)
type Reservation struct {
  Id    int
  Status string
  Datetime_start time.Time
  Datetime_end time.Time
  Color string
  Place int
  Is_repeatable bool
}
type ReservDB struct {
  Conn *sql.DB
}

func NewDBReserv(dataSourceName string) (*ReservDB, error) {
  conn, err := sql.Open("postgres", dataSourceName)
  if err != nil {
    return nil, err
  }

  if err := conn.Ping(); err != nil {
    return nil, err
  }

  return &ReservDB{Conn: conn}, nil
}
// Create добавляет нового пользователя в базу данных
func (db *ReservDB) Create_Reserv(reserv Reservation) (int, error) {
  var id int
  err := db.Conn.QueryRow("INSERT INTO reservation(id, status, datetime_start, datetime_end, color,place, is_repeatable) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",reserv.Id, reserv.Status, reserv.Datetime_start, reserv.Datetime_end, reserv.Color, reserv.Place, reserv.Is_repeatable).Scan(&id)
  if err != nil {
    return 0, err
  }
  return id, nil
}

func (db *ReservDB) Read_Reservinfo(id int) (Reservation , error) {
  var reserv Reservation
  err := db.Conn.QueryRow("SELECT id, status, datetime_start, datetime_end, color,place, is_repeatable FROM reservation WHERE id=$1", id).Scan(&reserv.Id, &reserv.Status, &reserv.Datetime_start, &reserv.Datetime_end, &reserv.Color, &reserv.Place, &reserv.Is_repeatable)
  if err != nil {
    return Reservation{}, err
  }
  return reserv, nil
}

// Update обновляет информацию о пользователе
func (db *ReservDB) Update_Reservinfo(reserv Reservation) error {
  _, err := db.Conn.Exec("UPDATE reservation SET status=$2, datetime_start=$3, datetime_end=$4, color=$5,place=$6, is_repeatable=$7 WHERE id=$1", reserv.Id, reserv.Status, reserv.Datetime_start, reserv.Datetime_end, reserv.Color, reserv.Place, reserv.Is_repeatable)
  return err
}

// Delete удаляет пользователя по ID
func (db *ReservDB) Delete_Reserv(id int) error {
  _, err := db.Conn.Exec("DELETE FROM reservation WHERE id=$1", id)
  return err
}
 
 
func (db *ReservDB) CloseDBReserv() error {
  return db.Conn.Close()
}

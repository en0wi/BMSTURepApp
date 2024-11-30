package db

import (
  "database/sql"
   _"github.com/lib/pq"
)
type User struct {
  Id    int
  Given_name  string
  Family_name  string
  Middle_name  string
  Student_group  string
  Phone_number string
  Description  string
  Tg  string
  Vk  string
  Is_admin  bool
  Is_banned bool
}

type DB struct {
  Conn *sql.DB
}

// NewDB создает новое соединение с базой данных
func NewDB(dataSourceName string) (*DB, error) {
  conn, err := sql.Open("postgres", dataSourceName)
  if err != nil {
    return nil, err
  }

  if err := conn.Ping(); err != nil {
    return nil, err
  }

  return &DB{Conn: conn}, nil
}

// Create добавляет нового пользователя в базу данных
func (db *DB) Create_User(user User) (int, error) {
  var id int
  err := db.Conn.QueryRow("INSERT INTO users(id,given_name,family_name,middle_name,student_group,phone_number,description,tg,vk,is_admin,is_banned) VALUES($1, $2, $3, $4, $5, $6, $7, $8,$9,$10,$11) RETURNING id",user.Id, user.Given_name,user.Family_name, user.Middle_name,user.Student_group,user.Phone_number,user.Description, user.Tg,user.Vk,user.Is_admin,user.Is_banned).Scan(&id)
  if err != nil {
    return 0, err
  }
  return id, nil
}

func (db *DB) Read_Userinfo(id int) (User , error) {
  var user User
  err := db.Conn.QueryRow("SELECT id, given_name, family_name, middle_name, student_group, phone_number, description, tg, vk, is_admin, is_banned FROM users WHERE id=$1", id).Scan(&user.Id, &user.Given_name, &user.Family_name, &user.Middle_name, &user.Student_group, &user.Phone_number, &user.Description, &user.Tg, &user.Vk, &user.Is_admin, &user.Is_banned)
  if err != nil {
    return User{}, err
  }
  return user, nil
}

// Update обновляет информацию о пользователе
func (db *DB) Update_Userinfo(user User) error {
  _, err := db.Conn.Exec("UPDATE users SET given_name=$2, family_name=$3, middle_name=$4,student_group=$5,phone_number=$6,description=$7, tg=$8, vk=$9, is_admin=$10, is_banned=$11 WHERE id=$1", user.Id, user.Given_name, user.Family_name,user.Middle_name,user.Student_group,user.Phone_number,user.Description, user.Tg,user.Vk,user.Is_admin,user.Is_banned)
  return err
}

// Delete удаляет пользователя по ID
func (db *DB) Delete_User(id int) error {
  _, err := db.Conn.Exec("DELETE FROM users WHERE id=$1", id)
  return err
}

// Close закрывает соединение с базой данных
func (db *DB) Close() error {
  return db.Conn.Close()
}



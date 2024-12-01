package db

import (
  "database/sql"
   _"github.com/lib/pq"
)
type Group struct {
  Id    int
  Group_name string
  Description string
  Is_academy bool
}

type GroupDB struct {
  Conn *sql.DB
}

func NewDBGroup(dataSourceName string) (*GroupDB, error) {
  conn, err := sql.Open("postgres", dataSourceName)
  if err != nil {
    return nil, err
  }

  if err := conn.Ping(); err != nil {
    return nil, err
  }

  return &GroupDB{Conn: conn}, nil
}
// Create добавляет нового пользователя в базу данных
func (db *GroupDB) Create_Group(group Group) (int, error) {
  var id int
  err := db.Conn.QueryRow("INSERT INTO groups(id, group_name, description, is_academy) VALUES($1, $2, $3, $4) RETURNING id",group.Id, group.Group_name,group.Description, group.Is_academy).Scan(&id)
  if err != nil {
    return 0, err
  }
  return id, nil
}

func (db *GroupDB) Read_Groupinfo(id int) (Group, error) {
  var group Group
  err := db.Conn.QueryRow("SELECT id, group_name, description, is_academy FROM groups WHERE id=$1", id).Scan(&group.Id, &group.Group_name,&group.Description,&group.Is_academy)
  if err != nil {
    return Group{}, err
  }
  return group, nil
}

// Update обновляет информацию о пользователе
func (db *GroupDB) Update_Groupinfo(group Group) error {
  _, err := db.Conn.Exec("UPDATE groups SET group_name=$2, description=$3, is_academy=$4 WHERE id=$1", group.Id, group.Group_name,group.Description, group.Is_academy)
  return err
}

// Delete удаляет пользователя по ID
func (db *GroupDB) Delete_Group(id int) error {
  _, err := db.Conn.Exec("DELETE FROM groups WHERE id=$1", id)
  return err
}
 
 
func (db *GroupDB) CloseDb() error {
  return db.Conn.Close()
}

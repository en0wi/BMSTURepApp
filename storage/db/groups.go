package db

import (
	structures "BMSTURepApp/internal/domain"
)

// Create_Group Create добавляет нового пользователя в базу данных
func (db *DB) Create_Group(group structures.Group) (int, error) {
	var id int
	err := db.conn.QueryRow("INSERT INTO groups(id, group_name, description, is_academy) VALUES($1, $2, $3, $4) RETURNING id", group.Id, group.Group_name, group.Description, group.Is_academy).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db *DB) Read_Groupinfo(id int) (structures.Group, error) {
	var group structures.Group
	err := db.conn.QueryRow("SELECT id, group_name, description, is_academy FROM groups WHERE id=$1", id).Scan(&group.Id, &group.Group_name, &group.Description, &group.Is_academy)
	if err != nil {
		return structures.Group{}, err
	}
	return group, nil
}

// Update_Groupinfo Update обновляет информацию о пользователе
func (db *DB) Update_Groupinfo(group structures.Group) error {
	_, err := db.conn.Exec("UPDATE groups SET group_name=$2, description=$3, is_academy=$4 WHERE id=$1", group.Id, group.Group_name, group.Description, group.Is_academy)
	return err
}

// Delete_Group Delete удаляет пользователя по ID
func (db *DB) Delete_Group(id int) error {
	_, err := db.conn.Exec("DELETE FROM groups WHERE id=$1", id)
	return err
}

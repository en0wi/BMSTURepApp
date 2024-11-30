package storage

import (
	structures "BMSTURepApp/internal/domain"
	_ "github.com/lib/pq"
)

// CreateUser Create добавляет нового пользователя в базу данных
func (db *DB) CreateUser(user structures.User) (int, error) {
	var id int
	err := db.conn.QueryRow("INSERT INTO users(id,given_name,family_name,middle_name,student_group,phone_number,description,tg,vk,is_admin,is_banned) VALUES($1, $2, $3, $4, $5, $6, $7, $8,$9,$10,$11) RETURNING id", user.Id, user.Given_name, user.Family_name, user.Middle_name, user.Student_group, user.Phone_number, user.Description, user.TelegramTag, user.VkLink, user.Is_admin, user.Is_banned).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db *DB) ReadUserinfo(id int) (structures.User, error) {
	var user structures.User
	err := db.conn.QueryRow("SELECT id, given_name, family_name, middle_name, student_group, phone_number, description, tg, vk, is_admin, is_banned FROM users WHERE id=$1", id).Scan(&user.Id, &user.Given_name, &user.Family_name, &user.Middle_name, &user.Student_group, &user.Phone_number, &user.Description, &user.TelegramTag, &user.VkLink, &user.Is_admin, &user.Is_banned)
	if err != nil {
		return structures.User{}, err
	}
	return user, nil
}

// UpdateUserinfo Update обновляет информацию о пользователе
func (db *DB) UpdateUserinfo(user structures.User) error {
	_, err := db.conn.Exec("UPDATE users SET given_name=$2, family_name=$3, middle_name=$4,student_group=$5,phone_number=$6,description=$7, tg=$8, vk=$9, is_admin=$10, is_banned=$11 WHERE id=$1", user.Id, user.Given_name, user.Family_name, user.Middle_name, user.Student_group, user.Phone_number, user.Description, user.TelegramTag, user.VkLink, user.Is_admin, user.Is_banned)
	return err
}

// DeleteUser Delete удаляет пользователя по ID
func (db *DB) DeleteUser(id int) error {
	_, err := db.conn.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}

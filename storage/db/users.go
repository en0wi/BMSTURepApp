package db

import (
	structures "BMSTURepApp/internal/domain"
	_ "github.com/lib/pq"
)

// Create_User Create добавляет нового пользователя в базу данных
func (db *DB) Create_User(user structures.User) (int, error) {
	var id int
	err := db.conn.QueryRow("INSERT INTO users(id,first_name,middle_name,last_name,num_of_group,phone_number,description,tg_tag,vk_tag,is_admin,is_banned) VALUES($1, $2, $3, $4, $5, $6, $7, $8,$9,$10,$11) RETURNING id", user.Id, user.First_name, user.Middle_name, user.Last_name, user.Num_of_group, user.Phone_number, user.Description, user.Tg_tag, user.Vk_tag, user.Is_admin, user.Is_banned).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db *DB) Read_Userinfo(id int) (structures.User, error) {
	var user structures.User
	err := db.conn.QueryRow("SELECT id, first_name, middle_name, last_name, num_of_group, phone_number, description, tg_tag, vk_tag, is_admin, is_banned FROM users WHERE id=$1", id).Scan(&user.Id, &user.First_name, &user.Middle_name, &user.Last_name, &user.Num_of_group, &user.Phone_number, &user.Description, &user.Tg_tag, &user.Vk_tag, &user.Is_admin, &user.Is_banned)
	if err != nil {
		return structures.User{}, err
	}
	return user, nil
}

// Update обновляет информацию о пользователе
func (db *DB) Update_Userinfo(user structures.User) error {
	_, err := db.conn.Exec("UPDATE users SET first_name=$2, middle_name=$3, last_name=$4,num_of_group=$5,phone_number=$6,description=$7, tg_tag=$8, vk_tag=$9, is_admin=$10, is_banned=$11 WHERE id=$1", user.Id, user.First_name, user.Middle_name, user.Last_name, user.Num_of_group, user.Phone_number, user.Description, user.Tg_tag, user.Vk_tag, user.Is_admin, user.Is_banned)
	return err
}

// Delete удаляет пользователя по ID
func (db *DB) Delete_User(id int) error {
	_, err := db.conn.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}

// Close закрывает соединение с базой данных
func (db *DB) Close() error {
	return db.conn.Close()
}

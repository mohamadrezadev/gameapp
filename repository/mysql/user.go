package mysql

import (
	"GameApp/entity"
	"database/sql"
	"fmt"
	"time"
)

func (d *MySqlDb) IsPhoneNumberUinc(PhoneNumber string) (bool, error) {
	row := d.db.QueryRow(`select * from users where phone_number =? `, PhoneNumber)
	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, fmt.Errorf("can't scan query result: %w", err)
	}
	return false, nil
}

func (d *MySqlDb) RegisterUser(u entity.User) (entity.User, error) {
	res,err:=d.db.Exec(`insert into users(name,phone_number,password) values(?,?,?)`,u.Name,u.PhoneNumber,u.Password)
	if err!=nil{
		return entity.User{}, fmt.Errorf("can't execute command: %w", err)

	}
	id,_:=res.LastInsertId()
	u.ID=uint(id)

	return u,nil 

}
func (d *MySqlDb) GetUserByPhoneNumber(PhoneNumber string) (entity.User ,bool,error){
	row :=d.db.QueryRow(`select * from users where phone_number =?`,PhoneNumber)
	user ,err:=scanUser(row)
	if err!=nil{
		if err== sql.ErrNoRows{
			return entity.User{}, false, nil
		}
		return entity.User{}, false, fmt.Errorf("can't scan query result: %w", err)
	}
	return user, true, nil
}

func (d *MySqlDb) GetUserById(userId uint)(entity.User,error){
	row :=d.db.QueryRow(`select * from users where id =?`,userId)
	user ,err:=scanUser(row)
	if err!=nil{
		if err== sql.ErrNoRows{
			return entity.User{}, fmt.Errorf("record not found")
		}
		return entity.User{}, fmt.Errorf("can't scan query result: %w", err)

	}
	return user, nil
}

func scanUser(row *sql.Row) (entity.User, error) {
	// var createdAt []uint8
	var createdAt time.Time
	var user entity.User

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt)
	fmt.Println("createdAt", createdAt)
	return user, err
}

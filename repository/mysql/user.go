package mysql

import (
	"GameApp/entity"
	"GameApp/pkg/errmsg"
	"GameApp/pkg/richerror"
	"database/sql"
	"fmt"
	"time"
)

func (d *MySqlDb) IsPhoneNumberUnique(PhoneNumber string) ( bool, error) {

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
	res, err := d.db.Exec(`insert into users(name,phone_number,password) values(?,?,?)`, u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command: %w", err)

	}
	id, _ := res.LastInsertId()
	u.ID = uint(id)

	return u, nil

}
func (d *MySqlDb) GetUserByPhoneNumber(PhoneNumber string) (entity.User, error) {
	const op = "mysql.GetUserByPhoneNumber"

	row := d.db.QueryRow(`select * from users where phone_number =?`, PhoneNumber)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}
		return entity.User{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return user, nil
}

func (d *MySqlDb) GetUserById(userId uint) (entity.User, error) {
	const op = "mysql.GetUserByID"

	row := d.db.QueryRow(`select * from users where id =?`, userId)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}
		return entity.User{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)

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

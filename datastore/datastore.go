package datastore

import (
	"database/sql"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"Gofr_Zopsmart/model"
)

type user struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	PhoneNo     int    `json:"phone_number"`
	bill_due     int    `json:"bill_due"`
	BillNo      int    `json:"bill_number"`
}

func New() *user {
	return &user{}
}
func (s *user) GetByID(ctx *gofr.Context, id string) (*model.user, error) {
	var resp model.user

	err := ctx.DB().QueryRowContext(ctx, " SELECT user_id,user_name,phone_number,bill_due,bill_number FROM user where user_id= ?", id).
		Scan(&resp.userId, &resp.userName, &resp.PhoneNo, &resp.bill_due, &resp.bill_number)
	switch err {
	case sql.ErrNoRows:
		return &model.user{}, errors.EntityNotFound{Entity: "entries", ID: id}
	case nil:
		return &resp, nil
	default:
		return &model.user{}, err
	}
}
func (s *user) Create(ctx *gofr.Context, user *model.user) (*model.user, error) {
	var resp model.user
	result, err := ctx.DB().ExecContext(ctx, "INSERT INTO user (user_id, user_name, phone_number, bill_due, bill_number) VALUES (?,?,?,?,?)",
		user.userID, user.userName, user.PhoneNo, user.bill_due, user.bill_number)
	if err != nil {
		return &model.user{}, errors.DB{Err: err}
	}
	// Remove the line since lastInsertID is not being used
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return &model.user{}, errors.DB{Err: err}
	}
	// Set the ID in the response using entries.BookId
	resp.userId = int(lastInsertID)
	resp.userName = user.userName
	resp.PhoneNo = user.PhoneNo
	resp.bill_due = user.bill_due
	resp.bill_number = user.bill_number
	return &resp, nil
}
func (s *user) Update(ctx *gofr.Context, user *model.user) (*model.user, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE user SET  user_name= ? ,phone_number= ? , bill_due = ? , bill_number = ?   WHERE user_id= ?",
		user.userName, user.PhoneNo, user.bill_due, user.bill_number, user.userId)
	if err != nil {
		return &model.user{}, errors.DB{Err: err}
	}
	return user, nil
}
func (s *user) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM user where user_id= ?", id)
	if err != nil {
		return errors.DB{Err: err}
	}
	return nil
}

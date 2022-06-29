package users

import (
	"context"

	"github.com/JesusJMM/blog-plat-go/postgres"
	"github.com/vingarcia/ksql"
)

type UserRepository interface {
	Create(postgres.User) (postgres.User, error)
	ChangePassword(int, string) error
}

type UserRepo struct {
	db  *ksql.DB
	ctx context.Context
}

func New(db *ksql.DB, ctx context.Context) *UserRepo{
  return &UserRepo{
    db: db,
    ctx: ctx,
  }

}

func (r UserRepo) Create(user postgres.User) (postgres.User, error) {
	outUser := user
	pass, err := EncryptPassword(user.Password)
	if err != nil {
		return postgres.User{}, err
	}
	outUser.Password = pass
	err = r.db.Insert(r.ctx, postgres.UserTable, &outUser)
	return outUser, err
}

type updateUserPassword struct {
	ID       int    `ksql:"user_id"`
	Password *string `ksql:"password"`
}

func (r UserRepo) ChangePassword(id int, password string) error {
	pass, err := EncryptPassword(password)
	if err != nil {
		return err
	}
  return r.db.Patch(r.ctx, postgres.UserTable, updateUserPassword{
    ID: id,
    Password: &pass,
  })
}

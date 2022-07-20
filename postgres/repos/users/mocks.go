package users

import "github.com/JesusJMM/blog-plat-go/postgres"

type MockedUserRepo struct {
  CreateFn func(postgres.User) (postgres.User, error)
	ChangePasswordFn func(int, string) error
}

func NewMocked() MockedUserRepo {
  return MockedUserRepo{}
}

func (r MockedUserRepo) Create(user postgres.User) (postgres.User, error) {
  if r.CreateFn != nil {
    return r.CreateFn(user)
  }
  user.ID = 3
  return user, nil
}

func (r MockedUserRepo) ChangePassword(id int, password string) error{
  if r.ChangePasswordFn != nil {
    return r.ChangePasswordFn(id, password)
  }
  return nil
}

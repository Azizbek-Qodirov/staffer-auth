package service

import (
	"auth-service/models"
	"auth-service/postgresql/managers"
	"database/sql"

	"github.com/google/uuid"
)

type UserService struct {
	UM managers.UserManager
}

func NewUserService(conn *sql.DB) *UserService {
	return &UserService{UM: *managers.NewUserManager(conn)}
}

func (u *UserService) Register(req *models.RegisterReq) error {
	req.ID = uuid.NewString()
	req.Role = "user"
	return u.UM.Register(*req)
}

func (u *UserService) GetProfile(req *models.GetProfileReq) (*models.GetProfileResp, error) {
	return u.UM.Profile(*req)
}

func (u *UserService) EmailExists(email string) (bool, error) {
	return u.UM.EmailExists(email)
}

func (u *UserService) GetByID(id *models.GetProfileByIdReq) (*models.GetProfileByIdResp, error) {
	return u.UM.GetByID(id)
}

func (u *UserService) ChangeRole(req *models.ChangeRoleReq) error {
	return u.UM.ChangeRole(*req)
}

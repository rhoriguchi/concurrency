package controller

import (
	"fmt"
	"pkgArch/entity"
)

type userController struct {
}

func NewUserController() *userController {
	return &userController{}
}

func (u *userController) Notify(usr entity.User, msg string) {
	fmt.Printf("message sent to %v: %v\n", usr.Name, msg)
}

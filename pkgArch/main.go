package main

import (
	"fmt"
	"pkgArch/controller"
	"pkgArch/entity"
	"time"
)

/*
	Inspired by the article: https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1#.n2t1m3l6g

	GROUP BY FUNCTION

	The current package layout approach in this project is to group code by it’s functional type.
	For example, all models (=entity) go in one package, controllers go in another.

	There are two issues with this approach though.
		1. First, names are atrocious. We end up with type names like controller.UserController where we are
		   duplicating the package name in the type‘s name.

   		2. The bigger issue, however, is circular dependencies. Our different functional types may need to reference
		   each other. This only works if we have one-way dependencies but many times applications are not that
		   simple.

	GROUP BY MODULE

	Another approach would be grouping code by module instead of by function. For example, you may have a users package
	and an accounts package. This approach has the same issues: complicated naming with names like users.User and circular
	dependencies.

	TODO:

	Change the package layout to adhere to these 4 principles:

	1. Root package is for domain types (contradicts CLEAN architecture)
	2. Group subpackages by dependency (AccountController uses UserController)
	3. Use a shared mock subpackage
	4. Main package ties together dependencies

	As a side-dish attempt to fix the potential deadlock mentioned in controller/account.go
*/

func main() {
	usr1 := &entity.User{Name: "John", Email: "j@connor.name"}
	usr2 := &entity.User{Name: "Clara", Email: "c@hiller.name"}

	acnt1 := &entity.Account{
		Balance: 100,
		Owner:   usr1,
	}

	acnt2 := &entity.Account{
		Balance: 0,
		Owner:   usr2,
	}

	uc := controller.NewUserController()
	ac := &controller.AccountController{UserCtl: uc}

	err := ac.Transfer(acnt2, acnt1, -12, time.Now().Add(time.Millisecond*30))
	if err != nil {
		fmt.Printf("Transaction failed: %v\n", err)
	} else {
		fmt.Println("Transaction successful.")
	}
}

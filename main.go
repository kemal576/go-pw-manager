package main

import (
	"github.com/kemal576/go-pw-manager/repository"
	"github.com/kemal576/go-pw-manager/utils"
	_ "github.com/lib/pq"
)

func main() {
	db := repository.Conn()
	defer db.Close()

	userRepository, err := repository.NewUserRepository(db)
	utils.ErrorCheck(err)

	users, err2 := userRepository.GetAll()
	utils.ErrorCheck(err2)

	for _, u := range users {
		println("ID:", u.Id)
		println("FirstName:", u.FirstName)
		println("")
	}

}

package main

import (
	"fmt"
	"log"

	"github.com/kemal576/go-pw-manager/repository"
	"github.com/kemal576/go-pw-manager/service"
	"github.com/kemal576/go-pw-manager/utils"
	_ "github.com/lib/pq"
)

func main() {
	db, err := repository.Conn()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	userRepository, err := repository.NewUserRepository(db)
	utils.ErrorCheck(err)
	userService := service.NewUserService(userRepository)
	user := userService.GetById(1)
	fmt.Printf("Adı: %s\nSoyadı: %s", user.FirstName, user.LastName)

	/*userService.Create(&models.User{Id: 1, FirstName: "Test", LastName: "Testtt",
	Email: "deneme@gmail.com", Password: "şifre123"})*/

}

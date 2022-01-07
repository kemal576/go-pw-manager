package app

import (
	"github.com/kemal576/go-pw-manager/models"
	"github.com/kemal576/go-pw-manager/repository"
)

func GetAllLogins(repo repository.LoginRepository) ([]models.Login, error) {
	logins, err := repo.GetAll()
	if err != nil {
		return logins, err
	}

	return logins, nil
}

func GetLoginsByUserId(id int, repo repository.LoginRepository) ([]models.Login, error) {
	logins, err := repo.GetLoginsByUserId(id)
	if err != nil {
		return logins, err
	}

	for i, _ := range logins {
		var err error
		logins[i].Identity, err = Decrypt(logins[i].Identity)
		logins[i].Password, err = Decrypt(logins[i].Password)

		if err != nil {
			return nil, err
		}
	}

	return logins, nil
}

func GetLoginByUrl(id int, url string, repo repository.LoginRepository) (models.Login, error) {
	login, err := repo.GetLoginByUrl(id, url)
	if err != nil {
		return login, err
	}

	login.Identity, err = Decrypt(login.Identity)
	login.Password, err = Decrypt(login.Password)

	if err != nil {
		return login, err
	}

	return login, nil
}

func CreateLogin(loginDto models.LoginDTO, repo repository.LoginRepository) (int, error) {
	login := models.ToLogin(loginDto)
	var err error
	var id int

	login.Identity, err = Encrypt(login.Identity)
	login.Password, err = Encrypt(login.Password)

	if err != nil {
		return id, err
	}

	id, err = repo.Create(login)
	if err != nil {
		return id, err
	}

	return id, nil
}

func UpdateLogin(login models.Login, repo repository.LoginRepository) (int, error) {
	var err error
	var id int

	login.Identity, err = Encrypt(login.Identity)
	login.Password, err = Encrypt(login.Password)

	if err != nil {
		return id, err
	}

	err = repo.Update(&login)
	if err != nil {
		return id, err
	}

	return id, nil
}

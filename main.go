package main

import (
	"log"

	"github.com/kemal576/go-pw-manager/internal/app"
	"github.com/kemal576/go-pw-manager/internal/secret"
	_ "github.com/lib/pq"
)

func main() {
	/*dbconn, err := repository.Conn()
	if err != nil {
		log.Panic(err)
	}
	defer dbconn.Close()

	db := repository.New(dbconn)
	router := router.New(*db)

	println("Listening from http://localhost:5764")
	log.Fatal(http.ListenAndServe(":5764", router.Router))*/

	keyByte, err := secret.ReadSecret("aes", "enc_key")
	if err != nil {
		log.Fatal(err)
	}

	secretText := "ÇOK GİZLİ BİLGİ"
	println("ŞİFRELENECEK METİN:", secretText)
	encText, err := app.Encrypt(secretText, string(keyByte))
	if err != nil {
		log.Fatal(err)
	}
	println("ŞİFRELENMİŞ METİN:", encText)
	decText, err := app.Decrypt(encText, string(keyByte))
	if err != nil {
		log.Fatal(err)
	}

	println("ŞİFRESİ ÇÖZÜLMÜŞ METİN:", decText)
}

package seeder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"mercafacil-challenge/api/models"

	"github.com/jinzhu/gorm"
)

type Contact struct {
	Name      string `json:"name"`
	Cellphone string `json:"cellphone"`
}

type Contacts struct {
	Contacts []Contact `json: "contacts"`
}

func getContacts(client string) Contacts {
	jsonFile, err := os.Open("contacts-" + client + ".json")
	if err != nil {
		fmt.Printf("An error has occurred while loading json file: %s\n", err)
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	var contacts Contacts
	err = json.Unmarshal([]byte(data), &contacts)
	if err != nil {
		fmt.Printf("An error has occurred while unmarshal json object: %s\n", err)
	}
	return contacts
}

func mapToUser(contact *Contact, client string) models.User {
	var user models.User
	user.Name = contact.Name
	user.Cellphone = contact.Cellphone
	user.Prepare(client)
	return user
}

func Load(db *gorm.DB, client string) {
	log.Println("Loading contacts", client)
	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("Cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("Cannot migrate table: %v", err)
	}

	contacts := getContacts(client)
	for i := range contacts.Contacts {
		log.Println("Creating contact", i)
		contact := &contacts.Contacts[i]
		user := mapToUser(contact, client)

		err = db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			log.Fatalf("Cannot seed users table: %v", err)
		}
	}
}

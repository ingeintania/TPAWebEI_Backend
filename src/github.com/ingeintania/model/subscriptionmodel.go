package model

import (
	"github.com/ingeintania/database"
	//"github.com/ingeintania/db"
	"github.com/ingeintania/email"
	"log"
	"time"
)

type Subscription struct {
	ID int  `gorm:"PRIMARY_KEY"`
	Email string `gorm:"VARCHAR(100)"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time	`sql:index`
}

func init() {
	db, error := database.Connect()

	if error != nil{
		panic("No Conection")
	}
	db.AutoMigrate(&Subscription{})

	log.Println("Initialize Subscription Success")
}

func InsertNewSubscription(newEmail string) *Subscription {
	db, error := database.Connect()

	if error != nil{
		panic("No Conection")
	}

	newSubscription := &Subscription{
		Email: newEmail,
	}
	db.Save(newSubscription)

	to := []string{newEmail}
	email.SubscriptionEmail(to)

	log.Println("Insert New Subscription Success")
	return newSubscription
}

func SendSubscriptionToAll() []Subscription {
	db, error := database.Connect()
	defer db.Close()

	if error != nil{
		panic("No Conection")
	}


	var subscription []Subscription
	db.Find(&subscription)

	for i, _ := range subscription {
		to := []string{subscription[i].Email}
		email.SubscriptionEmail(to)
	}

	return subscription
}
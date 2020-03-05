package model

import "github.com/ingeintania/database"

type Order struct {
	Order_id			int		`gorm:"primary_key"`
	User_name		 	string	`gorm:"type:varchar(1000)"`
	User_email			string	`gorm:"type:varchar(1000)"`
	User_phone			string	`gorm:"type:varchar(1000)"`
	Ticket_quantity		int		`gorm:"type:int"`
	Ticket_title		string	`gorm:"type:varchar(1000)"`
	Ticket_name			string	`gorm:"type:varchar(1000)"`
	Ticket_nationality	string	`gorm:"type:varchar(1000)"`
	Order_payment_method	string	`gorm:"type:varchar(1000)"`
	Order_promo_code	string	`gorm:"type:varchar(1000)"`
	Order_price			int		`gorm:"type:int"`

}

func init(){
	db, error := database.Connect()

	if error != nil{
		panic("No Conection")
	}
	db.AutoMigrate(&Order{})
}

func GetAllOrder()([]Order, error){
	db, error := database.Connect()

	if error != nil{
		return nil, error
	}
	var allOrder []Order

	db.Find(&allOrder)

	defer db.Close()
	return allOrder, nil
}

func CreateOrder(user_name string, user_email string,
	user_phone string, ticket_quantity int,
	ticket_title string, ticket_name string,
	ticket_nationality string, order_payment_method string,
	order_promo_code string, order_price int)(*Order, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var order = Order{
		User_name:user_name,
		User_email:user_email,
		User_phone:user_phone,
		Ticket_quantity:ticket_quantity,
		Ticket_title:ticket_title,
		Ticket_name:ticket_name,
		Ticket_nationality:ticket_nationality,
		Order_payment_method:order_payment_method,
		Order_promo_code:order_promo_code,
		Order_price:order_price,
	}
	if db.NewRecord(order){
		db.Create(&order)
	}

	return &order,nil
}

package model

import (
	"fmt"
	"github.com/ingeintania/database"
)

type Blog struct {
	Blog_id				int			`gorm:"primary_key"`
	Blog_image_path 	string	`gorm:"type:varchar(100)"`
	Blog_title			string	`gorm:"type:varchar(100)"`
	Blog_content		string	`gorm:"type:varchar(100)"`
	Blog_viewers		int		`gorm:"type:int"`
	Blog_category		string	`gorm:"type:varchar(100)"`
	Blog_writer			string	`gorm:"type:varchar(100)"`
}

func init(){
	db, error := database.Connect()

	if error != nil{
		panic("No Conection")
	}
	db.AutoMigrate(&Blog{})
}

func GetAllBlog()([]Blog, error){
	db, error := database.Connect()

	if error != nil{
		return nil, error
	}
	var allBlog []Blog

	db.Find(&allBlog)

	defer db.Close()
	return allBlog, nil

}

func CreateBlog(image_path string, title string, content string, viewers int, category string, writer string)(*Blog, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var blog = Blog{
		Blog_image_path: image_path,
		Blog_title:      title,
		Blog_content:    content,
		Blog_viewers:    viewers,
		Blog_category:   category,
		Blog_writer:     writer,
	}
	if image_path == ""{
		fmt.Print(image_path)
		panic("DB SALAH")
	}
	if title == ""{
		fmt.Print(title)
		panic("DB SALAH")
	}
	if content == ""{
		fmt.Print(content)
		panic("DB SALAH")
	}
	if viewers == 0{
		fmt.Print(title)
		panic("DB SALAH")
	}
	if category == ""{
		fmt.Print(title)
		panic("DB SALAH")
	}
	if writer == ""{
		fmt.Print(title)
		panic("DB SALAH")
	}

	if db.NewRecord(blog){
		db.Create(&blog)
		fmt.Print("Blog has been added!")
	}

	return &blog,nil
}

func UpdateBlog(id int, image_path string, title string, content string, viewers int, category string, writer string)(*Blog, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var blog = Blog{
		Blog_id:id,
	}
	if image_path == ""{
		fmt.Print(image_path)
		panic("DB SALAH")
	}
	if title == ""{
		fmt.Print(title)
		panic("DB SALAH")
	}
	if content == ""{
		fmt.Print(content)
		panic("DB SALAH")
	}
	if viewers == 0{
		fmt.Print(title)
		panic("DB SALAH")
	}
	if category == ""{
		fmt.Print(title)
		panic("DB SALAH")
	}
	if writer == ""{
		fmt.Print(title)
		panic("DB SALAH")
	}

	db.Model(&blog).Where("blog_id=?", id).
		Update(map[string]interface{}{"blog_image_path": image_path, "blog_title":title,
			"blog_content":content, "blog_viewers":viewers,
			"blog_category":category, "blog_writer":writer,})
	fmt.Print("Blog has been updated!")

	return &blog,nil
}

func DeleteBlog(id int) (*Blog, error) {
	db, error := database.Connect()

	if error !=nil{
		return nil, error
	}
	defer db.Close()
	var blog = Blog{
		Blog_id:id,
	}

	db.Where("blog_id=?", id).Find(&blog)
	db.Delete(blog)

	return &blog,nil
}
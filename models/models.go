package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	//"github.com/astaxie/beedb"
	_ "github.com/Go-SQL-Driver/MySQL" 
)

type User struct {
	UserId	int
	UserName	string
	UserPassword	string
	UserIntroduction	string
}

// func getLink() beedb.Model {
// 	db, err := sql.Open("mysql", "root:root@tcp(192.168.1.81:3306)/test_my?charset=utf8")
// 	if err != nil {
// 		panic(err)
// 	}
// 	orm := beedb.New(db)
// 	return orm
// }
func initMysql() *sql.DB {  
	//打开数据库连接Open(驱动名,连接字符串)  
	db, err := sql.Open("mysql", "root:root@tcp(192.168.34.9:3306)/myapp?charset=utf8")  
	if err != nil {  
		log.Fatal(err)  
	}  
	return db  
  }
func FindUser(user User) *User{
	db := initMysql()  
	defer db.Close()
	row, err := db.Query("select * from `myapp`.`user`")  
    if err != nil {  
        log.Fatal(err)  
    }  
    var user1 User  
    for row.Next() {  
        row.Scan(&user1.UserId,&user1.UserName, &user1.UserPassword,&user1.UserIntroduction)  
	   // log.Println("id:", id, ",name:", name, "password:", password,"introduction:",introduction)  
	   if (user1.UserName == user.UserName) && (user1.UserPassword == user.UserPassword){
		   return &user1
	   }
	}
	return nil
	
}
func SaveUser(user User) error {
	//orm := getLink()
	db := initMysql()  
    defer db.Close()
	fmt.Println(user)
	//err := orm.Save(&user)
	_, err := db.Exec("insert into `myapp`.`user`(userName,userPassword,userIntroduction) values(?,?,?)", user.UserName,user.UserPassword,user.UserIntroduction)  
    //c, _ := result.RowsAffected()  
    //log.Println("add affected rows:", c)
	return err
}

func ValidateUser(user User) error {
	//orm := getLink()
	db := initMysql()  
	defer db.Close()
	//var u User
	//orm.Where("username=? and pwd=?", user.Username, user.Pwd).Find(&u)  
    row, err := db.Query("select userName,userPassword from `myapp`.`user`")  
    if err != nil {  
        log.Fatal(err)  
    }  
    var name string = ""      
    var password string = ""  
    for row.Next() {  
        row.Scan(&name, &password)  
	   // log.Println("id:", id, ",name:", name, "password:", password,"introduction:",introduction)  
	   if (name == user.UserName) && (password == user.UserPassword){
		   return nil
	   }
	}
	return errors.New("用户名或密码错误！")
	
}
package main

import (
	"fmt"
	"database/sql"
	"log"
	"bufio"
	"os"
	_"github.com/go-sql-driver/mysql"
)

type Users struct {
	id, name string
}

func Input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func ConnectDatabase() *sql.DB {
	dbDriver := "mysql"
	dbUser := "kiettran"
	dbPassword := "Kiet@123456"
	dbName := "testDB"

	// Connect to database 
	db, err := sql.Open(dbDriver, dbUser + ":" + dbPassword + "@tcp(127.0.0.1:3306)/" + dbName)

	if(err != nil) {
		panic(err.Error())
	} else {
		log.Println("Connect to database successfully!")
	}

	return db
}

func PrintMenu() {
	fmt.Println("______________________________")
	fmt.Println("|Menu Options:               |")
	fmt.Println("|1. Select user              |")
	fmt.Println("|2. Insert user              |")
	fmt.Println("|3. Update user              |")
	fmt.Println("|4. Delete user              |")
	fmt.Println("|Other. Exit                 |")
	fmt.Println("|____________________________|")
}

func Select(db *sql.DB, name string) {
	query := fmt.Sprintf("SELECT * FROM user WHERE name = '%s'", name)
	
	result,err := db.Query(query)
	if(err != nil) {
		panic(err.Error())
	}

	fmt.Printf("Id\tUsername\n")
	for result.Next() {
		var user Users
		result.Scan(&user.id, &user.name)
		fmt.Printf("%s\t%s\n", user.id, user.name)
	}
}

func Insert(db *sql.DB, user Users) {
	insert := fmt.Sprintf("INSERT INTO user(id,name) VALUES('%s','%s')", user.id, user.name)

	_,err := db.Exec(insert)
	if(err != nil) {
		panic(err.Error())
	}

	log.Printf("Insert User id = %s, name = %s successfully", user.id, user.name)
}

func Update(db *sql.DB, id string, name string) {
	update := fmt.Sprintf("UPDATE user SET name = '%s' WHERE id = '%s'", name, id)

	_,err := db.Exec(update)
	if(err != nil) {
		panic(err.Error())
	}

	log.Printf("Updated User with id = %s", id)
}

func Delete(db *sql.DB, id string) {
	delete := fmt.Sprintf("DELETE FROM user WHERE id = '%s'", id)

	_,err := db.Exec(delete)
	if(err != nil) {
		panic(err.Error())
	}
	
	log.Printf("Deleted User with id = %s", id)
}


func main() {
	fmt.Println("Start with MYSQL Server ...")

	db := ConnectDatabase()
	defer db.Close()

	var opt int
	flag := true
	for flag {
		PrintMenu()
		fmt.Printf("Enter option: ")
		fmt.Scanf("%d", &opt)

		switch(opt) {
		    case 1: {
				var name string
				fmt.Printf("Enter name of user: ")
				name = Input()
				fmt.Println(name)
				Select(db, name)
			}

		    case 2: {
				var user Users
				fmt.Printf("Enter id: ")
				fmt.Scanf("%s", &user.id)
				fmt.Printf("Enter name: ")
				user.name = Input()
				Insert(db, user)
		    }

		    case 3: {
				var id, up_name string
				fmt.Printf("Enter id: ")
				fmt.Scanf("%s", &id)
				fmt.Printf("Enter name update: ")
				up_name = Input()
				Update(db, id, up_name)
		    }

		    case 4: {
				var id string
				fmt.Printf("Enter id: ")
				fmt.Scanf("%s", &id)
				Delete(db, id)
		    }

		    default: {
				log.Printf("Exit!\n")
				flag = false
			} 
		}
	}
}
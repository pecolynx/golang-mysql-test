package main

import (
	"fmt"

	"github.com/pecolynx/golang-mysql-test/pkg/application/administrator"
	"github.com/pecolynx/golang-mysql-test/pkg/domain"
	"github.com/pecolynx/golang-mysql-test/pkg/gateway"
)

func main() {
	db, err := gateway.OpenMySQL(&gateway.MySQLConfig{
		User:     "user",
		Password: "password",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "testdb",
	})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := gateway.NewUserRepository(db)
	id, err := userRepo.AddUser(domain.NewUserAddParameter("taro", 15))
	if err != nil {
		panic(err)
	}

	// time.Sleep(3 * time.Second)
	fmt.Println(id)
	updated, err := userRepo.UpdateUser(id, domain.NewUserUpdateParameter("jiro", 16))
	if err != nil {
		panic(err)
	}

	fmt.Println(updated)
	fmt.Println(userRepo)

	deleted, err := userRepo.RemoveUser(100)
	if err != nil {
		panic(err)
	}
	fmt.Println(deleted)

	users, err := userRepo.FindUnderageUsers(10, 0)
	if err != nil {
		panic(err)
	}
	for _, u := range users {
		fmt.Println(u)
	}

	repository := gateway.NewRepository(db)
	repositoryFactory := gateway.NewRepositoryFactory()
	adminUserService := administrator.NewUserService(repository, repositoryFactory)
	adminUserService.FindUserByID(id)
}

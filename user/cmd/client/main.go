package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/HotPotatoC/twitter-clone/user/rpc/user"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	client = user.NewUserServiceProtobufClient("http://localhost:"+os.Getenv("PORT"), &http.Client{})
	ctx    = context.Background()
)

func getUserByID(id string) {
	user, err := client.FindUserByID(ctx, &user.FindUserByIDRequest{UserId: id})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", user)
}

func getUserByEmail(email string) {
	user, err := client.FindUserByEmail(ctx, &user.FindUserByEmailRequest{Email: email})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", user)
}

func createUser() *user.User {
	now := time.Now().String()
	req := &user.CreateUserRequest{
		Name:             now,
		ScreenName:       "user_" + now,
		Password:         "123123",
		Email:            "user_" + now + "@mail.com",
		Bio:              "lorem ipsum",
		Location:         "Saint. " + now,
		Website:          "user_" + now + ".com",
		ProfileImageUrl:  "",
		ProfileBannerUrl: "",
		BirthDate:        timestamppb.Now(),
	}

	response, err := client.CreateUser(ctx, req)
	if err != nil {
		panic(err)
	}

	return response.User
}

func deleteUser(id string) bool {
	response, err := client.DeleteUser(ctx, &user.DeleteUserRequest{UserId: id})
	if err != nil {
		panic(err)
	}

	return response.GetSuccess()
}

func main() {
	user := createUser()
	getUserByEmail(user.GetEmail())
	deleteUser(user.GetUserId())
	getUserByID(user.GetUserId())
}

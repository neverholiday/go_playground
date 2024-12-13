package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	googleoauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func main() {
	jsonKey, err := os.ReadFile("credentials.json")
	if err != nil {
		panic(err)
	}
	config, err := google.ConfigFromJSON(jsonKey, googleoauth2.UserinfoEmailScope)
	if err != nil {
		panic(err)
	}
	tokenJson, err := os.ReadFile("token.json")
	if err != nil {
		panic(err)
	}
	var token oauth2.Token
	err = json.Unmarshal(tokenJson, &token)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	client := config.Client(ctx, &token)

	srv, err := googleoauth2.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		panic(err)
	}
	mailInfo, err := srv.Userinfo.Get().Context(ctx).Do()
	if err != nil {
		panic(err)
	}
	fmt.Printf("login with mail: %s\n", mailInfo.Email)
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func main() {
	ctx := context.Background()

	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	// client := getClient(config)
	token, err := tokenFromFile("token1.json")
	if err != nil {
		log.Fatalf("Unable to read token: %v", err)

	}

	client := config.Client(ctx, token)

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}
	// fmt.Println(srv)

	about, err := srv.About.Get().Fields("*").Do()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println("storage quota = ", about.StorageQuota.Limit)
	fmt.Println("storage usage = ", about.StorageQuota.Usage)
	percentUsage := float32(about.StorageQuota.Usage) / float32(about.StorageQuota.Limit)
	percentUsage = percentUsage * 100
	fmt.Println("percent usage = ", percentUsage)

	allFiles, err := srv.Files.List().Fields("nextPageToken, files(id, name)").Do()
	// r, err := srv.Files.List().Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	fmt.Println("Files:")
	if len(allFiles.Files) == 0 {
		fmt.Println("No files found. Create new folder")

	} else {
		for _, i := range allFiles.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}

	parentFolderName := "TestUploadFromDrive"
	q := fmt.Sprintf("name='%s'", parentFolderName)
	r, err := srv.Files.List().Fields("nextPageToken, files(id, name)").Q(q).Do()
	// r, err := srv.Files.List().Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found. Create new folder")
		folder, err := srv.Files.Create(&drive.File{Name: parentFolderName, MimeType: "application/vnd.google-apps.folder"}).Do()
		if err != nil {
			log.Fatal("cannot create folder")
		}
		fmt.Println("create folder: ", folder.Id)

	} else {
		for _, i := range r.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}

	parentFolderId := r.Files[0].Id

	nowTime := time.Now()
	fileName := fmt.Sprintf("%v.txt", nowTime.Unix())
	content := strings.NewReader(fmt.Sprintf("content of this file eiei. created at %v", nowTime.Format(time.DateTime)))

	f, err := srv.Files.Create(&drive.File{Name: fileName, Parents: []string{parentFolderId}}).Media(content).Do()
	if err != nil {
		log.Fatalf("unable to upload: %v", err)
	}
	fmt.Println("upload file: ", f.Name)

}

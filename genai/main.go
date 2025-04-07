package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/auth/credentials"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/genai"
)

func main() {
	var cfg struct {
		APIKey          string `envconfig:"API_KEY" required:"true"`
		ModelID         string `envconfig:"MODEL_ID" required:"true"`
		ProjectID       string `envconfig:"PROJECT_ID" required:"true"`
		Location        string `envconfig:"LOCATION" required:"true"`
		CredentialsFile string `envconfig:"CREDENTIALS_FILE" required:"true"`
	}

	err := envconfig.Process("GENAI", &cfg)
	if err != nil {
		panic(err)
	}

	creds, err := credentials.DetectDefault(
		&credentials.DetectOptions{
			Scopes:          []string{"https://www.googleapis.com/auth/cloud-platform"},
			CredentialsFile: cfg.CredentialsFile,
		},
	)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	client, err := genai.NewClient(
		ctx,
		&genai.ClientConfig{
			Backend:     genai.BackendVertexAI,
			Location:    cfg.Location,
			Project:     cfg.ProjectID,
			Credentials: creds,
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(client.ClientConfig().Backend)
	fmt.Println(string(client.ClientConfig().Credentials.JSON()))

	p, err := client.Models.List(ctx, nil)
	if err != nil {
		panic(err)
	}
	for _, m := range p.Items {
		fmt.Println(m.Name)
	}

	fmt.Println("------------------------------------------")

	// model := "gemini-2.0-flash-lite-001"
	model := fmt.Sprintf(
		"projects/%s/locations/us-central1/endpoints/%s",
		cfg.ProjectID,
		cfg.ModelID,
	)
	// prompt := "Introduce yourself in 1-2 sentences."
	prompt := "Before you can use Windows to its fullest, you will need to activate it. Activation can be done automatically over the internet. Click the Activation icon in the System Tray to start the process. . If you selected not to update automatically, you should run Windows Update as soon as possible. This will ensure that you have the latest security and stability fixes. If you chose to automatically update, your computer will start downloading and installing updates as soon as it is connected to the internet. . Most of your hardware should be installed automatically, but you may have to get drivers for more specialized hardware, or download the latest versions from the manufacturers. You can see what needs drivers from the Device Manager. . While Microsoft provides a free antivirus solution called Microsoft Essentials, it is barebones and not a strong deterrent against viruses. Instead, install a third-party antivirus program that will help to protect your computer and information. You can find both free and paid antivirus software. Once you’ve got Windows updated and protected, you can start installing the programs you need. Keep in mind that not every program that you used in older versions of Windows will be compatible with Windows Vista.\n\nProvide a summary of the article in two or three sentences:\n\n"

	t, err := client.Models.CountTokens(
		ctx,
		model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("total token: %d", t.TotalTokens))

	resp, err := client.Models.GenerateContent(
		ctx,
		model,
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Text())

}

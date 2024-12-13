package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/sashabaranov/go-openai"
)

type EnvCfg struct {
	ApiKey string `envconfig:"OPENAI_APIKEY" required:"true"`
}

func main() {
	var envCfg EnvCfg
	err := envconfig.Process("MYAPP", &envCfg)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	fmt.Println(envCfg.ApiKey)

	client := openai.NewClient(envCfg.ApiKey)

	resp, err := client.CreateSpeech(ctx, openai.CreateSpeechRequest{
		Model: openai.SpeechModel(openai.TTSModel1),
		Voice: openai.VoiceAlloy,
		Input: "ว่าจะใด๋",
	})

	if err != nil {
		panic(err)
	}

	// _, err = io.Copy(outFile, resp.ReadCloser)
	// if err != nil {
	// 	panic(err)
	// }

	soundData, err := io.ReadAll(resp)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("new.mp3", soundData, 0644)
	if err != nil {
		panic(err)
	}

}

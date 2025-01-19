package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/emiago/diago"
	"github.com/emiago/sipgo"
	"github.com/emiago/sipgo/sip"
	"github.com/kelseyhightower/envconfig"
)

type EnvCfg struct {
	User       string `envconfig:"USER" required:"true"`
	Password   string `envconfig:"PASSWORD" required:"true"`
	ServerHost string `envconfig:"SERVER_HOST" required:"true"`
	ServerPort string `envconfig:"SERVER_PORT" required:"true"`
}

func main() {

	var envCfg EnvCfg
	err := envconfig.Process("GO_SIP_PHONE_REGISTER", &envCfg)
	if err != nil {
		panic(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var uri sip.Uri
	err = sip.ParseUri(
		fmt.Sprintf("sip:%v@%v:%v",
			envCfg.User,
			envCfg.ServerHost,
			envCfg.ServerPort,
		),
		&uri,
	)
	if err != nil {
		panic(err)
	}

	ua, err := sipgo.NewUA()
	if err != nil {
		panic(err)
	}

	tu := diago.NewDiago(
		ua,
		diago.WithTransport(diago.Transport{
			Transport: "tcp",
			BindHost:  "localhost",
			BindPort:  5060,
		}))

	// tu.Register(ctx, sip.Uri{
	// 	User: ,
	// })

}

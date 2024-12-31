package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/emiago/diago"
	"github.com/emiago/sipgo"
	"github.com/emiago/sipgo/sip"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	log.Logger = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.StampMicro,
	}).With().Timestamp().Logger().Level(0)

	sip.SIPDebug = true

	ua, err := sipgo.NewUA()
	if err != nil {
		panic(err)
	}
	defer ua.Close()

	var uri sip.Uri

	err = sip.ParseUri("sip:111@127.0.0.1:5090;transport=tcp", &uri)
	if err != nil {
		panic(err)
	}

	tu := diago.NewDiago(ua, diago.WithTransport(diago.Transport{
		Transport: "tcp",
		BindHost:  "localhost",
		BindPort:  5060,
	}))
	client, err := tu.Invite(ctx, uri, diago.InviteOptions{})
	if err != nil {
		panic(err)
	}
	_ = client

}
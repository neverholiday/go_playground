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

	tu := diago.NewDiago(ua, diago.WithTransport(diago.Transport{
		Transport: "tcp",
		BindHost:  "0.0.0.0",
		BindPort:  5060,
	}))

	tu.Serve(ctx, func(inDialog *diago.DialogServerSession) {
		log.Info().Str("id", inDialog.ID).Msg("New dialog request")

		inDialog.Progress()
		inDialog.Answer()
		inDialog.Bye(ctx)

		defer log.Info().Str("id", inDialog.ID).Msg("Dialog finished")
	})

}

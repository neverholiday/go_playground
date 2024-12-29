package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/emiago/diago"
	"github.com/emiago/sipgo"
	"github.com/rs/zerolog/log"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	ua, err := sipgo.NewUA()
	if err != nil {
		panic(err)
	}
	defer ua.Close()

	tu := diago.NewDiago(ua)
	tu.Serve(ctx, func(inDialog *diago.DialogServerSession) {
		log.Info().Str("id", inDialog.ID).Msg("New dialog request")
		defer log.Info().Str("id", inDialog.ID).Msg("Dialog finished")
	})

}

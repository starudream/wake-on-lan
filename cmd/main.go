package main

import (
	"flag"

	"github.com/starudream/go-lib/log"

	"github.com/starudream/wake-on-lan/wol"
)

type Params struct {
	Addr string
	Port string
	Mac  string
}

var params = &Params{}

func init() {
	log.Attach("app", "wake-on-lan")

	flag.StringVar(&params.Addr, "addr", "255.255.255.255", "broadcast ip")
	flag.StringVar(&params.Port, "port", "9", "")
	flag.StringVar(&params.Mac, "mac", "", "mac address")

	flag.Parse()
}

func main() {
	err := wol.Send(params.Addr, params.Port, params.Mac)
	if err != nil {
		log.Error().Msgf("send failed: %v", err)
	} else {
		log.Info().Msg("send success")
	}
}

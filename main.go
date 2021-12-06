package main

import (
	"flag"
	"os"

	"github.com/go-sdk/lib/app"
	"github.com/go-sdk/lib/log"
)

type Config struct {
	Addr string
	Port string
	Mac  string

	Packet MagicPacket

	Help bool
}

var config = &Config{}

func init() {
	flag.StringVar(&config.Addr, "addr", "255.255.255.255", "broadcast ip")
	flag.StringVar(&config.Port, "port", "9", "")
	flag.StringVar(&config.Mac, "mac", "", "mac address")

	flag.BoolVar(&config.Help, "help", false, "instructions for use")
	flag.Parse()

	if config.Help {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {
	a := app.New("wake-on-lan")
	defer a.Recover()

	a.Add(Wake)

	err := a.Once()
	if err != nil {
		log.Fatal(err)
	}
}

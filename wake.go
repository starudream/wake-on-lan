package main

import (
	"fmt"
	"net"

	"github.com/go-sdk/lib/log"
)

const (
	MacAddressLength         = 6
	MagicPacketHeaderLength  = 6
	MagicPacketPayloadSize   = 16
	MagicPacketPayloadLength = MagicPacketPayloadSize * MacAddressLength
	MagicPacketLength        = MagicPacketHeaderLength + MagicPacketPayloadLength
)

type MagicPacket [MagicPacketLength]byte

func Wake() error {
	if err := Init(); err != nil {
		return err
	}

	conn, err := net.Dial("udp", fmt.Sprintf("%s:%s", config.Addr, config.Port))
	if err != nil {
		return err
	}
	defer conn.Close()

	n, err := conn.Write(config.Packet[:])
	if err != nil {
		return err
	}
	if n == MagicPacketLength {
		log.Info("send success")
	} else {
		log.Error("send fail, length: %d", n)
	}

	return nil
}

func Init() error {
	if config.Mac == "" {
		return fmt.Errorf("mac address is empty")
	}

	hw, err := net.ParseMAC(config.Mac)
	if err != nil {
		return err
	}

	if len(hw) != MacAddressLength {
		return fmt.Errorf("%s is not a IEEE 802 MAC-48 address", config.Mac)
	}

	for i := 0; i < MagicPacketHeaderLength; i++ {
		config.Packet[i] = 0xff
	}

	for i := 0; i < MagicPacketPayloadSize; i++ {
		for j := 0; j < MacAddressLength; j++ {
			config.Packet[MagicPacketHeaderLength+MacAddressLength*i+j] = hw[j]
		}
	}

	return nil
}

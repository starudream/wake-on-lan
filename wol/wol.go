package wol

import (
	"fmt"
	"net"
)

const (
	MacAddressLength         = 6
	MagicPacketHeaderLength  = 6
	MagicPacketPayloadSize   = 16
	MagicPacketPayloadLength = MagicPacketPayloadSize * MacAddressLength
	MagicPacketLength        = MagicPacketHeaderLength + MagicPacketPayloadLength
)

type MagicPacket [MagicPacketLength]byte

func GenMagicPacket(mac string) (MagicPacket, error) {
	if mac == "" {
		return MagicPacket{}, fmt.Errorf("mac address is empty")
	}

	hw, err := net.ParseMAC(mac)
	if err != nil {
		return MagicPacket{}, err
	}

	if len(hw) != MacAddressLength {
		return MagicPacket{}, fmt.Errorf("mac address has invalid length")
	}

	packet := MagicPacket{}

	for i := 0; i < MagicPacketHeaderLength; i++ {
		packet[i] = 0xff
	}

	for i := 0; i < MagicPacketPayloadSize; i++ {
		for j := 0; j < MacAddressLength; j++ {
			packet[MagicPacketHeaderLength+MacAddressLength*i+j] = hw[j]
		}
	}

	return packet, nil
}

func Send(addr, port, mac string) error {
	packet, err := GenMagicPacket(mac)
	if err != nil {
		return err
	}

	conn, err := net.Dial("udp", fmt.Sprintf("%s:%s", addr, port))
	if err != nil {
		return err
	}
	defer conn.Close()

	n, err := conn.Write(packet[:])
	if err != nil {
		return err
	}

	if n != MagicPacketLength {
		return fmt.Errorf("wrote %d bytes instead of %d", n, MagicPacketLength)
	}

	return nil
}

package ftp

import (
	"github.com/chrjoh/pcapparse/util"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// Parse go through the given pcap file and selects ftp logins and save in the given output file
func Parse(inputFunc string, outputFunc string) {
	if handle, err := pcap.OpenOffline(inputFunc); err != nil {
		panic(err)
	} else {
		defer handle.Close()
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		ftpResult := NewFtpHandler()
		for packet := range packetSource.Packets() {
			if util.IsTcpPacket(packet) {
				ftpResult.HandlePacket(packet)
			}
		}
		ftpResult.Dump(outputFunc)
	}
}

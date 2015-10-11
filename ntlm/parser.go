package ntlm

import (
	"github.com/chrjoh/pcapparse/util"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// Parse the given pcap file for ntlm challenge/responsees and stores them on the given output file
func Parse(inputFunc string, outputFunc string) {
	if handle, err := pcap.OpenOffline(inputFunc); err != nil {
		panic(err)
	} else {
		ntlmResult := NewNtlmHandler()
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			if util.IsTcpPacket(packet) {
				ntlmResult.handlePacket(packet)
			}
		}
		ntlmResult.dump(outputFunc)
	}
}
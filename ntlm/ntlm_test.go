package ntlm

import (
	"strings"
	"testing"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var challengePacket []byte = []byte{
	0x00, 0x0c, 0x29, 0xd8, 0x37, 0x94, 0x00, 0x50,
	0x56, 0xc0, 0x00, 0x01, 0x08, 0x00, 0x45, 0x00,
	0x02, 0x11, 0x65, 0x65, 0x40, 0x00, 0x40, 0x06,
	0xbc, 0x24, 0xc0, 0xa8, 0x4b, 0x01, 0xc0, 0xa8,
	0x4b, 0x0b, 0x00, 0x50, 0xc0, 0x15, 0xcf, 0x0d,
	0x6f, 0x48, 0xbf, 0xd7, 0xcb, 0xc6, 0x50, 0x18,
	0x00, 0xed, 0x28, 0x9f, 0x00, 0x00, 0x48, 0x54,
	0x54, 0x50, 0x2f, 0x31, 0x2e, 0x31, 0x20, 0x34,
	0x30, 0x31, 0x20, 0x55, 0x6e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x0d,
	0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a,
	0x20, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x6f,
	0x66, 0x74, 0x2d, 0x49, 0x49, 0x53, 0x2f, 0x36,
	0x2e, 0x30, 0x0d, 0x0a, 0x44, 0x61, 0x74, 0x65,
	0x3a, 0x20, 0x57, 0x65, 0x64, 0x2c, 0x20, 0x31,
	0x32, 0x20, 0x53, 0x65, 0x70, 0x20, 0x32, 0x30,
	0x31, 0x32, 0x20, 0x31, 0x33, 0x3a, 0x30, 0x36,
	0x3a, 0x35, 0x35, 0x20, 0x47, 0x4d, 0x54, 0x0d,
	0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2d, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x20, 0x74,
	0x65, 0x78, 0x74, 0x2f, 0x68, 0x74, 0x6d, 0x6c,
	0x0d, 0x0a, 0x57, 0x57, 0x57, 0x2d, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x3a, 0x20, 0x4e, 0x54, 0x4c, 0x4d,
	0x20, 0x54, 0x6c, 0x52, 0x4d, 0x54, 0x56, 0x4e,
	0x54, 0x55, 0x41, 0x41, 0x43, 0x41, 0x41, 0x41,
	0x41, 0x42, 0x67, 0x41, 0x47, 0x41, 0x44, 0x67,
	0x41, 0x41, 0x41, 0x41, 0x46, 0x41, 0x6f, 0x6d,
	0x69, 0x45, 0x53, 0x49, 0x7a, 0x52, 0x46, 0x56,
	0x6d, 0x64, 0x34, 0x67, 0x41, 0x41, 0x41, 0x41,
	0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x49, 0x41,
	0x41, 0x67, 0x41, 0x41, 0x2b, 0x41, 0x41, 0x41,
	0x41, 0x42, 0x51, 0x4c, 0x4f, 0x44, 0x67, 0x41,
	0x41, 0x41, 0x41, 0x39, 0x54, 0x41, 0x45, 0x30,
	0x41, 0x51, 0x67, 0x41, 0x43, 0x41, 0x41, 0x59,
	0x41, 0x55, 0x77, 0x42, 0x4e, 0x41, 0x45, 0x49,
	0x41, 0x41, 0x51, 0x41, 0x57, 0x41, 0x46, 0x4d,
	0x41, 0x54, 0x51, 0x42, 0x43, 0x41, 0x43, 0x30,
	0x41, 0x56, 0x41, 0x42, 0x50, 0x41, 0x45, 0x38,
	0x41, 0x54, 0x41, 0x42, 0x4c, 0x41, 0x45, 0x6b,
	0x41, 0x56, 0x41, 0x41, 0x45, 0x41, 0x42, 0x49,
	0x41, 0x63, 0x77, 0x42, 0x74, 0x41, 0x47, 0x49,
	0x41, 0x4c, 0x67, 0x42, 0x73, 0x41, 0x47, 0x38,
	0x41, 0x59, 0x77, 0x42, 0x68, 0x41, 0x47, 0x77,
	0x41, 0x41, 0x77, 0x41, 0x6f, 0x41, 0x48, 0x4d,
	0x41, 0x5a, 0x51, 0x42, 0x79, 0x41, 0x48, 0x59,
	0x41, 0x5a, 0x51, 0x42, 0x79, 0x41, 0x44, 0x49,
	0x41, 0x4d, 0x41, 0x41, 0x77, 0x41, 0x44, 0x4d,
	0x41, 0x4c, 0x67, 0x42, 0x7a, 0x41, 0x47, 0x30,
	0x41, 0x59, 0x67, 0x41, 0x75, 0x41, 0x47, 0x77,
	0x41, 0x62, 0x77, 0x42, 0x6a, 0x41, 0x47, 0x45,
	0x41, 0x62, 0x41, 0x41, 0x46, 0x41, 0x42, 0x49,
	0x41, 0x63, 0x77, 0x42, 0x74, 0x41, 0x47, 0x49,
	0x41, 0x4c, 0x67, 0x42, 0x73, 0x41, 0x47, 0x38,
	0x41, 0x59, 0x77, 0x42, 0x68, 0x41, 0x47, 0x77,
	0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x3d,
	0x3d, 0x0d, 0x0a, 0x58, 0x2d, 0x50, 0x6f, 0x77,
	0x65, 0x72, 0x65, 0x64, 0x2d, 0x42, 0x79, 0x3a,
	0x20, 0x41, 0x53, 0x50, 0x2e, 0x4e, 0x43, 0x30,
	0x43, 0x44, 0x37, 0x42, 0x37, 0x38, 0x30, 0x32,
	0x43, 0x37, 0x36, 0x37, 0x33, 0x36, 0x45, 0x39,
	0x42, 0x32, 0x36, 0x46, 0x42, 0x31, 0x39, 0x42,
	0x45, 0x42, 0x32, 0x44, 0x33, 0x36, 0x32, 0x39,
	0x30, 0x42, 0x39, 0x46, 0x46, 0x39, 0x41, 0x34,
	0x36, 0x45, 0x44, 0x44, 0x41, 0x35, 0x45, 0x54,
	0x0d, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2d, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x3a, 0x20, 0x30, 0x0d, 0x0a, 0x0d, 0x0a,
}

var responsePacket []byte = []byte{

	0x00, 0x50, 0x56, 0xc0, 0x00, 0x01, 0x00, 0x0c,
	0x29, 0xd8, 0x37, 0x94, 0x08, 0x00, 0x45, 0x00,
	0x03, 0x24, 0x02, 0x03, 0x40, 0x00, 0x80, 0x06,
	0xde, 0x73, 0xc0, 0xa8, 0x4b, 0x0b, 0xc0, 0xa8,
	0x4b, 0x01, 0xc0, 0x15, 0x00, 0x50, 0xbf, 0xd7,
	0xcb, 0xc6, 0xcf, 0x0d, 0x71, 0x31, 0x50, 0x18,
	0x00, 0xfe, 0xed, 0x64, 0x00, 0x00, 0x47, 0x45,
	0x54, 0x20, 0x2f, 0x77, 0x70, 0x61, 0x64, 0x2e,
	0x64, 0x61, 0x74, 0x20, 0x48, 0x54, 0x54, 0x50,
	0x2f, 0x31, 0x2e, 0x31, 0x0d, 0x0a, 0x41, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x3a, 0x20, 0x2a, 0x2f,
	0x2a, 0x0d, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x2d,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x3a, 0x20, 0x4d,
	0x6f, 0x7a, 0x69, 0x6c, 0x6c, 0x61, 0x2f, 0x35,
	0x2e, 0x30, 0x20, 0x28, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x3b, 0x20,
	0x49, 0x45, 0x20, 0x31, 0x31, 0x2e, 0x30, 0x3b,
	0x20, 0x57, 0x69, 0x6e, 0x36, 0x34, 0x3b, 0x20,
	0x54, 0x72, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2f,
	0x37, 0x2e, 0x30, 0x29, 0x0d, 0x0a, 0x48, 0x6f,
	0x73, 0x74, 0x3a, 0x20, 0x77, 0x70, 0x61, 0x64,
	0x0d, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x4b, 0x65,
	0x65, 0x70, 0x2d, 0x41, 0x6c, 0x69, 0x76, 0x65,
	0x0d, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
	0x20, 0x4e, 0x54, 0x4c, 0x4d, 0x20, 0x54, 0x6c,
	0x52, 0x4d, 0x54, 0x56, 0x4e, 0x54, 0x55, 0x41,
	0x41, 0x44, 0x41, 0x41, 0x41, 0x41, 0x47, 0x41,
	0x41, 0x59, 0x41, 0x4a, 0x6f, 0x41, 0x41, 0x41,
	0x41, 0x4f, 0x41, 0x51, 0x34, 0x42, 0x73, 0x67,
	0x41, 0x41, 0x41, 0x42, 0x6f, 0x41, 0x47, 0x67,
	0x42, 0x59, 0x41, 0x41, 0x41, 0x41, 0x44, 0x67,
	0x41, 0x4f, 0x41, 0x48, 0x49, 0x41, 0x41, 0x41,
	0x41, 0x61, 0x41, 0x42, 0x6f, 0x41, 0x67, 0x41,
	0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
	0x44, 0x41, 0x41, 0x51, 0x41, 0x41, 0x42, 0x51,
	0x4b, 0x49, 0x6f, 0x67, 0x59, 0x42, 0x73, 0x52,
	0x30, 0x41, 0x41, 0x41, 0x41, 0x50, 0x62, 0x41,
	0x46, 0x63, 0x44, 0x35, 0x56, 0x61, 0x6a, 0x41,
	0x73, 0x43, 0x6f, 0x75, 0x7a, 0x7a, 0x76, 0x61,
	0x53, 0x56, 0x2b, 0x56, 0x55, 0x41, 0x54, 0x41,
	0x42, 0x42, 0x41, 0x45, 0x6b, 0x41, 0x55, 0x67,
	0x41, 0x74, 0x41, 0x46, 0x4d, 0x41, 0x56, 0x51,
	0x42, 0x53, 0x41, 0x45, 0x59, 0x41, 0x4c, 0x51,
	0x42, 0x51, 0x41, 0x45, 0x4d, 0x41, 0x5a, 0x41,
	0x42, 0x79, 0x41, 0x46, 0x38, 0x41, 0x5a, 0x51,
	0x42, 0x32, 0x41, 0x47, 0x6b, 0x41, 0x62, 0x41,
	0x42, 0x56, 0x41, 0x45, 0x77, 0x41, 0x51, 0x51,
	0x42, 0x4a, 0x41, 0x46, 0x49, 0x41, 0x4c, 0x51,
	0x42, 0x54, 0x41, 0x46, 0x55, 0x41, 0x55, 0x67,
	0x42, 0x47, 0x41, 0x43, 0x30, 0x41, 0x55, 0x41,
	0x42, 0x44, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
	0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
	0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
	0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
	0x41, 0x41, 0x41, 0x4e, 0x2f, 0x4f, 0x68, 0x4f,
	0x6a, 0x30, 0x64, 0x77, 0x2b, 0x4b, 0x4e, 0x53,
	0x35, 0x42, 0x4f, 0x74, 0x6b, 0x35, 0x49, 0x6a,
	0x41, 0x42, 0x41, 0x51, 0x41, 0x41, 0x41, 0x41,
	0x41, 0x41, 0x41, 0x4f, 0x2f, 0x44, 0x30, 0x69,
	0x32, 0x6b, 0x74, 0x4e, 0x41, 0x42, 0x4a, 0x4c,
	0x58, 0x59, 0x31, 0x75, 0x65, 0x59, 0x6f, 0x4f,
	0x77, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x67,
	0x41, 0x47, 0x41, 0x46, 0x4d, 0x41, 0x54, 0x51,
	0x42, 0x43, 0x41, 0x41, 0x45, 0x41, 0x46, 0x67,
	0x42, 0x54, 0x41, 0x45, 0x30, 0x41, 0x51, 0x67,
	0x41, 0x74, 0x41, 0x46, 0x51, 0x41, 0x54, 0x77,
	0x42, 0x50, 0x41, 0x45, 0x77, 0x41, 0x53, 0x77,
	0x42, 0x4a, 0x41, 0x46, 0x51, 0x41, 0x42, 0x41,
	0x41, 0x53, 0x41, 0x48, 0x4d, 0x41, 0x62, 0x51,
	0x42, 0x69, 0x41, 0x43, 0x34, 0x41, 0x62, 0x41,
	0x42, 0x76, 0x41, 0x47, 0x4d, 0x41, 0x59, 0x51,
	0x42, 0x73, 0x41, 0x41, 0x4d, 0x41, 0x4b, 0x41,
	0x42, 0x7a, 0x41, 0x47, 0x55, 0x41, 0x63, 0x67,
	0x42, 0x32, 0x41, 0x47, 0x55, 0x41, 0x63, 0x67,
	0x41, 0x79, 0x41, 0x44, 0x41, 0x41, 0x4d, 0x41,
	0x41, 0x7a, 0x41, 0x43, 0x34, 0x41, 0x63, 0x77,
	0x42, 0x74, 0x41, 0x47, 0x49, 0x41, 0x4c, 0x67,
	0x42, 0x73, 0x41, 0x47, 0x38, 0x41, 0x59, 0x77,
	0x42, 0x68, 0x41, 0x47, 0x77, 0x41, 0x42, 0x51,
	0x41, 0x53, 0x41, 0x48, 0x4d, 0x41, 0x62, 0x51,
	0x42, 0x69, 0x41, 0x43, 0x34, 0x41, 0x62, 0x41,
	0x42, 0x76, 0x41, 0x47, 0x4d, 0x41, 0x59, 0x51,
	0x42, 0x73, 0x41, 0x41, 0x67, 0x41, 0x4d, 0x41,
	0x41, 0x77, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
	0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
	0x41, 0x41, 0x4d, 0x41, 0x41, 0x41, 0x37, 0x54,
	0x6d, 0x51, 0x6e, 0x33, 0x49, 0x75, 0x38, 0x72,
	0x6b, 0x57, 0x57, 0x31, 0x42, 0x55, 0x43, 0x79,
	0x71, 0x61, 0x69, 0x56, 0x62, 0x70, 0x6c, 0x32,
	0x44, 0x37, 0x4c, 0x53, 0x61, 0x77, 0x6a, 0x33,
	0x6d, 0x62, 0x6c, 0x6b, 0x39, 0x42, 0x4c, 0x37,
	0x41, 0x4b, 0x41, 0x42, 0x41, 0x41, 0x41, 0x41,
	0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
	0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
	0x41, 0x41, 0x41, 0x41, 0x6b, 0x41, 0x45, 0x67,
	0x42, 0x49, 0x41, 0x46, 0x51, 0x41, 0x56, 0x41,
	0x42, 0x51, 0x41, 0x43, 0x38, 0x41, 0x64, 0x77,
	0x42, 0x77, 0x41, 0x47, 0x45, 0x41, 0x5a, 0x41,
	0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
	0x41, 0x41, 0x41, 0x41, 0x3d, 0x3d, 0x0d, 0x0a,
	0x0d, 0x0a,
}
var challengeResponseLc = "dr_evil::ULAIR-SURF-PC:1122334455667788:dfce84e8f4770f8a352e413ad9392230:0101000000000000efc3d22da4b4d00124b5d8d6e798a0ec000000000200060053004d0042000100160053004d0042002d0054004f004f004c004b00490054000400120073006d0062002e006c006f00630061006c000300280073006500720076006500720032003000300033002e0073006d0062002e006c006f00630061006c000500120073006d0062002e006c006f00630061006c000800300030000000000000000000000000300000ed39909f722ef2b9165b50540b2a9a8956e99760fb2d26b08f799b964f412fb00a001000000000000000000000000000000000000900120048005400540050002f0077007000610064000000000000000000\n"

func TestChallengeResponse(t *testing.T) {
	h := NewNtlmHandler()
	h.HandlePacket(createIPv4TCPPacket(challengePacket))
	h.HandlePacket(createIPv4TCPPacket(responsePacket))

	for _, pair := range h.serverResponsePairs {
		data, _ := pair.getResponseData()
		lc := data.LcString()
		if strings.Compare(lc, challengeResponseLc) != 0 {
			t.Fatalf("wanted: %v got: %v", challengeResponseLc, lc)
		}
	}
}

func createIPv4TCPPacket(packet []byte) gopacket.Packet {
	return gopacket.NewPacket(packet, layers.LinkTypeEthernet, gopacket.DecodeOptions{Lazy: true, NoCopy: true})
}

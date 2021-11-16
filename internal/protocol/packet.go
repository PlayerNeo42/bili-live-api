package protocol

import (
	"encoding/json"
	"github.com/agentanderson/bili-live-api/internal/util"
)

type Packet struct {
	underlyingBytes []byte // underlyingBytes is the build result

	PacketLength    int // PacketLength will be calculated during build()
	HeaderLength    int // HeaderLength is currently a fixed number 16
	ProtocolVersion int
	Operation       int
	SequenceID      int
	Body            []byte
}

func NewPacket(protocolVersion int, operation int, body []byte) Packet {
	return Packet{
		ProtocolVersion: protocolVersion,
		Operation:       operation,
		Body:            body,
	}
}

func NewJSONPacket(operation int, body []byte) Packet {
	return NewPacket(JSON, operation, body)
}

func NewPacketFromBytes(data []byte) Packet {
	l := util.BytesToInt(data[0:4])
	if l != len(data) {
		log.WithField("data", data).Error("数据包丢失,不可能错误")
	}
	pv := util.BytesToInt(data[6:8])
	op := util.BytesToInt(data[8:12])
	body := data[16:l]
	packet := NewPacket(pv, op, body)

	return packet
}

func (p Packet) Parse() []Packet {
	switch p.ProtocolVersion {
	case Popularity:
		fallthrough
	case JSON:
		return []Packet{p}
	case Zlib:
		return Slice(zlibParser(p.Body))
	case Brotli:
		return Slice(brotliParser(p.Body))
	default:
		log.Error("未知protover")
	}
	return nil
}

func (p *Packet) Unmarshal(v interface{}) error {
	return json.Unmarshal(p.Body, v)
}

func (p *Packet) Build() []byte {
	if p.underlyingBytes != nil {
		return p.underlyingBytes
	}

	base := []byte{0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}

	p.PacketLength = len(p.Body) + 16

	util.WriteToBytes(p.PacketLength, base[0:4])
	// header length is a fixed number 16
	util.WriteToBytes(p.ProtocolVersion, base[6:8])
	util.WriteToBytes(p.Operation, base[8:12])
	// sequence id is a fixed number 1

	p.underlyingBytes = append(base, p.Body...)
	return p.underlyingBytes
}

// Decode aka NewPacketFromBytes
func Decode(data []byte) Packet {
	return NewPacketFromBytes(data)
}

// Encode aka Packet.Build
func Encode(packet Packet) []byte {
	return packet.Build()
}

func Slice(data []byte) []Packet {
	var packets []Packet
	total := len(data)
	cursor := 0
	for cursor < total {
		l := util.BytesToInt(data[cursor : cursor+4])
		packets = append(packets, Decode(data[cursor:cursor+l]))
		cursor += l
	}
	return packets
}

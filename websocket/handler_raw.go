package websocket

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/andybalholm/brotli"

	"github.com/spelling-dragon/bili-live-api/dto"
	"github.com/spelling-dragon/bili-live-api/log"
)

func parseAndHandle(p *dto.WSPayload, client *Client) error {
	payloads := rawDataParserMap[p.ProtocolVersion](p)
	for _, payload := range payloads {
		handler, ok := opCodeHandlerMap[payload.Operation]
		if !ok {
			return fmt.Errorf("未知的操作码:%d Body:%s", payload.Operation, string(p.Body))
		}
		handler(payload, client)
	}
	return nil
}

type rawDataParser func(*dto.WSPayload) []*dto.WSPayload

var rawDataParserMap = map[dto.ProtoVer]rawDataParser{
	dto.Brotli:     brotliDecompressor(),
	dto.Zlib:       zlibDecompressor(),
	dto.JSON:       jsonParser(),
	dto.Popularity: popularityParser(),
}

func jsonParser() rawDataParser {
	return func(p *dto.WSPayload) []*dto.WSPayload {
		return []*dto.WSPayload{p}
	}
}

func popularityParser() rawDataParser {
	return func(p *dto.WSPayload) []*dto.WSPayload {
		return []*dto.WSPayload{p}
	}
}

func zlibDecompressor() rawDataParser {
	return func(p *dto.WSPayload) []*dto.WSPayload {
		data := p.Body
		buf := bytes.NewReader(data)
		reader, err := zlib.NewReader(buf)
		if err != nil {
			log.Errorf("zlib解压错误: %v", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(reader)
		if err != nil {
			log.Errorf("zlib解压错误: %v", err)
			os.Exit(1)
		}
		return sliceAndParse(b)
	}
}

func brotliDecompressor() rawDataParser {
	return func(p *dto.WSPayload) []*dto.WSPayload {
		log.Debug("使用Brotli解压")
		buf := bytes.NewReader(p.Body)
		reader := brotli.NewReader(buf)
		b, err := io.ReadAll(reader)
		if err != nil {
			log.Errorf("brotli解压错误: %v", err)
			os.Exit(1)
		}
		return sliceAndParse(b)
	}
}

// Decode 将二进制数据转换为 *dto.WSPayload 类型
func Decode(data []byte) *dto.WSPayload {
	l := binary.BigEndian.Uint32(data[0:4])
	pv := binary.BigEndian.Uint16(data[6:8])
	op := binary.BigEndian.Uint32(data[8:12])
	body := data[16:l]
	return &dto.WSPayload{
		RawBytes:        data,
		PacketLength:    int(l),
		HeaderLength:    16,
		ProtocolVersion: dto.ProtoVer(pv),
		Operation:       dto.OPCode(op),
		Body:            body,
	}
}

// Encode 将 *dto.WSPayload 类型转换为二进制数据
func Encode(p *dto.WSPayload) []byte {
	if p.RawBytes != nil {
		return p.RawBytes
	}
	header := []byte{0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}

	p.PacketLength = len(p.Body) + 16

	binary.BigEndian.PutUint32(header[0:4], uint32(p.PacketLength))
	// header length is a fixed number 16
	binary.BigEndian.PutUint16(header[6:8], uint16(p.ProtocolVersion))
	binary.BigEndian.PutUint32(header[8:12], uint32(p.Operation))
	// sequence id is a fixed number 1

	p.RawBytes = append(header, p.Body...)
	return p.RawBytes
}

func sliceAndParse(data []byte) []*dto.WSPayload {
	total := len(data)
	var wsPayloads []*dto.WSPayload
	cursor := 0
	for cursor < total {
		l := int(binary.BigEndian.Uint32(data[cursor : cursor+4]))
		wsPayloads = append(wsPayloads, Decode(data[cursor:cursor+l]))
		cursor += l
	}
	return wsPayloads
}

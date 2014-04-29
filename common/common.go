package common

import (
	"os"
	"fmt"
    "hash/crc32"
    "bytes"
    "encoding/binary"
)

//define states
const STATE_UNKNOWN =  3
const STATE_CRITICAL = 2
const STATE_WARNING = 1
const STATE_OK = 0

//packet type
const QUERY_PACKET = 1
const RESPONSE_PACKET = 2

//max buffer size 
const MAX_PACKETBUFFER_LENGTH = 1024

//check the version definitions in common.h
const VERSION_TWO = 2

const HELLO_COMMAND = "version"

const PROGRAM_VERSION = "0.02"

type NrpePacket struct {
    Packet_version int16
    Packet_type int16
    Crc32_value uint32
    Result_code int16
    Command_buffer [MAX_PACKETBUFFER_LENGTH]byte
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func Docrc32(pkt NrpePacket) uint32 {
    buf := new(bytes.Buffer)
    if err := binary.Write(buf, binary.LittleEndian, &pkt); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return crc32.ChecksumIEEE(buf.Bytes())
}



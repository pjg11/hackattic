package main

import (
	"log"
	"bytes"
	"encoding/json"
	"encoding/base64"
	"encoding/binary"
	
	"github.com/pjg11/hackattic/connection"
)

type Problem struct {
	Bytes 		string 		`json:"bytes"`
}

type Solution struct {
	Signed 		int32 		`json:"int"`
	Unsigned 	uint32 		`json:"uint"`
	Short 		int32		`json:"short"` 	// gotcha: the length of the short is same as int32, but the last two bytes are zeros
	Float 		float32		`json:"float"`
	Double 		float64 	`json:"double"`
	BigDouble 	float64 	`json:"big_endian_double"`
}

func main() {
    resp, err := connection.Challenge("help_me_unpack")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Close()

	var prob Problem
	_ = json.NewDecoder(resp).Decode(&prob)
	data, _ := base64.StdEncoding.DecodeString(prob.Bytes)

	var sol Solution
	binary.Read(bytes.NewBuffer(data[:]), binary.LittleEndian, &sol)
	binary.Read(bytes.NewBuffer(data[len(data)-8:]), binary.BigEndian, &(sol.BigDouble)) // the BigDouble is to be read in Network Byte Order = Big Endian
	b, _ := json.Marshal(sol)

	connection.Solve("help_me_unpack", b)
}
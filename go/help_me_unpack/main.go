package main

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"

	"github.com/peeley/hackattic/lib"
)

type ChallengeInput struct {
	Bytes string `json:"bytes"`
}

type ChallengeOutput struct {
	Int             int32   `json:"int"`
	Uint            uint32  `json:"uint"`
	Short           int16   `json:"short"`
	Float           float64 `json:"float"`
	Double          float64 `json:"double"`
	BigEndianDouble float64 `json:"big_endian_double"`
}

const (
	Problem = "help_me_unpack"
	Token = "999222e3f1907dfd"
)

func main() {
	input := lib.GetChallengeInput(Problem, Token)

	var inputJson ChallengeInput
	json.Unmarshal(input, &inputJson)

	decodedBytes, err := base64.StdEncoding.DecodeString(inputJson.Bytes)

	if err != nil {
		panic(err)
	}

	int := int32(binary.LittleEndian.Uint32(decodedBytes[:4]))
	uint := binary.LittleEndian.Uint32(decodedBytes[4:8])
	short := int16(binary.LittleEndian.Uint16(decodedBytes[8:12]))
	float := float64(math.Float32frombits(binary.LittleEndian.Uint32(decodedBytes[12:16])))
	double := math.Float64frombits(binary.LittleEndian.Uint64(decodedBytes[16:24]))
	bigEndianDouble := math.Float64frombits(binary.BigEndian.Uint64(decodedBytes[24:]))

	output := ChallengeOutput{
		int,
		uint,
		short,
		float,
		double,
		bigEndianDouble,
	}

	outputString, err := json.Marshal(output)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(outputString))

	lib.SubmitChallengeSolution(Problem, Token, string(outputString))
}

package test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	tool "go-contract-decode-tool/decode"

	eth_abi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestDecoe(t *testing.T) {
	def := `[{"name" : "method", "type": "function", "outputs": [{ "type": "bool" }] }]`
	abi, err := eth_abi.JSON(strings.NewReader(def))
	if err != nil {
		panic(err)
	}
	data := "0000000000000000000000000000000000000000000000000000000000000001"
	encodeData, err := hex.DecodeString(data)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	fmt.Println(abi.Methods)
	decodeData, err := abi.Unpack("method", encodeData)
	if err != nil {
		panic(err)
	}
	fmt.Println(decodeData)
}

func TestHash(t *testing.T) {
	/*
		steps: 1. get 4 bytes to match method
			   2. get reset of data
	*/

	methodName := "getConversationList(address,uint256,uint256)"

	methodId := crypto.Keccak256Hash([]byte(methodName))

	id := methodId.Bytes()[:4]
	fmt.Println(hex.EncodeToString(id))
}

func TestGetOutput(t *testing.T) {
	def := `[{"name" : "method", "type": "function", "outputs": [{ "type": "bool" }] }]`
	abi, err := eth_abi.JSON(strings.NewReader(def))
	if err != nil {
		panic(err)
	}
	data := "0000000000000000000000000000000000000000000000000000000000000001"
	encodeData, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}

	out, err := tool.DecodeTxOutputDataByMethodName(abi, "method", encodeData)
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

func TestGetInput(t *testing.T) {
	def := `[{"inputs": [{"internalType": "address", "name": "user", "type": "address"}, {"internalType": "uint256", "name": "start", "type": "uint256"}, {"internalType": "uint256", "name": "count", "type": "uint256"}], "name": "getConversationList", "outputs": [], "stateMutability": "view", "type": "function"}]`
	abi, err := eth_abi.JSON(strings.NewReader(def))
	if err != nil {
		panic(err)
	}
	data := "000000000000000000000000f4e63c5efb0a04801d1208253198fe6a10148f5c0000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a"
	encodeData, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}
	out, err := tool.DecodeTxInputDataByMethodName(abi, "getConversationList", encodeData)
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

func TestDecodeDataByOrigingData(t *testing.T) {
	def := `[{"inputs": [{"internalType": "address", "name": "user", "type": "address"}, {"internalType": "uint256", "name": "start", "type": "uint256"}, {"internalType": "uint256", "name": "count", "type": "uint256"}], "name": "getConversationList", "outputs": [], "stateMutability": "view", "type": "function"}]`
	abi, err := eth_abi.JSON(strings.NewReader(def))
	if err != nil {
		panic(err)
	}
	data := "9557f0dc000000000000000000000000f4e63c5efb0a04801d1208253198fe6a10148f5c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a"
	encodeData, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}
	out, err := tool.DecodeTxInputDataByOriginData(abi, encodeData)
	if err != nil {
		panic(err)
	}
	res, _ := json.Marshal(out)
	fmt.Println(string(res))
}

func TestDecodeEventDataByOriginData(t *testing.T) {
	b, err := ioutil.ReadFile("./test.json")
	if err != nil {
		panic(err)
	}
	abi, err := eth_abi.JSON(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	data := "00000000000000000000000014dc79964da2c08b23698b3d3cc7ca32193d995500000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000006377411f00000000000000000000000023618e81e3f5cdf7f54c3d65f7fbc0abf5b21e8f000000000000000000000000000000000000000000000000000000000000004073656e642066726f6d207573657245696768742d3078313464433739393634646132433038623233363938423344336363374361333231393364393935352d32"
	encodeData, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}
	out, err := tool.DecodeTxInputDataByMethodName(abi, "MessageSend", encodeData)
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

func TestCount(t *testing.T) {
	var x = 1
	var i = 1
	for ; i <= 128; i += i {
		x += x
	}
	fmt.Print(x)
}

func TestSum(t *testing.T) {
	var words = []string{"cat", "slient", "listen", "kitten", "salient"}
	var length = len(words)
	var i = 0

	var hasCheck = make(map[string]bool)
	var count = 0
	for ; i < length; i++ {
		if hasCheck[words[i]] {
			continue
		}
		for j := i + 1; j < length; j++ {
			if check(words[i], words[j]) {
				count += 1
			}
		}
		hasCheck[words[i]] = true
	}
	fmt.Print(length - count)
}

func check(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if strings.Count(a, string([]byte(a)[i])) != strings.Count(b, string([]byte(a)[i])) {
			return false
		}
	}
	return true
}

func TestA(t *testing.T) {
	var a = "A"
	b, _ := strconv.Atoi(a)

	fmt.Println(b)
	fmt.Println([]byte(a))
}

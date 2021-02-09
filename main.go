package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	"aicdev.com/crack_xor_encryption/xor-cracker/core"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	message      = kingpin.Flag("message", "encrypted message to attack (base64)").Short('m').Required().String()
	maxKeyLength = kingpin.Flag("length", "max key length you want to test for").Short('l').Required().Int()
)

func main() {
	kingpin.Parse()

	fmt.Println("decode message from hex")
	m, err := hex.DecodeString(*message)
	core.Check(err)

	fmt.Println("compute reference metrics")
	ft := []string{
		"./data/alice_adventures_in_wonderland.txt",
		"./data/the_adeventure_of_sherlock_holmes.txt",
		"./data/pride_and_prejudice.txt",
	}

	rb, rl := core.GetReferenceTextMetrics(ft)

	keyMap := make(map[int]float32)
	ms := string(m)

	fmt.Println("process input message")
	for i := 2; i < *maxKeyLength; i++ {
		keyMap[i] = core.ProcessInputMessage(ms, i)
	}

	fmt.Println("computed hamming distances. key length should be one of these")
	computedKeys := core.GetTopValuesFromKeyMap(keyMap, 5)

	td := [][]string{}

	for i, v := range computedKeys {
		mg := core.GetMessageGroups(ms, v.KeyLength)
		crk := core.CrackKey(mg, rb, rl)
		ko, ks := core.GetStringKeys(crk)
		td = append(td, []string{strconv.Itoa(i + 1), strconv.Itoa(v.KeyLength), fmt.Sprintf("%f", v.Score), ko + ", " + ks})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rank", "Key-Length", "Score", "Possible keys"})
	for _, v := range td {
		table.Append(v)
	}
	table.Render()
}

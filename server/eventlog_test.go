package server

import (
	"encoding/hex"
	"fmt"
	"github.com/google/go-attestation/attest"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/ProsaicSatsuma/go-tpm-tools/client"
	"github.com/ProsaicSatsuma/go-tpm-tools/internal"
	tpmpb "github.com/ProsaicSatsuma/go-tpm-tools/proto"
	"github.com/google/go-tpm/tpm2"
)

func TestParseEventLog(t *testing.T) {
	absPath, err := filepath.Abs("./test/debian_10_binary_bios_measurements")
	if err != nil {
		t.Fatalf("failed to get abs path: %v", err)
	}
	evtLog, err := ioutil.ReadFile(absPath)
	if err != nil {
		t.Fatalf("failed to read event log: %v", err)
	}

	// Fetched from the tpm2-tools eventlog command.
	pcr0 := decodeHexOrFatal(t, "0f2d3a2a1adaa479aeeca8f5df76aadc41b862ea")
	pcr1 := decodeHexOrFatal(t, "b1676439cac1531683990fefe2218a43239d6fe8")
	pcr2 := decodeHexOrFatal(t, "b2a83b0ebf2f8374299a5b2bdfc31ea955ad7236")
	pcr3 := decodeHexOrFatal(t, "b2a83b0ebf2f8374299a5b2bdfc31ea955ad7236")
	pcr4 := decodeHexOrFatal(t, "1eb30816474a3f144e99b24e4ad480b2e51fd9e1")
	pcr5 := decodeHexOrFatal(t, "019079179dbc0eb5992c500dcf8a095910ac590d")
	pcr6 := decodeHexOrFatal(t, "b2a83b0ebf2f8374299a5b2bdfc31ea955ad7236")
	pcr7 := decodeHexOrFatal(t, "9e6c57e850f371c2a7fe02bca552149363952318")

	pcrs := tpmpb.Pcrs{
		Hash: tpmpb.HashAlgo_SHA1,
		Pcrs: map[uint32][]byte{
			0: pcr0,
			1: pcr1,
			2: pcr2,
			3: pcr3,
			4: pcr4,
			5: pcr5,
			6: pcr6,
			7: pcr7,
		},
	}

	var events []attest.Event
	events, err = ParseAndVerifyEventLog(evtLog, &pcrs)
	if err != nil {
		t.Errorf("failed to parse and verify log: %v", err)
	}

	fmt.Printf("event log size: %v\n\n", len(events))

	for i := 0; i < len(events); i++ {
		fmt.Printf("Event type: %v\n", events[i].Type)
		fmt.Printf("Event digest: %v\n", hex.EncodeToString(events[i].Digest))
		fmt.Printf("Applies to PCR %v\n", events[i].Index)
		fmt.Printf("%v\n", events[i].Data)
		fmt.Printf("\n")
	}
}

func TestSystemParseEventLog(t *testing.T) {
	rwc := internal.GetTPM(t)
	defer client.CheckedClose(t, rwc)
	evtLog := internal.GetEventLog(t, rwc)

	sel := tpm2.PCRSelection{Hash: tpm2.AlgSHA1, PCRs: []int{0, 4, 5, 7}}
	pcrs, err := client.ReadPCRs(rwc, sel)
	if err != nil {
		t.Errorf("failed to read PCRs: %v", err)
	}

	_, err = ParseAndVerifyEventLog(evtLog, pcrs)
	if err != nil {
		t.Errorf("failed to parse and verify log: %v", err)
	}
}

func decodeHexOrFatal(tb testing.TB, hexStr string) []byte {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		tb.Fatalf("failed to decode hex string (%v): %v", hexStr, err)
	}
	return bytes
}

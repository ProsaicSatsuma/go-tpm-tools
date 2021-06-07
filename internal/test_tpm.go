package internal

import (
	"io"
	"io/ioutil"
	"path/filepath"
	"sync"
	"testing"

	"github.com/google/go-attestation/attest"
	"github.com/ProsaicSatsuma/go-tpm-tools/simulator"
	"github.com/google/go-tpm/tpm2"
	"github.com/google/go-tpm/tpmutil"
)

// Only open the TPM device once. Reopening the device causes issues on Linux.
var (
	tpm  io.ReadWriteCloser
	lock sync.Mutex
)

const (
	testEventLogAlg = tpm2.AlgSHA1
)

type noClose struct {
	io.ReadWriter
}

func (n noClose) Close() error {
	return nil
}

// GetTPM is a cross-platform testing helper function that retrives the
// appropriate TPM device from the flags passed into "go test".
func GetTPM(tb testing.TB) io.ReadWriteCloser {
	tb.Helper()
	if useRealTPM() {
		lock.Lock()
		defer lock.Unlock()
		if tpm == nil {
			var err error
			if tpm, err = getRealTPM(); err != nil {
				tb.Fatalf("Failed to open TPM: %v", err)
			}
		}
		return noClose{tpm}
	}

	simulator, err := simulator.Get()
	if err != nil {
		tb.Fatalf("Simulator initialization failed: %v", err)
	}
	return simulator
}

// GetEventLog is a testing helper function that gets the TCG event log
// on supported systems, if using a real TPM, or a test event log, if not.
//
// Note that GetEventLog may have side effects.
// If a test requests a test event log, GetEventLog extends
// the test event log's events into the simulator.
func GetEventLog(tb testing.TB, rw io.ReadWriter) []byte {
	if useRealTPM() {
		eventLog, err := getRealEventLog()
		if err != nil {
			tb.Skipf("Failed to get system event log: %v", err)
		}
		return eventLog
	}

	absPath, err := filepath.Abs("../server/test/debian_10_binary_bios_measurements")
	if err != nil {
		tb.Fatalf("Failed to get abs path: %v", err)
	}
	eventLog, err := ioutil.ReadFile(absPath)
	if err != nil {
		tb.Fatalf("Failed to get test event log: %v", err)
	}

	// Extend event log events on simulator TPM.
	simulateEventLogEvents(tb, rw, eventLog)
	return eventLog
}

// simulateEventLogEvents simulates the events in the the test event log
// "server/test/debian_10_binary_bios_measurements" by parsing the log
// and manually extending the TPM.
func simulateEventLogEvents(tb testing.TB, rw io.ReadWriter, eventLog []byte) {
	attestEventLog, err := attest.ParseEventLog(eventLog)
	if err != nil {
		tb.Fatalf("Failed to parse test event log: %v", err)
	}

	events := attestEventLog.Events(attest.HashSHA1)
	for _, event := range events {
		extendOnePcr(tb, rw, event.Index, event.Digest)
	}
}

func extendOnePcr(tb testing.TB, rw io.ReadWriter, pcr int, hash []byte) {
	err := tpm2.PCRExtend(rw, tpmutil.Handle(pcr), testEventLogAlg, hash, "")
	if err != nil {
		tb.Fatalf("PCRExtend failed: %v", err)
	}
}

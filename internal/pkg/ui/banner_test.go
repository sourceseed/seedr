package ui

import (
	"testing"
	"bytes"
	"log"
	"os"
	"github.com/stretchr/testify/assert"
)

// Make sure the banner does not have tabs
func TestPrintBanner(t *testing.T) {
	output := captureOutput(func() {
		PrintBanner()
	})

	assert.NotContains(t, output, "\t")
}

func captureOutput(f func()) string {
    var buf bytes.Buffer
    log.SetOutput(&buf)
    f()
    log.SetOutput(os.Stderr)
    return buf.String()
}
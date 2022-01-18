package vtasample_test

import (
	"testing"

	"github.com/knsh14/vtasample"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, vtasample.Analyzer, "a", "b")
}

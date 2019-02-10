package balance

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	l "github.com/LaMonF/FDJ_SLACK/log"
)

func TestBalance_getBalanceValue(t *testing.T) {
	// Init
	content := []byte("666")
	dir, err := ioutil.TempDir("", "fdjSlackTEST")
	if err != nil {
		l.Error(err)
	}

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "test.*.fdjSlack")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		l.Error(err)
	}

	result := getBalanceValue(tmpfn)
	if result != 666 {
		t.Fail()
	}
}

func TestBalance_SetBalanceValue(t *testing.T) {
	// Init
	dir, err := ioutil.TempDir("", "fdjSlackTEST")
	if err != nil {
		l.Error(err)
	}
	defer os.RemoveAll(dir) // clean up
	tmpfn := filepath.Join(dir, "test.*.fdjSlack")

	const EXPECTED_VALUE = 999
	balance := NewBalance(tmpfn)
	balance.SetBalanceValue(EXPECTED_VALUE)
	if balance.Value != EXPECTED_VALUE {
		t.Fail()
	}

	dat, err := ioutil.ReadFile(tmpfn)
	formattedString := strings.Replace(string(dat), "\n", "", -1)
	value, _ := strconv.ParseFloat(formattedString, 64)
	if value != EXPECTED_VALUE {
		t.Fail()
	}

}

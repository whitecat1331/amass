package amass

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strings"
	"testing"
)

type testValues struct {
	Domain  string
	Domains []string
	Output  string
}

func NewTestingValues() testValues {
	domains := strings.Split(os.Getenv("DOMAINS"), ",")
	return testValues{
		Domain:  domains[0],
		Domains: domains,
		Output:  os.Getenv("OUTPUT"),
	}

}

var tv = NewTestingValues()

// var cliArgs = []string{"-silent, "-nocolor", "-d", tv.Domain}
var cliArgs = []string{"-nocolor", "-d", os.Getenv("DOMAINS")}

func TestRunEnumCommandPrivate(t *testing.T) {
	enumChan := make(chan string)
	go runEnumCommand(enumChan, cliArgs)
	for domain := range enumChan {
		t.Log(domain)
	}
}

func TestRunEnumCommandPublic(t *testing.T) {
	result := RunEnumCommand(cliArgs)
	t.Log(result)
}

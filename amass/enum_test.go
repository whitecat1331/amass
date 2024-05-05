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

// []string{"-silent", "-brute",
// "-nocolor", "-o", output, "-d", domain}
func TestAmassEnum(t *testing.T) {
	// args := []string{"-silent", "-brute", "-nocolor"}
	args := []string{"-brute", "-silent", "-nocolor", "-v",
		"-o", "amass_enum.log"}
	t.Log(args)
	t.Logf("Domains: %#v", tv.Domains)
	out := RunEnumCommand(tv.Domains, args...)
	for i := range out {
		t.Log(ParseFQDN(i))
	}
}

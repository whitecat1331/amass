package amass

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"testing"
)

func TestEnumDomainFinder(t *testing.T) {
	domains := EnumAllDomain(os.Getenv("AMASS_DOMAIN"), os.Getenv("AMASS_OUTPUT"))
	fmt.Println(domains)
}

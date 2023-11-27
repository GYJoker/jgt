package encrypt

import (
	"fmt"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	pwd := GeneratePassword("BankAdmin0!@", "lfakj23mkfca0skl;af")

	fmt.Println(pwd)
}

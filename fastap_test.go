package marshal

import (
	"fmt"
	"testing"

	"github.com/nazarifard/syncpool"
)

func TestFastape(t *testing.T) {
	//go:generate fastape string
	tap := NewFastap(stringTape{}, syncpool.StringPool)
	buff, _ := tap.Encode("aaaaaaaaaaaaaaaa")
	v, _, _ := tap.Decode(buff.Bytes())
	fmt.Println(*v)
	if *v != "aaaaaaaaaaaaaaaa" {
		t.Error("stringTape failed")
	}
	buff.Free()
}

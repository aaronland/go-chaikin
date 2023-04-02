package chaikin

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"testing"
)

func TestSmooth(t *testing.T) {

	// https://observablehq.com/@pamacha/chaikins-algorithm

	input := [][2]float64{
		[2]float64{185.6, 115.34581311988063},
		[2]float64{371.2, 426.3801223498234},
		[2]float64{556.8, 396.20750147765983},
		[2]float64{742.4, 447.7259263695299},
	}

	close := true
	output := Smooth(input, 6, close)

	enc, err := json.Marshal(output)

	if err != nil {
		t.Fatalf("Failed to marshal data, %v", err)
	}

	hash := sha256.Sum256(enc)
	str_hash := fmt.Sprintf("%x", hash)

	expected_hash := "62631f005a15212e5ff7ffe34f9a2dd9af5f0866b37229d6c05a07fb35d77d99"

	if str_hash != expected_hash {
		t.Fatalf("Unexpected hash for output, %s", str_hash)
	}

}

package credential

import (
	"testing"

	"github.com/goccy/go-json"

	"github.com/gobuffalo/packr/v2"
	"github.com/stretchr/testify/assert"
)

// These test vectors are taken from the vc-data-model spec examples
// e.g. https://www.w3.org/TR/vc-data-model/#example-a-simple-example-of-a-verifiable-credential
const (
	VCTestVector1 string = "vc-example-1.json"
	VCTestVector2 string = "vc-example-11.json"
	VCTestVector3 string = "vc-example-20.json"
	VCTestVector4 string = "vc-example-21.json"
	VPTestVector1 string = "vp-example-2.json"
	VPTestVector2 string = "vp-example-22.json"
)

var (
	box           = packr.New("VC & VP Test Vectors", "/test_vectors")
	vcTestVectors = []string{VCTestVector1, VCTestVector2, VCTestVector3, VCTestVector4}
	vpTestVectors = []string{VPTestVector1, VPTestVector2}
)

// Before running, you'll need to execute `mage packr`
func TestVCVectors(t *testing.T) {
	// round trip serialize and de-serialize from json to our object model
	for _, tv := range vcTestVectors {
		gotTestVector, err := getTestVector(tv)
		assert.NoError(t, err)

		var vc VerifiableCredential
		err = json.Unmarshal([]byte(gotTestVector), &vc)
		assert.NoError(t, err)

		assert.NoError(t, vc.IsValid())
		assert.False(t, vc.IsEmpty())

		vcBytes, err := json.Marshal(vc)
		assert.NoError(t, err)
		assert.JSONEqf(t, gotTestVector, string(vcBytes), "error message %s")
	}
}

// Before running, you'll need to execute `mage packr`
func TestVPVectors(t *testing.T) {
	// round trip serialize and de-serialize from json to our object model
	for _, tv := range vpTestVectors {
		gotTestVector, err := getTestVector(tv)
		assert.NoError(t, err)

		var vp VerifiablePresentation
		err = json.Unmarshal([]byte(gotTestVector), &vp)
		assert.NoError(t, err)

		assert.NoError(t, vp.IsValid())
		assert.False(t, vp.IsEmpty())

		vpBytes, err := json.Marshal(vp)
		assert.NoError(t, err)
		assert.JSONEqf(t, gotTestVector, string(vpBytes), "error message %s")
	}
}

func getTestVector(fileName string) (string, error) {
	return box.FindString(fileName)
}

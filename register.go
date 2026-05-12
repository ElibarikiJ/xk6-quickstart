package quickstart

import (
	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/ecdsa"
	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/privatekey"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/starkbank", new(StarkBank))
}

type StarkBank struct{}

// SignMessage takes a PEM private key and a message string, returning a Base64 signature
func (s *StarkBank) SignMessage(pemKey string, message string) string {
	// Import the key using Stark Bank's internal PEM parser
	privKey := privatekey.FromPem(pemKey)
	
	// Generate a deterministic signature (RFC 6979)
	signature := ecdsa.Sign(message, &privKey)
	
	// Return the Base64 DER encoded signature required by Stark Bank APIs
	return signature.ToBase64()
}

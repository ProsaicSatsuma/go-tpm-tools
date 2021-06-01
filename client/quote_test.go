package client

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rsa"
	"io"
	"testing"

	"github.com/google/go-tpm-tools/internal"
	"github.com/google/go-tpm/tpm2"
)

func TestQuote(t *testing.T) {
	rwc := internal.GetTPM(t)
	defer CheckedClose(t, rwc)

	tests := []struct {
		name   string
		getKey func(io.ReadWriter) (*Key, error)
	}{
		{"AK-ECC", AttestationKeyECC},
		{"AK-RSA", AttestationKeyRSA},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ak, err := test.getKey(rwc)
			if err != nil {
				t.Errorf("failed to generate AK: %v", err)
			}
			defer ak.Close()

			selpcr := tpm2.PCRSelection{
				Hash: tpm2.AlgSHA256,
				PCRs: []int{7},
			}
			quoted, err := ak.Quote(selpcr, []byte("test"))
			if err != nil {
				t.Errorf("failed to quote: %v", err)
			}
			sig, err := tpm2.DecodeSignature(bytes.NewBuffer(quoted.GetRawSig()))
			if err != nil {
				t.Errorf("signature decoding failed: %v", err)
			}

			switch pub := ak.PublicKey().(type) {
			case *ecdsa.PublicKey:
				hash, err := sig.ECC.HashAlg.Hash()
				if err != nil {
					t.Fatalf("not a valid hash type: %v", sig.ECC.HashAlg)
				}

				hashCon := hash.New()
				hashCon.Write(quoted.GetQuote())
				if !ecdsa.Verify(pub, hashCon.Sum(nil)[:], sig.ECC.R, sig.ECC.S) {
					t.Errorf("ECC signature verification failed")
				}
			case *rsa.PublicKey:
				hash, err := sig.RSA.HashAlg.Hash()
				if err != nil {
					t.Fatalf("not a valid hash type: %v", sig.RSA.HashAlg)
				}

				hashCon := hash.New()
				hashCon.Write(quoted.GetQuote())
				if err = rsa.VerifyPKCS1v15(pub, hash, hashCon.Sum(nil), []byte(sig.RSA.Signature)); err != nil {
					t.Errorf("RSA signature verification failed: %v", err)
				}
			}
		})
	}

}

func TestQuoteShouldFailWithNonSigningKey(t *testing.T) {
	rwc := internal.GetTPM(t)
	defer CheckedClose(t, rwc)

	srk, err := StorageRootKeyRSA(rwc)
	if err != nil {
		t.Errorf("failed to generate SRK: %v", err)
	}
	defer srk.Close()

	selpcr := tpm2.PCRSelection{
		Hash: tpm2.AlgSHA1,
		PCRs: []int{7},
	}
	_, err = srk.Quote(selpcr, []byte("test"))
	if err == nil {
		t.Errorf("Quote with a non-signing key should fail")
	}
}

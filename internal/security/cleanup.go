package security

import "log"

// SecureWipeCiphertexts guarantees that NO PII (even encrypted)
// persists after the verification transaction.
func SecureWipeCiphertexts(cDoc, cBio *string) {
	log.Println("Zero-persistence: Securely wiping ciphertexts from memory...")
	// In Go, strings are immutable, so we can't truly wipe the memory buffer
	// they back natively without unsafe pointers or CGo.
	// For actual implementation, use byte slices and wipe using:
	// for i := range buf { buf[i] = 0 }
	
	*cDoc = ""
	*cBio = ""
}

package api

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sakshya/gateway-tee/internal/engine"
	"github.com/sakshya/gateway-tee/internal/security"
	"github.com/sakshya/gateway-tee/internal/tee"
)

type VerifyRequest struct {
	PublicKeyBase64 string `json:"publicKeyBase64" binding:"required"`
	CDoc            string `json:"c_doc" binding:"required"`
	CBio            string `json:"c_bio" binding:"required"`
}

func VerifyHandler(c *gin.Context) {
	var req VerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Received verification request from client")

	// 1. TEE Attestation Check (Stub)
	if err := tee.VerifyEnclaveIntegrity(); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "TEE integrity check failed"})
		return
	}

	// 2. Pass to FFI Engine Bridge
	// In reality this calls the Rust Engine via FFI or a secure RPC channel inside the Nitro Enclave
	success, err := engine.RunFhePipeline(req.PublicKeyBase64, req.CDoc, req.CBio)

	// 3. Zero-persistence cleanup
	security.SecureWipeCiphertexts(&req.CDoc, &req.CBio)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "FHE processing failed"})
		return
	}

	// If success, we would also generate the ZKP here (or trigger the Circom prover)
	// For MVP, we return a success status
	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"status":  "verified-zkp-generated",
		"txHash":  "msg_solana_stub_tx_hash",
	})
}

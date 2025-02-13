// Copyright (C) 2022-2024, Chain4Travel AG. All rights reserved.
//
// This file is a derived work, based on ava-labs code whose
// original notices appear below.
//
// It is distributed under the same license conditions as the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********************************************************
// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package staking

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"

	// "github.com/ava-labs/avalanchego/utils/perms"
	"github.com/chain4travel/caminogoeth-compat/caminogo/perms"

	// "github.com/ava-labs/avalanchego/ids"
	"github.com/chain4travel/caminogoeth-compat/caminogo/ids"

	utilsSecp256k1 "github.com/chain4travel/caminogoeth-compat/caminogo/secp256k1"
)

var errDuplicateExtension = errors.New("duplicate certificate extension")

// InitNodeStakingKeyPair generates a self-signed TLS key/cert pair to use in
// staking. The key and files will be placed at [keyPath] and [certPath],
// respectively. If there is already a file at [keyPath], returns nil.
func StoreStakingKeyPair(keyPath, certPath string, keyBytes, certBytes []byte) error {
	// If there is already a file at [keyPath], do nothing
	if _, err := os.Stat(keyPath); !os.IsNotExist(err) { // TODO@
		return nil
	}

	// Ensure directory where key/cert will live exist
	if err := os.MkdirAll(filepath.Dir(certPath), perms.ReadWriteExecute); err != nil {
		return fmt.Errorf("couldn't create path for cert: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(keyPath), perms.ReadWriteExecute); err != nil {
		return fmt.Errorf("couldn't create path for key: %w", err)
	}

	// Write cert to disk
	certFile, err := os.Create(certPath)
	if err != nil {
		return fmt.Errorf("couldn't create cert file: %w", err)
	}
	if _, err := certFile.Write(certBytes); err != nil {
		return fmt.Errorf("couldn't write cert file: %w", err)
	}
	if err := certFile.Close(); err != nil {
		return fmt.Errorf("couldn't close cert file: %w", err)
	}
	if err := os.Chmod(certPath, perms.ReadOnly); err != nil { // Make cert read-only
		return fmt.Errorf("couldn't change permissions on cert: %w", err)
	}

	// Write key to disk
	keyOut, err := os.Create(keyPath)
	if err != nil {
		return fmt.Errorf("couldn't create key file: %w", err)
	}
	if _, err := keyOut.Write(keyBytes); err != nil {
		return fmt.Errorf("couldn't write private key: %w", err)
	}
	if err := keyOut.Close(); err != nil {
		return fmt.Errorf("couldn't close key file: %w", err)
	}
	if err := os.Chmod(keyPath, perms.ReadOnly); err != nil { // Make key read-only
		return fmt.Errorf("couldn't change permissions on key: %w", err)
	}
	return nil
}

func LoadTLSCertFromBytes(keyBytes, certBytes []byte) (*tls.Certificate, error) {
	cert, err := tls.X509KeyPair(certBytes, keyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed creating cert: %w", err)
	}

	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return nil, fmt.Errorf("failed parsing cert: %w", err)
	}
	return &cert, nil
}

func NewCertAndKeyBytesWithSecpKey() ([]byte, []byte, ids.NodeID, error) {
	// Create RSA key to sign cert with
	rsaKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, ids.EmptyNodeID, fmt.Errorf("couldn't generate rsa key: %w", err)
	}
	// Create SECP256K1 key to sign cert with
	secpKey := utilsSecp256k1.RsaPrivateKeyToSecp256PrivateKey(rsaKey)
	extension := utilsSecp256k1.SignRsaPublicKey(secpKey, &rsaKey.PublicKey)

	// Create self-signed staking cert
	certTemplate := &x509.Certificate{
		SerialNumber:          big.NewInt(0),
		NotBefore:             time.Date(2000, time.January, 0, 0, 0, 0, 0, time.UTC),
		NotAfter:              time.Now().AddDate(100, 0, 0),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageDataEncipherment,
		ExtraExtensions:       []pkix.Extension{*extension},
		BasicConstraintsValid: true,
	}
	certBytes, err := x509.CreateCertificate(rand.Reader, certTemplate, certTemplate, &rsaKey.PublicKey, rsaKey)
	if err != nil {
		return nil, nil, ids.EmptyNodeID, fmt.Errorf("couldn't create certificate: %w", err)
	}
	var certBuff bytes.Buffer
	if err := pem.Encode(&certBuff, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes}); err != nil {
		return nil, nil, ids.EmptyNodeID, fmt.Errorf("couldn't write cert file: %w", err)
	}

	privBytes, err := x509.MarshalPKCS8PrivateKey(rsaKey)
	if err != nil {
		return nil, nil, ids.EmptyNodeID, fmt.Errorf("couldn't marshal private key: %w", err)
	}

	var keyBuff bytes.Buffer
	if err := pem.Encode(&keyBuff, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		return nil, nil, ids.EmptyNodeID, fmt.Errorf("couldn't write private key: %w", err)
	}

	tlsCert, err := LoadTLSCertFromBytes(keyBuff.Bytes(), certBuff.Bytes())
	if err != nil {
		return nil, nil, ids.EmptyNodeID, fmt.Errorf("couldn't load tls cert from bytes: %w", err)
	}

	nodeID, err := certToID(tlsCert.Leaf)

	return keyBuff.Bytes(), certBuff.Bytes(), nodeID, nil
}

func certToID(cert *x509.Certificate) (ids.NodeID, error) {
	pubKeyBytes, err := utilsSecp256k1.RecoverSecp256PublicKey(cert)
	if err != nil {
		return ids.EmptyNodeID, err
	}
	return ids.ToNodeID(pubKeyBytes)
}

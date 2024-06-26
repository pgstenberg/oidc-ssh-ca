package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net"
)

func addrToPort(addr string) (int, error) {
	tcpAddr, err := net.ResolveTCPAddr("", addr)
	if err != nil {
		return 0, err
	}

	return tcpAddr.Port, nil
}
func generateSshCommand(user string, host string, port int, command string) string {
	if port == 22 {
		return fmt.Sprintf("ssh %s@%s '%s'", user, host, command)
	}

	return fmt.Sprintf("ssh %s@%s -p %d '%s'", user, host, port, command)
}

func jsonMarshalUnmarshal[T any](value any, target *T) error {
	inrec, err := json.Marshal(value)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(inrec, target); err != nil {
		return err
	}
	return nil
}

func stringSliceToBytes(target []string) [][]byte {
	l := [][]byte{}
	for _, v := range target {
		l = append(l, []byte(v))
	}
	return l
}

func generatePrivateKey(bitSize int) ([]byte, error) {

	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	d := pem.EncodeToMemory(&pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(privateKey),
	})

	return d, nil
}

package orderer

import (
	ordPb "github.com/hyperledger/fabric-protos-go/orderer"
)

// Orderer represents the details of an orderer node,
// including its host, port, and certificates.
type Orderer struct {
	Host         string
	Port         uint32
	Certificates [][]byte
}

// Pool describes common interface for pool of orderer clients
type Pool interface {
	// Get returns orderer broadcast client by orderer definition
	Get(orderer *Orderer) (ordPb.AtomicBroadcastClient, error)
	// Close gracefully closes all internal connections
	Close() error
}

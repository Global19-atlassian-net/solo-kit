// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: api/external/envoy/api/v2/core/http_uri.proto

package core

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/mitchellh/hashstructure"
	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
)

// Hash function
func (m *HttpUri) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	if _, err = hasher.Write([]byte(m.GetUri())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetTimeout(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	switch m.HttpUpstreamType.(type) {

	case *HttpUri_Cluster:

		if _, err = hasher.Write([]byte(m.GetCluster())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"go.uber.org/zap"
)

type TestingSnapshot struct {
	Mocks                MocksByNamespace
	Fakes                FakesByNamespace
	Anothermockresources AnothermockresourcesByNamespace
}

func (s TestingSnapshot) Clone() TestingSnapshot {
	return TestingSnapshot{
		Mocks:                s.Mocks.Clone(),
		Fakes:                s.Fakes.Clone(),
		Anothermockresources: s.Anothermockresources.Clone(),
	}
}

func (s TestingSnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashMocks(),
		s.hashFakes(),
		s.hashAnothermockresources(),
	)
}

func (s TestingSnapshot) hashMocks() uint64 {
	return hashutils.HashAll(s.Mocks.List().AsInterfaces()...)
}

func (s TestingSnapshot) hashFakes() uint64 {
	return hashutils.HashAll(s.Fakes.List().AsInterfaces()...)
}

func (s TestingSnapshot) hashAnothermockresources() uint64 {
	return hashutils.HashAll(s.Anothermockresources.List().AsInterfaces()...)
}

func (s TestingSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("mocks", s.hashMocks()))
	fields = append(fields, zap.Uint64("fakes", s.hashFakes()))
	fields = append(fields, zap.Uint64("anothermockresources", s.hashAnothermockresources()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}

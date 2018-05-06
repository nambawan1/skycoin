// Code generated by mockery v1.0.0
package visor

import (
	"github.com/boltdb/bolt"
	"github.com/stretchr/testify/mock"

	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/visor/historydb"
)

// historyerMock is an autogenerated mock type for the historyer type
type historyerMock struct {
	mock.Mock
}

func newHistoryerMock() *historyerMock {
	return &historyerMock{}
}

// GetAddrTxns provides a mock function with given fields: tx, address
func (_m *historyerMock) GetAddrTxns(tx *bolt.Tx, address cipher.Address) ([]historydb.Transaction, error) {
	ret := _m.Called(tx, address)

	var r0 []historydb.Transaction
	if rf, ok := ret.Get(0).(func(*bolt.Tx, cipher.Address) []historydb.Transaction); ok {
		r0 = rf(tx, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]historydb.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bolt.Tx, cipher.Address) error); ok {
		r1 = rf(tx, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddrUxOuts provides a mock function with given fields: tx, address
func (_m *historyerMock) GetAddrUxOuts(tx *bolt.Tx, address cipher.Address) ([]*historydb.UxOut, error) {
	ret := _m.Called(tx, address)

	var r0 []*historydb.UxOut
	if rf, ok := ret.Get(0).(func(*bolt.Tx, cipher.Address) []*historydb.UxOut); ok {
		r0 = rf(tx, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*historydb.UxOut)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bolt.Tx, cipher.Address) error); ok {
		r1 = rf(tx, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransaction provides a mock function with given fields: tx, hash
func (_m *historyerMock) GetTransaction(tx *bolt.Tx, hash cipher.SHA256) (*historydb.Transaction, error) {
	ret := _m.Called(tx, hash)

	var r0 *historydb.Transaction
	if rf, ok := ret.Get(0).(func(*bolt.Tx, cipher.SHA256) *historydb.Transaction); ok {
		r0 = rf(tx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*historydb.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bolt.Tx, cipher.SHA256) error); ok {
		r1 = rf(tx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUxout provides a mock function with given fields: tx, uxid
func (_m *historyerMock) GetUxout(tx *bolt.Tx, uxid cipher.SHA256) (*historydb.UxOut, error) {
	ret := _m.Called(tx, uxid)

	var r0 *historydb.UxOut
	if rf, ok := ret.Get(0).(func(*bolt.Tx, cipher.SHA256) *historydb.UxOut); ok {
		r0 = rf(tx, uxid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*historydb.UxOut)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bolt.Tx, cipher.SHA256) error); ok {
		r1 = rf(tx, uxid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseBlock provides a mock function with given fields: tx, b
func (_m *historyerMock) ParseBlock(tx *bolt.Tx, b *coin.Block) error {
	ret := _m.Called(tx, b)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bolt.Tx, *coin.Block) error); ok {
		r0 = rf(tx, b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ParsedHeight provides a mock function with given fields: tx
func (_m *historyerMock) ParsedHeight(tx *bolt.Tx) (int64, error) {
	ret := _m.Called(tx)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*bolt.Tx) int64); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bolt.Tx) error); ok {
		r1 = rf(tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetIfNeed provides a mock function with given fields: tx
func (_m *historyerMock) ResetIfNeed(tx *bolt.Tx) error {
	ret := _m.Called(tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bolt.Tx) error); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

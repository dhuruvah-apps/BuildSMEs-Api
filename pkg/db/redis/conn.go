package redis

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/pkg/logger"
	"github.com/dgraph-io/badger"
)

type BadgerStore struct {
	db  *badger.DB
	ttl time.Duration
}

const BDGSSTPREFIX = "sst:"
const BDGU64PREFIX = "u64:"

var dbSstPrefix = []byte(BDGSSTPREFIX)
var dbU64Prefix = []byte(BDGU64PREFIX)
var ErrKeyNotFound = errors.New("not found")

// Returns new redis client
func NewRedisClient(cfg *config.Config, logger logger.Logger) *BadgerStore {
	if err := os.MkdirAll(cfg.Redis.RedisAddr, 0774); err != nil {
		return nil
	}

	opts := badger.DefaultOptions(cfg.Redis.RedisAddr)

	db, err := badger.Open(opts)
	if err != nil {
		logger.Error("Bagder store error")
	}

	store := &BadgerStore{db: db, ttl: time.Duration(cfg.Session.Expire) * time.Second}

	return store
}

func (b *BadgerStore) Close() error {
	return b.db.Close()
}

// Get a value in StableStore.
func (b *BadgerStore) Get(k []byte) ([]byte, error) {
	return b.GetRaw(sstKeyOf(k))
}

// Set a key/value in StableStore.
func (b *BadgerStore) Set(k []byte, v []byte) error {
	return b.SetRaw(sstKeyOf(k), v)
}

// Set a key/value in StableStore.
func (b *BadgerStore) Del(k []byte) error {
	return b.DeleteRaw(sstKeyOf(k))
}

// SetUint64 is like Set, but handles uint64 values
func (b *BadgerStore) SetUint64(key []byte, val uint64) error {
	return b.SetRaw(u64KeyOf(key), uint64ToBytes(val))
}

// GetUint64 is like Get, but handles uint64 values
func (b *BadgerStore) GetUint64(key []byte) (uint64, error) {
	val, err := b.GetRaw(u64KeyOf(key))
	if err != nil {
		return 0, err
	}
	return bytesToUint64(val), nil
}

func (b *BadgerStore) GetRaw(k []byte) ([]byte, error) {
	txn := b.db.NewTransaction(false)
	defer txn.Discard()
	item, err := txn.Get(k)
	if item == nil {
		return nil, ErrKeyNotFound
	}
	if err != nil {
		return nil, err
	}
	v, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}
	if err := txn.Commit(); err != nil {
		return nil, err
	}
	return append([]byte(nil), v...), nil
}

func (b *BadgerStore) SetRaw(k []byte, v []byte) error {
	return b.db.Update(func(txn *badger.Txn) error {
		return txn.SetEntry(badger.NewEntry(k, v).WithTTL(b.ttl))
	})
}

func (b *BadgerStore) DeleteRaw(key []byte) error {
	return b.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}

func sstKeyOf(rawKey []byte) []byte {
	key := fmt.Sprintf("%s%s", dbSstPrefix, hex.EncodeToString(rawKey))

	return []byte(key)
}

func bytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func uint64ToBytes(u uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, u)
	return buf
}

func u64KeyOf(rawKey []byte) []byte {
	key := fmt.Sprintf("%s%s", dbU64Prefix, hex.EncodeToString(rawKey))

	return []byte(key)
}

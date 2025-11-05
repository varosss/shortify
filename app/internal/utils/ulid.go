package utils

import (
	"math/rand"
	"sync"
	"time"

	"github.com/oklog/ulid/v2"
)

var (
	entropyMu sync.Mutex
	entropy   = ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
)

const SHORT_ULID_LEN = 10

func GenerateShortULID() string {
	entropyMu.Lock()
	defer entropyMu.Unlock()

	id := ulid.MustNew(ulid.Timestamp(time.Now()), entropy)
	code := id.String()

	return code[len(code)-SHORT_ULID_LEN:]
}

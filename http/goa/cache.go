package goa

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

var cacheOffset uint64

// HTTPCacheKey builds the cache key for the fizz/buzz method
// in addition with the classic parameters, it adds an automatic expiration of cacheDelay.
// You can expires a responses by using the
func HTTPCacheKey(cacheDelay time.Duration, string1 string, string2 string, int1 int, int2 int, limit int) string {
	// cached for
	return strconv.Itoa(int(time.Now().Truncate(cacheDelay).Unix())) + "_" +
		// cache offset to expire the cache manually
		strconv.Itoa(int(atomic.LoadUint64(&cacheOffset))) + "_" +
		string1 + "_" +
		string2 + "_" +
		strconv.Itoa(int1) + "_" +
		strconv.Itoa(int2) + "_" +
		strconv.Itoa(limit)
}

// ExpiresAllKeys expireds all cache content
func ExpiresAllKeys() {
	cc := atomic.LoadUint64(&cacheOffset)
	cc += rand.Uint64()
	atomic.StoreUint64(&cacheOffset, cc)
}

func ParamsFromCacheKey(k string) (string1, string2 string, int1, int2, limit int, err error) {

	var res []string
	if res = strings.SplitN(k, "_", 7); len(res) != 7 {
		err = errors.New("invalid key, must have 8 part")
		return
	}
	string1 = res[2]
	string2 = res[3]

	var i int
	if i, err = strconv.Atoi(res[4]); err != nil {
		return
	}
	int1 = i
	if i, err = strconv.Atoi(res[5]); err != nil {
		return
	}
	int2 = i
	if i, err = strconv.Atoi(res[6]); err != nil {
		return
	}
	limit = i
	return
}

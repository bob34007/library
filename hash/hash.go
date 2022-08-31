package hash

import (
	"hash/fnv"

	"github.com/pingcap/tidb/util/hack"
)

func String2Hash(s string) (uint64, error) {
	h := fnv.New64a()
	ss := hack.Slice(s)
	_, err := h.Write(ss)
	if err != nil {
		return 0, err
	}
	return h.Sum64(), nil
}

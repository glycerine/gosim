package go125

import "github.com/glycerine/gosim/gosimruntime"

func MathRand_runtime_rand() uint64 {
	return gosimruntime.Fastrand64()
}

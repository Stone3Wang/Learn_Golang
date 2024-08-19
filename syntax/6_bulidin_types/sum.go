package main

func Sum(vals []int) int {
	var res int
	for _, v := range vals {
		res += v
	}
	return res
}

func SumInt64(vals []int64) int64 {
	var res int64
	for _, v := range vals {
		res += v
	}
	return res
}
func SumInt32(vals []int32) int32 {
	var res int32
	for _, v := range vals {
		res += v
	}
	return res
}

func Keys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for K := range m {
		keys = append(keys, K)
	}
	return keys
}

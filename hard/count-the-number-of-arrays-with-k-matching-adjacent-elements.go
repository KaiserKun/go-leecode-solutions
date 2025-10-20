package hard

// link: https://leetcode.cn/problems/count-the-number-of-arrays-with-k-matching-adjacent-elements/
// thought: 组合数学，先选出k个位置放相同元素，然后剩下的位置随便放
const (
	mod = 1e9 + 7
	mx  = 100000
)

var F, invF [mx]int

func qpow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func init() {
	F[0] = 1
	for i := 1; i < mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF[mx-1] = qpow(F[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}

func comb(n, m int) int {
	return F[n] * invF[m] % mod * invF[n-m] % mod
}

func countGoodArrays(n int, m int, k int) int {
	return comb(n-1, k) * m % mod * qpow(m-1, n-k-1) % mod
}

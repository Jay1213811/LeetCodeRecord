package main
//快速幂模版题
func superPow(a int, b []int) int {
	var dfs func(a, mod int, b[] int) int
	dfs = func(a, mod int, b[] int) int {
		if len(b) == 0 || a == 1 {
			return 1
		}
		return quickPow(dfs(a, mod, b[:len(b)-1]), 10, mod) * quickPow(a, b[len(b)-1], mod) % mod
	}

	mod := 1337
	a %= mod
	return dfs(a, mod, b)
}

func quickPow(a, b, mod int) int {
	ans := 1
	a %= mod
	for b != 0 {
		if (b & 1) == 1 {
			ans = ans * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return ans
}

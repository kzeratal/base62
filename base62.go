package base62

var charactersSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Encode(num uint64) string {
	res := ""
	for num > 0 {
		res = string(charactersSet[num%62]) + res
		num /= 62
	}
	return res
}
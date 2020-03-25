package iteration


func Repeat(s string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += s
	}
	return repeated
}

func Revert(s string) string {
	var reversed []rune
	r := []rune(s)

	for i := len(s) - 1; i >=0; i-- {
		reversed = append(reversed, r[i])
	}
	return string(reversed)
}

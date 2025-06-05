package gen

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

var (
	lower    = "abcdefghijklmnopqrstuvwxyz"
	upper    = strings.ToUpper(lower)
	numbers  = "0123456789"
	specials = "@#$%&*()_+-={}[]|<>.,~"
)

type PswdInfo struct {
	Len     int
	Lower   bool
	Upper   bool
	Numbers bool
	Special bool
}

func NewPswdInfo(l int, lower, upper, num, spe bool) PswdInfo {
	return PswdInfo{
		Len:     l,
		Lower:   lower,
		Upper:   upper,
		Numbers: num,
		Special: spe,
	}
}

func setDictionary(info PswdInfo) []string {
	dict := []string{}
	if info.Lower {
		dict = append(dict, lower)
	}
	if info.Upper {
		dict = append(dict, upper)
	}
	if info.Numbers {
		dict = append(dict, numbers)
	}
	if info.Special {
		dict = append(dict, specials)
	}
	return dict
}

func Gen(info PswdInfo) {
	if !info.Lower && !info.Upper && !info.Numbers && !info.Special {
		info.Lower = true
		info.Upper = true
		info.Numbers = true
		info.Special = true
	}

	dict := setDictionary(info)

	pswd := make([]rune, info.Len)
	for i := range pswd {
    charset := dict[rand.IntN(len(dict))]
		pswd[i] = rune(charset[rand.IntN(len(charset))])
	}

	fmt.Println(string(pswd))

  for i := range pswd {
		pswd[i] = rune(0)
	}
}

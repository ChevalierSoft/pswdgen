package gen

import (
	"fmt"
	"math/rand"
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

func setDictionary(info *PswdInfo) string {
	dict := ""
	if info.Lower {
		dict += lower
	}
	if info.Upper {
		dict += upper
	}
	if info.Numbers {
		dict += numbers
	}
	if info.Special {
		dict += specials
	}
	return dict
}

func Gen(info *PswdInfo) {
	if !info.Lower && !info.Upper && !info.Numbers && !info.Special {
		info.Lower = true
		info.Upper = true
		info.Numbers = true
		info.Special = true
	}
	lib := setDictionary(info)
	pswd := make([]rune, info.Len)
	for i := range pswd {
		pswd[i] = rune(lib[rand.Intn(len(lib))])
	}
	fmt.Println(string(pswd))
}

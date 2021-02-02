package regcheck

import (
	"net/url"
	"regexp"
	"strconv"
	"time"
	"unicode"
)

/************************* 自定义类型 ************************/
//数字+字母  不限制大小写
func IsID(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[0-9a-zA-Z]+$", s)
		if false == b {
			return b
		}
	}
	return b
}

//数字+字母+符号 6~30位
func IsPwd(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[0-9a-zA-Z@.]+$", s)
		if false == b {
			return b
		}
	}
	return b
}

/************************* 数字类型 ************************/
//纯整数
func IsInteger(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[0-9]+$", s)
		if false == b {
			return b
		}
	}
	return b
}

//纯小数
func IsDecimals(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^\\d+\\.[0-9]+$", s)
		if false == b {
			return b
		}
	}
	return b
}

//家用电话（不带前缀） 最高8位
func IsTelephone(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[0-9]{11}$", s)
		if false == b {
			return b
		}
	}
	return b
}

func IsCertIfID(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[0-9]{18}|([0-9]{17}x|X)$", s)
		if false == b {
			return b
		}
	}
	return b
}

/************************* 英文类型 *************************/
//仅小写
func IsEngishLowCase(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[a-z]+$", s)
		if false == b {
			return b
		}
	}
	return b
}

//仅大写
func IsEnglishCap(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[A-Z]+$", s)
		if false == b {
			return b
		}
	}
	return b
}

//大小写混合
func IsEnglish(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[A-Za-z]+$", s)
		if false == b {
			return b
		}
	}
	return b
}

//邮箱 最高30位
func IsEmail(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", s)
		if false == b {
			return b
		}
	}
	return b
}

//汉字
func IsHan(str ...string) bool {
	set := []*unicode.RangeTable{unicode.Han}
	rue := []rune(str[0])
	for _, r := range rue {
		if !unicode.IsOneOf(set, r) {
			return false
		}
	}
	return true
}

//汉字 字母 数字
func IsHanLM(str ...string) bool {
	set := []*unicode.RangeTable{unicode.Han, unicode.L, unicode.M, unicode.N}
	for _, s := range str {
		rue := []rune(s)
		for _, r := range rue {
			if !unicode.IsOneOf(set, r) {
				return false
			}
		}
	}
	return true
}

func IsPrint(str ...string) bool {
	for _, s := range str {
		rue := []rune(s)
		for _, r := range rue {
			if !unicode.IsPrint(r) {
				return false
			}
		}
	}
	return true
}

func IsURL(str ...string) bool {
	_, err := url.ParseRequestURI(str[0])
	if err != nil {
		return false
	}
	return true
}

func IsMMYY(expdt string) bool {
	if len(expdt) != 4 {
		return false
	}
	mth, err := strconv.Atoi(expdt[:2])
	if err != nil {
		return false
	}
	if mth < 1 || mth > 12 {
		return false
	}
	year, err := strconv.Atoi(expdt[2:])
	if err != nil {
		return false
	}

	now := time.Now().Year() % 100
	if year > now+10 || year < now {
		return false
	}
	return true
}

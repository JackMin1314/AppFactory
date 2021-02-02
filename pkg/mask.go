package tools

//获取掩码data为元数据，start ，end表示保留首位明码的位数,num表示掩码的位数
func GetMask(data string, start, end, num int) (mask string) {
	runeArr := []rune(data)
	length := len(runeArr)
	if start <= length && end <= length { //允许  呵呵打  变为   呵呵打***呵呵打
		for i := 0; i < num; i++ {
			mask += "*"
		}
		return string(runeArr[:start]) + mask + string(runeArr[length-end:])
	}
	return
}

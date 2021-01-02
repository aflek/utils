package utils


//СharIsNum - функция проверки: Символ есть число
func СharIsNum(s byte) bool {
	
	if s> 47 && s < 58 {
		return true
	}
	return false
}

//StrIsNum - функция проверки: Срока состоит только из чисел
func StrIsNum(s string) bool {
    for i:=0; i<len(s); i++ {
		c:= s[i]
        if !СharIsNum(c) {
            return false
        }
    }
    return true
}
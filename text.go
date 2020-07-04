package mycaptcha




const txtChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// GetRandText 生成随机字符的字符串
func GetRandText(n int) string {
	l:=len(txtChars)
	
	res:=make([]byte,n)
	for i:=0;i<n;i++{
		res[i]=txtChars[r.Intn(l)]
	}
	
	return string(res)
}
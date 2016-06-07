package goutility

func CheckErr(err error) bool {
	if err != nil {
		panic(err)
		return true
	}
	return false
}

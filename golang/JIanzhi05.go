package main

/* func replaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
} */

func replaceSpace(s string) string {
	tmp := []byte(s)
	ret := []byte{}
	for _, each := range tmp {
		if each == byte(' ') {
			ret = append(ret, []byte("%20")...)
		} else {
			ret = append(ret, each)
		}
	}
	return string(ret)

}

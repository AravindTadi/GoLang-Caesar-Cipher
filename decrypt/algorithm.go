package decrypt

func Decmsg(h string) string {
	v := ""
	for _, i := range h {
		ascii := int(i)
		conv := string(ascii - 3)
		v = v + conv

	}

	return v
}

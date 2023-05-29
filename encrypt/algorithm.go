package encrypt

func Encmsg(s string) string {
	encmssg := ""
	for _, i := range s {
		ascii := int(i)
		msg_ := string(ascii + 3)
		encmssg = encmssg + msg_

	}
	return encmssg
}

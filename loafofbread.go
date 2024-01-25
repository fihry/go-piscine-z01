package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	} else if len(str) < 5 {
		return "Invalid Output\n"
	} else {
		jump := ""
		i := 0
		for _, e := range str {
			if i%6 != 5 && e == ' ' {
				continue
			}
			if i%6 == 5 {
				jump += " "
			} else {
				jump += string(e)
			}
			i++
		}
		for i := len(jump) - 1; i >= 0; i-- {
			if jump[i] != ' ' {
				jump = jump[:i+1]
				break
			}
		}
		return jump + "\n"
	}
}

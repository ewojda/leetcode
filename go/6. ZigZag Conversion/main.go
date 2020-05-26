package main

func main() {}

func convert1(s string, numRows int) string {
	pillarlen := numRows
	angledlen := numRows - 2
	//gaplen := (pillarlen - angledlen) / 2
	//gap := strings.Repeat(string(byte(0)), gaplen)

	getColumns := func() []string {
		columns := make([]string, len(s)/(pillarlen+angledlen)+2)
		pillar := false
		remaining := s
		for {
			pillar = !pillar
			if pillar {
				if len(remaining) >= pillarlen {
					columns = append(columns, remaining[:pillarlen])
					remaining = remaining[:pillarlen]
				} else if len(remaining) > 0 {
					columns = append(columns, remaining)
					remaining = ""
				} else {
					break
				}
			} else {
				//Angled
				if len(remaining) >= angledlen {
					//columns = append(columns, gap + strings.)
				} else if len(remaining) > 0 {

				} else {
					break
				}
			}
		}
		return nil
	}
	result := ""
	columns := getColumns()
	for row := 0; row < pillarlen; row++ {
		for column := 0; column < len(columns); column++ {
			if char := columns[column][row]; char != 0 {
				result += string(char)
			}
		}
	}
	return result
}

func convert(s string, numRows int) string {
	if len(s) <= 2 {
		return s
	}
	pillarlen := numRows
	angledlen := numRows - 2
	if angledlen < 0 {
		angledlen = 0
	}
	gaplen := 1
	charAt := func(row, column int) (byte, bool) {
		if column%2 == 0 {
			//Straight
			i := (column/2)*pillarlen + (column/2)*angledlen + row
			if i < len(s) {
				return s[i], true
			}
		} else {
			//Angled
			if row < gaplen || row >= angledlen+gaplen {
				return 0, false
			}

			i := (column/2+1)*pillarlen + (column/2)*angledlen + angledlen - 1 - (row - gaplen)
			if i < len(s) {
				return s[i], true
			}
		}
		return 0, false
	}
	columns := len(s)/(pillarlen+angledlen)*2 + 2
	result := make([]byte, 0, len(s))
	for row := 0; row < pillarlen; row++ {
		for column := 0; column < columns; column++ {
			if char, ok := charAt(row, column); ok {
				result = append(result, char)
			}
		}
	}
	return string(result)
}

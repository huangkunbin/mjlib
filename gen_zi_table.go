package mjlib

var zi_gui_tested = [9]*map[int]bool{}
var zi_gui_eye_tested = [9]*map[int]bool{}

// 1→0
// 10→10
// 2→110
// 20→1110
// 3→11110
// 30→111110
// 4→1111110
// 40→11111110
func zi_check_add(cards []int, gui_num int, eye bool) bool {
	key := 0
	p := -1

	for i := 0; i < 7; i++ {
		if cards[i] > 0 {
			if cards[i] > 4 {
				key = 0
				break
			}
			p++
			switch cards[i] {
			case 2:
				key |= 0x3 << uint32(p)
				p += 2
			case 3:
				key |= 0xF << uint32(p)
				p += 4
			case 4:
				key |= 0x3F << uint32(p)
				p += 6
			}
			key |= 0x1 << uint32(p)
			p++
		}
	}

	var m *map[int]bool
	if !eye {
		m = zi_gui_tested[gui_num]
	} else {
		m = zi_gui_eye_tested[gui_num]
	}
	_, ok := (*m)[key]

	if ok {
		return false
	}

	(*m)[key] = true

	if key == 0 {
		return true
	}

	MTableMgr.Add(key, gui_num, eye, false)

	return true
}

func parse_zi_table_sub(cards []int, num int, eye bool) {
	for i := 0; i < 7; i++ {
		if cards[i] == 0 {
			continue
		}

		cards[i]--

		if !zi_check_add(cards, num, eye) {
			cards[i]++
			continue
		}

		if num < 8 {
			parse_zi_table_sub(cards, num+1, eye)
		}
		cards[i]++
	}
}

func parse_zi_table(cards []int, eye bool) {
	if !zi_check_add(cards, 0, eye) {
		return
	}
	parse_zi_table_sub(cards, 1, eye)
}

func gen_3(cards []int, level int, eye bool, maxLevel int) {
	for i := 0; i < 7; i++ {
		if cards[i] > 3 {
			continue
		}
		cards[i] += 3

		parse_zi_table(cards, eye)
		if level < 5 {
			gen_3(cards, level+1, eye, maxLevel)
		}

		cards[i] -= 3
	}
}

func gen_zi_table(maxLevel int) {
	for i := 0; i < 9; i++ {
		zi_gui_tested[i] = &map[int]bool{}
		zi_gui_eye_tested[i] = &map[int]bool{}
	}

	cards := []int{
		0, 0, 0, 0, 0, 0, 0,
	}

	// 无眼
	// fmt.Printf("无眼表生成开始\n")
	gen_3(cards, 1, false, maxLevel)
	// fmt.Printf("无眼表生成结束\n")

	// 有眼
	// fmt.Printf("有眼表生成开始\n")
	for i := 0; i < 7; i++ {
		cards[i] = 2
		// fmt.Printf("将 %d \n", i)
		parse_zi_table(cards, true)
		gen_3(cards, 1, true, maxLevel)
		cards[i] = 0
	}
	// fmt.Printf("有眼表生成结束\n")

	zi_gui_tested = [9]*map[int]bool{}
	zi_gui_eye_tested = [9]*map[int]bool{}

	// fmt.Printf("表数据存储开始\n")
	// MTableMgr.DumpziTable()
	// fmt.Printf("表数据存储结束\n")
}

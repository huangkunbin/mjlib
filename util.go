package mjlib

type utilMod struct {
}

func (this *utilMod) GetEye(cards []int32) int {
	for i := 0; i < 34; i++ {
		if cards[i] >= 2 {
			cards[i] -= 2

			var gui, eye_num int32

			gui, eye_num = HuMod._split(cards, 0, 0, 8, true, 0)
			if gui != 0 || eye_num != 0 {
				cards[i] += 2
				continue
			}
			gui, eye_num = HuMod._split(cards, 0, 9, 17, true, 0)
			if gui != 0 || eye_num != 0 {
				cards[i] += 2
				continue
			}
			gui, eye_num = HuMod._split(cards, 0, 18, 26, true, 0)
			if gui != 0 || eye_num != 0 {
				cards[i] += 2
				continue
			}
			gui, eye_num = HuMod._split(cards, 0, 27, 33, false, 0)
			if gui != 0 || eye_num != 0 {
				cards[i] += 2
				continue
			}

			cards[i] += 2
			return i
		}
	}

	return -1
}

func (this *utilMod) GetTingCards(cards []int32, gui_1 int, gui_2 int) (tingCards []int) {
	var (
		gui_num_1 int32
		gui_num_2 int32
	)

	if gui_1 != INVALID_CARD {
		gui_num_1 = cards[gui_1]
		cards[gui_1] = 0
	}

	if gui_2 != INVALID_CARD {
		gui_num_2 = cards[gui_2]
		cards[gui_2] = 0
	}

	for i := 0; i < 34; i++ {

		if i != gui_1 && i != gui_2 {
			if cards[i]+1 > 4 {
				continue
			}
			cards[i]++
		} else {
			if i == gui_1 {
				if gui_num_1+1 > 4 {
					continue
				}
				gui_num_1++
			}
			if i == gui_2 {
				if gui_num_2+1 > 4 {
					continue
				}
				gui_num_2++
			}
		}

		if HuMod.split(cards, gui_num_1+gui_num_2) {
			tingCards = append(tingCards, i)
		}

		if i != gui_1 && i != gui_2 {
			cards[i]--
		} else {
			if i == gui_1 {
				gui_num_1--
			}
			if i == gui_2 {
				gui_num_2--
			}
		}

	}

	if gui_1 != INVALID_CARD {
		cards[gui_1] = gui_num_1
	}

	if gui_2 != INVALID_CARD {
		cards[gui_2] = gui_num_2
	}

	return
}

func (this *utilMod) GetCardListByNum(cards []int32, num int32) (cardList []int) {
	for i := 0; i < 34; i++ {
		if cards[i] == num {
			cardList = append(cardList, i)
		}
	}
	return
}

func (this *utilMod) GetAAAList(cards []int32) (aaaList []int) {
	for i := 0; i < 34; i++ {
		if cards[i] >= 3 {
			aaaList = append(aaaList, i)
		}
	}
	return
}

// IsZuHeLong
func (g *GameLogic) IsZuHeLong(cardIndex []int16, cardCount int16, weaveItem []WeaveItem, weaveCount int32, currentCard int32) bool {
	// 手牌牌数不对直接返回false
	if cardCount%3 != 2 {
		return false
	}
	index := -1
	lastIndex := []int{-1, -1, -1}

	cardIndexTemp := make([]int16, g.maxIndex)
	copy(cardIndexTemp, cardIndex)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if cardIndexTemp[i*9+j] > 0 && cardIndexTemp[i*9+j+3] > 0 && cardIndexTemp[i*9+j+6] > 0 {
				if index == -1 {
					index = j
					if i == 0 {
						lastIndex[i] = index
						cardIndexTemp[i*9+j]--
						cardIndexTemp[i*9+j+3]--
						cardIndexTemp[i*9+j+6]--
					} else {
						for in, v := range lastIndex {
							if in == i {
								lastIndex[i] = index
								cardIndexTemp[i*9+j]--
								cardIndexTemp[i*9+j+3]--
								cardIndexTemp[i*9+j+6]--
								break
							}
							if v == index {
								return false
							}
						}
					}
				} else {
					return false
				}
			}
			index = -1
		}
		if lastIndex[i] == -1 {
			return false
		}
	}
	// fmt.Println("cardIndexTemp", cardIndexTemp)
	var pos [14]int16
	// 剩余手牌查表看是否和牌
	key := g.CalcKey(cardIndexTemp, pos[0:])
	cardType := g.mahjong.GetType(key)
	if cardType == nil {
		return false
	}
	return true
}
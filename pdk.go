// OnNewTimer 定时器
func (t *Table) OnNewTimer(id int, parameter interface{}) error {
	t.log.Infoln("OnNewTimer", t.tableFrame.GetRoomId(), t.currentStatus, t.roomType, t.currentUser, t.timeoutStat, id)
	if id != 1 {
		return nil
	}
	if t.currentStatus == pdk.StatusPlay {
		if t.currentUser != define.INVALID_CHAIR_ID {
			userItem := t.tableFrame.GetTableUserItem(int(t.currentUser))
			if userItem == nil {
				return nil
			}
			var outPokerResult OutPokerResult
			turnPokerCount := t.turnPokerCount
			if t.currentUser == t.bankerUser && t.handCount[t.currentUser] == t.handPokerCount {
				turnPokerCount = -1
			}
			can := t.gameLogic.SearchOutPoker(t.handPoker[t.currentUser], t.handCount[t.currentUser], t.turnPoker, turnPokerCount, &outPokerResult)
			outCard := &message.PdkOutPoker{}
			if can {
				for i := range outPokerResult.CanOutPoker {
					outCard.Type = proto.Int32(1)
					outCard.Poker = outPokerResult.CanOutPoker[i]
					t.log.Infoln("OnNewTimer OutCard", outCard)
					data, err := proto.Marshal(outCard)
					if err != nil {
						return err
					}
					if t.OnUserOutPoker(data, userItem) {
						break
					}
				}
			} else {
				outCard.Type = proto.Int32(2)

				t.log.Infoln("OnNewTimer OutCard", outCard)
				data, err := proto.Marshal(outCard)
				if err != nil {
					return err
				}

				t.OnUserOutPoker(data, userItem)
			}
			if !userItem.IsOnline() {
				t.timeoutStat[userItem.GetChairID()]++
			}
			return nil
		}
	}
	return nil
}
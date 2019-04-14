package deck

// sort Deck first by suit, and then by face...
// ==========================DESC===========================//
type sortDeckDESC Deck

func (sd sortDeckDESC) Len() int {
	return len(sd)
}

func (sd sortDeckDESC) Less(i, j int) bool {
	// (suit) order the numbers from least to greatest (in ascending)
	// (face) order the numbers from greatest to least (in descending)

	if sd[i].suit < sd[j].suit {
		return true
	} else if sd[i].suit == sd[j].suit {
		if sd[i].face > sd[j].face {
			return true
		}
		return false
	}
	return false
}

func (sd sortDeckDESC) Swap(i, j int) {
	sd[i], sd[j] = sd[j], sd[i]
}

// ==========================ASC===========================//

type sortDeckASC Deck

func (sd sortDeckASC) Len() int {
	return len(sd)
}

func (sd sortDeckASC) Less(i, j int) bool {
	// (suit) order the numbers from greatest to least (in descending)
	// (face) order the numbers from least to greatest (in ascending)

	if sd[i].suit > sd[j].suit {
		return true
	} else if sd[i].suit == sd[j].suit {
		if sd[i].face < sd[j].face {
			return true
		}
		return false
	}
	return false
}

func (sd sortDeckASC) Swap(i, j int) {
	sd[i], sd[j] = sd[j], sd[i]
}

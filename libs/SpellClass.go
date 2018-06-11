package libs

import (
	"net"
)

func Fireball(Caster User, Recipient string, L map[string]net.Conn) {
	Power := 10
	Cost := 5
	//Check to make sure they are both Online
	if _, ok := L[Recipient]; ok {
		//check Mana
		if RetrieveMana(Caster.Username) >= Cost {
			// Lower Mana
			ChangeMana(Recipient, Cost, Caster.Conn)
			// Change Health
			ChangeHealth(Recipient, Power, L[Recipient])

			if RetrieveHealth(Recipient) >= 0 {
				ServerPrivateMessage(L[Recipient], "You were fireballed by "+Caster.Username)
				ServerPrivateMessage(Caster.Conn, "You fireballed "+Recipient)

			} else {
				ServerPrivateMessage(L[Recipient], "You Died!")
			}

		} else {
			ServerPrivateMessage(Caster.Conn, "Not enough mana!")
		}

	} else {
		ServerPrivateMessage(Caster.Conn, Recipient+", does not appear to be online!")
	}

}
func FrostBolt(Caster User, Recipient string, L map[string]net.Conn) {
	Power := 10
	Cost := 5
	//Check to make sure they are both Online
	if _, ok := L[Recipient]; ok {
		//check Mana
		if RetrieveMana(Caster.Username) >= Cost {
			// Lower Mana
			ChangeMana(Recipient, Cost, Caster.Conn)
			// Change Health
			ChangeHealth(Recipient, Power, L[Recipient])

			if RetrieveHealth(Recipient) >= 0 {
				ServerPrivateMessage(L[Recipient], "You were fireballed by "+Caster.Username)
				ServerPrivateMessage(Caster.Conn, "You fireballed "+Recipient)

			} else {
				ServerPrivateMessage(L[Recipient], "You Died!")
			}

		} else {
			ServerPrivateMessage(Caster.Conn, "Not enough mana!")
		}

	} else {
		ServerPrivateMessage(Caster.Conn, Recipient+", does not appear to be online!")
	}

}
func ArcaneMissles(Caster User, Recipient string, L map[string]net.Conn) {
	Power := 10
	Cost := 5
	//Check to make sure they are both Online
	if _, ok := L[Recipient]; ok {
		//check Mana
		if RetrieveMana(Caster.Username) >= Cost {
			// Lower Mana
			ChangeMana(Recipient, Cost, Caster.Conn)
			// Change Health
			ChangeHealth(Recipient, Power, L[Recipient])

			if RetrieveHealth(Recipient) >= 0 {
				ServerPrivateMessage(L[Recipient], "You were fireballed by "+Caster.Username)
				ServerPrivateMessage(Caster.Conn, "You fireballed "+Recipient)

			} else {
				ServerPrivateMessage(L[Recipient], "You Died!")
			}

		} else {
			ServerPrivateMessage(Caster.Conn, "Not enough mana!")
		}

	} else {
		ServerPrivateMessage(Caster.Conn, Recipient+", does not appear to be online!")
	}

}
func NickySmash(Caster User, Recipient string, L map[string]net.Conn) {
	Power := 10
	Cost := 5
	//Check to make sure they are both Online
	if _, ok := L[Recipient]; ok {
		//check Mana
		if RetrieveMana(Caster.Username) >= Cost {
			// Lower Mana
			ChangeMana(Recipient, Cost, Caster.Conn)
			// Change Health
			ChangeHealth(Recipient, Power, L[Recipient])

			if RetrieveHealth(Recipient) >= 0 {
				ServerPrivateMessage(L[Recipient], "You were fireballed by "+Caster.Username)
				ServerPrivateMessage(Caster.Conn, "You fireballed "+Recipient)

			} else {
				ServerPrivateMessage(L[Recipient], "You Died!")
			}

		} else {
			ServerPrivateMessage(Caster.Conn, "Not enough mana!")
		}

	} else {
		ServerPrivateMessage(Caster.Conn, Recipient+", does not appear to be online!")
	}

}

package gift

type ConfigEntity struct {
	ID       uint64
	GiftList []*Gift
}

type Gift struct {
	Title string
	Pic   string
	Desc  string
}

func (c *ConfigEntity) AllowSendGift() bool {
	return true
}

func (c *ConfigEntity) HasQB() bool {
	for _, gift := range c.GiftList {
		if gift.Desc == "QB" {
			return true
		}
	}

	return false
}

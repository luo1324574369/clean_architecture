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

func (s *ConfigEntity) HasQB() bool {
	for _, gift := range s.GiftList {
		if gift.Desc == "QB" {
			return true
		}
	}

	return false
}

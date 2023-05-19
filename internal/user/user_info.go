package user

type Info struct {
	Uid     int64
	GameUid int64
	isSign  bool
}

func (i *Info) SetSign(isSign bool) {
	i.isSign = isSign
}

package user

type Info struct {
	Uid     int64
	GameUid int64
	isSign  bool
}

func (i *Info) setSign(isSign bool) {
	i.isSign = isSign
}

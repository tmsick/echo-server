package domain

type (
	UserID       string
	UserName     string
	UserEmail    string
	UserPassword string
)

func (i UserID) String() string {
	return string(i)
}

func (n UserName) String() string {
	return string(n)
}

func (e UserEmail) String() string {
	return string(e)
}

func (p UserPassword) String() string {
	return string(p)
}

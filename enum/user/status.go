package user

type Status int

const (
	Guest Status = iota
	ProvisionalMember
	RegularMember
	Withdrawn
)

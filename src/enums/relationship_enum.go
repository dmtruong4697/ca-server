package enums

type Status string

const (
	FRIEND_REQUEST_PENDING Status = "FRIEND_REQUEST_PENDING"
	FRIEND                 Status = "FRIEND"
	BLOCKED                Status = "BLOCKED"
)

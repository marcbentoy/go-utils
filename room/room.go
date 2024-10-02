package sdkroom

type Room struct {
	Name string
}

func NewRoom(name string) *Room {
	return &Room{
		Name: name,
	}
}

package service

type userService struct {
}

var UserService = userService{}

func (user userService) Test() *map[string]int {
	println("Hello, my name is lhf")

	result := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}

	return &result
}

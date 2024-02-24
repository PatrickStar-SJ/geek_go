package main

func UseUser() {
	u, err := GetUser(12)
	if err != nil {
		println(err)
		return
	}
	// u.Name
	println(u.Name)
}

type Optional[T any] struct {
	Val T
}

func (o Optional[T]) Apply(fn func(t T)) {
	var val any = o.Val
	if val == nil {
		return
	}
	fn(o.Val)
}

func OfNullable[T any](t T) Optional[T] {
	return Optional[T]{
		Val: t,
	}
}

// GetUser 最佳实践，没有 error，*User 一定不为 nil
func GetUser(id int64) (*User, error) {
	return &User{}, nil
}

func Component() {
	var u User // 这个时候。
	// Address就已经初始化了，零值（但不是 nil）
	println(u.Friend.name) // panic，因为 Friend 是指针，所以初始化的是 nil

	var u1 User = User{
		Friend: &Friend{},
	}
	println(u1.Friend.name)
}

type User struct {
	Name string
	Address
	*Friend
}

type Address struct {
}

type Friend struct {
	name string
}

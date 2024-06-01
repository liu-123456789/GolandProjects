package main

type Fish struct {
}

type FakeFish Fish

func (f Fish) Swim() {
	println("test1")

}

func (f FakeFish) FakeSwim() {
	println("test2")
}
func UserFish() {
	f1 := Fish{}
	f1.Swim()

	f2 := FakeFish{}
	f2.FakeSwim()

	f1 = Fish(f2)
	f2 = FakeFish(f1)
	println(f1, f2)
}

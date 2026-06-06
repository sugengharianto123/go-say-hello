package go_say_hello

func SayHello() string {
	return "Hello !"
}

func SayLoveYou(name string) string {
	return "I love you " + name + "!"
}

func SayHappyBirthday(name string, age int) string {
	return "Happy Birthday " + name + "! You are now " + string(age) + "."
}

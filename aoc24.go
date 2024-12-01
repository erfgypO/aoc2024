package aoc24

import "reflect"

type Aoc24 struct{}

func Run(key string) {
	reflect.ValueOf(&Aoc24{}).MethodByName(key).Call([]reflect.Value{})
}

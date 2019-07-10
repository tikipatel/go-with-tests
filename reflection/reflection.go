package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
	// fn("i still can't belive South Korea beat Germany 2-0 to to put them last in their group.")
}

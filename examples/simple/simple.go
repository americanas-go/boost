package simple

// ExampleStruct Lorem ipsum dolor sit amet, consectetur adipiscing elit
// @B App xpto
// @B Type FUNC
// @B HandlerType HTTP
type ExampleStruct struct {
}

func (t *ExampleStruct) FooStructMethod() {
}

// FooMethod Lorem ipsum dolor sit amet, consectetur adipiscing elit
// @B App xpto
// @B Type FUNC
// @B HandlerType HTTP
// @B Paths /,/foo
// @B Method POST
// @B Consumes application/json,application/yaml
// @B Produces application/json
// @B QueryParams name=foo,type=bool&name=bar,type=string
// @B PathParams name=foo,type=string
// @B HeaderParams name=foo,type=string&bar
// @B Body github.com/americanas-go/boost/examples/multi.Request
// @B Response 201
func FooMethod() {
}

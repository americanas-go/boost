//go:generate go-enum -f=$GOFILE --marshal
package main

// ENUM(HTTP,FUNCTION,GRPC)
type AppType int

type Annotations map[string][]string

func (m *Annotations) ToSpec() (spec Spec, err error) {
	for k, as := range *m {

	}
}

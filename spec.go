//go:generate go-enum -f=$GOFILE --marshal
package main

// ENUM(HTTP,FUNCTION,GRPC)
type AppType int

// ENUM(INTERFACE,FUNCTION)
type Type int

type AppSpec interface {
	SetFromAnnotations(string, []Annotation)
}

type GRPCSpec struct {
	Name        string
	Description string
	Options     struct {
		GoPackage   string
		JavaPackage string
	}
	Methods []struct {
		Name string
		// Type    string (UNARY, STREAM)
		Handler         string
		Package         string
		RelativePackage string
		Type            Type
		Message         struct {
			Input  string
			Output string
		}
	}
}

func (s *GRPCSpec) SetFromAnnotations(s2 string, an []Annotation) {}

type FunctionSpec struct {
	Name            string
	Description     string
	Handler         string
	Type            Type
	Package         string
	RelativePackage string
}

func (s *FunctionSpec) SetFromAnnotations(s2 string, an []Annotation) {}

type HTTPSpec struct {
	Name        string
	Description string
	Endpoints   []EndpointHTTPSpec
}

func (s *HTTPSpec) SetFromAnnotations(s2 string, an []Annotation) {}

type EndpointHTTPSpec struct {
	Paths           []string
	Method          string
	Package         string
	RelativePackage string
	Description     string
	Handler         string
	Type            Type
	Consumes        []string
	Produces        []string
	Parameters      []ParameterHTTPSpec
	Responses       map[int]ResponseHTTPSpec
}

type ResponseHTTPSpec struct {
	Description string
	Schema      string
}

type ParameterHTTPSpec struct {
	Name        string
	Description string
	Source      string
	Type        string
	Required    bool
	Schema      string
	Validations struct{}
}

type Spec struct {
	Apps []*AppsSpec
}

type AppsSpec struct {
	Type AppType
	Spec AppSpec
}

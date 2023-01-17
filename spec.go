package main

type Spec struct {
	Name        string
	Description string
	Apps        []struct {
		Type string
		/*
			Spec struct {
				Name        string
				Description string
				Options     struct {
					GoPackage   string
					JavaPackage string
				}
				Methods []struct {
					Name    string
					Type    string
					Handler string
					Message struct {
						Input  string
						Output string
					}
				}
			}
			Spec struct {
				Name        string
				Description string
				Handler     string
			}
		*/
		Spec struct {
			Name        string
			Description string
			Endpoints   []struct {
				Path        string
				Method      string
				Description string
				Handler     string
				Consumes    []string
				Produces    []string
				Parameters  []struct {
					Name        string
					Description string
					Source      string
					Type        string
					Required    bool
					Schema      string
					Validations struct {
					}
				}
				Responses []struct {
					Num201 struct {
						Description string
						Schema      string
					}
				}
			}
		}
	}
}

func New() {

}

package resources

type Resource interface {
	GetResourceType() string
}

type BaseResource struct {
	ID         string       `firestore:"id,omitempty" json:"id,omitempty"`
	Identifier []Identifier `firestore:"identifier,omitempty" json:"identifier"`
}

type Code interface {
	GetCodes() []Code
	String() string
}

func UnmarshalCode(c Code, raw string) Code {
	for _, code := range c.GetCodes() {
		if raw == code.String() {
			return code
		}
	}

	return nil
}

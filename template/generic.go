package template

const TemplateTypeGeneric TemplateType = "generic"

type GenericTemplate struct {
	// Title is limited to 45 characters
	Title    string `json:"title" yaml:"title"`
	ItemURL  string `json:"item_url,omitempty"`
	ImageURL string `json:"image_url,omitempty" yaml:"image"`
	// Subtitle is limited to 80 characters
	Subtitle string   `json:"subtitle,omitempty" yaml:"text,omitempty"`
	Buttons  []Button `json:"buttons,omitempty" yaml:"buttons,omitempty"`
}

func (GenericTemplate) Type() TemplateType {
	return TemplateTypeGeneric
}

func (GenericTemplate) SupportsButtons() bool {
	return true
}

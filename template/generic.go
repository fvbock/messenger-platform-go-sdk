package template

import "fmt"

const TemplateTypeGeneric TemplateType = "generic"

type GenericTemplate struct {
	// Title is limited to 45 characters
	Title    string `json:"title" yaml:"title,omitempty"`
	ItemURL  string `json:"item_url,omitempty"`
	ImageURL string `json:"image_url,omitempty" yaml:"image,omitempty"`
	// Subtitle is limited to 80 characters
	Subtitle string   `json:"subtitle,omitempty" yaml:"text,omitempty"`
	Buttons  []Button `json:"buttons,omitempty" yaml:"buttons,omitempty"`
}

func (t *GenericTemplate) String() string {
	return fmt.Sprintf(`template:
	Title:    %v
	ItemURL:  %v
	ImageURL: %v
	Text:     %v
	Buttons:  %v
`, t.Title, t.ItemURL, t.ImageURL, t.Subtitle, t.Buttons)
}

func (GenericTemplate) Type() TemplateType {
	return TemplateTypeGeneric
}

func (GenericTemplate) SupportsButtons() bool {
	return true
}

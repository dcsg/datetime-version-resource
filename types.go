package resource

type Source struct {
	Timezone string `json:"timezone,omitempty"`
	Format   string `json:"format,omitempty"`
}

type Version struct {
	Version string `json:"version"`
}

type Metadata struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CheckRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

type InRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

type InResponse struct {
	Version Version `json:"version"`
	Metadata []Metadata `json:"metadata"`
}

type OutRequest struct {
	Source  Source  `json:"source"`
}

type OutResponse struct {
	Version  Version  `json:"version"`
	Metadata []Metadata `json:"metadata"`
}

var SourceDefaults = Source{
	Timezone: "UTC",
	Format: "20060102-150405",
}

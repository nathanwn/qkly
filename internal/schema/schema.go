package schema

type InputType string

const (
	StdinInput = "stdin"
	FileInput  = "file"
	RegexInput = "regex"
)

type InputConfiguration struct {
	Type     InputType `json:"type"`
	FileName string    `json:"fileName"`
	Pattern  string    `json:"pattern"`
}

type OutputType string

const (
	StdoutOutput OutputType = "stdout"
	FileOutput              = "file"
)

type OutputConfiguration struct {
	Type     OutputType `json:"type"`
	FileName string     `json:"fileName"`
}

type TestFileType string

const (
	TestFileTypeInput  TestFileType = "input"
	TestFileTypeOutput TestFileType = "output"
)

type TestData struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type TaskData struct {
	Name         string              `json:"name"`
	Group        string              `json:"group"`
	Url          string              `json:"url"`
	Interactive  bool                `json:"interactive"`
	MemoryLimit  int                 `json:"memoryLimit"`
	TimeLimit    int                 `json:"timeLimit"`
	Tests        []TestData          `json:"tests"`
	TestType     string              `json:"testType"`
	InputConfig  InputConfiguration  `json:"input"`
	OutputConfig OutputConfiguration `json:"output"`
}

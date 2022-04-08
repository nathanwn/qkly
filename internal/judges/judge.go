package judges

type Judge interface {
	Name() string
	ContestAndTaskId(url string) (string, string)
}

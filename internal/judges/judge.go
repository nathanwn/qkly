package judges

type Judge interface {
	Id() string
	ContestAndTaskId(url string) (string, string)
}

package gotool

// Commander ...
type Commander interface {
	ExecuteCommand(string, ...string) (string, error)
}

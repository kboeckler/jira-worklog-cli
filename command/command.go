package command

type Command interface {
	Execute(params string) (string, error)
}

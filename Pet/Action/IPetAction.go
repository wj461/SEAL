package Action

type IPetAction interface {
	Update(windowName string) error
}

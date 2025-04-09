package Pet

type IPetAction interface {
	Update(windowName string) error
}

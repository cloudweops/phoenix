package ioc

type IocObject interface {
	// Init initializes the object
	Init() error
	// Name returns the name of the object
	Name() string
	// Priority returns the priority of the object
	Priority() int
}

// IocObjectImpl implements IocObject
type IocObjectImpl struct {
}

// Init initializes the IocObjectImpl object
// It returns an error if something goes wrong
func (i *IocObjectImpl) Init() error {
	// The function body is intentionally left blank
	return nil
}

// Name returns the name of the object
func (i *IocObjectImpl) Name() string {
	return "Nil"
}

// Priority returns the priority of the object
func (i *IocObjectImpl) Priority() int {
	return 0
}

package ioc

// Binder A fluent interface for creating descriptors
type Binder interface {
	// Bind binds the given name to the structure.  prototype can be the struct to be created
	// or can be a pointer to the struct to be created.  prototype is NOT the structure
	// that will be used as a service by Dargo.  If prototype implements DargoInitializer
	// then the DargoInitialize method will be called on it prior to being given to other
	// services
	Bind(name string, prototype any) Binder
	// BindWithCreator binds the given name to a creation function
	BindWithCreator(name string, bindMethod func() (any, error)) Binder
	// BindConstant binds the exact constant as-is into the ServiceLocator
	BindConstant(name string, constant any) Binder
}

type MyBinder struct {
	BindMap            map[string]any
	BindWithCreatorMap map[string]func() (any, error)
	BindConstantMap    map[string]any
}

var _ Binder = (*MyBinder)(nil)

func NewMyBinder() *MyBinder {
	var mbinder = &MyBinder{
		BindMap:            make(map[string]any),
		BindWithCreatorMap: make(map[string]func() (any, error)),
		BindConstantMap:    make(map[string]any),
	}
	return mbinder
}

func (b *MyBinder) Bind(name string, prototype any) Binder {
	b.BindMap[name] = prototype
	return b
}
func (b *MyBinder) BindWithCreator(name string, bindMethod func() (any, error)) Binder {
	b.BindConstantMap[name] = bindMethod
	return b
}
func (b *MyBinder) BindConstant(name string, constant any) Binder {
	b.BindConstantMap[name] = constant
	return b
}

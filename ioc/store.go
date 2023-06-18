package ioc

import (
	"fmt"
	"sort"
)

var store = NewDefaultStore()

const DefaultNamespace = "default"

// InitIocObject initializes the default store for dependency injection
// Returns an error if the initialization fails
func InitIocObject() error {
	return store.InitIocObject()
}

// RegistryObject registers an IocObject to the default namespace in the store.
// It is a shortcut function calling RegistryObjectWithNamespace with the default namespace.
//
// obj: the IocObject to be registered.
func RegistryObject(obj IocObject) {
	RegistryObjectWithNamespace(DefaultNamespace, obj)
}

// RegistryObjectWithNamespace registers the given object to a specific namespace in the IOC container.
// It takes in a string representing the namespace and an IocObject interface representing the object to register.
func RegistryObjectWithNamespace(namespace string, obj IocObject) {
	store.Namespace(namespace).Add(obj)
}

// GetObject returns the object with the given name, default namespace is used.
func GetObject(name string) IocObject {
	return GetObjectWithNamespace(DefaultNamespace, name)
}

// GetObjectWithNamespace returns the object with the given namespace and name
func GetObjectWithNamespace(namespace string, name string) IocObject {
	obj := store.Namespace(namespace).Get(name)
	if obj == nil {
		panic(fmt.Sprintf("object %s.%s not found", namespace, name))
	}
	return obj
}

type DefaultStore struct {
	store map[string]*IocObjectSet
}

// NewDefaultStore creates a new DefaultStore
func NewDefaultStore() *DefaultStore {
	return &DefaultStore{
		store: make(map[string]*IocObjectSet),
	}
}

// InitIocObject initializes all managed objects in the store.
// Returns an error if initialization of any object fails.
func (s *DefaultStore) InitIocObject() error {
	// Iterate over all namespaces and the objects within them.
	for ns, objects := range s.store {
		// Sort the objects within each namespace.
		objects.Sort()

		// Iterate over each object in the namespace and initialize it.
		for _, obj := range objects.Items {
			if err := obj.Init(); err != nil {
				// If initialization of an object fails, return an error.
				return fmt.Errorf("failed to init object %s.%s: %v", ns, obj.Name(), err)
			}
		}
	}

	// Return nil if all objects were successfully initialized.
	return nil
}

// Namespace returns the IocObjectSet for the given namespace.
func (s *DefaultStore) Namespace(namespace string) *IocObjectSet {
	if _, ok := s.store[namespace]; !ok {
		s.store[namespace] = NewIocObjectSet()
	}
	return s.store[namespace]
}

// ShowRegisteredObjectNames returns the names of all managed objects in the store.
func (s *DefaultStore) ShowRegisteredObjectNames() (names []string) {
	for ns, objSet := range s.store {
		for _, obj := range objSet.Items {
			names = append(names, fmt.Sprintf("%s.%s", ns, obj.Name()))
		}
	}
	return
}

// IocObjectSet is a set of IocObject
type IocObjectSet struct {
	Items []IocObject
}

// NewIocObjectSet creates a new IocObjectSet
func NewIocObjectSet() *IocObjectSet {
	return &IocObjectSet{
		Items: make([]IocObject, 0),
	}
}

// Add adds an IocObject to the IocObjectSet
func (s *IocObjectSet) Add(obj IocObject) {
	if s.Exist(obj.Name()) {
		panic(fmt.Sprintf("ioc obj %s has been registered", obj.Name()))
	}
	s.Items = append(s.Items, obj)
}

// Exist checks if the IocObjectSet has the given name.
func (s *IocObjectSet) Exist(name string) bool {
	obj := s.Get(name)
	return obj != nil
}

// Get returns the IocObjectSet object with the given name.
func (s *IocObjectSet) Get(name string) IocObject {
	// Iterate over all items in the set and return the one with the matching name.
	for _, obj := range s.Items {
		if obj.Name() == name {
			return obj
		}
	}
	// Return nil if no object with the given name was found.
	return nil
}

// Len returns the length of the IocObjectSet
func (s *IocObjectSet) Len() int {
	return len(s.Items)
}

// Less compares two elements of the IocObjectSet
func (s *IocObjectSet) Less(i, j int) bool {
	return s.Items[i].Priority() > s.Items[j].Priority()
}

// Swap swaps two elements of the IocObjectSet
func (s *IocObjectSet) Swap(i, j int) {
	s.Items[i], s.Items[j] = s.Items[j], s.Items[i]
}

// Sort function sorts the IocObjectSet by object priority.
// It uses one call to data.Len to determine n and O(n*log(n)) calls to data.Less and data.Swap for sorting.
// The sort is not guaranteed to be stable.
func (s *IocObjectSet) Sort() {
	sort.Sort(s)
}

// ObjectNames returns the names of all registered IocObject
func (s *IocObjectSet) ObjectNames() (names []string) {
	for _, obj := range s.Items {
		names = append(names, obj.Name())
	}
	return
}

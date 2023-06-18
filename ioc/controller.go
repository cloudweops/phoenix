package ioc

import "google.golang.org/grpc"

const ControllerNamespace = "controller"

type GRPCController interface {
	Registry(*grpc.Server)
}

// RegistryController Registers an IocObject to the controller namespace.
// obj: The IocObject to be registered.
func RegistryController(obj IocObject) {
	RegistryObjectWithNamespace(ControllerNamespace, obj)
}

// GetController Returns an IocObject from the controller namespace with the given name.
func GetController(name string) IocObject {
	return GetObjectWithNamespace(ControllerNamespace, name)
}

// ListControllerObjectNames Returns a list of all IocObject names in the controller namespace.
func ListControllerObjectNames() (names []string) {
	return store.Namespace(ControllerNamespace).ObjectNames()
}

// LoadGrpcController registers all objects to the gRPC server.
// It iterates over all objects in the controller namespace of the IOC container
// and attempts to cast them to the GRPCController interface type.
// If the object is a GRPCController, it is registered using the grpc server.
func LoadGrpcController(server *grpc.Server) {
	// Get all objects in the controller namespace
	objects := store.Namespace(ControllerNamespace)

	// Iterate over all objects to register them to the grpc server
	for _, obj := range objects.Items {
		// Attempt to cast the object to GRPCController interface type
		c, ok := obj.(GRPCController)
		if !ok {
			// If the object is not a GRPCController, skip it
			continue
		}
		// Register the GRPCController to the grpc server
		c.Registry(server)
	}
}

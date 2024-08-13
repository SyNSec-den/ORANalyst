package condition_collection

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/packages"
)

type typeTracker struct {
	trackedTypes map[types.Type]bool
}

func getField(t types.Type, i int) (types.Type, string, error) {
	switch t := t.(type) {
	case *types.Basic:
		return t, "", fmt.Errorf("cannot get field of basic type: %v", t)
	case *types.Array:
		return t.Elem(), "", fmt.Errorf("cannot get field of array type: %v", t)
	case *types.Slice:
		return t.Elem(), "", fmt.Errorf("cannot get field of slice type: %v", t)
	case *types.Struct:
		return t.Field(i).Type(), t.Field(i).Name(), nil
	case *types.Pointer:
		return getField(t.Elem(), i)
	case *types.Interface:
		return t, "", fmt.Errorf("cannot get field of interface type: %v", t)
	case *types.Map:
		return t.Elem(), "", fmt.Errorf("cannot get field of map type: %v", t)
	case *types.Chan:
		return t.Elem(), "", fmt.Errorf("cannot get field of channel type: %v", t)
	case *types.Named:
		underlyingType, fieldName, err := getField(t.Underlying(), i)
		if _, ok := t.Underlying().(*types.Struct); ok {
			fieldName = t.Obj().Name() + "." + fieldName // Prepend struct name
		}
		return underlyingType, fieldName, err
	case *types.Tuple:
		return t.At(i).Type(), "", nil
	case *types.Signature:
		return t, "", fmt.Errorf("cannot get field of function type: %v", t)
	case *types.Union:
		return t, "", fmt.Errorf("cannot get field of union type: %v", t)
	default:
		return t, "", fmt.Errorf("unknown type: %v", t)
	}
}

func LookUpTypeFromPackage(pkg *packages.Package, typeName string) (types.Type, error) {
	if pkg == nil || pkg.Errors != nil {
		return nil, fmt.Errorf("package not found")
	}

	obj := pkg.Types.Scope().Lookup(typeName)
	if obj == nil {
		return nil, fmt.Errorf("type not found")
	}
	pointerType := types.NewPointer(obj.Type())
	return pointerType, nil
}

func NewTypeTracker(typs []types.Type, pkg packages.Package) *typeTracker {
	tt := &typeTracker{
		trackedTypes: make(map[types.Type]bool),
	}
	for _, t := range typs {
		tt.buildTypeMap(t, pkg)
	}
	return tt
}

func (tt *typeTracker) buildTypeMap(t types.Type, pkg packages.Package) {
	if _, ok := tt.trackedTypes[t]; ok {
		return
	}

	switch t := t.(type) {
	case *types.Basic:
		return
	case *types.Array:
		tt.buildTypeMap(t.Elem(), pkg)
	case *types.Slice:
		tt.buildTypeMap(t.Elem(), pkg)
	case *types.Struct:
		tt.trackedTypes[t] = true
		for i := 0; i < t.NumFields(); i++ {
			tt.buildTypeMap(t.Field(i).Type(), pkg)
		}
	case *types.Pointer:
		tt.trackedTypes[t] = true
		tt.buildTypeMap(t.Elem(), pkg)
	case *types.Interface:
		tt.trackedTypes[t] = true
		// find whether all types in the package that implements the interface
		scope := pkg.Types.Scope()
		for _, name := range scope.Names() {
			obj := scope.Lookup(name)
			named, ok := obj.Type().(*types.Named)
			if !ok {
				continue
			}
			if types.Implements(named, t) || types.Implements(named.Underlying(), t) {
				tt.buildTypeMap(named, pkg)
			}
		}
	case *types.Map:
		tt.buildTypeMap(t.Elem(), pkg)
	case *types.Chan:
		tt.buildTypeMap(t.Elem(), pkg)
	case *types.Named:
		tt.trackedTypes[t] = true
		tt.buildTypeMap(t.Underlying(), pkg)
	case *types.Tuple:
		tt.trackedTypes[t] = true
		for i := 0; i < t.Len(); i++ {
			tt.buildTypeMap(t.At(i).Type(), pkg)
		}
	case *types.Signature:
		return
	case *types.Union:
		return
	default:
		panic(fmt.Sprintf("unknown type: %v", t))
	}
}

func (tt *typeTracker) isTrackedType(t types.Type) bool {
	res := tt.isTrackedTypeImpl(t)
	fmt.Printf("isTrackedType type: %v - %v, is tracked: %v\n", t, describeType(t), res)
	return res
}

func (tt *typeTracker) isTrackedTypeImpl(t types.Type) bool {
	if track, ok := tt.trackedTypes[t]; ok {
		return track
	}

	switch t := t.(type) {
	case *types.Basic:
		return false
	case *types.Array:
		return tt.isTrackedType(t.Elem())
	case *types.Slice:
		return tt.isTrackedType(t.Elem())
	case *types.Struct:
		return false
	case *types.Pointer:
		return tt.isTrackedType(t.Elem())
	case *types.Interface:
		return false
	case *types.Map:
		return tt.isTrackedType(t.Elem())
	case *types.Chan:
		return tt.isTrackedType(t.Elem())
	case *types.Named:
		return tt.isTrackedType(t.Underlying())
	case *types.Tuple:
		return false
	case *types.Signature:
		if t.Recv() == nil {
			return false
		}
		return tt.isTrackedType(t.Recv().Type())
	case *types.Union:
		return false
	default:
		panic(fmt.Sprintf("unknown type: %v", t))
	}
}

func describeType(t types.Type) string {
	switch t := t.(type) {
	case *types.Basic:
		// It's a basic type (int, float, string, etc.)
		return "Basic type: " + t.Name()
	case *types.Array:
		// It's an array type
		return "Array type"
	case *types.Slice:
		// It's a slice type
		return "Slice type"
	case *types.Struct:
		// It's a struct type
		return "Struct type"
	case *types.Pointer:
		// It's a pointer type
		return "Pointer type, points to: " + describeType(t.Elem())
	case *types.Interface:
		// It's an interface type
		return "Interface type"
	case *types.Map:
		// It's a map type
		return "Map type"
	case *types.Chan:
		// It's a channel type
		return "Channel type"
	case *types.Named:
		// It's a named type (custom type like a typedef or a struct)
		return "Named type: " + t.Obj().Name()
	case *types.Tuple:
		// It's a tuple (used in multi-return functions)
		return "Tuple type"
	case *types.Signature:
		// It's a function signature (used for functions)
		return "Function type"
	case *types.Union:
		// It's a union type (introduced in Go 1.18)
		return "Union type"
	default:
		// Unknown or a type we haven't checked for
		return "Unknown type"
	}
}

package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

var decoderInterface = reflect.TypeOf(new(Decoder)).Elem()

type demoPtrImplements [8]byte

func (self *demoPtrImplements) DecodeRLP(*string) error {
	return nil
}

type Decoder interface {
	DecodeRLP(*string) error
}

func TestPtrImplements(t *testing.T) {
	demoSlice := []byte{0x01, 0x02}
	// output: slice
	fmt.Println("kind = ", reflect.ValueOf(demoSlice).Kind())

	var demoArray [20]byte
	// output: array
	fmt.Println("kind = ", reflect.ValueOf(demoArray).Kind())

	b := demoPtrImplements{0x01, 0x02, 0x03}
	// output: array
	fmt.Println("kind = ", reflect.ValueOf(b).Kind())

	typeB := reflect.TypeOf(b)
	// output: uint8
	fmt.Println("===== Elem of type b is ", typeB.Elem())
	// output: array
	fmt.Println("===== kind of type b is ", typeB.Kind())
	// output: *demoPtrImplements
	fmt.Println("===== ptr of type b is ", reflect.PtrTo(typeB))
	// output: true
	fmt.Println("===== implement decorderInterface?", reflect.PtrTo(typeB).Implements(decoderInterface))

	r := reflect.ValueOf(b)
	rtype := r.Type()

	eType := rtype.Elem()
	// output: uint8
	fmt.Println("===== ptr of value b is ", eType)
	// output: *uint8
	fmt.Println("===== ptr of value b is ", reflect.PtrTo(eType))
	// output: false
	fmt.Println("===== implement decorderInterface?", reflect.PtrTo(eType).Implements(decoderInterface))
}

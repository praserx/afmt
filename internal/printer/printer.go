package printer

import (
	"fmt"
	"reflect"
	// "github.com/kr/pretty"
)

type Printer struct {
	ignoreUnsupported bool
}

func NewPrinter(opts ...Option) *Printer {
	// Default values
	var options = &Options{
		IgnoreUnsupported: true,
	}

	for _, opt := range opts {
		opt(options)
	}

	return &Printer{
		ignoreUnsupported: options.IgnoreUnsupported,
	}

	return nil
}

func (p *Printer) printStruct(inspect interface{}) error {

	if reflect.TypeOf(inspect).Kind() == reflect.Struct {

		val := reflect.ValueOf(inspect)

		for i := 0; i < val.NumField(); i++ {
			switch val.Field(i).Kind() {
			case reflect.Int:
				p.printNumber(val.Field(i))
			default:
				if p.ignoreUnsupported == false {
					return fmt.Errorf("Type %s", reflect.TypeOf(val.Field(i)))
				}
			}
		}

		return nil
	}

	return fmt.Errorf("Value is not a structure")

	// if reflect.ValueOf(inspect).Kind() == reflect.Struct {

	// 	t := reflect.TypeOf(inspect).Name()
	// 	query := fmt.Sprintf("insert into %s values(", t)
	// 	v := reflect.ValueOf(inspect)
	// 	for i := 0; i < v.NumField(); i++ {
	// 		switch v.Field(i).Kind() {
	// 		case reflect.Int:
	// 			if i == 0 {
	// 				query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
	// 			} else {
	// 				query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
	// 			}
	// 		case reflect.String:
	// 			if i == 0 {
	// 				query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
	// 			} else {
	// 				query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
	// 			}
	// 		case reflect.Struct:
	// 			fmt.Println("struct")
	// 		default:
	// 			fmt.Println("Unsupported type")
	// 			return
	// 		}
	// 	}
	// 	query = fmt.Sprintf("%s)", query)
	// 	fmt.Println(query)
	// 	return

	// }
	// fmt.Println("unsupported type")

	// pretty.Println(inspect)
	// pretty.Println(reflect.TypeOf(inspect))
	// pretty.Println(reflect.ValueOf(inspect))

	// fmt.Println(reflect.TypeOf(inspect))
	// fmt.Println(reflect.ValueOf(inspect))

	// fmt.Println(reflect.ValueOf(inspect))
}

func (p *Printer) printArray(inspect interface{}) {

}

func (p *Printer) printString(inspect interface{}) {

}

func (p *Printer) printNumber(inspect interface{}) {

}

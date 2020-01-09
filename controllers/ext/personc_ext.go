package ext

import (
	// "fmt"
	"net/http"
	// "reflect"
)

//===================================================================================================
// base Person entity controller extension-point code generated on 09 Jan 20 16:53 CST
//===================================================================================================

// CtrlPersonCreateExt provides access to the ControllerCreateExt extension-point interface
type CtrlPersonCreateExt struct {
	ControllerCreateExt
}

// CtrlPersonGetExt provides access to the ControllerGetExt extension-point interface
type CtrlPersonGetExt struct {
	ControllerGetExt
}

// CtrlPersonUpdateExt provides access to the ControllerUpdateExt extension-point interface
type CtrlPersonUpdateExt struct {
	ControllerUpdateExt
}

// PersonCtrlExt provides access to the Person implementations of the following interfaces:
//   CtrlCreateExt
//   CtrlUpdateExt
//   CtrlGetExt
type PersonCtrlExt struct {
	CrtEp CtrlPersonCreateExt
	UpdEp CtrlPersonUpdateExt
	GetEp CtrlPersonGetExt
}

var personCtrlExp PersonCtrlExt

// InitPersonCtrlExt initializes the person entity's controller
// extension-point interface implementations.
func InitPersonCtrlExt() *PersonCtrlExt {
	personCtrlExp = PersonCtrlExt{}
	return &personCtrlExp
}

//------------------------------------------------------------------------------------------
// ControllerCreateExt extension-point interface implementation for entity Person
//------------------------------------------------------------------------------------------

// BeforeFirst extension-point implementation for entity Person
// TODO: implement checks and document them here
func (crtEP *CtrlPersonCreateExt) BeforeFirst(w http.ResponseWriter, r *http.Request) error {

	return nil
}

// AfterBodyDecode extension-point implementation for entity Person
// TODO: implement application logic and document it here
func (crtEP *CtrlPersonCreateExt) AfterBodyDecode(ent interface{}) error {

	// fmt.Println("TypeOf ent:", reflect.TypeOf(ent))
	// fmt.Println("ValueOf ent:", reflect.ValueOf(ent))
	// p := ent.(*models.Person)
	//
	// make changes / validate the content struct pointer (p) here
	// p.<field_name> = "A new value"

	return nil
}

// BeforeResponse extension-point implementation for entity Person
// TODO: implement application logic and document it here
func (crtEP *CtrlPersonCreateExt) BeforeResponse(ent interface{}) error {

	// fmt.Println("TypeOf ent:", reflect.TypeOf(ent))
	// fmt.Println("ValueOf ent:", reflect.ValueOf(ent))
	// p := ent.(*models.Person)
	//
	// make changes / validate the content struct pointer (p) here
	// p.<field_name> = p.<field_name> + "."

	return nil
}

//------------------------------------------------------------------------------------------
// ControllerUpdateExt extension-point interface implementation for entity Person
//------------------------------------------------------------------------------------------

// BeforeFirst extension-point implementation for entity Person
// TODO: implement checks and document them here
func (updEP *CtrlPersonUpdateExt) BeforeFirst(w http.ResponseWriter, r *http.Request) error {

	return nil
}

// AfterBodyDecode extension-point implementation for entity Person
// TODO: implement application logic and document it here
func (updEP *CtrlPersonUpdateExt) AfterBodyDecode(ent interface{}) error {

	// fmt.Println("TypeOf ent:", reflect.TypeOf(ent))
	// fmt.Println("ValueOf ent:", reflect.ValueOf(ent))
	// p := ent.(*models.Person)
	//
	// make changes / validate the content struct pointer (p) here
	// p.<field_name> = "An updated value"
	return nil
}

// BeforeResponse extension-point implementation for entity Person
// TODO: implement application logic and document it here
func (updEP *CtrlPersonUpdateExt) BeforeResponse(ent interface{}) error {

	// fmt.Println("TypeOf ent:", reflect.TypeOf(ent))
	// fmt.Println("ValueOf ent:", reflect.ValueOf(ent))
	// p := ent.(*models.Person)
	//
	// make changes / validate the content struct pointer (p) here
	// p.<field_name> = p.<field_name> + "."

	return nil
}

//------------------------------------------------------------------------------------------
// ControllerGetExt extension-point interface implementation for entity Person
//------------------------------------------------------------------------------------------

// BeforeFirst extension-point implementation for entity Person
// TODO: implement checks and document them here
func (getEP *CtrlPersonGetExt) BeforeFirst(w http.ResponseWriter, r *http.Request) error {

	return nil
}

// BeforeModelCall extension-point implementation for entity Person
// TODO: implement application logic and document it here
func (getEP *CtrlPersonGetExt) BeforeModelCall(ent interface{}) error {

	// fmt.Println("TypeOf ent:", reflect.TypeOf(ent))
	// fmt.Println("ValueOf ent:", reflect.ValueOf(ent))
	// p := ent.(*models.Person)
	//
	// make changes / validate the content struct pointer (p) here
	// p.<field_name> = "A new value"

	return nil
}

// BeforeResponse extension-point implementation for entity Person
// TODO: implement application logic and document it here
func (getEP *CtrlPersonGetExt) BeforeResponse(ent interface{}) error {

	// p := ent.(*models.Person)
	//
	// make changes / validate the content struct pointer (p) here
	// p.<field_name> = p.<field_name> + "."

	return nil
}

package models

import (
	// "fmt"
	// "reflect"
	"jiffy_tests/GeneratedPersonApp/models/ext"
)

//===================================================================================================
// base Person entity model extension-point code generated on 09 Jan 20 16:53 CST
//===================================================================================================

// MdlPersonCreateExt provides access to the ModelCreateExt extension-point interface
type MdlPersonCreateExt struct {
	ext.ModelCreateExt
}

// MdlPersonUpdateExt provides access to the ControllerUpdateExt extension-point interface
type MdlPersonUpdateExt struct {
	ext.ModelUpdateExt
}

// MdlPersonGetExt provides access to the ControllerGetExt extension-point interface
type MdlPersonGetExt struct {
	ext.ModelGetExt
}

// PersonMdlExt provides access to the Person implementations of the following interfaces:
//   MdlCreateExt
//   MdlUpdateExt
//   MdlGetExt
type PersonMdlExt struct {
	CrtEp MdlPersonCreateExt
	UpdEp MdlPersonUpdateExt
	GetEp MdlPersonGetExt
}

var personMdlExp PersonMdlExt

// InitPersonMdlExt initializes the person entity's model
// extension-point interface implementations.
func InitPersonMdlExt() *PersonMdlExt {
	personMdlExp = PersonMdlExt{}
	return &personMdlExp
}

//----------------------------------------------------------------------------
// ModelCreateExt interface implementation for entity Person
//----------------------------------------------------------------------------

// BeforeDB model extension-point implementation for entity Person
// TODO: implement pre-ORM call logic and document it here
func (crtEP *MdlPersonCreateExt) BeforeDB(ent interface{}) error {

	// fmt.Println("TypeOf ent:", reflect.TypeOf(ent))
	// fmt.Println("ValueOf ent:", reflect.ValueOf(ent))
	// p := ent.(*models.Person)

	// make changes / validate the content struct pointer (p) here
	// p.Name = "A new field value"
	return nil
}

// AfterDB model extension-point implementation for entity Person
// TODO: implement post-ORM call logic and document it here
func (crtEP *MdlPersonCreateExt) AfterDB(ent interface{}) error {

	// fmt.Println("TypeOf ent:", reflect.TypeOf(ent))
	// fmt.Println("ValueOf ent:", reflect.ValueOf(ent))
	// p := ent.(*models.Person)

	// make changes / validate the content struct pointer (p) here
	// p.Name = "A new field value"
	return nil
}

//----------------------------------------------------------------------------
// ModelUpdateExt interface implementation for entity Person
//----------------------------------------------------------------------------

// BeforeDB extension-point implementation for entity Person
// TODO: implement pre-ORM call logic and document it here
func (updEP *MdlPersonUpdateExt) BeforeDB(ent interface{}) error {

	// fmt.Println("TypeOf ent:", reflect.TypeOf(ent))
	// fmt.Println("ValueOf ent:", reflect.ValueOf(ent))
	// p := ent.(*models.Person)

	// make changes / validate the content struct pointer (p) here
	// p.Name = "A new field value"
	return nil
}

// AfterDB extension-point implementation for entity Person
// TODO: implement post-ORM call logic and document it here
func (updEP *MdlPersonUpdateExt) AfterDB(ent interface{}) error {

	// fmt.Println("TypeOf ent:", reflect.TypeOf(ent))
	// fmt.Println("ValueOf ent:", reflect.ValueOf(ent))
	// p := ent.(*models.Person)

	// make changes / validate the content struct pointer (p) here
	// p.Name = "A new field value"
	return nil
}

//----------------------------------------------------------------------------
// ModelGetExt interface implementation for entity Person
//----------------------------------------------------------------------------

// BeforeDB extension-point implementation for entity Person
// TODO: implement pre-ORM call logic and document it here
func (getEP *MdlPersonGetExt) BeforeDB(ent interface{}) error {

	// fmt.Println("TypeOf ent:", reflect.TypeOf(ent))
	// fmt.Println("ValueOf ent:", reflect.ValueOf(ent))
	// p := ent.(*models.Person)

	// make changes / validate the content struct pointer (p) here
	// p.Name = "A new field value"
	return nil
}

// AfterDB extension-point implementation for entity Person
// TODO: implement post-ORM call logic and document it here
func (getEP *MdlPersonGetExt) AfterDB(ent interface{}) error {

	// fmt.Println("TypeOf ent:", reflect.TypeOf(ent))
	// fmt.Println("ValueOf ent:", reflect.ValueOf(ent))
	// p := ent.(*Person)

	// make changes / validate the content struct pointer (p) here
	// p.Name = "A new field value"
	return nil
}

package models

//=============================================================================================
// base Person entity model code generated on 09 Jan 20 16:53 CST
//=============================================================================================

import (
	"fmt"

	"github.com/1414C/lw"
	"github.com/1414C/sqac"
)

// Person structure
type Person struct {
	ID           uint64   `json:"id" db:"id" sqac:"primary_key:inc;start:10000000"`
	Href         string   `json:"href" db:"href" sqac:"-"`
	Name         *string  `json:"name,omitempty" db:"name" sqac:"nullable:true;index:non-unique"`
	Age          *uint    `json:"age,omitempty" db:"age" sqac:"nullable:true"`
	Weight       *float64 `json:"weight,omitempty" db:"weight" sqac:"nullable:true"`
	ValidLicense *bool    `json:"valid_license,omitempty" db:"valid_license" sqac:"nullable:true;index:non-unique"`
}

// PersonDB is a CRUD-type interface specifically for dealing with Persons.
type PersonDB interface {
	Create(person *Person) error
	Update(person *Person) error
	Delete(person *Person) error
	Get(person *Person) error
	GetPersons(params []sqac.GetParam, cmdMap map[string]interface{}) ([]Person, uint64) // uint64 holds $count result
	GetPersonsByName(op string, Name string) []Person
	GetPersonsByAge(op string, Age uint) []Person
	GetPersonsByWeight(op string, Weight float64) []Person
	GetPersonsByValidLicense(op string, ValidLicense bool) []Person
}

// personValidator checks and normalizes data prior to
// db access.
type personValidator struct {
	PersonDB
}

// personValFunc type is the prototype for discrete Person normalization
// and validation functions that will be executed by func runPersonValidationFuncs(...)
type personValFunc func(*Person) error

// PersonService is the public interface to the Person entity
type PersonService interface {
	PersonDB
}

// private service for person
type personService struct {
	PersonDB
}

// personSqac is a sqac-based implementation of the PersonDB interface.
type personSqac struct {
	handle sqac.PublicDB
	ep     PersonMdlExt
}

var _ PersonDB = &personSqac{}

// newPersonValidator returns a new personValidator
func newPersonValidator(pdb PersonDB) *personValidator {
	return &personValidator{
		PersonDB: pdb,
	}
}

// runPersonValFuncs executes a list of discrete validation
// functions against a person.
func runPersonValFuncs(person *Person, fns ...personValFunc) error {

	// iterate over the slice of function names and execute
	// each in-turn.  the order in which the lists are made
	// can matter...
	for _, fn := range fns {
		err := fn(person)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewPersonService needs some work:
func NewPersonService(handle sqac.PublicDB) PersonService {

	ps := &personSqac{
		handle: handle,
		ep:     *InitPersonMdlExt(),
	}

	pv := newPersonValidator(ps) // *db
	return &personService{
		PersonDB: pv,
	}
}

// ensure consistency (build error if delta exists)
var _ PersonDB = &personValidator{}

//-------------------------------------------------------------------------------------------------------
// CRUD-type model methods for Person
//-------------------------------------------------------------------------------------------------------
//
// Create validates and normalizes data used in the person creation.
// Create then calls the creation code contained in PersonService.
func (pv *personValidator) Create(person *Person) error {

	// perform normalization and validation -- comment out checks that are not required
	// note that the check calls are generated as a straight enumeration of the entity
	// structure.  It may be neccessary to adjust the calling order depending on the
	// relationships between the fields in the entity structure.
	err := runPersonValFuncs(person,
		pv.normvalName,
		pv.normvalAge,
		pv.normvalWeight,
		pv.normvalValidLicense,
	)

	if err != nil {
		return err
	}
	return pv.PersonDB.Create(person)
}

// Update validates and normalizes the content of the Person
// being updated by way of executing a list of predefined discrete
// checks.  if the checks are successful, the entity is updated
// on the db via the ORM.
func (pv *personValidator) Update(person *Person) error {

	// perform normalization and validation -- comment out checks that are not required
	// note that the check calls are generated as a straight enumeration of the entity
	// structure.  It may be neccessary to adjust the calling order depending on the
	// relationships between the fields in the entity structure.
	err := runPersonValFuncs(person,
		pv.normvalName,
		pv.normvalAge,
		pv.normvalWeight,
		pv.normvalValidLicense,
	)

	if err != nil {
		return err
	}
	return pv.PersonDB.Update(person)
}

// Delete is passed through to the ORM with no real
// validations.  id is checked in the controller.
func (pv *personValidator) Delete(person *Person) error {

	return pv.PersonDB.Delete(person)
}

// Get is passed through to the ORM with no real
// validations.  id is checked in the controller.
func (pv *personValidator) Get(person *Person) error {

	return pv.PersonDB.Get(person)
}

// GetPersons is passed through to the ORM with no validation
func (pv *personValidator) GetPersons(params []sqac.GetParam, cmdMap map[string]interface{}) ([]Person, uint64) {

	return pv.PersonDB.GetPersons(params, cmdMap)
}

//-------------------------------------------------------------------------------------------------------
// internal personValidator funcs
//-------------------------------------------------------------------------------------------------------
// These discrete functions are used to normalize and validate the Entity fields
// from with in the Create and Update methods.  See the comments in the model's
// Create and Update methods for details regarding use.

// normvalName normalizes and validates field Name
func (pv *personValidator) normvalName(person *Person) error {

	// TODO: implement normalization and validation for Name
	return nil
}

// normvalAge normalizes and validates field Age
func (pv *personValidator) normvalAge(person *Person) error {

	// TODO: implement normalization and validation for Age
	return nil
}

// normvalWeight normalizes and validates field Weight
func (pv *personValidator) normvalWeight(person *Person) error {

	// TODO: implement normalization and validation for Weight
	return nil
}

// normvalValidLicense normalizes and validates field ValidLicense
func (pv *personValidator) normvalValidLicense(person *Person) error {

	// TODO: implement normalization and validation for ValidLicense
	return nil
}

//-------------------------------------------------------------------------------------------------------
// internal person Simple Query Validator funcs
//-------------------------------------------------------------------------------------------------------
// Simple query normalization and validation occurs in the controller to an
// extent, as the URL has to be examined closely in order to determine what to call
// in the model.  This section may be blank if no model fields were marked as
// selectable in the <models>.json file.
// GetPersonsByName is passed through to the ORM with no validation.
func (pv *personValidator) GetPersonsByName(op string, name string) []Person {

	// TODO: implement normalization and validation for the GetPersonsByName call.
	// TODO: typically no modifications are required here.
	return pv.PersonDB.GetPersonsByName(op, name)
}

// GetPersonsByAge is passed through to the ORM with no validation.
func (pv *personValidator) GetPersonsByAge(op string, age uint) []Person {

	// TODO: implement normalization and validation for the GetPersonsByAge call.
	// TODO: typically no modifications are required here.
	return pv.PersonDB.GetPersonsByAge(op, age)
}

// GetPersonsByWeight is passed through to the ORM with no validation.
func (pv *personValidator) GetPersonsByWeight(op string, weight float64) []Person {

	// TODO: implement normalization and validation for the GetPersonsByWeight call.
	// TODO: typically no modifications are required here.
	return pv.PersonDB.GetPersonsByWeight(op, weight)
}

// GetPersonsByValidLicense is passed through to the ORM with no validation.
func (pv *personValidator) GetPersonsByValidLicense(op string, valid_license bool) []Person {

	// TODO: implement normalization and validation for the GetPersonsByValidLicense call.
	// TODO: typically no modifications are required here.
	return pv.PersonDB.GetPersonsByValidLicense(op, valid_license)
}

//-------------------------------------------------------------------------------------------------------
// ORM db CRUD access methods
//-------------------------------------------------------------------------------------------------------
//
// Create a new Person in the database via the ORM
func (ps *personSqac) Create(person *Person) error {

	err := ps.ep.CrtEp.BeforeDB(person)
	if err != nil {
		return err
	}
	err = ps.handle.Create(person)
	if err != nil {
		return err
	}
	err = ps.ep.CrtEp.AfterDB(person)
	if err != nil {
		return err
	}
	return err
}

// Update an existng Person in the database via the ORM
func (ps *personSqac) Update(person *Person) error {

	err := ps.ep.UpdEp.BeforeDB(person)
	if err != nil {
		return err
	}
	err = ps.handle.Update(person)
	if err != nil {
		return err
	}
	err = ps.ep.UpdEp.AfterDB(person)
	if err != nil {
		return err
	}
	return err
}

// Delete an existing Person in the database via the ORM
func (ps *personSqac) Delete(person *Person) error {
	return ps.handle.Delete(person)
}

// Get an existing Person from the database via the ORM
func (ps *personSqac) Get(person *Person) error {

	err := ps.ep.GetEp.BeforeDB(person)
	if err != nil {
		return err
	}
	err = ps.handle.GetEntity(person)
	if err != nil {
		return err
	}
	err = ps.ep.GetEp.AfterDB(person)
	if err != nil {
		return err
	}
	return err
}

// Get all existing Persons from the db via the ORM
func (ps *personSqac) GetPersons(params []sqac.GetParam, cmdMap map[string]interface{}) ([]Person, uint64) {

	var err error

	// create a slice to read into
	persons := []Person{}

	// call the ORM
	result, err := ps.handle.GetEntitiesWithCommands(persons, params, cmdMap)
	if err != nil {
		lw.Warning("PersonModel GetPersons() error: %s", err.Error())
		return nil, 0
	}

	// check to see what was returned
	switch result.(type) {
	case []Person:
		persons = result.([]Person)

		// call the extension-point
		for i := range persons {
			err = ps.ep.GetEp.AfterDB(&persons[i])
			if err != nil {
				lw.Warning("PersonModel GetPersons AfterDB() error: %s", err.Error())
			}
		}
		return persons, 0

	case int64:
		return nil, uint64(result.(int64))

	case uint64:
		return nil, result.(uint64)

	default:
		return nil, 0

	}
}

//-------------------------------------------------------------------------------------------------------
// ORM db simple selector access methods
//-------------------------------------------------------------------------------------------------------
//
// Get all existing PersonsByName from the db via the ORM
func (ps *personSqac) GetPersonsByName(op string, Name string) []Person {

	var persons []Person
	var c string

	switch op {
	case "EQ":
		c = "name = ?"
	case "LIKE":
		c = "name like ?"
	default:
		return nil
	}
	qs := fmt.Sprintf("SELECT * FROM person WHERE %s;", c)
	err := ps.handle.Select(&persons, qs, Name)
	if err != nil {
		lw.Warning("GetPersonsByName got: %s", err.Error())
		return nil
	}

	if ps.handle.IsLog() {
		lw.Info("GetPersonsByName found: %v based on (%s %v)", persons, op, Name)
	}

	// call the extension-point
	for i := range persons {
		err = ps.ep.GetEp.AfterDB(&persons[i])
		if err != nil {
			lw.Warning("PersonModel Getpersons AfterDB() error: %s", err.Error())
		}
	}
	return persons
}

// Get all existing PersonsByAge from the db via the ORM
func (ps *personSqac) GetPersonsByAge(op string, Age uint) []Person {

	var persons []Person
	var c string

	switch op {
	case "EQ":
		c = "age = ?"
	case "LT":
		c = "age < ?"
	case "GT":
		c = "age > ?"
	default:
		return nil
	}
	qs := fmt.Sprintf("SELECT * FROM person WHERE %s;", c)
	err := ps.handle.Select(&persons, qs, Age)
	if err != nil {
		lw.Warning("GetPersonsByAge got: %s", err.Error())
		return nil
	}

	if ps.handle.IsLog() {
		lw.Info("GetPersonsByAge found: %v based on (%s %v)", persons, op, Age)
	}

	// call the extension-point
	for i := range persons {
		err = ps.ep.GetEp.AfterDB(&persons[i])
		if err != nil {
			lw.Warning("PersonModel Getpersons AfterDB() error: %s", err.Error())
		}
	}
	return persons
}

// Get all existing PersonsByWeight from the db via the ORM
func (ps *personSqac) GetPersonsByWeight(op string, Weight float64) []Person {

	var persons []Person
	var c string

	switch op {
	case "EQ":
		c = "weight = ?"
	case "LT":
		c = "weight < ?"
	case "LE":
		c = "weight <= ?"
	case "GT":
		c = "weight > ?"
	case "GE":
		c = "weight >= ?"
	default:
		return nil
	}
	qs := fmt.Sprintf("SELECT * FROM person WHERE %s;", c)
	err := ps.handle.Select(&persons, qs, Weight)
	if err != nil {
		lw.Warning("GetPersonsByWeight got: %s", err.Error())
		return nil
	}

	if ps.handle.IsLog() {
		lw.Info("GetPersonsByWeight found: %v based on (%s %v)", persons, op, Weight)
	}

	// call the extension-point
	for i := range persons {
		err = ps.ep.GetEp.AfterDB(&persons[i])
		if err != nil {
			lw.Warning("PersonModel Getpersons AfterDB() error: %s", err.Error())
		}
	}
	return persons
}

// Get all existing PersonsByValidLicense from the db via the ORM
func (ps *personSqac) GetPersonsByValidLicense(op string, ValidLicense bool) []Person {

	var persons []Person
	var c string

	switch op {
	case "EQ":
		c = "valid_license = ?"
	case "NE":
		c = "valid_license != ?"
	default:
		return nil
	}
	qs := fmt.Sprintf("SELECT * FROM person WHERE %s;", c)
	err := ps.handle.Select(&persons, qs, ValidLicense)
	if err != nil {
		lw.Warning("GetPersonsByValidLicense got: %s", err.Error())
		return nil
	}

	if ps.handle.IsLog() {
		lw.Info("GetPersonsByValidLicense found: %v based on (%s %v)", persons, op, ValidLicense)
	}

	// call the extension-point
	for i := range persons {
		err = ps.ep.GetEp.AfterDB(&persons[i])
		if err != nil {
			lw.Warning("PersonModel Getpersons AfterDB() error: %s", err.Error())
		}
	}
	return persons
}

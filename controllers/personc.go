package controllers

//=============================================================================================
// base Person entity controller code generated on 09 Jan 20 16:53 CST
//=============================================================================================

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"jiffy_tests/GeneratedPersonApp/controllers/ext"
	"jiffy_tests/GeneratedPersonApp/models"

	"github.com/1414C/lw"
	"github.com/1414C/sqac"
	"github.com/gorilla/mux"
)

// PersonController is the person controller type for route binding
type PersonController struct {
	ps   models.PersonService
	ep   ext.PersonCtrlExt
	svcs models.Services
}

// NewPersonController creates a new PersonController
func NewPersonController(ps models.PersonService, svcs models.Services) *PersonController {
	return &PersonController{
		ps:   ps,
		ep:   *ext.InitPersonCtrlExt(),
		svcs: svcs,
	}
}

// Create facilitates the creation of a new Person.  This method is bound
// to the gorilla.mux router in main.go.
//
// POST /person
func (pc *PersonController) Create(w http.ResponseWriter, r *http.Request) {

	var err error
	var pm models.Person

	// TODO: implement extension-point if required
	// TODO: safe to comment this block out if the extension-point is not needed
	err = pc.ep.CrtEp.BeforeFirst(w, r)
	if err != nil {
		lw.ErrorWithPrefixString("PersonController CreateBeforeFirst() error:", err)
		respondWithError(w, http.StatusBadRequest, "personc: Invalid request")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pm); err != nil {
		lw.ErrorWithPrefixString("Person Create:", err)
		respondWithError(w, http.StatusBadRequest, "personc: Invalid request payload")
		return
	}
	defer r.Body.Close()

	// TODO: implement extension-point if required
	// TODO: safe to comment this block out if the extension-point is not needed
	err = pc.ep.CrtEp.AfterBodyDecode(&pm)
	if err != nil {
		lw.ErrorWithPrefixString("PersonController CreateAfterBodyDecode() error:", err)
		respondWithError(w, http.StatusBadRequest, "personc: Invalid request payload")
		return
	}

	// fill the model
	person := models.Person{
		Name:         pm.Name,
		Age:          pm.Age,
		Weight:       pm.Weight,
		ValidLicense: pm.ValidLicense,
	}

	// build a base urlString for the JSON Body self-referencing Href tag
	urlString := buildHrefStringFromCRUDReq(r, true)

	// call the Create method on the person model
	err = pc.ps.Create(&person)
	if err != nil {
		lw.ErrorWithPrefixString("Person Create:", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	person.Href = urlString + strconv.FormatUint(uint64(person.ID), 10)

	// TODO: implement extension-point if required
	// TODO: safe to comment this block out if the extension-point is not needed
	err = pc.ep.CrtEp.BeforeResponse(&person)
	if err != nil {
		lw.ErrorWithPrefixString("PersonController CreateBeforeResponse() error:", err)
		respondWithError(w, http.StatusBadRequest, "personc: Invalid request")
		return
	}
	respondWithJSON(w, http.StatusCreated, person)
}

// Update facilitates the update of an existing Person.  This method is bound
// to the gorilla.mux router in main.go.
//
// PUT /person:id
func (pc *PersonController) Update(w http.ResponseWriter, r *http.Request) {

	var err error
	var pm models.Person

	// TODO: implement extension-point if required
	// TODO: safe to comment this block out if the extension-point is not needed
	err = pc.ep.UpdEp.BeforeFirst(w, r)
	if err != nil {
		lw.ErrorWithPrefixString("PersonController UpdateBeforeFirst() error:", err)
		respondWithError(w, http.StatusBadRequest, "personc: Invalid request")
		return
	}

	// get the parameter(s)
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		lw.ErrorWithPrefixString("Person Update:", err)
		respondWithError(w, http.StatusBadRequest, "Invalid person id")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pm); err != nil {
		lw.ErrorWithPrefixString("Person Update:", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// TODO: implement extension-point if required
	// TODO: safe to comment this block out if the extension-point is not needed
	err = pc.ep.UpdEp.AfterBodyDecode(&pm)
	if err != nil {
		lw.ErrorWithPrefixString("PersonController UpdateAfterBodyDecode() error:", err)
		respondWithError(w, http.StatusBadRequest, "personc: Invalid request payload")
		return
	}

	// fill the model
	person := models.Person{
		Name:         pm.Name,
		Age:          pm.Age,
		Weight:       pm.Weight,
		ValidLicense: pm.ValidLicense,
	}

	// build a base urlString for the JSON Body self-referencing Href tag
	urlString := buildHrefStringFromCRUDReq(r, false)
	person.ID = id

	// call the update method on the model
	err = pc.ps.Update(&person)
	if err != nil {
		lw.ErrorWithPrefixString("Person Update:", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	person.Href = urlString

	// TODO: implement extension-point if required
	// TODO: safe to comment this block out if the extension-point is not needed
	err = pc.ep.UpdEp.BeforeResponse(&person)
	if err != nil {
		lw.ErrorWithPrefixString("PersonController UpdateBeforeResponse() error:", err)
		respondWithError(w, http.StatusBadRequest, "personc: Invalid request")
		return
	}
	respondWithJSON(w, http.StatusCreated, person)
}

// Get facilitates the retrieval of an existing Person.  This method is bound
// to the gorilla.mux router in main.go.
//
// GET /person/:id
func (pc *PersonController) Get(w http.ResponseWriter, r *http.Request) {

	var err error

	// TODO: implement extension-point if required
	// TODO: safe to comment this block out if the extension-point is not needed
	err = pc.ep.GetEp.BeforeFirst(w, r)
	if err != nil {
		lw.Warning("PersonController GetBeforeFirst() error: %s", err.Error())
		respondWithError(w, http.StatusBadRequest, "personc: Invalid request")
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		lw.Warning("Person Get: %s", err.Error())
		respondWithError(w, http.StatusBadRequest, "Invalid person ID")
		return
	}

	// build a base urlString for the JSON Body self-referencing Href tag
	urlString := buildHrefStringFromCRUDReq(r, false)

	person := models.Person{
		ID: id,
	}

	// TODO: implement extension-point if required
	// TODO: safe to comment this block out if the extension-point is not needed
	err = pc.ep.GetEp.BeforeModelCall(&person)
	if err != nil {
		lw.Warning("PersonController GetBeforeModelCall() error: %s", err.Error())
		respondWithError(w, http.StatusBadRequest, "personc: Invalid request")
		return
	}

	err = pc.ps.Get(&person)
	if err != nil {
		lw.Warning(err.Error())
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	person.Href = urlString

	// TODO: implement extension-point if required
	// TODO: safe to comment this block out if the extension-point is not needed
	err = pc.ep.GetEp.BeforeResponse(&person)
	if err != nil {
		lw.Warning("PersonController GetBeforeResponse() error: %s", err.Error())
		respondWithError(w, http.StatusBadRequest, "personc: Invalid request")
		return
	}
	respondWithJSON(w, http.StatusCreated, person)
}

// Delete facilitates the deletion of an existing Person.  This method is bound
// to the gorilla.mux router in main.go.
//
// DELETE /person/:id
func (pc *PersonController) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		lw.ErrorWithPrefixString("Person Delete:", err)
		respondWithError(w, http.StatusBadRequest, "Invalid Person ID")
		return
	}

	person := models.Person{
		ID: id,
	}

	err = pc.ps.Delete(&person)
	if err != nil {
		lw.ErrorWithPrefixString("Person Delete:", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithHeader(w, http.StatusAccepted)
}

// getPersonSet is used by all PersonSet queries as a means of injecting parameters
// returns ([]Person, $count, countRequested, error)
func (pc *PersonController) getPersonSet(w http.ResponseWriter, r *http.Request, params []sqac.GetParam) ([]models.Person, uint64, bool, error) {

	var mapCommands map[string]interface{}
	var err error
	var urlString string
	var persons []models.Person
	var count uint64
	countReq := false

	// check for mux.vars
	vars := mux.Vars(r)

	// parse commands ($cmd) if any
	if len(vars) > 0 && vars != nil {
		mapCommands, err = parseRequestCommands(vars)
		if err != nil {
			return nil, 0, false, err
		}
	}

	// $count trumps all other commands
	if mapCommands != nil {
		_, ok := mapCommands["count"]
		if ok {
			countReq = true
		}
		persons, count = pc.ps.GetPersons(params, mapCommands)
	} else {
		persons, count = pc.ps.GetPersons(params, nil)
	}

	// retrieved []Person and not asked to $count
	if persons != nil && countReq == false {
		for i, l := range persons {
			persons[i].Href = urlString + strconv.FormatUint(uint64(l.ID), 10)
		}
		return persons, 0, countReq, nil
	}

	// $count was requested, which trumps all other commands
	if countReq == true {
		return nil, count, countReq, nil
	}

	// fallthrough and return nothing
	return nil, 0, countReq, nil
}

// GetPersons facilitates the retrieval of all existing Persons.  This method is bound
// to the gorilla.mux router in main.go.
//
// GET /persons
// GET /persons/$count | $limit=n $offset=n $orderby=<field_name> ($asc|$desc)
func (pc *PersonController) GetPersons(w http.ResponseWriter, r *http.Request) {

	var persons []models.Person
	var count uint64
	countReq := false

	// build base Href; common for each selected row
	urlString := buildHrefBasic(r, true)

	// call the common getPersonSet method
	persons, count, countReq, err := pc.getPersonSet(w, r, nil)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf(`{"GetPersons": "%s"}`, err))
		return
	}

	// retrieved []Person and not asked to $count
	if persons != nil && countReq == false {
		for i, l := range persons {
			persons[i].Href = urlString + strconv.FormatUint(uint64(l.ID), 10)
		}
		respondWithJSON(w, http.StatusOK, persons)
		return
	}

	// $count was requested, which trumps all other commands
	if countReq == true {
		respondWithCount(w, http.StatusOK, count)
		return
	}

	// fallthrough and return nothing
	respondWithJSON(w, http.StatusOK, "[]")
}

// GetPersonsByName facilitates the retrieval of existing
// Persons based on Name.
// GET /persons/name(OP 'searchString')
// GET /persons/name(OP 'searchString')/$count | $limit=n $offset=n $orderby=<field_name> ($asc|$desc)
func (pc *PersonController) GetPersonsByName(w http.ResponseWriter, r *http.Request) {

	// get the name parameter
	vars := mux.Vars(r)
	searchValue := vars["name"]
	if searchValue == "" {
		respondWithError(w, http.StatusBadRequest, "missing search criteria")
		return
	}

	// adjust operator and predicate if neccessary
	op, predicate, err := buildStringQueryComponents(searchValue)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf(`{"GetPersonsByName": "%s"}`, err))
		return
	}

	// build GetParam
	p := sqac.GetParam{
		FieldName:    "name",
		Operand:      op,
		ParamValue:   predicate,
		NextOperator: "",
	}
	params := []sqac.GetParam{}
	params = append(params, p)

	// build base Href; common for each selected row
	urlString := buildHrefBasic(r, true)

	// call the common Person GetSet method
	persons, count, countReq, err := pc.getPersonSet(w, r, params)
	if persons != nil && countReq == false {
		for i, l := range persons {
			persons[i].Href = urlString + "person/" + strconv.FormatUint(uint64(l.ID), 10)
		}
		respondWithJSON(w, http.StatusOK, persons)
		return
	}

	if countReq == true {
		respondWithCount(w, http.StatusOK, count)
		return
	}
	respondWithJSON(w, http.StatusOK, "[]")
}

// GetPersonsByAge facilitates the retrieval of existing
// Persons based on Age.

// GET /persons/age(OP searchValue)
// GET /persons/age(OP searchValue)/$count | $limit=n $offset=n $orderby=<field_name> ($asc|$desc)
func (pc *PersonController) GetPersonsByAge(w http.ResponseWriter, r *http.Request) {

	// get the age parameter
	vars := mux.Vars(r)
	searchValue := vars["age"]
	if searchValue == "" {
		respondWithError(w, http.StatusBadRequest, "missing search criteria")
		return
	}

	// adjust operator and predicate if neccessary
	op, predicate, err := buildUIntQueryComponent(searchValue)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf(`{"GetPersonsByAge": "%s"}`, err))
		return
	}

	// build GetParam
	p := sqac.GetParam{
		FieldName:    "age",
		Operand:      op,
		ParamValue:   predicate,
		NextOperator: "",
	}
	params := []sqac.GetParam{}
	params = append(params, p)

	// build base Href; common for each selected row
	urlString := buildHrefBasic(r, true)

	// call the common Person GetSet method
	persons, count, countReq, err := pc.getPersonSet(w, r, params)
	if persons != nil && countReq == false {
		for i, l := range persons {
			persons[i].Href = urlString + "person/" + strconv.FormatUint(uint64(l.ID), 10)
		}
		respondWithJSON(w, http.StatusOK, persons)
		return
	}

	if countReq == true {
		respondWithCount(w, http.StatusOK, count)
		return
	}
	respondWithJSON(w, http.StatusOK, "[]")
}

// GetPersonsByWeight facilitates the retrieval of existing
// Persons based on Weight.

// GET /persons/weight(OP searchValue)
// GET /persons/weight(OP searchValue)/$count | $limit=n $offset=n $orderby=<field_name> ($asc|$desc)
func (pc *PersonController) GetPersonsByWeight(w http.ResponseWriter, r *http.Request) {

	// get the weight parameter
	vars := mux.Vars(r)
	searchValue := vars["weight"]
	if searchValue == "" {
		respondWithError(w, http.StatusBadRequest, "missing search criteria")
		return
	}

	// adjust operator and predicate if neccessary
	op, predicate, err := buildFloat64QueryComponent(searchValue)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf(`{"GetPersonsByWeight": "%s"}`, err))
		return
	}

	// build GetParam
	p := sqac.GetParam{
		FieldName:    "weight",
		Operand:      op,
		ParamValue:   predicate,
		NextOperator: "",
	}
	params := []sqac.GetParam{}
	params = append(params, p)

	// build base Href; common for each selected row
	urlString := buildHrefBasic(r, true)

	// call the common Person GetSet method
	persons, count, countReq, err := pc.getPersonSet(w, r, params)
	if persons != nil && countReq == false {
		for i, l := range persons {
			persons[i].Href = urlString + "person/" + strconv.FormatUint(uint64(l.ID), 10)
		}
		respondWithJSON(w, http.StatusOK, persons)
		return
	}

	if countReq == true {
		respondWithCount(w, http.StatusOK, count)
		return
	}
	respondWithJSON(w, http.StatusOK, "[]")
}

// GetPersonsByValidLicense facilitates the retrieval of existing
// Persons based on ValidLicense.

// GET /persons/valid_license(OP searchValue)
// GET /persons/valid_license(OP searchValue)/$count | $limit=n $offset=n $orderby=<field_name> ($asc|$desc)
func (pc *PersonController) GetPersonsByValidLicense(w http.ResponseWriter, r *http.Request) {

	// get the valid_license parameter
	vars := mux.Vars(r)
	searchValue := vars["valid_license"]
	if searchValue == "" {
		respondWithError(w, http.StatusBadRequest, "missing search criteria")
		return
	}

	// adjust operator and predicate if neccessary
	op, predicate, err := buildBoolQueryComponents(searchValue)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf(`{"GetPersonsByValidLicense": "%s"}`, err))
		return
	}

	// build GetParam
	p := sqac.GetParam{
		FieldName:    "valid_license",
		Operand:      op,
		ParamValue:   predicate,
		NextOperator: "",
	}
	params := []sqac.GetParam{}
	params = append(params, p)

	// build base Href; common for each selected row
	urlString := buildHrefBasic(r, true)

	// call the common Person GetSet method
	persons, count, countReq, err := pc.getPersonSet(w, r, params)
	if persons != nil && countReq == false {
		for i, l := range persons {
			persons[i].Href = urlString + "person/" + strconv.FormatUint(uint64(l.ID), 10)
		}
		respondWithJSON(w, http.StatusOK, persons)
		return
	}

	if countReq == true {
		respondWithCount(w, http.StatusOK, count)
		return
	}
	respondWithJSON(w, http.StatusOK, "[]")
}

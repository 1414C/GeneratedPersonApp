package main_test

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"jiffy_tests/GeneratedPersonApp/appobj"
	"jiffy_tests/GeneratedPersonApp/models"
)

// SessionData contains session management vars
type SessionData struct {
	jwtToken     string
	client       *http.Client
	log          bool
	ID           uint64
	baseURL      string
	testURL      string
	testEndPoint string
	usrName      string
	usrID        uint64
}

var (
	sessionData SessionData
	certFile    = flag.String("cert", "mycert1.cer", "A PEM encoded certificate file.")
	keyFile     = flag.String("key", "mycert1.key", "A PEM encoded private key file.")
	caFile      = flag.String("CA", "myCA.cer", "A PEM encoded CA's certificate file.")
)

var a appobj.AppObj

func TestMain(m *testing.M) {

	// parse flags
	logFlag := flag.Bool("log", false, "extended log")
	useHttpsFlag := flag.Bool("https", false, "true == use https")
	addressFlag := flag.String("address", "localhost:3000", "address:port to connect to")
	u := flag.String("u", "admin", "user name")
	passwd := flag.String("passwd", "", "passwd")
	flag.Parse()

	sessionData.log = *logFlag

	// initialize client / transport
	err := sessionData.initializeClient(*useHttpsFlag)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	// build base url
	err = sessionData.buildURL(*useHttpsFlag, *addressFlag)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	// this method was implemented prior to the end-point authorization build
	// and does not presently work.  I think the test should run with a user
	// and password specified from the command line. :(
	// // create test usr
	// err = sessionData.createUsr()
	// if err != nil {
	// 	log.Fatalf("%s\n", err.Error())
	// }

	// login / get jwt
	err = sessionData.getJWT(*u, *passwd)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	code := m.Run()

	// // delete test usr
	// err = sessionData.deleteUsr()
	// if err != nil {
	//	log.Fatalf("%s\n", err.Error())
	//}

	os.Exit(code)

}

// initialize client / transport
func (sd *SessionData) initializeClient(useHttps bool) error {

	// https
	if useHttps {
		// Load client cert
		cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
		if err != nil {
			return err
		}

		// Load CA cert
		caCert, err := ioutil.ReadFile(*caFile)
		if err != nil {
			return err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		// Setup HTTPS client
		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
		}
		tlsConfig.BuildNameToCertificate()
		transport := &http.Transport{TLSClientConfig: tlsConfig}
		sd.client = &http.Client{Transport: transport,
			Timeout: time.Second * 10,
		}
	}
	// http
	sd.client = &http.Client{
		Timeout: time.Second * 10,
	}
	return nil
}

// buildURL builds a url based on flag parameters
//
// internal
func (sd *SessionData) buildURL(useHttps bool, address string) error {

	sd.baseURL = "http"
	if useHttps {
		sd.baseURL = sd.baseURL + "s"
	}
	sd.baseURL = sd.baseURL + "://" + address
	return nil
}

// createUsr creates a test usr for the application
//
// POST - /usr
func (sd *SessionData) createUsr() error {

	url := sd.baseURL + "/usr"

	// create unique usr name
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}
	sessionData.usrName = fmt.Sprintf("%X%X%X%X%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	jsonStr := fmt.Sprintf("{\"email\":\"%s@1414c.io\",\"password\":\"woofwoof\"}", sessionData.usrName)

	// var jsonBody = []byte(`{"email":"testusr123@1414c.io", "password":"woofwoof"}`)
	var jsonBody = []byte(jsonStr)
	fmt.Println("creating usr:", string(jsonBody))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	if sd.log {
		fmt.Println("POST request Headers:", req.Header)
	}

	resp, err := sd.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var usr models.Usr
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&usr); err != nil {
		return err
	}

	sessionData.usrID = usr.ID

	if sd.log {
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
	}
	return nil
}

// deleteUsr deletes the test usr
//
// DELETE - /usr/:id
func (sd *SessionData) deleteUsr() error {

	idStr := fmt.Sprint(sessionData.usrID)
	// url := "https://localhost:8080/usr/" + idStr
	fmt.Println("deleting usr:", sessionData.usrName, sessionData.usrID)
	url := sessionData.baseURL + "/usr/" + idStr
	var jsonBody = []byte(`{}`)
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonBody))
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sessionData.jwtToken)

	if sessionData.log {
		fmt.Println("sessionData.ID:", string(sessionData.ID))
		fmt.Println("DELETE URL:", url)
		fmt.Println("DELETE request Headers:", req.Header)
	}

	resp, err := sessionData.client.Do(req)
	if err != nil {
		fmt.Printf("Test was unable to DELETE /usr/%d. Got %s.\n", sessionData.usrID, err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		fmt.Printf("DELETE /usr{:id} expected http status code of 201 - got %d", resp.StatusCode)
		return err
	}
	return nil
}

// getJWT authenticates and get JWT
//
// POST - /usr/login
func (sd *SessionData) getJWT(u, p string) error {

	type jwtResponse struct {
		Token string `json:"token"`
	}

	// url := "https://localhost:8080/usr/login"
	url := sessionData.baseURL + "/usr/login"

	jsonStr := ""
	if u != "" {
		jsonStr = fmt.Sprintf("{\"email\":\"%s\",\"password\":\"%s\"}", u, p)
	} else {
		jsonStr = fmt.Sprintf("{\"email\":\"%s@1414c.io\",\"password\":\"woofwoof\"}", sessionData.usrName)
	}

	// var jsonStr = []byte(`{"email":"bunnybear10@1414c.io", "password":"woofwoof"}`)
	// jsonStr := fmt.Sprintf("{\"email\":\"%s@1414c.io\",\"password\":\"woofwoof\"}", sessionData.usrName)
	fmt.Println("using usr:", jsonStr)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	if sd.log {
		fmt.Println("POST request Headers:", req.Header)
	}

	resp, err := sd.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var j jwtResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&j); err != nil {
		return err
	}

	sd.jwtToken = j.Token

	if sd.log {
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
	}
	return nil
}

// testSelectableField is used to test the endpoint access to an entity field
// that has been marked as Selectable in the model file.  access will be tested
// for each of the supported operations via multiple calls to this method.
// The selection data provided in the end-point string is representitive of
// the field data-type only, and it is not expected that the string or
// number types will return a data payload in the response body.  Consequently,
// only the http status code in the response is examined.
//
// GET - sd.testURL
func (sd *SessionData) testSelectableField(t *testing.T) {

	var jsonStr = []byte(`{}`)
	req, _ := http.NewRequest("GET", sd.testURL, bytes.NewBuffer(jsonStr))
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sessionData.jwtToken)

	if sessionData.log {
		fmt.Println("GET URL:", sd.testURL)
		fmt.Println("GET request Headers:", req.Header)
	}

	resp, err := sessionData.client.Do(req)
	if err != nil {
		t.Errorf("Test was unable to GET %s. Got %s.\n", sd.testEndPoint, err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("GET %s expected http status code of 200 - got %d", sd.testEndPoint, resp.StatusCode)
	}
}

// TestCreatePerson attempts to create a new Person on the db
//
// POST /person
func TestCreatePerson(t *testing.T) {

	// url := "https://localhost:8080/person"
	url := sessionData.baseURL + "/person"

	var jsonStr = []byte(`{"name":"string_value",
"age":500000,
"weight":1900.99,
"valid_license":true}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sessionData.jwtToken)

	if sessionData.log {
		fmt.Println("POST request Headers:", req.Header)
	}

	resp, err := sessionData.client.Do(req)
	if err != nil {
		t.Errorf("Test was unable to POST /person. Got %s.\n", err.Error())
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("GET /person response Body:", string(body))
		t.Errorf("Test was unable to POST /person. Got %s.\n", err.Error())
	}
	defer resp.Body.Close()

	var e models.Person
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&e); err != nil {
		t.Errorf("Test was unable to decode the result of POST /person. Got %s.\n", err.Error())
	}

	//============================================================================================
	// TODO: implement validation of the created entity here
	//============================================================================================
	if e.Name == nil || *e.Name != "string_value" {
		t.Errorf("inconsistency detected in POST /person field Name.")
	}

	if e.Age == nil || *e.Age != 500000 {
		t.Errorf("inconsistency detected in POST /person field Age.")
	}

	if e.Weight == nil || *e.Weight != 1900.99 {
		t.Errorf("inconsistency detected in POST /person field Weight.")
	}

	if e.ValidLicense == nil || *e.ValidLicense != true {
		t.Errorf("inconsistency detected in POST /person field ValidLicense.")
	}

	if e.ID != 0 {
		sessionData.ID = e.ID
	} else {
		log.Printf("ID value of 0 detected - subsequent test cases will run with ID == 0!")
	}
}

// TestGetPersons attempts to read all persons from the db
//
// GET /persons
func TestGetPersons(t *testing.T) {

	// url := "https://localhost:8080/persons"
	url := sessionData.baseURL + "/persons"
	jsonStr := []byte(`{}`)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sessionData.jwtToken)

	if sessionData.log {
		fmt.Println("GET /persons request Headers:", req.Header)
	}

	// client := &http.Client{}
	resp, err := sessionData.client.Do(req)
	if err != nil {
		t.Errorf("Test was unable to GET /persons. Got %s.\n", err.Error())
	}
	defer resp.Body.Close()

	if sessionData.log {
		fmt.Println("GET /persons response Status:", resp.Status)
		fmt.Println("GET /persons response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("GET /persons response Body:", string(body))
	}
}

// TestGetPerson attempts to read person/{:id} from the db
// using the id created in this entity's TestCreate function.
//
// GET /person/{:id}
func TestGetPerson(t *testing.T) {

	idStr := fmt.Sprint(sessionData.ID)
	// url := "https://localhost:8080/person/" + idStr
	url := sessionData.baseURL + "/person/" + idStr
	jsonStr := []byte(`{}`)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sessionData.jwtToken)

	if sessionData.log {
		fmt.Println("GET /person request Headers:", req.Header)
	}

	// client := &http.Client{}
	resp, err := sessionData.client.Do(req)
	if err != nil {
		t.Errorf("Test was unable to GET /person/%d. Got %s.\n", sessionData.ID, err.Error())
	}
	defer resp.Body.Close()

	if sessionData.log {
		fmt.Println("GET /person response Status:", resp.Status)
		fmt.Println("GET /person response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("GET /person response Body:", string(body))
	}
}

// TestUpdatePerson attempts to update an existing Person on the db
//
// PUT /person/{:id}
func TestUpdatePerson(t *testing.T) {

	idStr := fmt.Sprint(sessionData.ID)
	// url := "https://localhost:8080/person/" + idStr
	url := sessionData.baseURL + "/person/" + idStr

	var jsonStr = []byte(`{"name":"string_update",
"age":999999,
"weight":8888.88,
"valid_license":false}`)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sessionData.jwtToken)

	if sessionData.log {
		fmt.Println("POST request Headers:", req.Header)
	}

	resp, err := sessionData.client.Do(req)
	if err != nil {
		t.Errorf("Test was unable to PUT /person/{:id}. Got %s.\n", err.Error())
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("PUT /person{:id} expected http status code of 201 - got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	var e models.Person
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&e); err != nil {
		t.Errorf("Test was unable to decode the result of PUT /person. Got %s.\n", err.Error())
	}

	//============================================================================================
	// TODO: implement validation of the updated entity here
	//============================================================================================
	if e.Name == nil || *e.Name != "string_update" {
		t.Errorf("inconsistency detected in POST /person field Name.")
	}

	if e.Age == nil || *e.Age != 999999 {
		t.Errorf("inconsistency detected in POST /person field Age.")
	}

	if e.Weight == nil || *e.Weight != 8888.88 {
		t.Errorf("inconsistency detected in POST /person field Weight.")
	}

	if e.ValidLicense == nil || *e.ValidLicense != false {
		t.Errorf("inconsistency detected in POST /person field ValidLicense.")
	}

	if e.ID != 0 {
		sessionData.ID = e.ID
	} else {
		log.Printf("ID value of 0 detected - subsequent test cases will run with ID == 0!")
	}
}

// TestDeletePerson attempts to delete the new Person on the db
//
// DELETE /person/{:id}
func TestDeletePerson(t *testing.T) {

	idStr := fmt.Sprint(sessionData.ID)
	// url := "https://localhost:8080/person/" + idStr
	url := sessionData.baseURL + "/person/" + idStr
	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sessionData.jwtToken)

	if sessionData.log {
		fmt.Println("sessionData.ID:", string(sessionData.ID))
		fmt.Println("DELETE URL:", url)
		fmt.Println("DELETE request Headers:", req.Header)
	}

	resp, err := sessionData.client.Do(req)
	if err != nil {
		t.Errorf("Test was unable to DELETE /person/%d. Got %s.\n", sessionData.ID, err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("DELETE /person{:id} expected http status code of 201 - got %d", resp.StatusCode)
	}
}

func TestGetPersonsByName(t *testing.T) {

	// http://127.0.0.1:<port>/persons/name(OP '<sel_string>')
	sessionData.testEndPoint = "/persons/name(EQ 'test_string')"
	sessionData.testURL = sessionData.baseURL + sessionData.testEndPoint
	sessionData.testSelectableField(t)

	sessionData.testEndPoint = "/persons/name(LIKE 'test_string')"
	sessionData.testURL = sessionData.baseURL + sessionData.testEndPoint
	sessionData.testSelectableField(t)

} // end func {Name string name  false 0  false false false nonUnique EQ,LIKE     sqac:"nullable:true;index:non-unique" json:"name,omitempty" false false false}

func TestGetPersonsByAge(t *testing.T) {

	// http://127.0.0.1:<port>/persons/age(OP XXX)
	sessionData.testEndPoint = "/persons/age(EQ 77)"
	sessionData.testURL = sessionData.baseURL + sessionData.testEndPoint
	sessionData.testSelectableField(t)

	sessionData.testEndPoint = "/persons/age(LT 77)"
	sessionData.testURL = sessionData.baseURL + sessionData.testEndPoint
	sessionData.testSelectableField(t)

	sessionData.testEndPoint = "/persons/age(GT 77)"
	sessionData.testURL = sessionData.baseURL + sessionData.testEndPoint
	sessionData.testSelectableField(t)

} // end func {Age uint age  false 0  false false false  EQ,LT,GT <nil>    sqac:"nullable:true" json:"age,omitempty" false false false}

func TestGetPersonsByWeight(t *testing.T) {

	// http://127.0.0.1:<port>/persons/weight(OP xxx.yyy)
	sessionData.testEndPoint = "/persons/weight(EQ 55.44)"
	sessionData.testURL = sessionData.baseURL + sessionData.testEndPoint
	sessionData.testSelectableField(t)

	sessionData.testEndPoint = "/persons/weight(LT 55.44)"
	sessionData.testURL = sessionData.baseURL + sessionData.testEndPoint
	sessionData.testSelectableField(t)

	sessionData.testEndPoint = "/persons/weight(LE 55.44)"
	sessionData.testURL = sessionData.baseURL + sessionData.testEndPoint
	sessionData.testSelectableField(t)

	sessionData.testEndPoint = "/persons/weight(GT 55.44)"
	sessionData.testURL = sessionData.baseURL + sessionData.testEndPoint
	sessionData.testSelectableField(t)

	sessionData.testEndPoint = "/persons/weight(GE 55.44)"
	sessionData.testURL = sessionData.baseURL + sessionData.testEndPoint
	sessionData.testSelectableField(t)

} // end func {Weight float64 weight  false 0  false false false  EQ,LT,LE,GT,GE <nil>    sqac:"nullable:true" json:"weight,omitempty" false false false}

func TestGetPersonsByValidLicense(t *testing.T) {

	// http://127.0.0.1:<port>/persons/valid_license(OP true|false)
	sessionData.testEndPoint = "/persons/valid_license(EQ true)"
	sessionData.testURL = sessionData.baseURL + sessionData.testEndPoint
	sessionData.testSelectableField(t)

	sessionData.testEndPoint = "/persons/valid_license(EQ false)"
	sessionData.testURL = sessionData.baseURL + sessionData.testEndPoint
	sessionData.testSelectableField(t)

	sessionData.testEndPoint = "/persons/valid_license(NE true)"
	sessionData.testURL = sessionData.baseURL + sessionData.testEndPoint
	sessionData.testSelectableField(t)

	sessionData.testEndPoint = "/persons/valid_license(NE false)"
	sessionData.testURL = sessionData.baseURL + sessionData.testEndPoint
	sessionData.testSelectableField(t)

} // end func {ValidLicense bool valid_license  false 0  false false false nonUnique EQ,NE <nil>    sqac:"nullable:true;index:non-unique" json:"valid_license,omitempty" false false false}

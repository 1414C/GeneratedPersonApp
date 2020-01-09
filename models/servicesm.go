package models

//=============================================================================================
// Services model code generated on 09 Jan 20 16:53 CST
//=============================================================================================

import (
	"github.com/1414C/sqac"
)

// Services contains the set of services used by the application
type Services struct {
	Usr       UsrService
	UsrGroup  UsrGroupService
	Auth      AuthService
	GroupAuth GroupAuthService
	Person    PersonService
	// Product ProductService
	handle sqac.PublicDB
}

// ServicesConfig function type
type ServicesConfig func(*Services) error

// WithSqac provides a function that will return a sqac services closure
func WithSqac(dialect, connectionInfo string, dbLog bool) ServicesConfig {
	return func(s *Services) error {
		handle := sqac.Create(dialect, false, dbLog, connectionInfo)
		s.handle = handle
		return nil
	}
}

// WithLogMode sets the sqac debugging log mode
func WithLogMode(mode bool) ServicesConfig {
	return func(s *Services) error {
		s.handle.Log(mode)
		return nil
	}
}

// WithUsr creates a Usr service
func WithUsr(pepper string) ServicesConfig {
	return func(s *Services) error {
		s.Usr = NewUsrService(s.handle, pepper)
		return nil
	}
}

// WithUsrGroup creates a UsrGroup service
func WithUsrGroup() ServicesConfig {
	return func(s *Services) error {
		s.UsrGroup = NewUsrGroupService(s.handle)
		return nil
	}
}

// WithAuth creates a Auth service
func WithAuth() ServicesConfig {
	return func(s *Services) error {
		s.Auth = NewAuthService(s.handle)
		return nil
	}
}

// WithGroupAuth creates a GroupAuth service
func WithGroupAuth() ServicesConfig {
	return func(s *Services) error {
		s.GroupAuth = NewGroupAuthService(s.handle)
		return nil
	}
}

// WithPerson creates a Person service
func WithPerson() ServicesConfig {
	return func(s *Services) error {
		s.Person = NewPersonService(s.handle)
		return nil
	}
}

// NewServices creates a Services object using the dialect and connectionInfo
// to create a db connection and share it across the set of services
// in the Services object.  ServicesConfig == func(*Services) error
func NewServices(cfgs ...ServicesConfig) (*Services, error) {

	var s Services
	for _, cfg := range cfgs {
		if err := cfg(&s); err != nil {
			return nil, err
		}
	}
	return &s, nil
}

// Close the db connection
func (s *Services) Close() error {
	return s.handle.Close()
}

// DestructiveReset - drop all tables immediately and rebuild them
func (s *Services) DestructiveReset() error {
	return s.handle.DestructiveResetTables(Person{})
}

// AlterAllTables runs AlterTables for each listed entity.  Supports additive columns only.
func (s *Services) AlterAllTables() error {
	return s.handle.AlterTables(Person{}, Usr{}, UsrGroup{}, Auth{}, GroupAuth{})
}

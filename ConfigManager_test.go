package main

import "testing"


func TestNewManagedConfigurationFromFile(t *testing.T) {

	conf := NewManagedConfiguration("testdata/complete.yaml")

	if conf.configuration.Base.Port != 8080 {
		t.Errorf("Application port mismatch from yaml. Expected %v got %v", 8080, conf.configuration.Base.Port)
	}

	if conf.configuration.Base.MgmtPort != 9000 {
		t.Errorf("Management port mismatch from yaml. Expected %v got %v", 9000, conf.configuration.Base.MgmtPort)
	}

	if conf.configuration.Postgres.Port != 5432 {
		t.Errorf("postgres port mismatch. Expected %v got %v", 5432, conf.configuration.Postgres.Port)
	}

}
func TestNewManagedConfigurationWithDefaults(t *testing.T) {
	conf := NewManagedConfiguration()

	if conf.configuration.Base.Port != 8080 {
		t.Errorf("Application port mismatch from yaml. Expected %v got %v", 8080, conf.configuration.Base.Port)
	}

	if conf.configuration.Postgres.Port != 5432 {
		t.Errorf("postgres port mismatch. Expected %v got %v", 5432, conf.configuration.Postgres.Port)
	}
}

func TestNewManagedConfigurationBaseOnly(t *testing.T) {
	conf := NewManagedConfiguration("testdata/base-only.yaml")

	if conf.configuration.Base.Port != 8080 {
		t.Errorf("Application port mismatch from yaml. Expected %v got %v", 8080, conf.configuration.Base.Port)
	}

	if conf.configuration.Base.MgmtPort != 9000 {
		t.Errorf("Management port mismatch from yaml. Expected %v got %v", 9000, conf.configuration.Base.MgmtPort)
	}

	if conf.configuration.Base.BasePath != "/path" {
		t.Errorf("basepath mismatch. Expected %v got %v", "/path", conf.configuration.Base.BasePath)
	}
}
func TestNewManagedConfigurationPostgresOnly(t *testing.T) {

}

func TestNewManagedConfigurationMongoOnly(t *testing.T) {

}

func TestNewManagedConfigurationCachingOnly(t *testing.T) {

}

func TestNewManagedConfigurationSessionOnly(t *testing.T) {

}

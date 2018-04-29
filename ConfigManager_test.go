package ms_util

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
	conf := NewManagedConfiguration("testdata/postgres.yaml")

	if conf.configuration.Postgres.Host != "example.com" {
		t.Errorf("postgres host mismatch from yaml. Expected %v got %v", "example.com", conf.configuration.Postgres.Host)
	}
	if conf.configuration.Postgres.Database != "testdb" {
		t.Errorf("postgres database mismatch from yaml. Expected %v got %v", "testdb", conf.configuration.Postgres.Database)
	}
	if conf.configuration.Postgres.MaxCon != 100 {
		t.Errorf("postgres connection mismatch from yaml. Expected %v got %v", 100, conf.configuration.Postgres.MaxCon)
	}
}

func TestNewManagedConfigurationMongoOnly(t *testing.T) {
	conf := NewManagedConfiguration("testdata/mongo.yaml")

	if conf.configuration.Mongo.Host != "example.com" {
		t.Errorf("mongo host mismatch from yaml. Expected %v got %v", "example.com", conf.configuration.Mongo.Host)
	}
	if conf.configuration.Mongo.Database != "testdb" {
		t.Errorf("Mongo database mismatch from yaml. Expected %v got %v", "testdb", conf.configuration.Mongo.Database)
	}
	if conf.configuration.Mongo.Port != 27000 {
		t.Errorf("Mongo port mismatch from yaml. Expected %v got %v", 27000, conf.configuration.Mongo.Port)
	}
}

func TestNewManagedConfigurationCachingOnly(t *testing.T) {

	conf := NewManagedConfiguration("testdata/caching.yaml")

	if conf.configuration.Cache.EnableCaching {
		t.Errorf("caching mismatch from yaml. Expected %v got %v", false, conf.configuration.Cache.EnableCaching)
	}

}

func TestNewManagedConfigurationSessionOnly(t *testing.T) {

	conf := NewManagedConfiguration("testdata/session.yaml")

	if conf.configuration.Session.AutogenerateKeyset{
		t.Errorf("session keyset mismatch from yaml. Expected %v got %v", false, conf.configuration.Session.AutogenerateKeyset)
	}

}

func TestPanicOnInvalidYaml(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Didn't panic on invalid yaml")
		}
	}()
	NewManagedConfiguration("testdata/invalid.yaml")
}
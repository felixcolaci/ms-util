package ms_util

import "testing"


func TestNewManagedConfigurationFromFile(t *testing.T) {

	conf := NewManagedConfiguration("testdata/complete.yaml")

	if conf.Configuration.Base.Port != 8080 {
		t.Errorf("Application port mismatch from yaml. Expected %v got %v", 8080, conf.Configuration.Base.Port)
	}

	if conf.Configuration.Base.MgmtPort != 9000 {
		t.Errorf("Management port mismatch from yaml. Expected %v got %v", 9000, conf.Configuration.Base.MgmtPort)
	}

	if conf.Configuration.Postgres.Port != 5432 {
		t.Errorf("postgres port mismatch. Expected %v got %v", 5432, conf.Configuration.Postgres.Port)
	}

}
func TestNewManagedConfigurationWithDefaults(t *testing.T) {
	conf := NewManagedConfiguration()

	if conf.Configuration.Base.Port != 8080 {
		t.Errorf("Application port mismatch from yaml. Expected %v got %v", 8080, conf.Configuration.Base.Port)
	}

	if conf.Configuration.Postgres.Port != 5432 {
		t.Errorf("postgres port mismatch. Expected %v got %v", 5432, conf.Configuration.Postgres.Port)
	}
}

func TestNewManagedConfigurationBaseOnly(t *testing.T) {
	conf := NewManagedConfiguration("testdata/base-only.yaml")

	if conf.Configuration.Base.Port != 8080 {
		t.Errorf("Application port mismatch from yaml. Expected %v got %v", 8080, conf.Configuration.Base.Port)
	}

	if conf.Configuration.Base.MgmtPort != 9000 {
		t.Errorf("Management port mismatch from yaml. Expected %v got %v", 9000, conf.Configuration.Base.MgmtPort)
	}

	if conf.Configuration.Base.BasePath != "/path" {
		t.Errorf("basepath mismatch. Expected %v got %v", "/path", conf.Configuration.Base.BasePath)
	}
}
func TestNewManagedConfigurationPostgresOnly(t *testing.T) {
	conf := NewManagedConfiguration("testdata/postgres.yaml")

	if conf.Configuration.Postgres.Host != "example.com" {
		t.Errorf("postgres host mismatch from yaml. Expected %v got %v", "example.com", conf.Configuration.Postgres.Host)
	}
	if conf.Configuration.Postgres.Database != "testdb" {
		t.Errorf("postgres database mismatch from yaml. Expected %v got %v", "testdb", conf.Configuration.Postgres.Database)
	}
	if conf.Configuration.Postgres.MaxCon != 100 {
		t.Errorf("postgres connection mismatch from yaml. Expected %v got %v", 100, conf.Configuration.Postgres.MaxCon)
	}
}

func TestNewManagedConfigurationMongoOnly(t *testing.T) {
	conf := NewManagedConfiguration("testdata/mongo.yaml")

	if conf.Configuration.Mongo.Host != "example.com" {
		t.Errorf("mongo host mismatch from yaml. Expected %v got %v", "example.com", conf.Configuration.Mongo.Host)
	}
	if conf.Configuration.Mongo.Database != "testdb" {
		t.Errorf("Mongo database mismatch from yaml. Expected %v got %v", "testdb", conf.Configuration.Mongo.Database)
	}
	if conf.Configuration.Mongo.Port != 27000 {
		t.Errorf("Mongo port mismatch from yaml. Expected %v got %v", 27000, conf.Configuration.Mongo.Port)
	}
}

func TestNewManagedConfigurationCachingOnly(t *testing.T) {

	conf := NewManagedConfiguration("testdata/caching.yaml")

	if conf.Configuration.Cache.EnableCaching {
		t.Errorf("caching mismatch from yaml. Expected %v got %v", false, conf.Configuration.Cache.EnableCaching)
	}

}

func TestNewManagedConfigurationSessionOnly(t *testing.T) {

	conf := NewManagedConfiguration("testdata/session.yaml")

	if conf.Configuration.Session.AutogenerateKeyset{
		t.Errorf("session keyset mismatch from yaml. Expected %v got %v", false, conf.Configuration.Session.AutogenerateKeyset)
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
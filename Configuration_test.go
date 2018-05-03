package ms_util

import "testing"

func TestInitBaseWithDefaults(t *testing.T) {

	testconfig := Configuration{}

	testconfig.initBaseWithDefaults()

	if testconfig.Base.Port != 8080 {
		t.Errorf("Error application port. Expected %v got %v", 8080, testconfig.Base.Port)
	}

}

func TestInitPostgresConfigWithDefaults(t *testing.T) {
	testconfig := Configuration{}

	testconfig.initPostgresWithDefaults()

	if testconfig.Postgres.MaxCon != 10 {
		t.Errorf("Error default connection pool size. Expected %v got %v", 10, testconfig.Postgres.MaxCon)
	}

	if testconfig.Postgres.Host != "localhost" {
		t.Errorf("Error default host mismatch. Expected %v got %v", "localhost", testconfig.Postgres.Host)
	}

	if testconfig.Postgres.Port != 5432 {
		t.Errorf("Error default port mismatch. Expected %v got %v", 5432, testconfig.Postgres.Port)
	}

	if testconfig.Postgres.ReinitSchema {
		t.Errorf("reinit schema was set to true expected false")
	}

	if testconfig.Postgres.UseSsl {
		t.Errorf("use ssl was set to true expected false")
	}
}

func TestInitMongoConfigWithDefaults(t *testing.T) {
	conf := Configuration{}
	conf.initMongoWithDefaults()

	if conf.Mongo.Username != "" {
		t.Errorf("username should be empty")
	}

	if conf.Mongo.Password != "" {
		t.Errorf("password should be empty")
	}

	if conf.Mongo.Port != 27017 {
		t.Errorf("mongo default port mismatch. Expected %v got %v", 27017, conf.Mongo.Port)
	}

	if conf.Mongo.Database != "dev-db" {
		t.Errorf("default database mismatch. Expected %v got %v", "dev-db", conf.Mongo.Database)
	}

	if conf.Mongo.Host != "localhost" {
		t.Errorf("default host mispatch. expected %v git %v", "localhost", conf.Mongo.Host)
	}

}

func TestInitCachingConfWithDefaults(t *testing.T) {
	conf := Configuration{}
	conf.initCachingWithDefaults()

	if !conf.Cache.EnableCaching {
		t.Errorf("caching should be enabled by default")
	}
}

func TestInitSessionConfigWithDefaults(t *testing.T) {
	conf := Configuration{}
	conf.initSessionWithDefaults()

	if conf.Session.CookieMaxAge != 3600 {
		t.Errorf("cookie max age mismatch. Expected %v got %v", 3600, conf.Session.CookieMaxAge)
	}
	if conf.Session.CookiePath != "/" {
		t.Errorf("cookie path mismatch. Expected %v got %v", "/", conf.Session.CookiePath)
	}

	if !conf.Session.AutogenerateKeyset {
		t.Errorf("cookie encryption keyset should be generated by default")
	}

	if conf.Session.BlockKeyPath != "" {
		t.Errorf("Path to blockkey should be empty by default")
	}
	if conf.Session.HashkeyPath != "" {
		t.Errorf("Path to hashkey should be empty by default")
	}
}
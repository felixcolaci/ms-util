package ms_util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//Basic Service Configuration
type BaseServiceConfiguration struct {
	//Name of the application
	Name string `yaml:"name"`
	//port of the appliation
	//defaults to 8080
	Port int `yaml:"port"`
	//mgmt port of the application
	//defaults to nil
	MgmtPort int `yaml:"mgmt-port"`
	//basePath of the application
	//defaults to nil
	BasePath string `yaml:"base-path"`
}

//OAuthConfiguration used to connect to an authorization server
type OAuthConfiguration struct {
	//Endpoint from the authorization server used to retrieve token
	TokenEndpoint string `yaml:"token-endpoint"`
	//Endpoint from the authorization server used for authorization redirect
	AuthorizeEndpoint string `yaml:"authorize-endpoint"`
	//The Client Id of the application
	ClientId string `yaml:"client-id"`
	//The client secret of the application
	ClientSecret string `yaml:"client-secret"`
	//Space delimited string of scopes to be requested from authorization server
	Scope string `yaml:"scope"`
	//response type parameter for authorize request
	ResponseType string `yaml:"response-type"`
	//redirect uri for the authorize request
	RedirectUri string `yaml:"redirect-uri"`
}

/*
Configuration to connect to postgres
*/
type PostgresConfig struct {
	//Use to connect to the database
	Username string `yaml:"username"`
	//Users password
	Password string `yaml:"password"`
	//The port which the database accepts connections on
	Port int `yaml:"port"`
	//The name of the database
	Database string `yaml:"database"`
	//The Hostname of the Database
	Host string `yaml:"host"`
	//If set to true the database will be truncated and schema dropped. Defaults to false
	ReinitSchema bool `yaml:"reinit-schema"`
	//If set to true ssl will be used for the connection
	UseSsl bool `yaml:"use-ssl"`
	//Max size of the connection pool
	MaxCon int `yaml:"max-connections"`
	//Max idle connections
	MaxIdleCon int `yaml:"max-idle-connections"`
	//Connection lifetime in minutes
	MaxConLifetime int `yaml:"max-con-lifetime"`
}

//Configuration for establishment of a mongo db connection
type MongoConf struct {
	//defaults to localhost
	Host string `yaml:"host"`
	//defaults to 27017
	Port int `yaml:"port"`
	//defaults to dev-db
	Database string `yaml:"database"`
	//defaults to nil
	Username string `yaml:"username"`
	//defaults to nil
	Password string `yaml:"password"`
}

type SessionHandlingConf struct {
	//If set to true the keyset used to encrypt cookies will be generated upon startup
	//and the possibly provided paths will be ignored.
	//defaults to true
	AutogenerateKeyset bool `yaml:"auto-generate-keyset"`
	//Path to the hashkey
	//defaults to nil
	HashkeyPath string `yaml:"hashkey"`
	//Path to the blockkey
	//defaults to nil
	BlockKeyPath string `yaml:"blockkey"`
	//base path for all generated cookies
	//defaults to "/"
	CookiePath string `yaml:"cookie-path"`
	//Max age of generated cookies
	//defaults to 3600 seconds
	CookieMaxAge int `yaml:"cookie-max-age"`
	//Name of session
	CookieName string `yaml:"cookie-name"`
	//http only mode
	HttpOnly bool `yaml:"cookie-http-only"`
}

type CachingConf struct {
	//If set to true caching will be enabled for all reading endpoints
	//currently the only supported caching method is in memory
	//Defaults to true
	EnableCaching bool `yaml:"caching-enabled"`
}

type Configuration struct {
	Base           BaseServiceConfiguration `yaml:"base,flow"`
	Session        SessionHandlingConf      `yaml:"session,flow"`
	Cache          CachingConf              `yaml:"cache,flow"`
	Postgres       PostgresConfig           `yaml:"postgres,flow"`
	Mongo          MongoConf                `yaml:"mongo,flow"`
	Authentication OAuthConfiguration       `yaml:"authentication,flow"`
}

type ConfigManager struct {
	Configuration *Configuration
}

func (c *Configuration) initSessionWithDefaults() {
	c.Session.AutogenerateKeyset = true
	c.Session.CookiePath = "/"
	c.Session.CookieMaxAge = 3600
}

func (c *Configuration) initBaseWithDefaults() {
	c.Base.Port = 8080
}

func (c *Configuration) initPostgresWithDefaults() {
	c.Postgres.Database = "dev-db"
	c.Postgres.Host = "localhost"
	c.Postgres.Port = 5432
	c.Postgres.Username = "postgres"
	c.Postgres.Password = "postgres"
	c.Postgres.ReinitSchema = false
	c.Postgres.UseSsl = false
	c.Postgres.MaxCon = 10
}

func (c *Configuration) initMongoWithDefaults() {
	c.Mongo.Database = "dev-db"
	c.Mongo.Host = "localhost"
	c.Mongo.Port = 27017
}

func (c *Configuration) initCachingWithDefaults() {
	c.Cache.EnableCaching = true
}

func (c *ConfigManager) ReinitializeConfig(args ...string) {
	c.Configuration = NewManagedConfiguration(args...).Configuration
}

func NewManagedConfiguration(args ...string) *ConfigManager {
	conf := ConfigManager{}

	path := "application.yaml"
	if len(args) > 0 {
		path = args[0]
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {

		config := Configuration{}
		config.initBaseWithDefaults()
		config.initPostgresWithDefaults()
		config.initMongoWithDefaults()
		config.initSessionWithDefaults()
		config.initCachingWithDefaults()

		conf.Configuration = &config
	} else {
		err := yaml.Unmarshal(file, &conf.Configuration)
		if err != nil {
			panic(err)
		}
	}

	return &conf
}

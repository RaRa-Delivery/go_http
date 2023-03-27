package framework

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	routing "github.com/qiangxue/fasthttp-routing"
)

var currentContext *Framework

type Framework struct {
	env        string
	flavour    string
	appBaseDir string
	port       string
	config     map[string]interface{}
	Mongo      map[string]MongoConnection
	httpSystem HTTPSystem
	Router     *routing.Router
}

/**
ENV Preparation
*/
func (f *Framework) LoadEnv() {
	f.env = os.Getenv("GO_ENV")
	if len(f.env) == 0 {
		f.env = "development"
	}
	if len(f.flavour) == 0 {
		f.flavour = "debug"
	}
}

func (f *Framework) LoadPort() {
	f.port = os.Getenv("FORCE_PORT")
	if len(f.port) != 0 {
		return
	}
	f.port = os.Getenv("PORT")
	if len(f.port) == 0 {
		f.port = "8080"
	}
}

func (f *Framework) GetEnv() string {
	return f.env
}

func (f *Framework) GetFlavour() string {
	return f.flavour
}

/**
Config management
*/
func (f *Framework) SetConfig(conf map[string]interface{}) {
	f.config = conf
}

func (f *Framework) GetConfValue(propName string) interface{} {
	return f.config[propName]
}

/**
HTTP
*/
func (f *Framework) setHTTPSystem(system HTTPSystem) {
	f.httpSystem = system
	f.Router = system.Router
}

func (f *Framework) Listen() {
	f.httpSystem.listen()
}

/**
DB
*/
func (f *Framework) InitMongo(descriptions []MongoConnectionDescription) {
	// load mongo session.
	f.Mongo = make(map[string]MongoConnection)
	for _, connectionDescription := range descriptions {
		log.Println("Connecting Mongo Db '" + connectionDescription.Name + "' -> " + connectionDescription.Description)
		camsDbMongoConnection := MongoConnection{
			Name: connectionDescription.Name,
			URL:  os.Getenv(connectionDescription.EnvVarName),
		}
		err := camsDbMongoConnection.init()
		if err == nil {
			f.Mongo[connectionDescription.Name] = camsDbMongoConnection
		} else if connectionDescription.CanFail == true {
			log.Println(err)
			log.Println("Not failing the connection as CanFail is true.")
		} else {
			// Panic Error
			panic(err)
		}
	}
}

/**
LOGGER Functions
*/
func (f *Framework) CoreLog(message interface{}) {
	f.Log("[FRAMEWORK]", message)
}

func (f *Framework) Debug(message interface{}) {
	f.Log("[DEBUG]", message)
}

func (f *Framework) Trace(message interface{}) {
	f.Log("[TRACE]", message)
}

func (f *Framework) Info(message interface{}) {
	f.Log("[INFO]", message)
}

func (f *Framework) Error(message interface{}) {
	f.Log("[ERROR]", message)
}

func (f *Framework) Warn(message interface{}) {
	f.Log("[WARNING]", message)
}

func (f *Framework) Log(emblem string, message interface{}) {
	bold := color.New(color.Bold).SprintFunc()
	fmt.Printf(bold("%s %+v\n"), bold(emblem), message)
}

func (f *Framework) GetBaseDir() string {
	return f.appBaseDir
}

func (f *Framework) SetAsMainContext() {
	currentContext = f
}

func GetCurrentAppContext() *Framework {
	return currentContext
}

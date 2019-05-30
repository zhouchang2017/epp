package config

import (
	"github.com/davecgh/go-spew/spew"
	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

const (
	RUN_MODE_LOCAL     = "local"
	RUN_MODE_CONTAINER = "container"
	RUN_MODE_K8S       = "k8s"
)

func init() {
	loadConf()
}

var Config *config

type config struct {
	ServiceName     string `yaml:"service_name"`
	ProjectRealPath string
	DB              dbconfig `yaml:"db"`
}

func loadConf() {

	runMode := os.Getenv("RUN_MODE")
	log.Println("run mode:", runMode)

	var confFile string
	switch runMode {
	case RUN_MODE_LOCAL:
		confFile = "./dev.yaml"
	case RUN_MODE_CONTAINER:
		confFile = "./container.yaml"
	case RUN_MODE_K8S:
		confFile = "./k8s.yaml"
	default:
		log.Fatalln("unsuppoer run mode! supports:[local,container,k8s]")
	}

	conf, _ := ioutil.ReadFile(confFile)
	if err := yaml.Unmarshal(conf, &Config); err != nil {
		log.Fatalln("conf load failed", err)
	}

	log.Println(spew.Sdump(Config))
}

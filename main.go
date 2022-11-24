package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type ConfigFile struct {
	Server      string
	CookieName  string `yaml:"cookieName"`
	Credentials []Credential
}

type Credential struct {
	Username string
	Password string
}

type MappingFile struct {
	Mappings map[string]int
}

var mappingsFile MappingFile
var configFile ConfigFile
var nextIndex = 0

func readFile(filename string) []byte {
	filename = os.Getenv("CCC_DATA") + filename

	log.Infof("Reading file: %v", filename)

	_, err := os.Stat(filename)

	if err != nil {
		os.Create(filename)
	}

	f, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("%v", err)
		return []byte{}
	}

	return f
}

func assignUUID(res http.ResponseWriter) (error, string) {
	if nextIndex == len(configFile.Credentials) {
		return errors.New("No more capacity left."), ""
	}

	id := fmt.Sprintf("%s", uuid.New())

	c := http.Cookie{
		Name:    configFile.CookieName,
		Value:   id,
		Expires: time.Now().Add(365 * 24 * time.Hour),
		Path:    "/",
		Secure:  false,
	}

	log.Infof("Setting cookie %v = %v", configFile.CookieName, id)

	mappingsFile.Mappings[id] = nextIndex

	log.Infof("Assigned: %v", nextIndex)

	nextIndex += 1

	mmappings, err := yaml.Marshal(&mappingsFile)

	if err != nil {
		log.Fatalf("%v", err)
	}

	err = os.WriteFile(os.Getenv("CCC_DATA")+"/mappings.yaml", []byte(string(mmappings)), 0644)

	if err != nil {
		log.Fatalf("%v", err)
	}

	http.SetCookie(res, &c)

	return nil, id
}

func getCookie(req *http.Request) string {
	for _, cookie := range req.Cookies() {
		if cookie.Name == configFile.CookieName {
			log.Infof("Got Cookie %v", cookie.Value)

			return cookie.Value
		}
	}

	return ""
}

func getCredentials(uuid string) (string, string) {
	if _, ok := mappingsFile.Mappings[uuid]; !ok {
		return "notfound!", "notfound!"
	}

	mapping := mappingsFile.Mappings[uuid]

	creds := configFile.Credentials[mapping]

	return creds.Username, creds.Password
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	if getCookie(req) == "" {
		err, id := assignUUID(res)

		if err == nil {
			fmt.Fprintf(res, "You have just been assigned %v. Please refresh this page.\n\n", id)
		} else {
			log.Errorf("%v", err)
			fmt.Fprintf(res, "%v", err)
		}
	} else {
		username, password := getCredentials(getCookie(req))

		fmt.Fprintf(res, "Your UUID is: %v\n\n", getCookie(req))
		fmt.Fprintf(res, "Your username is: %v\n\n", username)
		fmt.Fprintf(res, "Your password is: %v\n\n", password)
		fmt.Fprintf(res, "Console address: %v", configFile.Server)
	}
}

func main() {
	log.Infof("CookieCredentialCannon")

	dataDir := os.Getenv("CCC_DATA")
	log.Infof("Datadir: %v", dataDir)

	err := yaml.UnmarshalStrict(readFile("mappings.yaml"), &mappingsFile)

	if err != nil {
		log.Fatalf("%v", err)
	}

	if mappingsFile.Mappings == nil {
		mappingsFile.Mappings = make(map[string]int)
	}

	err = yaml.UnmarshalStrict(readFile("config.yaml"), &configFile)

	if err != nil {
		log.Fatalf("%v", err)
	}

	nextIndex = len(mappingsFile.Mappings)

	log.Infof("Startup config: %+v", configFile)
	log.Infof("Next index: %v", nextIndex)

	//	http.Handle("/data", http.FileServer(http.Dir(dataDir)))
	http.HandleFunc("/", handleIndex)

	log.Error(http.ListenAndServe(":8080", nil))
}

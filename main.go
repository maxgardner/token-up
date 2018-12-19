package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

var isTest bool

func init() {
	flag.BoolVar(&isTest, "test", false, "output example config file rather than updating official one")
}

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Error getting current user: %v\n", err)
	}
	kbDir := filepath.Join(usr.HomeDir, "/.kube")
	dwnDir := filepath.Join(usr.HomeDir, "/Downloads")

	configFile, err := ioutil.ReadFile(filepath.Join(kbDir, "/config"))
	if err != nil {
		log.Fatalf("Error reading config file: %v\n", err)
	}

	config := &Config{}
	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		log.Fatalf("Failed to unmarshal main config file: %v\n", err)
	}

	downloads, err := ioutil.ReadDir(dwnDir)
	if err != nil {
		log.Fatalf("Failed to read downloads directory: %v\n", err)
	}
	cfCount := 0

	for _, file := range downloads {
		if strings.Contains(file.Name(), "kube-config") && !file.IsDir() {
			cfCount++
			cfFile, err := ioutil.ReadFile(filepath.Join(dwnDir, file.Name()))
			if err != nil {
				log.Printf("Failed to read config file %s: %v\n", file.Name(), err)
				continue
			}

			cluster := strings.Replace(file.Name(), "kube-config-", "", 1)
			cf := &Config{}
			err = yaml.Unmarshal(cfFile, cf)
			if err != nil {
				log.Fatalf("Failed to unmarshal data from %s: %v\n", file.Name(), err)
			}

			for i, clstr := range config.Users {
				if strings.Contains(clstr.Name, cluster) {
					log.Printf("Updating token for %s\n", cluster)
					config.Users[i].User.AuthProvider.Config.IDToken = cf.Users[0].User.AuthProvider.Config.IDToken
					os.Remove(filepath.Join(dwnDir, file.Name()))
				}
			}
		}
	}

	if cfCount == 0 {
		log.Printf("No config files found in %s\n", dwnDir)
		return
	}

	newConfig, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("Error marshalling new config data: %v\n", err)
	}

	outPath := filepath.Join(kbDir, "/config")

	if isTest {
		outPath = "./test/config"
	}

	err = ioutil.WriteFile(outPath, newConfig, os.FileMode(0777))
	if err != nil {
		log.Fatalf("Error updating config file: %v\n", err)
	}
}

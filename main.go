package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"checkVersionJetBrains/src/application"
	"checkVersionJetBrains/src/domain"
)

func main() {
	configFile := application.CheckParam(os.Args)
	fmt.Println("")
	fmt.Println(strings.Repeat("-", 75))
	installedProducts := getLocalVersions(configFile)
	releaseData := getReleaseData()
	checkVersions(releaseData, installedProducts)
	fmt.Println(strings.Repeat("-", 75))
}

func checkVersions(releaseData domain.Products, installedProducts domain.LocalProducts) {
	for _, installedProduct := range installedProducts {
		var latestProduct domain.LatestProduct
		for _, product := range releaseData.Products {
			if product.Name != installedProduct.Name {
				continue
			}
			latestProduct.Name = product.Name

			if latestProduct.Name == "" {
				continue
			}
			for _, channel := range product.Channels {
				if channel.Status != "release" {
					continue
				}
				for _, build := range channel.Build {
					if latestProduct.ReleaseDate < build.ReleaseDate {
						latestProduct.ReleaseDate = build.ReleaseDate
						latestProduct.Version = build.Version
						latestProduct.Url = channel.Url
					}
				}

			}
		}
		if latestProduct.Version != installedProduct.Version {
			fmt.Printf("* %10s, installed version %s mismatch latest %s\nURL: %s\n", installedProduct.Name, installedProduct.Version, latestProduct.Version, latestProduct.Url)
		} else {
			fmt.Printf("* %10s: latest version %s installed\n", installedProduct.Name, installedProduct.Version)
		}
	}
}

func getReleaseData() domain.Products {
	res, err := http.Get(domain.CheckUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	releaseData := domain.Products{}
	err = xml.Unmarshal(data, &releaseData)
	if err != nil {
		log.Fatal(err)
	}
	return releaseData
}

func getLocalVersions(configFile string) domain.LocalProducts {
	var installedWithVersion domain.LocalProducts
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	config := domain.ConfigFile{}
	err = yaml.Unmarshal(data, &config) //nolint:typecheck
	if err != nil {
		log.Fatal(err)
	}
	basePath := application.FullQualifiedPath(config.BasePath)
	for _, product := range config.ProductConfigs {
		data, err := os.ReadFile(filepath.Join(basePath, product.Path, "product-info.json"))
		if err != nil {
			log.Fatal(err)
		}
		version := struct {
			Version string `json:"version"`
		}{}
		err = json.Unmarshal(data, &version)
		if err != nil {
			log.Fatal(err)
		}
		installedWithVersion = append(installedWithVersion, domain.LocalProduct{
			Name:    product.Name,
			Version: version.Version,
		})
	}
	return installedWithVersion
}

package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"checkVersionJetBrains/src/domain"
)

func main() {
	fmt.Println("")
	fmt.Println(strings.Repeat("-", 55))
	releaseData := getReleaseData()
	installedProducts := getLocalVersions(domain.GetInstalledProducts())
	checkVersions(releaseData, installedProducts)
	fmt.Println(strings.Repeat("-", 55))
}

func checkVersions(releaseData domain.Products, installedProducts domain.InstalledProducts) {
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
					}
				}

			}
		}
		if latestProduct.Version != installedProduct.Version {
			fmt.Printf("%10s, installed version %s mismatch latest %s\n", installedProduct.Name, installedProduct.Version, latestProduct.Version)
		} else {
			fmt.Printf("%10s: latest version %s installed\n", installedProduct.Name, installedProduct.Version)
		}
	}
}

func getReleaseData() domain.Products {
	res, err := http.Get(domain.CheckUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
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

func getLocalVersions(products domain.InstalledProducts) domain.InstalledProducts {
	var installedWithVersion domain.InstalledProducts
	for _, product := range products {
		data, err := ioutil.ReadFile(filepath.Join(domain.BasePath, product.File))
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
		installedWithVersion = append(installedWithVersion, domain.InstalledProduct{
			Name:    product.Name,
			File:    product.File,
			Version: version.Version,
		})
	}
	return installedWithVersion
}

package domain

const BasePath = "/usr/local"

type (
	LatestProduct struct {
		Name string
		ReleaseDate int
		Version string
	}
	InstalledProducts []InstalledProduct

	InstalledProduct struct {
		Name    string
		File    string
		Version string
}

	Products struct {
		Products []Product `xml:"product"`
	}
	Product struct {
		Name     string    `xml:"name,attr"`
		Channels []Channel `xml:"channel"`
	}
	Channel struct {
		Status string `xml:"status,attr"`
		Build  []Build  `xml:"build"`
	}
	Build struct {
		//Number      string `xml:"number,attr"`
		Version     string `xml:"version,attr"`
		ReleaseDate int `xml:"releaseDate,attr"`
	}
)

func GetInstalledProducts() InstalledProducts {
	return InstalledProducts{
		struct {
			Name string
			File string
			Version string
		}{
			Name: "GoLand",
			File: "Go/GoLand/product-info.json",
		},
		{
			Name: "WebStorm",
			File: "Javascript/WebStorm/product-info.json",
		},
	}

}

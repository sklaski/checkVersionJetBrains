package domain

const BasePath = "/usr/local"
const CheckUrl = "https://www.jetbrains.com/updates/updates.xml"

type (
	LatestProduct struct {
		Name        string
		ReleaseDate int
		Version     string
		Url         string
	}
	InstalledProducts []InstalledProduct

	InstalledProduct struct {
		Name    string
		File    string
		Version string
		Url     string
	}

	Products struct {
		Products []Product `xml:"product"`
	}
	Product struct {
		Name     string    `xml:"name,attr"`
		Channels []Channel `xml:"channel"`
	}
	Channel struct {
		Status string  `xml:"status,attr"`
		Url    string  `xml:"url,attr"`
		Build  []Build `xml:"build"`
	}
	Build struct {
		//Number      string `xml:"number,attr"`
		Version     string `xml:"version,attr"`
		ReleaseDate int    `xml:"releaseDate,attr"`
	}
)

func GetInstalledProducts() InstalledProducts {
	return InstalledProducts{
		struct {
			Name    string
			File    string
			Version string
			Url     string
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

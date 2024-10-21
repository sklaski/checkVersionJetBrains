package domain

const CheckUrl = "https://www.jetbrains.com/updates/updates.xml"

type (
	LatestProduct struct {
		Name        string
		ReleaseDate int
		Version     string
		Url         string
	}

	LocalProducts []LocalProduct

	LocalProduct struct {
		Name    string
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
		Status string  `xml:"status,attr"`
		Url    string  `xml:"url,attr"`
		Build  []Build `xml:"build"`
	}
	Build struct {
		Version     string `xml:"version,attr"`
		ReleaseDate int    `xml:"releaseDate,attr"`
	}
)

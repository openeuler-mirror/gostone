package conf

var (
	Url      = ""
	Version2 = map[string]interface{}{}
	Version3 = map[string]interface{}{}
)

func InitUrl(url string) {
	Url = url
	initVersion()
}

func initVersion() {
	Version2 = map[string]interface{}{
		"id":      "v2.0",
		"status":  "deprecated",
		"updated": "2016-08-04T00:00:00Z",
		"links": []interface{}{
			map[string]interface{}{
				"rel":  "self",
				"href": Url + "/v2.0/",
			},
			map[string]interface{}{
				"rel":  "describedby",
				"type": "text/html",
				"href": "https://docs.openstack.org/",
			},
		},
		"media-types": []interface{}{
			map[string]interface{}{
				"base": "application/json",
				"type": "application/vnd.openstack.identity-v2.0+json",
			},
		},
	}
	Version3 = map[string]interface{}{
		"id":      "v3.10",
		"status":  "stable",
		"updated": "2018-02-28T00:00:00Z",
		"links": []interface{}{
			map[string]interface{}{
				"rel":  "self",
				"href": Url + "/v3/",
			},
		},
		"media-types": []interface{}{
			map[string]interface{}{
				"base": "application/json",
				"type": "application/vnd.openstack.identity-v3+json",
			},
		},
	}

}

package kongapi

/*
{
    "hostname": "",
    "lua_version": "LuaJIT 2.1.0-alpha",
    "plugins": {
        "available_on_server": [
            ...
        ],
        "enabled_in_cluster": [
            ...
        ]
    },
    "configuration" : {
        ...
    },
    "tagline": "Welcome to Kong",
    "version": "0.6.0"
}
*/

// Information retrieves generic details about a node
type Information struct {
	Hostname   string `json:"hostname"`
	LuaVersion string `json:"lua_version"`
	Tagline    string `json:"tagline"`
	Version    string `json:"version"`
}

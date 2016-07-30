package kongapi

/*
{
    "server": {
        "total_requests": 3,
        "connections_active": 1,
        "connections_accepted": 1,
        "connections_handled": 1,
        "connections_reading": 0,
        "connections_writing": 1,
        "connections_waiting": 0
    },
    "database": {
        "apis": 2,
        "consumers": 0,
        "plugins": 2,
        "nodes": 1,
        ...
    }
}

*/

// Status retrieves usage information about a node
type Status struct {
	Server   Server   `json:"server"`
	Database Database `json:"server"`
}

// Server encapsulates metrics about the nginx HTTP/S server
type Server struct {
	TotalRequests       int `json:"server"`
	ConnectionsActive   int `json:"connections_active"`
	ConnectionsAccepted int `json:"connections_accepted"`
	ConnectionsHandled  int `json:"connections_handled"`
	ConnectionsReading  int `json:"connections_reading"`
	ConnectionsWriting  int `json:"connections_writing"`
	ConnectionsWaiting  int `json:"connections_waiting"`
}

// Database encapsulates metrics about the database collections
type Database struct {
	APIs      int `json:"apis"`
	Consumers int `json:"consumers"`
	Plugins   int `json:"plugins"`
	Nodes     int `json:"nodes"`
}

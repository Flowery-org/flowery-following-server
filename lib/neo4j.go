package lib

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"sync"
)

type Neo4jClient struct {
	Instance *FollowingClient
}

//* Singleton Pattern

var (
	inst *Neo4jClient
	mu   sync.Mutex
)

func (c *Neo4jClient) Connect(uri, username, password string) error {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return err
	}

	c.Instance = &FollowingClient{driver: driver}
	return nil
}

func GetNeo4jClientInstance() *Neo4jClient {
	if inst == nil {
		mu.Lock()
		defer mu.Unlock()
		if inst == nil { // Thread Safety
			inst = &Neo4jClient{}
		}
	}
	return inst
}

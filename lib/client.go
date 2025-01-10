package lib

import (
	"context"
	"flowery-following-server/dto"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowingClient struct {
	driver neo4j.Driver
}

// CreateUser creates a new user node in GDB
func (manager *FollowingClient) CreateUser(ctx context.Context, user dto.User) error {
	session := manager.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer closeSession(&session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := `
			CREATE (user:User {id: $id})
			RETURN user
		`

		params := map[string]interface{}{
			"id": user.Id,
		}

		_, err := tx.Run(query, params)

		return nil, err
	})

	return err
}

// FollowUser creates a FOLLOWS relationship between two users
func (manager *FollowingClient) FollowUser(ctx context.Context, rel dto.CreateRelation) error {
	session := manager.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer closeSession(&session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := `
			MATCH (follower:User {id: $followerId})		
			MATCH (following:User {id: $followingId})
			MERGE (follower)-[rel: FOLLOWS {createdAt: $createdAt}]->(following)
			return rel
		`

		params := map[string]interface{}{
			"followerId":  rel.FollowerId,
			"followingId": rel.FollowingId,
			"createdAt":   rel.CreatedAt,
		}

		_, err := tx.Run(query, params)
		return nil, err
	})

	return err
}

func (manager *FollowingClient) UnfollowUser(ctx context.Context, rel dto.DeleteRelation) error {
	session := manager.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer closeSession(&session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := `
			MATCH (follower:User {id: $followerId})-[rel: FOLLOWS]->(following: User {id: $followingId})
			DELETE rel
		`

		params := map[string]interface{}{
			"followerId":  rel.FollowerId,
			"followingId": rel.FollowingId,
		}

		_, err := tx.Run(query, params)
		return nil, err
	})

	return err

}

//func (manager *FollowingClient) DeleteUser(ctx context.Context, rel dto.CreateRelation) error {
//
//}

func closeSession(session *neo4j.Session) {
	s := *session
	err := s.Close()
	if err != nil {
		//* TODO Error Handling
		panic("Failed to close session: " + err.Error())
	}
}

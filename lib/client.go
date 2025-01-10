package lib

import (
	"context"
	"flowery-following-server/dto"
	"fmt"
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

func (manager *FollowingClient) GetAllFollowers(ctx context.Context, userId string) ([]dto.User, error) {
	session := manager.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer closeSession(&session)

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := `
			MATCH (follower:User)-[:FOLLOWS]->(user: User {id: $uid})
			RETURN follower.id as id	
		`

		result, err := tx.Run(query, map[string]interface{}{"uid": userId})

		if err != nil {
			return nil, err
		}

		var followers []dto.User
		for result.Next() {
			record := result.Record()
			id, _ := record.Get("id")
			followers = append(followers, dto.User{
				Id: id.(string),
			})
		}

		return followers, nil
	})

	if err != nil {
		return nil, err
	}

	followers, ok := result.([]dto.User)

	if !ok {
		return nil, fmt.Errorf("Failed to convert result to []dto.User")
	}

	return followers, nil
}

func (manager *FollowingClient) GetAllFollowings(ctx context.Context, userId string) ([]dto.User, error) {
	session := manager.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer closeSession(&session)

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := `
			MATCH (user:User {id: $uid})-[:FOLLOWS]->(following: User)
			RETURN following.id as id	
		`

		result, err := tx.Run(query, map[string]interface{}{"uid": userId})

		if err != nil {
			return nil, err
		}

		var followers []dto.User
		for result.Next() {
			record := result.Record()
			id, _ := record.Get("id")
			followers = append(followers, dto.User{
				Id: id.(string),
			})
		}

		return followers, nil
	})

	if err != nil {
		return nil, err
	}

	followers, ok := result.([]dto.User)

	if !ok {
		return nil, fmt.Errorf("Failed to convert result to []dto.User")
	}

	return followers, nil
}

func (manager *FollowingClient) DeleteUser(ctx context.Context, userId string) error {
	session := manager.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer closeSession(&session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := `
			MATCH (user:User {id: $uid})
			DELETE user
		`

		params := map[string]interface{}{"uid": userId}

		_, err := tx.Run(query, params)
		return nil, err
	})

	return err

}

func closeSession(session *neo4j.Session) {
	s := *session
	err := s.Close()
	if err != nil {
		//* TODO Error Handling
		panic("Failed to close session: " + err.Error())
	}
}

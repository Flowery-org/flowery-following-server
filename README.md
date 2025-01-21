# Flowery-Following-Server

> A server dedicated to provide follow / following service.

## 1. Overview
`Flowery-following-server` is a server that provides follow / following service. Based on GDB and `gin` server, it provides various follow-related features with high performance.

The followings are the features to be implemented.
- [x] Enlist new user (Sign-up)
- [x] Follow and unfollow users
- [x] Query followers and followings
- [x] Remove existing user
- [ ] Block user

The server is currently runs on port `13456`, but it is hard-coded. It needs to be decoupled into environments in near future. 

## 2. Tech Stack
The followings are the tech stacks that the server uses.

- `Gin`: A light-weighted go server framework.
- `Neo4J`: A GDB for following-follower service.
- `Docker` (TODO): A container tool to containerize the application.

## 3. API Docs
The followings are the document for API usage.

#### i) `PUT /v1/user`
- Enlist new user into GDB.
- Request Body:
```json
{
  "id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
```
```bash
curl -X PUT localhost:13456/v1/user -H "Content-Type: application/json" -d '{"id": "a"}' 
```
- Response Body
```json
{
  "ok": true
}
```
- Error Code
  - `500`: Internal Error; Failed to read request body or create user.

#### ii) `DELETE /v1/user`
- Delete existing user into GDB.
- Request Body:
```json
{
  "id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
```
```bash
curl -X DELETE localhost:13456/v1/user -H "Content-Type: application/json" -d '{"id": "a"}'
```
- Response Body
```json
{
  "ok": true
}
```
- Error Code
    - `500`: Internal Error; Failed to read request body or delete user

#### iii) `PUT /v1/rel`
- Create a follower-following relationship between user.
- Request Body:
```json
{
  "followerId": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
  "followingId": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
```
```bash
curl -X PUT localhost:13456/v1/rel -H "Content-Type: application/json" -d '{"followerId": "a", "followingId": "b"}'
```
- Response Body
```json
{
  "ok": true
}
```
- Error Code
    - `500`: Internal Error; Failed to read request.
    - `400`: Bad Input; User with such id doesn't exist.

#### iv) `DELETE /v1/rel`
- Delete the existing follower-following relationship between user.
- Request Body:
```json
{
  "followerId": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
  "followingId": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
```
```bash
curl -X DELETE localhost:13456/v1/rel -H "Content-Type: application/json" -d '{"followerId": "a", "followingId": "b"}'```
- Response Body
```json
{
  "ok": true
}
```
- Error Code
    - `500`: Internal Error; Failed to read request.
    - `400`: Bad Input; User with such id doesn't exist.

#### v) `GET /v1/rel/followings?userId={USER_ID}`
- Get all the followings of user
```bash
 curl 'localhost:13456/v1/rel/followings?userId=a'
```
 - Response Body
```json
[
  {
    "id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  },
  //...
]
```
- Error Code
    - `500`: Internal Error; Failed to read request or read from GDB.

#### v) `GET /v1/rel/followers?userId={USER_ID}`
- Get all the followers of user
```bash
 curl 'localhost:13456/v1/rel/followers?userId=a'
```
- Response Body
```json
[
  {
    "id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  },
  //...
]
```
- Error Code
    - `500`: Internal Error; Failed to read request or read from GDB.

## 4. TODOs
- The followings are TODOs to complete in near future.

- [ ] Dockerize
- [ ] Prevent duplicate user creation
- [ ] Error Handling
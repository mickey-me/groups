package main

import "golang.org/x/crypto/ed25519"

//alternatives
type Role string
type RequestType string
type MembershipRequestType string

//type Hash string
//type EncryptedHash string
type Data string

//type PublicKey []byte  //---> already implemented in crypto/ed25519
//type PrivateKey []byte //---> already implemented in crypto/ed25519

type Group struct {
	//	Id          string
	Name        string
	Description string
	Data        map[RequestType]Data

	Creator UserID

	Members               []UserID
	Memberships           map[Role][]UserID
	Authorizations        map[RequestType][]Role
	Membersauthorizations map[UserID][]RequestType
}

type DataBaseRequest struct {
	RequestNum int32
	Signature  Signature
}

type DataBaseResponse struct {
	Accepted bool
}

type DataBaseMessage struct {
	MessageStatus MessageStatus
}

type GroupCreationRequest struct {
	DataBaseRequest
	Group Group
}

type GroupCreationResponse struct {
	Accepted bool
}

type GroupCreationMessage struct {
	DataBaseMessage
	Request  GroupCreationRequest
	Response GroupCreationResponse
}

type MembershipRequest struct {
	DataBaseRequest
}

type MembershipResponse struct {
}

type MembershipMessage struct {
	DataBaseMessage
	Request  MembershipRequest
	Response MembershipResponse
}

type UserRequest struct {
	DataBaseRequest
	UserID     UserID
	PubllicKey ed25519.PublicKey
}

type UserResponse struct {
	Accepted bool
}

type UserMessage struct {
	DataBaseMessage
	Request  UserRequest
	Response UserResponse
}

// type Membership struct {
// 	types map[string]Role
// }

// type Role struct {
// 	members []UserID
// }

type Signature struct {
	Hash          [32]byte
	Encryptedhash []byte
}

//enums
type MessageStatus int

const (
	Received  MessageStatus = 0
	Confirmed MessageStatus = 1
	Sent      MessageStatus = 2
	Succeeded MessageStatus = 3
	Failed    MessageStatus = 4
)
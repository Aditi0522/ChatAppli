package repo

import (
	"time"
	"github.com/google/uuid"
)

type ConversationType string 

const (
	ConversationTypeDirect ConversationType = "direct"
	ConversationTypeGroup ConversationType = "group"
)

type MemberType string

const (
	MemberTypeAdmin MemberType = "admin"
	MemberTypeMember MemberType = "member"
)

type MessageKind string

const (
	MessageKindText  MessageKind = "text"
	MessageKindMedia MessageKind = "media"
)

type conversation struct {
      ID  uuid.UUID
	  Type  ConversationType
	  CreatedAt  time.Time
	  Title   *string
}

type members struct {
      ConvId uuid.UUID
	  UserId uuid.UUID
	  JoinedAt time.Time
	  LeftAt *time.Time
	  LastReadMessageId *uuid.UUID
	  Role MemberType
}

type messages struct {
      ID uuid.UUID
	  ConvId uuid.UUID
	  SenderId uuid.UUID
	  Kind MessageKind
	  CreatedAt time.Time
	  EditedAt *time.Time
	  DeletedAt *time.Time
	  Body *string
}
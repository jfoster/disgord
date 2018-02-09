package discord

import "github.com/andersfylling/snowflake"

type Role struct {
	ID          snowflake.ID `json:"id"`
	Name        string       `json:"name"`
	Managed     bool         `json:"managed"`
	Mentionable bool         `json:"mentionable"`
	Hoist       bool         `json:"hoist"`
	Color       int          `json:"color"`
	Position    int          `json:"position"`
	Permissions uint64       `json:"permissions"`
}

func NewRole() *Role {
	return &Role{}
}

func (r *Role) Clear() {

}

type RoleEvent struct {
	Role    *Role        `json:"role"`
	GuildID snowflake.ID `json:"guild_id"`
}

type RoleDeleteEvent struct {
	RoleID  snowflake.ID `json:"role_id"`
	GuildID snowflake.ID `json:"guild_id"`
}

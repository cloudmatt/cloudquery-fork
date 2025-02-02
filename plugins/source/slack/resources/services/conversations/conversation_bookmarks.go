// Code generated by codegen; DO NOT EDIT.

package conversations

import (
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ConversationBookmarks() *schema.Table {
	return &schema.Table{
		Name:        "slack_conversation_bookmarks",
		Description: `https://api.slack.com/methods/bookmarks.list`,
		Resolver:    fetchConversationBookmarks,
		Multiplex:   client.TeamMultiplex,
		Columns: []schema.Column{
			{
				Name:     "team_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveTeamID,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "channel_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ChannelID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "title",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Title"),
			},
			{
				Name:     "link",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Link"),
			},
			{
				Name:     "emoji",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Emoji"),
			},
			{
				Name:     "icon_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IconURL"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "date_created",
				Type:     schema.TypeTimestamp,
				Resolver: client.JSONTimeResolver("Created"),
			},
			{
				Name:     "date_updated",
				Type:     schema.TypeTimestamp,
				Resolver: client.JSONTimeResolver("Updated"),
			},
			{
				Name:     "rank",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Rank"),
			},
			{
				Name:     "last_updated_by_user_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastUpdatedByUserID"),
			},
			{
				Name:     "last_updated_by_team_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastUpdatedByTeamID"),
			},
			{
				Name:     "shortcut_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ShortcutID"),
			},
			{
				Name:     "entity_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EntityID"),
			},
			{
				Name:     "app_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AppID"),
			},
		},
	}
}

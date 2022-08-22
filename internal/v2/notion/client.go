package notion

import (
	"os"

	_ "github.com/gmarcha/discord-integration/internal/v2/env"
	"github.com/jomei/notionapi"
)

var (
	Client      *notionapi.Client
	notionToken notionapi.Token
)

func init() {

	notionToken = notionapi.Token(os.Getenv("NOTION_TOKEN"))
}

func init() {

	Client = notionapi.NewClient(notionToken)
}

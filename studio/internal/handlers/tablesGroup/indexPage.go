package tablesGroup

import (
	"github.com/gofiber/fiber/v2"
	coreUtils "github.com/nonamecat19/go-orm/core/utils"
	"github.com/nonamecat19/go-orm/studio/internal/utils"
	tablesView "github.com/nonamecat19/go-orm/studio/internal/view/tables"
)

func TablesPage(c *fiber.Ctx) error {
	sharedData := utils.GetSharedData(c)

	tables := make([]tablesView.Table, len(sharedData.Tables))
	for i, table := range sharedData.Tables {
		name := table.Info()
		tables[i] = tablesView.Table{
			Title: coreUtils.ToHumanCase(name),
			ID:    name,
		}
	}

	props := tablesView.TablePageProps{
		Tables: tables,
		Prefix: sharedData.Prefix,
	}

	return utils.Render(c, tablesView.TablesPage(props))
}

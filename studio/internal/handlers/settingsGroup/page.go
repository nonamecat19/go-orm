package settingsGroup

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nonamecat19/go-orm/studio/internal/utils"
	"github.com/nonamecat19/go-orm/studio/internal/view/settings"
)

func SettingsPage(c *fiber.Ctx) error {
	sharedData := utils.GetSharedData(c)
	return utils.Render(c, settings.SettingsPage(settings.SettingsPageProps{
		Prefix:       sharedData.Prefix,
		AssetsPrefix: sharedData.AssetsPrefix,
	}))
}

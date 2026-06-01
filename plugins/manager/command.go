package manager

import (
	"fmt"

	"wxbot/engine/control"
	"wxbot/engine/pkg/log"
	"wxbot/engine/robot"
)

var command Command

type Command struct {
	MenuMode string `gorm:"column:menu_mode;default:'1'"` // 菜单模式，默认模式一
}

func registerCommand() {
	engine := control.Register("command", &control.Options{
		Alias: "菜单管理",
		Help: "权限:\n" +
			"仅限机器人管理员\n\n" +
			"指令:\n" +
			"* 设置菜单模式[1|2] -> 默认为模式1文本输出，模式2为网页输出(需要配置公网地址)",
	})

	if err := db.CreateAndFirstOrCreate("command", &command); err != nil {
		log.Fatalf("create command table failed: %v", err)
	}

	engine.OnRegex(`设置菜单模式([1-2])`, robot.AdminPermission).SetBlock(true).Handle(func(ctx *robot.Ctx) {
		menuMode := ctx.State["regex_matched"].([]string)[1]
		if err := db.Orm.Table("command").Where("1=1").Update("menu_mode", menuMode).Error; err != nil {
			ctx.ReplyTextAndAt("设置菜单模式失败")
			return
		}
		command.MenuMode = menuMode
		ctx.ReplyTextAndAt("设置菜单模式成功")
	})

	// 菜单输出
	engine.OnFullMatchGroup([]string{"menu", "菜单"}).SetBlock(true).Handle(func(ctx *robot.Ctx) {
		c := ctx.State["manager"].(*control.Control)
		options := MenuOptions{WxId: ctx.Event.FromUniqueID}
		for _, m := range c.Manager.M {
			if m.Options.HideMenu {
				continue
			}
			options.Menus = append(options.Menus, struct {
				Name      string `json:"name"`
				Alias     string `json:"alias"`
				Priority  uint64 `json:"priority"`
				Describe  string `json:"describe"`
				DefStatus bool   `json:"defStatus"`
				CurStatus bool   `json:"curStatus"`
			}{
				Name:      m.Service,
				Alias:     m.Options.Alias,
				Priority:  m.Options.Priority,
				Describe:  m.Options.Help,
				DefStatus: !m.Options.DisableOnDefault,
				CurStatus: m.IsEnabledIn(ctx.Event.FromUniqueID),
			})
		}

		if command.MenuMode == "" {
			command.MenuMode = "1"
		}

		switch command.MenuMode {
		case "1":
			// 🔔实现方案一(默认方案)：直接输出菜单
			menus := "当前支持的功能有: \n"
			for i := range options.Menus {
				menu := ""
				menu += "服务名: %s\n"
				menu += "别称: %s\n"
				menu += "默认开启状态: %v\n"
				menu += "当前开启状态: %v\n\n"
				menus += fmt.Sprintf(menu, options.Menus[i].Name, options.Menus[i].Alias, options.Menus[i].DefStatus, options.Menus[i].CurStatus)
			}
			ctx.ReplyTextAndAt(menus)
		case "2":
			// 🔔实现方案二：web输出菜单，需要在config.yaml中配置公网环境，否则打不开
			address := ctx.Bot.GetConfig().ServerAddress
			address = fmt.Sprintf("%s/menu?wxid=%s", address, ctx.Event.FromUniqueID)
			ctx.ReplyShareLink(ctx.Bot.GetConfig().BotNickname, "机器人当前所有的指令都在这里哦！", "https://imgbed.link/file/10160", address)
		}
	})
}

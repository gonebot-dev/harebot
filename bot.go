package main

import (
	onebotv11 "github.com/gonebot-dev/goneadapter-onebotv11"
	"github.com/gonebot-dev/gonebot"
	echo "github.com/gonebot-dev/goneplugin-echo"
	status "github.com/gonebot-dev/goneplugin-status"
)

func main() {
	gonebot.LoadPlugin(&echo.Echo)
	gonebot.LoadPlugin(&status.Status)
	gonebot.LoadAdapter(&onebotv11.OneBotV11)
	gonebot.Run()
}

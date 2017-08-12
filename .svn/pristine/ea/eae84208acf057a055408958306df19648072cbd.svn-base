package httpserver

import (
	"monitorweb/backend"
	"monitorweb/database"
	"monitorweb/model"
	"monitorweb/packets"

	"monitorweb/assets/public"
	"monitorweb/assets/templates"

	log "github.com/Sirupsen/logrus"
	"github.com/go-macaron/bindata"

	"gopkg.in/macaron.v1"
)

func StartHttpServer(mqtt *backend.Backend, db *database.DBEngine, mqttconf model.MqttEntity, addr string) {
	macaron.Env = macaron.PROD
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Map(mqtt)
	m.Map(db)
	m.Map(mqttconf)
	m.Use(macaron.Static("public",
		macaron.StaticOptions{
			FileSystem: bindata.Static(bindata.Options{
				Asset:      public.Asset,
				AssetDir:   public.AssetDir,
				AssetNames: public.AssetNames,
				AssetInfo:  public.AssetInfo,
				Prefix:     "",
			}),
		},
	))
	m.Use(macaron.Renderer(macaron.RenderOptions{
		TemplateFileSystem: bindata.Templates(bindata.Options{
			Asset:      templates.Asset,
			AssetDir:   templates.AssetDir,
			AssetNames: templates.AssetNames,
			AssetInfo:  templates.AssetInfo,
			Prefix:     "templates",
		}),
	}))

	//首页
	m.Get("/", appserver)
	m.Post("/queryrtdata", queryrtdata)

	//详情页面
	m.Get("/dtdata", dtdata)

	//亮灯
	m.Get("/open", openled)
	//灭灯
	m.Get("/close", closeled)

	m.RunAddr(addr)
}

//首页
func appserver(ctx *macaron.Context, mqttEntity model.MqttEntity) {
	log.Info("monitor first web")
	ctx.Data["mqtt"] = mqttEntity
	ctx.HTML(200, "rtdata")
}

//详情页面
func dtdata(ctx *macaron.Context, db *database.DBEngine) {
	mac := ctx.Query("mac")

	datalist, _ := db.QueryDetailData(mac)
	ctx.Data["list"] = datalist
	ctx.Data["mac"] = mac

	log.Info("mac = ", mac)
	//log.Info("list = ", datalist)
	ctx.HTML(200, "dtdata")

}

//POST server queryrtdata 查询最新实时数据
func queryrtdata(ctx *macaron.Context, db *database.DBEngine) {
	lastlist, _ := db.QueryAllLastData()
	//log.Info("lastlist = ", lastlist)
	ctx.JSON(200, lastlist)
}

//亮灯
// for test http://localhost:2017/open?mac=b827ebd212b7
func openled(ctx *macaron.Context, mqtt *backend.Backend) string {
	mac := ctx.Query("mac")
	var pkt packets.LedPacket
	pkt.LED = true
	pkt.Mac = mac
	mqtt.PublishPacket(pkt)
	return "Success"
}

//灭灯
//for test http://localhost:2017/close?mac=b827ebd212b7
func closeled(ctx *macaron.Context, mqtt *backend.Backend) string {
	mac := ctx.Query("mac")
	var pkt packets.LedPacket
	pkt.LED = false
	pkt.Mac = mac
	mqtt.PublishPacket(pkt)
	return "Success"
}

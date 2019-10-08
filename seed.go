package database

import (
	"log"
	"math/rand"
	"strconv"

	utils "banners.utils"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

func Seed() {
	db, err := sqlx.Open("sqlserver", ConnectionString)

	if err != nil {
		log.Fatal(err)
	} else {
		InitializeTables(*db)
	}
	defer db.Close()
}

func InitializeTables(connection sqlx.DB) {
	utils.ExecuteQuery(DestroyBannerTable, connection)
	utils.ExecuteQuery(DestroyZonesTable, connection)

	utils.ExecuteQuery(BannersSchema, connection)
	utils.ExecuteQuery(ZonesSchema, connection)

	utils.ExecuteQuery(StoredProcedureQuery, connection)

	PopulateZones(connection)
	PopulateBanners(connection)
}

func PopulateBanners(connection sqlx.DB) {
	for i := 0; i < 20; i++ {
		id := strconv.Itoa(i)
		zoneId := strconv.Itoa(rand.Intn(6))
		var s string
		s = "INSERT INTO [dbo].[banners]([zoneid], [id],[date],[background_color],[background_image],[button_text],[button_color],[button_background],[title],[description],[text_align],[link],[link_isExternal])VALUES(" + zoneId + ", " + id + ",'2019-09-13T08:00:00Z','bg_col', 'bg_img', 'btn_txt', 'btn_col', 'btn_bg', 'title', 'description', 'left', 'link', 'true')"
		utils.ExecuteQuery(s, connection)
	}
}

func PopulateZones(connection sqlx.DB) {
	var s string
	s = "INSERT INTO [dbo].[zones]([id],[deviceId],[pageId],[languageCode],[width],[height])VALUES(1, 1, 1, 'UK', 1920, 400)"
	utils.ExecuteQuery(s, connection)
	s = "INSERT INTO [dbo].[zones]([id],[deviceId],[pageId],[languageCode],[width],[height])VALUES(2, 1, 1, 'IT', 1920, 400)"
	utils.ExecuteQuery(s, connection)
	s = "INSERT INTO [dbo].[zones]([id],[deviceId],[pageId],[languageCode],[width],[height])VALUES(3, 2, 1, 'IT', 1920, 400)"
	utils.ExecuteQuery(s, connection)
	s = "INSERT INTO [dbo].[zones]([id],[deviceId],[pageId],[languageCode],[width],[height])VALUES(4, 1, 2, 'UK', 1920, 400)"
	utils.ExecuteQuery(s, connection)
	s = "INSERT INTO [dbo].[zones]([id],[deviceId],[pageId],[languageCode],[width],[height])VALUES(5, 1, 2, 'IT', 1920, 400)"
	utils.ExecuteQuery(s, connection)
}

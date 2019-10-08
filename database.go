package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	domain "banners.domain"
	utils "banners.utils"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

func SelectBanners(language string, pageId int, deviceId int) []domain.Banner {

	db, err := sqlx.Open("sqlserver", ConnectionString)

	var ErrorMessage string

	banners := []domain.Banner{}

	if err != nil {
		log.Fatal(err)
	} else {
		spResponseRows, spError := db.Queryx(ExecuteStoreProcedure, sql.Named("LanguageCode", language), sql.Named("PageId", pageId), sql.Named("DeviceTypeId", deviceId), sql.Named("ErrorMessage", &ErrorMessage))
		if spError != nil {
			log.Fatal(spError)
		} else {

			for spResponseRows.Next() {
				banner := domain.Banner{}
				err := spResponseRows.StructScan(&banner)
				if err != nil {
					log.Fatalln(err)
				} else {
					banners = append(banners, banner)
				}
			}
		}
	}
	defer db.Close()
	return banners
}

func SelectZones(banners []domain.Banner) []domain.Zone {
	ints := GetZonesIdsFromBanners(banners)
	zones := []domain.Zone{}
	for index := 0; index < len(ints); index++ {
		zone := SelectZone(ints[index])
		zoneBanners := GetBannersByZoneId(banners, ints[index])
		zoneValue := domain.NewZone(zone.ZoneId, zone.DeviceId, zone.PageId, zone.LanguageCode, zone.Width, zone.Height, zoneBanners)
		zones = append(zones, zoneValue)
	}
	return zones
}

func SelectZone(zoneId int) domain.Zone {
	db, err := sqlx.Open("sqlserver", ConnectionString)
	var returnValue domain.Zone
	zoneIdString := strconv.Itoa(zoneId)
	if err != nil {
		log.Fatal(err)
	}
	zones := []domain.Zone{}
	query := "SELECT * FROM zones WHERE id = " + zoneIdString

	selectError := db.Select(&zones, query)
	if selectError != nil {
		fmt.Printf(selectError.Error() + "\n")
	} else {
		returnValue = zones[0]
	}
	defer db.Close()

	return returnValue
}

func GetZonesIdsFromBanners(banners []domain.Banner) []int {
	var ints []int
	for index := 0; index < len(banners); index++ {
		if !utils.Contains(ints, banners[index].ZoneId) {
			ints = append(ints, banners[index].ZoneId)
		}
	}
	return ints
}

func GetBannersByZoneId(banners []domain.Banner, zoneId int) []domain.Banner {
	returnValue := []domain.Banner{}
	for index := 0; index < len(banners); index++ {
		if banners[index].ZoneId == zoneId {
			returnValue = append(returnValue, banners[index])
		}
	}
	return returnValue
}

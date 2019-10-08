package database

var ConnectionString = "Server=SERVER_NAME;Database=DATABASE_NAME;Trusted_Connection=True;"

var DestroyBannerTable = "IF OBJECT_ID('dbo.banners', 'U') IS NOT NULL DROP TABLE dbo.banners"
var DestroyZonesTable = "IF OBJECT_ID('dbo.zones', 'U') IS NOT NULL DROP TABLE dbo.zones"

var BannersSchema = "CREATE TABLE banners(zoneId INT, id INT, date NVARCHAR(50), background_color NVARCHAR(50), background_image  NVARCHAR(50), button_text NVARCHAR(50),button_color NVARCHAR(50),button_background NVARCHAR(50),title NVARCHAR(50),description NVARCHAR(50),text_align NVARCHAR(50),link NVARCHAR(50),link_isExternal NVARCHAR(50))"
var ZonesSchema = "CREATE TABLE zones(id INT, deviceId INT, pageId INT, languageCode NVARCHAR(50), width INT, height INT)"

var ExecuteStoreProcedure = "EXEC cms.spGetBanners @LanguageCode, @PageId, @DeviceTypeId, @ErrorMessage OUTPUT;"

var StoredProcedureQuery = `CREATE OR ALTER PROCEDURE cms.spGetBanners
	
	@LanguageCode VARCHAR(30),
	@PageId INT,
	@DeviceTypeId INT,
	@ErrorMessage VARCHAR(1000) output
	
	AS
	 BEGIN

	 IF EXISTS (SELECT * FROM zones Z WHERE Z.languageCode = @LanguageCode AND Z.pageId = @PageId AND Z.deviceId = @DeviceTypeId)
	 BEGIN
		 SET NOCOUNT ON;
		 SELECT
		 [zoneId]
		,[date]
		,[background_color]
		,[background_image]
		,[button_text]
		,[button_color]
		,[button_background]
		,[title]
		,[description]
		,[text_align]
		,[link]
		,[link_isExternal]
		 FROM banners B
		 INNER JOIN zones Z ON B.zoneId = Z.id
		 WHERE Z.languageCode = @LanguageCode AND Z.deviceId = @DeviceTypeId AND Z.pageId = @PageId;
		 END 
		 ELSE 
		 BEGIN
		 SET @ErrorMessage = 'Required data not here.'
		 END
	 END;`

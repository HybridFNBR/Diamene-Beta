package addons

import (
	"github.com/gin-gonic/gin"
)

// DO NOT EDIT THE FUNC NAME OR THE PACKAGE NAME OR IT WILL NOT WORK
func CustomRoutes(r *gin.Engine) {
	r.GET("/api/inventory/v3/:deploymentId/players/:playerId/POIDiscoveryPlayerPersistence", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"binary": nil,
			"inventory": gin.H{
				"playerId":      "00027b91959a4c57a1272efcc4d7480f",
				"inventoryName": "poidiscoveryplayerpersistence",
				"prefix":        "/",
				"versionId":     nil,
				"instance":      "4e4a7ec1-9e79-4d22-872d-2678d0867437",
				"contents": gin.H{
					"/POIDiscovery_PlayerState":      "{\\\"StructName\\\":\\\"POIDiscoveryPersistenceData\\\",\\\"StructData\\\":{\\\"mapStates\\\":{\\\"hera_V2_Terrain\\\":{\\\"discoveredPOIs\\\":[{\\\"tagName\\\":\\\"Athena.Location.POI.Generic.34\\\"},{\\\"tagName\\\":\\\"Athena.Location.UnNamedPOI.Landmark.62\\\"},{\\\"tagName\\\":\\\"Athena.Location.POI.Generic.30\\\"},{\\\"tagName\\\":\\\"Athena.Location.POI.Generic.36\\\"},{\\\"tagName\\\":\\\"Athena.Location.UnNamedPOI.Landmark.80\\\"},{\\\"tagName\\\":\\\"Athena.Location.UnNamedPOI.Landmark.58\\\"},{\\\"tagName\\\":\\\"Athena.Location.UnNamedPOI.Landmark.55\\\"}],\\\"bMigrationComplete\\\":false,\\\"bMilestoneGranted\\\":true,\\\"version\\\":1},\\\"athena_Terrain_S09\\\":{\\\"discoveredPOIs\\\":[{\\\"tagName\\\":\\\"Athena.Location.POI.FatalFields\\\"},{\\\"tagName\\\":\\\"Athena.Location.POI.SaltySprings\\\"},{\\\"tagName\\\":\\\"Athena.Location.POI.PleasantPark\\\"},{\\\"tagName\\\":\\\"Athena.Location.POI.DustyDivot\\\"},{\\\"tagName\\\":\\\"Athena.Location.POI.LootLake\\\"}],\\\"bMigrationComplete\\\":true,\\\"bMilestoneGranted\\\":true,\\\"version\\\":1}}}}",
					"/POIDiscovery_PlayerState.meta": "{\\\"version\\\":2,\\\"serializer\\\":\\\"UStructAsJsonString\\\",\\\"buildid\\\":\\\"\\\",\\\"timestamp\\\":1790631312,\\\"metadata\\\":{}}",
				},
				"types": gin.H{},
			},
			"continuationToken": nil,
		})
	})
}

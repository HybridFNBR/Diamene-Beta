package timeline

import (
	"Diamene/managers"
	"Diamene/types"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func Timeline(r *gin.Engine) {
	r.GET("/fortnite/api/calendar/v1/timeline", func(c *gin.Context) {
		var calendar types.FCalendarDownload
		versionInfo := managers.GetVersionInfo(c)

		timeline, _ := os.ReadFile("./responses/timeline/timeline.json")
		_ = json.Unmarshal(timeline, &calendar)

		keychainFile, _ := os.ReadFile("./responses/keychain.json")
		_ = json.Unmarshal(keychainFile, &managers.Keychain)

		clientEvents := calendar.Channels["client-events"]

		managers.LoadConfig()
		if managers.Config.GeneralSettings.BRefreshKeychainViaTimeline == true {
			managers.CompareAndUpdateKeychain()
		}

		var CertVersions = map[int]int{
			3700114: 1,
			3724489: 1,
			3729133: 1,
			3741772: 1,
			3757339: 1,
			3775276: 1,
			3790078: 1,
			3807424: 2,
			3825894: 2,
			3841827: 2,
			3847564: 2,
			3858292: 2,
			3870737: 2,
			3889387: 2,
		}
		if cl, ok := CertVersions[versionInfo.VersionLegacy]; ok {
			versionInfo.VersionGlobal = cl
		}

		calendar.Channels["client-events"].States[0].ActiveEvents = append(calendar.Channels["client-events"].States[0].ActiveEvents,
			types.FEventRecord{
				EventType:   fmt.Sprintf("EventFlag.Season%d", versionInfo.VersionGlobal),
				ActiveUntil: "9999-12-01T21:10:00.000Z",
				ActiveSince: "2020-11-21T07:00:00.000Z",
			},
			types.FEventRecord{
				EventType:   fmt.Sprintf("EventFlag.LobbySeason%d", versionInfo.VersionGlobal),
				ActiveUntil: "9999-12-01T21:10:00.000Z",
				ActiveSince: "2020-11-21T07:00:00.000Z",
			},
			types.FEventRecord{
				EventType:   "GGL05",
				ActiveUntil: time.Now().UTC().Add(60 * time.Second).Format("2006-01-02T15:04:05.000Z"),
				ActiveSince: time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
			},
			types.FEventRecord{
				EventType:   "Slab1",
				ActiveUntil: "9999-12-01T21:10:00.000Z",
				ActiveSince: "2020-11-21T07:00:00.000Z",
			},
			types.FEventRecord{
				EventType:   "Slab2",
				ActiveUntil: "9999-12-01T21:10:00.000Z",
				ActiveSince: "2020-11-21T07:00:00.000Z",
			},
			types.FEventRecord{
				EventType:   "Slab3",
				ActiveUntil: "9999-12-01T21:10:00.000Z",
				ActiveSince: "2020-11-21T07:00:00.000Z",
			},
			types.FEventRecord{
				EventType:   "heard0stone78hole65stick178feet", // Comic Reader
				ActiveUntil: "9999-12-01T21:10:00.000Z",
				ActiveSince: "2020-11-21T07:00:00.000Z",
			},
			types.FEventRecord{
				EventType:   "EventFlag.Event_Osiris",
				ActiveUntil: "9999-12-01T21:10:00.000Z",
				ActiveSince: "2020-11-21T07:00:00.000Z",
			},
		)
		var versionEventFlags = map[string][]string{
			"Cert": {
				"EventFlag.LobbyWinterDecor",
			},

			"4.5": {
				"EventFlag.BR_S4_Geode_Countdown",
				"EventFlag.BR_S4_Oddity_01_Event",
				"EventFlag.BR_S4_Oddity_02_Event",
				"EventFlag.BR_S4_Oddity_03_Event",
				"EventFlag.BR_S4_Oddity_04_Event",
				"EventFlag.BR_S4_Oddity_05_Event",
				"EventFlag.BR_S4_Oddity_06_Event",
				"EventFlag.BR_S4_Oddity_07_Event",
				"EventFlag.BR_S4_Oddity_08_Event",
				"EventFlag.BR_S5_Oddity_Tomato_Event",
				"EventFlag.BR_S4_Rift_Growth",
			},

			"5.10": {
				"EventFlag.BR_S5_Rift_Corrupt",
				"EventFlag.BR_S5_Rift_Shrink",
				"EventFlag.Anniversary2018_BR",
				"EventFlag.BirthdayBattleBus",
				"EventFlag.SupplyDropBDay",
				"EventFlag.BR_S5_RiskyReels_Event",
			},

			"5.40": {
				"EventFlag.BR_S5_RiskyReels_Event",
				"EventFlag.BR_S5_Sphere_Spawn",
				"EventFlag.BR_S5_Cube_TurnOn",
				"EventFlag.BR_S5_Cube_Lightning",
				//"EventFlag.BR_S5_Cube_MoveTo1",
				//"EventFlag.BR_S5_Cube_Rune1",
				//"EventFlag.BR_S5_Cube_MoveTo2",
				//"EventFlag.BR_S5_Cube_Rune2",
				//"EventFlag.BR_S5_Cube_MoveTo3",
				//"EventFlag.BR_S5_Cube_Rune3",
				//"EventFlag.BR_S5_Cube_MoveTo4",
				//"EventFlag.BR_S5_Cube_Rune4",
				//"EventFlag.BR_S5_Cube_MoveTo5",
				//"EventFlag.BR_S5_Cube_Rune5",
				//"EventFlag.BR_S5_Cube_MoveTo6",
				//"EventFlag.BR_S5_Cube_Rune6",
				//"EventFlag.BR_S5_Cube_MoveTo7",
				//"EventFlag.BR_S5_Cube_Rune7",
				//"EventFlag.BR_S5_Cube_MoveTo8",
			},

			"5.40|5.41": {
				"EventFlag.LTM_Heist",
			},

			"6.00": {
				//"EventFlag.BR_S6_Island_MoveTo1"
				//"EventFlag.BR_S6_Island_Rune1",
				//"EventFlag.BR_S6_Island_MoveTo2"
			},

			"6.01": {
				//"EventFlag.BR_S6_Island_Rune2"
				//"EventFlag.BR_S6_Island_MoveTo3"
				//"EventFlag.BR_S6_Island_Rune3",
				//"EventFlag.BR_S6_Island_MoveTo4"
			},

			"6.02": {
				//"EventFlag.BR_S6_Island_Rune4"
				//"EventFlag.BR_S6_Island_MoveTo5"
				//"EventFlag.BR_S6_Island_Rune5",
				//"EventFlag.BR_S6_Island_MoveTo6",
			},

			"6.10": {
				"EventFlag.BR_S6_Island_Beacon",
				//"EventFlag.BR_S6_Island_MoveTo7",
				//"EventFlag.BR_S6_Island_Rune7",
				//"EventFlag.BR_S6_Island_MoveTo8"
				//"EventFlag.BR_S6_Island_Rune8",
			},

			"6.20|6.21": {
				/* Halloween EventFlags */
				"HW_01",
				"HW_02",
				"HW_03",
				"HW_04",
				"HW_05",
				"HW_06",
				"HW_07",
				"EventFlag.LobbySeason6Halloween",
				"EventFlag.LTM_Fortnitemares",
				"EventFlag.LTM_LilKevin",
			},

			"6.22": {
				"EventFlag.LTM_LilKevin",
			},

			"7.10": {
				"EventFlag.LTM_14DaysOfFortnite", //WinterFest
				"HLDY_01",
				"HLDY_02",
				"HLDY_03",
				"HLDY_04",
				"HLDY_05",
				"HLDY_06",
				"HLDY_07",
				"R0",
				"TopSecret",
				"SP",
			},

			"7.20": {
				"EventFlag.LTM_14DaysOfFortnite", //WinterFest
				"warning",                        //TV Warning
				"Countdown",                      //TV Countdown
				"P1",                             // Ice King Event related
				"WDK",                            // Ice King Apeaars
				"EventFlag.LTM_WinterDeimos",
			},

			"7.30": {
				"F0",                     // Marshmello Event countdown
				"F1",                     // Marshmello Event countdown
				"FLevel",                 // Marshmello Event Level
				"FEST_POSTER",            // Marshmello event posters
				"roadie_0",               // ?
				"EventFlag.LTE_Festivus", // Marshmello event
				"S7_B",                   // ?
				"S7_C",                   // ?
				"HP",                     // ?
			},

			"7.40": {
				"EventFlag.LTE_S7_OverTime",
			},

			"8.10": {
				"EventFlag.LTM_HighStakes", // Event_HighStakes
			},

			"8.30": {
				"EventFlag.LTE_BootyBay", // Event_S8_BootyBay
			},

			"8.40": {
				"EventFlag.LTM_Goose",
			},

			"8.50": {
				"EventFlag.LTE_S8_ExtraCredit",
				"EventFlag.LTM_Ashton", // Event_Ashton
			},

			"8.51": {
				"EventFlag.LTM_Ashton", // Event_Ashton
				"EventFlag.LTE_S8_ExtraCredit",
				"EventFlag.UnvaultingCountdown", // Unvaulting Event countdown
				"GWS0",                          // Unvaulting Event related
				"GWS1",                          //  Unvaulting Event related
				"GW1",                           //  Unvaulting Event related
				"GW2",                           //  Unvaulting Event related
				"GW3",                           // Unvaulting Event related
				"GWW1",                          //  Unvaulting Event related
				"GWW2",                          //  Unvaulting Event related
				"GWW3",                          //  Unvaulting Event related
				"GWW4",                          //  Unvaulting Event related
				"GWW5",                          //  Unvaulting Event related
			},

			"9.01": {
				"Cat0",
				"EventFlag.LTM_Wax",
			},

			"9.10": {
				"Cat1",
				"EventFlag.LTM_Vigilante",
			},

			"9.20": {
				"CatSpot1", // Cattus Spotting Ocean
				"EventFlag.LTM_Mash",
			},

			"9.21": {
				"MAS1", // MonsterAttackSmall1
				"MAS2", // MonsterAttackSmall2
				"EventFlag.LTM_Mash",
			},

			"9.30|9.40|9.41": {
				"CDTime",     // Final Showdown countdown
				"DoggusProp", //Doggus Propaganda Posters
				"CVD1",       // Final Showdown countdown
				"CatSpot2",   // Cattus Spotting 2
				"STP",        // Stranger Things Portals(Mega Mall)
			},

			"9.30": {
				"EventFlag.LTM_14DaysOfSummer",
			},

			"9.40": {
				"EventFlag.LTE_S9_Overtime",
				//"EventFlag.Anniversary2019_BR", // Anniversary 2019 (BR)
				"EventFlag.LTE_WorldCup", // Event_WorldCup
				"EventFlag.WorldCupBattleBus",
				//"EventFlag.SupplyDropBDay",     // Anniversary 2019 (BR)
				//"EventFlag.BirthdayBattleBus2nd",
			},

			"9.41": {
				"EventFlag.LTE_S9_Overtime",
				"EventFlag.Anniversary2019_BR", // Anniversary 2019 (BR)
				"EventFlag.SupplyDropBDay",     // Anniversary 2019 (BR)
				"EventFlag.BirthdayBattleBus2nd",
				//"EventFlag.WorldCupBattleBus",
			},

			"10.00|10.10|10.20|10.30|10.31|10.40": {
				"EventFlag.Season10_UrgentMission_1",
				"EventFlag.Season10_UrgentMission_2",
				"EventFlag.Season10_UrgentMission_3",
				"EventFlag.Season10_UrgentMission_4",
				"EventFlag.Season10_UrgentMission_5",
				"EventFlag.Season10_UrgentMission_6",
				"EventFlag.Season10_UrgentMission_7",
				"EventFlag.Season10_UrgentMission_8",
				"EventFlag.Season10_UrgentMission_9",
				"EventFlag.Season10_UrgentMission_10",
			},

			"10.00": {
				"VLB_1",      // Visitor-Home Depot Setup
				"VLB_2",      // Visitor-Home Depot Junk Collection
				"VLB_3",      // Visitor-Home Depot Lab Complete
				"VLB_BCN_01", // Visitor-Home Depot Beacon 1 Assembly
			},

			"10.10": {
				"VLB_BCN_02", // Visitor-Home Depot Beacon 2 Assembly
				"VLB_BCN_03", // Visitor-Home Depot Beacon 3 Assembly
			},

			"10.20": {
				"EventFlag.S10_Oak",
				"ITB",        // IT Balloon
				"VLB_BCN_04", // Visitor-Home Depot Beacon 4 Assembly
				"VLB_BCN_05", // Visitor-Home Depot Beacon 5 Assembly
				"VLB_MCH_01", // Visitor-Home Depot Mech Scanning 1
				"VLB_MCH_02", // Visitor-Home Depot Mech Scanning 2
				"VLB_MCH_03", // Visitor-Home Depot Mech Scanning 3
				"VLB_MCH_04", // Visitor-Home Depot Mech Scanning 4
				"VLB_MCH_05", // Visitor-Home Depot Mech Scanning 5
			},

			"10.30": {
				"EventFlag.S10_Mystery",
				"VLB_BCN_06", // Visitor-Home Depot Beacon 6 Assembly
				"VLB_MCH_06", // Visitor-Home Depot Mech Scanning 6
				"VLB_MCH_07", // Visitor-Home Depot Mech Scanning 7
				"VLB_MCH_08", // Visitor-Home Depot Mech Scanning 8
				"VLB_MCH_09", // Visitor-Home Depot Mech Scanning 9
				"VLB_MCH_10", // Visitor-Home Depot Mech Scanning 10
				"VLB_MCH_11", // Visitor-Home Depot Mech Scanning 11
				"VLB_MCH_12", // Visitor-Home Depot Mech Scanning 12
				"VLB_RKT_01", // Visitor-Home Depot Rocket 1 Assembly
				"VLB_RKT_02", // Visitor-Home Depot Rocket 2 Assembly
				"VLB_RKT_03", // Visitor-Home Depot Rocket 3 Assembly
				"VLB_RKT_04", // Visitor-Home Depot Rocket 4 Assembly
				"VLB_RKT_05", // Visitor-Home Depot Rocket 5 Assembly
				"VLB_RKT_06", // Visitor-Home Depot Rocket 6 Assembly
				"VLB_RKT_07", // Visitor-Home Depot Rocket 7 (Thusters) Assembly
			},

			"10.31|10.40": {
				"EventFlag.LTE_BlackMonday",
				"EventFlag.S10_Mystery",
				"RBBMS", // Black Monday Tilted POI Swap
				//"survey_br_nick_001", // nightnight blackhole lobby mini event
			},

			"10.40": {
				"EventFlag.LTE_S10_Overtime",
				"EventFlag.LTM_ZebraWallet",
				"RKT_Final",
				"NN1",
				//"survey_br_nick_001", // nightnight blackhole lobby mini event
			},

			"11.10": {
				"EventFlag.LTM_Hyena",
				"EventFlag.LTE_Fortnitemares2019",
				"COL_01",                       // ColumbusLoadMap
				"EventFlag.HalloweenBattleBus", // Event_Fortnitemares_BattleBus
				"HW2019",                       // FortniteMares 2019
				"HW2019Leadup",                 // FortniteMares 2019 Leadup
			},

			"11.20": {
				"GALP", // Galileo Promo
				//"Risky_Default", // Gal - Risky Reels Default Foundation
				"WP1", // WinterSnowFXLight
			},

			"11.30": {
				"EventFlag.LTE_CoinCollectXP",
				"EventFlag.LTE_Galileo",
				"Risky_Default",
				"GALCD",                    // C2S1-related (Winterfest 2019)
				"GALP",                     // C2S1-related (Winterfest 2019)
				"EventFlag.Winterfest",     // Winterfest_Tree
				"EventFlag.LTE_WinterFest", // Event_WinterFest
				"TopSecret",                // Holiday_BattleBus_2019
				"WD1",                      // Winter Decorations 1
				"WD1C",                     // Winter Decorations 1 Client
				"WD2",                      // Winter Decorations 2
				"WD3",                      // Winter Decorations 3
				"WD4",                      // Winter Decorations 4
				"WD5",                      // Winter Decorations 5
				"WDT1",                     // Winter Delivery Trucks 1
				"WDT2",                     // Winter Delivery Trucks 2
				"WDT3",                     // Winter Delivery Trucks 3
				"WDT4",                     // Winter Delivery Trucks 4
				"WRRD1",                    // Winter Retail Row Store Deco 1
				"WSFX2",                    // WinterSnowFXheavy
				"WSFX3",                    // WinterSnowFXstorms
				"WSFX3",                    // WinterSnowFXstorms2
				"WSFX3F",                   // WinterSnowFall Stage 3
				"WSH1",                     // Winter Salty House Update
				"WSL0",                     // Winter Snow Level 0
				"WSL1",                     // Winter Snow Level 1
				"WSL2",                     // Winter Snow Level 2
				"WSL3",                     // Winter Snow Level 3
				"WSM1",                     // Winter Snowmen 1
				"WSM2",                     // Winter Snowmen 2
				"WSM3",                     // Winter Snowmen 3
				"WSM4",                     // Winter Snowmen 4
				"WTR1",                     // Winter Trees 1
				"WTR2",                     // Winter Trees 2
				"WTR3",                     // Winter Trees 3
				"WTR4",                     // Winter Trees 4
			},

			"11.31": {
				"EventFlag.Winterfest",     // Winterfest_Tree
				"EventFlag.LTE_WinterFest", // Event_WinterFest
				"TopSecret",                // Holiday_BattleBus_2019
				"WD1",                      // Winter Decorations 1
				"WD1C",                     // Winter Decorations 1 Client
				"WD2",                      // Winter Decorations 2
				"WD3",                      // Winter Decorations 3
				"WD4",                      // Winter Decorations 4
				"WD5",                      // Winter Decorations 5
				"WDT1",                     // Winter Delivery Trucks 1
				"WDT2",                     // Winter Delivery Trucks 2
				"WDT3",                     // Winter Delivery Trucks 3
				"WDT4",                     // Winter Delivery Trucks 4
				"WRRD1",                    // Winter Retail Row Store Deco 1
				"WSFX2",                    // WinterSnowFXheavy
				"WSFX3",                    // WinterSnowFXstorms
				"WSFX3",                    // WinterSnowFXstorms2
				"WSFX3F",                   // WinterSnowFall Stage 3
				"WSH1",                     // Winter Salty House Update
				"WSL0",                     // Winter Snow Level 0
				"WSL1",                     // Winter Snow Level 1
				"WSL2",                     // Winter Snow Level 2
				"WSL3",                     // Winter Snow Level 3
				"WSM1",                     // Winter Snowmen 1
				"WSM2",                     // Winter Snowmen 2
				"WSM3",                     // Winter Snowmen 3
				"WSM4",                     // Winter Snowmen 4
				"WTR1",                     // Winter Trees 1
				"WTR2",                     // Winter Trees 2
				"WTR3",                     // Winter Trees 3
				"WTR4",                     // Winter Trees 4
			},

			"11.50": {
				"EventFlag.LTM_Sourdropper", // Event_Sourdropper
				"GP1",                       // Gold Phase 1
				"GP2",                       // Gold Phase 2
				"GP3",                       // Gold Phase 3
				"GP4",                       // Gold Phase 4
				"GP5",                       // Gold Phase 5
			},

			"12.20": {
				"EventFlag.LTE_TechRelease_12.20_HF", // Event_TechRelease_12.20_HF
			},

			"12.21": {
				"EventFlag.LTE_Oro", // Event_Oro
			},

			"12.30": {
				"DPY01",                        // Donut Yacht
				"EventFlag.Donut",              // Donut SupplyDrop
				"EventFlag.DonutBattleBus",     // Donut BattleBus
				"EventFlag.LTE_TechRelease_12", // Event_TechRelease_12.30_HF
				"FEC01",                        // FE Cables 1
				"FEC02",                        // FE Cables 2
			},

			"12.40|12.41": {
				"EventFlag.LTE_JerkyChallenges",
				"FEC01", // FE Cables 1
				"FEC02", // FE Cables 2
				"FEC03", // FE Cables 3
				"JCD01", // Jerky Countdown Timer
				"JCH01", // Jerky Countdown Hold 1
				"JH01",  // Jerky Heads Appear
				"JL1",   // Jerky Leadup 1
				"JL2",   // Jerky Leadup 2
				"JLL",   // Jerky LoadLevel
				"JP01",  // Jerky Posters Appear
				"JP02",  // Jerky Posters Swap
			},

			"12.60|12.61": {
				"EventFlag.LTE_StormTheAgency",
				"FEC01",  // FE Cables 1
				"FEC02",  // FE Cables 2
				"FEC03",  // FE Cables 3
				"FCD01",  // Fritter Countdown Timer
				"FHS01",  // Hatches Stage 1
				"FHS02",  // Hatches Stage 2
				"FLA01",  // Fritter Lobby Device Active
				"FLCD01", // Fritter Lobby Countdown
				"FLL01",  // Fritter Load Level
			},

			"13.40": {
				"SM1",     //scorch marks
				"LS0",     // Landmark Sphere 0
				"LS1",     // Landmark Sphere 1
				"LS2",     // Landmark Sphere 2
				"LS3",     // Landmark Sphere 3
				"LS_POST", // LS_POST
				"TC1",     //thorhammer 'poi'
				"TR1",     //rift in the sky
				"CBS1",    //Coral buddies stage 1
			},

			"14.00": {
				"SPJ01", //"Sky Sparkle for Junior
				"HSS0",  // Hightower Sphere Stage 0
				"HSS1",  // Hightower Sphere Stage 1
				"HSS2",  // Hightower Sphere Stage 2
				"HSS3",  // Hightower Sphere Stage 3
				"RLS0",  // Rift Landmark Sphere 0
				"RLS1",  // Rift Landmark Sphere 1
				"RLS2",  // Rift Landmark Sphere 2
				"RLS3",  // Rift Landmark Sphere 3
				"SL_01", // Stark Lab - SL_01
				"SL_02", // Stark Lab - SL_02

				"LS2",  //Galactus Countdown(Ch2S4 related)
				"HSS2", //Sky Distant Object(Ch2S4 related)
			},

			"14.20": {
				"EventFlag.Anniversary2020_BR",
				"EventFlag.BirthdayBattleBus3rd", //Event_Birthday_BattleBus3
				"SPJ01",                          //Sky Sparkle for Junior
				"BE01",                           // Birthday Event
				"BVJ01",                          // Begin visible in sky for Junior
				"EventFlag.BusUpgrade1",          // Stark Bus Mod Stage 1
				"HC_01",                          // Hellicarrier - Bus Mod
				"HC_02",                          // Hellicarrier - Mod Stage1
				"SH01",                           // Stark Hangar - SH01
				"SH02",                           // Stark Hangar - SH02
				"TE01",                           // Turbo Event
				"TE_NE",                          // TurboNotLive Event
			},

			"14.30": {
				"EventFlag.BusUpgrade1",
				//"SPJ02", //Sky Card for Junior(Ch2S4 related)
			},

			"14.40": {
				//"SPJ02",
				"EventFlag.HalloweenBattleBus",     // Event_Fortnitemares_BattleBus
				"EventFlag.LTE_Fortnitemares_2020", // Event_Fortnitemares_Challenges
				"Yrk1",                             // York - 01
			},

			"14.50": {
				"EventFlag.BusUpgrade2",
				"BBMS2",
			},

			"14.60": {
				"GALCD",
				"EventFlag.BusUpgrade3",
				"BBMS3", // Battle Bus Swap - Stage 3
				"FLA01", // C2S4-related (Timers)
				"FSE01", // C2S4-related (Timers)
				//"FSE02", // C2S4-related (Timers)
				"JSE02", // C2S4-related (Timers)
				"JFE01", // C2S4-related (Countdown)
				//"SPJ03", // C2S4-related (Sky Head For Junior)

			},

			"15.00": {
				"NewFire1", // Campfire New 1
				"NewFire2", // Campfire New 2
				"OldFire1", // Campfire Old 1
			},

			"15.10:": {
				"Wtr20Dec1",  // Winter Decorations 1
				"Wtr20Tree1", // Holiday Trees 1
				"Wtr20Dec2",  // Winter Decorations 2
				"Wtr20Dec3",  // Winter Decorations 3
				"Wtr20Tree2", // Holiday Trees 2
				"Wtr20Tree3", // Holiday Trees 3
				"WSL1",       //Winter Snow Level 1
				"WSL2",       //Winter Snow Level 2
				"WSL3",       //Winter Snow Level 3
				"HE01",       // Snowmando Outposts
				"HE02",       // Remaining Snowmando Outpost
				"NewFire3",   // Campfire New 3
				"NewFire4",   // Campfire New 4
				"OldFire2",   // Campfire Old 2
				"OldFire3",   // Campfire Old 3
				"TopSecret",  // Holiday_BattleBus_2020
			},

			"15.20": {
				"NewFire5", // Campfire New 5
				"OldFire4", // Campfire Old 4
			},

			"16.00": {
				"RPFS1", // Dino Egg Cracking
			},

			"16.10": {
				"FTF01", // Fortified POI: Misty Meadows
				"FTF02", // Fortified POI: Sweaty Sands
				"FTF03", // Fortified POI: Retail Row
			},

			"16.50": {
				"EventFlag.Foreshadowing_Quest_03",    // S16 Foreshadowing Quest 3
				"ADP07",                               // 16.50 Foreshadowing Abductions
				"EventFlag.Foreshadowing_Quest_05",    // S16 Foreshadowing Quest 5
				"EventFlag.Foreshadowing_Quest_02",    // S16 Foreshadowing Quest 2
				"EventFlag.Foreshadowing_Quest_01",    // S16 Foreshadowing Quest 1
				"CC03",                                // Crop Circle 03
				"CC02",                                // Crop Circle 02
				"CC01",                                // Crop Circle 01
				"EventFlag.Event_NBA_Challenges",      // Event_NBA_Challenges
				"EventFlag.LTQ_S16_Legendary_Week_11", // S16 Legendary Quest Week 11
			},
			"17.10": {
				"AVOC1",           // Egg spawn 1
				"AVOC3",           // Egg spawn 3
				"Bush_Spawn_On",   // Bush_Spawn_On
				"Faux_BB",         // Faux_BB
				"Faux_DD",         // Faux_DD
				"GRAN01",          // Mural Stage 1
				"GRAN02",          // Mural Stage 2
				"IOCC",            // IO_Camp Crater
				"IOCH",            // IO_Camp Holly
				"IOCS",            // IO_Camp Sweaty
				"PMVE",            // Mothership Patrol
				"STL01",           // Holly Settlement House 1
				"STL02",           // Holly Settlement House 2
				"STL03",           // Holly Settlement House 3
				"SWPLSA",          // Sweaty Party Landingship
				"SummerHenchTime", // SummerHenchTime
				"WLCM1",           // Welcome Party full-swing
			},

			"17.20": {
				"AVOC4",  // Egg spawn 4
				"AVOC5",  // Egg spawn 5
				"FWCS",   // Foray !!!
				"IOCCLS", // IO_Camp Crater Phase Two
				"PMVE",   // Mothership Patrol
			},

			"17.20|17.21|17.30": {
				"DPOQ",  // Bushranger grave appears (RIP)
				"BEL01", // Leadup In-game/Leadup In-Lobby
				"BET01", // Rift Tour-related (Countdown)
				"BEL02", // Leadup In-game/Leadup In-Lobby
				"BEP02", // Rift Tour Enable Pads?
				"BPLS",  // Rift Tour Party Landingship
				"BTP",   // Rift Tour Teaser Poster
				"BAP",   // Rift Tour Announce Poster
				"BTL01", // Rift Tour TOD Leadup 1)
				"RL04",  // RocketLeague
			},

			"17.40": {
				//"ABDCRL", // Coral: Abduction
				"BPAQ", // Bomb Parts Assembling
				//"FMVE",   // Mothership Move To: Farm
				"SKFD", // Skyfire Disabled Week 11 17.40
			},

			"17.50": {
				"KEL01",  // Operation Sky Fire Event-related
				"KEL02",  // Operation Sky Fire Event-related
				"ABDFRM", //Farm: Abduction(Ch2S7 related)
			},

			"18.00": {
				"EventFlag.BirthdayBattleBus4th",
				"FNBD21",
				"Tandem_Active_W", // Tandem_Active_W
				"WEBS01",          // War Effort Cool Stations
				"WENW01",          // War Effort New Weapons
				"WET01",           // War Effort Turrets
				"WEUV01",          // War Effort Unvault 1
				"WEUV02",          // War Effort Unvault 2
				"WEUV03",          // War Effort Unvault 3
			},

			"18.10": {
				"CNV_04",                    // Convoy Stage 4
				"CNV_05",                    // Convoy Stage 5
				"CNV_06",                    // Convoy Stage 6
				"EventFlag.Event_SoundWave", // Event_SoundWave
				"SWSP",                      // Soundwave Series Poster
				"Tandem_Active_BM",          // Tandem_Active_BM
				"Tandem_Active_GF",          // Tandem_Active_GF
				"Tandem_Active_N",           // Tandem_Active_N
				"Tandem_Active_RAV",         // Tandem_Active_RAV
			},

			"18.20": {
				"EventFlag.HalloweenBattleBus", // Event_Fortnitemares_BattleBus_2021
				"FNM21",                        // FNMares 2021
				"CNV_07",                       // Convoy Stage 7
				"CNV_08",                       // Convoy Stage 8
				"CNV_09",                       // Convoy Stage 9
			},

			"18.21": {
				"EventFlag.HalloweenBattleBus", // Event_Fortnitemares_BattleBus_2021
				"FNM21",                        // FNMares 2021
				//"CT01",                         // Cube Town Phase 1
			},

			"18.30": {
				//"CT02",  // Cube Town Phase 2
				//"CT03",  // Cube Town Phase 3
				//"CTQ03", // Cube Town Queen Phase 3
			},

			"18.40": {
				"CTQ04",                 // Cube Town Queen Phase 4
				"EventFlag.HeadbandBus", // Event_HeadBand_BattleBus
				"GF01",                  // Guava Fort 1
				"GF02",                  // Guava Fort 2
				"GF03",                  // Guava Fort 3
				"GGL01",                 // Guava In-Game Leadup
				"GGL02",                 // Guava Lobby Leadup
				"GNOM_01",               // Gnomes Embark
			},

			"19.01": {
				"IODS_1",                   // IO_Drill 01
				"IODS_2",                   // IO_Drill 02
				"EventFlag.LTE_WinterFest", // WinterFest 2021
				"WF_IG_Avail",              // WinterFest 2021
				"2021_SM02",                // Snow Stage 02
				"2021_SM03",                // Snow Stage 03
				"2021_SM04",                // Snow Stage 04
				"2021_SM05",                // Snow Stage 05
				"AHE01",                    // Artemis Winterfest - Suburban Deco
				"AHT01",                    // Artemis Winterfest - Tree Deco
				"ASM01",                    // Artemis Winterfest - Snowman Phase 1
				"ASM02",                    // Artemis Winterfest - Snowman Phase 2
				"ASM03",                    // Artemis Winterfest - Snowman Phase 3
				"ASM04",                    // Artemis Winterfest - Snowman Phase 4
				"ASM05",                    // Artemis Winterfest - Snowman Phase 5
				"ASM06",                    // Artemis Winterfest - Snowman Phase 6
				"ASM07",                    // Artemis Winterfest - Snowman Phase 7
				"ASM08",                    // Artemis Winterfest - Snowman Phase 8
				"ASM09",                    // Artemis Winterfest - Snowman Phase 9
				"TopSecret",                // Holiday Battle Bus 2021
				"c_2021_SM02",              // Snow Stage 02 Client
				"c_2021_SM03",              // Snow Stage 03 Client
				"c_2021_SM04",              // Snow Stage 04 Client
				"c_2021_SM05",              // Snow Stage 05 Client
			},

			"19.10": {
				"IODS_3", // IO_Drill 03
			},

			"19.20": {
				"IODS_4", // IO_Drill 04
				"IODS_5", // IO_Drill 05
			},

			"19.30|19.40": {
				"2022_Milkshake_Hole_001", // Event_Milkshake_Hole_001
				"MLQ22",                   // Milkshake Event
				"2022_Milkshake_Hole_002", // Event_Milkshake_Hole_002
				"2022_Milkshake_Hole_003", // Event_Milkshake_Hole_003
				"MLQR22",                  // Milkshake Random Event
			},

			"20.00": {
				"WE_TBYB_BBRG",  //War Effort: Boogie Bomb vs. Rift To Go TBYB
				"WE_TBYB_BSB",   //War Effort: Balloons vs. Bubble Shield TBYB
				"WE_TRT_RR",     //War Effort: Turrets - Rocky Reels
				"WE_BB_RR",      //War Effort: Battle Bus - Rocky Reels
				"WE_CPA",        //War Effort: Choppa Funding
				"WE_BSB",        //War Effort: Balloons vs. Bubble Shield Funding
				"WE_TRT_TLT",    //War Effort: Turrets - Tilted Towers
				"WE_BB_TLT",     //War Effort: Battle Bus - Tilted Towers
				"WE_BBRG",       //War Effort: Boogie Bomb vs. Rift To Go Funding
				"WE_TRT_CCR",    //War Effort: Turrets - Coney Crossroads
				"WE_BB_CCR",     //War Effort: Battle Bus - Coney Crossroads
				"WE_TRT_CDC",    //War Effort: Turrets - Condo Canyon
				"WE_BB_CDC",     //War Effort: Battle Bus - Condo Canyon
				"WE_TRT_SCT",    //War Effort: Turrets - Sanctuary
				"WE_TRT",        //War Effort: Turret Funding
				"WE_BB",         //War Effort: Battle Bus Funding
				"WE_BB_SWS",     //War Effort: Battle Bus - Workshop
				"WE_BB_SCT",     //War Effort: Battle Bus - Sanctuary
				"WE_TRT_BGL",    //War Effort: Turrets - Bugle
				"WE_BB_BGL",     //War Effort: Battle Bus - Bugle
				"WE_RDCAR",      //War Effort: Combat vs. Red Dot Funding
				"WE_TBYB_RDCAR", //War Effort: Combat vs. Red Dot TBYB
				"ResolveBus",    //Resolve Battle Bus 2022
				"LF01",          //Lantern Fest
			},

			"20.10": {
				"BS_2_C", //Battle For Condo CLIENT
				"BS_1_C", //Battle For Bugle CLIENT
				"BS_2",   //Battle For Condo
				"BS_1",   //Battle For Bugle
			},

			"20.20": {
				"BS_4_C", //Battle For Rocky CLIENT
				"BS_3_C", //Battle For Coney CLIENT
				"BS_4",   //Battle For Rocky
				"BS_3",   //Battle For Coney
			},

			"20.30": {
				"BS_6_C", //Post Battles CLIENT
				"BS_5_C", //Battle For Tilted CLIENT
				"BS_6",   //Post Battles
				"BS_5",   //Battle For Tilted
			},

			"20.40": {
				"AL01",   // Armadillo In-Game Leadup
				"AL02",   // Armadillo Lobby Leadup
				"ARLU01", // ArmadilloRoll Lead Up
			},

			"21.00": {
				"CelebrationBus", // Celebration Battle Bus
				"GS_01",          // Grapple_Stops_On
				"RBCLG01",        // Logjam Bloom (Closed)
				"RRBB01",         // Roots_Butter_01
				"RRBB02",         // Roots_Butter_02
				"RRLG01",         // Roots_LogJam_01
				"RRLG02",         // Roots_LogJam_02
			},

			"21.10": {
				"BB01",   //Reality Bloom: Butter Barn
				"RBCT01", //Temple Bloom (Closed)
				"RRT01",  //Roots_Temple_01
				"RRT02",  //Roots_Temple_02
				"RRSS01", //Roots_Sleepy_01
			},

			"21.20": {
				"NSSC01", //No Sweat Summer Concert (Construction)
				"RBSS01", //Sleepy Bloom (Closed)
				"RRSS02", //Roots_Sleepy_02,
				"RRCC01", //Roots_Condo_01
			},

			"21.30": {
				"EventFlag.Event_NoSweatSummer", // Event_S21_NoSweatSummer
				"STMN_PH",                       //Stamina Preheat
				"RBB01",                         // Bugle Bloom (Closed)
				"RRB01",                         //Roots_Bugle_01
				"RRB02",                         //Roots_Bugle_02
				"RBCC02",                        //Condo Bloom (Open)
				"RBSS02",                        //Sleepy Bloom (Open)
				"RRCC02",                        //Roots_Condo_02
				"NSSC02",                        // No Sweat Summer Concert (Constructed)
				"NSRT0",                         //Rebuild Tilted Voting: First Round
				"NSRT1",                         //Rebuild Tilted Voting: Second Round,
				"NSRT2",                         //Rebuild Tilted Voting: Final Round
				"RBTP1S2",                       //Rebuild Tilted Plot 1 Stage 2
				"RBTP2S2",                       //Rebuild Tilted Plot 2 Stage 2
				"RBTP3S2",                       //Rebuild Tilted Plot 3 Stage 2,

			},

			"21.40|21.50|21.51": {
				"Event_S21_Stamina", // Dragon Ball tab(21.40)
			},

			"21.40": {
				"StaminaBus", //Stamina Battle Bus
				"RBTP1F",     //Rebuild Tilted Plot 1 Final
				"RBTP2F",     //Rebuild Tilted Plot 2 Final
				"RBTP3F",     //Rebuild Tilted Plot 3 Final
			},

			"21.50": {
				"RainbowBus", //Rainbow Bus
				"HB01",       //Hyper Bloom (Locked)
			},
			"22.00": {
				"FNBD05",         // Fortnite Birthday Map Changes
				"BirthdayBus5th", // 5th Birthday Battle Bus
			},

			"22.20": {
				"EventFlag.HalloweenBattleBus", // Fortnitemares Battlebus
			},

			"22.10": {
				"S22_Vortex_Active",
			},

			"22.30": {
				"EventFlag.FNCSBus", //FNCS BattleBus
			},

			"22.40": {
				"AB_SF", // Airborne Soccer Fields
				"RL01",  // Fracture In-Game Leadup
				"RL02",  // Fracture Radish Lobby Leadup
				"RL12",  // Fracture Lobby Leadup 2
				"RL13",  // Fracture Lobby Leadup 3
				"RL14",  // Fracture Lobby Leadup 4
				"RL15",  // Fracture Lobby Leadup 5
				"RL16",  // Fracture Lobby Leadup 6
				"RL22",  // Fracture In-Game Leadup 2
				"RL23",  // Fracture In-Game Leadup 3
				"RL24",  // Fracture In-Game Leadup 4
				"RL25",  // Fracture In-Game Leadup 5
				"RL26",  // Fracture In-Game Leadup 6
				"RC01",  // Fracture Core
				"RC02",  // Fracture Core
				"RC03",  // Fracture Core
				"RC04",  // Fracture Core
				"RC05",  // Fracture Core
				"RC06",  // Fracture Core
			},

			"23.10": {
				"CalendarEvent_Season23_Winterfest", // Winterfest 2022
				"EventFlag.LTE_WinterFestTab",       // Winterfest 2022
				"Winter",                            // Winterfest Cosmetics SupplyDrop/BattleBus
				"WinterDeco",                        // Winter Decorations
			},

			"23.20": {
				"RGCONSTUCTION02", // RiftGate_ConstructionPhase_02
			},

			"23.30": {
				"ArcticIce", // Event_ArcticIce
			},

			"23.40": {
				"EventFlag.Event_Vaultbreakers", // Most Wanted tab
				"EventFlag.Event_S23_ProgressiveBackbling",
			},

			"23.50": {
				"EventFlag.Event_Vaultbreakers", // Most Wanted tab
				"EventFlag.Event_S23_FindIt",    // Event_S23_FindIt
				"EventFlag.Event_S23_ProgressiveBackbling",
				"RG01",         // RiftGate Stage 1
				"RG02",         // RiftGate Stage 2
				"ST01",         // Skytear Stage 1
				"ST02",         // Skytear Stage 2
				"FNCSMajorBus", //FNCS Major BattleBus
				"EventFlag.Event_S23Cipher",
			},

			"24.00|24.10": {
				"EventFlag.Event_S24_LanternFest", // Latern Fest 2023
				"Lanternfest",                     // Latern Fest 2023
				"EventFlag.Event_S24_SunBurst",    // Sunburst
				"EventFlag.Event_S24_SpringFling", // Springfling
			},

			"24.30|24.40": {
				"EventFlag.FNCSMajorBus",       // Event_FNCSTakeover_BattleBus
				"EventFlag.Event_PlotTwist",    // Star Wars 2023 tab
				"EventFlag.Event_PlotTwist_E2", // Event_PlotTwist_E2
				"EventFlag.Event_PlotTwist_E3", // Event_PlotTwist_E3
				"MLQ24",                        // Event_S24_Milkshake
			},

			"25.11": {
				"14DAYS",      // 14 Days of Summer: Propping
				"14DAYSMURAL", // 14 Days of Summer: Mural
			},

			"25.30": {
				"EventFlag.Event_BelongTreaty", // Jujutsu Kaisen tab
			},

			"26.10": {
				"FNBD2023", // FN Birthday 2023(6th Birthday)
				"EventFlag.BirthdayBus6th",
			},

			"26.30": {
				"FNMBus", // FortniteMares Battle Bus
			},

			"27.11|31.41|32.11": {
				"DL01", // Durian & Kiln Event-related (Countdown)
				"DL02", // Durian & Kiln Event-related (Countdown)
			},

			"28.01": {
				"Winter", // Winterfest 2023
			},

			"28.10": {
				"EventFlag.Event_LinedNotebook_Teaser", // TMNT Tab countdown
			},

			"28.20|28.30": {
				"EventFlag.Event_LinedNotebook", // TMNT mini pass
				/*
				   "CH5S1CPPE", //Pre-Emergence Event(Central Picnic)
				   "CH5S1CPCE", //Crater event - Pre-Emergence Event (Central Picnic)
				   "CH5S1CPTE", //Thump Event - Pre-Emergence Event (Central Picnic)
				   "CH5S1CPFH", //Final Hour - Pre-Emergence Event (Central Picnic)
				   "CH5S1CPFP", //Titan Hand Appears(Central Picnic)
				   "CH5S1CPPP", //Primary Events(Central Picnic)
				   "CH5S1CPEP", //Post-Chain Event(Central Picnic)

				    Flags labled "Pre-Emergence Event" need to have the "CH5S1CPPE" flag to work event flags not labled with the "Pre-Emergence Event" need to have "CH5S1CPPE" disabled to work
				*/
			},

			"29.00|29.01|29.10": {
				"EventFlag.Event_S29_SeasonalActivation",
				"EventFlag.Event_S29_MidasActivation",
			},

			"29.20|29.30": {
				"EventFlag.Event_S29_ColdDay", // Avatar mini pass + cabbages
				"AtlaShrines",                 // Avatar map changes
				"AtlaScrolls",                 // Avatar map changes
				"AtlaChests",                  // Avatar map changes
			},

			"29.40": {
				/*
				   "EventFlag.Event_Osiris",
				   "SUPERSPORT_BUILDUP_1",
				   "SUPERSPORT_BUILDUP_2",
				   "SUPERSPORT_BUILDUP_3",
				   "SUPERSPORT_CHARGE_1",
				   "SUPERSPORT_STRIKE_1",
				   "SUPERSPORT_CHARGE_2",
				   "SUPERSPORT_STRIKE_2"
				   "SUPERSPORT_CHARGE_3",
				   "SUPERSPORT_STRIKE_3"
				   "SUPERSPORT_SANDSTORM",
				   "SUPERSPORT_LIGHT_2",
				*/
			},

			"30.20|30.30": {
				"EventFlag.Event_Greenhouse", //Pirates Of the Caribbean Mini Pass
				"GreenhouseActive",
				"RE_01_Cutout",
				"RE_02_Cutout",
				"RE_03_Cutout",
				"RE_01", //Nitro Pipeline Stage 1
				"RE_02", //Nitro Pipeline Stage 2
				"RE_03", //Nitro Pipeline Stage 3
			},

			"31.20": {
				"EventFlag.Event_S31_Birthday",
			},

			"31.40|31.41": {
				"EventFlag.Event_S31_NearAcorn",
			},

			"31.20|31.30": {
				"EventFlag.Event_ButtercakePass",
			},

			"31.41": {
				"KilnB",              //Kiln Buildup
				"Kiln_Preload_Event", //Kiln Preload Event
			},

			"33.11": {
				"EventFlag.S33_WinterFestTab", //2024 winterfest
				"WF_IG_Avail",                 //2024 winterfest
			},

			"33.30": {
				"GrillMoon_Day1", //GrillMoon Mini Event Buildup Stages
				"GrillMoon_Day2", //GrillMoon Mini Event Buildup Stages
				"GrillMoon_Day3", //GrillMoon Mini Event Buildup Stages
				"GrillMoon_Day4", //GrillMoon Mini Event Buildup Stages
				"GrillMoon_Day5", //GrillMoon Mini Event Buildup Stages
			},

			"34.30|34.40": {
				/*"GiftedTraveler_Phase1",          //GiftedTraveler Buildup Stages(Daigo Portal Beams)
				"GiftedTraveler_Phase2",          //GiftedTraveler Buildup Stages(Daigo Portal Beams)
				"GiftedTraveler_Phase3",          //GiftedTraveler Buildup Stages(Daigo Portal Beams)
				"GiftedTraveler_Phase4",          //GiftedTraveler Buildup Stages(Daigo Portal Beams)
				"GiftedTraveler_FinalPhase",*/ //GiftedTraveler Buildup Stages(Daigo Portal Beams)
			},

			"35.20": {
				"RHDEvent", //Ripehoneydew Event Countdown
			},

			"36.30": {
				/*
				   "PC_01", //Pistol Canary Stage 1
				   "PC_02", //Pistol Canary Stage 2
				   "PC_03", //Pistol Canary Stage 3
				   "Event_S33_Winterfest_BrickSoup_Sequence", //Pistol Canary Sequence
				   "Event_S33_Winterfest_BrickSoup_PreLoad_Sequence", //Pistol Canary Pre Load Sequence
				   "PC_Event", //Pistol Canary Event
				   "EventStart", //Pistol Canary Event Start
				   "PC_After", //Pistol Canary After Event
				   "PC_TentacleStage2", //Pistol Canary Tentacle Stage 2
				*/
				"PC_Timer", //Pistol Canary Countdown
			},

			"37.31": {
				"SM_Monolith_2",
				"SM_Logo2",
				"SM_PostEvent",
				"SM_VoiceLog_5",
			},

			"38.11": {
				"SM_Timer",
			},

			"39.10|39.11": {
				"DC_Orb",
			},

			"39.11": {
				"EventFlag.S39_WinterFestTab", //Winterfest Tab 2025
				"WF_IG_Avail",
			},

			"39.20": {
				"DC_RiftRotate3",
			},

			"39.50|39.51": {
				"DC_3ShardsVisible",
			},

			"40.00|40.10|40.20|40.30|40.40": {
				"EventFlag.Event_Rivals_Act_04",  //Showdown Act4
				"EventFlag.Event_Rivals_Gate_01", //Showdown Act1 Mid
				"EventFlag.Event_Rivals_Gate_02", //Showdown Act1 End
				"EventFlag.Event_Rivals_Gate_03", //Showdown Act2 Mid
				"EventFlag.Event_Rivals_Gate_04", //Showdown Act2 End
				"EventFlag.Event_Rivals_Gate_05", //Showdown Act3 End
				"EventFlag.Event_Rivals_MS_00",   //Showdown Milestone 1
				"EventFlag.Event_Rivals_MS_01",   //Showdown Milestone 2
				"EventFlag.Event_Rivals_MS_02",   //Showdown Milestone 3
				"EventFlag.Event_Rivals_MS_03",   //Showdown Milestone 4
				"EventFlag.Event_Rivals_MS_04",   //Showdown Milestone 5
				"EventFlag.Event_Rivals_MS_05",   //Showdown Milestone 6
				"EventFlag.Event_Rivals_MS_06",   //Showdown Milestone 7
				"EventFlag.Event_Rivals_MS_07",   //Showdown Milestone 8
				"EventFlag.Event_Rivals_MS_08",   //Showdown Milestone 9
				"EventFlag.Event_Rivals_MS_09",   //Showdown Milestone 10
				"EventFlag.Event_Rivals_MS_10",   //Showdown Milestone 11
				"EventFlag.Event_Rivals_MS_11",   //Showdown Milestone 12
			},

			"40.10": {
				"EventFlag.Event_Foolnite_Activation", // April Fools 2026
				"FF_BHMActivate",                      // April Fools 2026
				"DC_Rift_BlindBird",
				"Figment_BlindBird",
			},

			"40.40": {
				"EventFlag.Event_MorningBellQuests_S40_Pt1",
				"EventFlag.Event_MorningBellQuests_S40_Pt2",
				"EventFlag.Event_FlameCeleb_S40",
				"DC_4ShardsVisible",
			},

			"40.41": {
				"Squid_to_ST",
				"DC_5ShardsVisible",
			},

			"41.30": {
				"PonyBackerBuildUp", //Unstable Event Countdown
			},

			"42.10|42.20": { //assuming first stage is active for 2 versions
				"DustyUpdate_A",
				//"DustyUpdate_B"
				//"DustyUpdate_C"
				//"DustyUpdate_D" //might make it a config but this is fine for now
			},
		}

		if versionInfo.VersionGlobal == 12 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.LTE_SpyGames",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 15 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S15_Legendary_Week_01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S15_Legendary_Week_02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S15_Legendary_Week_03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S15_Legendary_Week_04",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S15_Legendary_Week_05",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S15_Legendary_Week_06",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S15_Legendary_Week_07",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S15_Legendary_Week_08",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S15_Legendary_Week_09",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S15_Legendary_Week_10",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S15_Legendary_Week_11",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S15_Legendary_Week_12",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S15_Legendary_Week_13",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S15_Legendary_Week_14",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S15_Legendary_Week_15",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_Spire_Challenges",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Foreshadowing_Quest_01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Foreshadowing_Quest_02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Foreshadowing_Quest_03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Foreshadowing_Quest_05",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_NBA_Challenges",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 16 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S16_Legendary_Week_01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S16_Legendary_Week_02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S16_Legendary_Week_03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S16_Legendary_Week_04",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S16_Legendary_Week_05",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S16_Legendary_Week_06",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S16_Legendary_Week_07",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S16_Legendary_Week_08",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S16_Legendary_Week_09",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S16_Legendary_Week_10",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S16_Legendary_Week_11",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S16_Legendary_Week_12",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_HiddenRole",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_OperationSnowdown",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_PlumRetro",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 17 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S17_Legendary_Week_01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S17_Legendary_Week_02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S17_Legendary_Week_03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S17_Legendary_Week_04",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S17_Legendary_Week_05",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S17_Legendary_Week_06",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S17_Legendary_Week_07",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S17_Legendary_Week_08",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S17_Legendary_Week_09",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S17_Legendary_Week_10",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S17_Legendary_Week_11",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S17_Legendary_Week_12",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S17_Legendary_Week_13",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S17_Legendary_Week_14",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_CosmicSummer",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_O2_Challenges",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S17_CB_Radio",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTE_Cashier_Quests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTE_Buffet_PartyTrooper",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTE_Buffet_Cosmetics",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTE_Buffet_PreQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTE_Buffet_Attend",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTE_Buffet_PostQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S17_Yeet_Week",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S17_Sneak_Week",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 18 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S18_Repeatable_Weekly",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S18_Repeatable_Weekly",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTE_Season18_TextileQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTE_Season18_BirthdayQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_Fornitemares_2021",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S18_Repeatable_Weekly_06",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_HordeRush",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S18_Repeatable_Weekly_07",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S18_Repeatable_Weekly_08",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTQ_S18_Repeatable_Weekly_09",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.S18_WildWeek_Bargain",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.S18_WildWeek_Shadows",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
			if managers.Config.TimelineConfigs.CH2S8.CubeTownStage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "CT01",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
			if managers.Config.TimelineConfigs.CH2S8.CubeTownStage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "CT02",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
			if managers.Config.TimelineConfigs.CH2S8.CubeTownStage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "CT03",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
		}

		if versionInfo.VersionGlobal == 19 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_S19_DeviceQuestsPart1",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S19_DeviceQuestsPart2",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S19_DeviceQuestsPart3",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S19_Gow_Quests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.S19_WinterfestCrewGrant",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SoundWave",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTM_CE_Q",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTM_SG_Q",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTM_UV_Q",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 20 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.NoBuildQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "LF01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S20_EmicidaQuest",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S20_LevelUpPack",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S20_KlomboQuest_0",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S20_KlomboQuest_1",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S20_KlomboQuest_2",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S20_May4thQuest",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S20_KlomboQuest",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S20_Armadillo",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S20_AliQuest",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 21 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_NoSweatSummer",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S21_IslandHopper",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.S21_WildWeekend_BargainBin",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S21_CRRocketQuest",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S21_LevelUpPack",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S21_WildWeeks_Fire",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S21_WildWeeks_Kondor",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S21_FallFest",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S21_GenQuest",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S21_WildWeeks_BargainBin",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 22 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_S22NarrativePart1",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S22_Headset",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S22_BirthdayQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.S22FortnitemaresQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S22_LevelUpPack",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.S22ChromePunkQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S22_DistantEcho",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.S22HordeRush",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S22_FNCSQuest",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.S22_Radish_RewardGrantEvent",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S22_VistaQuest",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S22_RedOasis",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.S22_WildWeek_Bargain",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S22_NarrativePart2",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.S22_WildWeek_Avian",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 23 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_ZeroWeek",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_SunBurst",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23NarrativePart1",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Weekly_01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Weekly_02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Weekly_03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Weekly_04",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Weekly_05",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Weekly_06",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Weekly_07",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Weekly_08",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTE_WinterFest",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_RebootRally",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Emerald",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Lettuce",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23NarrativePart2",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "ArcticIce",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23NarrativePart3",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_LevelUpPack",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Supercharged_A",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Supercharged_B",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Supercharged_C",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Supercharged_D",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Supercharged_E",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Supercharged_F",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Supercharged_G",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Supercharged_H",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_FindIt",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23NarrativePart5",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_Overtime",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23Cipher",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23NarrativePart4BonusGoal",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23NarrativePart4",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S23_CreedQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 25 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_BelongTreaty",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S25_LevelUpPack",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S25_RebootRally",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S25_14DOS",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S25_AloeCrouton",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S25_Maze",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 26 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_S26_BirthdayQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S26_FNM",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S26_Mash",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S26_Intertwine",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S26_LevelUpPack",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 27 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.S27_RufusQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.S27RaiderPinkSherbetQuestRemoval",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S27_RebootRally",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 28 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMarPack",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMarSpeedRuns01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_LinedNotebook",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S28DelMarQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S28_Winterfest",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S28_RebootRally",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S28_TryRocketRacing",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S28_LevelUpPass",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S28_Prelude",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS01_SparksPack",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS02_Week01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS02_Week02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS02_Week03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS02_Week04",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS02_Week05",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS02_Week06",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS02_Week07",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS02_Week08",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS02_Week09",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 29 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_S29_Gallery",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_Season01_Kickoff",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_Season01_PartyPlay",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_Season01_Rank",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_Osiris",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S29_RebootRally",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S29_ColdDayPrelude",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S29_LevelUpPass",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S29_SeasonalActivation",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S29_WhiplashWW",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS03_Week01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS03_Week02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS03_Week03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS03_Week04",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS03_Week05",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS03_Week06",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS03_Week07",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 30 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_S30_Beanstalk_Quests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTM_Hyena",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTM_Vigilante",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LTM_ZebraWallet",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_Adidas",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_Inferno",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_QuestPack_02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_Stainless",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_Rank_02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksSpotlight_S04_A",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S30_RebootRally",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S30_BlastBerryQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_GreenhousePrelude",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S30_LevelUpPass",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S30_AllSweat",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S30_FlatWare",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S30_StoryQuests_P6",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S30_StoryQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S30_Stainless",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_Sparks_Stainless01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_Sparks_Stainless02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS04",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS04_Week01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS04_Week02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS04_Week03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS04_Week05",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS04_Week08",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Beanstalk_01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 31 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_S31_Sweatember",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_Adidas",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_Inferno",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S31_NearAcorn",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_QuestPack_02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_Season03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_Stainless",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_TimeTrials_01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_TimeTrials_02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_TimeTrials_03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_Rank_02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_Season03_Rank",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SpooktacularPass",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S31_Birthday",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S31_QuestDeckBed",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S31_FoundQuests_GreenTown",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S31_FoundQuests_PB_P02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "Event_S31_FoundQuests_PB_P01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S31_FoundQuests_TroutWrist",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_RedFish_Quest",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S31_Mash",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S31_RebootRally",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S31_LevelUpPass",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S31_NearAcorn",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S31_RankedScythe",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S31_ToadJam",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS05",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS05_Week01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS05_Week02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS05_Week03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS05_Week04",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS05_Week05",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS05_Week06",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS05_Week07",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS05_Week08",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS05_Week09",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS05_Week10",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS05_Week11",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 32 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_DelMar_TimeTrials_04",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S32_RebootRally",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S32_ScytheGold",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS06",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS06_Week01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS06_Week02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS06_Week03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS06_Week04",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS06_Week05",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS06_Week06",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS06_Week07",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS06_Week08",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS06_Week09",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS06_Week10",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_SparksS06_Week11",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 33 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_S33_BlondeJaw",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S33_Scenario_BatterBoi",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_Scenario_Shell",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S33_NarrativeBattleQuest",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S33_ChatQuest",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.S33_WinterFest",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 34 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_S34_Bling",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S34_CollectionQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S34_Lanternfest",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S34_LTE_Quests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S34_SummerRoadtrip_Quests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S34_RoomCharge",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S34_SprocketPoppy",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S34_StoryQuests_EOS_01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S34_StoryQuests_EOS_02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S34_StoryQuests_P1",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S34_WeaponExpertiseSniper",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S34_WeaponExpertise",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 36 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_S36_MidSeasonCharacter_Quests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S36_IslandStories",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S36_14Days",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_PistolCanary_Quests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_PistolCanary_Quests_After",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_PistolCanary_Quests_After_2",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S36_WeaponExpertise",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S36_SummerRoadtrip_Quests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 37 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_S37_Birthday",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S37_ArrowTube",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S37_CommunityQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S37_Delulu",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S37_PulseLancer",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S37_WeaponExpertise",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S37_Mash",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S37_StoryQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S37_StoryLimerockEvent",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 38 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_S38_Cinematic",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S38_Delulu",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S38_Stories_P01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S38_Stories_P02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S38_Stories_P03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S38_Stories_P04",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S38_Stories",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S37_StoryLimerockEvent",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 39 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_S39_BattleBus",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S39_HoneyFiber",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S39_Story_P01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S39_Story_P02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S39_Story_P03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S39_Story_P04",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S39_UmbraMolt",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S39_LnL_DailyLogin_P01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S39_LnL_DailyLogin_P02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S39_WinterFest",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}
		if versionInfo.VersionGlobal == 40 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_MorningBellQuests_S40_Pt1",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_MorningBellQuests_S40_Pt2",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S40_BR_CommunityDay_Quests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.LobbySeason40",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S40_BR_Elites_Quests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S40_LlamaRama",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S40_SpearShawl",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S40_STWPromo",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 41 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.PearlTaffy_Quests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S41_SpriteMastery_P01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S41_SpriteMastery_P02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S41_SpriteMastery_P03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S41_SpriteMastery_P04",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S41_Story_P01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S41_Story_P02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S41_Story_P03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.41_30_Story_P01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.41_30_Story_P02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.41_30_Story_P03",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.41_30_Story_P04",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.VersionGlobal == 42 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "EventFlag.Event_S42_Audiologs_4200",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S42_Birthday",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S42_BlockStackCosmetics",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S42_SpriteMastery_P01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S42_SpriteMastery_P02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S42_CosmicThunder_TutorialQuests",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S42_NarrowFlea",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.PreviousSeason_SpriteDust",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.ProgressionTrack_S42",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S42_Story_BlockStack_BonusGoals",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S42_Story_BlockStack_P01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S42_Story_BlockStack_P02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S42_Story_HeroicTale_BonusGoals",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S42_Story_HeroicTale",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S42_Story_PhoenixKick_BonusGoals",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S42_Story_PhoenixKick",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S42_Story_SheerWill_BonusGoals",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S42_Story_SheerWill_P01",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.Event_S42_Story_SheerWill_P02",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
					types.FEventRecord{
						EventType:   "EventFlag.ProgressionTrack",
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					},
				)
		}

		if versionInfo.Version == 14.30 || versionInfo.Version == 14.40 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "SPJ02",
						ActiveUntil: time.Now().UTC().Add(1 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
						ActiveSince: time.Now().UTC().Add(-1 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
					},
				)
		}

		if versionInfo.Version == 14.50 || versionInfo.Version == 14.60 {
			calendar.Channels["client-events"].States[0].ActiveEvents =
				append(calendar.Channels["client-events"].States[0].ActiveEvents,
					types.FEventRecord{
						EventType:   "SPJ03",
						ActiveUntil: time.Now().UTC().Add(1 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
						ActiveSince: time.Now().UTC().Add(-1 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
					},
				)
		}

		versionToString := fmt.Sprintf("%.2f", versionInfo.Version)
		for key, flags := range versionEventFlags {
			versions := strings.Split(key, "|")
			if slices.Contains(versions, versionToString) {
				for _, flag := range flags {
					clientEvents.States[0].ActiveEvents = append(clientEvents.States[0].ActiveEvents, types.FEventRecord{
						EventType:   flag,
						ActiveUntil: "9999-12-01T21:10:00.000Z",
						ActiveSince: "2020-11-21T07:00:00.000Z",
					})
				}
			}
		}
		if versionInfo.VersionGlobal == 17 {
			if managers.Config.TimelineConfigs.CH2S7.MoveMotherShipToSlurpySwamp == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "SMVE",
							ActiveUntil: time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval+5) * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: time.Now().UTC().Add(-120 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
						},
					)
			}
			if managers.Config.TimelineConfigs.CH2S7.MotherShipSlurpySwampDoorsOpen == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "SDOPN",
							ActiveUntil: time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval+5) * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: time.Now().UTC().Add(-120 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
						},
					)
			}
			if managers.Config.TimelineConfigs.CH2S7.MotherShipSlurpySwampAbductionWarmup == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "SDOPN",
							ActiveUntil: time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval+5) * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: time.Now().UTC().Add(-120 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
						},
						types.FEventRecord{
							EventType:   "SWRMDUP",
							ActiveUntil: time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval+5) * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: time.Now().UTC().Add(-120 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
						},
					)
			}

			if managers.Config.TimelineConfigs.CH2S7.MotherShipSlurpySwampAbduction == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "ABDSLP",
							ActiveUntil: time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval+5) * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: time.Now().UTC().Add(-120 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
						},
					)
			}

			if managers.Config.TimelineConfigs.CH2S7.MoveMotherShipToCoralCastle == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "CMVE",
							ActiveUntil: time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval+5) * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: time.Now().UTC().Add(-120 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
						},
					)
			}
			if managers.Config.TimelineConfigs.CH2S7.MotherShipCoralCastleDoorsOpen == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "CDOPN",
							ActiveUntil: time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval+5) * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: time.Now().UTC().Add(-120 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
						},
					)
			}
			if managers.Config.TimelineConfigs.CH2S7.MotherShipCoralCastleAbductionWarmup == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "CDOPN",
							ActiveUntil: time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval+5) * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: time.Now().UTC().Add(-120 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
						},
						types.FEventRecord{
							EventType:   "CWRMDUP",
							ActiveUntil: time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval+5) * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: time.Now().UTC().Add(-120 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
						},
					)
			}

			if managers.Config.TimelineConfigs.CH2S7.MotherShipCoralCastleAbduction == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "ABDCRL",
							ActiveUntil: time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval+5) * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: time.Now().UTC().Add(-120 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
						},
					)
			}

			if managers.Config.TimelineConfigs.CH2S7.MoveMotherShipToCornyComplex == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "FMVE",
							ActiveUntil: time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval+5) * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: time.Now().UTC().Add(-120 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
						},
					)
			}
			if managers.Config.TimelineConfigs.CH2S7.MotherShipCornyComplexDoorsOpen == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "FDOPN",
							ActiveUntil: time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval+5) * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: time.Now().UTC().Add(-120 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
						},
					)
			}
			if managers.Config.TimelineConfigs.CH2S7.MotherShipCornyComplexAbductionWarmup == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "FDOPN",
							ActiveUntil: time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval+5) * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: time.Now().UTC().Add(-120 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
						},
						types.FEventRecord{
							EventType:   "FWRMDUP",
							ActiveUntil: time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval+5) * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: time.Now().UTC().Add(-120 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
						},
					)
			}

			if managers.Config.TimelineConfigs.CH2S7.MotherShipCornyComplexAbduction == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "ABDFRM",
							ActiveUntil: time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval+5) * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: time.Now().UTC().Add(-120 * time.Hour).Format("2006-01-02T15:04:05.000Z"),
						},
					)
			}
		}
		if versionInfo.VersionGlobal == 27 {
			if managers.Config.TimelineConfigs.S27MapStages.Stage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "RufusWeek2",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
			if managers.Config.TimelineConfigs.S27MapStages.Stage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "RufusWeek3",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
			if managers.Config.TimelineConfigs.S27MapStages.Stage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "RufusWeek4",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
		}
		if versionInfo.VersionGlobal == 32 {
			if managers.Config.TimelineConfigs.RemixMapStages.Stage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "ClydeSeason1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "Week1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.RemixMapStages.Stage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "ClydeSeason2Part1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "Week2",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.RemixMapStages.Stage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "ClydeSeason2Part2",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "Week3",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.RemixMapStages.Stage4 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "KL1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "Week4",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "ClydeSeason3Part1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.RemixMapStages.JuiceWRLDPOIStage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "KL1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.RemixMapStages.JuiceWRLDPOIStage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "KL2",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.RemixMapStages.JuiceWRLDPOIStage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "KL3",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
		}

		if versionInfo.VersionGlobal == 10 || versionInfo.VersionGlobal == 42 {
			if managers.Config.TimelineConfigs.CH1S10.BEnableRuinHouseBeacon == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "RBPRH",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "RBBRH",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH1S10.BEnableGothemCityBeacon == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "RBBMP",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "RBBM",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH1S10.BEnableGreasyGroveBeacon == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "RBPGG",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "RBBGG",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH1S10.BEnableMoistyPalmsBeacon == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "RBPMP",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "RBBMP",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH1S10.BEnableFrenzyFieldsBeacon == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "RBPFI",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "RBBFI",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH1S10.BEnableOakBeacon == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "RBPBO",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "RBBO",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH1S10.BEnableVoidBeacon == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "RBPVA",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "RBBVA",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH1S10.BEnableRetailRowBeacon == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "RBPRR",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "RBBRR",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH1S10.BEnableFloatingCubeIsland == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "RBFI",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
		}

		if versionInfo.VersionGlobal == 28 {
			if managers.Config.TimelineConfigs.TMNTTabStages.Stage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "EventFlag.Event_LinedNotebook_Teaser1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.TMNTTabStages.Stage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "EventFlag.Event_LinedNotebook_Teaser2",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.TMNTTabStages.Stage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "EventFlag.Event_LinedNotebook_Teaser3",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
		}

		if versionInfo.VersionGlobal == 29 {
			if managers.Config.TimelineConfigs.AvatarIcebergStages.Stage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "AtlaIceberg1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.AvatarIcebergStages.Stage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "AtlaIceberg2",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.AvatarIcebergStages.Stage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "AtlaIceberg3",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
		}

		if versionInfo.Version == 9.3 {
			if managers.Config.TimelineConfigs.CH1S9MechStages.Stage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "Tob0", // Mech Stage 1
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "Tob01", // Robot Construction Site
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH1S9MechStages.Stage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "Tob1", // Mech Stage 2
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "Tob01", // Robot Construction Site
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH1S9MechStages.Stage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "Tob2", // Mech Stage 3
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "Tob01", // Robot Construction Site
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH1S9MechStages.Stage4 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "Tob3", // Mech Stage 4
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "Tob01", // Robot Construction Site
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH1S9MechStages.Stage5 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "Tob4", // Mech Stage 5
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "Tob01", // Robot Construction Site
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH1S9MechStages.Stage6 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "Tob5", // Mech Stage 6
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "Tob01", // Robot Construction Site
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH1S9MechStages.Stage7 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "Tob6", // Mech Stage 7
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "Tob01", // Robot Construction Site
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH1S9MechStages.Stage8 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "Tob7", // Mech Stage 8
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "Tob01", // Robot Construction Site
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH1S9MechStages.Stage9 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "Tob8", // Mech Stage 9
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "Tob01", // Robot Construction Site
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
		}

		if versionInfo.VersionGlobal == 13 {

			if managers.Config.TimelineConfigs.S13WaterLevel.WaterLevel0 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "WL0", // Water Level 0
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.S13WaterLevel.WaterLevel1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "WL1", // Water Level 1
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.S13WaterLevel.WaterLevel2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "WL2", // Water Level 2
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.S13WaterLevel.WaterLevel3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "WL3", // Water Level 3
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.S13WaterLevel.WaterLevel4 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "WL4", // Water Level 4
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.S13WaterLevel.WaterLevel5 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "WL5", // Water Level 5
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.S13WaterLevel.WaterLevel6 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "WL6", // Water Level 6
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.S13WaterLevel.WaterLevel7 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "WL7", // Water Level 7
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
		}

		if versionInfo.Version == 12.40 || versionInfo.Version == 12.41 {
			if managers.Config.TimelineConfigs.JerkyConcertStages.Stage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "JS01", // Jerky Stage Delivery Trucks
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.JerkyConcertStages.Stage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "JS02", // Jerky Stage Partially Built
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.JerkyConcertStages.Stage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "JS03", //Jerky Stage Finished Stage
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

		}

		if versionInfo.VersionGlobal == 31 {
			if managers.Config.TimelineConfigs.CH5S4.RemixPreludeEventBuildUpStage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "StageA",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH5S4.RemixPreludeEventBuildUpStage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "StageB",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH5S4.RemixPreludeEventBuildUpStage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "StageC",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH5S4.RemixPreludeEventBuildUpStage4 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "StageD",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
		}

		if versionInfo.VersionGlobal == 35 {
			if managers.Config.TimelineConfigs.CH6MS1.DeathStarStage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "DS_0",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH6MS1.DeathStarStage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "DS_1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH6MS1.DeathStarStage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "DS_2",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH6MS1.DeathStarStage4 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "DS_3",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
		}

		if versionInfo.VersionGlobal == 34 {
			if managers.Config.TimelineConfigs.CH6S2.DeathStarStage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "GiftedTraveler_Observer_Phase1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH6S2.DeathStarStage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "GiftedTraveler_Observer_Phase2",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH6S2.DeathStarStage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "GiftedTraveler_Observer_Phase3",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH6S2.DeathStarStage4 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "GiftedTraveler_Observer_Phase4",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH6S2.DeathStarStage5 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "GiftedTraveler_Observer_Phase5",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
		}

		if versionInfo.VersionGlobal == 37 {
			if managers.Config.TimelineConfigs.CH6S4.AlienOverlordsEventBuildupStage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "LimeRock_Tree1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH6S4.AlienOverlordsEventBuildupStage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "LimeRock_Tree2",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH6S4.AlienOverlordsEventBuildupStage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "LimeRock_Tree3",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH6S4.AlienOverlordsEventBuildupTreeFX1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "LimeRock_Tree3",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "LimeRock_TreeFX_1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH6S4.AlienOverlordsEventBuildupTreeFX2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "LimeRock_Tree3",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "LimeRock_TreeFX_2",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

		}

		if versionInfo.VersionGlobal == 38 {
			if managers.Config.TimelineConfigs.CH6MS2.HomerSimpsonCloudStage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "BC_Cloud1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH6MS2.HomerSimpsonCloudStage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "BC_Cloud2",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH6MS2.HomerSimpsonCloudStage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "BC_Cloud3",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH6MS2.HomerSimpsonOcean == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "FT_Head",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
		}

		if versionInfo.VersionGlobal == 39 {
			if managers.Config.TimelineConfigs.CH7S1.BattleBusConstructionStage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "EventFlag.Event_DragoncartNarrative_Stage01",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S1.BattleBusConstructionStage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "EventFlag.Event_DragoncartNarrative_Stage02",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S1.BattleBusConstructionStage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "EventFlag.Event_DragoncartNarrative_Stage03",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S1.BattleBusConstructionStage4 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "EventFlag.Event_DragoncartNarrative_Stage04",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S1.DarkHarvesterOcean == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "FlyingOrb_Ocean",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S1.DarkHarvesterStarterIsland == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "FlyingOrb_StarterIsland",
							ActiveUntil: time.Now().UTC().Add(60 * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S1.Shard1DeliveryStage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "DC_ShardDelivery1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
			if managers.Config.TimelineConfigs.CH7S1.Shard1DeliveryStage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "DC_ShardDelivery2",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
			if managers.Config.TimelineConfigs.CH7S1.Shard1Delivered == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "DC_ShardDelivery3",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S1.RedRiftStage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "RL_Rift1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S1.RedRiftStage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "RL_Rift2",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
		}

		if versionInfo.VersionGlobal == 40 {
			if managers.Config.TimelineConfigs.CH7S2.OGDarkHarvesterStage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "BlindBird_P1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S2.OGDarkHarvesterStage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "BlindBird_P2_Pre",
							ActiveUntil: time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval) * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S2.OGDarkHarvesterStage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "BlindBird_P3",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S2.OGDarkHarvesterStage4 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "BlindBird_P4_Pre",
							ActiveUntil: time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval) * time.Second).Format("2006-01-02T15:04:05.000Z"),
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S2.EliteStrongholdEnableRift == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "RF_RiftIn2",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S2.EliteStrongholdDarkHarvesterStage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "RF_OrbIn2",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S2.Shard4DeliveryStage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "DC_ShardDelivery4",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S2.Shard4Delivered == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "DC_ShardDelivered4",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "DC_4ShardsVisible",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S2.Shard5DeliveryStage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "DC_Oasis_Delivery1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S2.Shard5DeliveryStage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "DC_Oasis_Delivery2",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S2.Shard5DeliveryStage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "DC_Oasis_Delivery3",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S2.Shard5Delivered == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "DC_5ShardsVisible",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "DC_Oasis_Delivered",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S2.Shard6DeliveryStage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "DC_Squid_Delivery1",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "DC_5ShardsVisible",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S2.Shard6DeliveryStage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "DC_Squid_Delivery2",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "DC_5ShardsVisible",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S2.Shard6DeliveryStage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "DC_5ShardsVisible",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "DC_Squid_Delivery3",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S2.Shard6Delivered == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "DC_Squid_Delivered",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
						types.FEventRecord{
							EventType:   "DC_6ShardsVisible",
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
		}

		if versionInfo.Version == 41.30 {
			if managers.Config.TimelineConfigs.CH7S3OGMechStages.Stage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "Tob1", // Mech Stage 1
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S3OGMechStages.Stage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "Tob2", // Mech Stage 2
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S3OGMechStages.Stage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "Tob3", // Mech Stage 3
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S3OGMechStages.Stage4 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "Tob4", // Mech Stage 4
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S3OGMechStages.Stage5 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "Tob5", // Mech Stage 5
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S3OGMechStages.Stage6 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "Tob6", // Mech Stage 6
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.CH7S3OGMechStages.Stage7 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "Tob7", // Mech Stage 7
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
		}

		if versionInfo.VersionGlobal == 11 && versionInfo.Version < 11.30 {
			if managers.Config.TimelineConfigs.RiskyReelsStarWarsBuildUpStages.Stage1 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "GalRisky_001", // Star Wars Event Risky Buildup Stage 1
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.RiskyReelsStarWarsBuildUpStages.Stage2 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "GalRisky_002", // Star Wars Event Risky Buildup Stage 2
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}

			if managers.Config.TimelineConfigs.RiskyReelsStarWarsBuildUpStages.Stage3 == true {
				calendar.Channels["client-events"].States[0].ActiveEvents =
					append(calendar.Channels["client-events"].States[0].ActiveEvents,
						types.FEventRecord{
							EventType:   "GalRisky_003", // Star Wars Event Risky Buildup Stage 3
							ActiveUntil: "9999-12-01T21:10:00.000Z",
							ActiveSince: "2020-11-21T07:00:00.000Z",
						},
					)
			}
		}

		calendar.Channels["client-events"].States[0].State["seasonNumber"] = versionInfo.VersionGlobal
		calendar.Channels["client-events"].States[0].State["seasonTemplateId"] = fmt.Sprintf("AthenaSeason:athenaseason%d", versionInfo.VersionGlobal)
		calendar.Channels["client-events"].States[0].State["dailyStoreEnd"] = time.Now().UTC().Add(managers.Config.GeneralSettings.ItemShopRefreshInterval * time.Second).Format("2006-01-02T15:04:05.000Z")
		calendar.Channels["client-events"].States[0].State["weeklyStoreEnd"] = time.Now().UTC().Add(managers.Config.GeneralSettings.ItemShopRefreshInterval * time.Second).Format("2006-01-02T15:04:05.000Z")
		clientEvents.CacheExpire = time.Now().UTC().Add(time.Duration(managers.Config.TimelineConfigs.TimelineRefreshInterval) * time.Second).Format("2006-01-02T15:04:05.000Z")

		calendar.Channels["client-events"] = clientEvents
		calendar.CurrentTime = time.Now().UTC().Format("2006-01-02T15:04:05.000Z")

		tk := calendar.Channels["tk"]
		calendar.Channels["tk"].States[0].State["k"] = managers.Keychain
		tk.CacheExpire = clientEvents.CacheExpire
		calendar.Channels["tk"] = tk

		calendarJson, _ := json.Marshal(calendar)
		c.Data(http.StatusOK, "application/json", calendarJson)

	})
}

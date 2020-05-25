package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (

	// distasters
	earthquakeswitch, tornadoswitch, meteorswitch bool
	// magic
	teleportswitch, rainoffireswitch, shockblockswitch, distortionswitch bool
	// miniboss
	minibossDie   bool
	minibosskills int
	// audio
	startIntroMusic, checkMusic, musicOn, levelMusic, levelEndMusic, enemyDieSound, hpLossSound, dropCoinSound, coinCollectedSound, shootSound, fxOn bool
	chooseMusic                                                                                                                                      int
	levelTune, introTune, levelEndTune                                                                                                               rl.Music

	// level end
	menuTextSpaceColor2 = rl.Green
	enemyendv2          = rl.NewVector2(700, 360)
	enemyend2v2         = rl.NewVector2(-64, 0)
	// level end
	enemyshopv2  = rl.NewVector2(700, 360)
	enemyshop2v2 = rl.NewVector2(-64, 0)
	// create level
	horizplaton, horizplatlr bool
	horizplatH, horizplatV   int

	// images
	weapongameud bool
	// launcher
	launchermenucount             = 1
	fullscreenOn                  = true
	averageDiffOn                 = true
	pixellauncherfade             = float32(0.5)
	selecty                       = int32(55)
	enemy7launcherV2              = rl.NewVector2(-64, 0)
	dinolauncherV2                = rl.NewVector2(-10, 122)
	framecountlauncher            int
	weaponud, pixellauncherfadeOn bool
	weaponlauncherV2              rl.Vector2
	// weather
	cloudsactive, fogactive, rainactive, snowactive, frograinactive bool
	// backgrounds
	choosebackcount int
	// difficulty
	easyDiffOn, difficultDiffOn bool

	// disasters
	createTornado, meteorOn, meteorLR, createMeteor, rainoffireOn, rainoffireTimerOn, shockBlocksOn, createShockBlocks, createDistortion bool
	tornadoStartBlock, meteorStartBlock, tornadoLBlock, tornadoRBlock, rainoffireTimer, rainoffireTimerFinal                             int

	// weather
	rainFrogOn, frogRainOn, tornadoLR, snowOn, createSnow bool
	snowTimer                                             int

	// shop items
	mushroomShieldOn, pigCoinHasDropped, pigShieldOn, pigeonShieldFlash1, pigeonShieldFlash2, pigeonShieldFlash3, knightJumpOn, knightBackFadeOn, skeletonShield1, skeletonShield2, skeletonShield3, skeletonShield4, skeletonHasDropped1, skeletonHasDropped2, skeletonHasDropped3, sawbladeLR, rainFrogTimerOn bool
	pigCoinTimer1, pigCoinTimer2, pigCoinTimer3, knightJumpHeight, skeletonDropTimer1, skeletonDropTimer2, skeletonDropTimer3, cannonTimer, rainFrogTimer                                                                                                                                                        int
	knightBackFade                                                                                                                                                                                                                                                                                               = float32(0.7)

	// level end
	noCoinsTextOn bool

	// shop
	shopPrice1, shopPrice2, shopPrice3                                                          int
	backShopRec1UpDown, backShopRec2UpDown, backShopRec3UpDown, backShopRec4UpDown, shuffleDone bool
	descTxtColor                                                                                rl.Color

	backShopRecY                                                                                                           = int32(0)
	backShopRecH1, backShopRecH2, backShopRecH3, backShopRecH4, backShopRecY1, backShopRecY2, backShopRecY3, backShopRecY4 int32

	// mini boss
	miniBossHP                                                                                                                                                                 = 5
	miniBossOn1, miniBossOn2, miniBossOn3, miniBossUpOn, miniBossJumpOn, miniBossFallOn, miniBossRollOn, miniBossWallJumpOn, miniBossJumpTimerOn, boss5RollOn, miniBossHPpause bool
	currentMiniBossBlock                                                                                                                                                       int
	bossTextX                                                                                                                                                                  = int32(screenW + 10)

	// debugging
	enemyMovementTEST string
	// check stuck enemies
	checkEnemy1On, checkEnemy2On, checkEnemy3On, checkEnemy4On, checkEnemy5On bool
	enemyPosition1MAP                                                         = make([]int, 20)
	enemyPosition2MAP                                                         = make([]int, 20)
	enemyPosition3MAP                                                         = make([]int, 20)

	// create level
	extraShapesOn, extraFeaturesOn, groundBlocksOn bool

	// power ups
	powerUpCollected, powerUpRotation1, powerUpRotation2, invisibleOn, bombOn bool
	powerUpCurrentActive                                                      int
	powerUpTimer5                                                             = 5

	bombBlock int

	// choose dino menu
	borderFade   = float32(1.0)
	borderFadeOn bool

	// options menu
	menuTextSpaceColor                                                                       = rl.Yellow
	menuColumnSelect                                                                         = 1
	disastersOn, tornadoOn, somethingHereOn, magicOn, electricShocksOn, platformDistortionOn bool
	// miniboss

	miniBossType int

	// enemy special
	enemy4SpecialOn, enemy5Shield2, enemy5Shield3, enemy7SpecialOn                                                                         bool
	enemy5Shield1                                                                                                                          = true
	enemy7Horizontal, enemy7Vertical, enemy7RoofFall1, enemy7RoofFall2, enemy7RoofFall3, enemy7RoofFall4, enemy7RoofFall5, enemy7RoofFall6 int

	// pixel noise scan lines
	pixelNoiseOn, switchScanLines, noiseLinesOn bool

	// MARK: intro story screen

	spaceToStart bool

	noiseLineX1, noiseLineX2, noiseLineX3, noiseLineX4, noiseLineX5, noiseLineX6, noiseLineX7, noiseLineX8, noiseLineX9, noiseLineX10, noiseLineX11, noiseLineX12 int32

	noiseLineX1Change, noiseLineX2Change, noiseLineX3Change, noiseLineX4Change, noiseLineX5Change, noiseLineX6Change, noiseLineX7Change, noiseLineX8Change, noiseLineX9Change, noiseLineX10Change, noiseLineX11Change, noiseLineX12Change int32

	noiseLineDistance1, noiseLineDistance2, noiseLineDistance3, noiseLineDistance4, noiseLineDistance5, noiseLineDistance6, noiseLineDistance7, noiseLineDistance8, noiseLineDistance9, noiseLineDistance10, noiseLineDistance11, noiseLineDistance12 int

	noiseLineLR1, noiseLineLR2, noiseLineLR3, noiseLineLR4, noiseLineLR5, noiseLineLR6, noiseLineLR7, noiseLineLR8, noiseLineLR9, noiseLineLR10, noiseLineLR11, noiseLineLR12 bool

	logoScount                                                                 = rInt(14, 21)
	logoStep1, logoStep2, logoStep3, logoStep4, logoSflash, noiseLinesScreenOn bool

	rectangleFade       = float32(1.0)
	introTextFade       = float32(0.0)
	introTextSpaceColor = rl.Yellow

	introCount2 = 4
	introCount1 = 2

	logoCount1                                                                                                     = 10
	raylibV2                                                                                                       = rl.NewVector2(-500, 340)
	gopherV2                                                                                                       = rl.NewVector2(-500, 314)
	introCirclesTimer                                                                                              int
	introBulletV2                                                                                                  rl.Vector2
	exclamationIntroV2                                                                                             = rl.NewVector2(-20, 330)
	dinoIntroV2                                                                                                    = rl.NewVector2(-16, 350)
	enemy1IntroV2                                                                                                  = rl.NewVector2(680, 348)
	introStoryScreenOn, introScreenShake, createBulletV2on, startLogosOn, logo1On, logo2On, logo3On, introPauseOff bool
	introStoryScreenTextY                                                                                          = int32(screenH - 100)
	introStoryScreenTextY2                                                                                         = int32(screenH - 40)
	introStoryScreenTextY3                                                                                         = int32(screenH + 20)
	introStoryScreenTextY4                                                                                         = int32(screenH + 80)
	introStoryScreenTextSize                                                                                       = int32(40)
	// add level shape
	addLevelShapeOn bool
	// weather
	weatherOn bool

	// level end
	levelEndMenuCount          int
	weaponLevelEndV2           = rl.NewVector2(148, 43)
	weaponUpDown, flashButtons bool
	// MARK: active special item
	pigeonStartBlock, sawbladeStartBlock, activeSpecialDirection, fireballStartBlock, cherriesStartBlock, poisonballStartBlock, iceballStartBlock, cannonStartBlock, petMushroomStartBlock, mushroomJumpHeight, petGreenPigStartBlock, watermelonStartBlock, petSkeletonStartBlock, petSlimeStartBlock, petKnightStartBlock, petRedSlimeStartBlock, petBatStartBlock, appleStartBlock int
	activeSpecialItem, checkPigeonBlock                                                                                                                                                                                                                                                                                                                                               string
	activeSpecialOn, activeSpecialActive, activeSpecialComplete, pigeonLeftRight, pigeonOn, trampolineOn, petMushroomLeftRight, mushroomJumpOn, exclamationOn, propellorOn, petGreenPigLeftRight, petSkeletonLeftRight, petSlimeLeftRight, petKnightLeftRight, petRedSlimeLeftRight, petBatLeftRight                                                                                  bool
	// teleport
	teleportOn, teleportActive                                        bool
	teleportPostion1, teleportPostion2, teleportTimer, teleportTimer2 int
	// shop
	shopOn                          bool
	coinsTextColor                  = rl.Gold
	shopMenuBorder                  = rl.Fade(rl.Red, 0.6)
	shopMenuCount                   int
	shopItem1, shopItem2, shopItem3 string
	// falling blocks
	fallingBlocksOn        bool
	currentFallingBlock    int
	fallingBlockTypeHOLDER string
	// mini boss
	miniBossTimer1, miniBossTimer2, miniBossTimer3, miniBossCurrentBlock, miniBossTimerText1, miniBossTimerText2, miniBossTimerText3, miniBossCountdown int
	miniBossRightLeft, miniBossOn, miniBossTextOn, miniBOSSActive                                                                                       bool
	// weapons
	weaponTextColor                                                                  rl.Color
	dropNewWeapon                                                                    bool
	dropWeaponCount, dropWeaponChange, currentDropWeaponBlock, weaponTextColorChange int
	dropWeaponHOLDER1, dropWeaponHOLDER2, currentPlayerWeapon                        string
	// help
	helpOn bool
	// MARK: game start
	dinoSelectMarkV2                                                                                                                           = rl.NewVector2(float32(screenW/2-93), float32(screenH/2+151))
	dinoRunV2                                                                                                                                  rl.Vector2
	gameStartOn, createIntroCircles, startIntroTimer, dinoIntroLeft, dinoIntroRight, chooseCharacterIntro, leftRightChooseDinoRec, flickerText bool
	frameCountGameStart, introTimerInterval, chooseCharacter, chooseMenuOption                                                                 int
	dinoTextColor                                                                                                                              rl.Color
	movingRecChooseDinoX                                                                                                                       = screenW / 3
	// MARK: weather / effects

	earthquakeCount, earthquakeTimer1, earthquakeTimerText1, earthquakeTimer2, earthquakeTimerText2, earthquakeTimer3, earthquakeTimerText3, earthquakeCountdown int
	cloudsOn, fogOn, fogRightLeft, rainOn, earthquakesOn, earthquakeFallOn, earthquakeTextOn, earthquakesActive                                                  bool
	fogX, fogY                                                                                                                                                   int32
	// platforms
	platformTypeName, platformTypeNameTL, floorTypeNameTL string
	movingPlatformsOn                                     bool
	// objects
	dropCoin, coinHasDropped1, coinHasDropped2, coinHasDropped3, coinHasDropped4, coinHasDropped5, coinHasDropped6, coinHasDropped7, coinHasDropped8, coinHasDropped9, coinHasDropped10 bool
	coinCurrentBlock                                                                                                                                                                    int
	// power ups
	powerUpVehicleCurrentBlock, powerUpTimer, currentPowerUpBlock, powerUpSlots, activePowerUpCount int
	powerUpDropped, powerUpHasDropped                                                               bool
	powerUpType, powerUpTypeTL                                                                      string
	// effects

	screenShake, screenShake1, screenShake2, screenShake3, screenShake4, screenShake5 bool
	screenShakeTimer                                                                  int
	// MARK: enemies variables
	clearEnemyBlock, enemiesScreenCount, killzCount, totalKillz, onScreenEnemiesCount, enemyTypeGenerate, enemyNumber int
	equalizeEnemyNumbers                                                                                              bool
	// draw screen
	drawScreenCurrentBlock, drawScreenCurrentBlockWeather, drawScreenCurrentBlockWeatherNEXT int
	// MARK: maps
	launcherpixelMAP         = make([]bool, 102400)
	shockBlockMAP            = make([]string, 5264)
	snowMAP                  = make([]string, 5264)
	rainoffireMAP            = make([]string, 5264)
	meteorMAP                = make([]string, 5264)
	tornadoMAP               = make([]string, 5264)
	rainFrogMAP              = make([]string, 5264)
	platformsEffectsMAP      = make([]string, 5264)
	effectsMAP               = make([]string, 5264)
	noiseLinesMAP            = make([]int32, 12)
	noiseLinesDistanceMAP    = make([]int, 12)
	noiseLinesLRMAP          = make([]bool, 12)
	pixelNoiseMAP            = make([]bool, 65472)
	starsMAP                 = make([]string, 5264)
	starsFadeMAP             = make([]int, 5264)
	activeSpecialMAP         = make([]string, 5264)
	shopItemsMAP             = []string{"petMushroom", "petGreenPig", "petBat", "petPigeon", "petKnight", "trampoline", "petSkeleton", "fireball", "watermelon", "cannon", "propellor", "petSlime", "petRedSlime", "poisonball", "sawblade", "heart", "exclamation", "greenHeart", "iceball", "apple", "cherries"}
	teleportsMAP             = make([]string, 5264)
	earthquakesMAP           = make([]string, 5264)
	miniBossMAP              = make([]string, 5264)
	miniBossEffectsMAP       = make([]string, 5264)
	rainMAP                  = make([]string, 5264)
	backgroundObjectsMAP     = make([]string, 5264)
	deadEnemiesCirclesRadius = make([]float32, 15)
	deadEnemiesCirclesV2     = make([]rl.Vector2, 15)
	weaponsMAP               = make([]string, 5264)
	colorsMAP                = make([]rl.Color, 10)
	introScreenCircleXMAP    = make([]int32, 30)
	introScreenCircleYMAP    = make([]int32, 30)
	introScreenCircleRadius  = make([]float32, 30)
	playerDiedPointsMAP      = make([]int, 30)
	playerDiedCirclesMAP     = make([]float32, 30)

	cloudsMAP           = make([]string, 526400)
	platformsMAP        = make([]string, 5264)
	objectsMAP          = make([]string, 5264)
	powerUpsMAP         = make([]string, 5264)
	powerUpVehicleMAP   = make([]string, 5264)
	bulletsMAP          = make([]string, 5264)
	enemiesMovementMAP  = make([]string, 5264)
	enemiesDirectionMAP = make([]string, 5264)
	deadEnemiesMAP      = make([]string, 5264)
	enemiesMAP          = make([]string, 5264)
	levelMAP            = make([]string, 5264)
	playerMAP           = make([]string, 5264)

	// MARK: images

	// images shock blocks
	shockBlockIMG = rl.NewRectangle(1082, 79, 21, 21)
	// images meteor
	meteorIMG = rl.NewRectangle(432, 534, 48, 48)
	// images weather
	snow1IMG     = rl.NewRectangle(1134, 743, 77, 81)
	snow2IMG     = rl.NewRectangle(1132, 829, 77, 81)
	snow3IMG     = rl.NewRectangle(1134, 915, 77, 81)
	snow4IMG     = rl.NewRectangle(1223, 874, 40, 40)
	snow5IMG     = rl.NewRectangle(1223, 916, 40, 40)
	snow6IMG     = rl.NewRectangle(1223, 960, 40, 40)
	rainFrogIMG  = rl.NewRectangle(668, 198, 25, 29)
	rainFrogRIMG = rl.NewRectangle(776, 199, 25, 29)
	// images effects
	flameIMG       = rl.NewRectangle(1045, 2, 25, 16)
	poisonFlameIMG = rl.NewRectangle(1045, 21, 25, 16)
	// images power ups
	bombIMG          = rl.NewRectangle(154, 200, 32, 30)
	specialItemIMG   = rl.NewRectangle(0, 236, 32, 30)
	randomPowerUpIMG = rl.NewRectangle(0, 300, 32, 30)
	invisibleIMG     = rl.NewRectangle(56, 201, 32, 30)
	hpPowerUpIMG     = rl.NewRectangle(0, 332, 32, 30)
	// intro images
	raylibIMG = rl.NewRectangle(1784, 611, 256, 256)
	gopherIMG = rl.NewRectangle(1568, 887, 480, 137)
	// images speech bubbles
	exclamationIMG = rl.NewRectangle(1295, 128, 32, 17)
	hiIMG          = rl.NewRectangle(1304, 171, 18, 17)
	swearingIMG    = rl.NewRectangle(1297, 128, 24, 17)
	loserIMG       = rl.NewRectangle(1326, 194, 32, 17)
	// images shop items
	petMushroomIMG  = rl.NewRectangle(1281, 227, 26, 26)
	petBatIMG       = rl.NewRectangle(1403, 211, 16, 13)
	petPigeonIMG    = rl.NewRectangle(1482, 223, 38, 29)
	petPigeonUpIMG  = rl.NewRectangle(1264, 261, 38, 29)
	petKnightIMG    = rl.NewRectangle(1413, 168, 16, 16)
	petKnightLIMG   = rl.NewRectangle(1172, 234, 16, 16)
	petRedSlimeIMG  = rl.NewRectangle(1262, 87, 14, 15)
	petGreenManIMG  = rl.NewRectangle(1240, 224, 39, 24)
	petRedManIMG    = rl.NewRectangle(1461, 221, 32, 24)
	petGreenPigIMG  = rl.NewRectangle(1678, 218, 24, 25)
	greenHeartIMG   = rl.NewRectangle(1399, 194, 16, 12)
	petSkeletonIMG  = rl.NewRectangle(1518, 163, 16, 16)
	petSkeletonLIMG = rl.NewRectangle(1252, 434, 16, 16)
	petSlimeIMG     = rl.NewRectangle(1407, 144, 16, 12)
	propellorIMG    = rl.NewRectangle(1413, 122, 23, 10)
	trampolineIMG   = rl.NewRectangle(1283, 86, 27, 29)
	cherriesIMG     = rl.NewRectangle(1504, 83, 32, 22)
	appleIMG        = rl.NewRectangle(1504, 114, 32, 22)
	watermelonIMG   = rl.NewRectangle(1504, 141, 32, 20)
	heartShopIMG    = rl.NewRectangle(1595, 169, 12, 13)
	cannonIMG       = rl.NewRectangle(1663, 190, 44, 21)
	sawbladeIMG     = rl.NewRectangle(1740, 166, 38, 38)
	fireballIMG     = rl.NewRectangle(1980, 0, 68, 9)
	fireballLIMG    = rl.NewRectangle(1374, 11, 68, 13)
	poisonballIMG   = rl.NewRectangle(1985, 26, 63, 9)
	iceballIMG      = rl.NewRectangle(1986, 50, 62, 9)
	// images teleport
	teleportIMG = rl.NewRectangle(8, 612, 44, 52)
	// images background objects
	backObjGround1IMG  = rl.NewRectangle(834, 302, 47, 76)
	backObjGround2IMG  = rl.NewRectangle(882, 300, 47, 76)
	backObjGround3IMG  = rl.NewRectangle(930, 305, 92, 59)
	backObjGround4IMG  = rl.NewRectangle(713, 367, 67, 66)
	backObjGround5IMG  = rl.NewRectangle(784, 379, 68, 63)
	backObjGround6IMG  = rl.NewRectangle(853, 389, 18, 49)
	backObjGround7IMG  = rl.NewRectangle(872, 389, 18, 49)
	backObjGround8IMG  = rl.NewRectangle(894, 406, 25, 31)
	backObjGround9IMG  = rl.NewRectangle(932, 402, 35, 34)
	backObjGround10IMG = rl.NewRectangle(969, 406, 53, 42)
	backObjGround11IMG = rl.NewRectangle(969, 456, 53, 42)
	backObjGround12IMG = rl.NewRectangle(717, 443, 37, 68)
	backObjGround13IMG = rl.NewRectangle(755, 451, 63, 58)
	backObjGround14IMG = rl.NewRectangle(818, 444, 63, 66)
	backObjGround15IMG = rl.NewRectangle(988, 369, 31, 33)
	backObjGround16IMG = rl.NewRectangle(644, 375, 70, 65)
	backObjGround17IMG = rl.NewRectangle(661, 451, 54, 61)
	backObjGround18IMG = rl.NewRectangle(915, 513, 109, 134)
	backObjGround19IMG = rl.NewRectangle(805, 512, 111, 137)
	backObjGround20IMG = rl.NewRectangle(672, 513, 128, 154)
	backObjGround21IMG = rl.NewRectangle(615, 523, 50, 138)
	backObjGround22IMG = rl.NewRectangle(597, 664, 81, 62)
	backObjGround23IMG = rl.NewRectangle(622, 726, 51, 61)
	backObjGround24IMG = rl.NewRectangle(682, 677, 125, 115)
	backObjGround25IMG = rl.NewRectangle(595, 788, 72, 101)
	backObjGround26IMG = rl.NewRectangle(669, 804, 66, 90)
	backObjGround27IMG = rl.NewRectangle(740, 799, 70, 131)
	backObjGround28IMG = rl.NewRectangle(813, 797, 207, 183)
	backObjGround29IMG = rl.NewRectangle(757, 953, 50, 61)
	backObjGround30IMG = rl.NewRectangle(668, 921, 73, 91)
	backObjGround31IMG = rl.NewRectangle(580, 908, 77, 105)
	backObjGround32IMG = rl.NewRectangle(489, 873, 91, 138)
	backObjGround33IMG = rl.NewRectangle(818, 662, 65, 26)
	backObjGround34IMG = rl.NewRectangle(889, 656, 67, 31)
	backObjGround35IMG = rl.NewRectangle(958, 658, 65, 30)
	backObjGround36IMG = rl.NewRectangle(580, 908, 77, 105)
	backObjGround37IMG = rl.NewRectangle(963, 692, 59, 31)
	backObjGround38IMG = rl.NewRectangle(922, 691, 35, 33)
	backObjGround39IMG = rl.NewRectangle(827, 696, 91, 28)
	backObjGround40IMG = rl.NewRectangle(836, 736, 50, 59)
	backObjGround41IMG = rl.NewRectangle(893, 734, 66, 61)
	backObjGround42IMG = rl.NewRectangle(966, 728, 55, 66)
	backObjGround43IMG = rl.NewRectangle(841, 247, 30, 48)
	backObjGround44IMG = rl.NewRectangle(713, 274, 45, 81)

	backObj1IMG  = rl.NewRectangle(764, 269, 20, 20)
	backObj2IMG  = rl.NewRectangle(784, 247, 20, 39)
	backObj3IMG  = rl.NewRectangle(810, 250, 25, 38)
	backObj4IMG  = rl.NewRectangle(760, 294, 32, 30)
	backObj5IMG  = rl.NewRectangle(795, 295, 30, 30)
	backObj6IMG  = rl.NewRectangle(803, 328, 20, 25)
	backObj7IMG  = rl.NewRectangle(776, 331, 16, 26)
	backObj8IMG  = rl.NewRectangle(758, 332, 16, 24)
	backObj9IMG  = rl.NewRectangle(916, 382, 20, 13)
	backObj10IMG = rl.NewRectangle(940, 364, 42, 17)
	backObj11IMG = rl.NewRectangle(946, 381, 32, 18)
	backObj12IMG = rl.NewRectangle(886, 443, 16, 20)
	backObj13IMG = rl.NewRectangle(883, 466, 18, 22)
	backObj14IMG = rl.NewRectangle(884, 489, 18, 19)
	backObj15IMG = rl.NewRectangle(906, 451, 18, 18)
	backObj16IMG = rl.NewRectangle(906, 469, 18, 17)
	backObj17IMG = rl.NewRectangle(906, 490, 19, 18)
	backObj18IMG = rl.NewRectangle(927, 442, 21, 19)
	backObj19IMG = rl.NewRectangle(927, 461, 20, 16)
	backObj20IMG = rl.NewRectangle(928, 477, 18, 16)
	backObj21IMG = rl.NewRectangle(928, 493, 19, 16)
	backObj22IMG = rl.NewRectangle(950, 444, 17, 15)
	backObj23IMG = rl.NewRectangle(950, 476, 17, 15)
	backObj24IMG = rl.NewRectangle(949, 493, 19, 13)
	// images miscellaneous
	flyingRaddishIMG = rl.NewRectangle(843, 0, 30, 38)
	// images enemies
	enemy9IMG   = rl.NewRectangle(1469, 257, 52, 36) // elephant
	enemy8IMG   = rl.NewRectangle(1600, 333, 36, 36) // ghost
	enemy7IMG   = rl.NewRectangle(1241, 339, 43, 30) // spike man
	enemy7URIMG = rl.NewRectangle(1088, 970, 30, 43) // spike man
	enemy7ULIMG = rl.NewRectangle(1056, 972, 30, 43) // spike man
	enemy7UPIMG = rl.NewRectangle(1218, 479, 43, 30) // spike man
	enemy6IMG   = rl.NewRectangle(1134, 298, 46, 36) // bat
	enemy5IMG   = rl.NewRectangle(1348, 430, 48, 48) // tree
	enemy4IMG   = rl.NewRectangle(1587, 476, 32, 34) // rock man
	enemy3IMG   = rl.NewRectangle(1108, 379, 34, 44) // flower
	enemy2IMG   = rl.NewRectangle(1591, 378, 32, 36) // chicken
	enemy2RIMG  = rl.NewRectangle(1585, 516, 32, 36) // chicken
	enemy1IMG   = rl.NewRectangle(514, 72, 28, 28)   // mushroom
	// images mini boss
	miniBoss1IMG      = rl.NewRectangle(1, 518, 32, 44)    // pink rabbit
	miniBoss1RIMG     = rl.NewRectangle(1, 568, 32, 44)    // pink rabbit
	miniBoss2IMG      = rl.NewRectangle(0, 706, 32, 32)    // pink man
	miniBoss2RIMG     = rl.NewRectangle(0, 667, 32, 32)    // pink man
	miniBoss2RollIMG  = rl.NewRectangle(395, 705, 32, 32)  // pink man roll
	miniBoss2RRollIMG = rl.NewRectangle(395, 673, 32, 32)  // pink man roll
	miniBoss3IMG      = rl.NewRectangle(0, 779, 32, 32)    // blue man
	miniBoss3RIMG     = rl.NewRectangle(0, 741, 32, 32)    // blue man
	miniBoss3UpIMG    = rl.NewRectangle(396, 745, 32, 32)  // blue man up wall
	miniBoss3RUpIMG   = rl.NewRectangle(396, 778, 32, 32)  // blue man up wall
	miniBoss4RIMG     = rl.NewRectangle(0, 817, 32, 32)    // tree stump
	miniBoss4IMG      = rl.NewRectangle(0, 852, 32, 32)    // tree stump
	miniBoss5IMG      = rl.NewRectangle(0, 925, 32, 32)    // frog man
	miniBoss5RIMG     = rl.NewRectangle(0, 888, 32, 32)    // frog man
	miniBoss5RollIMG  = rl.NewRectangle(238, 813, 32, 32)  // frog man roll
	miniBoss5RRollIMG = rl.NewRectangle(238, 852, 32, 32)  // frog man roll
	miniBoss6IMG      = rl.NewRectangle(1050, 514, 40, 50) // chicken
	miniBoss7IMG      = rl.NewRectangle(640, 110, 36, 36)  // duck
	miniBoss7RIMG     = rl.NewRectangle(639, 151, 36, 36)  // duck
	// images popups
	getArrowIMG = rl.NewRectangle(0, 205, 17, 21)
	// images keyboard
	arrowKeysIMG = rl.NewRectangle(0, 363, 229, 151)
	f1KeyIMG     = rl.NewRectangle(231, 363, 78, 75)
	escKeyIMG    = rl.NewRectangle(231, 437, 78, 75)
	altKeyIMG    = rl.NewRectangle(316, 377, 94, 65)
	ctrlKeyIMG   = rl.NewRectangle(316, 444, 94, 65)
	// images ground tiles
	ground1IMG  = rl.NewRectangle(336, 143, 48, 48)
	ground2IMG  = rl.NewRectangle(384, 143, 48, 48)
	ground3IMG  = rl.NewRectangle(432, 143, 48, 48)
	ground4IMG  = rl.NewRectangle(480, 143, 48, 48)
	ground5IMG  = rl.NewRectangle(528, 143, 48, 48)
	ground6IMG  = rl.NewRectangle(576, 143, 48, 48)
	ground7IMG  = rl.NewRectangle(336, 191, 48, 48)
	ground8IMG  = rl.NewRectangle(384, 191, 48, 48)
	ground9IMG  = rl.NewRectangle(432, 191, 48, 48)
	ground10IMG = rl.NewRectangle(480, 191, 48, 48)
	ground11IMG = rl.NewRectangle(528, 191, 48, 48)
	ground12IMG = rl.NewRectangle(576, 191, 48, 48)
	ground13IMG = rl.NewRectangle(336, 239, 48, 48)
	ground14IMG = rl.NewRectangle(384, 239, 48, 48)
	ground15IMG = rl.NewRectangle(432, 239, 48, 48)
	ground16IMG = rl.NewRectangle(480, 239, 48, 48)
	ground17IMG = rl.NewRectangle(528, 239, 48, 48)
	ground18IMG = rl.NewRectangle(576, 239, 48, 48)
	ground19IMG = rl.NewRectangle(336, 287, 48, 48)
	ground20IMG = rl.NewRectangle(384, 287, 48, 48)
	ground21IMG = rl.NewRectangle(432, 287, 48, 48)
	ground22IMG = rl.NewRectangle(480, 287, 48, 48)
	ground23IMG = rl.NewRectangle(528, 287, 48, 48)
	ground24IMG = rl.NewRectangle(576, 287, 48, 48)
	// images platform tiles
	platform1TileIMG  = rl.NewRectangle(334, 19, 16, 16)
	platform2TileIMG  = rl.NewRectangle(349, 19, 16, 16)
	platform3TileIMG  = rl.NewRectangle(364, 19, 16, 16)
	platform4TileIMG  = rl.NewRectangle(379, 19, 16, 16)
	platform5TileIMG  = rl.NewRectangle(395, 19, 16, 16)
	platform6TileIMG  = rl.NewRectangle(409, 19, 16, 16)
	platform7TileIMG  = rl.NewRectangle(425, 19, 16, 16)
	platform8TileIMG  = rl.NewRectangle(334, 35, 16, 16)
	platform9TileIMG  = rl.NewRectangle(349, 34, 16, 16)
	platform10TileIMG = rl.NewRectangle(365, 35, 16, 16)
	platform11TileIMG = rl.NewRectangle(381, 35, 16, 16)
	platform12TileIMG = rl.NewRectangle(397, 35, 16, 16)
	platform13TileIMG = rl.NewRectangle(412, 35, 16, 16)
	platform14TileIMG = rl.NewRectangle(428, 35, 16, 16)
	platform15TileIMG = rl.NewRectangle(334, 51, 16, 16)
	platform16TileIMG = rl.NewRectangle(350, 51, 16, 16)
	platform17TileIMG = rl.NewRectangle(366, 51, 16, 16)
	platform18TileIMG = rl.NewRectangle(382, 51, 16, 16)
	platform19TileIMG = rl.NewRectangle(398, 51, 16, 16)
	platform20TileIMG = rl.NewRectangle(414, 51, 16, 16)
	platform21TileIMG = rl.NewRectangle(430, 51, 16, 16)
	platform22TileIMG = rl.NewRectangle(334, 67, 16, 16)
	platform23TileIMG = rl.NewRectangle(350, 67, 16, 16)
	platform24TileIMG = rl.NewRectangle(366, 67, 16, 16)
	platform25TileIMG = rl.NewRectangle(382, 67, 16, 16)
	platform26TileIMG = rl.NewRectangle(398, 67, 16, 16)
	platform27TileIMG = rl.NewRectangle(414, 67, 16, 16)
	platform28TileIMG = rl.NewRectangle(430, 67, 16, 16)
	platform29TileIMG = rl.NewRectangle(446, 67, 16, 16)
	platform30TileIMG = rl.NewRectangle(462, 67, 16, 16)
	platform31TileIMG = rl.NewRectangle(334, 83, 16, 16)
	platform32TileIMG = rl.NewRectangle(350, 83, 16, 16)
	platform33TileIMG = rl.NewRectangle(366, 83, 16, 16)
	platform34TileIMG = rl.NewRectangle(382, 83, 16, 16)
	platform35TileIMG = rl.NewRectangle(398, 83, 16, 16)
	platform36TileIMG = rl.NewRectangle(414, 83, 16, 16)
	platform37TileIMG = rl.NewRectangle(430, 83, 16, 16)
	platform38TileIMG = rl.NewRectangle(446, 83, 16, 16)
	platform39TileIMG = rl.NewRectangle(462, 83, 16, 16)
	platform40TileIMG = rl.NewRectangle(334, 99, 16, 16)
	platform41TileIMG = rl.NewRectangle(350, 99, 16, 16)
	platform42TileIMG = rl.NewRectangle(366, 99, 16, 16)
	platform43TileIMG = rl.NewRectangle(382, 99, 16, 16)
	// images objects
	coinIMG = rl.NewRectangle(334, 0, 16, 16)
	// images backgrounds
	backgroundIMG = rl.NewRectangle(0, 0, 1366, 768)
	// images weapons
	weapon1IMG  = rl.NewRectangle(416, 501, 14, 9)
	weapon1LIMG = rl.NewRectangle(519, 489, 14, 9)
	weapon2IMG  = rl.NewRectangle(432, 500, 16, 10)
	weapon2LIMG = rl.NewRectangle(501, 488, 16, 10)
	weapon3IMG  = rl.NewRectangle(449, 502, 16, 7)
	weapon3LIMG = rl.NewRectangle(484, 490, 16, 7)
	weapon4IMG  = rl.NewRectangle(466, 500, 16, 10)
	weapon4LIMG = rl.NewRectangle(467, 488, 16, 10)
	weapon5IMG  = rl.NewRectangle(484, 499, 14, 12)
	weapon5LIMG = rl.NewRectangle(451, 487, 14, 12)
	weapon6IMG  = rl.NewRectangle(500, 500, 16, 10)
	weapon6LIMG = rl.NewRectangle(433, 488, 16, 10)
	weapon7IMG  = rl.NewRectangle(517, 500, 16, 10)
	weapon7LIMG = rl.NewRectangle(416, 488, 16, 10)
	// images players
	heartIMG       = rl.NewRectangle(456, 0, 56, 48)
	shadowIMG      = rl.NewRectangle(5, 193, 18, 7)
	dinoGreenRIMG  = rl.NewRectangle(4, 4, 16, 18)
	dinoGreenLIMG  = rl.NewRectangle(317, 100, 16, 18)
	dinoRedRIMG    = rl.NewRectangle(5, 28, 16, 18)
	dinoRedLIMG    = rl.NewRectangle(315, 124, 16, 18)
	dinoYellowRIMG = rl.NewRectangle(4, 52, 16, 18)
	dinoYellowLIMG = rl.NewRectangle(317, 148, 16, 18)
	dinoBlueRIMG   = rl.NewRectangle(4, 76, 16, 18)
	dinoBlueLIMG   = rl.NewRectangle(317, 172, 16, 18)
	dinoV2         rl.Vector2
	shadowV2       rl.Vector2
	// images resources
	imgs              rl.Texture2D
	backgroundTexture rl.Texture2D
	// player
	jumpActive, fallActive, hpLossPause, playerDied                                                                                       bool
	playerCurrentBlock, playerVertical, playerHorizontal, jumpHeight, playerHP, playerHPmax, collisionCount, playerDiedBlock, playerCoins int
	playerDirection, playerGroundBlockL, playerGroundBlockR                                                                               string
	// timers
	frameCount, secondsCountdown int
	// menu
	menuOn, gameMenuOn, grayscaleOn, altColorsOn, chalkOn, pencilOn, inkOn, standardBackOn, randomBackOn bool
	menuSelectBoxY                                                                                       = int32(95)
	menuSelectNumber                                                                                     = 1
	// debugging
	debuggingOn bool
	// game
	dinoType                                                       string
	level, backgroundSelect, totalLevels                           int
	pauseOn, introScreenOn, levelEnd, creditsScreenOn, scanLinesOn bool
	fadeAmount                                                     = float32(1.0)
	// screen
	camera         rl.Camera2D
	cameraIntro    rl.Camera2D
	cameraShop     rl.Camera2D
	cameraShop2    rl.Camera2D
	cameraShop3    rl.Camera2D
	cameraShop4    rl.Camera2D
	cameraShop5    rl.Camera2D
	cameraMiniBoss rl.Camera2D
	screenW        = int32(1366)
	screenH        = int32(768)
)

/*

	1366 X 768 = 86 X 48 16px blocks
	2X ZOOM therefore = 43 X 24 = 1032

	border + 4 blocks = 47 X 28
	8px X2 = 94 X 56




*/
// MARK: start new level
func startNewLevel() {
	levelEndMusic = false
	clearMAPS()
	cNEWLEVEL()
	cLEVEL()
	cENEMIES()
	cCLOUDS()
	cRAIN()
	randomColors()
	chooseLevelMusic()
	levelMusic = true
	levelEnd = false
	pauseOn = false
}

// MARK: clear maps
func clearMAPS() {
	for a := 0; a < 5264; a++ {
		effectsMAP[a] = ""
		teleportsMAP[a] = ""
		activeSpecialMAP[a] = ""
		weaponsMAP[a] = ""
		platformsMAP[a] = ""
		powerUpsMAP[a] = ""
		objectsMAP[a] = ""
		bulletsMAP[a] = ""
		enemiesMovementMAP[a] = ""
		deadEnemiesMAP[a] = ""
		enemiesMAP[a] = ""
		playerMAP[a] = ""
		powerUpsMAP[a] = ""
		powerUpVehicleMAP[a] = ""
		backgroundObjectsMAP[a] = ""
		miniBossMAP[a] = ""
		earthquakesMAP[a] = ""
		levelMAP[a] = ""
		platformsEffectsMAP[a] = ""
		tornadoMAP[a] = ""
		meteorMAP[a] = ""
		shockBlockMAP[a] = ""
		snowMAP[a] = ""
		rainFrogMAP[a] = ""
		rainoffireMAP[a] = ""
	}
	for a := 0; a < 526400; a++ {
		cloudsMAP[a] = ""
	}
}

// MARK: new level variables
func cNEWLEVEL() {
	chooseeffects()
	chooseweather()
	enemyendv2 = rl.NewVector2(700, 360)
	enemyend2v2 = rl.NewVector2(-64, 0)
	enemyshopv2 = rl.NewVector2(700, 360)
	enemyshop2v2 = rl.NewVector2(-64, 0)
	horizplaton = false
	movingPlatformsOn = false
	extraFeaturesOn = false
	extraShapesOn = false
	snowTimer = 2
	rainoffireTimerFinal = 5
	rainFrogTimer = 1
	cannonTimer = 5
	skeletonDropTimer1 = rInt(10, 22)
	skeletonDropTimer2 = rInt(23, 35)
	skeletonDropTimer3 = rInt(36, 48)
	skeletonShield1 = true
	skeletonShield2 = false
	skeletonShield3 = false
	skeletonShield4 = false
	pigeonShieldFlash1 = true
	pigeonShieldFlash2 = false
	pigeonShieldFlash3 = false
	petSkeletonStartBlock = 3988
	petKnightStartBlock = 3988
	pigeonStartBlock = 276
	petBatStartBlock = 960
	petMushroomStartBlock = 3800
	petGreenPigStartBlock = 3800
	noCoinsTextOn = false
	shuffleDone = false
	backShopRecY1 = rInt32(-100, int(screenH+100))
	backShopRecY2 = rInt32(-100, int(screenH+100))
	backShopRecY3 = rInt32(-100, int(screenH+100))
	backShopRecY4 = rInt32(-100, int(screenH+100))
	for {
		if backShopRecY1 == backShopRecY2 || backShopRecY1 == backShopRecY3 || backShopRecY1 == backShopRecY4 {
			if backShopRecY1 > screenH/2 {
				backShopRecY1--
			} else {
				backShopRecY1++
			}
		}
		if backShopRecY2 == backShopRecY3 || backShopRecY2 == backShopRecY4 {
			if backShopRecY2 > screenH/2 {
				backShopRecY2--
			} else {
				backShopRecY2++
			}
		}
		if backShopRecY3 == backShopRecY4 {
			if backShopRecY3 > screenH/2 {
				backShopRecY3--
			} else {
				backShopRecY3++
			}
		}
		if backShopRecY1 != backShopRecY2 && backShopRecY2 != backShopRecY3 && backShopRecY3 != backShopRecY4 && backShopRecY1 != backShopRecY4 && backShopRecY2 != backShopRecY4 && backShopRecY3 != backShopRecY1 {
			break
		}
	}

	backShopRecH1 = rInt32(20, 50)
	backShopRecH2 = rInt32(4, 12)
	backShopRecH3 = rInt32(10, 20)
	backShopRecH4 = rInt32(2, 7)

	miniBossHP = 5
	miniBossOn1 = false
	miniBossOn2 = false
	miniBossOn3 = false

	checkEnemy1On = true
	checkEnemy2On = false
	checkEnemy3On = false

	introPauseOff = false
	equalizeEnemyNumbers = false
	chooseTeleport := flipcoin()
	if chooseTeleport {
		teleportActive = true
	} else {
		teleportActive = false
	}
	teleportTimer = rInt(10, 24)
	teleportTimer2 = teleportTimer + 10

	levelEndMenuCount = 1
	totalLevels++
	killzCount = 0
	shopMenuCount = 1
	earthquakeCount = 2
	miniBossTimer1 = rInt(39, 45)
	miniBossTimer2 = rInt(25, 30)
	miniBossTimer3 = rInt(10, 14)
	miniBossTimerText1 = miniBossTimer1 + 3
	miniBossTimerText2 = miniBossTimer2 + 3
	miniBossTimerText3 = miniBossTimer3 + 3
	earthquakeTimer1 = miniBossTimerText1 + rInt(2, 7)
	earthquakeTimer2 = miniBossTimerText2 + rInt(2, 7)
	earthquakeTimer3 = miniBossTimerText3 + rInt(2, 7)
	earthquakeTimerText1 = earthquakeTimer1 + 3
	earthquakeTimerText2 = earthquakeTimer2 + 3
	earthquakeTimerText3 = earthquakeTimer3 + 3
	screenShakeTimer = 4
	currentPlayerWeapon = "weapon1TL"
	dropWeaponCount = rInt(3, 10)
	dinoRunV2 = rl.NewVector2(float32(screenW/3+100), float32(screenH/2+130))
	drawScreenCurrentBlockWeatherNEXT = 18804
	powerUpTimer = rInt(300, 500)
	playerHP = 3
	playerHPmax = playerHP
	playerCurrentBlock = 3700
	playerMAP[playerCurrentBlock] = "playerTL"
	playerDirection = "right"
	secondsCountdown = 60
	camera.Zoom = 2.0
	cameraIntro.Zoom = 6.0
	cameraShop.Zoom = 5.0
	cameraShop2.Zoom = 3.4
	cameraShop3.Zoom = 2.4
	cameraShop4.Zoom = 6.0
	cameraShop5.Zoom = 4.2
	cameraMiniBoss.Zoom = 3.0
	changeBackgroundImage()

}

// MARK: start game variables
func cSTARTGAMEVARIABLES() {

	changediff()
	chooseeffects()
	chooseweather()
	changeBackgroundImage()
	pixelNoiseOn = true
	scanLinesOn = true
	noiseLinesOn = true
	snowTimer = 2
	rainoffireTimerFinal = 5
	rainFrogTimer = 1
	cannonTimer = 5
	skeletonDropTimer1 = rInt(10, 22)
	skeletonDropTimer2 = rInt(23, 35)
	skeletonDropTimer3 = rInt(36, 48)
	skeletonShield1 = true
	skeletonShield2 = false
	skeletonShield3 = false
	skeletonShield4 = false
	pigeonShieldFlash1 = true
	pigeonShieldFlash2 = false
	pigeonShieldFlash3 = false
	petSkeletonStartBlock = 3988
	petKnightStartBlock = 3988
	pigeonStartBlock = 276
	petBatStartBlock = 960
	petGreenPigStartBlock = 3800
	petMushroomStartBlock = 3800
	noCoinsTextOn = false
	shuffleDone = false
	backShopRecH1 = rInt32(20, 50)
	backShopRecH2 = rInt32(4, 12)
	backShopRecH3 = rInt32(10, 20)
	backShopRecH4 = rInt32(2, 7)
	backShopRecY1 = rInt32(-100, int(screenH+100))
	backShopRecY2 = rInt32(-100, int(screenH+100))
	backShopRecY3 = rInt32(-100, int(screenH+100))
	backShopRecY4 = rInt32(-100, int(screenH+100))
	for {
		if backShopRecY1 == backShopRecY2 || backShopRecY1 == backShopRecY3 || backShopRecY1 == backShopRecY4 {
			if backShopRecY1 > screenH/2 {
				backShopRecY1--
			} else {
				backShopRecY1++
			}
		}
		if backShopRecY2 == backShopRecY3 || backShopRecY2 == backShopRecY4 {
			if backShopRecY2 > screenH/2 {
				backShopRecY2--
			} else {
				backShopRecY2++
			}
		}
		if backShopRecY3 == backShopRecY4 {
			if backShopRecY3 > screenH/2 {
				backShopRecY3--
			} else {
				backShopRecY3++
			}
		}
		if backShopRecY1 != backShopRecY2 && backShopRecY2 != backShopRecY3 && backShopRecY3 != backShopRecY4 && backShopRecY1 != backShopRecY4 && backShopRecY2 != backShopRecY4 && backShopRecY3 != backShopRecY1 {
			break
		}

	}
	miniBossHP = 5
	miniBossOn1 = false
	miniBossOn2 = false
	miniBossOn3 = false
	checkEnemy1On = true
	checkEnemy2On = false
	checkEnemy3On = false
	introPauseOff = false
	logo1On = true
	startLogosOn = true
	equalizeEnemyNumbers = false
	chooseTeleport := flipcoin()
	if chooseTeleport {
		teleportActive = true
	} else {
		teleportActive = false
	}
	teleportTimer = rInt(10, 24)
	teleportTimer2 = teleportTimer + 10
	levelEndMenuCount = 1
	totalLevels++
	killzCount = 0
	shopMenuCount = 1
	earthquakeCount = 2
	// earthquakesOn = true
	miniBossTimer1 = rInt(39, 45)
	miniBossTimer2 = rInt(25, 30)
	miniBossTimer3 = rInt(10, 14)
	miniBossTimerText1 = miniBossTimer1 + 3
	miniBossTimerText2 = miniBossTimer2 + 3
	miniBossTimerText3 = miniBossTimer3 + 3
	earthquakeTimer1 = miniBossTimerText1 + rInt(2, 7)
	earthquakeTimer2 = miniBossTimerText2 + rInt(2, 7)
	earthquakeTimer3 = miniBossTimerText3 + rInt(2, 7)
	earthquakeTimerText1 = earthquakeTimer1 + 3
	earthquakeTimerText2 = earthquakeTimer2 + 3
	earthquakeTimerText3 = earthquakeTimer3 + 3
	screenShakeTimer = 4
	currentPlayerWeapon = "weapon1TL"
	dropWeaponCount = rInt(3, 10)
	chooseMenuOption = 1
	dinoRunV2 = rl.NewVector2(float32(screenW/3+100), float32(screenH/2+130))
	dinoIntroRight = true
	pauseOn = true
	gameStartOn = true
	drawScreenCurrentBlockWeatherNEXT = 18804
	powerUpTimer = rInt(300, 500)
	playerHP = 3
	playerHPmax = playerHP
	playerCurrentBlock = 3700
	playerMAP[playerCurrentBlock] = "playerTL"
	playerDirection = "right"
	secondsCountdown = 60
	camera.Zoom = 2.0
	cameraIntro.Zoom = 6.0
	cameraShop.Zoom = 5.0
	cameraShop2.Zoom = 3.4
	cameraShop3.Zoom = 2.4
	cameraShop4.Zoom = 6.0
	cameraShop5.Zoom = 4.2
	cameraMiniBoss.Zoom = 3.0
}

// MARK: launcher window
func pixellauncherback() {
	for a := 0; a < 102400; a++ {
		placepixel := rolldice() + rolldice() + rolldice() + rolldice()
		if placepixel == 24 {
			launcherpixelMAP[a] = true
		}
	}
}
func launcher() {

	pixellauncherback()

	rl.InitWindow(320, 320, "spl@ta launcher")
	imgs = rl.LoadTexture("imgs_splata.png")
	rl.SetTargetFPS(30)
	for !rl.WindowShouldClose() {
		framecountlauncher++

		if framecountlauncher%4 == 0 {
			dinoYellowRIMG.X += 24
			if dinoYellowRIMG.X >= 100 {
				dinoYellowRIMG.X = 4
			}
		}
		if framecountlauncher%3 == 0 {
			enemy7UPIMG.X += 44
			if enemy7UPIMG.X >= 1560 {
				enemy7UPIMG.X = 1218
			}
		}
		if framecountlauncher%4 == 0 {

			if pixellauncherfadeOn {
				pixellauncherfade -= 0.1
				if pixellauncherfade <= 0 {
					pixellauncherfadeOn = false
				}
			} else {
				pixellauncherfade += 0.1
				if pixellauncherfade >= 0.5 {
					pixellauncherfadeOn = true
				}
			}

			if weaponud {
				weaponud = false
			} else {
				weaponud = true
			}
		}
		if weaponud {
			weaponlauncherV2 = rl.NewVector2(dinolauncherV2.X+12, dinolauncherV2.Y+4)
		} else {
			weaponlauncherV2 = rl.NewVector2(dinolauncherV2.X+12, dinolauncherV2.Y+5)
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		pixelx := int32(0)
		pixely := int32(0)
		for a := 0; a < 102400; a++ {
			drawpixel := launcherpixelMAP[a]
			if drawpixel {
				rl.DrawPixel(pixelx, pixely, rl.Fade(rl.Maroon, pixellauncherfade))
			}

			pixelx++
			if pixelx == 320 {
				pixely++
				pixelx = 0
			}
		}
		dinolauncherV2.X++
		if dinolauncherV2.X > 330 {
			dinolauncherV2.X = -10
		}
		enemy7launcherV2.X -= 2
		if enemy7launcherV2.X <= -120 {
			enemy7launcherV2.X = 190
		}
		camera.Zoom = 2.0
		rl.BeginMode2D(camera)
		rl.DrawTextureRec(imgs, weapon7IMG, weaponlauncherV2, rl.Fade(rl.White, 0.3))
		rl.DrawTextureRec(imgs, dinoYellowRIMG, dinolauncherV2, rl.Fade(rl.White, 0.3))
		rl.DrawTextureRec(imgs, enemy7UPIMG, enemy7launcherV2, rl.Fade(rl.White, 0.3))

		rl.EndMode2D()
		if selecty < 236 {
			rl.DrawRectangle(0, selecty, 320, 30, rl.Fade(rl.Red, 0.2))
			rl.DrawRectangle(0, 280, 320, 40, rl.White)
			rl.DrawText("start spl@ta", 90, 290, 20, rl.Black)
		} else {
			rl.DrawRectangle(0, selecty, 320, 40, rl.Fade(rl.Red, 0.2))
			rl.DrawText("start spl@ta", 90, 290, 20, rl.White)
		}
		if rl.IsKeyPressed(rl.KeyDown) {
			if selecty < 85 {
				selecty += 30
				launchermenucount++
			} else if selecty == 85 {
				selecty += 90
				launchermenucount++
			} else if selecty >= 175 && selecty < 230 {
				selecty += 30
				launchermenucount++
			} else if selecty == 235 {
				selecty += 45
				launchermenucount++
			}
		}
		if rl.IsKeyPressed(rl.KeyUp) {

			if selecty > 84 && selecty < 175 {
				selecty -= 30
				launchermenucount--
			} else if selecty == 175 {
				selecty -= 90
				launchermenucount--
			} else if selecty > 175 && selecty < 236 {
				selecty -= 30
				launchermenucount--
			} else if selecty == 280 {
				selecty -= 45
				launchermenucount--
			}
		}
		if rl.IsKeyPressed(rl.KeySpace) {

			switch launchermenucount {
			case 1:
				if fullscreenOn {
					fullscreenOn = false
				} else {
					fullscreenOn = true
				}
			case 2:
				if fullscreenOn {
					fullscreenOn = false
				} else {
					fullscreenOn = true
				}
			case 3:
				if easyDiffOn {
					easyDiffOn = false
					averageDiffOn = true
				} else {
					easyDiffOn = true
					averageDiffOn = false
					difficultDiffOn = false
				}
			case 4:
				if averageDiffOn {
					averageDiffOn = false
					difficultDiffOn = true
					easyDiffOn = false
				} else {
					averageDiffOn = true
					difficultDiffOn = false
					easyDiffOn = false
				}
			case 5:
				if difficultDiffOn {
					difficultDiffOn = false
					easyDiffOn = true
					averageDiffOn = false
				} else {
					difficultDiffOn = true
					averageDiffOn = false
					easyDiffOn = false
				}
			}
		}
		rl.DrawText("screen", 30, 10, 40, rl.White)
		rl.DrawText("fullscreen", 50, 60, 20, rl.White)
		rl.DrawText("window", 50, 90, 20, rl.White)
		rl.DrawText("difficulty", 30, 130, 40, rl.White)
		rl.DrawText("easy", 50, 180, 20, rl.White)
		rl.DrawText("average", 50, 210, 20, rl.White)
		rl.DrawText("difficult", 50, 240, 20, rl.White)
		boxx := int32(250)
		boxy := int32(60)
		for a := 0; a < 2; a++ {
			rl.DrawRectangle(boxx, boxy, 20, 20, rl.White)
			rl.DrawRectangle(boxx+2, boxy+2, 16, 16, rl.Black)
			boxy += 30
		}
		boxx = 250
		boxy = 180
		for a := 0; a < 3; a++ {
			rl.DrawRectangle(boxx, boxy, 20, 20, rl.White)
			rl.DrawRectangle(boxx+2, boxy+2, 16, 16, rl.Black)
			boxy += 30
		}
		boxx = 250
		boxy = 60
		if fullscreenOn {
			rl.DrawRectangle(boxx+2, boxy+2, 16, 16, rl.Orange)
		}
		boxy += 30
		if fullscreenOn == false {
			rl.DrawRectangle(boxx+2, boxy+2, 16, 16, rl.Orange)
		}
		boxy = 180
		if easyDiffOn {
			rl.DrawRectangle(boxx+2, boxy+2, 16, 16, rl.Orange)
		}
		boxy += 30
		if averageDiffOn {
			rl.DrawRectangle(boxx+2, boxy+2, 16, 16, rl.Orange)
		}
		boxy += 30
		if difficultDiffOn {
			rl.DrawRectangle(boxx+2, boxy+2, 16, 16, rl.Orange)
		}

		if selecty == 280 {
			rl.SetExitKey(rl.KeySpace) // key to end the game and close window
		} else {
			rl.SetExitKey(rl.KeyEnd) // key to end the game and close window
		}
		//	menucountTEXT := strconv.Itoa(launchermenucount)
		//	rl.DrawText(menucountTEXT, 240, 2, 20, rl.Green)
		//	selectyTEXT := strconv.Itoa(int(selecty))
		//	rl.DrawText(selectyTEXT, 240, 2, 20, rl.Green)
		rl.EndDrawing()
	} // end WindowShouldClose
	rl.CloseWindow()
}

// MARK: difficulty functions
func changediff() {

	if easyDiffOn {
		weatherOn = false
		cloudsactive = false
		fogactive = false
		rainactive = false
		snowactive = false
		frograinactive = false

		disastersOn = false
		meteorswitch = true

		magicOn = false
		teleportswitch = true
		distortionswitch = true

	} else if averageDiffOn {
		weatherOn = false
		cloudsactive = true
		fogactive = false
		rainactive = true
		snowactive = true
		frograinactive = false

		disastersOn = false
		meteorswitch = true
		tornadoswitch = true

		magicOn = false
		teleportswitch = true
		distortionswitch = true
		shockblockswitch = true

	} else if difficultDiffOn {
		weatherOn = true
		cloudsactive = true
		fogactive = true
		rainactive = true
		snowactive = true
		frograinactive = true

		disastersOn = true
		earthquakeswitch = true
		tornadoswitch = true
		meteorswitch = true

		magicOn = true
		teleportswitch = true
		rainoffireswitch = true
		shockblockswitch = true
		distortionswitch = true
	}

}

// MARK: tornado functions
func moveTornado() {

	if tornadoLR {

		tornadoHorizontal := tornadoLBlock / 94
		tornadoLVertical := tornadoLBlock - (94 * tornadoHorizontal)

		if tornadoLVertical > 4 {

			for a := 0; a < 5264; a++ {
				if tornadoMAP[a] == "tornado" {
					tornadoMAP[a] = ""
					tornadoMAP[a-1] = "tornado"
				}
			}
			tornadoLBlock--
			tornadoRBlock--

		} else {
			tornadoLR = false
		}
	} else {
		tornadoHorizontal := tornadoRBlock / 94
		tornadoRVertical := tornadoRBlock - (94 * tornadoHorizontal)

		if tornadoRVertical < 90 {

			for a := 5263; a > 0; a-- {
				if tornadoMAP[a] == "tornado" {
					tornadoMAP[a] = ""
					tornadoMAP[a+1] = "tornado"
				}

			}
			tornadoRBlock++
			tornadoLBlock++
		} else {
			tornadoLR = true
		}
	}
}

// MARK: meteor functions

func moveMeteorDown() {

	meteorHorizontal := meteorStartBlock / 94

	if meteorHorizontal < 42 {

		meteorMAP[meteorStartBlock] = ""
		meteorMAP[meteorStartBlock+1] = ""
		meteorMAP[meteorStartBlock+2] = ""
		meteorMAP[meteorStartBlock+3] = ""
		meteorMAP[meteorStartBlock+94] = ""
		meteorMAP[meteorStartBlock+95] = ""
		meteorMAP[meteorStartBlock+96] = ""
		meteorMAP[meteorStartBlock+97] = ""
		meteorMAP[meteorStartBlock+188] = ""
		meteorMAP[meteorStartBlock+189] = ""
		meteorMAP[meteorStartBlock+190] = ""
		meteorMAP[meteorStartBlock+191] = ""

		if meteorLR {
			meteorStartBlock += 94
			chooseLR := rolldice()
			if chooseLR == 6 {
				meteorStartBlock--
			}

			meteorMAP[meteorStartBlock] = "meteorTL"
			meteorMAP[meteorStartBlock+1] = "meteor"
			meteorMAP[meteorStartBlock+2] = "meteor"
			meteorMAP[meteorStartBlock+3] = "meteor"
			meteorMAP[meteorStartBlock+94] = "meteor"
			meteorMAP[meteorStartBlock+95] = "meteor"
			meteorMAP[meteorStartBlock+96] = "meteor"
			meteorMAP[meteorStartBlock+97] = "meteor"
			meteorMAP[meteorStartBlock+188] = "meteor"
			meteorMAP[meteorStartBlock+189] = "meteor"
			meteorMAP[meteorStartBlock+190] = "meteor"
			meteorMAP[meteorStartBlock+191] = "meteor"

			levelMAP[meteorStartBlock] = ""
			levelMAP[meteorStartBlock+1] = ""
			levelMAP[meteorStartBlock+2] = ""
			levelMAP[meteorStartBlock+3] = ""
			levelMAP[meteorStartBlock+94] = ""
			levelMAP[meteorStartBlock+95] = ""
			levelMAP[meteorStartBlock+96] = ""
			levelMAP[meteorStartBlock+97] = ""
			levelMAP[meteorStartBlock+188] = ""
			levelMAP[meteorStartBlock+189] = ""
			levelMAP[meteorStartBlock+190] = ""
			levelMAP[meteorStartBlock+191] = ""

			platformsMAP[meteorStartBlock] = ""
			platformsMAP[meteorStartBlock+1] = ""
			platformsMAP[meteorStartBlock+2] = ""
			platformsMAP[meteorStartBlock+3] = ""
			platformsMAP[meteorStartBlock+94] = ""
			platformsMAP[meteorStartBlock+95] = ""
			platformsMAP[meteorStartBlock+96] = ""
			platformsMAP[meteorStartBlock+97] = ""
			platformsMAP[meteorStartBlock+188] = ""
			platformsMAP[meteorStartBlock+189] = ""
			platformsMAP[meteorStartBlock+190] = ""
			platformsMAP[meteorStartBlock+191] = ""

			backgroundObjectsMAP[meteorStartBlock] = ""
			backgroundObjectsMAP[meteorStartBlock+1] = ""
			backgroundObjectsMAP[meteorStartBlock+2] = ""
			backgroundObjectsMAP[meteorStartBlock+3] = ""
			backgroundObjectsMAP[meteorStartBlock+94] = ""
			backgroundObjectsMAP[meteorStartBlock+95] = ""
			backgroundObjectsMAP[meteorStartBlock+96] = ""
			backgroundObjectsMAP[meteorStartBlock+97] = ""
			backgroundObjectsMAP[meteorStartBlock+188] = ""
			backgroundObjectsMAP[meteorStartBlock+189] = ""
			backgroundObjectsMAP[meteorStartBlock+190] = ""
			backgroundObjectsMAP[meteorStartBlock+191] = ""

		} else {
			meteorStartBlock += 94
			chooseLR := rolldice()
			if chooseLR == 6 {
				meteorStartBlock++
			}
			meteorMAP[meteorStartBlock] = "meteorTL"
			meteorMAP[meteorStartBlock+1] = "meteor"
			meteorMAP[meteorStartBlock+2] = "meteor"
			meteorMAP[meteorStartBlock+3] = "meteor"
			meteorMAP[meteorStartBlock+94] = "meteor"
			meteorMAP[meteorStartBlock+95] = "meteor"
			meteorMAP[meteorStartBlock+96] = "meteor"
			meteorMAP[meteorStartBlock+97] = "meteor"
			meteorMAP[meteorStartBlock+188] = "meteor"
			meteorMAP[meteorStartBlock+189] = "meteor"
			meteorMAP[meteorStartBlock+190] = "meteor"
			meteorMAP[meteorStartBlock+191] = "meteor"

			levelMAP[meteorStartBlock] = ""
			levelMAP[meteorStartBlock+1] = ""
			levelMAP[meteorStartBlock+2] = ""
			levelMAP[meteorStartBlock+3] = ""
			levelMAP[meteorStartBlock+94] = ""
			levelMAP[meteorStartBlock+95] = ""
			levelMAP[meteorStartBlock+96] = ""
			levelMAP[meteorStartBlock+97] = ""
			levelMAP[meteorStartBlock+188] = ""
			levelMAP[meteorStartBlock+189] = ""
			levelMAP[meteorStartBlock+190] = ""
			levelMAP[meteorStartBlock+191] = ""

			platformsMAP[meteorStartBlock] = ""
			platformsMAP[meteorStartBlock+1] = ""
			platformsMAP[meteorStartBlock+2] = ""
			platformsMAP[meteorStartBlock+3] = ""
			platformsMAP[meteorStartBlock+94] = ""
			platformsMAP[meteorStartBlock+95] = ""
			platformsMAP[meteorStartBlock+96] = ""
			platformsMAP[meteorStartBlock+97] = ""
			platformsMAP[meteorStartBlock+188] = ""
			platformsMAP[meteorStartBlock+189] = ""
			platformsMAP[meteorStartBlock+190] = ""
			platformsMAP[meteorStartBlock+191] = ""

			backgroundObjectsMAP[meteorStartBlock] = ""
			backgroundObjectsMAP[meteorStartBlock+1] = ""
			backgroundObjectsMAP[meteorStartBlock+2] = ""
			backgroundObjectsMAP[meteorStartBlock+3] = ""
			backgroundObjectsMAP[meteorStartBlock+94] = ""
			backgroundObjectsMAP[meteorStartBlock+95] = ""
			backgroundObjectsMAP[meteorStartBlock+96] = ""
			backgroundObjectsMAP[meteorStartBlock+97] = ""
			backgroundObjectsMAP[meteorStartBlock+188] = ""
			backgroundObjectsMAP[meteorStartBlock+189] = ""
			backgroundObjectsMAP[meteorStartBlock+190] = ""
			backgroundObjectsMAP[meteorStartBlock+191] = ""

		}

	} else {
		meteorMAP[meteorStartBlock] = "meteorTL"
		meteorMAP[meteorStartBlock+1] = "meteor"
		meteorMAP[meteorStartBlock+2] = "meteor"
		meteorMAP[meteorStartBlock+3] = "meteor"
		meteorMAP[meteorStartBlock+94] = "meteor"
		meteorMAP[meteorStartBlock+95] = "meteor"
		meteorMAP[meteorStartBlock+96] = "meteor"
		meteorMAP[meteorStartBlock+97] = "meteor"
		meteorMAP[meteorStartBlock+188] = "meteor"
		meteorMAP[meteorStartBlock+189] = "meteor"
		meteorMAP[meteorStartBlock+190] = "meteor"
		meteorMAP[meteorStartBlock+191] = "meteor"

		levelMAP[meteorStartBlock+1] = "floor"
		levelMAP[meteorStartBlock+2] = "floor"
		levelMAP[meteorStartBlock+94] = "floor"
		levelMAP[meteorStartBlock+95] = "floor"
		levelMAP[meteorStartBlock+96] = "floor"
		levelMAP[meteorStartBlock+97] = "floor"
		levelMAP[meteorStartBlock+188] = "floor"
		levelMAP[meteorStartBlock+189] = "floor"
		levelMAP[meteorStartBlock+190] = "floor"
		levelMAP[meteorStartBlock+191] = "floor"

		platformsEffectsMAP[meteorStartBlock+187] = "fire6"
		platformsEffectsMAP[meteorStartBlock+186] = "fire6"
		platformsEffectsMAP[meteorStartBlock+185] = "fire6"
		platformsEffectsMAP[meteorStartBlock+184] = "fire6"
		platformsEffectsMAP[meteorStartBlock+183] = "fire6"
		platformsEffectsMAP[meteorStartBlock+194] = "fire6"
		platformsEffectsMAP[meteorStartBlock+195] = "fire6"
		platformsEffectsMAP[meteorStartBlock+196] = "fire6"
		platformsEffectsMAP[meteorStartBlock+197] = "fire6"
		platformsEffectsMAP[meteorStartBlock+198] = "fire6"

	}

}

// MARK: extra shape functions

func clearAroundBlock(currentBlock int) {

	currentBlockHOLDER := currentBlock

	// current block
	levelMAP[currentBlock] = ""
	levelMAP[currentBlock+1] = ""
	levelMAP[currentBlock+94] = ""
	levelMAP[currentBlock+95] = ""

	platformsMAP[currentBlock] = ""
	platformsMAP[currentBlock+1] = ""
	platformsMAP[currentBlock+94] = ""
	platformsMAP[currentBlock+95] = ""

	// above block

	currentBlock = currentBlockHOLDER
	currentBlock -= 94 * 2

	levelMAP[currentBlock] = ""
	levelMAP[currentBlock+1] = ""
	levelMAP[currentBlock+94] = ""
	levelMAP[currentBlock+95] = ""

	platformsMAP[currentBlock] = ""
	platformsMAP[currentBlock+1] = ""
	platformsMAP[currentBlock+94] = ""
	platformsMAP[currentBlock+95] = ""

	// below block

	currentBlock = currentBlockHOLDER
	currentBlock += 94 * 2

	levelMAP[currentBlock] = ""
	levelMAP[currentBlock+1] = ""
	levelMAP[currentBlock+94] = ""
	levelMAP[currentBlock+95] = ""

	platformsMAP[currentBlock] = ""
	platformsMAP[currentBlock+1] = ""
	platformsMAP[currentBlock+94] = ""
	platformsMAP[currentBlock+95] = ""

	// left block
	currentBlock = currentBlockHOLDER
	currentBlock -= 2

	levelMAP[currentBlock] = ""
	levelMAP[currentBlock+1] = ""
	levelMAP[currentBlock+94] = ""
	levelMAP[currentBlock+95] = ""

	platformsMAP[currentBlock] = ""
	platformsMAP[currentBlock+1] = ""
	platformsMAP[currentBlock+94] = ""
	platformsMAP[currentBlock+95] = ""

	// above left block
	currentBlock -= (94 * 2)

	levelMAP[currentBlock] = ""
	levelMAP[currentBlock+1] = ""
	levelMAP[currentBlock+94] = ""
	levelMAP[currentBlock+95] = ""

	platformsMAP[currentBlock] = ""
	platformsMAP[currentBlock+1] = ""
	platformsMAP[currentBlock+94] = ""
	platformsMAP[currentBlock+95] = ""

	// below left block
	currentBlock += (94 * 4)

	levelMAP[currentBlock] = ""
	levelMAP[currentBlock+1] = ""
	levelMAP[currentBlock+94] = ""
	levelMAP[currentBlock+95] = ""

	platformsMAP[currentBlock] = ""
	platformsMAP[currentBlock+1] = ""
	platformsMAP[currentBlock+94] = ""
	platformsMAP[currentBlock+95] = ""

	// right block

	currentBlock = currentBlockHOLDER
	currentBlock += 2

	levelMAP[currentBlock] = ""
	levelMAP[currentBlock+1] = ""
	levelMAP[currentBlock+94] = ""
	levelMAP[currentBlock+95] = ""

	platformsMAP[currentBlock] = ""
	platformsMAP[currentBlock+1] = ""
	platformsMAP[currentBlock+94] = ""
	platformsMAP[currentBlock+95] = ""

	// above right block
	currentBlock -= (94 * 2)

	levelMAP[currentBlock] = ""
	levelMAP[currentBlock+1] = ""
	levelMAP[currentBlock+94] = ""
	levelMAP[currentBlock+95] = ""

	platformsMAP[currentBlock] = ""
	platformsMAP[currentBlock+1] = ""
	platformsMAP[currentBlock+94] = ""
	platformsMAP[currentBlock+95] = ""

	// below  right block
	currentBlock += (94 * 4)

	levelMAP[currentBlock] = ""
	levelMAP[currentBlock+1] = ""
	levelMAP[currentBlock+94] = ""
	levelMAP[currentBlock+95] = ""

	platformsMAP[currentBlock] = ""
	platformsMAP[currentBlock+1] = ""
	platformsMAP[currentBlock+94] = ""
	platformsMAP[currentBlock+95] = ""

}

func checkAroundPlatformBlock(currentBlock int) bool {

	blockIsNearby := false

	// current platform block
	block1 := levelMAP[currentBlock]
	block2 := levelMAP[currentBlock+1]
	block3 := levelMAP[currentBlock+94]
	block4 := levelMAP[currentBlock+95]

	// right platform block
	newCurrentBlock := currentBlock + 2

	block5 := levelMAP[newCurrentBlock]
	block6 := levelMAP[newCurrentBlock+1]
	block7 := levelMAP[newCurrentBlock+94]
	block8 := levelMAP[newCurrentBlock+95]

	// left platform block
	newCurrentBlock = currentBlock - 2

	block9 := levelMAP[newCurrentBlock]
	block10 := levelMAP[newCurrentBlock+1]
	block11 := levelMAP[newCurrentBlock+94]
	block12 := levelMAP[newCurrentBlock+95]

	// bottom platform block
	newCurrentBlock = currentBlock + (2 * 94)

	block13 := levelMAP[newCurrentBlock]
	block14 := levelMAP[newCurrentBlock+1]
	block15 := levelMAP[newCurrentBlock+94]
	block16 := levelMAP[newCurrentBlock+95]

	// bottom right platform block
	newCurrentBlock = currentBlock + ((2 * 94) + 1)

	block21 := levelMAP[newCurrentBlock]
	block22 := levelMAP[newCurrentBlock+1]
	block23 := levelMAP[newCurrentBlock+94]
	block24 := levelMAP[newCurrentBlock+95]

	// top platform block
	newCurrentBlock = currentBlock - (2 * 94)

	block17 := levelMAP[newCurrentBlock]
	block18 := levelMAP[newCurrentBlock+1]
	block19 := levelMAP[newCurrentBlock+94]
	block20 := levelMAP[newCurrentBlock+95]

	if block1 == "floor" || block2 == "floor" || block3 == "floor" || block4 == "floor" || block5 == "floor" || block6 == "floor" || block7 == "floor" || block8 == "floor" || block9 == "floor" || block10 == "floor" || block11 == "floor" || block12 == "floor" || block13 == "floor" || block14 == "floor" || block15 == "floor" || block16 == "floor" || block17 == "floor" || block18 == "floor" || block19 == "floor" || block20 == "floor" || block21 == "floor" || block22 == "floor" || block23 == "floor" || block24 == "floor" {

		blockIsNearby = true

	}

	return blockIsNearby

}

// MARK: noise lines functions

func cNOISELINES() {

	for a := 0; a < 12; a++ {
		noiseLinesMAP[a] = rInt32(0, int(screenW))
		noiseLinesLRMAP[a] = flipcoin()
		noiseLinesDistanceMAP[a] = rInt(100, 300)
	}
}

// MARK: pixel noise functions
func cPIXELNOISE() {

	for a := 0; a < 880; a++ {
		pixelNoiseMAP[a] = false
	}
	for a := 0; a < 880; a++ {
		placePixelNoise := rolldice() + rolldice()
		if placePixelNoise == 12 {
			pixelNoiseMAP[a] = true
		}

	}

}

// MARK: stars functions
func cSTARS() {
	for a := 0; a < 5264; a++ {

		placeStar := rolldice()

		if placeStar == 6 {

			chooseStarType := rInt(1, 4)

			if chooseStarType == 1 {
				starsMAP[a] = "star4"
			} else if chooseStarType == 2 {
				starsMAP[a] = "star3"
			} else if chooseStarType == 3 {
				starsMAP[a] = "star2"
			}

			chooseStarFade := rInt(0, 4)
			starsFadeMAP[a] = chooseStarFade

		}

	}
}

// MARK: falling block functions

func clearFallingBlock() {

	levelMAP[currentFallingBlock] = ""
	levelMAP[currentFallingBlock+1] = ""
	levelMAP[currentFallingBlock+94] = ""
	levelMAP[currentFallingBlock+95] = ""
	platformsMAP[currentFallingBlock] = ""
	platformsMAP[currentFallingBlock+1] = ""
	platformsMAP[currentFallingBlock+94] = ""
	platformsMAP[currentFallingBlock+95] = ""

}

func moveFallingBlockDown() {
	levelMAP[currentFallingBlock] = "fallingBlock"
	levelMAP[currentFallingBlock+1] = "fallingBlock"
	levelMAP[currentFallingBlock+94] = "fallingBlock"
	levelMAP[currentFallingBlock+95] = "fallingBlock"
	platformsMAP[currentFallingBlock] = fallingBlockTypeHOLDER
}

func cFALLINGBLOCKS() {

	fallingBlockStartBlock := rInt(194, 268)

	platformsMAP[fallingBlockStartBlock] = platformTypeNameTL
	levelMAP[fallingBlockStartBlock] = "fallingBlock"
	levelMAP[fallingBlockStartBlock+1] = "fallingBlock"
	levelMAP[fallingBlockStartBlock+94] = "fallingBlock"
	levelMAP[fallingBlockStartBlock+95] = "fallingBlock"

	fallingBlockStartBlock += 2

	platformsMAP[fallingBlockStartBlock] = platformTypeNameTL
	levelMAP[fallingBlockStartBlock] = "fallingBlock"
	levelMAP[fallingBlockStartBlock+1] = "fallingBlock"
	levelMAP[fallingBlockStartBlock+94] = "fallingBlock"
	levelMAP[fallingBlockStartBlock+95] = "fallingBlock"

	fallingBlockStartBlock += ((94 * 2) - 2)

	platformsMAP[fallingBlockStartBlock] = platformTypeNameTL
	levelMAP[fallingBlockStartBlock] = "fallingBlock"
	levelMAP[fallingBlockStartBlock+1] = "fallingBlock"
	levelMAP[fallingBlockStartBlock+94] = "fallingBlock"
	levelMAP[fallingBlockStartBlock+95] = "fallingBlock"

	fallingBlockStartBlock += 2

	platformsMAP[fallingBlockStartBlock] = platformTypeNameTL
	levelMAP[fallingBlockStartBlock] = "fallingBlock"
	levelMAP[fallingBlockStartBlock+1] = "fallingBlock"
	levelMAP[fallingBlockStartBlock+94] = "fallingBlock"
	levelMAP[fallingBlockStartBlock+95] = "fallingBlock"

}

// MARK: mini boss functions

func killMiniboss() {
	deadEnemiesMAP[miniBossCurrentBlock] = "deadEnemy1_4"
	minibosskills++
	if screenShake == false {
		screenShake = true
	}
	if enemyDieSound == false {
		enemyDieSound = true
	}
}

func clearMINIBOSS() {
	miniBossMAP[miniBossCurrentBlock] = ""
	miniBossMAP[miniBossCurrentBlock+94] = ""
	miniBossMAP[miniBossCurrentBlock+(94*2)] = ""
	miniBossMAP[miniBossCurrentBlock+(94*3)] = ""
	miniBossMAP[(miniBossCurrentBlock + 1)] = ""
	miniBossMAP[(miniBossCurrentBlock+1)+94] = ""
	miniBossMAP[(miniBossCurrentBlock+1)+(94*2)] = ""
	miniBossMAP[(miniBossCurrentBlock+1)+(94*3)] = ""
}
func upMINIBOSS() {
	miniBossMAP[miniBossCurrentBlock] = "miniBossTL"
	miniBossMAP[miniBossCurrentBlock+94] = "miniBoss"
	miniBossMAP[miniBossCurrentBlock+(94*2)] = "miniBoss"
	miniBossMAP[miniBossCurrentBlock+(94*3)] = "miniBoss"
	miniBossCurrentBlock++
	miniBossMAP[miniBossCurrentBlock] = "miniBoss"
	miniBossMAP[miniBossCurrentBlock+94] = "miniBoss"
	miniBossMAP[miniBossCurrentBlock+(94*2)] = "miniBoss"
	miniBossMAP[miniBossCurrentBlock+(94*3)] = "miniBoss"
}
func moveMiniBossDown() {

	if miniBossCurrentBlock < 3762 {
		clearMINIBOSS()
		miniBossCurrentBlock += 94 * 2
		upMINIBOSS()
	} else {
		miniBossFallOn = false
	}
}

func moveMiniBossUp() {
	if miniBossCurrentBlock > 564 {
		clearMINIBOSS()
		miniBossCurrentBlock -= 94 * 2
		upMINIBOSS()
	} else {
		miniBossJumpOn = false
	}
}

func moveMiniBossRight() {
	if miniBossType == 6 {
		clearMINIBOSS()
		if miniBossCurrentBlock < 1880 {
			miniBossCurrentBlock++
			upMINIBOSS()
		} else if miniBossCurrentBlock >= 1880 {
			miniBossCurrentBlock--
			upMINIBOSS()
		}
	} else {
		clearMINIBOSS()
		if miniBossCurrentBlock > 3852 {
			miniBossCurrentBlock = 3852
			miniBossRightLeft = false
			miniBossCurrentBlock--
			upMINIBOSS()
		} else {
			miniBossCurrentBlock++
			upMINIBOSS()
		}
	}
}

func moveMiniBossLeft() {

	if miniBossType == 6 {
		clearMINIBOSS()
		if miniBossCurrentBlock < 1880 {
			miniBossCurrentBlock++
			upMINIBOSS()
		} else if miniBossCurrentBlock >= 1880 {
			miniBossCurrentBlock--
			upMINIBOSS()
		}

	} else {
		clearMINIBOSS()
		if miniBossCurrentBlock < 3762 && miniBossCurrentBlock > 3572 {
			miniBossCurrentBlock = 3762
			miniBossRightLeft = true
			miniBossCurrentBlock++
			upMINIBOSS()
		} else {
			miniBossCurrentBlock--
			upMINIBOSS()
		}
	}
}

func cMINIBOSS() {

	for a := 0; a < 5264; a++ {
		miniBossMAP[a] = ""
	}

	miniBossType = rInt(1, 8)

	miniBossRightLeft = flipcoin()

	if miniBossType == 6 {
		miniBossMAP[756] = "miniBossTL"
		miniBossMAP[756+94] = "miniBoss"
		miniBossMAP[756+(94*2)] = "miniBoss"
		miniBossMAP[756+(94*3)] = "miniBoss"
		miniBossMAP[757] = "miniBoss"
		miniBossMAP[757+94] = "miniBoss"
		miniBossMAP[757+(94*2)] = "miniBoss"
		miniBossMAP[757+(94*3)] = "miniBoss"
	} else {
		if miniBossRightLeft {
			miniBossMAP[1506] = "miniBossTL"
			miniBossMAP[1506+94] = "miniBoss"
			miniBossMAP[1506+(94*2)] = "miniBoss"
			miniBossMAP[1506+(94*3)] = "miniBoss"
			miniBossMAP[1507] = "miniBoss"
			miniBossMAP[1507+94] = "miniBoss"
			miniBossMAP[1507+(94*2)] = "miniBoss"
			miniBossMAP[1507+(94*3)] = "miniBoss"
		} else {
			miniBossMAP[1594] = "miniBossTL"
			miniBossMAP[1594+94] = "miniBoss"
			miniBossMAP[1594+(94*2)] = "miniBoss"
			miniBossMAP[1594+(94*3)] = "miniBoss"
			miniBossMAP[1595] = "miniBoss"
			miniBossMAP[1595+94] = "miniBoss"
			miniBossMAP[1595+(94*2)] = "miniBoss"
			miniBossMAP[1595+(94*3)] = "miniBoss"
		}
	}
}

// MARK: background image functions

func changeBackgroundImage() {

	if standardBackOn && altColorsOn && grayscaleOn && pencilOn && inkOn && chalkOn {
		randomBackOn = true
	}

	if randomBackOn == false {
		choosebackcount = 5

		if standardBackOn == false {
			choosebackcount--
		}
		if altColorsOn == false {
			choosebackcount--
		}
		if grayscaleOn == false {
			choosebackcount--
		}
		if pencilOn == false {
			choosebackcount--
		}
		if inkOn == false {
			choosebackcount--
		}
		if chalkOn == false {
			choosebackcount--
		}
		if choosebackcount <= 0 {
			randomBackOn = true
		}
	}

	if randomBackOn {

		chooseBackground := rInt(1, 55)

		switch chooseBackground {
		case 1:
			backgroundTexture = rl.LoadTexture("back1.png")
		case 2:
			backgroundTexture = rl.LoadTexture("back1_2.png")
		case 3:
			backgroundTexture = rl.LoadTexture("back1_3.png")
		case 4:
			backgroundTexture = rl.LoadTexture("back1_4.png")
		case 5:
			backgroundTexture = rl.LoadTexture("back2.png")
		case 6:
			backgroundTexture = rl.LoadTexture("back2_2.png")
		case 7:
			backgroundTexture = rl.LoadTexture("back2_3.png")
		case 8:
			backgroundTexture = rl.LoadTexture("back2_4.png")
		case 9:
			backgroundTexture = rl.LoadTexture("back3.png")
		case 10:
			backgroundTexture = rl.LoadTexture("back3_2.png")
		case 11:
			backgroundTexture = rl.LoadTexture("back3_3.png")
		case 12:
			backgroundTexture = rl.LoadTexture("back3_4.png")
		case 13:
			backgroundTexture = rl.LoadTexture("back4.png")
		case 14:
			backgroundTexture = rl.LoadTexture("back4_2.png")
		case 15:
			backgroundTexture = rl.LoadTexture("back4_3.png")
		case 16:
			backgroundTexture = rl.LoadTexture("back4_4.png")
		case 17:
			backgroundTexture = rl.LoadTexture("back5.png")
		case 18:
			backgroundTexture = rl.LoadTexture("back5_2.png")
		case 19:
			backgroundTexture = rl.LoadTexture("back5_3.png")
		case 20:
			backgroundTexture = rl.LoadTexture("back5_4.png")
		case 21:
			backgroundTexture = rl.LoadTexture("back6.png")
		case 22:
			backgroundTexture = rl.LoadTexture("back6_2.png")
		case 23:
			backgroundTexture = rl.LoadTexture("back6_3.png")
		case 24:
			backgroundTexture = rl.LoadTexture("back6_4.png")
		case 25:
			backgroundTexture = rl.LoadTexture("back7.png")
		case 26:
			backgroundTexture = rl.LoadTexture("back7_2.png")
		case 27:
			backgroundTexture = rl.LoadTexture("back7_3.png")
		case 28:
			backgroundTexture = rl.LoadTexture("back7_4.png")
		case 29:
			backgroundTexture = rl.LoadTexture("back8.png")
		case 30:
			backgroundTexture = rl.LoadTexture("back8_2.png")
		case 31:
			backgroundTexture = rl.LoadTexture("back8_3.png")
		case 32:
			backgroundTexture = rl.LoadTexture("back8_4.png")
		case 33:
			backgroundTexture = rl.LoadTexture("back9.png")
		case 34:
			backgroundTexture = rl.LoadTexture("back9_2.png")
		case 35:
			backgroundTexture = rl.LoadTexture("back9_3.png")
		case 36:
			backgroundTexture = rl.LoadTexture("back9_4.png")
		case 37:
			backgroundTexture = rl.LoadTexture("back1_5.png")
		case 38:
			backgroundTexture = rl.LoadTexture("back1_6.png")
		case 39:
			backgroundTexture = rl.LoadTexture("back2_5.png")
		case 40:
			backgroundTexture = rl.LoadTexture("back2_6.png")
		case 41:
			backgroundTexture = rl.LoadTexture("back3_5.png")
		case 42:
			backgroundTexture = rl.LoadTexture("back3_6.png")
		case 43:
			backgroundTexture = rl.LoadTexture("back4_5.png")
		case 44:
			backgroundTexture = rl.LoadTexture("back4_6.png")
		case 45:
			backgroundTexture = rl.LoadTexture("back5_5.png")
		case 46:
			backgroundTexture = rl.LoadTexture("back5_6.png")
		case 47:
			backgroundTexture = rl.LoadTexture("back6_5.png")
		case 48:
			backgroundTexture = rl.LoadTexture("back6_6.png")
		case 49:
			backgroundTexture = rl.LoadTexture("back7_5.png")
		case 50:
			backgroundTexture = rl.LoadTexture("back7_6.png")
		case 51:
			backgroundTexture = rl.LoadTexture("back8_5.png")
		case 52:
			backgroundTexture = rl.LoadTexture("back8_6.png")
		case 53:
			backgroundTexture = rl.LoadTexture("back9_5.png")
		case 54:
			backgroundTexture = rl.LoadTexture("back9_6.png")
		}
	} else {
		backgroundSelect = rInt(1, 10)
		if choosebackcount == 1 {
			switch backgroundSelect {
			case 1:
				if grayscaleOn {
					backgroundTexture = rl.LoadTexture("back1_2.png")
				} else if altColorsOn {
					backgroundTexture = rl.LoadTexture("back1_3.png")
				} else if pencilOn {
					backgroundTexture = rl.LoadTexture("back1_4.png")
				} else if inkOn {
					backgroundTexture = rl.LoadTexture("back1_5.png")
				} else if chalkOn {
					backgroundTexture = rl.LoadTexture("back1_6.png")
				} else {
					backgroundTexture = rl.LoadTexture("back1.png")
					standardBackOn = true
				}
			case 2:
				if grayscaleOn {
					backgroundTexture = rl.LoadTexture("back2_2.png")
				} else if altColorsOn {
					backgroundTexture = rl.LoadTexture("back2_3.png")
				} else if pencilOn {
					backgroundTexture = rl.LoadTexture("back2_4.png")
				} else if inkOn {
					backgroundTexture = rl.LoadTexture("back2_5.png")
				} else if chalkOn {
					backgroundTexture = rl.LoadTexture("back2_6.png")
				} else {
					backgroundTexture = rl.LoadTexture("back2.png")
					standardBackOn = true
				}
			case 3:
				if grayscaleOn {
					backgroundTexture = rl.LoadTexture("back3_2.png")
				} else if altColorsOn {
					backgroundTexture = rl.LoadTexture("back3_3.png")
				} else if pencilOn {
					backgroundTexture = rl.LoadTexture("back3_4.png")
				} else if inkOn {
					backgroundTexture = rl.LoadTexture("back3_5.png")
				} else if chalkOn {
					backgroundTexture = rl.LoadTexture("back3_6.png")
				} else {
					backgroundTexture = rl.LoadTexture("back3.png")
					standardBackOn = true
				}
			case 4:
				if grayscaleOn {
					backgroundTexture = rl.LoadTexture("back4_2.png")
				} else if altColorsOn {
					backgroundTexture = rl.LoadTexture("back4_3.png")
				} else if pencilOn {
					backgroundTexture = rl.LoadTexture("back4_4.png")
				} else if inkOn {
					backgroundTexture = rl.LoadTexture("back4_5.png")
				} else if chalkOn {
					backgroundTexture = rl.LoadTexture("back4_6.png")
				} else {
					backgroundTexture = rl.LoadTexture("back4.png")
					standardBackOn = true
				}
			case 5:
				if grayscaleOn {
					backgroundTexture = rl.LoadTexture("back5_2.png")
				} else if altColorsOn {
					backgroundTexture = rl.LoadTexture("back5_3.png")
				} else if pencilOn {
					backgroundTexture = rl.LoadTexture("back5_4.png")
				} else if inkOn {
					backgroundTexture = rl.LoadTexture("back5_5.png")
				} else if chalkOn {
					backgroundTexture = rl.LoadTexture("back5_6.png")
				} else {
					backgroundTexture = rl.LoadTexture("back5.png")
					standardBackOn = true
				}
			case 6:
				if grayscaleOn {
					backgroundTexture = rl.LoadTexture("back6_2.png")
				} else if altColorsOn {
					backgroundTexture = rl.LoadTexture("back6_3.png")
				} else if pencilOn {
					backgroundTexture = rl.LoadTexture("back6_4.png")
				} else if inkOn {
					backgroundTexture = rl.LoadTexture("back6_5.png")
				} else if chalkOn {
					backgroundTexture = rl.LoadTexture("back6_6.png")
				} else {
					backgroundTexture = rl.LoadTexture("back6.png")
					standardBackOn = true
				}
			case 7:
				if grayscaleOn {
					backgroundTexture = rl.LoadTexture("back7_2.png")
				} else if altColorsOn {
					backgroundTexture = rl.LoadTexture("back7_3.png")
				} else if pencilOn {
					backgroundTexture = rl.LoadTexture("back7_4.png")
				} else if inkOn {
					backgroundTexture = rl.LoadTexture("back7_5.png")
				} else if chalkOn {
					backgroundTexture = rl.LoadTexture("back7_6.png")
				} else {
					backgroundTexture = rl.LoadTexture("back7.png")
					standardBackOn = true
				}
			case 8:
				if grayscaleOn {
					backgroundTexture = rl.LoadTexture("back8_2.png")
				} else if altColorsOn {
					backgroundTexture = rl.LoadTexture("back8_3.png")
				} else if pencilOn {
					backgroundTexture = rl.LoadTexture("back8_4.png")
				} else if inkOn {
					backgroundTexture = rl.LoadTexture("back8_5.png")
				} else if chalkOn {
					backgroundTexture = rl.LoadTexture("back8_6.png")
				} else {
					backgroundTexture = rl.LoadTexture("back8.png")
					standardBackOn = true
				}
			case 9:
				if grayscaleOn {
					backgroundTexture = rl.LoadTexture("back9_2.png")
				} else if altColorsOn {
					backgroundTexture = rl.LoadTexture("back9_3.png")
				} else if pencilOn {
					backgroundTexture = rl.LoadTexture("back9_4.png")
				} else if inkOn {
					backgroundTexture = rl.LoadTexture("back9_5.png")
				} else if chalkOn {
					backgroundTexture = rl.LoadTexture("back9_6.png")
				} else {
					backgroundTexture = rl.LoadTexture("back9.png")
					standardBackOn = true
				}
			}
		} else if choosebackcount > 1 {

			backselected := false
			for {
				choooseback := rolldice()
				if choooseback == 1 {
					switch backgroundSelect {
					case 1:
						if grayscaleOn {
							backgroundTexture = rl.LoadTexture("back1_2.png")
							backselected = true
						}
					case 2:
						if grayscaleOn {
							backgroundTexture = rl.LoadTexture("back2_2.png")
							backselected = true
						}
					case 3:
						if grayscaleOn {
							backgroundTexture = rl.LoadTexture("back3_2.png")
							backselected = true
						}
					case 4:
						if grayscaleOn {
							backgroundTexture = rl.LoadTexture("back4_2.png")
							backselected = true
						}
					case 5:
						if grayscaleOn {
							backgroundTexture = rl.LoadTexture("back5_2.png")
							backselected = true
						}
					case 6:
						if grayscaleOn {
							backgroundTexture = rl.LoadTexture("back6_2.png")
							backselected = true
						}
					case 7:
						if grayscaleOn {
							backgroundTexture = rl.LoadTexture("back7_2.png")
							backselected = true
						}
					case 8:
						if grayscaleOn {
							backgroundTexture = rl.LoadTexture("back8_2.png")
							backselected = true
						}
					case 9:
						if grayscaleOn {
							backselected = true
						}
					}
				} else if choooseback == 2 {
					switch backgroundSelect {
					case 1:
						if altColorsOn {
							backgroundTexture = rl.LoadTexture("back1_3.png")
							backselected = true
						}
					case 2:
						if altColorsOn {
							backgroundTexture = rl.LoadTexture("back2_3.png")
							backselected = true
						}
					case 3:
						if altColorsOn {
							backgroundTexture = rl.LoadTexture("back3_3.png")
							backselected = true
						}
					case 4:
						if altColorsOn {
							backgroundTexture = rl.LoadTexture("back4_3.png")
							backselected = true
						}
					case 5:
						if altColorsOn {
							backgroundTexture = rl.LoadTexture("back5_3.png")
							backselected = true
						}
					case 6:
						if altColorsOn {
							backgroundTexture = rl.LoadTexture("back6_3.png")
							backselected = true
						}
					case 7:
						if altColorsOn {
							backgroundTexture = rl.LoadTexture("back7_3.png")
							backselected = true
						}
					case 8:
						if altColorsOn {
							backgroundTexture = rl.LoadTexture("back8_3.png")
							backselected = true
						}
					case 9:
						if altColorsOn {
							backgroundTexture = rl.LoadTexture("back9_3.png")
							backselected = true
						}
					}
				} else if choooseback == 3 {
					switch backgroundSelect {
					case 1:
						if pencilOn {
							backgroundTexture = rl.LoadTexture("back1_4.png")
							backselected = true
						}
					case 2:
						if pencilOn {
							backgroundTexture = rl.LoadTexture("back2_4.png")
							backselected = true
						}
					case 3:
						if pencilOn {
							backgroundTexture = rl.LoadTexture("back3_4.png")
							backselected = true
						}
					case 4:
						if pencilOn {
							backgroundTexture = rl.LoadTexture("back4_4.png")
							backselected = true
						}
					case 5:
						if pencilOn {
							backgroundTexture = rl.LoadTexture("back5_4.png")
							backselected = true
						}
					case 6:
						if pencilOn {
							backgroundTexture = rl.LoadTexture("back6_4.png")
							backselected = true
						}
					case 7:
						if pencilOn {
							backgroundTexture = rl.LoadTexture("back7_4.png")
							backselected = true
						}
					case 8:
						if pencilOn {
							backgroundTexture = rl.LoadTexture("back8_4.png")
							backselected = true
						}
					case 9:
						if pencilOn {
							backgroundTexture = rl.LoadTexture("back9_4.png")
							backselected = true
						}
					}
				} else if choooseback == 4 {
					switch backgroundSelect {
					case 1:
						if inkOn {
							backgroundTexture = rl.LoadTexture("back1_5.png")
							backselected = true
						}
					case 2:
						if inkOn {
							backgroundTexture = rl.LoadTexture("back2_5.png")
							backselected = true
						}
					case 3:
						if inkOn {
							backgroundTexture = rl.LoadTexture("back3_5.png")
							backselected = true
						}
					case 4:
						if inkOn {
							backgroundTexture = rl.LoadTexture("back4_5.png")
							backselected = true
						}
					case 5:
						if inkOn {
							backgroundTexture = rl.LoadTexture("back5_5.png")
							backselected = true
						}
					case 6:
						if inkOn {
							backgroundTexture = rl.LoadTexture("back6_5.png")
							backselected = true
						}
					case 7:
						if inkOn {
							backgroundTexture = rl.LoadTexture("back7_5.png")
							backselected = true
						}
					case 8:
						if inkOn {
							backgroundTexture = rl.LoadTexture("back8_5.png")
							backselected = true
						}
					case 9:
						if inkOn {
							backgroundTexture = rl.LoadTexture("back9_5.png")
							backselected = true
						}
					}
				} else if choooseback == 5 {
					switch backgroundSelect {
					case 1:
						if chalkOn {
							backgroundTexture = rl.LoadTexture("back1_6.png")
							backselected = true
						}
					case 2:
						if chalkOn {
							backgroundTexture = rl.LoadTexture("back2_6.png")
							backselected = true
						}
					case 3:
						if chalkOn {
							backgroundTexture = rl.LoadTexture("back3_6.png")
							backselected = true
						}
					case 4:
						if chalkOn {
							backgroundTexture = rl.LoadTexture("back4_6.png")
							backselected = true
						}
					case 5:
						if chalkOn {
							backgroundTexture = rl.LoadTexture("back5_6.png")
							backselected = true
						}
					case 6:
						if chalkOn {
							backgroundTexture = rl.LoadTexture("back6_6.png")
							backselected = true
						}
					case 7:
						if chalkOn {
							backgroundTexture = rl.LoadTexture("back7_6.png")
							backselected = true
						}
					case 8:
						if chalkOn {
							backgroundTexture = rl.LoadTexture("back8_6.png")
							backselected = true
						}
					case 9:
						if chalkOn {
							backgroundTexture = rl.LoadTexture("back9_6.png")
							backselected = true
						}
					}
				} else if choooseback == 6 {

					switch backgroundSelect {
					case 1:
						if standardBackOn {
							backgroundTexture = rl.LoadTexture("back1.png")
							backselected = true
						}
					case 2:
						if standardBackOn {
							backgroundTexture = rl.LoadTexture("back2.png")
							backselected = true
						}
					case 3:
						if standardBackOn {
							backgroundTexture = rl.LoadTexture("back3.png")
							backselected = true
						}
					case 4:
						if standardBackOn {
							backgroundTexture = rl.LoadTexture("back4.png")
							backselected = true
						}
					case 5:
						if standardBackOn {
							backgroundTexture = rl.LoadTexture("back5.png")
							backselected = true
						}
					case 6:
						if standardBackOn {
							backgroundTexture = rl.LoadTexture("back6.png")
							backselected = true
						}
					case 7:
						if standardBackOn {
							backgroundTexture = rl.LoadTexture("back7.png")
							backselected = true
						}
					case 8:
						if standardBackOn {
							backgroundTexture = rl.LoadTexture("back8.png")
							backselected = true
						}
					case 9:
						if standardBackOn {
							backgroundTexture = rl.LoadTexture("back9.png")
							backselected = true
						}
					}

				}
				if backselected == true {
					break
				}
			}

		}
	}
}

// MARK: dead enemies functions
func deadEnemyCircles() {
	for a := 0; a < 10; a++ {
		deadEnemiesCirclesRadius[a] = rFloat32(30, 100)
		deadEnemyCircleX := rFloat32(20, 1340)
		deadEnemyCircleY := rFloat32(20, 740)
		deadEnemiesCirclesV2[a] = rl.NewVector2(deadEnemyCircleX, deadEnemyCircleY)
	}
}

// MARK: intro functions
func randomColors() {
	colorsMAP[0] = rl.Red
	colorsMAP[1] = rl.Black
	colorsMAP[2] = rl.Gold
	colorsMAP[3] = rl.Red
	colorsMAP[4] = rl.Black
	colorsMAP[5] = rl.Black
	colorsMAP[6] = rl.Gold
	colorsMAP[7] = rl.Red
	colorsMAP[8] = rl.Red
	colorsMAP[9] = rl.Red

	rand.Shuffle(len(colorsMAP), func(i, j int) { colorsMAP[i], colorsMAP[j] = colorsMAP[j], colorsMAP[i] })

}

// MARK: choose effects
func chooseeffects() {

	earthquakesOn = false
	tornadoOn = false
	meteorOn = false
	teleportOn = false
	rainoffireOn = false
	shockBlocksOn = false
	disastersOn = false

	if disastersOn {
		if rolldice() == 6 {
			choose := rInt(1, 4)
			if choose == 1 {
				earthquakesOn = true
				tornadoOn = false
				meteorOn = false
			} else if choose == 2 {
				earthquakesOn = false
				tornadoOn = true
				meteorOn = false
			} else if choose == 3 {
				earthquakesOn = false
				tornadoOn = false
				meteorOn = true
			}
		}

	}
	if magicOn {
		if rolldice() == 6 {
			choose := rInt(1, 5)
			if choose == 1 {
				teleportOn = true
				rainoffireOn = false
				shockBlocksOn = false
				platformDistortionOn = false
			} else if choose == 2 {
				teleportOn = false
				rainoffireOn = true
				shockBlocksOn = false
				platformDistortionOn = false
			} else if choose == 3 {
				teleportOn = false
				rainoffireOn = false
				shockBlocksOn = true
				platformDistortionOn = false
			} else if choose == 4 {
				teleportOn = false
				rainoffireOn = false
				shockBlocksOn = false
				platformDistortionOn = true
			}
		}
		if difficultDiffOn {
			if rolldice() == 6 {
				choose := rInt(1, 5)
				if choose == 1 {
					if teleportOn == false {
						teleportOn = true
					} else {
						rainoffireOn = true
					}
				} else if choose == 2 {
					if rainoffireOn == false {
						rainoffireOn = true
					} else {
						shockBlocksOn = true
					}
				} else if choose == 3 {
					if shockBlocksOn == false {
						shockBlocksOn = true
					} else {
						platformDistortionOn = true
					}
				} else if choose == 4 {
					if platformDistortionOn == false {
						platformDistortionOn = true
					} else {
						teleportOn = true
					}
				}
			}
		}

	}

}

// MARK: weather functions
func chooseweather() {
	cloudsOn = false
	fogOn = false
	rainOn = false
	snowOn = false
	frogRainOn = false
	rainFrogOn = false

	choose := rolldice()

	switch choose {

	case 1:
		if cloudsactive {
			cloudsOn = true
		}
	case 2:
		if fogactive {
			fogOn = true
		}
	case 3:
		if rainactive {
			rainOn = true
		}
	case 4:
		if snowactive {
			snowOn = true
		}
	case 5:
		if frograinactive {
			frogRainOn = true
			rainFrogOn = true
		}
	case 6:
		choose2 := rolldice()
		switch choose2 {
		case 1:
			if cloudsactive {
				cloudsOn = true
			}
		case 2:
			if fogactive {
				fogOn = true
			}
		case 3:
			if rainactive {
				rainOn = true
			}
		case 4:
			if snowactive {
				snowOn = true
			}
		case 5:
			if frograinactive {
				frogRainOn = true
			}
		}
	}

	if difficultDiffOn {
		another := flipcoin()

		if another {

			choose = rolldice()

			switch choose {

			case 1:
				if cloudsactive {
					cloudsOn = true
				}
			case 2:
				if fogactive {
					fogOn = true
				}
			case 3:
				if rainactive {
					rainOn = true
				}
			case 4:
				if snowactive {
					snowOn = true
				}
			case 5:
				if frograinactive {
					frogRainOn = true
					rainFrogOn = true
				}
			case 6:
				choose2 := rolldice()
				switch choose2 {
				case 1:
					if cloudsactive {
						cloudsOn = true
					}
				case 2:
					if fogactive {
						fogOn = true
					}
				case 3:
					if rainactive {
						rainOn = true
					}
				case 4:
					if snowactive {
						snowOn = true
					}
				case 5:
					if frograinactive {
						frogRainOn = true
					}
				}
			}
		}

	}

	if rainOn && snowOn {
		rainOn = false
	}

	if totalLevels < 3 && fogOn {
		fogOn = false
	}

}
func upRAIN() {
	for a := 0; a < 188; a++ {
		placeRain := rolldice()
		if placeRain == 6 {
			rainMAP[a] = "rain"
		}
	}
}
func cRAIN() {
	for a := 0; a < 4228; a++ {
		placeRain := rolldice()
		if placeRain == 6 {
			rainMAP[a] = "rain"
		}
	}
}

func cCLOUDS() {

	cloudNumber := rInt(300, 400)

	cloudSpace := 9400 / cloudNumber
	cloudStartPoint := (rInt(5, 10) * 9400) + (rInt(5, 10))
	cloudStartPointHOLDER := cloudStartPoint

	for a := 0; a < cloudNumber; a++ {

		cloudMaxWidth := rInt(10, cloudSpace-4)
		cloudMaxHeight := rInt(10, 20)
		cloudHalfHeight := cloudMaxHeight / 2

		cloudStartPoint += (cloudMaxHeight / 2) * 9400
		cloudStartPointHOLDER2 := cloudStartPoint

		for b := 0; b < cloudMaxWidth; b++ {
			cloudsMAP[cloudStartPoint+b] = "cloud"
		}

		newCloudWidth := cloudMaxWidth

		for b := 0; b < cloudHalfHeight; b++ {
			cloudStartPoint -= 9400
			changeCloudWidth := rInt(1, 3)
			newCloudWidth = newCloudWidth - (changeCloudWidth * 2)
			cloudStartPoint += changeCloudWidth
			for c := 0; c < newCloudWidth; c++ {
				cloudsMAP[cloudStartPoint+c] = "cloud"
			}

			if newCloudWidth <= 4 {
				break
			}
		}

		cloudStartPoint = cloudStartPointHOLDER2
		newCloudWidth = cloudMaxWidth

		for b := 0; b < cloudHalfHeight; b++ {
			cloudStartPoint += 9400
			changeCloudWidth := rInt(1, 3)
			newCloudWidth = newCloudWidth - (changeCloudWidth * 2)
			cloudStartPoint += changeCloudWidth
			for c := 0; c < newCloudWidth; c++ {
				cloudsMAP[cloudStartPoint+c] = "cloud"
			}

			if newCloudWidth <= 4 {
				break
			}
		}

		cloudStartPoint = cloudStartPointHOLDER + cloudMaxWidth + rInt(4, 12)
		cloudStartPointHorizontal := cloudStartPoint / 9400

		if cloudStartPointHorizontal > 5 && cloudStartPointHorizontal < 30 {

			cloudStartPoint += (rInt(-3, 4) * 9400)

		}

		cloudStartPointHOLDER = cloudStartPoint
	}

}

// MARK: weapon functions
func clearWeapon() {
	weaponsMAP[currentDropWeaponBlock] = ""
	weaponsMAP[currentDropWeaponBlock+1] = ""
}
func moveWeaponDown() {
	currentDropWeaponBlock += 94
	weaponsMAP[currentDropWeaponBlock] = dropWeaponHOLDER1
	weaponsMAP[currentDropWeaponBlock+1] = dropWeaponHOLDER2
}

// MARK: object functions
func clearCoin() {
	objectsMAP[coinCurrentBlock] = ""
	objectsMAP[coinCurrentBlock+1] = ""
	objectsMAP[coinCurrentBlock+94] = ""
	objectsMAP[coinCurrentBlock+95] = ""
}
func moveCoinDown() {
	coinCurrentBlock += 94
	objectsMAP[coinCurrentBlock] = "coinTL"
	objectsMAP[coinCurrentBlock+1] = "coin"
	objectsMAP[coinCurrentBlock+94] = "coin"
	objectsMAP[coinCurrentBlock+95] = "coin"
}

// MARK: power up functions
func clearPowerUp() {

	powerUpTypeTL = powerUpsMAP[currentPowerUpBlock]
	powerUpType = powerUpsMAP[currentPowerUpBlock+1]

	powerUpsMAP[currentPowerUpBlock] = ""
	powerUpsMAP[currentPowerUpBlock+1] = ""
	powerUpsMAP[currentPowerUpBlock+94] = ""
	powerUpsMAP[currentPowerUpBlock+95] = ""

}

func movePowerUpDown() {

	newPowerUpBlockPosition := currentPowerUpBlock + 94

	powerUpsMAP[newPowerUpBlockPosition] = powerUpTypeTL
	powerUpsMAP[newPowerUpBlockPosition+1] = powerUpType
	powerUpsMAP[newPowerUpBlockPosition+94] = powerUpType
	powerUpsMAP[newPowerUpBlockPosition+95] = powerUpType

}

func cPOWERUPS() {

	powerUpBlock := rInt(10, 31)
	powerUpBlock = powerUpBlock * 94

	powerUpVehicleMAP[powerUpBlock] = "powerUpVehicleTL"
	powerUpVehicleMAP[powerUpBlock+1] = "powerUpVehicle"
	powerUpVehicleMAP[powerUpBlock+94] = "powerUpVehicle"
	powerUpVehicleMAP[powerUpBlock+95] = "powerUpVehicle"

}

func movePowerUpVehicleRight() {

	powerUpVehicleHorizontal := powerUpVehicleCurrentBlock / 94
	powerUpVehicleVertical := powerUpVehicleCurrentBlock - (powerUpVehicleHorizontal * 94)

	if powerUpVehicleVertical < 91 {

		powerUpVehicleCurrentBlock++

		powerUpVehicleMAP[powerUpVehicleCurrentBlock] = "powerUpVehicleTL"
		powerUpVehicleMAP[powerUpVehicleCurrentBlock+1] = "powerUpVehicle"
		powerUpVehicleMAP[powerUpVehicleCurrentBlock+94] = "powerUpVehicle"
		powerUpVehicleMAP[powerUpVehicleCurrentBlock+95] = "powerUpVehicle"

		if powerUpVehicleHorizontal > 6 && powerUpVehicleHorizontal < 35 {

			upDown := flipcoin()

			if upDown {
				clearPowerUpVehicle()
				powerUpVehicleCurrentBlock -= 94
				powerUpVehicleMAP[powerUpVehicleCurrentBlock] = "powerUpVehicleTL"
				powerUpVehicleMAP[powerUpVehicleCurrentBlock+1] = "powerUpVehicle"
				powerUpVehicleMAP[powerUpVehicleCurrentBlock+94] = "powerUpVehicle"
				powerUpVehicleMAP[powerUpVehicleCurrentBlock+95] = "powerUpVehicle"
			} else {
				clearPowerUpVehicle()
				powerUpVehicleCurrentBlock += 94
				powerUpVehicleMAP[powerUpVehicleCurrentBlock] = "powerUpVehicleTL"
				powerUpVehicleMAP[powerUpVehicleCurrentBlock+1] = "powerUpVehicle"
				powerUpVehicleMAP[powerUpVehicleCurrentBlock+94] = "powerUpVehicle"
				powerUpVehicleMAP[powerUpVehicleCurrentBlock+95] = "powerUpVehicle"
			}

		} else if powerUpVehicleHorizontal <= 6 {
			clearPowerUpVehicle()
			powerUpVehicleCurrentBlock += 94
			powerUpVehicleMAP[powerUpVehicleCurrentBlock] = "powerUpVehicleTL"
			powerUpVehicleMAP[powerUpVehicleCurrentBlock+1] = "powerUpVehicle"
			powerUpVehicleMAP[powerUpVehicleCurrentBlock+94] = "powerUpVehicle"
			powerUpVehicleMAP[powerUpVehicleCurrentBlock+95] = "powerUpVehicle"

		} else if powerUpVehicleHorizontal >= 35 {
			clearPowerUpVehicle()
			powerUpVehicleCurrentBlock -= 94
			powerUpVehicleMAP[powerUpVehicleCurrentBlock] = "powerUpVehicleTL"
			powerUpVehicleMAP[powerUpVehicleCurrentBlock+1] = "powerUpVehicle"
			powerUpVehicleMAP[powerUpVehicleCurrentBlock+94] = "powerUpVehicle"
			powerUpVehicleMAP[powerUpVehicleCurrentBlock+95] = "powerUpVehicle"

		}

	} else {
		clearPowerUpVehicle()
	}

	if powerUpDropped == false && powerUpHasDropped == false {

		dropPowerUp := rInt(0, 1000)

		if dropPowerUp > 980 {
			powerUpDropped = true
		}

	}
	if powerUpDropped == false && powerUpHasDropped == false && powerUpVehicleVertical > 80 {
		powerUpDropped = true
	}

	// MARK: create powerup
	if powerUpDropped {

		powerUpType := rInt(1, 11)

		powerUpName := "powerUp" + strconv.Itoa(powerUpType)
		powerUpNameTL := powerUpName + "TL"

		powerUpPosition := powerUpVehicleCurrentBlock + (94 * 3)

		powerUpsMAP[powerUpPosition] = powerUpNameTL
		powerUpsMAP[powerUpPosition+1] = powerUpName
		powerUpsMAP[powerUpPosition+94] = powerUpName
		powerUpsMAP[powerUpPosition+95] = powerUpName

		powerUpDropped = false
		powerUpHasDropped = true
	}

}

func clearPowerUpVehicle() {
	powerUpVehicleMAP[powerUpVehicleCurrentBlock] = ""
	powerUpVehicleMAP[powerUpVehicleCurrentBlock+1] = ""
	powerUpVehicleMAP[powerUpVehicleCurrentBlock+94] = ""
	powerUpVehicleMAP[powerUpVehicleCurrentBlock+95] = ""
}

// MARK: enemy functions

func moveEnemyUpRight() {

	outerClearEnemyBlock := clearEnemyBlock - (94 * 2)
	enemyBoundaryLineCount := 0
	for a := 0; a < 16; a++ {
		enemiesMAP[outerClearEnemyBlock] = "boundary"
		outerClearEnemyBlock++
		enemyBoundaryLineCount++
		if enemyBoundaryLineCount == 4 {
			enemyBoundaryLineCount = 0
			outerClearEnemyBlock += 90
		}
	}

	enemiesMAP[clearEnemyBlock-93] = "enemyTL"
	enemiesMAP[clearEnemyBlock-92] = "enemy"
	enemiesMAP[clearEnemyBlock+1] = "enemy"
	enemiesMAP[clearEnemyBlock+2] = "enemy"

}

func drawNewEnemyPosition(newEnemyPosition int) {

	outerClearEnemyBlock := newEnemyPosition - 95
	enemyBoundaryLineCount := 0
	for a := 0; a < 16; a++ {
		enemiesMAP[outerClearEnemyBlock] = "boundary"
		outerClearEnemyBlock++
		enemyBoundaryLineCount++
		if enemyBoundaryLineCount == 4 {
			enemyBoundaryLineCount = 0
			outerClearEnemyBlock += 90
		}
	}

	enemiesMAP[newEnemyPosition] = "enemyTL"
	enemiesMAP[newEnemyPosition+1] = "enemy"
	enemiesMAP[newEnemyPosition+94] = "enemy"
	enemiesMAP[newEnemyPosition+95] = "enemy"
}

func findEnemyTL(collsionBlock int) int {

	enemyTLPostion := collsionBlock

	currentCollisionBlock := enemiesMAP[collsionBlock]

	if currentCollisionBlock == "enemyTL" {
		enemyTLPostion = collsionBlock
	} else if enemiesMAP[collsionBlock-94] == "enemyTL" {
		enemyTLPostion = collsionBlock - 94
	} else if enemiesMAP[collsionBlock-95] == "enemyTL" {
		enemyTLPostion = collsionBlock - 95
	} else if enemiesMAP[collsionBlock-1] == "enemyTL" {
		enemyTLPostion = collsionBlock - 1
	}

	return enemyTLPostion

}

func killEnemy() {
	deadEnemiesMAP[clearEnemyBlock] = "deadEnemy1_4"
	enemiesScreenCount--
	killzCount++
	totalKillz++
	if screenShake == false {
		screenShake = true
	}
	if enemyDieSound == false {
		enemyDieSound = true
	}

}

func moveEnemyUpStuck() {

	outerClearEnemyBlock := clearEnemyBlock - (94*5 + 1)
	enemyBoundaryLineCount := 0
	for a := 0; a < 16; a++ {
		enemiesMAP[outerClearEnemyBlock] = "boundary"
		outerClearEnemyBlock++
		enemyBoundaryLineCount++
		if enemyBoundaryLineCount == 4 {
			enemyBoundaryLineCount = 0
			outerClearEnemyBlock += 90
		}
	}
	enemiesMAP[clearEnemyBlock-(94*4)] = "enemyTL"
	enemiesMAP[clearEnemyBlock-(94*4-1)] = "enemy"
	enemiesMAP[clearEnemyBlock-(94*3)] = "enemy"
	enemiesMAP[clearEnemyBlock-(94*3+1)] = "enemy"

}

func checkEnemyStuckOnBlock(currentEnemyBlock int) {

	TLblock := levelMAP[currentEnemyBlock]
	TRblock := levelMAP[currentEnemyBlock+1]
	BLblock := levelMAP[currentEnemyBlock+94]
	BRblock := levelMAP[currentEnemyBlock+95]

	if TLblock == "floor" || TRblock == "floor" || BLblock == "floor" || BRblock == "floor" {

		enemyHorizontal := currentEnemyBlock

		if enemyHorizontal > 6 {
			clearEnemyBlock = currentEnemyBlock
			movementHOLDER := enemiesMovementMAP[currentEnemyBlock]
			clearEnemy()
			enemiesMovementMAP[currentEnemyBlock-(94*4)] = movementHOLDER
			moveEnemyUpStuck()

		}

	}
}

func checkRightEnemy(currentEnemyBlock int) bool {

	enemyIsRight := false

	rightEnemyBlock1 := enemiesMAP[currentEnemyBlock-91]
	rightEnemyBlock2 := enemiesMAP[currentEnemyBlock+3]
	rightEnemyBlock3 := enemiesMAP[currentEnemyBlock+97]
	rightEnemyBlock4 := enemiesMAP[currentEnemyBlock+191]

	if rightEnemyBlock1 == "boundary" || rightEnemyBlock2 == "boundary" || rightEnemyBlock3 == "boundary" || rightEnemyBlock4 == "boundary" {
		enemyIsRight = true
	}
	return enemyIsRight

}
func checkLeftEnemy(currentEnemyBlock int) bool {

	enemyIsLeft := false

	leftEnemyBlock1 := enemiesMAP[currentEnemyBlock-96]
	leftEnemyBlock2 := enemiesMAP[currentEnemyBlock-2]
	leftEnemyBlock3 := enemiesMAP[currentEnemyBlock+92]
	leftEnemyBlock4 := enemiesMAP[currentEnemyBlock+186]

	if leftEnemyBlock1 == "boundary" || leftEnemyBlock2 == "boundary" || leftEnemyBlock3 == "boundary" || leftEnemyBlock4 == "boundary" {
		enemyIsLeft = true
	}
	return enemyIsLeft

}
func checkLeftEnemy2(currentEnemyBlock int) bool {

	enemyIsLeft := false

	leftEnemyBlock1 := enemiesMAP[currentEnemyBlock-100]
	leftEnemyBlock2 := enemiesMAP[currentEnemyBlock-6]
	leftEnemyBlock3 := enemiesMAP[currentEnemyBlock+88]
	leftEnemyBlock4 := enemiesMAP[currentEnemyBlock+182]

	if leftEnemyBlock1 == "boundary" || leftEnemyBlock2 == "boundary" || leftEnemyBlock3 == "boundary" || leftEnemyBlock4 == "boundary" {
		enemyIsLeft = true
	}
	return enemyIsLeft

}
func checkAboveEnemy(currentEnemyBlock int) bool {
	enemyIsAbove := false

	currentEnemyBlockHorizontal := currentEnemyBlock / 94

	if currentEnemyBlockHorizontal > 3 {

		aboveEnemyBlock1 := enemiesMAP[currentEnemyBlock-189]
		aboveEnemyBlock2 := enemiesMAP[currentEnemyBlock-188]
		aboveEnemyBlock3 := enemiesMAP[currentEnemyBlock-187]
		aboveEnemyBlock4 := enemiesMAP[currentEnemyBlock-186]

		if aboveEnemyBlock1 == "boundary" || aboveEnemyBlock2 == "boundary" || aboveEnemyBlock3 == "boundary" || aboveEnemyBlock4 == "boundary" {
			enemyIsAbove = true
		}

	}
	return enemyIsAbove
}

func checkBelowEnemy(currentEnemyBlock int) bool {
	enemyIsBelow := false

	belowEnemyBlock1 := enemiesMAP[currentEnemyBlock+281]
	belowEnemyBlock2 := enemiesMAP[currentEnemyBlock+282]
	belowEnemyBlock3 := enemiesMAP[currentEnemyBlock+283]
	belowEnemyBlock4 := enemiesMAP[currentEnemyBlock+284]

	if belowEnemyBlock1 == "boundary" || belowEnemyBlock2 == "boundary" || belowEnemyBlock3 == "boundary" || belowEnemyBlock4 == "boundary" {

		enemyIsBelow = true
	}
	return enemyIsBelow
}

func moveEnemyTeleport() int {

	newEnemyPosition := rInt(568, 1400)

	for {

		checkLevelBlock := levelMAP[newEnemyPosition]
		checkLevelBlock1 := levelMAP[newEnemyPosition+1]
		checkLevelBlock2 := levelMAP[newEnemyPosition+94]
		checkLevelBlock3 := levelMAP[newEnemyPosition+95]

		newEnemyPosition++

		if checkLevelBlock == "" && checkLevelBlock1 == "" && checkLevelBlock2 == "" && checkLevelBlock3 == "" {
			break

		}

	}

	outerClearEnemyBlock := newEnemyPosition - 95
	enemyBoundaryLineCount := 0
	for a := 0; a < 16; a++ {
		enemiesMAP[outerClearEnemyBlock] = "boundary"
		outerClearEnemyBlock++
		enemyBoundaryLineCount++
		if enemyBoundaryLineCount == 4 {
			enemyBoundaryLineCount = 0
			outerClearEnemyBlock += 90
		}
	}

	enemiesMAP[newEnemyPosition] = "enemyTL"
	enemiesMAP[newEnemyPosition+1] = "enemy"
	enemiesMAP[newEnemyPosition+94] = "enemy"
	enemiesMAP[newEnemyPosition+95] = "enemy"

	return newEnemyPosition

}

func moveEnemyDown() {
	outerClearEnemyBlock := clearEnemyBlock - 1
	enemyBoundaryLineCount := 0
	for a := 0; a < 16; a++ {
		enemiesMAP[outerClearEnemyBlock] = "boundary"
		outerClearEnemyBlock++
		enemyBoundaryLineCount++
		if enemyBoundaryLineCount == 4 {
			enemyBoundaryLineCount = 0
			outerClearEnemyBlock += 90
		}
	}
	enemiesMAP[clearEnemyBlock+94] = "enemyTL"
	enemiesMAP[clearEnemyBlock+95] = "enemy"
	enemiesMAP[clearEnemyBlock+(94*2)] = "enemy"
	enemiesMAP[(clearEnemyBlock+1)+(94*2)] = "enemy"
}

func moveEnemyUp() {
	outerClearEnemyBlock := clearEnemyBlock - (94*2 + 1)
	enemyBoundaryLineCount := 0
	for a := 0; a < 16; a++ {
		enemiesMAP[outerClearEnemyBlock] = "boundary"
		outerClearEnemyBlock++
		enemyBoundaryLineCount++
		if enemyBoundaryLineCount == 4 {
			enemyBoundaryLineCount = 0
			outerClearEnemyBlock += 90
		}
	}
	enemiesMAP[clearEnemyBlock-94] = "enemyTL"
	enemiesMAP[clearEnemyBlock-93] = "enemy"
	enemiesMAP[clearEnemyBlock] = "enemy"
	enemiesMAP[clearEnemyBlock+1] = "enemy"
}
func moveEnemyLeft() {
	outerClearEnemyBlock := clearEnemyBlock - 96
	enemyBoundaryLineCount := 0
	for a := 0; a < 16; a++ {
		enemiesMAP[outerClearEnemyBlock] = "boundary"
		outerClearEnemyBlock++
		enemyBoundaryLineCount++
		if enemyBoundaryLineCount == 4 {
			enemyBoundaryLineCount = 0
			outerClearEnemyBlock += 90
		}
	}
	enemiesMAP[clearEnemyBlock-1] = "enemyTL"
	enemiesMAP[clearEnemyBlock] = "enemy"
	enemiesMAP[clearEnemyBlock+93] = "enemy"
	enemiesMAP[clearEnemyBlock+94] = "enemy"
}
func moveEnemyRight() {
	outerClearEnemyBlock := clearEnemyBlock - 94
	enemyBoundaryLineCount := 0
	for a := 0; a < 16; a++ {
		enemiesMAP[outerClearEnemyBlock] = "boundary"
		outerClearEnemyBlock++
		enemyBoundaryLineCount++
		if enemyBoundaryLineCount == 4 {
			enemyBoundaryLineCount = 0
			outerClearEnemyBlock += 90
		}
	}
	enemiesMAP[clearEnemyBlock+1] = "enemyTL"
	enemiesMAP[clearEnemyBlock+2] = "enemy"
	enemiesMAP[clearEnemyBlock+95] = "enemy"
	enemiesMAP[clearEnemyBlock+96] = "enemy"
}

func clearEnemy() {

	enemiesMovementMAP[clearEnemyBlock] = ""
	outerClearEnemyBlock := clearEnemyBlock - 95
	enemyBoundaryLineCount := 0
	for a := 0; a < 16; a++ {
		enemiesMAP[outerClearEnemyBlock] = ""
		outerClearEnemyBlock++
		enemyBoundaryLineCount++
		if enemyBoundaryLineCount == 4 {
			enemyBoundaryLineCount = 0
			outerClearEnemyBlock += 90
		}
	}

}

// MARK: player functions
func clearPlayer() {
	playerMAP[playerCurrentBlock] = ""
	playerMAP[playerCurrentBlock+1] = ""
	playerMAP[playerCurrentBlock+94] = ""
	playerMAP[playerCurrentBlock+95] = ""
}

// MARK: create enemies
func cENEMIES() {

	if totalLevels < 3 {
		enemyTypeGenerate = rInt(1, 4)
	} else if totalLevels < 6 && totalLevels >= 3 {
		if easyDiffOn {
			enemyTypeGenerate = rInt(1, 5)
		} else if averageDiffOn {
			enemyTypeGenerate = rInt(1, 6)
		} else if difficultDiffOn {
			enemyTypeGenerate = rInt(1, 6)
		}
	} else if totalLevels < 9 && totalLevels >= 6 {
		if easyDiffOn {
			enemyTypeGenerate = rInt(1, 5)
		} else if averageDiffOn {
			enemyTypeGenerate = rInt(1, 6)
		} else if difficultDiffOn {
			enemyTypeGenerate = rInt(1, 7)
		}
	} else if totalLevels < 11 && totalLevels >= 9 {
		if easyDiffOn {
			enemyTypeGenerate = rInt(1, 6)
		} else if averageDiffOn {
			enemyTypeGenerate = rInt(1, 7)
		} else if difficultDiffOn {
			enemyTypeGenerate = rInt(1, 8)
		}
	} else if totalLevels < 13 && totalLevels >= 11 {
		if easyDiffOn {
			enemyTypeGenerate = rInt(1, 7)
		} else if averageDiffOn {
			enemyTypeGenerate = rInt(1, 8)
		} else if difficultDiffOn {
			enemyTypeGenerate = rInt(1, 9)
		}
	} else if totalLevels < 15 && totalLevels >= 13 {
		if easyDiffOn {
			enemyTypeGenerate = rInt(1, 8)
		} else if averageDiffOn {
			enemyTypeGenerate = rInt(1, 9)
		} else if difficultDiffOn {
			enemyTypeGenerate = rInt(1, 10)
		}
	} else if totalLevels < 20 && totalLevels >= 15 {
		if easyDiffOn {
			enemyTypeGenerate = rInt(1, 8)
		} else if averageDiffOn {
			enemyTypeGenerate = rInt(1, 10)
		} else if difficultDiffOn {
			enemyTypeGenerate = rInt(1, 10)
		}
	} else if totalLevels >= 20 {
		enemyTypeGenerate = rInt(1, 10)
	}

	switch enemyTypeGenerate {
	case 1:
		enemyNumber = 5
		if averageDiffOn {
			enemyNumber++
		} else if difficultDiffOn {
			enemyNumber += 2
		}
		if totalLevels > 10 && totalLevels < 15 {
			enemiesScreenCount++
		} else if totalLevels > 15 && totalLevels < 20 {
			enemiesScreenCount += 2
		} else if totalLevels > 20 && totalLevels < 25 {
			enemiesScreenCount += 3
		} else if totalLevels > 25 {
			enemiesScreenCount += 4
		}
		enemiesScreenCount = enemyNumber
	case 2:
		enemyNumber = 5
		if averageDiffOn {
			enemyNumber++
		} else if difficultDiffOn {
			enemyNumber += 2
		}

		if totalLevels > 10 && totalLevels < 15 {
			enemiesScreenCount++
		} else if totalLevels > 15 && totalLevels < 20 {
			enemiesScreenCount += 2
		} else if totalLevels > 20 && totalLevels < 25 {
			enemiesScreenCount += 3
		} else if totalLevels > 25 {
			enemiesScreenCount += 4
		}
		enemiesScreenCount = enemyNumber
	case 3:
		enemyNumber = 3
		if averageDiffOn {
			enemyNumber++
		} else if difficultDiffOn {
			enemyNumber += 2
		}

		if totalLevels > 10 && totalLevels < 15 {
			enemiesScreenCount++
		} else if totalLevels > 15 && totalLevels < 20 {
			enemiesScreenCount += 2
		} else if totalLevels > 20 && totalLevels < 25 {
			enemiesScreenCount += 3
		} else if totalLevels > 25 {
			enemiesScreenCount += 4
		}
		enemiesScreenCount = enemyNumber
	case 4:
		enemyNumber = 4
		if averageDiffOn {
			enemyNumber++
		} else if difficultDiffOn {
			enemyNumber += 2
		}

		if totalLevels > 10 && totalLevels < 15 {
			enemiesScreenCount++
		} else if totalLevels > 15 && totalLevels < 20 {
			enemiesScreenCount += 2
		} else if totalLevels > 20 && totalLevels < 25 {
			enemiesScreenCount += 3
		} else if totalLevels > 25 {
			enemiesScreenCount += 4
		}
		enemiesScreenCount = enemyNumber
	case 5:
		enemyNumber = 5
		if averageDiffOn {
			enemyNumber++
		} else if difficultDiffOn {
			enemyNumber += 2
		}

		if totalLevels > 10 && totalLevels < 15 {
			enemiesScreenCount++
		} else if totalLevels > 15 && totalLevels < 20 {
			enemiesScreenCount += 2
		} else if totalLevels > 20 && totalLevels < 25 {
			enemiesScreenCount += 3
		} else if totalLevels > 25 {
			enemiesScreenCount += 4
		}
		enemiesScreenCount = enemyNumber
	case 6:
		enemyNumber = 5
		if averageDiffOn {
			enemyNumber++
		} else if difficultDiffOn {
			enemyNumber += 2
		}

		if totalLevels > 10 && totalLevels < 15 {
			enemiesScreenCount++
		} else if totalLevels > 15 && totalLevels < 20 {
			enemiesScreenCount += 2
		} else if totalLevels > 20 && totalLevels < 25 {
			enemiesScreenCount += 3
		} else if totalLevels > 25 {
			enemiesScreenCount += 4
		}
		enemiesScreenCount = enemyNumber
	case 7:
		enemyNumber = 4
		if averageDiffOn {
			enemyNumber++
		} else if difficultDiffOn {
			enemyNumber += 2
		}

		if totalLevels > 10 && totalLevels < 15 {
			enemiesScreenCount++
		} else if totalLevels > 15 && totalLevels < 20 {
			enemiesScreenCount += 2
		} else if totalLevels > 20 && totalLevels < 25 {
			enemiesScreenCount += 3
		} else if totalLevels > 25 {
			enemiesScreenCount += 4
		}

		enemy7RoofFall1 = rInt(24, 60)
		enemy7RoofFall2 = rInt(12, 50)
		enemy7RoofFall3 = rInt(40, 60)
		enemy7RoofFall4 = rInt(8, 40)
		enemy7RoofFall5 = rInt(20, 70)
		enemy7RoofFall6 = rInt(22, 62)

		enemiesScreenCount = enemyNumber
	case 8:
		enemyNumber = 5
		if averageDiffOn {
			enemyNumber++
		} else if difficultDiffOn {
			enemyNumber += 2
		}

		if totalLevels > 10 && totalLevels < 15 {
			enemiesScreenCount++
		} else if totalLevels > 15 && totalLevels < 20 {
			enemiesScreenCount += 2
		} else if totalLevels > 20 && totalLevels < 25 {
			enemiesScreenCount += 3
		} else if totalLevels > 25 {
			enemiesScreenCount += 4
		}
		enemiesScreenCount = enemyNumber
	case 9:
		enemyNumber = 5
		if averageDiffOn {
			enemyNumber++
		} else if difficultDiffOn {
			enemyNumber += 2
		}

		if totalLevels > 10 && totalLevels < 15 {
			enemiesScreenCount++
		} else if totalLevels > 15 && totalLevels < 20 {
			enemiesScreenCount += 2
		} else if totalLevels > 20 && totalLevels < 25 {
			enemiesScreenCount += 3
		} else if totalLevels > 25 {
			enemiesScreenCount += 4
		}
		enemiesScreenCount = enemyNumber

	}

	for {

		enemyPosition := rInt(5, 86)

		if enemiesMAP[enemyPosition] == "" {

			if enemiesMAP[enemyPosition+1] == "" && enemiesMAP[enemyPosition+2] == "" && enemiesMAP[enemyPosition+3] == "" {

				enemyBoundaryLineCount := 0
				enemyPositionHOLDER := enemyPosition
				for a := 0; a < 16; a++ {

					enemiesMAP[enemyPosition] = "boundary"
					enemyPosition++
					enemyBoundaryLineCount++
					if enemyBoundaryLineCount == 4 {
						enemyBoundaryLineCount = 0
						enemyPosition += 90
					}
				}
				innerEnemyPosition := enemyPositionHOLDER + 95

				enemiesMAP[innerEnemyPosition] = "enemyTL"
				enemiesMAP[innerEnemyPosition+1] = "enemy"
				enemiesMAP[innerEnemyPosition+94] = "enemy"
				enemiesMAP[innerEnemyPosition+95] = "enemy"

				rightLeft := flipcoin()

				if rightLeft {
					enemiesMovementMAP[innerEnemyPosition] = "left"
				} else {
					enemiesMovementMAP[innerEnemyPosition] = "right"
				}

				enemyNumber--
			}
		}
		if enemyNumber == 0 {
			break
		}

	}

}

// MARK: create level
func cLEVEL() {

	for a := 0; a < 5264; a++ {
		levelMAP[a] = ""
	}
	for a := 4230; a < 4700; a++ {
		levelMAP[a] = "floor"
	}

	floorType := rInt(1, 25)
	floorTypeNameTL = "ground" + strconv.Itoa(floorType) + "TL"

	for a := 4228; a < 4323; a += 6 {
		platformsMAP[a] = floorTypeNameTL
	}

	// place background objects ground level
	for a := 4042; a < 4136; a++ {

		placeBackgroundObject := rolldice() + rolldice()

		if placeBackgroundObject > 10 {

			backObjType := rInt(1, 45)
			backObjTypeName := "backObjGround" + strconv.Itoa(backObjType)

			backgroundObjectsMAP[a] = backObjTypeName
			a += 4
		}
	}
	// create platforms
	platformLayoutType := rInt(1, 16)

	switch platformLayoutType {

	case 1: // parallel lines

		platformStartBlock := rInt(668, 681)
		platformStartBlock += (rInt(1, 4) * 94)
		platformLength := rInt(30, 61)

		platformType := rInt(1, 44)
		platformTypeName = "platform" + strconv.Itoa(platformType)
		platformTypeNameTL = platformTypeName + "TL"

		for {

			for a := 0; a < platformLength; a++ {
				levelMAP[platformStartBlock+a] = "floor"
				levelMAP[platformStartBlock+(a+1)] = "floor"
				levelMAP[platformStartBlock+(a+94)] = "floor"
				levelMAP[platformStartBlock+(a+95)] = "floor"

				platformsMAP[platformStartBlock+a] = platformTypeNameTL
				platformsMAP[platformStartBlock+(a+1)] = platformTypeName
				platformsMAP[platformStartBlock+(a+94)] = platformTypeName
				platformsMAP[platformStartBlock+(a+95)] = platformTypeName

				// place background object
				placeBackgroundObject := rolldice() + rolldice()
				if placeBackgroundObject > 10 {
					backObjType := rInt(1, 25)
					backObjTypeName := "backObj" + strconv.Itoa(backObjType)
					backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = backObjTypeName
				}

				a++
			}

			platformStartBlockHorizontal := platformStartBlock / 94

			if platformStartBlockHorizontal > 30 {
				break
			}

			platformStartBlock += (rInt(8, 12) * 94)
			platformStartBlock += rInt(-4, 5)
			platformLength = rInt(30, 61)

		}

	case 2: // short random plaforms

		platformType := rInt(1, 42)
		platformTypeName = "platform" + strconv.Itoa(platformType)
		platformTypeNameTL = platformTypeName + "TL"
		platformsRows := rInt(4, 7)
		maxPlatformLength := (90 / platformsRows) - 2
		platformLength := rInt(4, maxPlatformLength)
		if platformLength%2 != 0 {
			platformLength++
		}
		platformStartBlock := (rInt(8, 15) * 94) + rInt(6, 10)
		platformStartBlockHOLDER := platformStartBlock

		for {

			for {
				for a := 0; a < platformLength; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					// place background object
					placeBackgroundObject := rolldice() + rolldice()
					if placeBackgroundObject > 10 {
						backObjType := rInt(1, 25)
						backObjTypeName := "backObj" + strconv.Itoa(backObjType)
						backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = backObjTypeName
					}

					a++

				}
				platformStartBlock += (rInt(4, 11) * 94) + (rInt(-2, 3))
				platformLength = rInt(4, maxPlatformLength)

				platformStartBlockHorizontal := platformStartBlock / 94

				if platformStartBlockHorizontal > 38 {
					break
				}
			}

			platformStartBlock = platformStartBlockHOLDER
			platformStartBlockHorizontal := platformStartBlock / 94
			platformStartBlockVertical := platformStartBlock - (platformStartBlockHorizontal * 94)

			if platformStartBlockVertical >= 70 {
				break
			}

			platformStartBlock += maxPlatformLength + (rInt(2, 5)) + (rInt(-2, 3) * 94)
			platformStartBlockHOLDER = platformStartBlock

		}

	case 3: // large columns X4

		platformType := rInt(1, 42)
		platformTypeName = "platform" + strconv.Itoa(platformType)
		platformTypeNameTL = platformTypeName + "TL"

		platformStartBlock := rInt(664, 675)
		platformStartBlock += (rInt(1, 6) * 94)

		platformStartBlockHOLDER := platformStartBlock

		for {

			platformLength := rInt(8, 15)
			if platformLength%2 != 0 {
				platformLength++
			}

			for a := 0; a < platformLength; a++ {
				levelMAP[platformStartBlock+a] = "floor"
				levelMAP[platformStartBlock+(a+1)] = "floor"
				levelMAP[platformStartBlock+(a+94)] = "floor"
				levelMAP[platformStartBlock+(a+95)] = "floor"

				platformsMAP[platformStartBlock+a] = platformTypeNameTL
				platformsMAP[platformStartBlock+(a+1)] = platformTypeName
				platformsMAP[platformStartBlock+(a+94)] = platformTypeName
				platformsMAP[platformStartBlock+(a+95)] = platformTypeName

				// place background object
				placeBackgroundObject := rolldice() + rolldice()
				if placeBackgroundObject > 10 {
					backObjType := rInt(1, 25)
					backObjTypeName := "backObj" + strconv.Itoa(backObjType)
					backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = backObjTypeName
				}
				a++

			}

			platformStartBlock += 94 * 4

			for {

				levelMAP[platformStartBlock] = "floor"
				platformsMAP[platformStartBlock] = platformTypeNameTL
				platformStartBlock++
				levelMAP[platformStartBlock] = "floor"
				platformsMAP[platformStartBlock] = platformTypeName
				platformStartBlock += 93
				levelMAP[platformStartBlock] = "floor"
				platformsMAP[platformStartBlock] = platformTypeName
				platformStartBlock++
				levelMAP[platformStartBlock] = "floor"
				platformsMAP[platformStartBlock] = platformTypeName

				platformStartBlockHorizontal := platformStartBlock / 94
				//	platformStartBlockVertical := platformStartBlock - (platformStartBlockHorizontal * 94)

				if platformStartBlockHorizontal > 30 {
					break
				}

				platformStartBlock += 94 * 4
				platformStartBlock--
			}

			platformStartBlock += 94 * 4
			platformStartBlock--

			for a := 0; a < platformLength; a++ {
				levelMAP[platformStartBlock+a] = "floor"
				levelMAP[platformStartBlock+(a+1)] = "floor"
				levelMAP[platformStartBlock+(a+94)] = "floor"
				levelMAP[platformStartBlock+(a+95)] = "floor"

				platformsMAP[platformStartBlock+a] = platformTypeNameTL
				platformsMAP[platformStartBlock+(a+1)] = platformTypeName
				platformsMAP[platformStartBlock+(a+94)] = platformTypeName
				platformsMAP[platformStartBlock+(a+95)] = platformTypeName

				// place background object
				placeBackgroundObject := rolldice() + rolldice()
				if placeBackgroundObject > 10 {
					backObjType := rInt(1, 25)
					backObjTypeName := "backObj" + strconv.Itoa(backObjType)
					backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = backObjTypeName
				}

				a++
			}

			platformStartBlock = platformStartBlockHOLDER
			platformStartBlock += platformLength - 2

			platformStartBlock += 94 * 4

			for {

				levelMAP[platformStartBlock] = "floor"
				platformsMAP[platformStartBlock] = platformTypeNameTL
				platformStartBlock++
				levelMAP[platformStartBlock] = "floor"
				platformsMAP[platformStartBlock] = platformTypeName
				platformStartBlock += 93
				levelMAP[platformStartBlock] = "floor"
				platformsMAP[platformStartBlock] = platformTypeName
				platformStartBlock++
				levelMAP[platformStartBlock] = "floor"
				platformsMAP[platformStartBlock] = platformTypeName

				platformStartBlockHorizontal := platformStartBlock / 94
				//	platformStartBlockVertical := platformStartBlock - (platformStartBlockHorizontal * 94)
				if platformStartBlockHorizontal > 30 {
					break
				}

				platformStartBlock += 94 * 4
				platformStartBlock--
			}

			platformStartBlock = platformStartBlockHOLDER + platformLength

			platformStartBlock += rInt(4, 9)

			platformStartBlockHOLDER = platformStartBlock

			platformStartBlockHorizontal := platformStartBlock / 94
			platformStartBlockVertical := platformStartBlock - (platformStartBlockHorizontal * 94)
			if platformStartBlockVertical > 70 {
				break
			}

		}
	case 4: // 3X short platforms left / 1 large column right

		platformType := rInt(1, 42)
		platformTypeName = "platform" + strconv.Itoa(platformType)
		platformTypeNameTL = platformTypeName + "TL"

		platformStartBlock := rInt(704, 720)
		platformStartBlock += (rInt(1, 6) * 94)

		platformStartBlockHOLDER := platformStartBlock

		platformLength := rInt(10, 15)
		if platformLength%2 != 0 {
			platformLength++
		}

		for a := 0; a < platformLength; a++ {
			levelMAP[platformStartBlock+a] = "floor"
			levelMAP[platformStartBlock+(a+1)] = "floor"
			levelMAP[platformStartBlock+(a+94)] = "floor"
			levelMAP[platformStartBlock+(a+95)] = "floor"

			platformsMAP[platformStartBlock+a] = platformTypeNameTL
			platformsMAP[platformStartBlock+(a+1)] = platformTypeName
			platformsMAP[platformStartBlock+(a+94)] = platformTypeName
			platformsMAP[platformStartBlock+(a+95)] = platformTypeName

			// place background object
			placeBackgroundObject := rolldice() + rolldice()
			if placeBackgroundObject > 10 {
				backObjType := rInt(1, 25)
				backObjTypeName := "backObj" + strconv.Itoa(backObjType)
				backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = backObjTypeName
			}
			a++
		}

		platformStartBlock += 94 * 4

		for {

			levelMAP[platformStartBlock] = "floor"
			platformsMAP[platformStartBlock] = platformTypeNameTL
			platformStartBlock++
			levelMAP[platformStartBlock] = "floor"
			platformsMAP[platformStartBlock] = platformTypeName
			platformStartBlock += 93
			levelMAP[platformStartBlock] = "floor"
			platformsMAP[platformStartBlock] = platformTypeName
			platformStartBlock++
			levelMAP[platformStartBlock] = "floor"
			platformsMAP[platformStartBlock] = platformTypeName

			platformStartBlockHorizontal := platformStartBlock / 94
			//	platformStartBlockVertical := platformStartBlock - (platformStartBlockHorizontal * 94)

			if platformStartBlockHorizontal > 30 {
				break
			}

			platformStartBlock += 94 * 4
			platformStartBlock--
		}

		platformStartBlock += 94 * 4
		platformStartBlock--

		for a := 0; a < platformLength; a++ {
			levelMAP[platformStartBlock+a] = "floor"
			levelMAP[platformStartBlock+(a+1)] = "floor"
			levelMAP[platformStartBlock+(a+94)] = "floor"
			levelMAP[platformStartBlock+(a+95)] = "floor"

			platformsMAP[platformStartBlock+a] = platformTypeNameTL
			platformsMAP[platformStartBlock+(a+1)] = platformTypeName
			platformsMAP[platformStartBlock+(a+94)] = platformTypeName
			platformsMAP[platformStartBlock+(a+95)] = platformTypeName

			// place background object
			placeBackgroundObject := rolldice() + rolldice()
			if placeBackgroundObject > 10 {
				backObjType := rInt(1, 25)
				backObjTypeName := "backObj" + strconv.Itoa(backObjType)
				backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = backObjTypeName
			}
			a++
		}

		platformStartBlock = platformStartBlockHOLDER
		platformStartBlock += platformLength - 2

		platformStartBlock += 94 * 4

		for {

			levelMAP[platformStartBlock] = "floor"
			platformsMAP[platformStartBlock] = platformTypeNameTL
			platformStartBlock++
			levelMAP[platformStartBlock] = "floor"
			platformsMAP[platformStartBlock] = platformTypeName
			platformStartBlock += 93
			levelMAP[platformStartBlock] = "floor"
			platformsMAP[platformStartBlock] = platformTypeName
			platformStartBlock++
			levelMAP[platformStartBlock] = "floor"
			platformsMAP[platformStartBlock] = platformTypeName

			platformStartBlockHorizontal := platformStartBlock / 94
			//	platformStartBlockVertical := platformStartBlock - (platformStartBlockHorizontal * 94)
			if platformStartBlockHorizontal > 30 {
				break
			}

			platformStartBlock += 94 * 4
			platformStartBlock--
		}

		platformStartBlock = rInt(664, 674)
		platformStartBlock += (rInt(1, 5) * 94)
		platformLength = rInt(4, 7)
		if platformLength%2 != 0 {
			platformLength++
		}
		platformStartBlockHOLDER = platformStartBlock

		for a := 0; a < 3; a++ {
			for {

				for a := 0; a < platformLength; a++ {
					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					// place background object
					placeBackgroundObject := rolldice() + rolldice()
					if placeBackgroundObject > 10 {
						backObjType := rInt(1, 25)
						backObjTypeName := "backObj" + strconv.Itoa(backObjType)
						backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = backObjTypeName
					}
					a++
				}

				platformStartBlockHorizontal := platformStartBlock / 94
				//	platformStartBlockVertical := platformStartBlock - (platformStartBlockHorizontal * 94)
				if platformStartBlockHorizontal > 30 {
					break
				}
				platformStartBlock += (rInt(4, 9) * 94)
			}

			platformStartBlock = platformStartBlockHOLDER + platformLength + 4
			platformLength = rInt(4, 9)
			if platformLength%2 != 0 {
				platformLength++
			}
			platformStartBlockHOLDER = platformStartBlock
		}

	case 5: // 2X inverted pyramids

		platformType := rInt(1, 42)
		platformTypeName = "platform" + strconv.Itoa(platformType)
		platformTypeNameTL = platformTypeName + "TL"
		platformStartBlock := rInt(1040, 1053)
		platformStartBlock += (rInt(1, 6) * 94)

		platformLength := rInt(20, 35)
		if platformLength%2 != 0 {
			platformLength++
		}

		platformStartBlockHOLDER := platformStartBlock + platformLength + 2

		for a := 0; a < platformLength; a++ {
			levelMAP[platformStartBlock+a] = "floor"
			levelMAP[platformStartBlock+(a+1)] = "floor"
			levelMAP[platformStartBlock+(a+94)] = "floor"
			levelMAP[platformStartBlock+(a+95)] = "floor"

			platformsMAP[platformStartBlock+a] = platformTypeNameTL
			platformsMAP[platformStartBlock+(a+1)] = platformTypeName
			platformsMAP[platformStartBlock+(a+94)] = platformTypeName
			platformsMAP[platformStartBlock+(a+95)] = platformTypeName

			// place background object
			placeBackgroundObject := rolldice() + rolldice()
			if placeBackgroundObject > 10 {
				backObjType := rInt(1, 25)
				backObjTypeName := "backObj" + strconv.Itoa(backObjType)
				backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = backObjTypeName
			}
			a++
		}

		for {
			platformStartBlock += 94 * 4
			platformStartBlock += 2

			platformLength = 2
			for a := 0; a < platformLength; a++ {
				levelMAP[platformStartBlock+a] = "floor"
				levelMAP[platformStartBlock+(a+1)] = "floor"
				levelMAP[platformStartBlock+(a+94)] = "floor"
				levelMAP[platformStartBlock+(a+95)] = "floor"

				platformsMAP[platformStartBlock+a] = platformTypeNameTL
				platformsMAP[platformStartBlock+(a+1)] = platformTypeName
				platformsMAP[platformStartBlock+(a+94)] = platformTypeName
				platformsMAP[platformStartBlock+(a+95)] = platformTypeName

				// place background object
				placeBackgroundObject := rolldice() + rolldice()
				if placeBackgroundObject > 10 {
					backObjType := rInt(1, 25)
					backObjTypeName := "backObj" + strconv.Itoa(backObjType)
					backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = backObjTypeName
				}
				a++
			}

			platformStartBlockHorizontal := platformStartBlock / 94
			//	platformStartBlockVertical := platformStartBlock - (platformStartBlockHorizontal * 94)
			if platformStartBlockHorizontal > 34 {
				break
			}
		}

		for {
			platformStartBlock -= 94 * 4
			platformStartBlock += 2

			platformLength = 2
			for a := 0; a < platformLength; a++ {
				levelMAP[platformStartBlock+a] = "floor"
				levelMAP[platformStartBlock+(a+1)] = "floor"
				levelMAP[platformStartBlock+(a+94)] = "floor"
				levelMAP[platformStartBlock+(a+95)] = "floor"

				platformsMAP[platformStartBlock+a] = platformTypeNameTL
				platformsMAP[platformStartBlock+(a+1)] = platformTypeName
				platformsMAP[platformStartBlock+(a+94)] = platformTypeName
				platformsMAP[platformStartBlock+(a+95)] = platformTypeName

				// place background object
				placeBackgroundObject := rolldice() + rolldice()
				if placeBackgroundObject > 10 {
					backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = "backgroundObject"
				}
				a++
			}

			platformStartBlockHorizontal := platformStartBlock / 94
			//	platformStartBlockVertical := platformStartBlock - (platformStartBlockHorizontal * 94)
			if platformStartBlockHorizontal < 16 {
				break
			}
		}

		platformStartBlock = platformStartBlockHOLDER

		platformLength = rInt(20, 35)
		if platformLength%2 != 0 {
			platformLength++
		}
		//	platformLengthHalf := platformLength / 2

		for a := 0; a < platformLength; a++ {
			levelMAP[platformStartBlock+a] = "floor"
			levelMAP[platformStartBlock+(a+1)] = "floor"
			levelMAP[platformStartBlock+(a+94)] = "floor"
			levelMAP[platformStartBlock+(a+95)] = "floor"

			platformsMAP[platformStartBlock+a] = platformTypeNameTL
			platformsMAP[platformStartBlock+(a+1)] = platformTypeName
			platformsMAP[platformStartBlock+(a+94)] = platformTypeName
			platformsMAP[platformStartBlock+(a+95)] = platformTypeName

			// place background object
			placeBackgroundObject := rolldice() + rolldice()
			if placeBackgroundObject > 10 {
				backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = "backgroundObject"
			}
			a++
		}

		for {
			platformStartBlock += 94 * 4
			platformStartBlock += 2

			platformLength = 2
			for a := 0; a < platformLength; a++ {
				levelMAP[platformStartBlock+a] = "floor"
				levelMAP[platformStartBlock+(a+1)] = "floor"
				levelMAP[platformStartBlock+(a+94)] = "floor"
				levelMAP[platformStartBlock+(a+95)] = "floor"

				platformsMAP[platformStartBlock+a] = platformTypeNameTL
				platformsMAP[platformStartBlock+(a+1)] = platformTypeName
				platformsMAP[platformStartBlock+(a+94)] = platformTypeName
				platformsMAP[platformStartBlock+(a+95)] = platformTypeName

				// place background object
				placeBackgroundObject := rolldice() + rolldice()
				if placeBackgroundObject > 10 {
					backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = "backgroundObject"
				}
				a++
			}

			platformStartBlockHorizontal := platformStartBlock / 94
			//	platformStartBlockVertical := platformStartBlock - (platformStartBlockHorizontal * 94)
			if platformStartBlockHorizontal > 34 {
				break
			}
		}

		for {
			platformStartBlock -= 94 * 4
			platformStartBlock += 2

			platformLength = 2
			for a := 0; a < platformLength; a++ {
				levelMAP[platformStartBlock+a] = "floor"
				levelMAP[platformStartBlock+(a+1)] = "floor"
				levelMAP[platformStartBlock+(a+94)] = "floor"
				levelMAP[platformStartBlock+(a+95)] = "floor"

				platformsMAP[platformStartBlock+a] = platformTypeNameTL
				platformsMAP[platformStartBlock+(a+1)] = platformTypeName
				platformsMAP[platformStartBlock+(a+94)] = platformTypeName
				platformsMAP[platformStartBlock+(a+95)] = platformTypeName

				// place background object
				placeBackgroundObject := rolldice() + rolldice()
				if placeBackgroundObject > 10 {
					backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = "backgroundObject"
				}
				a++
			}

			platformStartBlockHorizontal := platformStartBlock / 94
			//	platformStartBlockVertical := platformStartBlock - (platformStartBlockHorizontal * 94)
			if platformStartBlockHorizontal < 16 {
				break
			}
		}
	case 6: // boxes

		platformType := rInt(1, 42)
		platformTypeName = "platform" + strconv.Itoa(platformType)
		platformTypeNameTL = platformTypeName + "TL"
		platformStartBlock := rInt(1040, 1054)
		platformStartBlock += (rInt(1, 15) * 94)

		for {

			platformLength := rInt(6, 13)
			if platformLength%2 != 0 {
				platformLength++
			}

			platformStartBlockHOLDER := platformStartBlock + (platformLength - 2)

			for a := 0; a < platformLength; a++ {
				levelMAP[platformStartBlock+a] = "floor"
				levelMAP[platformStartBlock+(a+1)] = "floor"
				levelMAP[platformStartBlock+(a+94)] = "floor"
				levelMAP[platformStartBlock+(a+95)] = "floor"

				platformsMAP[platformStartBlock+a] = platformTypeNameTL
				platformsMAP[platformStartBlock+(a+1)] = platformTypeName
				platformsMAP[platformStartBlock+(a+94)] = platformTypeName
				platformsMAP[platformStartBlock+(a+95)] = platformTypeName

				// place background object
				placeBackgroundObject := rolldice() + rolldice()
				if placeBackgroundObject > 10 {
					backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = "backgroundObject"
				}

				a++
			}

			platformStartBlock += (94 * 2)

			for a := 0; a < platformLength; a++ {
				levelMAP[platformStartBlock] = "floor"
				levelMAP[platformStartBlock+1] = "floor"
				levelMAP[platformStartBlock+94] = "floor"
				levelMAP[platformStartBlock+95] = "floor"

				platformsMAP[platformStartBlock] = platformTypeNameTL
				platformsMAP[platformStartBlock+1] = platformTypeName
				platformsMAP[platformStartBlock+94] = platformTypeName
				platformsMAP[platformStartBlock+95] = platformTypeName

				a++
				platformStartBlock += (94 * 2)
			}

			for a := 0; a < platformLength; a++ {
				levelMAP[platformStartBlock+a] = "floor"
				levelMAP[platformStartBlock+(a+1)] = "floor"
				levelMAP[platformStartBlock+(a+94)] = "floor"
				levelMAP[platformStartBlock+(a+95)] = "floor"

				platformsMAP[platformStartBlock+a] = platformTypeNameTL
				platformsMAP[platformStartBlock+(a+1)] = platformTypeName
				platformsMAP[platformStartBlock+(a+94)] = platformTypeName
				platformsMAP[platformStartBlock+(a+95)] = platformTypeName
				a++
			}

			platformStartBlock = platformStartBlockHOLDER

			platformStartBlock += (94 * 2)

			for a := 0; a < platformLength; a++ {
				levelMAP[platformStartBlock] = "floor"
				levelMAP[platformStartBlock+1] = "floor"
				levelMAP[platformStartBlock+94] = "floor"
				levelMAP[platformStartBlock+95] = "floor"

				platformsMAP[platformStartBlock] = platformTypeNameTL
				platformsMAP[platformStartBlock+1] = platformTypeName
				platformsMAP[platformStartBlock+94] = platformTypeName
				platformsMAP[platformStartBlock+95] = platformTypeName

				a++
				platformStartBlock += (94 * 2)
			}

			platformStartBlockHorizontal := platformStartBlock / 94
			platformStartBlockVertical := platformStartBlock - (platformStartBlockHorizontal * 94)

			if platformStartBlockVertical > 70 {
				break
			}

			platformStartBlock = platformStartBlockHOLDER

			platformStartBlockHorizontal = platformStartBlock / 94

			if platformStartBlockHorizontal > 12 {

				platformStartBlock += (rInt(-4, 4) * 94)
				platformStartBlock += rInt(4, 10)

			} else {
				platformStartBlock += (rInt(1, 4) * 94)
				platformStartBlock += rInt(4, 10)
			}
		}
	case 7: // steps up right

		platformType := rInt(1, 42)
		platformTypeName = "platform" + strconv.Itoa(platformType)
		platformTypeNameTL = platformTypeName + "TL"
		platformStartBlock := rInt(3482, 3494)

		platformLength := 4

		for {

			for a := 0; a < platformLength; a++ {
				levelMAP[platformStartBlock+a] = "floor"
				levelMAP[platformStartBlock+(a+1)] = "floor"
				levelMAP[platformStartBlock+(a+94)] = "floor"
				levelMAP[platformStartBlock+(a+95)] = "floor"

				platformsMAP[platformStartBlock+a] = platformTypeNameTL
				platformsMAP[platformStartBlock+(a+1)] = platformTypeName
				platformsMAP[platformStartBlock+(a+94)] = platformTypeName
				platformsMAP[platformStartBlock+(a+95)] = platformTypeName

				// place background object
				placeBackgroundObject := rolldice() + rolldice()
				if placeBackgroundObject > 10 {
					backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = "backgroundObject"
				}

				a++
			}

			platformStartBlockHorizontal := platformStartBlock / 94

			if platformStartBlockHorizontal < 12 {
				break
			}

			platformStartBlock -= (94 * 2)
			platformStartBlock += 4

		}

	case 8: // steps up left

		platformType := rInt(1, 42)
		platformTypeName = "platform" + strconv.Itoa(platformType)
		platformTypeNameTL = platformTypeName + "TL"
		platformStartBlock := rInt(662, 674)

		platformLength := 4

		for {

			for a := 0; a < platformLength; a++ {
				levelMAP[platformStartBlock+a] = "floor"
				levelMAP[platformStartBlock+(a+1)] = "floor"
				levelMAP[platformStartBlock+(a+94)] = "floor"
				levelMAP[platformStartBlock+(a+95)] = "floor"

				platformsMAP[platformStartBlock+a] = platformTypeNameTL
				platformsMAP[platformStartBlock+(a+1)] = platformTypeName
				platformsMAP[platformStartBlock+(a+94)] = platformTypeName
				platformsMAP[platformStartBlock+(a+95)] = platformTypeName

				// place background object
				placeBackgroundObject := rolldice() + rolldice()
				if placeBackgroundObject > 10 {
					backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = "backgroundObject"
				}

				a++
			}

			platformStartBlockHorizontal := platformStartBlock / 94

			if platformStartBlockHorizontal > 38 {
				break
			}

			platformStartBlock += (94 * 2)
			platformStartBlock += 4

		}

	case 9: // steps pyramid

		platformType := rInt(1, 42)
		platformTypeName = "platform" + strconv.Itoa(platformType)
		platformTypeNameTL = platformTypeName + "TL"

		platformStartBlock := rInt(3482, 3494)

		platformLength := 4

		for {

			for a := 0; a < platformLength; a++ {
				levelMAP[platformStartBlock+a] = "floor"
				levelMAP[platformStartBlock+(a+1)] = "floor"
				levelMAP[platformStartBlock+(a+94)] = "floor"
				levelMAP[platformStartBlock+(a+95)] = "floor"

				platformsMAP[platformStartBlock+a] = platformTypeNameTL
				platformsMAP[platformStartBlock+(a+1)] = platformTypeName
				platformsMAP[platformStartBlock+(a+94)] = platformTypeName
				platformsMAP[platformStartBlock+(a+95)] = platformTypeName

				// place background object
				placeBackgroundObject := rolldice() + rolldice()
				if placeBackgroundObject > 10 {
					backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = "backgroundObject"
				}

				a++
			}

			platformStartBlockHorizontal := platformStartBlock / 94

			if platformStartBlockHorizontal < 24 {
				break
			}

			platformStartBlock -= (94 * 2)
			platformStartBlock += 4
		}
		for {

			for a := 0; a < platformLength; a++ {
				levelMAP[platformStartBlock+a] = "floor"
				levelMAP[platformStartBlock+(a+1)] = "floor"
				levelMAP[platformStartBlock+(a+94)] = "floor"
				levelMAP[platformStartBlock+(a+95)] = "floor"

				platformsMAP[platformStartBlock+a] = platformTypeNameTL
				platformsMAP[platformStartBlock+(a+1)] = platformTypeName
				platformsMAP[platformStartBlock+(a+94)] = platformTypeName
				platformsMAP[platformStartBlock+(a+95)] = platformTypeName

				// place background object
				placeBackgroundObject := rolldice() + rolldice()
				if placeBackgroundObject > 10 {
					backgroundObjectsMAP[(platformStartBlock+a)-(94*2)] = "backgroundObject"
				}

				a++
			}

			platformStartBlockHorizontal := platformStartBlock / 94

			if platformStartBlockHorizontal > 38 {
				break
			}

			platformStartBlock += (94 * 2)
			platformStartBlock += 4
		}

	case 10: // zigzag columns

		platformType := rInt(1, 42)
		platformTypeName = "platform" + strconv.Itoa(platformType)
		platformTypeNameTL = platformTypeName + "TL"
		platformStartBlock := rInt(3482, 3494)
		platformStartBlockHOLDER := platformStartBlock
		platformLength := 4

		platformDirectionCount := 0
		platformDirection := false

		platformColumnCount := rInt(2, 5)

		for {

			for {

				if platformDirection == false {
					for a := 0; a < platformLength; a++ {
						levelMAP[platformStartBlock+a] = "floor"
						levelMAP[platformStartBlock+(a+1)] = "floor"
						levelMAP[platformStartBlock+(a+94)] = "floor"
						levelMAP[platformStartBlock+(a+95)] = "floor"

						platformsMAP[platformStartBlock+a] = platformTypeNameTL
						platformsMAP[platformStartBlock+(a+1)] = platformTypeName
						platformsMAP[platformStartBlock+(a+94)] = platformTypeName
						platformsMAP[platformStartBlock+(a+95)] = platformTypeName

						a++
					}
					platformDirectionCount++
					platformStartBlock -= (94 * 2)
					platformStartBlock += 4
				}
				if platformDirection {
					for a := 0; a < platformLength; a++ {
						levelMAP[platformStartBlock+a] = "floor"
						levelMAP[platformStartBlock+(a+1)] = "floor"
						levelMAP[platformStartBlock+(a+94)] = "floor"
						levelMAP[platformStartBlock+(a+95)] = "floor"

						platformsMAP[platformStartBlock+a] = platformTypeNameTL
						platformsMAP[platformStartBlock+(a+1)] = platformTypeName
						platformsMAP[platformStartBlock+(a+94)] = platformTypeName
						platformsMAP[platformStartBlock+(a+95)] = platformTypeName
						a++
					}
					platformDirectionCount++
					platformStartBlock -= (94 * 2)
					platformStartBlock -= 4
				}

				if platformDirectionCount == 3 {

					if platformDirection {
						platformDirection = false
					} else {
						platformDirection = true
					}
					platformDirectionCount = 0

				}

				platformStartBlockHorizontal := platformStartBlock / 94

				if platformStartBlockHorizontal < 12 {
					platformColumnCount--
					platformDirectionCount = 0
					platformDirection = false
					platformStartBlock = platformStartBlockHOLDER
					platformStartBlock += rInt(20, 25)
					platformStartBlockHOLDER = platformStartBlock
					break
				}
			}

			if platformColumnCount == 0 {
				break
			}
		}

	case 11:
		// horizontal line top with gap

		platformType := rInt(1, 42)
		platformTypeName = "platform" + strconv.Itoa(platformType)
		platformTypeNameTL = platformTypeName + "TL"

		platformStartBlock := 940
		platformStartBlock += (rInt(-3, 4) * 94)
		platformStartBlockHOLDER := platformStartBlock

		for b := 0; b < 4; b++ {
			for a := 0; a < 46; a++ {
				clearAroundBlock(platformStartBlock)
				platformStartBlock++
			}

			platformStartBlock = platformStartBlockHOLDER

			for a := 0; a < 45; a++ {

				levelMAP[platformStartBlock+a] = "floor"
				levelMAP[platformStartBlock+(a+1)] = "floor"
				levelMAP[platformStartBlock+(a+94)] = "floor"
				levelMAP[platformStartBlock+(a+95)] = "floor"

				platformsMAP[platformStartBlock+a] = platformTypeNameTL
				platformsMAP[platformStartBlock+(a+1)] = platformTypeName
				platformsMAP[platformStartBlock+(a+94)] = platformTypeName
				platformsMAP[platformStartBlock+(a+95)] = platformTypeName

				platformStartBlock++
			}

			platformStartBlock = platformStartBlockHOLDER
			platformStartBlock += rInt(30, 50)
			gapLength := rInt(6, 15)

			// gap
			for a := 0; a < gapLength; a++ {

				levelMAP[platformStartBlock+a] = ""
				levelMAP[platformStartBlock+(a+1)] = ""
				levelMAP[platformStartBlock+(a+94)] = ""
				levelMAP[platformStartBlock+(a+95)] = ""

				platformsMAP[platformStartBlock+a] = ""
				platformsMAP[platformStartBlock+(a+1)] = ""
				platformsMAP[platformStartBlock+(a+94)] = ""
				platformsMAP[platformStartBlock+(a+95)] = ""

				platformStartBlock++
			}
			platformStartBlock = platformStartBlockHOLDER
			platformStartBlock += (rInt(6, 11) * 94)
			platformStartBlockHOLDER = platformStartBlock

		}

	case 12:
		platformType := rInt(1, 42)
		platformTypeName = "platform" + strconv.Itoa(platformType)
		platformTypeNameTL = platformTypeName + "TL"

		platformStartBlock := 940 + rInt(4, 20)
		platformStartBlock += (rInt(-3, 4) * 94)

		for {
			levelMAP[platformStartBlock] = "floor"
			levelMAP[platformStartBlock+1] = "floor"
			levelMAP[platformStartBlock+94] = "floor"
			levelMAP[platformStartBlock+95] = "floor"

			platformsMAP[platformStartBlock] = platformTypeNameTL
			platformsMAP[platformStartBlock+1] = platformTypeName
			platformsMAP[platformStartBlock+94] = platformTypeName
			platformsMAP[platformStartBlock+95] = platformTypeName

			anotherBlock := flipcoin()

			if anotherBlock {
				platformStartBlock += 2
				levelMAP[platformStartBlock] = "floor"
				levelMAP[platformStartBlock+1] = "floor"
				levelMAP[platformStartBlock+94] = "floor"
				levelMAP[platformStartBlock+95] = "floor"

				platformsMAP[platformStartBlock] = platformTypeNameTL
				platformsMAP[platformStartBlock+1] = platformTypeName
				platformsMAP[platformStartBlock+94] = platformTypeName
				platformsMAP[platformStartBlock+95] = platformTypeName
			}

			platformStartBlock += rInt(4, 15)
			platformStartBlock += (rInt(1, 3) * 94)

			blockHorizontal := platformStartBlock / 94

			if blockHorizontal > 40 {
				break
			}

		}

		// zigzag top
		platformStartBlock = rInt(1136, 1156)
		platformStartBlock += (rInt(0, 11) * 94)

		zigzagNumber := rInt(2, 6)
		zigzagNumberHOLDER := zigzagNumber

		platformStartBlockHOLDER := platformStartBlock

		// clear existing blocks
		for {
			for a := 0; a < 4; a++ {
				clearAroundBlock(platformStartBlock)
				platformStartBlock += 2
				platformStartBlock -= (94 * 2)
			}
			platformStartBlock += (94 * 4)
			clearAroundBlock(platformStartBlock)
			for a := 0; a < 3; a++ {
				clearAroundBlock(platformStartBlock)
				platformStartBlock += 2
				platformStartBlock += (94 * 2)
			}
			zigzagNumber--
			if zigzagNumber == 0 {
				break
			}

		}

		zigzagNumber = zigzagNumberHOLDER
		platformStartBlock = platformStartBlockHOLDER

		// draw platform shape
		for {
			for a := 0; a < 4; a++ {

				levelMAP[platformStartBlock] = "floor"
				levelMAP[platformStartBlock+1] = "floor"
				levelMAP[platformStartBlock+94] = "floor"
				levelMAP[platformStartBlock+95] = "floor"

				platformsMAP[platformStartBlock] = platformTypeNameTL
				platformsMAP[platformStartBlock+1] = platformTypeName
				platformsMAP[platformStartBlock+94] = platformTypeName
				platformsMAP[platformStartBlock+95] = platformTypeName

				platformStartBlock += 2
				platformStartBlock -= (94 * 2)

			}
			platformStartBlock += (94 * 4)

			for a := 0; a < 3; a++ {

				levelMAP[platformStartBlock] = "floor"
				levelMAP[platformStartBlock+1] = "floor"
				levelMAP[platformStartBlock+94] = "floor"
				levelMAP[platformStartBlock+95] = "floor"

				platformsMAP[platformStartBlock] = platformTypeNameTL
				platformsMAP[platformStartBlock+1] = platformTypeName
				platformsMAP[platformStartBlock+94] = platformTypeName
				platformsMAP[platformStartBlock+95] = platformTypeName

				platformStartBlock += 2
				platformStartBlock += (94 * 2)

			}
			zigzagNumber--

			if zigzagNumber == 0 {
				break
			}

		}

		platformStartBlock = 3948 + rInt(4, 20)
		platformStartBlock -= (rInt(0, 3) * 94)

		for {
			blockHorizontal := platformStartBlock / 94
			if blockHorizontal < 42 {
				clearAroundBlock(platformStartBlock)
			}
			levelMAP[platformStartBlock] = "floor"
			levelMAP[platformStartBlock+1] = "floor"
			levelMAP[platformStartBlock+94] = "floor"
			levelMAP[platformStartBlock+95] = "floor"

			platformsMAP[platformStartBlock] = platformTypeNameTL
			platformsMAP[platformStartBlock+1] = platformTypeName
			platformsMAP[platformStartBlock+94] = platformTypeName
			platformsMAP[platformStartBlock+95] = platformTypeName

			anotherBlock := flipcoin()

			if anotherBlock {
				platformStartBlock += 2
				blockHorizontal := platformStartBlock / 94

				if blockHorizontal < 42 {
					clearAroundBlock(platformStartBlock)
				}

				levelMAP[platformStartBlock] = "floor"
				levelMAP[platformStartBlock+1] = "floor"
				levelMAP[platformStartBlock+94] = "floor"
				levelMAP[platformStartBlock+95] = "floor"

				platformsMAP[platformStartBlock] = platformTypeNameTL
				platformsMAP[platformStartBlock+1] = platformTypeName
				platformsMAP[platformStartBlock+94] = platformTypeName
				platformsMAP[platformStartBlock+95] = platformTypeName
			}

			platformStartBlock += rInt(4, 15)
			platformStartBlock -= (rInt(1, 5) * 94)

			blockHorizontal = platformStartBlock / 94

			if blockHorizontal < 10 {
				break
			}

		}

	case 13:

		platformType := rInt(1, 42)
		platformTypeName = "platform" + strconv.Itoa(platformType)
		platformTypeNameTL = platformTypeName + "TL"

		platformStartBlock := 3770 + rInt(0, 10)
		platformStartBlock -= (rInt(0, 4) * 94)

		for {

			horizontalLength := rInt(3, 6)

			for a := 0; a < horizontalLength; a++ {
				levelMAP[platformStartBlock] = "floor"
				levelMAP[platformStartBlock+1] = "floor"
				levelMAP[platformStartBlock+94] = "floor"
				levelMAP[platformStartBlock+95] = "floor"

				platformsMAP[platformStartBlock] = platformTypeNameTL
				platformsMAP[platformStartBlock+1] = platformTypeName
				platformsMAP[platformStartBlock+94] = platformTypeName
				platformsMAP[platformStartBlock+95] = platformTypeName

				platformStartBlock += 2

				blockHorizontal := platformStartBlock / 94
				blockVertical := platformStartBlock - (blockHorizontal * 94)

				if blockVertical > 84 {
					break
				}

			}

			for {
				levelMAP[platformStartBlock] = "floor"
				levelMAP[platformStartBlock+1] = "floor"
				levelMAP[platformStartBlock+94] = "floor"
				levelMAP[platformStartBlock+95] = "floor"

				platformsMAP[platformStartBlock] = platformTypeNameTL
				platformsMAP[platformStartBlock+1] = platformTypeName
				platformsMAP[platformStartBlock+94] = platformTypeName
				platformsMAP[platformStartBlock+95] = platformTypeName

				platformStartBlock -= 94 * 2
				blockHorizontal := platformStartBlock / 94

				columnHeight := rInt(12, 25)

				if blockHorizontal < columnHeight {
					break
				}
			}

			horizontalLength = rInt(3, 6)

			for a := 0; a < horizontalLength; a++ {
				levelMAP[platformStartBlock] = "floor"
				levelMAP[platformStartBlock+1] = "floor"
				levelMAP[platformStartBlock+94] = "floor"
				levelMAP[platformStartBlock+95] = "floor"

				platformsMAP[platformStartBlock] = platformTypeNameTL
				platformsMAP[platformStartBlock+1] = platformTypeName
				platformsMAP[platformStartBlock+94] = platformTypeName
				platformsMAP[platformStartBlock+95] = platformTypeName

				platformStartBlock += 2

				blockHorizontal := platformStartBlock / 94
				blockVertical := platformStartBlock - (blockHorizontal * 94)

				if blockVertical > 84 {
					break
				}

			}

			for {
				levelMAP[platformStartBlock] = "floor"
				levelMAP[platformStartBlock+1] = "floor"
				levelMAP[platformStartBlock+94] = "floor"
				levelMAP[platformStartBlock+95] = "floor"

				platformsMAP[platformStartBlock] = platformTypeNameTL
				platformsMAP[platformStartBlock+1] = platformTypeName
				platformsMAP[platformStartBlock+94] = platformTypeName
				platformsMAP[platformStartBlock+95] = platformTypeName

				platformStartBlock += 94 * 2
				blockHorizontal := platformStartBlock / 94

				columnHeight := rInt(25, 40)
				if blockHorizontal > columnHeight {
					break
				}
			}

			horizontalLength = rInt(3, 6)

			for a := 0; a < horizontalLength; a++ {
				levelMAP[platformStartBlock] = "floor"
				levelMAP[platformStartBlock+1] = "floor"
				levelMAP[platformStartBlock+94] = "floor"
				levelMAP[platformStartBlock+95] = "floor"

				platformsMAP[platformStartBlock] = platformTypeNameTL
				platformsMAP[platformStartBlock+1] = platformTypeName
				platformsMAP[platformStartBlock+94] = platformTypeName
				platformsMAP[platformStartBlock+95] = platformTypeName

				platformStartBlock += 2

				blockHorizontal := platformStartBlock / 94
				blockVertical := platformStartBlock - (blockHorizontal * 94)

				if blockVertical > 84 {
					break
				}

			}

			blockHorizontal := platformStartBlock / 94
			blockVertical := platformStartBlock - (blockHorizontal * 94)

			if blockVertical > 84 {
				break
			}

		}

	case 14:

		platformType := rInt(1, 42)
		platformTypeName = "platform" + strconv.Itoa(platformType)
		platformTypeNameTL = platformTypeName + "TL"

		platformStartBlock := 3870 + rInt(0, 5)
		platformStartBlock -= (rInt(0, 4) * 94)

		for {

			for {
				levelMAP[platformStartBlock] = "floor"
				levelMAP[platformStartBlock+1] = "floor"
				levelMAP[platformStartBlock+94] = "floor"
				levelMAP[platformStartBlock+95] = "floor"

				platformsMAP[platformStartBlock] = platformTypeNameTL
				platformsMAP[platformStartBlock+1] = platformTypeName
				platformsMAP[platformStartBlock+94] = platformTypeName
				platformsMAP[platformStartBlock+95] = platformTypeName

				platformStartBlock += 2

				blockHorizontal := platformStartBlock / 94
				blockVertical := platformStartBlock - (blockHorizontal * 94)

				lengthPlatform := rInt(65, 90)

				if blockVertical > lengthPlatform {
					break
				}
			}

			platformStartBlock -= 94 * 2
			platformStartBlock -= 2

			verticalLength := rInt(2, 5)
			for a := 0; a < verticalLength; a++ {
				levelMAP[platformStartBlock] = "floor"
				levelMAP[platformStartBlock+1] = "floor"
				levelMAP[platformStartBlock+94] = "floor"
				levelMAP[platformStartBlock+95] = "floor"

				platformsMAP[platformStartBlock] = platformTypeNameTL
				platformsMAP[platformStartBlock+1] = platformTypeName
				platformsMAP[platformStartBlock+94] = platformTypeName
				platformsMAP[platformStartBlock+95] = platformTypeName

				platformStartBlock -= 94 * 2
				blockHorizontal := platformStartBlock / 94

				if blockHorizontal < 10 {
					break
				}
			}

			for {
				levelMAP[platformStartBlock] = "floor"
				levelMAP[platformStartBlock+1] = "floor"
				levelMAP[platformStartBlock+94] = "floor"
				levelMAP[platformStartBlock+95] = "floor"

				platformsMAP[platformStartBlock] = platformTypeNameTL
				platformsMAP[platformStartBlock+1] = platformTypeName
				platformsMAP[platformStartBlock+94] = platformTypeName
				platformsMAP[platformStartBlock+95] = platformTypeName

				platformStartBlock -= 2

				blockHorizontal := platformStartBlock / 94
				blockVertical := platformStartBlock - (blockHorizontal * 94)

				lengthPlatform := rInt(8, 25)

				if blockVertical < lengthPlatform {
					break
				}
			}

			verticalLength = rInt(2, 5)
			for a := 0; a < verticalLength; a++ {
				levelMAP[platformStartBlock] = "floor"
				levelMAP[platformStartBlock+1] = "floor"
				levelMAP[platformStartBlock+94] = "floor"
				levelMAP[platformStartBlock+95] = "floor"

				platformsMAP[platformStartBlock] = platformTypeNameTL
				platformsMAP[platformStartBlock+1] = platformTypeName
				platformsMAP[platformStartBlock+94] = platformTypeName
				platformsMAP[platformStartBlock+95] = platformTypeName

				platformStartBlock -= 94 * 2
				blockHorizontal := platformStartBlock / 94

				if blockHorizontal < 10 {
					break
				}
			}
			blockHorizontal := platformStartBlock / 94
			if blockHorizontal < 10 {
				break
			}
		}

	case 15:

		platformType := rInt(1, 42)
		platformTypeName = "platform" + strconv.Itoa(platformType)
		platformTypeNameTL = platformTypeName + "TL"

		platformStartBlock := 3870 + rInt(0, 5)
		platformStartBlock -= (rInt(0, 4) * 94)

		for {
			blockHorizontal := platformStartBlock / 94

			if blockHorizontal > 10 {
				horizontalLength := rInt(3, 5)
				for a := 0; a < horizontalLength; a++ {

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += 2

					blockHorizontal := platformStartBlock / 94
					blockVertical := platformStartBlock - (blockHorizontal * 94)

					if blockVertical > 84 {
						break
					}

				}

				verticalLength := rInt(3, 5)
				for a := 0; a < verticalLength; a++ {
					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock -= 94 * 2
					blockHorizontal := platformStartBlock / 94
					if blockHorizontal < 10 {
						break
					}
				}

				horizontalLength = rInt(3, 5)
				for a := 0; a < horizontalLength; a++ {

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock -= 2

					blockHorizontal := platformStartBlock / 94
					blockVertical := platformStartBlock - (blockHorizontal * 94)

					if blockVertical < 10 {
						break
					}

				}

				verticalLength = rInt(3, 5)
				for a := 0; a < verticalLength; a++ {
					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock -= 94 * 2
					blockHorizontal := platformStartBlock / 94
					if blockHorizontal < 10 {
						break
					}
				}

			} else {
				horizontalLength := rInt(12, 19)
				for a := 0; a < horizontalLength; a++ {

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += 2

					blockHorizontal := platformStartBlock / 94
					blockVertical := platformStartBlock - (blockHorizontal * 94)

					if blockVertical > 84 {
						break
					}
				}

				for {
					verticalLength := rInt(3, 5)
					for a := 0; a < verticalLength; a++ {
						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 94 * 2
						blockHorizontal := platformStartBlock / 94
						if blockHorizontal > 36 {
							break
						}
					}

					horizontalLength = rInt(3, 5)
					for a := 0; a < horizontalLength; a++ {

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock -= 2

						blockHorizontal := platformStartBlock / 94
						blockVertical := platformStartBlock - (blockHorizontal * 94)

						if blockVertical < 10 {
							break
						}

					}

					verticalLength = rInt(3, 5)
					for a := 0; a < verticalLength; a++ {
						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 94 * 2
						blockHorizontal := platformStartBlock / 94
						if blockHorizontal > 36 {
							break
						}
					}

					horizontalLength = rInt(3, 5)
					for a := 0; a < horizontalLength; a++ {

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 2

						blockHorizontal := platformStartBlock / 94
						blockVertical := platformStartBlock - (blockHorizontal * 94)

						if blockVertical > 84 {
							break
						}

					}
					blockHorizontal := platformStartBlock / 94

					if blockHorizontal > 36 {
						horizontalLength = rInt(3, 5)
						for a := 0; a < horizontalLength; a++ {

							levelMAP[platformStartBlock] = "floor"
							levelMAP[platformStartBlock+1] = "floor"
							levelMAP[platformStartBlock+94] = "floor"
							levelMAP[platformStartBlock+95] = "floor"

							platformsMAP[platformStartBlock] = platformTypeNameTL
							platformsMAP[platformStartBlock+1] = platformTypeName
							platformsMAP[platformStartBlock+94] = platformTypeName
							platformsMAP[platformStartBlock+95] = platformTypeName

							platformStartBlock += 2

							blockHorizontal := platformStartBlock / 94
							blockVertical := platformStartBlock - (blockHorizontal * 94)

							if blockVertical > 84 {
								break
							}

						}

						break
					}

				}

				blockHorizontal := platformStartBlock / 94
				blockVertical := platformStartBlock - (blockHorizontal * 94)

				if blockVertical > 84 {
					break
				}
			}

		}

	}

	// MARK: add extra level shapes / extra features
	// difficulty setting change this
	if totalLevels > 1 && totalLevels < 13 {
		if easyDiffOn {
			if totalLevels >= 8 {
				extraShapesOn = true
			}
			if totalLevels >= 12 {
				extraFeaturesOn = true
			}
		} else if averageDiffOn {
			if totalLevels >= 4 {
				extraShapesOn = true
			}
			if totalLevels >= 8 {
				extraFeaturesOn = true
			}
		} else if difficultDiffOn {
			if totalLevels >= 2 {
				extraShapesOn = true
			}
			if totalLevels >= 6 {
				extraFeaturesOn = true
			}
		}
	}
	if extraFeaturesOn {
		// falling blocks
		fallingBlocksOn = flipcoin()
		// moving platforms up/down
		movingPlatformsOn = flipcoin()
		// ground blocks
		groundBlocksOn = flipcoin()
		// horizontal moving platforms
		horizplaton = flipcoin()

		if horizplaton && movingPlatformsOn {

			choose := flipcoin()

			if choose {
				horizplaton = true
				movingPlatformsOn = false
			} else {
				horizplaton = false
				movingPlatformsOn = true
			}

		}
	}

	// horizontal platforms

	if horizplaton {

		horizplatblock := 1324
		horizplatL := 8

		horizplatblock += 94 * rInt(2, 20)

		for a := 1; a < horizplatL+1; a++ {
			platname := "movingplatH" + strconv.Itoa(a)
			platformsMAP[horizplatblock] = platname
			levelMAP[horizplatblock] = "floor"
			horizplatblock++
		}

	}

	if extraShapesOn {
		switch platformLayoutType {
		case 4:
			chooseShapeNumber := rInt(1, 6)
			switch chooseShapeNumber {
			case 1:
				// cross shape right
				chooseShape1or2 := flipcoin()
				if chooseShape1or2 {
					shapePosition := rInt(900, 930)
					shapeNumber := rInt(3, 6)
					shapePositionHOLDER := shapePosition
					shapeNumberHOLDER := shapeNumber
					for {

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							clearAroundBlock(shapePosition)
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = shapePositionHOLDER
					shapeNumber = shapeNumberHOLDER

					for {

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							levelMAP[shapePosition] = "floor"
							levelMAP[shapePosition+1] = "floor"
							levelMAP[shapePosition+94] = "floor"
							levelMAP[shapePosition+95] = "floor"

							platformsMAP[shapePosition] = platformTypeNameTL
							platformsMAP[shapePosition+1] = platformTypeName
							platformsMAP[shapePosition+94] = platformTypeName
							platformsMAP[shapePosition+95] = platformTypeName
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}
				} else {

					shapePosition := rInt(860, 890)
					shapeNumber := rInt(3, 6)

					shapePositionHOLDER := shapePosition
					shapeNumberHOLDER := shapeNumber

					for {

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							clearAroundBlock(shapePosition)
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = shapePositionHOLDER
					shapeNumber = shapeNumberHOLDER

					for {

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							levelMAP[shapePosition] = "floor"
							levelMAP[shapePosition+1] = "floor"
							levelMAP[shapePosition+94] = "floor"
							levelMAP[shapePosition+95] = "floor"

							platformsMAP[shapePosition] = platformTypeNameTL
							platformsMAP[shapePosition+1] = platformTypeName
							platformsMAP[shapePosition+94] = platformTypeName
							platformsMAP[shapePosition+95] = platformTypeName
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = rInt(900, 930)
					shapeNumber = rInt(3, 6)

					shapePositionHOLDER = shapePosition
					shapeNumberHOLDER = shapeNumber

					for {

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							clearAroundBlock(shapePosition)
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = shapePositionHOLDER
					shapeNumber = shapeNumberHOLDER

					for {

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							levelMAP[shapePosition] = "floor"
							levelMAP[shapePosition+1] = "floor"
							levelMAP[shapePosition+94] = "floor"
							levelMAP[shapePosition+95] = "floor"

							platformsMAP[shapePosition] = platformTypeNameTL
							platformsMAP[shapePosition+1] = platformTypeName
							platformsMAP[shapePosition+94] = platformTypeName
							platformsMAP[shapePosition+95] = platformTypeName
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

				}

			case 2:

				// zigzag top
				platformStartBlock := rInt(1136, 1156)
				platformStartBlock += (rInt(0, 11) * 94)

				zigzagNumber := rInt(2, 6)
				zigzagNumberHOLDER := zigzagNumber

				platformStartBlockHOLDER := platformStartBlock

				// clear existing blocks
				for {
					for a := 0; a < 4; a++ {
						clearAroundBlock(platformStartBlock)
						platformStartBlock += 2
						platformStartBlock -= (94 * 2)
					}
					platformStartBlock += (94 * 4)
					clearAroundBlock(platformStartBlock)
					for a := 0; a < 3; a++ {
						clearAroundBlock(platformStartBlock)
						platformStartBlock += 2
						platformStartBlock += (94 * 2)
					}
					zigzagNumber--
					if zigzagNumber == 0 {
						break
					}

				}

				zigzagNumber = zigzagNumberHOLDER
				platformStartBlock = platformStartBlockHOLDER

				// draw platform shape
				for {
					for a := 0; a < 4; a++ {

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 2
						platformStartBlock -= (94 * 2)

					}
					platformStartBlock += (94 * 4)

					for a := 0; a < 3; a++ {

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 2
						platformStartBlock += (94 * 2)

					}
					zigzagNumber--

					if zigzagNumber == 0 {
						break
					}

				}

			case 3:

				// 4 block square right

				platformStartBlock := rInt(706, 741)
				platformStartBlock += (rInt(0, 4) * 94)

				// check surrounding blocks
				for {
					if checkAroundPlatformBlock(platformStartBlock) {
						platformStartBlock += 4
						platformStartBlock += (94 * 2)
					}
					if checkAroundPlatformBlock(platformStartBlock) == false {
						break
					}
				}

				blockNumber := rInt(2, 6)
				platformStartBlockHOLDER := platformStartBlock

				// clear existing blocks
				for a := 0; a < blockNumber; a++ {

					clearAroundBlock(platformStartBlock)
					platformStartBlock += 2
					clearAroundBlock(platformStartBlock)

					platformStartBlock += (93 * 2)
					clearAroundBlock(platformStartBlock)

					platformStartBlock += 2
					clearAroundBlock(platformStartBlock)

					platformStartBlock -= 2
					clearAroundBlock(platformStartBlock)

					platformStartBlock += (94 * 6)

					if platformStartBlock > 3666 {
						break
					}

					clearAroundBlock(platformStartBlock)

				}

				// draw platform shape
				platformStartBlock = platformStartBlockHOLDER

				for a := 0; a < blockNumber; a++ {

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += 2

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += (93 * 2)

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += 2

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock -= 2

					platformStartBlock += (94 * 6)

					if platformStartBlock > 3854 {
						break
					}

				}

			case 4:
				// 4 block square right with zigzag

				platformStartBlock := rInt(706, 741)
				platformStartBlock += (rInt(0, 4) * 94)

				// check surrounding blocks
				for {
					if checkAroundPlatformBlock(platformStartBlock) {
						platformStartBlock += 4
						platformStartBlock += (94 * 2)
					}
					if checkAroundPlatformBlock(platformStartBlock) == false {
						break
					}
				}

				blockNumber := rInt(2, 6)
				platformStartBlockHOLDER := platformStartBlock

				// clear existing blocks
				for a := 0; a < blockNumber; a++ {

					clearAroundBlock(platformStartBlock)
					platformStartBlock += 2
					clearAroundBlock(platformStartBlock)

					platformStartBlock += (93 * 2)
					clearAroundBlock(platformStartBlock)

					platformStartBlock += 2
					clearAroundBlock(platformStartBlock)

					platformStartBlock -= 2
					clearAroundBlock(platformStartBlock)

					platformStartBlock += (94 * 6)

					if platformStartBlock > 3666 {
						break
					}

					clearAroundBlock(platformStartBlock)

				}

				// draw platform shape
				platformStartBlock = platformStartBlockHOLDER

				for a := 0; a < blockNumber; a++ {

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += 2

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += (93 * 2)

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += 2

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock -= 2

					platformStartBlock += (94 * 6)

					if platformStartBlock > 3854 {
						break
					}

				}

				platformStartBlock = rInt(1136, 1156)
				platformStartBlock += (rInt(0, 11) * 94)

				zigzagNumber := rInt(2, 6)
				zigzagNumberHOLDER := zigzagNumber

				platformStartBlockHOLDER = platformStartBlock

				// clear existing blocks
				for {
					for a := 0; a < 4; a++ {
						clearAroundBlock(platformStartBlock)
						platformStartBlock += 2
						platformStartBlock -= (94 * 2)
					}
					platformStartBlock += (94 * 4)
					clearAroundBlock(platformStartBlock)
					for a := 0; a < 3; a++ {
						clearAroundBlock(platformStartBlock)
						platformStartBlock += 2
						platformStartBlock += (94 * 2)
					}
					zigzagNumber--
					if zigzagNumber == 0 {
						break
					}

				}

				zigzagNumber = zigzagNumberHOLDER
				platformStartBlock = platformStartBlockHOLDER

				// draw platform shape
				for {
					for a := 0; a < 4; a++ {

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 2
						platformStartBlock -= (94 * 2)

					}
					platformStartBlock += (94 * 4)

					for a := 0; a < 3; a++ {

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 2
						platformStartBlock += (94 * 2)

					}
					zigzagNumber--

					if zigzagNumber == 0 {
						break
					}

				}

			case 5: // 2X right curl

				// first curl
				platformStartBlock := rInt(854, 886)
				platformStartBlock += (rInt(0, 25) * 94)

				platformStartBlockHOLDER := platformStartBlock

				lengthChange := 3

				for a := 0; a < lengthChange; a++ {
					platformStartBlock++
					clearAroundBlock(platformStartBlock)
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				clearAroundBlock(platformStartBlock)

				for a := 0; a < lengthChange; a++ {
					platformStartBlock += (94 * 2)
					clearAroundBlock(platformStartBlock)
				}

				platformStartBlock = platformStartBlockHOLDER
				lengthChange = 3

				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock++
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock-a] = "floor"
					levelMAP[platformStartBlock-(a+1)] = "floor"
					levelMAP[platformStartBlock-(a+94)] = "floor"
					levelMAP[platformStartBlock-(a+95)] = "floor"

					platformsMAP[platformStartBlock-a] = platformTypeNameTL
					platformsMAP[platformStartBlock-(a+1)] = platformTypeName
					platformsMAP[platformStartBlock-(a+94)] = platformTypeName
					platformsMAP[platformStartBlock-(a+95)] = platformTypeName

					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += (94 * 2)
				}

				// second curl

				platformStartBlock = rInt(1362, 1387)
				platformStartBlock += (rInt(0, 15) * 94)

				platformStartBlockHOLDER = platformStartBlock

				lengthChange = 3

				for a := 0; a < lengthChange; a++ {
					platformStartBlock++
					clearAroundBlock(platformStartBlock)
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				clearAroundBlock(platformStartBlock)

				for a := 0; a < lengthChange; a++ {
					platformStartBlock += (94 * 2)
					clearAroundBlock(platformStartBlock)
				}

				platformStartBlock = platformStartBlockHOLDER
				lengthChange = 3

				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock++
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock-a] = "floor"
					levelMAP[platformStartBlock-(a+1)] = "floor"
					levelMAP[platformStartBlock-(a+94)] = "floor"
					levelMAP[platformStartBlock-(a+95)] = "floor"

					platformsMAP[platformStartBlock-a] = platformTypeNameTL
					platformsMAP[platformStartBlock-(a+1)] = platformTypeName
					platformsMAP[platformStartBlock-(a+94)] = platformTypeName
					platformsMAP[platformStartBlock-(a+95)] = platformTypeName

					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += (94 * 2)
				}

				// additional shape
				chooseNextShape := flipcoin()

				if chooseNextShape {

					// zigzag top
					platformStartBlock = rInt(1136, 1156)
					platformStartBlock += (rInt(6, 16) * 94)

					zigzagNumber := rInt(2, 6)
					zigzagNumberHOLDER := zigzagNumber

					platformStartBlockHOLDER := platformStartBlock

					// clear existing blocks
					for {
						for a := 0; a < 4; a++ {
							clearAroundBlock(platformStartBlock)
							platformStartBlock += 2
							platformStartBlock -= (94 * 2)
						}
						platformStartBlock += (94 * 4)
						clearAroundBlock(platformStartBlock)
						for a := 0; a < 3; a++ {
							clearAroundBlock(platformStartBlock)
							platformStartBlock += 2
							platformStartBlock += (94 * 2)
						}
						zigzagNumber--
						if zigzagNumber == 0 {
							break
						}

					}

					zigzagNumber = zigzagNumberHOLDER
					platformStartBlock = platformStartBlockHOLDER

					// draw platform shape
					for {
						for a := 0; a < 4; a++ {

							levelMAP[platformStartBlock] = "floor"
							levelMAP[platformStartBlock+1] = "floor"
							levelMAP[platformStartBlock+94] = "floor"
							levelMAP[platformStartBlock+95] = "floor"

							platformsMAP[platformStartBlock] = platformTypeNameTL
							platformsMAP[platformStartBlock+1] = platformTypeName
							platformsMAP[platformStartBlock+94] = platformTypeName
							platformsMAP[platformStartBlock+95] = platformTypeName

							platformStartBlock += 2
							platformStartBlock -= (94 * 2)

						}
						platformStartBlock += (94 * 4)

						for a := 0; a < 3; a++ {

							levelMAP[platformStartBlock] = "floor"
							levelMAP[platformStartBlock+1] = "floor"
							levelMAP[platformStartBlock+94] = "floor"
							levelMAP[platformStartBlock+95] = "floor"

							platformsMAP[platformStartBlock] = platformTypeNameTL
							platformsMAP[platformStartBlock+1] = platformTypeName
							platformsMAP[platformStartBlock+94] = platformTypeName
							platformsMAP[platformStartBlock+95] = platformTypeName

							platformStartBlock += 2
							platformStartBlock += (94 * 2)

						}
						zigzagNumber--

						if zigzagNumber == 0 {
							break
						}

					}

				}

			} // end switch case 4

		case 5:

			chooseShapeNumber := rInt(1, 6)

			switch chooseShapeNumber {

			case 1:

				// cross shape right

				chooseShape1or2 := flipcoin()

				if chooseShape1or2 {

					shapePosition := rInt(900, 930)
					shapeNumber := rInt(3, 6)

					shapePositionHOLDER := shapePosition
					shapeNumberHOLDER := shapeNumber

					for {

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							clearAroundBlock(shapePosition)
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = shapePositionHOLDER
					shapeNumber = shapeNumberHOLDER

					for {

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							levelMAP[shapePosition] = "floor"
							levelMAP[shapePosition+1] = "floor"
							levelMAP[shapePosition+94] = "floor"
							levelMAP[shapePosition+95] = "floor"

							platformsMAP[shapePosition] = platformTypeNameTL
							platformsMAP[shapePosition+1] = platformTypeName
							platformsMAP[shapePosition+94] = platformTypeName
							platformsMAP[shapePosition+95] = platformTypeName
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}
				} else {

					shapePosition := rInt(860, 890)
					shapeNumber := rInt(3, 6)

					shapePositionHOLDER := shapePosition
					shapeNumberHOLDER := shapeNumber

					for {

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							clearAroundBlock(shapePosition)
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = shapePositionHOLDER
					shapeNumber = shapeNumberHOLDER

					for {

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							levelMAP[shapePosition] = "floor"
							levelMAP[shapePosition+1] = "floor"
							levelMAP[shapePosition+94] = "floor"
							levelMAP[shapePosition+95] = "floor"

							platformsMAP[shapePosition] = platformTypeNameTL
							platformsMAP[shapePosition+1] = platformTypeName
							platformsMAP[shapePosition+94] = platformTypeName
							platformsMAP[shapePosition+95] = platformTypeName
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = rInt(900, 930)
					shapeNumber = rInt(3, 6)

					shapePositionHOLDER = shapePosition
					shapeNumberHOLDER = shapeNumber

					for {

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							clearAroundBlock(shapePosition)
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = shapePositionHOLDER
					shapeNumber = shapeNumberHOLDER

					for {

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							levelMAP[shapePosition] = "floor"
							levelMAP[shapePosition+1] = "floor"
							levelMAP[shapePosition+94] = "floor"
							levelMAP[shapePosition+95] = "floor"

							platformsMAP[shapePosition] = platformTypeNameTL
							platformsMAP[shapePosition+1] = platformTypeName
							platformsMAP[shapePosition+94] = platformTypeName
							platformsMAP[shapePosition+95] = platformTypeName
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

				}

			case 2:

				// zigzag top
				platformStartBlock := rInt(1136, 1156)
				platformStartBlock += (rInt(9, 18) * 94)

				zigzagNumber := rInt(2, 6)
				zigzagNumberHOLDER := zigzagNumber

				platformStartBlockHOLDER := platformStartBlock

				// clear existing blocks
				for {
					for a := 0; a < 4; a++ {
						clearAroundBlock(platformStartBlock)
						platformStartBlock += 2
						platformStartBlock -= (94 * 2)
					}
					platformStartBlock += (94 * 4)
					clearAroundBlock(platformStartBlock)
					for a := 0; a < 3; a++ {
						clearAroundBlock(platformStartBlock)
						platformStartBlock += 2
						platformStartBlock += (94 * 2)
					}
					zigzagNumber--
					if zigzagNumber == 0 {
						break
					}

				}

				zigzagNumber = zigzagNumberHOLDER
				platformStartBlock = platformStartBlockHOLDER

				// draw platform shape
				for {
					for a := 0; a < 4; a++ {

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 2
						platformStartBlock -= (94 * 2)

					}
					platformStartBlock += (94 * 4)

					for a := 0; a < 3; a++ {

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 2
						platformStartBlock += (94 * 2)

					}
					zigzagNumber--

					if zigzagNumber == 0 {
						break
					}

				}

			case 3:

				// 4 block square right

				platformStartBlock := rInt(690, 710)
				platformStartBlock += (rInt(0, 4) * 94)

				for b := 0; b < 2; b++ {

					blockNumber := rInt(2, 6)
					platformStartBlockHOLDER := platformStartBlock

					// clear existing blocks
					for a := 0; a < blockNumber; a++ {

						clearAroundBlock(platformStartBlock)
						platformStartBlock += 2
						clearAroundBlock(platformStartBlock)

						platformStartBlock += (93 * 2)
						clearAroundBlock(platformStartBlock)

						platformStartBlock += 2
						clearAroundBlock(platformStartBlock)

						platformStartBlock -= 2
						clearAroundBlock(platformStartBlock)

						platformStartBlock += (94 * 6)

						if platformStartBlock > 3666 {
							break
						}

						clearAroundBlock(platformStartBlock)

					}

					// draw platform shape
					platformStartBlock = platformStartBlockHOLDER

					for a := 0; a < blockNumber; a++ {

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 2

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += (93 * 2)

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 2

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock -= 2

						platformStartBlock += (94 * 6)

						if platformStartBlock > 3854 {
							break
						}

					}

					platformStartBlock = platformStartBlockHOLDER
					platformStartBlock += rInt(20, 35)
					platformStartBlock += (rInt(-2, 3) * 94)

				}

			case 4: // 2X right curl

				// first curl
				platformStartBlock := rInt(854, 886)
				platformStartBlock += (rInt(0, 25) * 94)

				platformStartBlockHOLDER := platformStartBlock

				lengthChange := 3

				for a := 0; a < lengthChange; a++ {
					platformStartBlock++
					clearAroundBlock(platformStartBlock)
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				clearAroundBlock(platformStartBlock)

				for a := 0; a < lengthChange; a++ {
					platformStartBlock += (94 * 2)
					clearAroundBlock(platformStartBlock)
				}

				platformStartBlock = platformStartBlockHOLDER
				lengthChange = 3

				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock++
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock-a] = "floor"
					levelMAP[platformStartBlock-(a+1)] = "floor"
					levelMAP[platformStartBlock-(a+94)] = "floor"
					levelMAP[platformStartBlock-(a+95)] = "floor"

					platformsMAP[platformStartBlock-a] = platformTypeNameTL
					platformsMAP[platformStartBlock-(a+1)] = platformTypeName
					platformsMAP[platformStartBlock-(a+94)] = platformTypeName
					platformsMAP[platformStartBlock-(a+95)] = platformTypeName

					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += (94 * 2)
				}

				// second curl

				platformStartBlock = rInt(1362, 1387)
				platformStartBlock += (rInt(0, 15) * 94)

				platformStartBlockHOLDER = platformStartBlock

				lengthChange = 3

				for a := 0; a < lengthChange; a++ {
					platformStartBlock++
					clearAroundBlock(platformStartBlock)
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				clearAroundBlock(platformStartBlock)

				for a := 0; a < lengthChange; a++ {
					platformStartBlock += (94 * 2)
					clearAroundBlock(platformStartBlock)
				}

				platformStartBlock = platformStartBlockHOLDER
				lengthChange = 3

				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock++
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock-a] = "floor"
					levelMAP[platformStartBlock-(a+1)] = "floor"
					levelMAP[platformStartBlock-(a+94)] = "floor"
					levelMAP[platformStartBlock-(a+95)] = "floor"

					platformsMAP[platformStartBlock-a] = platformTypeNameTL
					platformsMAP[platformStartBlock-(a+1)] = platformTypeName
					platformsMAP[platformStartBlock-(a+94)] = platformTypeName
					platformsMAP[platformStartBlock-(a+95)] = platformTypeName

					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += (94 * 2)
				}

			} // end switch case 5

		case 6:

			chooseShapeNumber := rInt(1, 4)

			switch chooseShapeNumber {

			case 1:
				// zigzag top
				platformStartBlock := rInt(1136, 1156)
				platformStartBlock += (rInt(9, 18) * 94)

				zigzagNumber := rInt(2, 6)
				zigzagNumberHOLDER := zigzagNumber

				platformStartBlockHOLDER := platformStartBlock

				// clear existing blocks
				for {
					for a := 0; a < 4; a++ {
						clearAroundBlock(platformStartBlock)
						platformStartBlock += 2
						platformStartBlock -= (94 * 2)
					}
					platformStartBlock += (94 * 4)
					clearAroundBlock(platformStartBlock)
					for a := 0; a < 3; a++ {
						clearAroundBlock(platformStartBlock)
						platformStartBlock += 2
						platformStartBlock += (94 * 2)
					}
					zigzagNumber--
					if zigzagNumber == 0 {
						break
					}

				}

				zigzagNumber = zigzagNumberHOLDER
				platformStartBlock = platformStartBlockHOLDER

				// draw platform shape
				for {
					for a := 0; a < 4; a++ {

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 2
						platformStartBlock -= (94 * 2)

					}
					platformStartBlock += (94 * 4)

					for a := 0; a < 3; a++ {

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 2
						platformStartBlock += (94 * 2)

					}
					zigzagNumber--

					if zigzagNumber == 0 {
						break
					}
				}

			case 2:

				// first curl
				platformStartBlock := rInt(854, 886)
				platformStartBlock += (rInt(0, 25) * 94)

				platformStartBlockHOLDER := platformStartBlock

				lengthChange := 3

				for a := 0; a < lengthChange; a++ {
					platformStartBlock++
					clearAroundBlock(platformStartBlock)
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				clearAroundBlock(platformStartBlock)

				for a := 0; a < lengthChange; a++ {
					platformStartBlock += (94 * 2)
					clearAroundBlock(platformStartBlock)
				}

				platformStartBlock = platformStartBlockHOLDER
				lengthChange = 3

				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock++
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock-a] = "floor"
					levelMAP[platformStartBlock-(a+1)] = "floor"
					levelMAP[platformStartBlock-(a+94)] = "floor"
					levelMAP[platformStartBlock-(a+95)] = "floor"

					platformsMAP[platformStartBlock-a] = platformTypeNameTL
					platformsMAP[platformStartBlock-(a+1)] = platformTypeName
					platformsMAP[platformStartBlock-(a+94)] = platformTypeName
					platformsMAP[platformStartBlock-(a+95)] = platformTypeName

					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += (94 * 2)
				}

				// second curl

				platformStartBlock = rInt(1362, 1387)
				platformStartBlock += (rInt(0, 15) * 94)

				platformStartBlockHOLDER = platformStartBlock

				lengthChange = 3

				for a := 0; a < lengthChange; a++ {
					platformStartBlock++
					clearAroundBlock(platformStartBlock)
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				clearAroundBlock(platformStartBlock)

				for a := 0; a < lengthChange; a++ {
					platformStartBlock += (94 * 2)
					clearAroundBlock(platformStartBlock)
				}

				platformStartBlock = platformStartBlockHOLDER
				lengthChange = 3

				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock++
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock-a] = "floor"
					levelMAP[platformStartBlock-(a+1)] = "floor"
					levelMAP[platformStartBlock-(a+94)] = "floor"
					levelMAP[platformStartBlock-(a+95)] = "floor"

					platformsMAP[platformStartBlock-a] = platformTypeNameTL
					platformsMAP[platformStartBlock-(a+1)] = platformTypeName
					platformsMAP[platformStartBlock-(a+94)] = platformTypeName
					platformsMAP[platformStartBlock-(a+95)] = platformTypeName

					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += (94 * 2)
				}

			case 3:
				// horizontal line top with gap

				platformStartBlock := 940
				platformStartBlock += (rInt(-3, 4) * 94)
				platformStartBlockHOLDER := platformStartBlock

				for a := 0; a < 46; a++ {
					clearAroundBlock(platformStartBlock)
					platformStartBlock++
				}

				platformStartBlock = platformStartBlockHOLDER

				for a := 0; a < 46; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock++
				}

				platformStartBlock = platformStartBlockHOLDER
				platformStartBlock += rInt(30, 50)
				gapLength := rInt(6, 15)

				for a := 0; a < gapLength; a++ {

					levelMAP[platformStartBlock+a] = ""
					levelMAP[platformStartBlock+(a+1)] = ""
					levelMAP[platformStartBlock+(a+94)] = ""
					levelMAP[platformStartBlock+(a+95)] = ""

					platformsMAP[platformStartBlock+a] = ""
					platformsMAP[platformStartBlock+(a+1)] = ""
					platformsMAP[platformStartBlock+(a+94)] = ""
					platformsMAP[platformStartBlock+(a+95)] = ""

					platformStartBlock++
				}

			} // end switch case 6
		case 7:

			chooseShapeNumber := rInt(1, 4)

			switch chooseShapeNumber {

			case 1:
				// horizontal line top with gap

				platformStartBlock := 940
				platformStartBlock += (rInt(-3, 4) * 94)
				platformStartBlockHOLDER := platformStartBlock

				for a := 0; a < 46; a++ {
					clearAroundBlock(platformStartBlock)
					platformStartBlock++
				}

				platformStartBlock = platformStartBlockHOLDER

				for a := 0; a < 46; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock++
				}

				platformStartBlock = platformStartBlockHOLDER
				platformStartBlock += rInt(30, 50)
				gapLength := rInt(6, 15)

				for a := 0; a < gapLength; a++ {

					levelMAP[platformStartBlock+a] = ""
					levelMAP[platformStartBlock+(a+1)] = ""
					levelMAP[platformStartBlock+(a+94)] = ""
					levelMAP[platformStartBlock+(a+95)] = ""

					platformsMAP[platformStartBlock+a] = ""
					platformsMAP[platformStartBlock+(a+1)] = ""
					platformsMAP[platformStartBlock+(a+94)] = ""
					platformsMAP[platformStartBlock+(a+95)] = ""

					platformStartBlock++
				}

			case 2:

				// first curl
				platformStartBlock := rInt(854, 886)
				platformStartBlock += (rInt(0, 25) * 94)

				platformStartBlockHOLDER := platformStartBlock

				lengthChange := 3

				for a := 0; a < lengthChange; a++ {
					platformStartBlock++
					clearAroundBlock(platformStartBlock)
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				clearAroundBlock(platformStartBlock)

				for a := 0; a < lengthChange; a++ {
					platformStartBlock += (94 * 2)
					clearAroundBlock(platformStartBlock)
				}

				platformStartBlock = platformStartBlockHOLDER
				lengthChange = 3

				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock++
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock-a] = "floor"
					levelMAP[platformStartBlock-(a+1)] = "floor"
					levelMAP[platformStartBlock-(a+94)] = "floor"
					levelMAP[platformStartBlock-(a+95)] = "floor"

					platformsMAP[platformStartBlock-a] = platformTypeNameTL
					platformsMAP[platformStartBlock-(a+1)] = platformTypeName
					platformsMAP[platformStartBlock-(a+94)] = platformTypeName
					platformsMAP[platformStartBlock-(a+95)] = platformTypeName

					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += (94 * 2)
				}

				// second curl

				platformStartBlock = rInt(1362, 1387)
				platformStartBlock += (rInt(0, 15) * 94)

				platformStartBlockHOLDER = platformStartBlock

				lengthChange = 3

				for a := 0; a < lengthChange; a++ {
					platformStartBlock++
					clearAroundBlock(platformStartBlock)
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				clearAroundBlock(platformStartBlock)

				for a := 0; a < lengthChange; a++ {
					platformStartBlock += (94 * 2)
					clearAroundBlock(platformStartBlock)
				}

				platformStartBlock = platformStartBlockHOLDER
				lengthChange = 3

				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock++
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock-a] = "floor"
					levelMAP[platformStartBlock-(a+1)] = "floor"
					levelMAP[platformStartBlock-(a+94)] = "floor"
					levelMAP[platformStartBlock-(a+95)] = "floor"

					platformsMAP[platformStartBlock-a] = platformTypeNameTL
					platformsMAP[platformStartBlock-(a+1)] = platformTypeName
					platformsMAP[platformStartBlock-(a+94)] = platformTypeName
					platformsMAP[platformStartBlock-(a+95)] = platformTypeName

					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += (94 * 2)
				}

			case 3:
				// zigzag top
				platformStartBlock := rInt(1136, 1156)
				platformStartBlock += (rInt(9, 18) * 94)

				zigzagNumber := rInt(2, 6)
				zigzagNumberHOLDER := zigzagNumber

				platformStartBlockHOLDER := platformStartBlock

				// clear existing blocks
				for {
					for a := 0; a < 4; a++ {
						clearAroundBlock(platformStartBlock)
						platformStartBlock += 2
						platformStartBlock -= (94 * 2)
					}
					platformStartBlock += (94 * 4)
					clearAroundBlock(platformStartBlock)
					for a := 0; a < 3; a++ {
						clearAroundBlock(platformStartBlock)
						platformStartBlock += 2
						platformStartBlock += (94 * 2)
					}
					zigzagNumber--
					if zigzagNumber == 0 {
						break
					}

				}

				zigzagNumber = zigzagNumberHOLDER
				platformStartBlock = platformStartBlockHOLDER

				// draw platform shape
				for {
					for a := 0; a < 4; a++ {

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 2
						platformStartBlock -= (94 * 2)

					}
					platformStartBlock += (94 * 4)

					for a := 0; a < 3; a++ {

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 2
						platformStartBlock += (94 * 2)

					}
					zigzagNumber--

					if zigzagNumber == 0 {
						break
					}
				}

			case 4:

				// cross shape right

				chooseShape1or2 := flipcoin()

				if chooseShape1or2 {

					shapePosition := rInt(900, 930)
					shapeNumber := rInt(3, 6)

					shapePositionHOLDER := shapePosition
					shapeNumberHOLDER := shapeNumber

					for {

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							clearAroundBlock(shapePosition)
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = shapePositionHOLDER
					shapeNumber = shapeNumberHOLDER

					for {

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							levelMAP[shapePosition] = "floor"
							levelMAP[shapePosition+1] = "floor"
							levelMAP[shapePosition+94] = "floor"
							levelMAP[shapePosition+95] = "floor"

							platformsMAP[shapePosition] = platformTypeNameTL
							platformsMAP[shapePosition+1] = platformTypeName
							platformsMAP[shapePosition+94] = platformTypeName
							platformsMAP[shapePosition+95] = platformTypeName
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}
				} else {

					shapePosition := rInt(860, 890)
					shapeNumber := rInt(3, 6)

					shapePositionHOLDER := shapePosition
					shapeNumberHOLDER := shapeNumber

					for {

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							clearAroundBlock(shapePosition)
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = shapePositionHOLDER
					shapeNumber = shapeNumberHOLDER

					for {

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							levelMAP[shapePosition] = "floor"
							levelMAP[shapePosition+1] = "floor"
							levelMAP[shapePosition+94] = "floor"
							levelMAP[shapePosition+95] = "floor"

							platformsMAP[shapePosition] = platformTypeNameTL
							platformsMAP[shapePosition+1] = platformTypeName
							platformsMAP[shapePosition+94] = platformTypeName
							platformsMAP[shapePosition+95] = platformTypeName
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = rInt(900, 930)
					shapeNumber = rInt(3, 6)

					shapePositionHOLDER = shapePosition
					shapeNumberHOLDER = shapeNumber

					for {

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							clearAroundBlock(shapePosition)
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = shapePositionHOLDER
					shapeNumber = shapeNumberHOLDER

					for {

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							levelMAP[shapePosition] = "floor"
							levelMAP[shapePosition+1] = "floor"
							levelMAP[shapePosition+94] = "floor"
							levelMAP[shapePosition+95] = "floor"

							platformsMAP[shapePosition] = platformTypeNameTL
							platformsMAP[shapePosition+1] = platformTypeName
							platformsMAP[shapePosition+94] = platformTypeName
							platformsMAP[shapePosition+95] = platformTypeName
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

				}

			} // end switch case 7

		case 8:

			chooseShapeNumber := rInt(1, 4)

			switch chooseShapeNumber {

			case 1:
				// horizontal line top with gap

				platformStartBlock := 940
				platformStartBlock += (rInt(-3, 4) * 94)
				platformStartBlockHOLDER := platformStartBlock

				for a := 0; a < 46; a++ {
					clearAroundBlock(platformStartBlock)
					platformStartBlock++
				}
				platformStartBlock = platformStartBlockHOLDER

				for a := 0; a < 46; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock++
				}

				platformStartBlock = platformStartBlockHOLDER
				platformStartBlock += rInt(30, 50)
				gapLength := rInt(6, 15)

				for a := 0; a < gapLength; a++ {

					levelMAP[platformStartBlock+a] = ""
					levelMAP[platformStartBlock+(a+1)] = ""
					levelMAP[platformStartBlock+(a+94)] = ""
					levelMAP[platformStartBlock+(a+95)] = ""

					platformsMAP[platformStartBlock+a] = ""
					platformsMAP[platformStartBlock+(a+1)] = ""
					platformsMAP[platformStartBlock+(a+94)] = ""
					platformsMAP[platformStartBlock+(a+95)] = ""

					platformStartBlock++
				}

			case 2:

				// first curl
				platformStartBlock := rInt(854, 886)
				platformStartBlock += (rInt(0, 25) * 94)

				platformStartBlockHOLDER := platformStartBlock

				lengthChange := 3

				for a := 0; a < lengthChange; a++ {
					platformStartBlock++
					clearAroundBlock(platformStartBlock)
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				clearAroundBlock(platformStartBlock)

				for a := 0; a < lengthChange; a++ {
					platformStartBlock += (94 * 2)
					clearAroundBlock(platformStartBlock)
				}

				platformStartBlock = platformStartBlockHOLDER
				lengthChange = 3

				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock++
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock-a] = "floor"
					levelMAP[platformStartBlock-(a+1)] = "floor"
					levelMAP[platformStartBlock-(a+94)] = "floor"
					levelMAP[platformStartBlock-(a+95)] = "floor"

					platformsMAP[platformStartBlock-a] = platformTypeNameTL
					platformsMAP[platformStartBlock-(a+1)] = platformTypeName
					platformsMAP[platformStartBlock-(a+94)] = platformTypeName
					platformsMAP[platformStartBlock-(a+95)] = platformTypeName

					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += (94 * 2)
				}

				// second curl

				platformStartBlock = rInt(1362, 1387)
				platformStartBlock += (rInt(0, 15) * 94)

				platformStartBlockHOLDER = platformStartBlock

				lengthChange = 3

				for a := 0; a < lengthChange; a++ {
					platformStartBlock++
					clearAroundBlock(platformStartBlock)
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				clearAroundBlock(platformStartBlock)

				for a := 0; a < lengthChange; a++ {
					platformStartBlock += (94 * 2)
					clearAroundBlock(platformStartBlock)
				}

				platformStartBlock = platformStartBlockHOLDER
				lengthChange = 3

				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock++
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock-a] = "floor"
					levelMAP[platformStartBlock-(a+1)] = "floor"
					levelMAP[platformStartBlock-(a+94)] = "floor"
					levelMAP[platformStartBlock-(a+95)] = "floor"

					platformsMAP[platformStartBlock-a] = platformTypeNameTL
					platformsMAP[platformStartBlock-(a+1)] = platformTypeName
					platformsMAP[platformStartBlock-(a+94)] = platformTypeName
					platformsMAP[platformStartBlock-(a+95)] = platformTypeName

					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += (94 * 2)
				}

			case 3:
				// zigzag top
				platformStartBlock := rInt(1136, 1156)
				platformStartBlock += (rInt(9, 18) * 94)

				zigzagNumber := rInt(2, 6)
				zigzagNumberHOLDER := zigzagNumber

				platformStartBlockHOLDER := platformStartBlock

				// clear existing blocks
				for {
					for a := 0; a < 4; a++ {
						clearAroundBlock(platformStartBlock)
						platformStartBlock += 2
						platformStartBlock -= (94 * 2)
					}
					platformStartBlock += (94 * 4)
					clearAroundBlock(platformStartBlock)
					for a := 0; a < 3; a++ {
						clearAroundBlock(platformStartBlock)
						platformStartBlock += 2
						platformStartBlock += (94 * 2)
					}
					zigzagNumber--
					if zigzagNumber == 0 {
						break
					}

				}

				zigzagNumber = zigzagNumberHOLDER
				platformStartBlock = platformStartBlockHOLDER

				// draw platform shape
				for {
					for a := 0; a < 4; a++ {

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 2
						platformStartBlock -= (94 * 2)

					}
					platformStartBlock += (94 * 4)

					for a := 0; a < 3; a++ {

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 2
						platformStartBlock += (94 * 2)

					}
					zigzagNumber--

					if zigzagNumber == 0 {
						break
					}
				}

			case 4:

				// cross shape right

				chooseShape1or2 := flipcoin()

				if chooseShape1or2 {

					shapePosition := rInt(900, 930)
					shapeNumber := rInt(3, 6)

					shapePositionHOLDER := shapePosition
					shapeNumberHOLDER := shapeNumber

					for {

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							clearAroundBlock(shapePosition)
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = shapePositionHOLDER
					shapeNumber = shapeNumberHOLDER

					for {

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							levelMAP[shapePosition] = "floor"
							levelMAP[shapePosition+1] = "floor"
							levelMAP[shapePosition+94] = "floor"
							levelMAP[shapePosition+95] = "floor"

							platformsMAP[shapePosition] = platformTypeNameTL
							platformsMAP[shapePosition+1] = platformTypeName
							platformsMAP[shapePosition+94] = platformTypeName
							platformsMAP[shapePosition+95] = platformTypeName
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}
				} else {

					shapePosition := rInt(860, 890)
					shapeNumber := rInt(3, 6)

					shapePositionHOLDER := shapePosition
					shapeNumberHOLDER := shapeNumber

					for {

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							clearAroundBlock(shapePosition)
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = shapePositionHOLDER
					shapeNumber = shapeNumberHOLDER

					for {

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							levelMAP[shapePosition] = "floor"
							levelMAP[shapePosition+1] = "floor"
							levelMAP[shapePosition+94] = "floor"
							levelMAP[shapePosition+95] = "floor"

							platformsMAP[shapePosition] = platformTypeNameTL
							platformsMAP[shapePosition+1] = platformTypeName
							platformsMAP[shapePosition+94] = platformTypeName
							platformsMAP[shapePosition+95] = platformTypeName
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = rInt(900, 930)
					shapeNumber = rInt(3, 6)

					shapePositionHOLDER = shapePosition
					shapeNumberHOLDER = shapeNumber

					for {

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							clearAroundBlock(shapePosition)
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = shapePositionHOLDER
					shapeNumber = shapeNumberHOLDER

					for {

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							levelMAP[shapePosition] = "floor"
							levelMAP[shapePosition+1] = "floor"
							levelMAP[shapePosition+94] = "floor"
							levelMAP[shapePosition+95] = "floor"

							platformsMAP[shapePosition] = platformTypeNameTL
							platformsMAP[shapePosition+1] = platformTypeName
							platformsMAP[shapePosition+94] = platformTypeName
							platformsMAP[shapePosition+95] = platformTypeName
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

				}
			} // end switch case 8
		case 9:

			chooseShapeNumber := rInt(1, 4)

			switch chooseShapeNumber {

			case 1:
				// horizontal line top with gap

				platformStartBlock := 940
				platformStartBlock += (rInt(-3, 4) * 94)
				platformStartBlockHOLDER := platformStartBlock

				for a := 0; a < 46; a++ {
					clearAroundBlock(platformStartBlock)
					platformStartBlock++
				}
				platformStartBlock = platformStartBlockHOLDER

				for a := 0; a < 46; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock++
				}

				platformStartBlock = platformStartBlockHOLDER
				platformStartBlock += rInt(30, 50)
				gapLength := rInt(6, 15)

				for a := 0; a < gapLength; a++ {

					levelMAP[platformStartBlock+a] = ""
					levelMAP[platformStartBlock+(a+1)] = ""
					levelMAP[platformStartBlock+(a+94)] = ""
					levelMAP[platformStartBlock+(a+95)] = ""

					platformsMAP[platformStartBlock+a] = ""
					platformsMAP[platformStartBlock+(a+1)] = ""
					platformsMAP[platformStartBlock+(a+94)] = ""
					platformsMAP[platformStartBlock+(a+95)] = ""

					platformStartBlock++
				}

			case 2:

				// first curl
				platformStartBlock := rInt(854, 886)
				platformStartBlock += (rInt(0, 25) * 94)

				platformStartBlockHOLDER := platformStartBlock

				lengthChange := 3

				for a := 0; a < lengthChange; a++ {
					platformStartBlock++
					clearAroundBlock(platformStartBlock)
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				clearAroundBlock(platformStartBlock)

				for a := 0; a < lengthChange; a++ {
					platformStartBlock += (94 * 2)
					clearAroundBlock(platformStartBlock)
				}

				platformStartBlock = platformStartBlockHOLDER
				lengthChange = 3

				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock++
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock-a] = "floor"
					levelMAP[platformStartBlock-(a+1)] = "floor"
					levelMAP[platformStartBlock-(a+94)] = "floor"
					levelMAP[platformStartBlock-(a+95)] = "floor"

					platformsMAP[platformStartBlock-a] = platformTypeNameTL
					platformsMAP[platformStartBlock-(a+1)] = platformTypeName
					platformsMAP[platformStartBlock-(a+94)] = platformTypeName
					platformsMAP[platformStartBlock-(a+95)] = platformTypeName

					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += (94 * 2)
				}

				// second curl

				platformStartBlock = rInt(1362, 1387)
				platformStartBlock += (rInt(0, 15) * 94)

				platformStartBlockHOLDER = platformStartBlock

				lengthChange = 3

				for a := 0; a < lengthChange; a++ {
					platformStartBlock++
					clearAroundBlock(platformStartBlock)
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					clearAroundBlock(platformStartBlock)
					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				clearAroundBlock(platformStartBlock)

				for a := 0; a < lengthChange; a++ {
					platformStartBlock += (94 * 2)
					clearAroundBlock(platformStartBlock)
				}

				platformStartBlock = platformStartBlockHOLDER
				lengthChange = 3

				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock++
				}
				platformStartBlock++

				for a := 0; a < (lengthChange - 1); a++ {
					platformStartBlock -= (94 * 2)
					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName
					platformStartBlock--
				}
				lengthChange++
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock-a] = "floor"
					levelMAP[platformStartBlock-(a+1)] = "floor"
					levelMAP[platformStartBlock-(a+94)] = "floor"
					levelMAP[platformStartBlock-(a+95)] = "floor"

					platformsMAP[platformStartBlock-a] = platformTypeNameTL
					platformsMAP[platformStartBlock-(a+1)] = platformTypeName
					platformsMAP[platformStartBlock-(a+94)] = platformTypeName
					platformsMAP[platformStartBlock-(a+95)] = platformTypeName

					platformStartBlock--
				}
				platformStartBlock -= 2
				platformStartBlock += (94 * 2)
				for a := 0; a < lengthChange; a++ {

					levelMAP[platformStartBlock] = "floor"
					levelMAP[platformStartBlock+1] = "floor"
					levelMAP[platformStartBlock+94] = "floor"
					levelMAP[platformStartBlock+95] = "floor"

					platformsMAP[platformStartBlock] = platformTypeNameTL
					platformsMAP[platformStartBlock+1] = platformTypeName
					platformsMAP[platformStartBlock+94] = platformTypeName
					platformsMAP[platformStartBlock+95] = platformTypeName

					platformStartBlock += (94 * 2)
				}

			case 3:
				// zigzag top
				platformStartBlock := rInt(1136, 1156)
				platformStartBlock += (rInt(9, 18) * 94)

				zigzagNumber := rInt(2, 6)
				zigzagNumberHOLDER := zigzagNumber

				platformStartBlockHOLDER := platformStartBlock

				// clear existing blocks
				for {
					for a := 0; a < 4; a++ {
						clearAroundBlock(platformStartBlock)
						platformStartBlock += 2
						platformStartBlock -= (94 * 2)
					}
					platformStartBlock += (94 * 4)
					clearAroundBlock(platformStartBlock)
					for a := 0; a < 3; a++ {
						clearAroundBlock(platformStartBlock)
						platformStartBlock += 2
						platformStartBlock += (94 * 2)
					}
					zigzagNumber--
					if zigzagNumber == 0 {
						break
					}

				}

				zigzagNumber = zigzagNumberHOLDER
				platformStartBlock = platformStartBlockHOLDER

				// draw platform shape
				for {
					for a := 0; a < 4; a++ {

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 2
						platformStartBlock -= (94 * 2)

					}
					platformStartBlock += (94 * 4)

					for a := 0; a < 3; a++ {

						levelMAP[platformStartBlock] = "floor"
						levelMAP[platformStartBlock+1] = "floor"
						levelMAP[platformStartBlock+94] = "floor"
						levelMAP[platformStartBlock+95] = "floor"

						platformsMAP[platformStartBlock] = platformTypeNameTL
						platformsMAP[platformStartBlock+1] = platformTypeName
						platformsMAP[platformStartBlock+94] = platformTypeName
						platformsMAP[platformStartBlock+95] = platformTypeName

						platformStartBlock += 2
						platformStartBlock += (94 * 2)

					}
					zigzagNumber--

					if zigzagNumber == 0 {
						break
					}
				}

			case 4:

				// cross shape right

				chooseShape1or2 := flipcoin()

				if chooseShape1or2 {

					shapePosition := rInt(900, 930)
					shapeNumber := rInt(3, 6)

					shapePositionHOLDER := shapePosition
					shapeNumberHOLDER := shapeNumber

					for {

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							clearAroundBlock(shapePosition)
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = shapePositionHOLDER
					shapeNumber = shapeNumberHOLDER

					for {

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							levelMAP[shapePosition] = "floor"
							levelMAP[shapePosition+1] = "floor"
							levelMAP[shapePosition+94] = "floor"
							levelMAP[shapePosition+95] = "floor"

							platformsMAP[shapePosition] = platformTypeNameTL
							platformsMAP[shapePosition+1] = platformTypeName
							platformsMAP[shapePosition+94] = platformTypeName
							platformsMAP[shapePosition+95] = platformTypeName
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}
				} else {

					shapePosition := rInt(860, 890)
					shapeNumber := rInt(3, 6)

					shapePositionHOLDER := shapePosition
					shapeNumberHOLDER := shapeNumber

					for {

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							clearAroundBlock(shapePosition)
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = shapePositionHOLDER
					shapeNumber = shapeNumberHOLDER

					for {

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							levelMAP[shapePosition] = "floor"
							levelMAP[shapePosition+1] = "floor"
							levelMAP[shapePosition+94] = "floor"
							levelMAP[shapePosition+95] = "floor"

							platformsMAP[shapePosition] = platformTypeNameTL
							platformsMAP[shapePosition+1] = platformTypeName
							platformsMAP[shapePosition+94] = platformTypeName
							platformsMAP[shapePosition+95] = platformTypeName
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = rInt(900, 930)
					shapeNumber = rInt(3, 6)

					shapePositionHOLDER = shapePosition
					shapeNumberHOLDER = shapeNumber

					for {

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							clearAroundBlock(shapePosition)
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						clearAroundBlock(shapePosition)

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

					shapePosition = shapePositionHOLDER
					shapeNumber = shapeNumberHOLDER

					for {

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 2
						shapePosition -= 2

						for a := 0; a < 3; a++ {
							levelMAP[shapePosition] = "floor"
							levelMAP[shapePosition+1] = "floor"
							levelMAP[shapePosition+94] = "floor"
							levelMAP[shapePosition+95] = "floor"

							platformsMAP[shapePosition] = platformTypeNameTL
							platformsMAP[shapePosition+1] = platformTypeName
							platformsMAP[shapePosition+94] = platformTypeName
							platformsMAP[shapePosition+95] = platformTypeName
							shapePosition += 2
						}

						shapePosition += 94 * 2
						shapePosition -= 4

						levelMAP[shapePosition] = "floor"
						levelMAP[shapePosition+1] = "floor"
						levelMAP[shapePosition+94] = "floor"
						levelMAP[shapePosition+95] = "floor"

						platformsMAP[shapePosition] = platformTypeNameTL
						platformsMAP[shapePosition+1] = platformTypeName
						platformsMAP[shapePosition+94] = platformTypeName
						platformsMAP[shapePosition+95] = platformTypeName

						shapePosition += 94 * 5
						shapePosition -= 4
						shapePosition += rInt(-4, 5)
						shapeNumber--

						if shapeNumber == 0 || shapePosition > 3760 {
							break
						}

					}

				}
			} // end switch case 9

		} // end switch platformLayoutType

	}

	// MARK: ground blocks

	if groundBlocksOn {

		platformStartBlock := rInt(4055, 4070)
		platformStartBlockHOLDER := platformStartBlock

		for {

			chooseGroundBlock := flipcoin()

			if chooseGroundBlock {
				for a := 0; a < 2; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock++

				}
				platformStartBlock = platformStartBlockHOLDER + 4
			} else {

				for a := 0; a < 2; a++ {

					levelMAP[platformStartBlock+a] = "floor"
					levelMAP[platformStartBlock+(a+1)] = "floor"
					levelMAP[platformStartBlock+(a+94)] = "floor"
					levelMAP[platformStartBlock+(a+95)] = "floor"

					platformsMAP[platformStartBlock+a] = platformTypeNameTL
					platformsMAP[platformStartBlock+(a+1)] = platformTypeName
					platformsMAP[platformStartBlock+(a+94)] = platformTypeName
					platformsMAP[platformStartBlock+(a+95)] = platformTypeName

					platformStartBlock -= (94 * 2)
					platformStartBlock--

				}
				platformStartBlock = platformStartBlockHOLDER + 4

			}

			platformStartBlock += rInt(6, 13)
			platformStartBlockHOLDER = platformStartBlock

			blockHorizontal := platformStartBlock / 94
			blockVertical := platformStartBlock - (blockHorizontal * 94)

			if blockVertical > 78 {
				break
			}

		}

	}

	// MARK: moving platforms up down
	if movingPlatformsOn {

		movingPlatformsNumber := rInt(1, 3)
		movingPlatformStartBlock := (rInt(8, 15) * 94) + rInt(20, 47)

		for {

			movingPlatformLength := rInt(2, 5)

			if movingPlatformLength%2 != 0 {
				movingPlatformLength++
			}

			for a := 0; a < movingPlatformLength; a++ {

				levelMAP[movingPlatformStartBlock] = "floor"
				platformsMAP[movingPlatformStartBlock] = "movingPlatformUpTL"
				platformsMAP[movingPlatformStartBlock+1] = "movingPlatformUp"

				movingPlatformStartBlock += 2

			}

			movingPlatformsNumber--
			movingPlatformStartBlock += (rInt(8, 15) * 94) + rInt(6, 31)

			if movingPlatformsNumber == 0 {
				break
			}
		}

	}

}

func chooseLevelMusic() {
	choose := rInt(1, 6)
	switch choose {
	case 1:
		levelTune = rl.LoadMusicStream("track1.ogg")
	case 2:
		levelTune = rl.LoadMusicStream("track2.ogg")
	case 3:
		levelTune = rl.LoadMusicStream("track3.ogg")
	case 4:
		levelTune = rl.LoadMusicStream("track4.ogg")
	case 5:
		levelTune = rl.LoadMusicStream("track5.ogg")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // random numbers

	rl.SetTraceLog(rl.LogError) // hides INFO window

	launcher()

	// TESTING REMOVE
	//difficultDiffOn = true // TESTING REMOVE
	//	fullscreenOn = false // TESTING REMOVE

	rl.SetExitKey(rl.KeyEnd) // key to end the game and close window
	selecty = 55             // fix for close launcher window space key end

	cSTARS()
	deadEnemyCircles()
	cLEVEL()
	cENEMIES()
	cCLOUDS()
	cRAIN()
	randomColors()
	cSTARTGAMEVARIABLES()
	rl.InitWindow(screenW, screenH, "spl@ta")
	rl.InitAudioDevice()

	introTune = rl.LoadMusicStream("intro_tune.mp3")
	levelEndTune = rl.LoadMusicStream("endshop_tune.mp3")

	chooseLevelMusic()

	// MARK: fullscreen
	// fullscreenOn = true
	if fullscreenOn {
		rl.ToggleFullscreen()
	}
	// MARK: load images
	fogImageGeneration := rl.GenImagePerlinNoise(int(screenW*3), int(screenH), 0, 0, 100.0)
	fogImage := rl.LoadTextureFromImage(fogImageGeneration)
	rl.UnloadImage(fogImageGeneration)
	imgs = rl.LoadTexture("imgs_splata.png")
	changeBackgroundImage()

	// MARK: TO DO

	/*
		LEVELS >> EXTRA FEATURES / HORIZONTAL MOVING PLATFORMS

		END LEVEL KEY >> 1 COIN COST TO END LEVEL / ELSE IF NO COINS THEN KILLZ = 0

		HIGH SCORE LEADERBOARD >> POST TO WEBSITE

	*/

	rl.SetTargetFPS(60)

	// MARK: WindowShouldClose
	for !rl.WindowShouldClose() {

		if musicOn {
			if startIntroMusic {
				rl.PlayMusicStream(introTune)
				rl.UpdateMusicStream(introTune)
			}
			if levelMusic {
				rl.PlayMusicStream(levelTune)
				rl.UpdateMusicStream(levelTune)
			}
			if levelEndMusic {
				rl.PlayMusicStream(levelEndTune)
				rl.UpdateMusicStream(levelEndTune)
			}
		}

		frameCountGameStart++

		// switch scan lines
		if frameCountGameStart%15 == 0 {

			if switchScanLines {
				switchScanLines = false
			} else {
				switchScanLines = true
			}
		}

		// MARK: frameCount timers
		if pauseOn == false {
			frameCount++
			// countdown timer
			if frameCount%60 == 0 {
				secondsCountdown--
				if secondsCountdown == 0 {
					rl.StopMusicStream(levelTune)
					pauseOn = true
					levelEnd = true
				}
			}
			// power up timer
			powerUpTimer--
			if powerUpTimer <= 0 {
				powerUpTimer = rInt(500, 800)
				powerUpDropped = false
				powerUpHasDropped = false
				cPOWERUPS()
			}

		}

		playerHorizontal = playerCurrentBlock / 94
		playerVertical = playerCurrentBlock - (playerHorizontal * 94)

		// MARK: keys

		if pauseOn == false {

			// TESTING REMOVE
			if rl.IsKeyPressed(rl.KeyNine) {
				platformDistortionOn = true
				createShockBlocks = false
			}
			if rl.IsKeyPressed(rl.KeyEight) {
				shockBlocksOn = true
				createShockBlocks = false
			}
			if rl.IsKeyPressed(rl.KeySeven) {
				snowOn = true
				snowTimer = 2
			}
			if rl.IsKeyPressed(rl.KeySix) {
				rainoffireOn = true
				rainoffireTimer = 1
			}
			if rl.IsKeyPressed(rl.KeyFive) {
				meteorOn = true
				createMeteor = false
			}
			if rl.IsKeyPressed(rl.KeyFour) {
				tornadoOn = true
				createTornado = false
				tornadoStartBlock = 3988
			}
			if rl.IsKeyPressed(rl.KeyThree) {
				frogRainOn = true
				rainFrogOn = true
			}

			if rl.IsKeyPressed(rl.KeyF2) {
				if levelEnd {
					pauseOn = false
					levelEnd = false
				} else {
					pauseOn = true
					levelEnd = true
				}
			}

			// player keys

			if rl.IsKeyPressed(rl.KeySpace) {

				shootSound = true

				switch currentPlayerWeapon {

				case "weapon1TL":
					switch playerDirection {
					case "left":
						bulletsMAP[playerCurrentBlock-2] = "bulletL"
					case "right":
						bulletsMAP[playerCurrentBlock+3] = "bulletR"
					}
				case "weapon2TL":
					switch playerDirection {
					case "left":
						bulletsMAP[playerCurrentBlock-2] = "bulletL"
						bulletsMAP[playerCurrentBlock+92] = "bulletL"
					case "right":
						bulletsMAP[playerCurrentBlock+3] = "bulletR"
						bulletsMAP[playerCurrentBlock+97] = "bulletR"
					}
				case "weapon3TL":
					switch playerDirection {
					case "left":
						bulletsMAP[playerCurrentBlock-2] = "bulletL"
						bulletsMAP[playerCurrentBlock-4] = "bulletL"
					case "right":
						bulletsMAP[playerCurrentBlock+3] = "bulletR"
						bulletsMAP[playerCurrentBlock+5] = "bulletR"
					}
				case "weapon4TL":
					switch playerDirection {
					case "left":
						bulletsMAP[playerCurrentBlock-96] = "bulletL"
						bulletsMAP[playerCurrentBlock-2] = "bulletL"
						bulletsMAP[playerCurrentBlock+92] = "bulletL"
					case "right":
						bulletsMAP[playerCurrentBlock-91] = "bulletR"
						bulletsMAP[playerCurrentBlock+3] = "bulletR"
						bulletsMAP[playerCurrentBlock+97] = "bulletR"
					}
				case "weapon5TL":
					switch playerDirection {
					case "left":
						bulletsMAP[playerCurrentBlock-2] = "bulletL"
						bulletsMAP[playerCurrentBlock+92] = "bulletL"
						bulletsMAP[playerCurrentBlock-4] = "bulletL"
						bulletsMAP[playerCurrentBlock+90] = "bulletL"
					case "right":
						bulletsMAP[playerCurrentBlock+3] = "bulletR"
						bulletsMAP[playerCurrentBlock+97] = "bulletR"
						bulletsMAP[playerCurrentBlock+5] = "bulletR"
						bulletsMAP[playerCurrentBlock+99] = "bulletR"
					}
				case "weapon6TL":
					switch playerDirection {
					case "left":
						bulletsMAP[playerCurrentBlock-2] = "bulletL"
						bulletsMAP[playerCurrentBlock+90] = "bulletL"
					case "right":
						bulletsMAP[playerCurrentBlock+3] = "bulletR"
						bulletsMAP[playerCurrentBlock+99] = "bulletR"
					}
				case "weapon7TL":
					switch playerDirection {
					case "left":
						bulletsMAP[playerCurrentBlock-192] = "bulletL"
						bulletsMAP[playerCurrentBlock-98] = "bulletL"
						bulletsMAP[playerCurrentBlock-96] = "bulletL"
						bulletsMAP[playerCurrentBlock-2] = "bulletL"
						bulletsMAP[playerCurrentBlock+92] = "bulletL"
						bulletsMAP[playerCurrentBlock-4] = "bulletL"
						bulletsMAP[playerCurrentBlock+90] = "bulletL"
					case "right":

						bulletsMAP[playerCurrentBlock-91] = "bulletR"
						bulletsMAP[playerCurrentBlock+3] = "bulletR"
						bulletsMAP[playerCurrentBlock+97] = "bulletR"
						bulletsMAP[playerCurrentBlock+5] = "bulletR"
						bulletsMAP[playerCurrentBlock+99] = "bulletR"
					}

				}
			}
			if rl.IsKeyPressed(rl.KeyLeftAlt) || rl.IsKeyPressed(rl.KeyRightAlt) || rl.IsKeyPressed(rl.KeyN) || rl.IsKeyPressed(rl.KeyV) || rl.IsKeyPressed(rl.KeyB) {
				switch playerDirection {
				case "left":
					bulletsMAP[playerCurrentBlock-94] = "bulletU"
					bulletsMAP[playerCurrentBlock-95] = "bulletU"
				case "right":
					bulletsMAP[playerCurrentBlock-93] = "bulletU"
					bulletsMAP[playerCurrentBlock-92] = "bulletU"
				}
			}
			if rl.IsKeyPressed(rl.KeyRight) {
				playerDirection = "right"
				clearPlayer()
				playerCurrentBlock++
				if earthquakeFallOn == false {
					if playerCurrentBlock >= 4134 {
						playerCurrentBlock = 4044
					}
				}

			} else if rl.IsKeyDown(rl.KeyRight) {
				playerDirection = "right"

				clearPlayer()
				playerCurrentBlock++
			}
			if rl.IsKeyPressed(rl.KeyLeft) {
				playerDirection = "left"

				clearPlayer()
				playerCurrentBlock--

			} else if rl.IsKeyDown(rl.KeyLeft) {
				playerDirection = "left"

				clearPlayer()
				playerCurrentBlock--

			}
			if rl.IsKeyPressed(rl.KeyUp) {

				if playerHorizontal > 2 && fallActive == false {
					if propellorOn {
						jumpHeight = 40
					} else {
						jumpHeight = 12
					}

					jumpActive = true
				}
			}
			if rl.IsKeyPressed(rl.KeyDown) {
				if playerHorizontal < 43 {
					clearPlayer()
					playerCurrentBlock += 94
				}
			} else if rl.IsKeyDown(rl.KeyDown) {
				if playerHorizontal < 43 {
					clearPlayer()
					playerCurrentBlock += 94
				}
			}

		} // end pauseOn
		// options menu key
		if rl.IsKeyPressed(rl.KeyF1) {
			if gameMenuOn {
				gameMenuOn = false
				pauseOn = false
			} else {
				menuSelectNumber = 1
				menuColumnSelect = 1
				menuSelectBoxY = int32(95)
				gameMenuOn = true
				pauseOn = true

			}
		}
		// debugging menu key
		if rl.IsKeyPressed(rl.KeyKpDecimal) {
			if debuggingOn {
				debuggingOn = false
			} else {
				debuggingOn = true
			}
		}
		// update block numbers for movement
		// correct floor level

		playerGroundBlockL = levelMAP[playerCurrentBlock+(94*2)]
		playerGroundBlockR = levelMAP[playerCurrentBlock+((94*2)+1)]
		if jumpActive == false {
			if playerGroundBlockL != "floor" && playerGroundBlockR != "floor" {
				fallActive = true
			}
		}

		if horizplaton {
			if playerGroundBlockL == "floor" && platformsMAP[playerCurrentBlock+(94*2)] == "movingplatH1" || playerGroundBlockL == "floor" && platformsMAP[playerCurrentBlock+(94*2)] == "movingplatH2" || playerGroundBlockL == "floor" && platformsMAP[playerCurrentBlock+(94*2)] == "movingplatH3" || playerGroundBlockL == "floor" && platformsMAP[playerCurrentBlock+(94*2)] == "movingplatH4" || playerGroundBlockL == "floor" && platformsMAP[playerCurrentBlock+(94*2)] == "movingplatH5" || playerGroundBlockL == "floor" && platformsMAP[playerCurrentBlock+(94*2)] == "movingplatH6" || playerGroundBlockL == "floor" && platformsMAP[playerCurrentBlock+(94*2)] == "movingplatH7" || playerGroundBlockL == "floor" && platformsMAP[playerCurrentBlock+(94*2)] == "movingplatH8" {

				if horizplatlr {
					clearPlayer()
					playerCurrentBlock--
				}

			}

			if playerGroundBlockR == "floor" && platformsMAP[playerCurrentBlock+((94*2)+1)] == "movingplatH1" || playerGroundBlockR == "floor" && platformsMAP[playerCurrentBlock+((94*2)+1)] == "movingplatH2" || playerGroundBlockR == "floor" && platformsMAP[playerCurrentBlock+((94*2)+1)] == "movingplatH3" || playerGroundBlockR == "floor" && platformsMAP[playerCurrentBlock+((94*2)+1)] == "movingplatH4" || playerGroundBlockR == "floor" && platformsMAP[playerCurrentBlock+((94*2)+1)] == "movingplatH5" || playerGroundBlockR == "floor" && platformsMAP[playerCurrentBlock+((94*2)+1)] == "movingplatH6" || playerGroundBlockR == "floor" && platformsMAP[playerCurrentBlock+((94*2)+1)] == "movingplatH7" || playerGroundBlockR == "floor" && platformsMAP[playerCurrentBlock+((94*2)+1)] == "movingplatH8" {

				if horizplatlr == false {
					clearPlayer()
					playerCurrentBlock++
				}

			}
		}

		if playerGroundBlockL == "floor" && levelMAP[playerCurrentBlock+94] == "floor" {
			clearPlayer()
			playerCurrentBlock -= 94
		}
		if playerGroundBlockR == "floor" && levelMAP[playerCurrentBlock+95] == "floor" {
			clearPlayer()
			playerCurrentBlock -= 94
		}
		// top & bottom boundaries
		if playerCurrentBlock > 4134 {
			clearPlayer()
			playerCurrentBlock -= 94
		}
		if playerCurrentBlock < 188 {
			clearPlayer()
			playerCurrentBlock += 94
		}

		// jump
		if jumpActive {
			if jumpHeight != 0 {

				clearPlayer()
				playerCurrentBlock -= 94
				jumpHeight--

			} else {
				fallActive = true
				jumpActive = false
			}
		}
		// fall
		if fallActive == true {

			if playerGroundBlockL != "floor" && playerGroundBlockR != "floor" {
				clearPlayer()
				playerCurrentBlock += 94
			} else if playerGroundBlockL == "floor" || playerGroundBlockR == "floor" {
				fallActive = false
			}

		}

		// expand player
		playerMAP[playerCurrentBlock] = "playerTL"
		playerMAP[playerCurrentBlock+1] = "player"
		playerMAP[playerCurrentBlock+94] = "player"
		playerMAP[playerCurrentBlock+95] = "player"

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		// MARK:  game start

		// MARK: intro logos
		if introStoryScreenOn {

			if rl.IsKeyPressed(rl.KeySpace) {
				introStoryScreenOn = false
				introScreenOn = true
			}
			if startLogosOn {

				// background rectangle
				rl.DrawRectangle(0, 0, screenW, screenH, rl.Black)

				//	logoStep4 = true
				if logoStep4 {

					rl.DrawText("a game by nicholasimon", 340, 356, 60, rl.White)

					if frameCountGameStart%logoScount == 0 {
						if logoSflash {
							logoSflash = false
						} else {
							logoSflash = true
						}
					}
					if logoSflash {
						rl.DrawText("s", 874, 356, 60, rl.Maroon)
						rl.DrawText("s", 874, 356, 60, rl.Fade(rl.Black, 0.6))
						logoScount = rInt(14, 21)
					}

					if frameCountGameStart%60 == 0 {
						introCount2--
						if introCount2 <= 0 {
							introCount2 = 0
						}
					}
					if introCount2 == 0 {
						if rectangleFade < 1.0 {
							if frameCountGameStart%3 == 0 {
								rectangleFade += 0.02
							}
						} else if rectangleFade >= 0.9 {
							logoStep4 = false
							startLogosOn = false
						}

					}

				} else {
					if logoStep2 {

						rl.DrawText("built with", 595, 200, 40, rl.Fade(rl.Blue, 0.6))

						gopherV2 = rl.NewVector2(555, 256)
						rl.DrawTextureRec(imgs, raylibIMG, gopherV2, rl.White)
						if logoStep3 {
							if rectangleFade < 1.0 {
								if frameCountGameStart%3 == 0 {
									rectangleFade += 0.02
								}
							} else if rectangleFade >= 0.9 {
								logoStep3 = false
								rectangleFade = 0.0
								logoStep4 = true
							}
						} else {
							if rectangleFade > 0.0 {
								if frameCountGameStart%3 == 0 {
									rectangleFade -= 0.02
								}
							} else if rectangleFade <= 0.1 {
								logoStep3 = true
							}
						}
					} else {
						gopherV2 = rl.NewVector2(444, 256)
						rl.DrawTextureRec(imgs, gopherIMG, gopherV2, rl.White)
						if logoStep1 == false {
							if rectangleFade > 0.0 {
								if frameCountGameStart%3 == 0 {
									rectangleFade -= 0.02
								}
							} else if rectangleFade <= 0.9 {
								logoStep1 = true
							}
						}
						if logoStep1 {
							rl.DrawText("built with", 595, 200, 40, rl.Fade(rl.Blue, 0.6))
							if rectangleFade < 1.0 {
								if frameCountGameStart%3 == 0 {
									rectangleFade += 0.02
								}
							} else if rectangleFade >= 0.9 {
								logoStep2 = true
							}
						}
					}

				}
				// fade rectangle
				rl.DrawRectangle(0, 250, screenW, 300, rl.Fade(rl.Black, rectangleFade))

			} else {
				startIntroMusic = true

				if frameCountGameStart%30 == 0 {
					for a := 0; a < 5264; a++ {
						checkStarFade := starsFadeMAP[a]
						switch checkStarFade {
						case 3:
							starsFadeMAP[a] = 2
						case 2:
							starsFadeMAP[a] = 1
						case 1:
							starsFadeMAP[a] = 0
						case 0:
							starsFadeMAP[a] = 3
						}
					}
				}

				drawScreenCurrentBlock = 0
				drawScreenLineCount := 0
				drawScreenX := int32(0)
				drawScreenY := int32(0)
				rl.BeginMode2D(camera)

				for a := 0; a < 5264; a++ {
					checkStarsBlock := starsMAP[drawScreenCurrentBlock]
					checkStarFade := starsFadeMAP[drawScreenCurrentBlock]

					switch checkStarsBlock {
					case "star4":
						starV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-16))
						starSizeV2 := rl.NewVector2(float32(2), float32(2))
						switch checkStarFade {
						case 3:
							rl.DrawRectangleV(starV2, starSizeV2, rl.Yellow)
						case 2:
							rl.DrawRectangleV(starV2, starSizeV2, rl.Fade(rl.Yellow, 0.8))
						case 1:
							rl.DrawRectangleV(starV2, starSizeV2, rl.Fade(rl.Yellow, 0.6))
						case 0:
							rl.DrawRectangleV(starV2, starSizeV2, rl.Fade(rl.Yellow, 0.4))
						}
					case "star3":
						starV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-16))
						starSizeV2 := rl.NewVector2(float32(2), float32(2))
						switch checkStarFade {
						case 3:
							rl.DrawRectangleV(starV2, starSizeV2, rl.Red)
						case 2:
							rl.DrawRectangleV(starV2, starSizeV2, rl.Fade(rl.Red, 0.8))
						case 1:
							rl.DrawRectangleV(starV2, starSizeV2, rl.Fade(rl.Red, 0.6))
						case 0:
							rl.DrawRectangleV(starV2, starSizeV2, rl.Fade(rl.Red, 0.4))
						}
					case "star2":
						starV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-16))
						starSizeV2 := rl.NewVector2(float32(2), float32(2))
						switch checkStarFade {
						case 3:
							rl.DrawRectangleV(starV2, starSizeV2, rl.Blue)
						case 2:
							rl.DrawRectangleV(starV2, starSizeV2, rl.Fade(rl.Blue, 0.8))
						case 1:
							rl.DrawRectangleV(starV2, starSizeV2, rl.Fade(rl.Blue, 0.6))
						case 0:
							rl.DrawRectangleV(starV2, starSizeV2, rl.Fade(rl.Blue, 0.4))
						}
					}

					drawScreenCurrentBlock++
					drawScreenLineCount++
					drawScreenX += 8
					if drawScreenLineCount == 94 {
						drawScreenLineCount = 0
						drawScreenX = 0
						drawScreenY += 8
					}

				}

				// draw bullet
				if dinoIntroV2.X == 195 {
					createBulletV2on = true
				}
				if createBulletV2on {
					introBulletV2 = rl.NewVector2(dinoIntroV2.X+40, dinoIntroV2.Y+6)
					createBulletV2on = false
				}
				if dinoIntroV2.X > 198 {
					createBulletV2on = false
					if introBulletV2.X < enemy1IntroV2.X {
						introBulletV2.X += 3
						rl.DrawCircleV(introBulletV2, 3, rl.Gold)
					}
				}
				if introStoryScreenTextY < 400 {

					if introCirclesTimer < 92 {
						// draw enemy
						rl.DrawTextureRec(imgs, enemy1IMG, enemy1IntroV2, rl.White)
					}

					// draw dino
					rl.DrawTextureRec(imgs, exclamationIMG, exclamationIntroV2, rl.White)
					rl.DrawTextureRec(imgs, dinoGreenRIMG, dinoIntroV2, rl.White)

					// draw weapon
					if dinoIntroV2.X > 150 {
						weaponV2 := rl.NewVector2(dinoIntroV2.X+15, dinoIntroV2.Y+3)
						rl.DrawTextureRec(imgs, weapon2IMG, weaponV2, rl.White)

					}
				}
				rl.EndMode2D()

				// update timers

				if introStoryScreenTextY < 400 {
					if frameCountGameStart%8 == 0 {
						dinoGreenRIMG.X += 24
						if dinoGreenRIMG.X >= 100 {
							dinoGreenRIMG.X = 4
						}
					}
					if frameCountGameStart%15 == 0 {
						exclamationIMG.X += 34
						if exclamationIMG.X == 1397 {
							exclamationIMG.X = 1295
						}
					}
					if enemy1IntroV2.X > 400 {
						enemy1IntroV2.X--
					}
					if dinoIntroV2.X < 200 {
						dinoIntroV2.X++
						exclamationIntroV2.X++
					}
					if frameCountGameStart%9 == 0 {
						enemy1IMG.X += 32
						if enemy1IMG.X >= 1024 {
							enemy1IMG.X = 514

						}
					}
				}

				// intro text
				if introStoryScreenTextY4 > -50 {
					introStoryScreenTextY--
					introStoryScreenTextY2--
					introStoryScreenTextY3--
					introStoryScreenTextY4--
				}

				if introStoryScreenTextY4 <= -45 {
					introStoryScreenOn = false
					introScreenOn = true
				}

				rl.DrawText("it was a clear, starry nite somewhere in the galaxy...", 140, introStoryScreenTextY, introStoryScreenTextSize, rl.White)
				rl.DrawText("and the dinosaur home planet of Goofilbig was being", 146, introStoryScreenTextY2, introStoryScreenTextSize, rl.White)
				rl.DrawText("invaded by an assortment of odd creatures", 230, introStoryScreenTextY3, introStoryScreenTextSize, rl.White)
				rl.DrawText("now it is your job to save them...", 350, introStoryScreenTextY4, introStoryScreenTextSize, rl.White)

				if spaceToStart {
					rl.DrawText("press space to start", 350, 354, 60, rl.Fade(introTextSpaceColor, introTextFade))

					if introTextFade < 1.0 {
						introTextFade += 0.02
					}

					if frameCountGameStart%30 == 0 {

						if introTextSpaceColor == rl.Yellow {
							introTextSpaceColor = rl.Gold
						} else {
							introTextSpaceColor = rl.Yellow
						}

					}
				}

				if frameCountGameStart%6 == 0 {
					introCirclesTimer++
				}
				if introCirclesTimer == 92 {
					if introScreenShake == false {
						screenShake = true
						introScreenShake = true
					}
					deadEnemyRadius1 := deadEnemiesCirclesRadius[0]
					deadEnemies1V2 := deadEnemiesCirclesV2[0]
					deadEnemyRadius2 := deadEnemiesCirclesRadius[1]
					deadEnemies2V2 := deadEnemiesCirclesV2[1]
					deadEnemyRadius3 := deadEnemiesCirclesRadius[2]
					deadEnemies3V2 := deadEnemiesCirclesV2[2]
					deadEnemyRadius4 := deadEnemiesCirclesRadius[3]
					deadEnemies4V2 := deadEnemiesCirclesV2[3]
					deadEnemyRadius5 := deadEnemiesCirclesRadius[4]
					deadEnemies5V2 := deadEnemiesCirclesV2[4]
					deadEnemyRadius6 := deadEnemiesCirclesRadius[5]
					deadEnemies6V2 := deadEnemiesCirclesV2[5]
					deadEnemyRadius7 := deadEnemiesCirclesRadius[6]
					deadEnemies7V2 := deadEnemiesCirclesV2[6]
					deadEnemyRadius8 := deadEnemiesCirclesRadius[7]
					deadEnemies8V2 := deadEnemiesCirclesV2[7]
					deadEnemyRadius9 := deadEnemiesCirclesRadius[8]
					deadEnemies9V2 := deadEnemiesCirclesV2[8]
					deadEnemyRadius10 := deadEnemiesCirclesRadius[9]
					deadEnemies10V2 := deadEnemiesCirclesV2[9]
					deadEnemyRadius11 := deadEnemiesCirclesRadius[10]
					deadEnemies11V2 := deadEnemiesCirclesV2[10]
					deadEnemyRadius12 := deadEnemiesCirclesRadius[11]
					deadEnemies12V2 := deadEnemiesCirclesV2[11]
					deadEnemyRadius13 := deadEnemiesCirclesRadius[12]
					deadEnemies13V2 := deadEnemiesCirclesV2[12]
					deadEnemyRadius14 := deadEnemiesCirclesRadius[13]
					deadEnemies14V2 := deadEnemiesCirclesV2[13]
					deadEnemyRadius15 := deadEnemiesCirclesRadius[14]
					deadEnemies15V2 := deadEnemiesCirclesV2[14]

					rl.DrawCircleV(deadEnemies1V2, deadEnemyRadius1, rl.Fade(rl.Maroon, 0.8))
					rl.DrawCircleV(deadEnemies2V2, deadEnemyRadius2, rl.Fade(rl.Maroon, 0.8))
					rl.DrawCircleV(deadEnemies3V2, deadEnemyRadius3, rl.Fade(rl.Maroon, 0.8))
					rl.DrawCircleV(deadEnemies4V2, deadEnemyRadius4, rl.Fade(rl.Maroon, 0.8))
					rl.DrawCircleV(deadEnemies5V2, deadEnemyRadius5, rl.Fade(rl.Maroon, 0.8))
					rl.DrawCircleV(deadEnemies6V2, deadEnemyRadius6, rl.Fade(rl.Maroon, 0.8))
					rl.DrawCircleV(deadEnemies7V2, deadEnemyRadius7, rl.Fade(rl.Maroon, 0.8))
					rl.DrawCircleV(deadEnemies8V2, deadEnemyRadius8, rl.Fade(rl.Maroon, 0.8))
					rl.DrawCircleV(deadEnemies9V2, deadEnemyRadius9, rl.Fade(rl.Maroon, 0.8))
					rl.DrawCircleV(deadEnemies10V2, deadEnemyRadius10, rl.Fade(rl.Maroon, 0.8))
					rl.DrawCircleV(deadEnemies11V2, deadEnemyRadius11, rl.Fade(rl.Maroon, 0.8))
					rl.DrawCircleV(deadEnemies12V2, deadEnemyRadius12, rl.Fade(rl.Maroon, 0.8))
					rl.DrawCircleV(deadEnemies13V2, deadEnemyRadius13, rl.Fade(rl.Maroon, 0.8))
					rl.DrawCircleV(deadEnemies14V2, deadEnemyRadius14, rl.Fade(rl.Maroon, 0.8))
					rl.DrawCircleV(deadEnemies15V2, deadEnemyRadius15, rl.Fade(rl.Maroon, 0.8))
				}

				if introCirclesTimer == 93 {
					deadEnemyRadius1 := deadEnemiesCirclesRadius[0]
					deadEnemies1V2 := deadEnemiesCirclesV2[0]
					deadEnemyRadius2 := deadEnemiesCirclesRadius[1]
					deadEnemies2V2 := deadEnemiesCirclesV2[1]
					deadEnemyRadius3 := deadEnemiesCirclesRadius[2]
					deadEnemies3V2 := deadEnemiesCirclesV2[2]
					deadEnemyRadius4 := deadEnemiesCirclesRadius[3]
					deadEnemies4V2 := deadEnemiesCirclesV2[3]
					deadEnemyRadius5 := deadEnemiesCirclesRadius[4]
					deadEnemies5V2 := deadEnemiesCirclesV2[4]
					deadEnemyRadius6 := deadEnemiesCirclesRadius[5]
					deadEnemies6V2 := deadEnemiesCirclesV2[5]
					deadEnemyRadius7 := deadEnemiesCirclesRadius[6]
					deadEnemies7V2 := deadEnemiesCirclesV2[6]
					deadEnemyRadius8 := deadEnemiesCirclesRadius[7]
					deadEnemies8V2 := deadEnemiesCirclesV2[7]
					deadEnemyRadius9 := deadEnemiesCirclesRadius[8]
					deadEnemies9V2 := deadEnemiesCirclesV2[8]
					deadEnemyRadius10 := deadEnemiesCirclesRadius[9]
					deadEnemies10V2 := deadEnemiesCirclesV2[9]
					deadEnemyRadius11 := deadEnemiesCirclesRadius[10]
					deadEnemies11V2 := deadEnemiesCirclesV2[10]
					deadEnemyRadius12 := deadEnemiesCirclesRadius[11]
					deadEnemies12V2 := deadEnemiesCirclesV2[11]
					deadEnemyRadius13 := deadEnemiesCirclesRadius[12]
					deadEnemies13V2 := deadEnemiesCirclesV2[12]
					deadEnemyRadius14 := deadEnemiesCirclesRadius[13]
					deadEnemies14V2 := deadEnemiesCirclesV2[13]
					deadEnemyRadius15 := deadEnemiesCirclesRadius[14]
					deadEnemies15V2 := deadEnemiesCirclesV2[14]
					rl.DrawCircleV(deadEnemies1V2, deadEnemyRadius1-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies1V2, deadEnemyRadius1-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies1V2.X), int32(deadEnemies1V2.Y), deadEnemyRadius1-4, rl.Black)
					rl.DrawCircleV(deadEnemies2V2, deadEnemyRadius2-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies2V2, deadEnemyRadius2-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies2V2.X), int32(deadEnemies2V2.Y), deadEnemyRadius2-4, rl.Black)
					rl.DrawCircleV(deadEnemies3V2, deadEnemyRadius3-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies3V2, deadEnemyRadius3-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies3V2.X), int32(deadEnemies3V2.Y), deadEnemyRadius3-4, rl.Black)
					rl.DrawCircleV(deadEnemies4V2, deadEnemyRadius4-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies4V2, deadEnemyRadius4-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies4V2.X), int32(deadEnemies4V2.Y), deadEnemyRadius4-4, rl.Black)
					rl.DrawCircleV(deadEnemies5V2, deadEnemyRadius5-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies5V2, deadEnemyRadius5-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies5V2.X), int32(deadEnemies5V2.Y), deadEnemyRadius5-4, rl.Black)
					rl.DrawCircleV(deadEnemies6V2, deadEnemyRadius6-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies6V2, deadEnemyRadius6-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies6V2.X), int32(deadEnemies6V2.Y), deadEnemyRadius6-4, rl.Black)
					rl.DrawCircleV(deadEnemies7V2, deadEnemyRadius7-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies7V2, deadEnemyRadius7-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies7V2.X), int32(deadEnemies7V2.Y), deadEnemyRadius7-4, rl.Black)
					rl.DrawCircleV(deadEnemies8V2, deadEnemyRadius8-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies8V2, deadEnemyRadius8-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies8V2.X), int32(deadEnemies8V2.Y), deadEnemyRadius8-4, rl.Black)
					rl.DrawCircleV(deadEnemies9V2, deadEnemyRadius9-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies9V2, deadEnemyRadius9-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies9V2.X), int32(deadEnemies9V2.Y), deadEnemyRadius9-4, rl.Black)
					rl.DrawCircleV(deadEnemies10V2, deadEnemyRadius10-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies10V2, deadEnemyRadius10-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies10V2.X), int32(deadEnemies10V2.Y), deadEnemyRadius10-4, rl.Black)
					rl.DrawCircleV(deadEnemies11V2, deadEnemyRadius11-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies11V2, deadEnemyRadius11-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies11V2.X), int32(deadEnemies11V2.Y), deadEnemyRadius11-4, rl.Black)
					rl.DrawCircleV(deadEnemies12V2, deadEnemyRadius12-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies12V2, deadEnemyRadius12-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies12V2.X), int32(deadEnemies12V2.Y), deadEnemyRadius12-4, rl.Black)
					rl.DrawCircleV(deadEnemies13V2, deadEnemyRadius13-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies13V2, deadEnemyRadius13-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies13V2.X), int32(deadEnemies13V2.Y), deadEnemyRadius13-4, rl.Black)
					rl.DrawCircleV(deadEnemies14V2, deadEnemyRadius14-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies14V2, deadEnemyRadius14-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies14V2.X), int32(deadEnemies14V2.Y), deadEnemyRadius14-4, rl.Black)
					rl.DrawCircleV(deadEnemies15V2, deadEnemyRadius15-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies15V2, deadEnemyRadius15-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies15V2.X), int32(deadEnemies15V2.Y), deadEnemyRadius15-4, rl.Black)
				}

				if introCirclesTimer == 94 {
					spaceToStart = true
					deadEnemyRadius1 := deadEnemiesCirclesRadius[0]
					deadEnemies1V2 := deadEnemiesCirclesV2[0]
					deadEnemyRadius2 := deadEnemiesCirclesRadius[1]
					deadEnemies2V2 := deadEnemiesCirclesV2[1]
					deadEnemyRadius3 := deadEnemiesCirclesRadius[2]
					deadEnemies3V2 := deadEnemiesCirclesV2[2]
					deadEnemyRadius4 := deadEnemiesCirclesRadius[3]
					deadEnemies4V2 := deadEnemiesCirclesV2[3]
					deadEnemyRadius5 := deadEnemiesCirclesRadius[4]
					deadEnemies5V2 := deadEnemiesCirclesV2[4]
					deadEnemyRadius6 := deadEnemiesCirclesRadius[5]
					deadEnemies6V2 := deadEnemiesCirclesV2[5]
					deadEnemyRadius7 := deadEnemiesCirclesRadius[6]
					deadEnemies7V2 := deadEnemiesCirclesV2[6]
					deadEnemyRadius8 := deadEnemiesCirclesRadius[7]
					deadEnemies8V2 := deadEnemiesCirclesV2[7]
					deadEnemyRadius9 := deadEnemiesCirclesRadius[8]
					deadEnemies9V2 := deadEnemiesCirclesV2[8]
					deadEnemyRadius10 := deadEnemiesCirclesRadius[9]
					deadEnemies10V2 := deadEnemiesCirclesV2[9]
					deadEnemyRadius11 := deadEnemiesCirclesRadius[10]
					deadEnemies11V2 := deadEnemiesCirclesV2[10]
					deadEnemyRadius12 := deadEnemiesCirclesRadius[11]
					deadEnemies12V2 := deadEnemiesCirclesV2[11]
					deadEnemyRadius13 := deadEnemiesCirclesRadius[12]
					deadEnemies13V2 := deadEnemiesCirclesV2[12]
					deadEnemyRadius14 := deadEnemiesCirclesRadius[13]
					deadEnemies14V2 := deadEnemiesCirclesV2[13]
					deadEnemyRadius15 := deadEnemiesCirclesRadius[14]
					deadEnemies15V2 := deadEnemiesCirclesV2[14]
					rl.DrawCircleV(deadEnemies1V2, deadEnemyRadius1-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies1V2, deadEnemyRadius1-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies1V2.X), int32(deadEnemies1V2.Y), deadEnemyRadius1-6, rl.Black)
					rl.DrawCircleV(deadEnemies2V2, deadEnemyRadius2-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies2V2, deadEnemyRadius2-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies2V2.X), int32(deadEnemies2V2.Y), deadEnemyRadius2-6, rl.Black)
					rl.DrawCircleV(deadEnemies3V2, deadEnemyRadius3-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies3V2, deadEnemyRadius3-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies3V2.X), int32(deadEnemies3V2.Y), deadEnemyRadius3-6, rl.Black)
					rl.DrawCircleV(deadEnemies4V2, deadEnemyRadius4-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies4V2, deadEnemyRadius4-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies4V2.X), int32(deadEnemies4V2.Y), deadEnemyRadius4-6, rl.Black)
					rl.DrawCircleV(deadEnemies5V2, deadEnemyRadius5-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies5V2, deadEnemyRadius5-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies5V2.X), int32(deadEnemies5V2.Y), deadEnemyRadius5-6, rl.Black)
					rl.DrawCircleV(deadEnemies6V2, deadEnemyRadius6-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies6V2, deadEnemyRadius6-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies6V2.X), int32(deadEnemies6V2.Y), deadEnemyRadius6-6, rl.Black)
					rl.DrawCircleV(deadEnemies7V2, deadEnemyRadius7-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies7V2, deadEnemyRadius7-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies7V2.X), int32(deadEnemies7V2.Y), deadEnemyRadius7-6, rl.Black)
					rl.DrawCircleV(deadEnemies8V2, deadEnemyRadius8-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies8V2, deadEnemyRadius8-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies8V2.X), int32(deadEnemies8V2.Y), deadEnemyRadius8-6, rl.Black)
					rl.DrawCircleV(deadEnemies9V2, deadEnemyRadius9-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies9V2, deadEnemyRadius9-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies9V2.X), int32(deadEnemies9V2.Y), deadEnemyRadius9-6, rl.Black)
					rl.DrawCircleV(deadEnemies10V2, deadEnemyRadius10-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies10V2, deadEnemyRadius10-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies10V2.X), int32(deadEnemies10V2.Y), deadEnemyRadius10-6, rl.Black)
					rl.DrawCircleV(deadEnemies11V2, deadEnemyRadius11-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies11V2, deadEnemyRadius11-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies11V2.X), int32(deadEnemies11V2.Y), deadEnemyRadius11-6, rl.Black)
					rl.DrawCircleV(deadEnemies12V2, deadEnemyRadius12-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies12V2, deadEnemyRadius12-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies12V2.X), int32(deadEnemies12V2.Y), deadEnemyRadius12-6, rl.Black)
					rl.DrawCircleV(deadEnemies13V2, deadEnemyRadius13-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies13V2, deadEnemyRadius13-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies13V2.X), int32(deadEnemies13V2.Y), deadEnemyRadius13-6, rl.Black)
					rl.DrawCircleV(deadEnemies14V2, deadEnemyRadius14-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies14V2, deadEnemyRadius14-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies14V2.X), int32(deadEnemies14V2.Y), deadEnemyRadius14-6, rl.Black)
					rl.DrawCircleV(deadEnemies15V2, deadEnemyRadius15-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies15V2, deadEnemyRadius15-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies15V2.X), int32(deadEnemies15V2.Y), deadEnemyRadius15-6, rl.Black)

				}

				if screenShake {

					camera.Zoom = 2.3

					if screenShake2 == false {

						camera.Target.X = 60.0
						camera.Rotation = -3.0
						camera.Target.Y = 60.0
						screenShake2 = true
						screenShake3 = true
					}
					if screenShake3 {
						screenShakeTimer--

						if screenShakeTimer == 0 {
							camera.Target.X = 5.0
							camera.Rotation = 3.0
							screenShakeTimer = 4
							screenShake3 = false
							screenShake4 = true

						}

					}
					if screenShake4 {
						screenShakeTimer--
						if screenShakeTimer == 0 {
							camera.Target.X = 0.0
							camera.Target.Y = 0.0
							camera.Rotation = 0.0
							screenShakeTimer = 4
							camera.Zoom = 2.0
							screenShake4 = false
							screenShake = false
							screenShake2 = false

						}

					}

				}
			}
			// draw pixel noise
			if frameCountGameStart%2 == 0 {
				cPIXELNOISE()
			}

			lineCountPixelNoise := 0
			pixelNoiseY := int32(0)
			pixelNoiseX := int32(0)
			for a := 0; a < 880; a++ {

				if pixelNoiseMAP[a] == true {
					rl.DrawRectangle(pixelNoiseX, pixelNoiseY, 2, 2, rl.Black)
				}

				lineCountPixelNoise += 34
				pixelNoiseX += 34
				if lineCountPixelNoise > 1350 {
					lineCountPixelNoise = 0
					pixelNoiseX = 0
					pixelNoiseY += 34

				}
			}
			// draw scan lines
			linesY := int32(0)
			for a := 0; a < int(screenH); a++ {
				rl.DrawLine(0, linesY, screenW, linesY, rl.Fade(rl.Black, 0.5))
				linesY += 2
				a++
			}

			// noise lines
			if frameCountGameStart%60 == 0 {
				if noiseLinesScreenOn {
					noiseLinesScreenOn = false
				} else {
					noiseLinesScreenOn = true
				}
			}

			if noiseLinesScreenOn {
				for a := 0; a < noiseLineDistance1; a++ {
					noiseLineX1Change++
				}
				for a := 0; a < noiseLineDistance2; a++ {
					noiseLineX2Change++
				}
				for a := 0; a < noiseLineDistance3; a++ {
					noiseLineX3Change++
				}
				for a := 0; a < noiseLineDistance4; a++ {
					noiseLineX4Change++
				}
				if noiseLineLR1 {
					rl.DrawLine(noiseLineX1+noiseLineX1Change, 0, noiseLineX1+noiseLineX1Change, screenH, rl.Black)
				} else {
					rl.DrawLine(noiseLineX1-noiseLineX1Change, 0, noiseLineX1-noiseLineX1Change, screenH, rl.Black)
				}
				if noiseLineLR2 {
					rl.DrawLine(noiseLineX2+noiseLineX2Change, 0, noiseLineX1+noiseLineX2Change, screenH, rl.Black)
				} else {
					rl.DrawLine(noiseLineX2-noiseLineX2Change, 0, noiseLineX2-noiseLineX2Change, screenH, rl.Black)
				}
				if noiseLineLR3 {
					rl.DrawLine(noiseLineX3+noiseLineX3Change, 0, noiseLineX3+noiseLineX3Change, screenH, rl.Black)
				} else {
					rl.DrawLine(noiseLineX3-noiseLineX3Change, 0, noiseLineX3-noiseLineX3Change, screenH, rl.Black)
				}
				if noiseLineLR4 {
					rl.DrawLine(noiseLineX4+noiseLineX4Change, 0, noiseLineX4+noiseLineX4Change, screenH, rl.Black)
				} else {
					rl.DrawLine(noiseLineX4-noiseLineX4Change, 0, noiseLineX4-noiseLineX4Change, screenH, rl.Black)
				}

			} else {
				cNOISELINES()
				noiseLineX1Change = 0
				noiseLineX2Change = 0
				noiseLineX3Change = 0
				noiseLineX4Change = 0
				noiseLineX1 = noiseLinesMAP[0]
				noiseLineX2 = noiseLinesMAP[1]
				noiseLineX3 = noiseLinesMAP[2]
				noiseLineX4 = noiseLinesMAP[3]
				noiseLineDistance1 = noiseLinesDistanceMAP[0]
				noiseLineDistance2 = noiseLinesDistanceMAP[1]
				noiseLineDistance3 = noiseLinesDistanceMAP[2]
				noiseLineDistance4 = noiseLinesDistanceMAP[3]
				noiseLineLR1 = noiseLinesLRMAP[0]
				noiseLineLR2 = noiseLinesLRMAP[1]
				noiseLineLR3 = noiseLinesLRMAP[2]
				noiseLineLR4 = noiseLinesLRMAP[3]
			}

		} // end introScreen

		if gameStartOn {
			chooseCharacter = 1
			introStoryScreenOn = true
			gameStartOn = false
		}

		// MARK: update background circle radius
		if introScreenOn == true || levelEnd == true || helpOn == true || menuOn == true || gameMenuOn == true {
			if createIntroCircles == false {

				for a := 0; a < 30; a++ {
					b := rInt32(40, int(screenW-40))
					c := rInt32(40, int(screenH-40))
					d := rFloat32(20, 50)

					introScreenCircleXMAP[a] = b
					introScreenCircleYMAP[a] = c
					introScreenCircleRadius[a] = d
					createIntroCircles = true

				}

			}
		}
		// MARK: choose character
		if chooseCharacterIntro {

			pauseOn = true
			rl.DrawRectangle(0, 0, screenW, screenH, rl.Black)
			if rl.IsKeyPressed(rl.KeySpace) {
				switch chooseCharacter {
				case 1:
					dinoType = "greenDino"
				case 2:
					dinoType = "redDino"
				case 3:
					dinoType = "yellowDino"
				case 4:
					dinoType = "blueDino"
				}
				startIntroMusic = false
				introScreenOn = false
				introStoryScreenOn = false
				chooseCharacterIntro = false
				pauseOn = false
			}
			if rl.IsKeyPressed(rl.KeyLeft) {
				chooseCharacter--
				if chooseCharacter == 0 {
					chooseCharacter = 4
				}
			}
			if rl.IsKeyPressed(rl.KeyRight) {
				chooseCharacter++
				if chooseCharacter == 5 {
					chooseCharacter = 1
				}
			}
			chooseCharacterRecX := int32((screenW / 2) - 440)
			for a := 0; a < 4; a++ {
				rl.DrawRectangle(chooseCharacterRecX, 100, 200, 200, rl.Fade(rl.White, 0.1))
				rl.DrawRectangle(chooseCharacterRecX, 100, 200, 200, rl.Fade(rl.Black, 0.7))
				chooseCharacterRecX += 220
			}
			// dino images 6X zoom
			rl.BeginMode2D(cameraIntro)
			dinoGreenIntroV2 := rl.NewVector2(50, 24)
			dinoRedIntroV2 := rl.NewVector2(85, 24)
			dinoYellowIntroV2 := rl.NewVector2(123, 24)
			dinoBlueIntroV2 := rl.NewVector2(160, 24)

			rl.DrawTextureRec(imgs, dinoGreenRIMG, dinoGreenIntroV2, rl.White)
			rl.DrawTextureRec(imgs, dinoRedRIMG, dinoRedIntroV2, rl.White)
			rl.DrawTextureRec(imgs, dinoYellowRIMG, dinoYellowIntroV2, rl.White)
			rl.DrawTextureRec(imgs, dinoBlueRIMG, dinoBlueIntroV2, rl.White)

			if frameCountGameStart%8 == 0 {
				if chooseCharacter == 1 {
					dinoGreenRIMG.X += 24
					if dinoGreenRIMG.X >= 100 {
						dinoGreenRIMG.X = 4
					}
				} else if chooseCharacter == 2 {
					dinoRedRIMG.X += 24
					if dinoRedRIMG.X >= 100 {
						dinoRedRIMG.X = 4
					}
				} else if chooseCharacter == 3 {
					dinoYellowRIMG.X += 24
					if dinoYellowRIMG.X >= 100 {
						dinoYellowRIMG.X = 4
					}
				} else if chooseCharacter == 4 {
					dinoBlueRIMG.X += 24
					if dinoBlueRIMG.X >= 100 {
						dinoBlueRIMG.X = 4
					}
				}
			}

			rl.EndMode2D() // end dino images 6X zoom

			selectRecX := int32((screenW / 2) - 448)

			switch chooseCharacter {

			case 1:
				rl.DrawRectangle(selectRecX, 92, 216, 8, rl.Red)                      // top select rec
				rl.DrawRectangle(selectRecX, 300, 216, 8, rl.Red)                     // bottom select rec
				rl.DrawRectangle(selectRecX, 100, 8, 200, rl.Red)                     // left select rec
				rl.DrawRectangle(selectRecX+208, 100, 8, 200, rl.Red)                 // right select rec
				rl.DrawRectangle(selectRecX, 92, 216, 8, rl.Fade(rl.Black, 0.5))      // top shadow select rec
				rl.DrawRectangle(selectRecX, 300, 216, 8, rl.Fade(rl.Black, 0.5))     // bottom shadow select rec
				rl.DrawRectangle(selectRecX, 100, 8, 200, rl.Fade(rl.Black, 0.5))     // left shadow select rec
				rl.DrawRectangle(selectRecX+208, 100, 8, 200, rl.Fade(rl.Black, 0.5)) // right shadow select rec

				rl.DrawText("Genie", selectRecX+60, 320, 40, rl.White)

				// shadow other dino boxes
				rl.DrawRectangle(selectRecX+228, 100, 200, 200, rl.Fade(rl.Black, 0.7))
				rl.DrawRectangle(selectRecX+448, 100, 200, 200, rl.Fade(rl.Black, 0.7))
				rl.DrawRectangle(selectRecX+668, 100, 200, 200, rl.Fade(rl.Black, 0.7))

			case 2:
				rl.DrawRectangle(selectRecX+220, 92, 216, 8, rl.Red)                      // top select rec
				rl.DrawRectangle(selectRecX+220, 300, 216, 8, rl.Red)                     // bottom select rec
				rl.DrawRectangle(selectRecX+220, 100, 8, 200, rl.Red)                     // left select rec
				rl.DrawRectangle(selectRecX+208+220, 100, 8, 200, rl.Red)                 // right select rec
				rl.DrawRectangle(selectRecX+220, 92, 216, 8, rl.Fade(rl.Black, 0.5))      // top shadow select rec
				rl.DrawRectangle(selectRecX+220, 300, 216, 8, rl.Fade(rl.Black, 0.5))     // bottom shadow select rec
				rl.DrawRectangle(selectRecX+220, 100, 8, 200, rl.Fade(rl.Black, 0.5))     // left shadow select rec
				rl.DrawRectangle(selectRecX+208+220, 100, 8, 200, rl.Fade(rl.Black, 0.5)) // right shadow select rec

				rl.DrawText("Fred", selectRecX+280, 320, 40, rl.White)

				// shadow other dino boxes
				rl.DrawRectangle(selectRecX+8, 100, 200, 200, rl.Fade(rl.Black, 0.7))
				rl.DrawRectangle(selectRecX+448, 100, 200, 200, rl.Fade(rl.Black, 0.7))
				rl.DrawRectangle(selectRecX+668, 100, 200, 200, rl.Fade(rl.Black, 0.7))

			case 3:
				rl.DrawRectangle(selectRecX+440, 92, 216, 8, rl.Red)                      // top select rec
				rl.DrawRectangle(selectRecX+440, 300, 216, 8, rl.Red)                     // bottom select rec
				rl.DrawRectangle(selectRecX+440, 100, 8, 200, rl.Red)                     // left select rec
				rl.DrawRectangle(selectRecX+208+440, 100, 8, 200, rl.Red)                 // right select rec
				rl.DrawRectangle(selectRecX+440, 92, 216, 8, rl.Fade(rl.Black, 0.5))      // top shadow select rec
				rl.DrawRectangle(selectRecX+440, 300, 216, 8, rl.Fade(rl.Black, 0.5))     // bottom shadow select rec
				rl.DrawRectangle(selectRecX+440, 100, 8, 200, rl.Fade(rl.Black, 0.5))     // left shadow select rec
				rl.DrawRectangle(selectRecX+208+440, 100, 8, 200, rl.Fade(rl.Black, 0.5)) // right shadow select rec

				rl.DrawText("Othello", selectRecX+480, 320, 40, rl.White)

				// shadow other dino boxes
				rl.DrawRectangle(selectRecX+8, 100, 200, 200, rl.Fade(rl.Black, 0.7))
				rl.DrawRectangle(selectRecX+228, 100, 200, 200, rl.Fade(rl.Black, 0.7))
				rl.DrawRectangle(selectRecX+668, 100, 200, 200, rl.Fade(rl.Black, 0.7))

			case 4:
				rl.DrawRectangle(selectRecX+660, 92, 216, 8, rl.Red)                      // top select rec
				rl.DrawRectangle(selectRecX+660, 300, 216, 8, rl.Red)                     // bottom select rec
				rl.DrawRectangle(selectRecX+660, 100, 8, 200, rl.Red)                     // left select rec
				rl.DrawRectangle(selectRecX+208+660, 100, 8, 200, rl.Red)                 // right select rec
				rl.DrawRectangle(selectRecX+660, 92, 216, 8, rl.Fade(rl.Black, 0.5))      // top shadow select rec
				rl.DrawRectangle(selectRecX+660, 300, 216, 8, rl.Fade(rl.Black, 0.5))     // bottom shadow select rec
				rl.DrawRectangle(selectRecX+660, 100, 8, 200, rl.Fade(rl.Black, 0.5))     // left shadow select rec
				rl.DrawRectangle(selectRecX+208+660, 100, 8, 200, rl.Fade(rl.Black, 0.5)) // right shadow select rec

				rl.DrawText("Louie", selectRecX+720, 320, 40, rl.White)

				// shadow other dino boxes
				rl.DrawRectangle(selectRecX+8, 100, 200, 200, rl.Fade(rl.Black, 0.7))
				rl.DrawRectangle(selectRecX+228, 100, 200, 200, rl.Fade(rl.Black, 0.7))
				rl.DrawRectangle(selectRecX+448, 100, 200, 200, rl.Fade(rl.Black, 0.7))

			}
			// red background

			rl.DrawRectangle(0, screenH-224, screenW, 108, rl.Orange)
			rl.DrawRectangle(0, screenH-224, screenW, 108, rl.Fade(rl.Black, borderFade))

			if frameCountGameStart%2 == 0 {

				if borderFade >= 1.0 {
					borderFadeOn = false
				} else if borderFade <= 0.0 {
					borderFadeOn = true
				}

				if borderFadeOn {
					borderFade += 0.01
				} else {
					borderFade -= 0.01

				}

			}

			rl.DrawRectangle(0, screenH-220, screenW, 100, rl.Red)
			rl.DrawRectangle(0, screenH-220, screenW, 100, rl.Fade(rl.Black, 0.5))

			rl.DrawText("choose your dino", 438, screenH-197, 60, rl.Black)
			rl.DrawText("choose your dino", 440, screenH-199, 60, rl.Yellow)
			rl.DrawText("choose your dino", 441, screenH-200, 60, rl.White)

			rl.DrawRectangle(movingRecChooseDinoX, screenH-220, 10, 100, rl.Red)
			rl.DrawRectangle(movingRecChooseDinoX, screenH-220, 10, 100, rl.Fade(rl.Black, 0.5))

			if leftRightChooseDinoRec == false {
				if movingRecChooseDinoX <= 440 {
					leftRightChooseDinoRec = true
				}
			}
			if leftRightChooseDinoRec {
				if movingRecChooseDinoX >= screenW-370 {
					leftRightChooseDinoRec = false
				}
			}
			if leftRightChooseDinoRec {
				if frameCountGameStart%4 == 0 {
					movingRecChooseDinoX += 10
				}
			}

			if leftRightChooseDinoRec == false {
				if frameCountGameStart%4 == 0 {
					movingRecChooseDinoX -= 10
				}

			}

			// noise lines

			if frameCountGameStart%60 == 0 {
				if noiseLinesScreenOn {
					noiseLinesScreenOn = false
				} else {
					noiseLinesScreenOn = true
				}
			}

			if noiseLinesScreenOn {
				for a := 0; a < noiseLineDistance1; a++ {
					noiseLineX1Change++
				}
				for a := 0; a < noiseLineDistance2; a++ {
					noiseLineX2Change++
				}
				for a := 0; a < noiseLineDistance3; a++ {
					noiseLineX3Change++
				}
				for a := 0; a < noiseLineDistance4; a++ {
					noiseLineX4Change++
				}
				if noiseLineLR1 {
					rl.DrawLine(noiseLineX1+noiseLineX1Change, 0, noiseLineX1+noiseLineX1Change, screenH, rl.Black)
				} else {
					rl.DrawLine(noiseLineX1-noiseLineX1Change, 0, noiseLineX1-noiseLineX1Change, screenH, rl.Black)
				}
				if noiseLineLR2 {
					rl.DrawLine(noiseLineX2+noiseLineX2Change, 0, noiseLineX1+noiseLineX2Change, screenH, rl.Black)
				} else {
					rl.DrawLine(noiseLineX2-noiseLineX2Change, 0, noiseLineX2-noiseLineX2Change, screenH, rl.Black)
				}
				if noiseLineLR3 {
					rl.DrawLine(noiseLineX3+noiseLineX3Change, 0, noiseLineX3+noiseLineX3Change, screenH, rl.Black)
				} else {
					rl.DrawLine(noiseLineX3-noiseLineX3Change, 0, noiseLineX3-noiseLineX3Change, screenH, rl.Black)
				}
				if noiseLineLR4 {
					rl.DrawLine(noiseLineX4+noiseLineX4Change, 0, noiseLineX4+noiseLineX4Change, screenH, rl.Black)
				} else {
					rl.DrawLine(noiseLineX4-noiseLineX4Change, 0, noiseLineX4-noiseLineX4Change, screenH, rl.Black)
				}

			} else {
				cNOISELINES()

				noiseLineX1Change = 0
				noiseLineX2Change = 0
				noiseLineX3Change = 0
				noiseLineX4Change = 0
				noiseLineX1 = noiseLinesMAP[0]
				noiseLineX2 = noiseLinesMAP[1]
				noiseLineX3 = noiseLinesMAP[2]
				noiseLineX4 = noiseLinesMAP[3]
				noiseLineDistance1 = noiseLinesDistanceMAP[0]
				noiseLineDistance2 = noiseLinesDistanceMAP[1]
				noiseLineDistance3 = noiseLinesDistanceMAP[2]
				noiseLineDistance4 = noiseLinesDistanceMAP[3]
				noiseLineLR1 = noiseLinesLRMAP[0]
				noiseLineLR2 = noiseLinesLRMAP[1]
				noiseLineLR3 = noiseLinesLRMAP[2]
				noiseLineLR4 = noiseLinesLRMAP[3]

			}

			// draw pixel noise

			if frameCountGameStart%2 == 0 {
				cPIXELNOISE()
			}

			lineCountPixelNoise := 0
			pixelNoiseY := int32(0)
			pixelNoiseX := int32(0)
			for a := 0; a < 880; a++ {

				if pixelNoiseMAP[a] == true {
					rl.DrawRectangle(pixelNoiseX, pixelNoiseY, 2, 2, rl.Black)
				}

				lineCountPixelNoise += 34
				pixelNoiseX += 34
				if lineCountPixelNoise > 1350 {
					lineCountPixelNoise = 0
					pixelNoiseX = 0
					pixelNoiseY += 34

				}
			}

			// draw scan lines

			if switchScanLines {
				linesY := int32(0)
				for a := 0; a < int(screenH); a++ {
					rl.DrawLine(0, linesY, screenW, linesY, rl.Fade(rl.Black, 0.5))
					linesY += 2
					a++
				}
			} else {
				linesY := int32(1)
				for a := 0; a < int(screenH); a++ {
					rl.DrawLine(0, linesY, screenW, linesY, rl.Fade(rl.Black, 0.5))
					linesY += 2
					a++
				}

			}

		}

		// MARK: intro screen
		if introScreenOn {
			pauseOn = true
			startIntroMusic = true
			if introPauseOff == false {
				if frameCountGameStart%30 == 0 {
					introPauseOff = true
				}
			}

			if introPauseOff {
				if rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyLeftControl) || rl.IsKeyPressed(rl.KeyRightControl) {
					switch chooseMenuOption {
					case 1:
						introScreenOn = false
						gameStartOn = false
						chooseCharacterIntro = true
					case 2:
						menuOn = true
					case 3:
						helpOn = true
					case 4:
						creditsScreenOn = true
					case 5:
						os.Exit(0)
					}
				}
			}

			if rl.IsKeyPressed(rl.KeyDown) {
				chooseMenuOption++
				dinoSelectMarkV2.Y += 25
				if dinoSelectMarkV2.Y > float32(screenH/2+251) {
					dinoSelectMarkV2.Y = float32(screenH/2 + 151)
				}
				if chooseMenuOption == 6 {
					chooseMenuOption = 1
				}
			}
			if rl.IsKeyPressed(rl.KeyUp) {
				chooseMenuOption--
				dinoSelectMarkV2.Y -= 25
				if dinoSelectMarkV2.Y < float32(screenH/2+151) {
					dinoSelectMarkV2.Y = float32(screenH/2 + 251)
				}
				if chooseMenuOption == 0 {
					chooseMenuOption = 5
				}
			}

			rl.DrawRectangle(0, 0, screenW, screenH, rl.Black)
			rl.DrawText("spl@ta", screenW/4+62, screenH/3+3, 160, rl.DarkGray)
			rl.DrawText("spl@ta", screenW/4+64, screenH/3+1, 160, rl.Black)
			rl.DrawText("spl@ta", screenW/4+65, screenH/3, 160, rl.White)

			if frameCountGameStart%15 == 0 {
				dinoTextColor = colorsMAP[rInt(0, 10)]
			}
			colorTxt1, colorTxt2, colorTxt3, colorTxt4, colorTxt5 := rl.White, rl.White, rl.White, rl.White, rl.White

			switch chooseMenuOption {

			case 1:
				colorTxt1 = rl.Green
			case 2:
				colorTxt2 = rl.Green
			case 3:
				colorTxt3 = rl.Green
			case 4:
				colorTxt4 = rl.Green
			case 5:
				colorTxt5 = rl.Green
			}

			rl.DrawText("dinosaurz with gunz", screenW/3+15, screenH/2+53, 40, rl.Yellow)
			rl.DrawText("dinosaurz with gunz", screenW/3+17, screenH/2+51, 40, rl.Black)
			rl.DrawText("dinosaurz with gunz", screenW/3+18, screenH/2+50, 40, dinoTextColor)
			rl.DrawText("new game", screenW/2-70, screenH/2+150, 20, colorTxt1)
			rl.DrawText("options", screenW/2-70, screenH/2+175, 20, colorTxt2)
			rl.DrawText("help", screenW/2-70, screenH/2+200, 20, colorTxt3)
			rl.DrawText("credits", screenW/2-70, screenH/2+225, 20, colorTxt4)
			rl.DrawText("exit", screenW/2-70, screenH/2+250, 20, colorTxt5)
			rl.DrawText("©", screenW-261, screenH-47, 30, rl.White)
			rl.DrawText("2020 nicholasimon", screenW-230, screenH-40, 20, rl.White)

			// lines through text
			rl.DrawRectangle(screenW/3+10, screenH/2+64, 500, 2, rl.Black)
			rl.DrawRectangle(screenW/3+10, screenH/2+68, 500, 2, rl.Black)
			rl.DrawRectangle(screenW/3+10, screenH/2+72, 500, 2, rl.Black)
			rl.DrawRectangle(screenW/3+10, screenH/2+76, 500, 2, rl.Black)
			rl.DrawRectangle(screenW/3+10, screenH/2+80, 500, 2, rl.Black)

			rl.DrawRectangle(screenW/3+10, screenH/2+40, 500, 60, rl.Fade(rl.Black, fadeAmount))

			if frameCountGameStart%14 == 0 {

				fadeAmount -= 0.2
				if fadeAmount <= 0.1 {
					fadeAmount = 1.0
				}
			}

			rl.DrawTextureRec(imgs, dinoGreenRIMG, dinoSelectMarkV2, rl.White)

			if frameCountGameStart%9 == 0 {
				dinoGreenRIMG.X += 24
				if dinoGreenRIMG.X > 52 {
					dinoGreenRIMG.X = 4
				}
			}

			// draw splatter circles
			for a := 0; a < 30; a++ {
				introCircleX := introScreenCircleXMAP[a]
				introCircleY := introScreenCircleYMAP[a]
				introCircleRadius := introScreenCircleRadius[a]

				rl.DrawCircle(introCircleX, introCircleY, introCircleRadius, rl.Fade(rl.Red, 0.4))

			}

			// noise lines

			if frameCountGameStart%60 == 0 {
				if noiseLinesScreenOn {
					noiseLinesScreenOn = false
				} else {
					noiseLinesScreenOn = true
				}
			}

			if noiseLinesScreenOn {
				for a := 0; a < noiseLineDistance1; a++ {
					noiseLineX1Change++
				}
				for a := 0; a < noiseLineDistance2; a++ {
					noiseLineX2Change++
				}
				for a := 0; a < noiseLineDistance3; a++ {
					noiseLineX3Change++
				}
				for a := 0; a < noiseLineDistance4; a++ {
					noiseLineX4Change++
				}
				if noiseLineLR1 {
					rl.DrawLine(noiseLineX1+noiseLineX1Change, 0, noiseLineX1+noiseLineX1Change, screenH, rl.Black)
				} else {
					rl.DrawLine(noiseLineX1-noiseLineX1Change, 0, noiseLineX1-noiseLineX1Change, screenH, rl.Black)
				}
				if noiseLineLR2 {
					rl.DrawLine(noiseLineX2+noiseLineX2Change, 0, noiseLineX1+noiseLineX2Change, screenH, rl.Black)
				} else {
					rl.DrawLine(noiseLineX2-noiseLineX2Change, 0, noiseLineX2-noiseLineX2Change, screenH, rl.Black)
				}
				if noiseLineLR3 {
					rl.DrawLine(noiseLineX3+noiseLineX3Change, 0, noiseLineX3+noiseLineX3Change, screenH, rl.Black)
				} else {
					rl.DrawLine(noiseLineX3-noiseLineX3Change, 0, noiseLineX3-noiseLineX3Change, screenH, rl.Black)
				}
				if noiseLineLR4 {
					rl.DrawLine(noiseLineX4+noiseLineX4Change, 0, noiseLineX4+noiseLineX4Change, screenH, rl.Black)
				} else {
					rl.DrawLine(noiseLineX4-noiseLineX4Change, 0, noiseLineX4-noiseLineX4Change, screenH, rl.Black)
				}

			} else {
				cNOISELINES()

				noiseLineX1Change = 0
				noiseLineX2Change = 0
				noiseLineX3Change = 0
				noiseLineX4Change = 0
				noiseLineX1 = noiseLinesMAP[0]
				noiseLineX2 = noiseLinesMAP[1]
				noiseLineX3 = noiseLinesMAP[2]
				noiseLineX4 = noiseLinesMAP[3]
				noiseLineDistance1 = noiseLinesDistanceMAP[0]
				noiseLineDistance2 = noiseLinesDistanceMAP[1]
				noiseLineDistance3 = noiseLinesDistanceMAP[2]
				noiseLineDistance4 = noiseLinesDistanceMAP[3]
				noiseLineLR1 = noiseLinesLRMAP[0]
				noiseLineLR2 = noiseLinesLRMAP[1]
				noiseLineLR3 = noiseLinesLRMAP[2]
				noiseLineLR4 = noiseLinesLRMAP[3]

			}

			// draw pixel noise
			if frameCountGameStart%2 == 0 {
				cPIXELNOISE()
			}

			lineCountPixelNoise := 0
			pixelNoiseY := int32(0)
			pixelNoiseX := int32(0)
			for a := 0; a < 880; a++ {

				if pixelNoiseMAP[a] == true {
					rl.DrawRectangle(pixelNoiseX, pixelNoiseY, 2, 2, rl.Black)
				}

				lineCountPixelNoise += 34
				pixelNoiseX += 34
				if lineCountPixelNoise > 1350 {
					lineCountPixelNoise = 0
					pixelNoiseX = 0
					pixelNoiseY += 34

				}
			}
			// draw scan lines
			if switchScanLines {
				linesY := int32(0)
				for a := 0; a < int(screenH); a++ {
					rl.DrawLine(0, linesY, screenW, linesY, rl.Fade(rl.Black, 0.5))
					linesY += 2
					a++
				}
			} else {
				linesY := int32(1)
				for a := 0; a < int(screenH); a++ {
					rl.DrawLine(0, linesY, screenW, linesY, rl.Fade(rl.Black, 0.5))
					linesY += 2
					a++
				}

			}

			if frameCountGameStart%3 == 0 {
				for a := 0; a < 30; a++ {
					introCircleRadius := introScreenCircleRadius[a]
					introCircleRadius--
					introScreenCircleRadius[a] = introCircleRadius
				}
			}

			sumIntroRadius := float32(0)
			for _, introRadius := range introScreenCircleRadius {
				sumIntroRadius += introRadius
			}

			if sumIntroRadius <= 0 {
				createIntroCircles = false
			}

		}
		// MARK: help screen
		if helpOn {

			// background rectangle
			rl.DrawRectangle(0, 0, screenW, screenH, rl.Black)
			rl.DrawText("press LEFT CTRL or F1 to close", screenW/4-2, screenH-118, 40, rl.DarkGray)
			rl.DrawText("press LEFT CTRL or F1 to close", screenW/4, screenH-120, 40, rl.White)

			if rl.IsKeyPressed(rl.KeyC) || rl.IsKeyPressed(rl.KeyF1) || rl.IsKeyPressed(rl.KeyLeftControl) {
				helpOn = false
			}
			rl.DrawText("kill as many enemies in as you can in a minute", screenW/6-2, 52, 40, rl.DarkGray)
			rl.DrawText("kill as many enemies in as you can in a minute", screenW/6, 50, 40, rl.White)
			rl.DrawText("the more killz the more bonuses", screenW/6+98, 102, 40, rl.DarkGray)
			rl.DrawText("the more killz the more bonuses", screenW/6+100, 100, 40, rl.White)
			rl.DrawText("kill 10 bosses >> WINNER >> repeat", screenW/6+76, 152, 40, rl.DarkGray)
			rl.DrawText("kill 10 bosses >> WINNER >> repeat", screenW/6+78, 150, 40, rl.White)

			// text flicker
			if frameCountGameStart%24 == 0 {
				if flickerText {
					flickerText = false
				} else {
					flickerText = true
				}
			}

			// arrow keys image
			arrowKeysV2 := rl.NewVector2(float32(screenW/4), 300)
			rl.DrawTextureRec(imgs, arrowKeysIMG, arrowKeysV2, rl.Fade(rl.White, 0.6))
			rl.DrawText("jump", screenW/4+68, 252, 40, rl.Gold)
			rl.DrawText("jump", screenW/4+69, 251, 40, rl.Black)
			rl.DrawText("jump", screenW/4+70, 250, 40, rl.Lime)
			if flickerText {
				rl.DrawText("jump", screenW/4+70, 250, 40, rl.Fade(rl.Black, 0.3))
			}
			rl.DrawText("run left", screenW/8-2, 392, 40, rl.Gold)
			rl.DrawText("run left", screenW/8-1, 391, 40, rl.Black)
			rl.DrawText("run left", screenW/8, 390, 40, rl.Lime)
			if flickerText {
				rl.DrawText("run left", screenW/8, 390, 40, rl.Fade(rl.Black, 0.3))
			}
			rl.DrawText("run right", screenW/3+123, 392, 40, rl.Gold)
			rl.DrawText("run right", screenW/3+124, 391, 40, rl.Black)
			rl.DrawText("run right", screenW/3+125, 390, 40, rl.Lime)
			if flickerText {
				rl.DrawText("run right", screenW/3+125, 390, 40, rl.Fade(rl.Black, 0.3))
			}
			rl.DrawText("drop", screenW/4+68, 462, 40, rl.Gold)
			rl.DrawText("drop", screenW/4+69, 461, 40, rl.Black)
			rl.DrawText("drop", screenW/4+70, 460, 40, rl.Lime)
			if flickerText {
				rl.DrawText("drop", screenW/4+70, 460, 40, rl.Fade(rl.Black, 0.3))
			}
			rl.DrawText("down", screenW/4+68, 502, 40, rl.Gold)
			rl.DrawText("down", screenW/4+69, 501, 40, rl.Black)
			rl.DrawText("down", screenW/4+70, 500, 40, rl.Lime)
			if flickerText {
				rl.DrawText("down", screenW/4+70, 500, 40, rl.Fade(rl.Black, 0.3))
			}
			// ctrl key image
			ctrlKeyV2 := rl.NewVector2(float32(screenW/2+200), 250)
			rl.DrawTextureRec(imgs, ctrlKeyIMG, ctrlKeyV2, rl.Fade(rl.White, 0.6))
			rl.DrawText("shoot", screenW/2+303, 262, 40, rl.Gold)
			rl.DrawText("shoot", screenW/2+304, 261, 40, rl.Black)
			rl.DrawText("shoot", screenW/2+305, 260, 40, rl.Lime)
			if flickerText {
				rl.DrawText("shoot", screenW/2+305, 260, 40, rl.Fade(rl.Black, 0.3))
			}
			// alt key image
			altKeyV2 := rl.NewVector2(float32(screenW/2+200), 330)
			rl.DrawTextureRec(imgs, altKeyIMG, altKeyV2, rl.Fade(rl.White, 0.6))
			rl.DrawText("shoot up", screenW/2+303, 342, 40, rl.Gold)
			rl.DrawText("shoot up", screenW/2+304, 341, 40, rl.Black)
			rl.DrawText("shoot up", screenW/2+305, 340, 40, rl.Lime)
			if flickerText {
				rl.DrawText("shoot up", screenW/2+305, 340, 40, rl.Fade(rl.Black, 0.3))
			}
			// f1 key image
			f1KeyV2 := rl.NewVector2(float32(screenW/2+215), 410)
			rl.DrawTextureRec(imgs, f1KeyIMG, f1KeyV2, rl.Fade(rl.White, 0.6))
			rl.DrawText("menu", screenW/2+303, 427, 40, rl.Gold)
			rl.DrawText("menu", screenW/2+304, 426, 40, rl.Black)
			rl.DrawText("menu", screenW/2+305, 425, 40, rl.Lime)
			if flickerText {
				rl.DrawText("menu", screenW/2+305, 425, 40, rl.Fade(rl.Black, 0.3))
			}
			// esc key image
			escKeyV2 := rl.NewVector2(float32(screenW/2+215), 490)
			rl.DrawTextureRec(imgs, escKeyIMG, escKeyV2, rl.Fade(rl.White, 0.6))
			rl.DrawText("exit", screenW/2+305-2, 507, 40, rl.Gold)
			rl.DrawText("exit", screenW/2+305-1, 506, 40, rl.Black)
			rl.DrawText("exit", screenW/2+305, 505, 40, rl.Lime)
			if flickerText {
				rl.DrawText("exit", screenW/2+305, 505, 40, rl.Fade(rl.Black, 0.3))
			}
			// draw splatter circles
			for a := 0; a < 30; a++ {
				introCircleX := introScreenCircleXMAP[a]
				introCircleY := introScreenCircleYMAP[a]
				introCircleRadius := introScreenCircleRadius[a]

				rl.DrawCircle(introCircleX, introCircleY, introCircleRadius, rl.Fade(rl.Red, 0.4))

			}
			sumIntroRadius := float32(0)
			for _, introRadius := range introScreenCircleRadius {
				sumIntroRadius += introRadius
			}
			if sumIntroRadius <= 0 {
				createIntroCircles = false
			}
		}

		// MARK: options menu
		if menuOn || gameMenuOn {

			// background rectangle
			rl.DrawRectangle(0, 0, screenW, screenH, rl.Black)

			rl.DrawText("if you experience performance issues disable weather - it does not effect gameplay", 240, screenH-190, 20, rl.White)

			rl.DrawText("space to select", 510, screenH-147, 40, rl.DarkGray)
			rl.DrawText("space to select", 512, screenH-149, 40, rl.Black)
			rl.DrawText("space to select", 513, screenH-150, 40, menuTextSpaceColor)

			rl.DrawText("f1 to close", 570, screenH-97, 40, rl.DarkGray)
			rl.DrawText("f1 to close", 572, screenH-99, 40, rl.Black)
			rl.DrawText("f1 to close", 573, screenH-100, 40, menuTextSpaceColor)

			if frameCountGameStart%15 == 0 {
				if menuTextSpaceColor == rl.Yellow {
					menuTextSpaceColor = rl.Gold
				} else {
					menuTextSpaceColor = rl.Yellow
				}
			}

			if rl.IsKeyPressed(rl.KeyF1) {
				menuOn = false
			}

			if rl.IsKeyPressed(rl.KeyLeft) {
				menuColumnSelect--
				if menuColumnSelect == 0 {
					menuColumnSelect = 3
				}
				menuSelectBoxY = int32(95)
				menuSelectNumber = 1
			}
			if rl.IsKeyPressed(rl.KeyRight) {
				menuColumnSelect++
				if menuColumnSelect == 4 {
					menuColumnSelect = 1
				}
				menuSelectBoxY = int32(95)
				menuSelectNumber = 1
			}

			if rl.IsKeyPressed(rl.KeyDown) {
				if menuColumnSelect == 1 {
					if menuSelectBoxY < 270 {
						menuSelectBoxY += 30
						menuSelectNumber++
					} else if menuSelectBoxY >= 270 && menuSelectBoxY < 280 {
						menuSelectBoxY += 80
						menuSelectNumber++
					} else if menuSelectBoxY >= 280 && menuSelectBoxY < 500 {
						menuSelectBoxY += 30
						menuSelectNumber++
					}
				} else if menuColumnSelect == 2 {
					if menuSelectBoxY < 180 {
						menuSelectBoxY += 30
						menuSelectNumber++
					} else if menuSelectBoxY >= 180 && menuSelectBoxY < 190 {
						menuSelectBoxY += 80
						menuSelectNumber++
					} else if menuSelectBoxY >= 190 && menuSelectBoxY < 380 {
						menuSelectBoxY += 30
						menuSelectNumber++
					}
				} else if menuColumnSelect == 3 {
					if menuSelectBoxY < 150 {
						menuSelectBoxY += 30
						menuSelectNumber++
					} else if menuSelectBoxY >= 150 && menuSelectBoxY < 160 {
						menuSelectBoxY += 80
						menuSelectNumber++
					} else if menuSelectBoxY >= 230 && menuSelectBoxY < 260 {
						menuSelectBoxY += 30
						menuSelectNumber++
					} else if menuSelectBoxY >= 260 && menuSelectBoxY < 270 {
						menuSelectBoxY += 80
						menuSelectNumber++
					} else if menuSelectBoxY >= 270 && menuSelectBoxY < 380 {
						menuSelectBoxY += 30
						menuSelectNumber++
					}
				}

			}
			if rl.IsKeyPressed(rl.KeyUp) {
				if menuColumnSelect == 1 {
					if menuSelectBoxY > 95 && menuSelectBoxY < 355 {
						menuSelectBoxY -= 30
						menuSelectNumber--
					} else if menuSelectBoxY == 355 {
						menuSelectBoxY -= 80
						menuSelectNumber--
					} else if menuSelectBoxY > 355 {
						menuSelectBoxY -= 30
						menuSelectNumber--
					}
				} else if menuColumnSelect == 2 {
					if menuSelectBoxY > 95 && menuSelectBoxY < 265 {
						menuSelectBoxY -= 30
						menuSelectNumber--
					} else if menuSelectBoxY == 265 {
						menuSelectBoxY -= 80
						menuSelectNumber--

					} else if menuSelectBoxY > 265 {
						menuSelectBoxY -= 30
						menuSelectNumber--
					}
				} else if menuColumnSelect == 3 {
					if menuSelectBoxY > 95 && menuSelectBoxY < 235 {
						menuSelectBoxY -= 30
						menuSelectNumber--
					} else if menuSelectBoxY == 235 {
						menuSelectBoxY -= 80
						menuSelectNumber--
					} else if menuSelectBoxY > 235 && menuSelectBoxY < 345 {
						menuSelectBoxY -= 30
						menuSelectNumber--
					} else if menuSelectBoxY == 345 {
						menuSelectBoxY -= 80
						menuSelectNumber--
					} else if menuSelectBoxY > 345 {
						menuSelectBoxY -= 30
						menuSelectNumber--
					}

				}

			}
			if rl.IsKeyPressed(rl.KeySpace) {

				if menuColumnSelect == 1 {
					switch menuSelectNumber {
					case 1:
						if standardBackOn {
							standardBackOn = false
							if randomBackOn {
								randomBackOn = false
							}
							changeBackgroundImage()
						} else {
							standardBackOn = true
							changeBackgroundImage()
						}
					case 2:
						if altColorsOn {
							altColorsOn = false
							if randomBackOn {
								randomBackOn = false
							}
							changeBackgroundImage()
						} else {
							altColorsOn = true
							changeBackgroundImage()
						}
					case 3:
						if grayscaleOn {
							grayscaleOn = false
							if randomBackOn {
								randomBackOn = false
							}
							changeBackgroundImage()
						} else {
							grayscaleOn = true
							changeBackgroundImage()
						}
					case 4:
						if pencilOn {
							pencilOn = false
							if randomBackOn {
								randomBackOn = false
							}
							changeBackgroundImage()
						} else {
							pencilOn = true
							changeBackgroundImage()
						}
					case 5:
						if inkOn {
							inkOn = false
							if randomBackOn {
								randomBackOn = false
							}
							changeBackgroundImage()
						} else {
							inkOn = true
							changeBackgroundImage()
						}
					case 6:
						if chalkOn {
							chalkOn = false
							if randomBackOn {
								randomBackOn = false
							}
							changeBackgroundImage()
						} else {
							chalkOn = true
							changeBackgroundImage()
						}
					case 7:
						if randomBackOn {
							randomBackOn = false
							standardBackOn = true
							altColorsOn = false
							chalkOn = false
							grayscaleOn = false
							inkOn = false
							pencilOn = false
							changeBackgroundImage()
						} else {
							randomBackOn = true
							altColorsOn = true
							chalkOn = true
							standardBackOn = true
							grayscaleOn = true
							inkOn = true
							pencilOn = true
							changeBackgroundImage()
						}
					case 8:
						if weatherOn {
							weatherOn = false
							cloudsactive = false
							fogactive = false
							rainactive = false
							snowactive = false
							frograinactive = false
						} else {
							weatherOn = true
							cloudsactive = true
							fogactive = true
							rainactive = true
							snowactive = true
							frograinactive = true
						}
					case 9:
						if cloudsactive {
							cloudsactive = false
						} else {
							cloudsactive = true
						}
					case 10:
						if fogactive {
							fogactive = false
						} else {
							fogactive = true
						}
					case 11:
						if rainactive {
							rainactive = false
						} else {
							rainactive = true
						}
					case 12:
						if snowactive {
							snowactive = false
						} else {
							snowactive = true
						}
					case 13:
						if frograinactive {
							frograinactive = false
						} else {
							frograinactive = true
						}
					}
					if standardBackOn == true && altColorsOn == true && grayscaleOn == true && pencilOn == true && inkOn == true && chalkOn == true && randomBackOn == false {
						randomBackOn = true
						changeBackgroundImage()
					}
				} else if menuColumnSelect == 2 {

					switch menuSelectNumber {
					case 1:
						if disastersOn {
							disastersOn = false
							earthquakeswitch = false
							tornadoswitch = false
							meteorswitch = false
						} else {
							disastersOn = true
							earthquakeswitch = true
							tornadoswitch = true
							meteorswitch = true
						}
					case 2:
						if earthquakeswitch {
							earthquakeswitch = false
						} else {
							earthquakeswitch = true
						}
					case 3:
						if tornadoswitch {
							tornadoswitch = false
						} else {
							tornadoswitch = true
						}
					case 4:
						if meteorswitch {
							meteorswitch = false
						} else {
							meteorswitch = true
						}
					case 5:
						if magicOn {
							magicOn = false
							teleportswitch = false
							rainoffireswitch = false
							shockblockswitch = false
							distortionswitch = false
						} else {
							magicOn = true
							teleportswitch = true
							rainoffireswitch = true
							shockblockswitch = true
							distortionswitch = true
						}
					case 6:
						if teleportswitch {
							teleportswitch = false
						} else {
							teleportswitch = true
						}
					case 7:
						if rainoffireswitch {
							rainoffireswitch = false
						} else {
							rainoffireswitch = true
						}
					case 8:
						if shockblockswitch {
							shockblockswitch = false
						} else {
							shockblockswitch = true
						}
					case 9:
						if distortionswitch {
							distortionswitch = false
						} else {
							distortionswitch = true
						}
					}
				} else if menuColumnSelect == 3 {

					switch menuSelectNumber {
					case 1:
						if scanLinesOn {
							scanLinesOn = false
						} else {
							scanLinesOn = true
						}
					case 2:
						if pixelNoiseOn {
							pixelNoiseOn = false
						} else {
							pixelNoiseOn = true
						}
					case 3:
						if noiseLinesOn {
							noiseLinesOn = false
						} else {
							noiseLinesOn = true
						}
					case 4:
						if musicOn {
							musicOn = false
						} else {
							musicOn = true
						}
					case 5:
						if fxOn {
							fxOn = false
						} else {
							fxOn = true
						}
					case 7:
						if easyDiffOn {
							easyDiffOn = false
							averageDiffOn = true
							difficultDiffOn = false
							changediff()
							chooseweather()
						} else {
							easyDiffOn = true
							averageDiffOn = false
							difficultDiffOn = false
							changediff()
							chooseweather()
						}
					case 8:
						if averageDiffOn {
							averageDiffOn = false
							difficultDiffOn = true
							easyDiffOn = false
							changediff()
							chooseweather()
						} else {
							averageDiffOn = true
							easyDiffOn = false
							difficultDiffOn = false
							changediff()
							chooseweather()
						}
					case 9:
						if difficultDiffOn {
							difficultDiffOn = false
							averageDiffOn = false
							easyDiffOn = true
							changediff()
							chooseweather()
						} else {
							difficultDiffOn = true
							averageDiffOn = false
							easyDiffOn = false
							changediff()
							chooseweather()
						}
					}
				}

			}

			// correct duplicates / not on switches
			if standardBackOn && altColorsOn && grayscaleOn && pencilOn && inkOn && chalkOn {
				randomBackOn = true
			}
			if cloudsactive && fogactive && rainactive && snowactive && frograinactive {
				weatherOn = true
			}
			if earthquakeswitch && tornadoswitch && meteorswitch {
				disastersOn = true
			}
			if teleportswitch && rainoffireswitch && shockblockswitch && distortionswitch {
				magicOn = true
			}

			if difficultDiffOn {
				averageDiffOn = false
				easyDiffOn = false
			} else if averageDiffOn {
				easyDiffOn = false
				difficultDiffOn = false
			} else if easyDiffOn {
				averageDiffOn = false
				difficultDiffOn = false
			}

			// select rectange
			if menuColumnSelect == 1 {
				rl.DrawRectangle(175, menuSelectBoxY, 255, 30, rl.Fade(rl.Red, 0.4))
			}
			if menuColumnSelect == 2 {
				rl.DrawRectangle(555, menuSelectBoxY, 255, 30, rl.Fade(rl.Red, 0.4))
			}
			if menuColumnSelect == 3 {
				rl.DrawRectangle(935, menuSelectBoxY, 255, 30, rl.Fade(rl.Red, 0.4))
			}
			// left menu column text
			rl.DrawText("backgrounds", 150, 50, 40, rl.White)
			rl.DrawText("standard", 200, 100, 20, rl.White)
			rl.DrawText("alt colors", 200, 130, 20, rl.White)
			rl.DrawText("grayscale", 200, 160, 20, rl.White)
			rl.DrawText("pencil", 200, 190, 20, rl.White)
			rl.DrawText("ink", 200, 220, 20, rl.White)
			rl.DrawText("chalk", 200, 250, 20, rl.White)
			rl.DrawText("all backgrounds", 200, 280, 20, rl.White)
			rl.DrawText("weather", 150, 310, 40, rl.White)
			rl.DrawText("all weather", 200, 360, 20, rl.White)
			rl.DrawText("clouds", 200, 390, 20, rl.White)
			rl.DrawText("fog", 200, 420, 20, rl.White)
			rl.DrawText("rain", 200, 450, 20, rl.White)
			rl.DrawText("snow", 200, 480, 20, rl.White)
			rl.DrawText("rain of frogs", 200, 510, 20, rl.White)

			// left column boxes
			menuBoxX := int32(380)
			menuBoxY := int32(100)
			for a := 0; a < 7; a++ {
				rl.DrawRectangle(menuBoxX, menuBoxY, 20, 20, rl.White)
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Black)
				menuBoxY += 30
			}

			// left column boxes 2
			menuBoxX = int32(380)
			menuBoxY = int32(360)
			for a := 0; a < 6; a++ {
				rl.DrawRectangle(menuBoxX, menuBoxY, 20, 20, rl.White)
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Black)
				menuBoxY += 30
			}

			// center menu column text
			rl.DrawText("disasters", 530, 50, 40, rl.White)
			rl.DrawText("all disasters", 580, 100, 20, rl.White)
			rl.DrawText("earthquakes", 580, 130, 20, rl.White)
			rl.DrawText("tornado", 580, 160, 20, rl.White)
			rl.DrawText("meteor", 580, 190, 20, rl.White)
			rl.DrawText("magic", 530, 220, 40, rl.White)
			rl.DrawText("all magic", 580, 270, 20, rl.White)
			rl.DrawText("teleporter", 580, 300, 20, rl.White)
			rl.DrawText("rain of fire", 580, 330, 20, rl.White)
			rl.DrawText("shock blocks", 580, 360, 20, rl.White)
			rl.DrawText("distortion", 580, 390, 20, rl.White)

			// center column boxes
			menuBoxX = int32(760)
			menuBoxY = int32(100)
			for a := 0; a < 4; a++ {
				rl.DrawRectangle(menuBoxX, menuBoxY, 20, 20, rl.White)
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Black)
				menuBoxY += 30
			}

			// center column boxes 2
			menuBoxX = int32(760)
			menuBoxY = int32(270)
			for a := 0; a < 5; a++ {
				rl.DrawRectangle(menuBoxX, menuBoxY, 20, 20, rl.White)
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Black)
				menuBoxY += 30
			}

			// right menu column text
			rl.DrawText("display", 910, 50, 40, rl.White)
			rl.DrawText("scanlines", 960, 100, 20, rl.White)
			rl.DrawText("pixel noise", 960, 130, 20, rl.White)
			rl.DrawText("film lines", 960, 160, 20, rl.White)
			rl.DrawText("sound", 910, 190, 40, rl.White)
			rl.DrawText("music", 960, 240, 20, rl.White)
			rl.DrawText("sound fx", 960, 270, 20, rl.White)
			rl.DrawText("difficulty", 910, 300, 40, rl.White)
			rl.DrawText("easy", 960, 350, 20, rl.White)
			rl.DrawText("average", 960, 380, 20, rl.White)
			rl.DrawText("difficult", 960, 410, 20, rl.White)

			// right column boxes
			menuBoxX = int32(1140)
			menuBoxY = int32(100)
			for a := 0; a < 3; a++ {
				rl.DrawRectangle(menuBoxX, menuBoxY, 20, 20, rl.White)
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Black)
				menuBoxY += 30
			}

			// right column boxes 2
			menuBoxX = int32(1140)
			menuBoxY = int32(240)
			for a := 0; a < 2; a++ {
				rl.DrawRectangle(menuBoxX, menuBoxY, 20, 20, rl.White)
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Black)
				menuBoxY += 30
			}

			// right column boxes 3
			menuBoxX = int32(1140)
			menuBoxY = int32(350)
			for a := 0; a < 3; a++ {
				rl.DrawRectangle(menuBoxX, menuBoxY, 20, 20, rl.White)
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Black)
				menuBoxY += 30
			}

			menuBoxX = int32(380)
			menuBoxY = int32(100)

			if standardBackOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if altColorsOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if grayscaleOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if pencilOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if inkOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if chalkOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if randomBackOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 80
			if weatherOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if cloudsactive {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if fogactive {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if rainactive {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if snowactive {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if frograinactive {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}

			menuBoxX = int32(760)
			menuBoxY = int32(100)

			if disastersOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if earthquakeswitch {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if tornadoswitch {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if meteorswitch {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 80
			if magicOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if teleportswitch {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if rainoffireswitch {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if shockblockswitch {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if distortionswitch {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}

			menuBoxX = int32(1140)
			menuBoxY = int32(100)

			if scanLinesOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if pixelNoiseOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if noiseLinesOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 80
			if musicOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if fxOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 80
			if easyDiffOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if averageDiffOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}
			menuBoxY += 30
			if difficultDiffOn {
				rl.DrawRectangle(menuBoxX+2, menuBoxY+2, 16, 16, rl.Orange)
			}

			// draw splatter circles
			for a := 0; a < 30; a++ {
				introCircleX := introScreenCircleXMAP[a]
				introCircleY := introScreenCircleYMAP[a]
				introCircleRadius := introScreenCircleRadius[a]

				rl.DrawCircle(introCircleX, introCircleY, introCircleRadius, rl.Fade(rl.Red, 0.4))

			}

			sumIntroRadius := float32(0)
			for _, introRadius := range introScreenCircleRadius {
				sumIntroRadius += introRadius
			}

			if sumIntroRadius <= 0 {
				createIntroCircles = false
			}

		}

		// MARK: credits screen
		if creditsScreenOn {

			if rl.IsKeyPressed(rl.KeyC) || rl.IsKeyPressed(rl.KeyF1) || rl.IsKeyPressed(rl.KeyLeftControl) {
				creditsScreenOn = false
			}

			// background rectangle
			rl.DrawRectangle(0, 0, screenW, screenH, rl.Black)
			rl.DrawText("press LEFT CTRL or F1 to close", screenW/4-2, screenH-118, 40, rl.DarkGray)
			rl.DrawText("press LEFT CTRL or F1 to close", screenW/4, screenH-120, 40, rl.White)

			rl.DrawText("all assets available on itch.io", 200, 200, 20, rl.White)
			rl.DrawText("dino characters - thanks to @ScissorMarks", 200, 230, 20, rl.White)
			rl.DrawText("raylib video games library @raysan5", 200, 260, 20, rl.White)

			// draw splatter circles
			for a := 0; a < 30; a++ {
				introCircleX := introScreenCircleXMAP[a]
				introCircleY := introScreenCircleYMAP[a]
				introCircleRadius := introScreenCircleRadius[a]

				rl.DrawCircle(introCircleX, introCircleY, introCircleRadius, rl.Fade(rl.Red, 0.4))

			}

			sumIntroRadius := float32(0)
			for _, introRadius := range introScreenCircleRadius {
				sumIntroRadius += introRadius
			}

			if sumIntroRadius <= 0 {
				createIntroCircles = false
			}

		}

		// MARK: game backgrounds
		if pauseOn == false {
			backgroundV2 := rl.NewVector2(0, 0)
			rl.DrawTextureRec(backgroundTexture, backgroundIMG, backgroundV2, rl.White)
		}

		rl.BeginMode2D(camera) // MARK: BeginMode2D

		// MARK: draw screen layer 0
		drawScreenCurrentBlock = 192
		drawScreenLineCount := 0
		drawScreenX := int32(0)
		drawScreenY := int32(0)

		if pauseOn == false {
			for a := 0; a < 4128; a++ {
				checkBackgroundObjectsBlock := backgroundObjectsMAP[drawScreenCurrentBlock]

				switch checkBackgroundObjectsBlock {

				case "backObj1":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-2))
					rl.DrawTextureRec(imgs, backObj1IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj1IMG, backObjV2, rl.White)
				case "backObj2":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-16))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-18))
					rl.DrawTextureRec(imgs, backObj2IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj2IMG, backObjV2, rl.White)
				case "backObj3":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-16))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-18))
					rl.DrawTextureRec(imgs, backObj3IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj3IMG, backObjV2, rl.White)
				case "backObj4":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-10))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-12))
					rl.DrawTextureRec(imgs, backObj4IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj4IMG, backObjV2, rl.White)
				case "backObj5":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-10))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-12))
					rl.DrawTextureRec(imgs, backObj5IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj5IMG, backObjV2, rl.White)
				case "backObj6":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-5))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-7))
					rl.DrawTextureRec(imgs, backObj6IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj6IMG, backObjV2, rl.White)
				case "backObj7":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-6))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-8))
					rl.DrawTextureRec(imgs, backObj7IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj7IMG, backObjV2, rl.White)
				case "backObj8":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-5))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-7))
					rl.DrawTextureRec(imgs, backObj8IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj8IMG, backObjV2, rl.White)
				case "backObj9":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY+6))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY+4))
					rl.DrawTextureRec(imgs, backObj9IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj9IMG, backObjV2, rl.White)
				case "backObj10":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY+4))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY+2))
					rl.DrawTextureRec(imgs, backObj10IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj10IMG, backObjV2, rl.White)
				case "backObj11":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY+2))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY))
					rl.DrawTextureRec(imgs, backObj11IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj11IMG, backObjV2, rl.White)
				case "backObj12":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-2))
					rl.DrawTextureRec(imgs, backObj12IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj12IMG, backObjV2, rl.White)
				case "backObj13":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-2))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-4))
					rl.DrawTextureRec(imgs, backObj13IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj13IMG, backObjV2, rl.White)
				case "backObj14":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-2))
					rl.DrawTextureRec(imgs, backObj14IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj14IMG, backObjV2, rl.White)
				case "backObj15":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY+1))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-1))
					rl.DrawTextureRec(imgs, backObj15IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj15IMG, backObjV2, rl.White)
				case "backObj16":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY+2))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY))
					rl.DrawTextureRec(imgs, backObj16IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj16IMG, backObjV2, rl.White)
				case "backObj17":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY+2))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY))
					rl.DrawTextureRec(imgs, backObj17IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj17IMG, backObjV2, rl.White)
				case "backObj18":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-2))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-4))
					rl.DrawTextureRec(imgs, backObj18IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj18IMG, backObjV2, rl.White)
				case "backObj19":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY+2))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY))
					rl.DrawTextureRec(imgs, backObj19IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj19IMG, backObjV2, rl.White)
				case "backObj20":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY+2))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY))
					rl.DrawTextureRec(imgs, backObj20IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObj20IMG, backObjV2, rl.White)
				case "backObj21":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-60))
					rl.DrawTextureRec(imgs, backObj21IMG, backObjV2, rl.White)
				case "backObj22":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-60))
					rl.DrawTextureRec(imgs, backObj22IMG, backObjV2, rl.White)
				case "backObj23":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-60))
					rl.DrawTextureRec(imgs, backObj23IMG, backObjV2, rl.White)
				case "backObj24":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-60))
					rl.DrawTextureRec(imgs, backObj24IMG, backObjV2, rl.White)

				case "backObjGround1":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-44))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-46))
					rl.DrawTextureRec(imgs, backObjGround1IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround1IMG, backObjV2, rl.White)
				case "backObjGround2":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-44))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-46))
					rl.DrawTextureRec(imgs, backObjGround2IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround2IMG, backObjV2, rl.White)
				case "backObjGround3":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-38))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-40))
					rl.DrawTextureRec(imgs, backObjGround3IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround3IMG, backObjV2, rl.White)
				case "backObjGround4":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-48))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-50))
					rl.DrawTextureRec(imgs, backObjGround4IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround4IMG, backObjV2, rl.White)
				case "backObjGround5":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-44))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-46))
					rl.DrawTextureRec(imgs, backObjGround5IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround5IMG, backObjV2, rl.White)
				case "backObjGround6":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-30))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-32))
					rl.DrawTextureRec(imgs, backObjGround6IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround6IMG, backObjV2, rl.White)
				case "backObjGround7":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-30))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-32))
					rl.DrawTextureRec(imgs, backObjGround7IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround7IMG, backObjV2, rl.White)
				case "backObjGround8":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-12))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-14))
					rl.DrawTextureRec(imgs, backObjGround8IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround8IMG, backObjV2, rl.White)
				case "backObjGround9":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-16))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-18))
					rl.DrawTextureRec(imgs, backObjGround9IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround9IMG, backObjV2, rl.White)
				case "backObjGround10":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-24))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-26))
					rl.DrawTextureRec(imgs, backObjGround10IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround10IMG, backObjV2, rl.White)
				case "backObjGround11":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-24))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-26))
					rl.DrawTextureRec(imgs, backObjGround11IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround11IMG, backObjV2, rl.White)
				case "backObjGround12":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-50))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-52))
					rl.DrawTextureRec(imgs, backObjGround12IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround12IMG, backObjV2, rl.White)
				case "backObjGround13":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-40))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-42))
					rl.DrawTextureRec(imgs, backObjGround13IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround13IMG, backObjV2, rl.White)
				case "backObjGround14":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-48))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-50))
					rl.DrawTextureRec(imgs, backObjGround14IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround14IMG, backObjV2, rl.White)
				case "backObjGround15":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-10))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-12))
					rl.DrawTextureRec(imgs, backObjGround15IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround15IMG, backObjV2, rl.White)
				case "backObjGround16":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-44))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-46))
					rl.DrawTextureRec(imgs, backObjGround16IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround16IMG, backObjV2, rl.White)
				case "backObjGround17":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-40))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-42))
					rl.DrawTextureRec(imgs, backObjGround17IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround17IMG, backObjV2, rl.White)
				case "backObjGround18":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-110))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-112))
					rl.DrawTextureRec(imgs, backObjGround18IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround18IMG, backObjV2, rl.White)
				case "backObjGround19":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-120))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-122))
					rl.DrawTextureRec(imgs, backObjGround19IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround19IMG, backObjV2, rl.White)
				case "backObjGround20":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-134))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-136))
					rl.DrawTextureRec(imgs, backObjGround20IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround20IMG, backObjV2, rl.White)
				case "backObjGround21":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-120))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-122))
					rl.DrawTextureRec(imgs, backObjGround21IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround21IMG, backObjV2, rl.White)
				case "backObjGround22":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-40))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-40))
					rl.DrawTextureRec(imgs, backObjGround22IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround22IMG, backObjV2, rl.White)
				case "backObjGround23":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-42))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-44))
					rl.DrawTextureRec(imgs, backObjGround23IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround23IMG, backObjV2, rl.White)
				case "backObjGround24":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-96))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-98))
					rl.DrawTextureRec(imgs, backObjGround24IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround24IMG, backObjV2, rl.White)
				case "backObjGround25":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-82))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-84))
					rl.DrawTextureRec(imgs, backObjGround25IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround25IMG, backObjV2, rl.White)
				case "backObjGround26":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-70))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-72))
					rl.DrawTextureRec(imgs, backObjGround26IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround26IMG, backObjV2, rl.White)
				case "backObjGround27":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-110))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-112))
					rl.DrawTextureRec(imgs, backObjGround27IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround27IMG, backObjV2, rl.White)
				case "backObjGround28":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-162))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-164))
					rl.DrawTextureRec(imgs, backObjGround28IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround28IMG, backObjV2, rl.White)
				case "backObjGround29":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-38))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-40))
					rl.DrawTextureRec(imgs, backObjGround29IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround29IMG, backObjV2, rl.White)
				case "backObjGround30":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-70))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-72))
					rl.DrawTextureRec(imgs, backObjGround30IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround30IMG, backObjV2, rl.White)
				case "backObjGround31":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-86))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-88))
					rl.DrawTextureRec(imgs, backObjGround31IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround31IMG, backObjV2, rl.White)
				case "backObjGround32":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-120))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-122))
					rl.DrawTextureRec(imgs, backObjGround32IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround32IMG, backObjV2, rl.White)
				case "backObjGround33":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-8))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-10))
					rl.DrawTextureRec(imgs, backObjGround33IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround33IMG, backObjV2, rl.White)
				case "backObjGround34":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-12))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-14))
					rl.DrawTextureRec(imgs, backObjGround34IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround34IMG, backObjV2, rl.White)
				case "backObjGround35":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-12))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-14))
					rl.DrawTextureRec(imgs, backObjGround35IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround35IMG, backObjV2, rl.White)
				case "backObjGround36":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-86))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-88))
					rl.DrawTextureRec(imgs, backObjGround36IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround36IMG, backObjV2, rl.White)
				case "backObjGround37":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-12))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-14))
					rl.DrawTextureRec(imgs, backObjGround37IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround37IMG, backObjV2, rl.White)
				case "backObjGround38":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-14))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-16))
					rl.DrawTextureRec(imgs, backObjGround38IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround38IMG, backObjV2, rl.White)
				case "backObjGround39":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-10))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-12))
					rl.DrawTextureRec(imgs, backObjGround39IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround39IMG, backObjV2, rl.White)
				case "backObjGround40":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-40))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-42))
					rl.DrawTextureRec(imgs, backObjGround40IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround40IMG, backObjV2, rl.White)
				case "backObjGround41":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-42))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-44))
					rl.DrawTextureRec(imgs, backObjGround41IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround41IMG, backObjV2, rl.White)
				case "backObjGround42":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-44))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-46))
					rl.DrawTextureRec(imgs, backObjGround42IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround42IMG, backObjV2, rl.White)
				case "backObjGround43":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-30))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-32))
					rl.DrawTextureRec(imgs, backObjGround43IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround43IMG, backObjV2, rl.White)
				case "backObjGround44":
					backObjV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-60))
					backObjV2shadow := rl.NewVector2(float32(drawScreenX-3), float32(drawScreenY-62))
					rl.DrawTextureRec(imgs, backObjGround44IMG, backObjV2shadow, rl.Fade(rl.Black, 0.6))
					rl.DrawTextureRec(imgs, backObjGround44IMG, backObjV2, rl.White)

				}

				drawScreenCurrentBlock++
				drawScreenLineCount++
				drawScreenX += 8
				if drawScreenLineCount == 86 {
					drawScreenCurrentBlock += 8
					drawScreenLineCount = 0
					drawScreenX = 0
					drawScreenY += 8
				}
			}

			// MARK: draw screen layer 1
			drawScreenCurrentBlock = 192
			drawScreenLineCount := 0
			drawScreenX := int32(0)
			drawScreenY := int32(0)

			for a := 0; a < 4128; a++ {

				checkBulletsBlock := bulletsMAP[drawScreenCurrentBlock]
				checkPlayerBlock := playerMAP[drawScreenCurrentBlock]
				checkEnemyBlock := enemiesMAP[drawScreenCurrentBlock]
				checkPowerUpBlock := powerUpsMAP[drawScreenCurrentBlock]
				checkCoinBlock := objectsMAP[drawScreenCurrentBlock]
				checkPlatformsBlock := platformsMAP[drawScreenCurrentBlock]
				checkWeaponsBlock := weaponsMAP[drawScreenCurrentBlock]
				checkEarthquakesBlock := earthquakesMAP[drawScreenCurrentBlock]
				checkPlatformEffectsBlock := platformsEffectsMAP[drawScreenCurrentBlock]
				checkRainFrogBlock := rainFrogMAP[drawScreenCurrentBlock]

				if pigeonOn {
					checkPigeonBlock = activeSpecialMAP[drawScreenCurrentBlock]

					if checkPigeonBlock != "" && checkPowerUpBlock != "" {
						rl.DrawRectangle(100, 100, 100, 100, rl.Red)
					}

				}

				switch checkRainFrogBlock {

				case "rainFrog":
					rainFrogV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, rainFrogIMG, rainFrogV2, rl.White)
				case "rainFrogR":
					rainFrogV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, rainFrogRIMG, rainFrogV2, rl.White)
				}

				// pick up coin
				if checkPlayerBlock != "" && checkCoinBlock != "" {

					playerCoins++

					if checkCoinBlock == "coinTL" {
						coinCurrentBlock = drawScreenCurrentBlock
						clearCoin()
					} else if checkCoinBlock == "coin" {
						if objectsMAP[drawScreenCurrentBlock-94] == "coinTL" {
							coinCurrentBlock = drawScreenCurrentBlock - 94
							clearCoin()
						} else if objectsMAP[drawScreenCurrentBlock-1] == "coinTL" {
							coinCurrentBlock = drawScreenCurrentBlock - 1
							clearCoin()
						} else if objectsMAP[drawScreenCurrentBlock-95] == "coinTL" {
							coinCurrentBlock = drawScreenCurrentBlock - 95
							clearCoin()
						}
					}

					coinCollectedSound = true

				}

				// pick up weapon
				if checkPlayerBlock != "" && checkWeaponsBlock != "" {

					if checkWeaponsBlock == "weapon1TL" || checkWeaponsBlock == "weapon2TL" || checkWeaponsBlock == "weapon3TL" || checkWeaponsBlock == "weapon4TL" || checkWeaponsBlock == "weapon5TL" || checkWeaponsBlock == "weapon6TL" || checkWeaponsBlock == "weapon7TL" || checkWeaponsBlock == "weapon8TL" || checkWeaponsBlock == "weapon9TL" || checkWeaponsBlock == "weapon10TL" {

						currentPlayerWeapon = checkWeaponsBlock
						weaponsMAP[drawScreenCurrentBlock] = ""
						weaponsMAP[drawScreenCurrentBlock+1] = ""

					}

				}

				// pick up powerup
				if checkPlayerBlock != "" && checkPowerUpBlock != "" {

					if checkPowerUpBlock == "powerUp1TL" || checkPowerUpBlock == "powerUp1" {

						powerUpCollected = true
						powerUpCurrentActive = 1

						if checkPowerUpBlock == "powerUp1TL" {
							currentPowerUpBlock = drawScreenCurrentBlock
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp1" && powerUpsMAP[a-94] == "powerUp1TL" {
							currentPowerUpBlock = a - 94
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp1" && powerUpsMAP[a-95] == "powerUp1TL" {
							currentPowerUpBlock = a - 95
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp1" && powerUpsMAP[a-1] == "powerUp1TL" {
							currentPowerUpBlock = a - 1
							clearPowerUp()
						}
					} else if checkPowerUpBlock == "powerUp2TL" || checkPowerUpBlock == "powerUp2" {

						powerUpCollected = true
						powerUpCurrentActive = 2

						if checkPowerUpBlock == "powerUp2TL" {
							currentPowerUpBlock = drawScreenCurrentBlock
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp2" && powerUpsMAP[a-94] == "powerUp2TL" {
							currentPowerUpBlock = a - 94
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp2" && powerUpsMAP[a-95] == "powerUp2TL" {
							currentPowerUpBlock = a - 95
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp2" && powerUpsMAP[a-1] == "powerUp2TL" {
							currentPowerUpBlock = a - 1
							clearPowerUp()
						}
					} else if checkPowerUpBlock == "powerUp3TL" || checkPowerUpBlock == "powerUp3" {

						powerUpCollected = true
						powerUpCurrentActive = 3

						if checkPowerUpBlock == "powerUp3TL" {
							currentPowerUpBlock = drawScreenCurrentBlock
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp3" && powerUpsMAP[a-94] == "powerUp3TL" {
							currentPowerUpBlock = a - 94
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp3" && powerUpsMAP[a-95] == "powerUp3TL" {
							currentPowerUpBlock = a - 95
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp3" && powerUpsMAP[a-1] == "powerUp3TL" {
							currentPowerUpBlock = a - 1
							clearPowerUp()
						}
					} else if checkPowerUpBlock == "powerUp4TL" || checkPowerUpBlock == "powerUp4" {

						powerUpCollected = true
						powerUpCurrentActive = 4

						if checkPowerUpBlock == "powerUp4TL" {
							currentPowerUpBlock = drawScreenCurrentBlock
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp4" && powerUpsMAP[a-94] == "powerUp4TL" {
							currentPowerUpBlock = a - 94
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp4" && powerUpsMAP[a-95] == "powerUp4TL" {
							currentPowerUpBlock = a - 95
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp4" && powerUpsMAP[a-1] == "powerUp4TL" {
							currentPowerUpBlock = a - 1
							clearPowerUp()
						}
					} else if checkPowerUpBlock == "powerUp5TL" || checkPowerUpBlock == "powerUp5" {

						powerUpCollected = true
						powerUpCurrentActive = 5

						if checkPowerUpBlock == "powerUp5TL" {
							currentPowerUpBlock = drawScreenCurrentBlock
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp5" && powerUpsMAP[a-94] == "powerUp5TL" {
							currentPowerUpBlock = a - 94
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp5" && powerUpsMAP[a-95] == "powerUp5TL" {
							currentPowerUpBlock = a - 95
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp5" && powerUpsMAP[a-1] == "powerUp5TL" {
							currentPowerUpBlock = a - 1
							clearPowerUp()
						}
					} else if checkPowerUpBlock == "powerUp6TL" || checkPowerUpBlock == "powerUp6" {

						powerUpCollected = true
						powerUpCurrentActive = 6

						if checkPowerUpBlock == "powerUp6TL" {
							currentPowerUpBlock = drawScreenCurrentBlock
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp6" && powerUpsMAP[a-94] == "powerUp6TL" {
							currentPowerUpBlock = a - 94
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp6" && powerUpsMAP[a-95] == "powerUp6TL" {
							currentPowerUpBlock = a - 95
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp6" && powerUpsMAP[a-1] == "powerUp6TL" {
							currentPowerUpBlock = a - 1
							clearPowerUp()
						}
					} else if checkPowerUpBlock == "powerUp7TL" || checkPowerUpBlock == "powerUp7" {

						powerUpCollected = true
						powerUpCurrentActive = 7

						if checkPowerUpBlock == "powerUp7TL" {
							currentPowerUpBlock = drawScreenCurrentBlock
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp7" && powerUpsMAP[a-94] == "powerUp7TL" {
							currentPowerUpBlock = a - 94
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp7" && powerUpsMAP[a-95] == "powerUp7TL" {
							currentPowerUpBlock = a - 95
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp7" && powerUpsMAP[a-1] == "powerUp7TL" {
							currentPowerUpBlock = a - 1
							clearPowerUp()
						}
					} else if checkPowerUpBlock == "powerUp8TL" || checkPowerUpBlock == "powerUp8" {

						powerUpCollected = true
						powerUpCurrentActive = 8

						if checkPowerUpBlock == "powerUp8TL" {
							currentPowerUpBlock = drawScreenCurrentBlock
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp8" && powerUpsMAP[a-94] == "powerUp8TL" {
							currentPowerUpBlock = a - 94
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp8" && powerUpsMAP[a-95] == "powerUp8TL" {
							currentPowerUpBlock = a - 95
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp8" && powerUpsMAP[a-1] == "powerUp8TL" {
							currentPowerUpBlock = a - 1
							clearPowerUp()
						}
					} else if checkPowerUpBlock == "powerUp9TL" || checkPowerUpBlock == "powerUp9" {

						powerUpCollected = true
						powerUpCurrentActive = 9

						if checkPowerUpBlock == "powerUp9TL" {
							currentPowerUpBlock = drawScreenCurrentBlock
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp9" && powerUpsMAP[a-94] == "powerUp9TL" {
							currentPowerUpBlock = a - 94
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp9" && powerUpsMAP[a-95] == "powerUp9TL" {
							currentPowerUpBlock = a - 95
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp9" && powerUpsMAP[a-1] == "powerUp9TL" {
							currentPowerUpBlock = a - 1
							clearPowerUp()
						}
					} else if checkPowerUpBlock == "powerUp10TL" || checkPowerUpBlock == "powerUp10" {

						powerUpCollected = true
						powerUpCurrentActive = 10

						if checkPowerUpBlock == "powerUp10TL" {
							currentPowerUpBlock = drawScreenCurrentBlock
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp10" && powerUpsMAP[a-94] == "powerUp10TL" {
							currentPowerUpBlock = a - 94
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp10" && powerUpsMAP[a-95] == "powerUp10TL" {
							currentPowerUpBlock = a - 95
							clearPowerUp()
						} else if checkPowerUpBlock == "powerUp10" && powerUpsMAP[a-1] == "powerUp10TL" {
							currentPowerUpBlock = a - 1
							clearPowerUp()
						}
					}

				}

				if hpLossPause == false {
					if checkPlayerBlock == "playerTL" || checkPlayerBlock == "player" {
						if checkEnemyBlock == "enemyTL" || checkEnemyBlock == "enemy" {
							playerHP--
							hpLossPause = true
							hpLossSound = true
						}
					}
				}

				switch checkPlatformsBlock {

				case "movingplatH1":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform16TileIMG, platformTileV2, rl.White)
				case "movingplatH2":
				//	rl.DrawRectangle(drawScreenX, drawScreenY, 8, 8, rl.Green)
				case "movingplatH3":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform16TileIMG, platformTileV2, rl.White)
				case "movingplatH4":
					// rl.DrawRectangle(drawScreenX, drawScreenY, 8, 8, rl.Red)
				case "movingplatH5":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform16TileIMG, platformTileV2, rl.White)
				case "movingplatH6":
				//	rl.DrawRectangle(drawScreenX, drawScreenY, 8, 8, rl.Green)
				case "movingplatH7":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform16TileIMG, platformTileV2, rl.White)
				case "movingplatH8":
				//	rl.DrawRectangle(drawScreenX, drawScreenY, 8, 8, rl.Red)
				case "movingPlatformUpTL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform16TileIMG, platformTileV2, rl.White)
				case "movingPlatformDownTL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform16TileIMG, platformTileV2, rl.White)
				case "ground1TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground1IMG, platformTileV2, rl.White)
				case "ground2TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground2IMG, platformTileV2, rl.White)
				case "ground3TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground3IMG, platformTileV2, rl.White)
				case "ground4TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground4IMG, platformTileV2, rl.White)
				case "ground5TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground5IMG, platformTileV2, rl.White)
				case "ground6TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground6IMG, platformTileV2, rl.White)
				case "ground7TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground7IMG, platformTileV2, rl.White)
				case "ground8TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground8IMG, platformTileV2, rl.White)
				case "ground9TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground9IMG, platformTileV2, rl.White)
				case "ground10TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground10IMG, platformTileV2, rl.White)
				case "ground11TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground11IMG, platformTileV2, rl.White)
				case "ground12TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground12IMG, platformTileV2, rl.White)
				case "ground13TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground13IMG, platformTileV2, rl.White)
				case "ground14TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground14IMG, platformTileV2, rl.White)
				case "ground15TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground15IMG, platformTileV2, rl.White)
				case "ground16TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground16IMG, platformTileV2, rl.White)
				case "ground17TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground17IMG, platformTileV2, rl.White)
				case "ground18TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground18IMG, platformTileV2, rl.White)
				case "ground19TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground19IMG, platformTileV2, rl.White)
				case "ground20TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground20IMG, platformTileV2, rl.White)
				case "ground21TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground21IMG, platformTileV2, rl.White)
				case "ground22TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground22IMG, platformTileV2, rl.White)
				case "ground23TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground23IMG, platformTileV2, rl.White)
				case "ground24TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, ground24IMG, platformTileV2, rl.White)

				case "platform1TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform1TileIMG, platformTileV2, rl.White)
				case "platform2TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform2TileIMG, platformTileV2, rl.White)
				case "platform3TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform3TileIMG, platformTileV2, rl.White)
				case "platform4TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform4TileIMG, platformTileV2, rl.White)
				case "platform5TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform5TileIMG, platformTileV2, rl.White)
				case "platform6TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform6TileIMG, platformTileV2, rl.White)
				case "platform7TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform7TileIMG, platformTileV2, rl.White)
				case "platform8TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform8TileIMG, platformTileV2, rl.White)
				case "platform9TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform9TileIMG, platformTileV2, rl.White)
				case "platform10TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform10TileIMG, platformTileV2, rl.White)
				case "platform11TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform11TileIMG, platformTileV2, rl.White)
				case "platform12TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform12TileIMG, platformTileV2, rl.White)
				case "platform13TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform13TileIMG, platformTileV2, rl.White)
				case "platform14TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform14TileIMG, platformTileV2, rl.White)
				case "platform15TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform15TileIMG, platformTileV2, rl.White)
				case "platform16TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform17TileIMG, platformTileV2, rl.White)
				case "platform17TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform17TileIMG, platformTileV2, rl.White)
				case "platform18TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform18TileIMG, platformTileV2, rl.White)
				case "platform19TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform19TileIMG, platformTileV2, rl.White)
				case "platform20TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform20TileIMG, platformTileV2, rl.White)
				case "platform21TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform21TileIMG, platformTileV2, rl.White)
				case "platform22TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform22TileIMG, platformTileV2, rl.White)
				case "platform23TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform23TileIMG, platformTileV2, rl.White)
				case "platform24TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform24TileIMG, platformTileV2, rl.White)
				case "platform25TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform25TileIMG, platformTileV2, rl.White)
				case "platform26TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform26TileIMG, platformTileV2, rl.White)
				case "platform27TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform27TileIMG, platformTileV2, rl.White)
				case "platform28TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform28TileIMG, platformTileV2, rl.White)
				case "platform29TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform29TileIMG, platformTileV2, rl.White)
				case "platform30TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform30TileIMG, platformTileV2, rl.White)
				case "platform31TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform31TileIMG, platformTileV2, rl.White)
				case "platform32TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform32TileIMG, platformTileV2, rl.White)
				case "platform33TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform33TileIMG, platformTileV2, rl.White)
				case "platform34TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform34TileIMG, platformTileV2, rl.White)
				case "platform35TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform35TileIMG, platformTileV2, rl.White)
				case "platform36TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform36TileIMG, platformTileV2, rl.White)
				case "platform37TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform37TileIMG, platformTileV2, rl.White)
				case "platform38TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform38TileIMG, platformTileV2, rl.White)
				case "platform39TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform39TileIMG, platformTileV2, rl.White)
				case "platform40TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform40TileIMG, platformTileV2, rl.White)
				case "platform41TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform41TileIMG, platformTileV2, rl.White)
				case "platform42TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform42TileIMG, platformTileV2, rl.White)
				case "platform43TL":
					platformTileV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, platform43TileIMG, platformTileV2, rl.White)

				}

				switch checkPlatformEffectsBlock {

				case "fire", "fire2", "fire3", "fire4", "fire5", "fire6":
					//	rl.DrawRectangle(drawScreenX, drawScreenY, 8, 8, rl.Fade(rl.Orange, 0.8))
					platformTileV2 := rl.NewVector2(float32(drawScreenX-8), float32(drawScreenY-8))
					rl.DrawTextureRec(imgs, flameIMG, platformTileV2, rl.White)

				case "poison", "poison2", "poison3", "poison4", "poison5":
					platformTileV2 := rl.NewVector2(float32(drawScreenX-8), float32(drawScreenY-8))
					rl.DrawTextureRec(imgs, poisonFlameIMG, platformTileV2, rl.White)

				}

				// MARK: draw earthquakes

				switch checkEarthquakesBlock {

				case "earthquake":
					rl.DrawRectangle(drawScreenX, drawScreenY, 8, 8, rl.Fade(rl.Black, 0.9))
					rl.DrawRectangle(drawScreenX, drawScreenY, 8, 8, rl.Fade(rl.DarkBrown, 0.2))

				}

				// MARK: draw dropped weapons

				switch checkWeaponsBlock {

				case "weapon1TL":
					dropWeaponV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					popUpV2 := rl.NewVector2(float32(drawScreenX-2), float32(drawScreenY-36))
					rl.DrawTextureRec(imgs, getArrowIMG, popUpV2, rl.White)
					rl.DrawTextureRec(imgs, weapon1IMG, dropWeaponV2, rl.White)
					rl.DrawText("glockenspiel", drawScreenX-21, drawScreenY-11, 10, rl.Black)
					rl.DrawText("glockenspiel", drawScreenX-20, drawScreenY-12, 10, weaponTextColor)
				case "weapon2TL":
					dropWeaponV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					popUpV2 := rl.NewVector2(float32(drawScreenX-2), float32(drawScreenY-36))
					rl.DrawTextureRec(imgs, getArrowIMG, popUpV2, rl.White)
					rl.DrawTextureRec(imgs, weapon2IMG, dropWeaponV2, rl.White)
					rl.DrawText("cervelat", drawScreenX-13, drawScreenY-11, 10, rl.Black)
					rl.DrawText("cervelat", drawScreenX-12, drawScreenY-12, 10, weaponTextColor)
				case "weapon3TL":
					dropWeaponV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					popUpV2 := rl.NewVector2(float32(drawScreenX-2), float32(drawScreenY-36))
					rl.DrawTextureRec(imgs, getArrowIMG, popUpV2, rl.White)
					rl.DrawTextureRec(imgs, weapon3IMG, dropWeaponV2, rl.White)
					rl.DrawText("manzello", drawScreenX-12, drawScreenY-11, 10, rl.Black)
					rl.DrawText("manzello", drawScreenX-13, drawScreenY-12, 10, weaponTextColor)
				case "weapon4TL":
					dropWeaponV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					popUpV2 := rl.NewVector2(float32(drawScreenX-2), float32(drawScreenY-36))
					rl.DrawTextureRec(imgs, getArrowIMG, popUpV2, rl.White)
					rl.DrawTextureRec(imgs, weapon4IMG, dropWeaponV2, rl.White)
					rl.DrawText("pibcorn", drawScreenX-11, drawScreenY-11, 10, rl.Black)
					rl.DrawText("pibcorn", drawScreenX-10, drawScreenY-12, 10, weaponTextColor)
				case "weapon5TL":
					dropWeaponV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					popUpV2 := rl.NewVector2(float32(drawScreenX-2), float32(drawScreenY-36))
					rl.DrawTextureRec(imgs, getArrowIMG, popUpV2, rl.White)
					rl.DrawTextureRec(imgs, weapon5IMG, dropWeaponV2, rl.White)
					rl.DrawText("vandola", drawScreenX-15, drawScreenY-11, 10, rl.Black)
					rl.DrawText("vandola", drawScreenX-14, drawScreenY-12, 10, weaponTextColor)
				case "weapon6TL":
					dropWeaponV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					popUpV2 := rl.NewVector2(float32(drawScreenX-2), float32(drawScreenY-36))
					rl.DrawTextureRec(imgs, getArrowIMG, popUpV2, rl.White)
					rl.DrawTextureRec(imgs, weapon6IMG, dropWeaponV2, rl.White)
					rl.DrawText("zambomba", drawScreenX-15, drawScreenY-11, 10, rl.Black)
					rl.DrawText("zambomba", drawScreenX-14, drawScreenY-12, 10, weaponTextColor)
				case "weapon7TL":
					dropWeaponV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					popUpV2 := rl.NewVector2(float32(drawScreenX-2), float32(drawScreenY-36))
					rl.DrawTextureRec(imgs, getArrowIMG, popUpV2, rl.White)
					rl.DrawTextureRec(imgs, weapon7IMG, dropWeaponV2, rl.White)
					rl.DrawText("marimbaphone", drawScreenX-23, drawScreenY-11, 10, rl.Black)
					rl.DrawText("marimbaphone", drawScreenX-22, drawScreenY-12, 10, weaponTextColor)

				}

				switch checkPowerUpBlock {
				case "powerUp1TL":
					//	rl.DrawRectangle(drawScreenX, drawScreenY, 16, 16, rl.Fade(rl.Orange, 0.7))
					popUpV2 := rl.NewVector2(float32(drawScreenX+6), float32(drawScreenY-38))
					powerUpV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-14))
					rl.DrawTextureRec(imgs, randomPowerUpIMG, powerUpV2, rl.White)
					rl.DrawTextureRec(imgs, getArrowIMG, popUpV2, rl.White)
				case "powerUp2TL":
					//	rl.DrawRectangle(drawScreenX, drawScreenY, 16, 16, rl.Fade(rl.Orange, 0.7))
					popUpV2 := rl.NewVector2(float32(drawScreenX+6), float32(drawScreenY-38))
					powerUpV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-14))
					rl.DrawTextureRec(imgs, hpPowerUpIMG, powerUpV2, rl.White)
					rl.DrawTextureRec(imgs, getArrowIMG, popUpV2, rl.White)
				case "powerUp3TL":
					//	rl.DrawRectangle(drawScreenX, drawScreenY, 16, 16, rl.Fade(rl.Orange, 0.7))
					popUpV2 := rl.NewVector2(float32(drawScreenX+6), float32(drawScreenY-38))
					powerUpV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-14))
					rl.DrawTextureRec(imgs, specialItemIMG, powerUpV2, rl.White)
					rl.DrawTextureRec(imgs, getArrowIMG, popUpV2, rl.White)
				case "powerUp4TL":
					//	rl.DrawRectangle(drawScreenX, drawScreenY, 16, 16, rl.Fade(rl.Orange, 0.7))
					popUpV2 := rl.NewVector2(float32(drawScreenX+6), float32(drawScreenY-38))
					powerUpV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-14))
					rl.DrawTextureRec(imgs, bombIMG, powerUpV2, rl.White)
					rl.DrawTextureRec(imgs, getArrowIMG, popUpV2, rl.White)
				case "powerUp5TL":
					rl.DrawRectangle(drawScreenX, drawScreenY, 16, 16, rl.Fade(rl.Orange, 0.7))
				case "powerUp6TL":
					rl.DrawRectangle(drawScreenX, drawScreenY, 16, 16, rl.Fade(rl.Orange, 0.7))
				case "powerUp7TL":
					rl.DrawRectangle(drawScreenX, drawScreenY, 16, 16, rl.Fade(rl.Orange, 0.7))
				case "powerUp8TL":
					rl.DrawRectangle(drawScreenX, drawScreenY, 16, 16, rl.Fade(rl.Orange, 0.7))
				case "powerUp9TL":
					rl.DrawRectangle(drawScreenX, drawScreenY, 16, 16, rl.Fade(rl.Orange, 0.7))
				case "powerUp10TL":
					rl.DrawRectangle(drawScreenX, drawScreenY, 16, 16, rl.Fade(rl.Orange, 0.7))

				}

				switch checkCoinBlock {

				case "coinTL":
					coinV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))

					rl.DrawTextureRec(imgs, coinIMG, coinV2, rl.White)
				}

				switch checkBulletsBlock {
				case "bulletEnemyD":
					rl.DrawCircle(drawScreenX+6, drawScreenY, 3, rl.Purple)
					rl.DrawCircle(drawScreenX+6, drawScreenY+4, 5, rl.Purple)
				case "bulletEnemyR":
					rl.DrawCircle(drawScreenX, drawScreenY+4, 3, rl.Green)
					rl.DrawCircle(drawScreenX+6, drawScreenY+4, 5, rl.Green)
				case "bulletEnemyL":
					rl.DrawCircle(drawScreenX+12, drawScreenY+4, 3, rl.Green)
					rl.DrawCircle(drawScreenX+6, drawScreenY+4, 5, rl.Green)
				case "bulletL":
					rl.DrawCircle(drawScreenX+6, drawScreenY+4, 4, rl.Violet)
					rl.DrawCircle(drawScreenX+6, drawScreenY+4, 3, rl.Red)
					rl.DrawCircle(drawScreenX+6, drawScreenY+4, 2, rl.Orange)
				case "bulletR":
					rl.DrawCircle(drawScreenX+2, drawScreenY+4, 4, rl.Violet)
					rl.DrawCircle(drawScreenX+2, drawScreenY+4, 3, rl.Red)
					rl.DrawCircle(drawScreenX+2, drawScreenY+4, 2, rl.Orange)
				case "bulletU":
					rl.DrawCircle(drawScreenX+2, drawScreenY+4, 4, rl.Violet)
					rl.DrawCircle(drawScreenX+2, drawScreenY+4, 3, rl.Red)
					rl.DrawCircle(drawScreenX+2, drawScreenY+4, 2, rl.Orange)
				}

				drawScreenCurrentBlock++
				drawScreenLineCount++
				drawScreenX += 8
				if drawScreenLineCount == 86 {
					drawScreenCurrentBlock += 8
					drawScreenLineCount = 0
					drawScreenX = 0
					drawScreenY += 8
				}
			}

			// MARK: draw screen layer 2

			drawScreenCurrentBlock = 192
			drawScreenCurrentBlockWeather = drawScreenCurrentBlockWeatherNEXT
			drawScreenLineCount = 0
			drawScreenX = int32(0)
			drawScreenY = int32(0)

			for a := 0; a < 4128; a++ {

				checkEffectsBlock := effectsMAP[drawScreenCurrentBlock]
				checkPlayerBlock := playerMAP[drawScreenCurrentBlock]
				checkCloudsBlock := cloudsMAP[drawScreenCurrentBlockWeather]
				checkPowerUpVehicleBlock := powerUpVehicleMAP[drawScreenCurrentBlock]
				checkEnemyBlock := enemiesMAP[drawScreenCurrentBlock]
				checkEnemyMovementBlock := enemiesMovementMAP[drawScreenCurrentBlock]
				checkMiniBossEffectsBlock := miniBossEffectsMAP[drawScreenCurrentBlock]
				checkMiniBossBlock := miniBossMAP[drawScreenCurrentBlock]
				checkTeleportBlock := teleportsMAP[drawScreenCurrentBlock]
				checkActiveSpecialBlock := activeSpecialMAP[drawScreenCurrentBlock]
				checkTornadoBlock := tornadoMAP[drawScreenCurrentBlock]
				checkMeteorBlock := meteorMAP[drawScreenCurrentBlock]
				checkRainOfFireBlock := rainoffireMAP[drawScreenCurrentBlock]
				checkSnowBlock := snowMAP[drawScreenCurrentBlock]
				checkShockBlock := shockBlockMAP[drawScreenCurrentBlock]

				// player collision shock blocks
				if shockBlocksOn {
					if a > 564 {
						if checkShockBlock == "shockBlock" && playerMAP[a-94] == "player" {
							playerHP--
							hpLossPause = true
							hpLossSound = true
						} else if checkShockBlock == "shockBlock" && playerMAP[a-94*2] == "player" {
							playerHP--
							hpLossPause = true
							hpLossSound = true
						} else if checkShockBlock == "shockBlock" && playerMAP[a-94*3] == "player" {
							playerHP--
							hpLossPause = true
							hpLossSound = true
						} else if checkShockBlock == "shockBlock" && playerMAP[a-94*4] == "player" {
							playerHP--
							hpLossPause = true
							hpLossSound = true
						} else if checkShockBlock == "shockBlock" && playerMAP[a-94*5] == "player" {
							playerHP--
							hpLossPause = true
							hpLossSound = true
						} else if checkShockBlock == "shockBlock" && playerMAP[a] == "player" {
							playerHP--
							hpLossPause = true
							hpLossSound = true
						}
					}
				}

				// player collision rain of fire
				if rainoffireOn {
					if a > 564 {
						if checkRainOfFireBlock == "rainoffire" && playerMAP[a-94] == "player" {
							playerHP--
							hpLossPause = true
							hpLossSound = true
						} else if checkRainOfFireBlock == "rainoffire" && playerMAP[a-94*2] == "player" {
							playerHP--
							hpLossPause = true
							hpLossSound = true
						} else if checkRainOfFireBlock == "rainoffire" && playerMAP[a] == "player" {
							playerHP--
							hpLossPause = true
							hpLossSound = true
						}
					}
				}

				// player miniboss effects collisions
				if hpLossPause == false {
					if checkPlayerBlock == "playerTL" && checkMiniBossEffectsBlock != "" {
						playerHP--
						hpLossPause = true
						hpLossSound = true
					} else if checkPlayerBlock == "player" && checkMiniBossEffectsBlock != "" {
						playerHP--
						hpLossPause = true
						hpLossSound = true
					} else if checkPlayerBlock == "player" && checkMiniBossBlock != "" {
						playerHP--
						hpLossPause = true
						hpLossSound = true
					} else if checkPlayerBlock == "playerTL" && checkMiniBossBlock != "" {
						playerHP--
						hpLossPause = true
						hpLossSound = true
					}
				}

				// MARK: trampoline jump
				if trampolineOn {
					if jumpActive == false && fallActive == false {
						if checkPlayerBlock != "" && checkActiveSpecialBlock != "" {
							jumpHeight = 24
							jumpActive = true
						}
					}
				}
				// MARK: teleport move player
				if checkTeleportBlock != "" && checkPlayerBlock != "" {

					if teleportPostion1-playerCurrentBlock < 564 {
						clearPlayer()
						playerCurrentBlock = teleportPostion2 - 10
					} else {
						clearPlayer()
						playerCurrentBlock = teleportPostion1 - 10
					}
				}

				switch checkShockBlock {

				case "shockBlock":
					//	rl.DrawRectangle(drawScreenX, drawScreenY, 8, 8, rl.Fade(rl.Blue, 0.8))
					shockblockV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, shockBlockIMG, shockblockV2, rl.White)
				}

				switch checkMiniBossEffectsBlock {
				case "zap":
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 5, rl.Violet)
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 4, rl.Magenta)
				case "zap2":
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 4, rl.Violet)
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 3, rl.Magenta)
				case "zap3":
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 5, rl.Violet)
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 2, rl.Magenta)
				case "zap4":
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 2, rl.Violet)
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 1, rl.Magenta)
				case "zap5":
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 5, rl.DarkBlue)
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 4, rl.Blue)
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 3, rl.SkyBlue)
				case "zap6":
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 5, rl.DarkBlue)
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 4, rl.Blue)
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 3, rl.SkyBlue)
				case "zap7":
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 5, rl.DarkBrown)
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 4, rl.Orange)
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 3, rl.Gold)
				case "zap8":
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 5, rl.DarkBrown)
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 4, rl.Orange)
					rl.DrawCircle(drawScreenX+4, drawScreenY+4, 3, rl.Gold)
				}

				switch checkEnemyBlock {
				case "enemyTL":

					switch enemyTypeGenerate {
					case 1:
						enemyV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-8))
						rl.DrawTextureRec(imgs, enemy1IMG, enemyV2, rl.White)
					case 2:
						enemyV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-16))
						rl.DrawTextureRec(imgs, enemy2IMG, enemyV2, rl.White)
					case 3:
						enemyV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-24))
						rl.DrawTextureRec(imgs, enemy3IMG, enemyV2, rl.White)
					case 4:
						enemyV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-16))
						if checkEnemyMovementBlock != "enemy4Special" {
							rl.DrawTextureRec(imgs, enemy4IMG, enemyV2, rl.White)
						} else if checkEnemyMovementBlock == "enemy4Special" {
							rl.DrawTextureRec(imgs, enemy4IMG, enemyV2, rl.Black)
						} else if checkEnemyMovementBlock == "enemy4Special2" {
							rl.DrawTextureRec(imgs, enemy4IMG, enemyV2, rl.Green)
						}
					case 5:
						enemyV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-14))
						rl.DrawTextureRec(imgs, enemy5IMG, enemyV2, rl.White)
						if enemy5Shield1 {
							rl.DrawCircleLines(int32(enemyV2.X+22), int32(enemyV2.Y+18), 24, rl.Purple)
							rl.DrawCircleLines(int32(enemyV2.X+22), int32(enemyV2.Y+18), 26, rl.Blue)
							rl.DrawCircleLines(int32(enemyV2.X+22), int32(enemyV2.Y+18), 28, rl.White)

							rl.DrawCircleLines(int32(enemyV2.X), int32(enemyV2.Y), 4, rl.White)
							rl.DrawCircleLines(int32(enemyV2.X+40), int32(enemyV2.Y-4), 2, rl.Violet)
							rl.DrawCircleLines(int32(enemyV2.X+60), int32(enemyV2.Y+16), 4, rl.White)
							rl.DrawCircleLines(int32(enemyV2.X+70), int32(enemyV2.Y+20), 3, rl.Green)
							rl.DrawCircleLines(int32(enemyV2.X-20), int32(enemyV2.Y-20), 5, rl.Yellow)
						} else if enemy5Shield2 {
							rl.DrawCircleLines(int32(enemyV2.X+22), int32(enemyV2.Y+18), 24, rl.White)
							rl.DrawCircleLines(int32(enemyV2.X+22), int32(enemyV2.Y+18), 26, rl.Purple)
							rl.DrawCircleLines(int32(enemyV2.X+22), int32(enemyV2.Y+18), 28, rl.Blue)

							rl.DrawCircleLines(int32(enemyV2.X-20), int32(enemyV2.Y-10), 4, rl.Blue)
							rl.DrawCircleLines(int32(enemyV2.X+50), int32(enemyV2.Y+10), 6, rl.Orange)
							rl.DrawCircleLines(int32(enemyV2.X+10), int32(enemyV2.Y-22), 2, rl.Pink)
							rl.DrawCircleLines(int32(enemyV2.X-30), int32(enemyV2.Y-32), 3, rl.White)
							rl.DrawCircleLines(int32(enemyV2.X+60), int32(enemyV2.Y+4), 4, rl.Red)
						} else if enemy5Shield3 {
							rl.DrawCircleLines(int32(enemyV2.X+22), int32(enemyV2.Y+18), 24, rl.Blue)
							rl.DrawCircleLines(int32(enemyV2.X+22), int32(enemyV2.Y+18), 26, rl.Purple)
							rl.DrawCircleLines(int32(enemyV2.X+22), int32(enemyV2.Y+18), 28, rl.White)

							rl.DrawCircleLines(int32(enemyV2.X), int32(enemyV2.Y-30), 4, rl.Green)
							rl.DrawCircleLines(int32(enemyV2.X+10), int32(enemyV2.Y+24), 6, rl.Yellow)
							rl.DrawCircleLines(int32(enemyV2.X+4), int32(enemyV2.Y-12), 2, rl.Violet)
							rl.DrawCircleLines(int32(enemyV2.X-44), int32(enemyV2.Y+32), 3, rl.Blue)
							rl.DrawCircleLines(int32(enemyV2.X+2), int32(enemyV2.Y+31), 4, rl.White)

						}
					case 6:
						enemyV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-8))
						rl.DrawTextureRec(imgs, enemy6IMG, enemyV2, rl.White)
					case 7:

						if checkEnemyMovementBlock == "jump5" {
							enemyV2 := rl.NewVector2(float32(drawScreenX-2), float32(drawScreenY-2))
							rl.DrawTextureRec(imgs, enemy7ULIMG, enemyV2, rl.White)
						} else if checkEnemyMovementBlock == "jump6" {
							enemyV2 := rl.NewVector2(float32(drawScreenX-1), float32(drawScreenY-8))
							rl.DrawTextureRec(imgs, enemy7URIMG, enemyV2, rl.White)
						} else if checkEnemyMovementBlock == "roofR" || checkEnemyMovementBlock == "roofL" {
							enemyV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-2))
							rl.DrawTextureRec(imgs, enemy7UPIMG, enemyV2, rl.White)
						} else {
							enemyV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-8))
							rl.DrawTextureRec(imgs, enemy7IMG, enemyV2, rl.White)
						}

					case 8:
						enemyV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-20))
						rl.DrawTextureRec(imgs, enemy8IMG, enemyV2, rl.White)
					case 9:
						enemyV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-8))
						rl.DrawTextureRec(imgs, enemy9IMG, enemyV2, rl.White)

					}

				}

				// MARK: draw active special
				switch checkActiveSpecialBlock {

				case "zap12":
					rl.DrawCircle(drawScreenX, drawScreenY, 5, rl.DarkGreen)
					rl.DrawCircle(drawScreenX, drawScreenY, 4, rl.Green)
					rl.DrawCircle(drawScreenX, drawScreenY, 3, rl.Lime)
				case "zap11":
					rl.DrawCircle(drawScreenX, drawScreenY, 3, rl.Fade(rl.Purple, knightBackFade))
				case "zap10":
					rl.DrawCircle(drawScreenX, drawScreenY-2, 2, rl.White)
					rl.DrawCircle(drawScreenX, drawScreenY, 3, rl.White)
				case "zap9":
					rl.DrawCircle(drawScreenX, drawScreenY-2, 2, rl.Orange)
					rl.DrawCircle(drawScreenX, drawScreenY, 3, rl.Orange)
				case "appleTL":
					activeSpecialV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY+10))
					rl.DrawTextureRec(imgs, appleIMG, activeSpecialV2, rl.White)
				case "petBatTL":
					activeSpecialV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawCircle(drawScreenX+8, drawScreenY+6, 10, rl.Fade(rl.White, 0.4))
					rl.DrawCircleLines(drawScreenX+8, drawScreenY+6, 10, rl.White)
					rl.DrawTextureRec(imgs, petBatIMG, activeSpecialV2, rl.White)
				case "petRedSlimeTL":
					activeSpecialV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY+10))
					exclamationV2 := rl.NewVector2(float32(drawScreenX-4), float32(drawScreenY-10))
					rl.DrawTextureRec(imgs, exclamationIMG, exclamationV2, rl.White)
					rl.DrawTextureRec(imgs, petRedSlimeIMG, activeSpecialV2, rl.White)
				case "petKnightTL":
					activeSpecialV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY+2))
					rl.DrawCircle(drawScreenX+8, drawScreenY+10, 12, rl.Fade(rl.Purple, knightBackFade))
					if petKnightLeftRight {
						rl.DrawTextureRec(imgs, petKnightLIMG, activeSpecialV2, rl.White)
					} else {
						rl.DrawTextureRec(imgs, petKnightIMG, activeSpecialV2, rl.White)
					}
				case "petSlimeTL":
					activeSpecialV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY+4))
					rl.DrawCircle(drawScreenX+8, drawScreenY+10, 12, rl.Fade(rl.Lime, 0.4))
					rl.DrawCircleLines(drawScreenX+8, drawScreenY+10, 12, rl.Lime)
					rl.DrawTextureRec(imgs, petSlimeIMG, activeSpecialV2, rl.White)
				case "petSkeletonTL":
					activeSpecialV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					if skeletonShield1 {
						rl.DrawCircle(drawScreenX-4, drawScreenY-4, 2, rl.Fade(rl.Orange, 0.4))
						rl.DrawCircleLines(drawScreenX-4, drawScreenY-4, 2, rl.Orange)

						rl.DrawCircle(drawScreenX+12, drawScreenY-4, 2, rl.Fade(rl.Orange, 0.4))
						rl.DrawCircleLines(drawScreenX+12, drawScreenY-4, 2, rl.Orange)

						rl.DrawCircle(drawScreenX-8, drawScreenY-6, 3, rl.Fade(rl.Orange, 0.4))
						rl.DrawCircleLines(drawScreenX-8, drawScreenY-6, 3, rl.Orange)
					} else if skeletonShield2 {
						rl.DrawCircle(drawScreenX+6, drawScreenY-2, 2, rl.Fade(rl.Orange, 0.4))
						rl.DrawCircleLines(drawScreenX+6, drawScreenY-2, 2, rl.Orange)

						rl.DrawCircle(drawScreenX, drawScreenY-8, 2, rl.Fade(rl.Orange, 0.4))
						rl.DrawCircleLines(drawScreenX, drawScreenY-8, 2, rl.Orange)

						rl.DrawCircle(drawScreenX+24, drawScreenY+16, 3, rl.Fade(rl.Orange, 0.4))
						rl.DrawCircleLines(drawScreenX+24, drawScreenY+16, 3, rl.Orange)
					} else if skeletonShield3 {
						rl.DrawCircle(drawScreenX+24, drawScreenY+12, 2, rl.Fade(rl.Orange, 0.4))
						rl.DrawCircleLines(drawScreenX+24, drawScreenY+12, 2, rl.Orange)

						rl.DrawCircle(drawScreenX-6, drawScreenY-2, 2, rl.Fade(rl.Orange, 0.4))
						rl.DrawCircleLines(drawScreenX-6, drawScreenY-2, 2, rl.Orange)

						rl.DrawCircle(drawScreenX-10, drawScreenY, 3, rl.Fade(rl.Orange, 0.4))
						rl.DrawCircleLines(drawScreenX-10, drawScreenY, 3, rl.Orange)
					} else if skeletonShield4 {
						rl.DrawCircle(drawScreenX-8, drawScreenY-8, 3, rl.Fade(rl.Orange, 0.4))
						rl.DrawCircleLines(drawScreenX-8, drawScreenY-8, 3, rl.Orange)

						rl.DrawCircle(drawScreenX+12, drawScreenY-4, 2, rl.Fade(rl.Orange, 0.4))
						rl.DrawCircleLines(drawScreenX+12, drawScreenY-4, 2, rl.Orange)

						rl.DrawCircle(drawScreenX-8, drawScreenY-6, 3, rl.Fade(rl.Orange, 0.4))
						rl.DrawCircleLines(drawScreenX-8, drawScreenY-6, 3, rl.Orange)
					}
					if petSkeletonLeftRight {
						rl.DrawTextureRec(imgs, petSkeletonLIMG, activeSpecialV2, rl.White)
					} else {
						rl.DrawTextureRec(imgs, petSkeletonIMG, activeSpecialV2, rl.White)
					}
				case "watermelonTL":
					activeSpecialV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, watermelonIMG, activeSpecialV2, rl.White)
				case "cherriesTL":
					activeSpecialV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY-2))
					rl.DrawTextureRec(imgs, cherriesIMG, activeSpecialV2, rl.White)
				case "petGreenPigTL":
					activeSpecialV2 := rl.NewVector2(float32(drawScreenX-4), float32(drawScreenY-4))
					if pigShieldOn {
						rl.DrawCircle(drawScreenX+8, drawScreenY+6, 18, rl.Fade(rl.DarkGreen, 0.4))
						rl.DrawCircleLines(drawScreenX+8, drawScreenY+6, 18, rl.DarkGreen)
					}
					rl.DrawTextureRec(imgs, petGreenPigIMG, activeSpecialV2, rl.White)
				case "petMushroomTL":
					activeSpecialV2 := rl.NewVector2(float32(drawScreenX-6), float32(drawScreenY-8))
					if mushroomShieldOn {
						rl.DrawCircle(drawScreenX+8, drawScreenY+6, 18, rl.Fade(rl.Pink, 0.4))
						rl.DrawCircleLines(drawScreenX+8, drawScreenY+6, 18, rl.Pink)

						rl.DrawCircle(drawScreenX+32, drawScreenY-4, 2, rl.Fade(rl.Pink, 0.4))
						rl.DrawCircleLines(drawScreenX+40, drawScreenY-12, 2, rl.Pink)

						rl.DrawCircle(drawScreenX-12, drawScreenY-4, 3, rl.Fade(rl.Pink, 0.4))
						rl.DrawCircleLines(drawScreenX-12, drawScreenY-4, 3, rl.Pink)

						rl.DrawCircle(drawScreenX-16, drawScreenY+4, 3, rl.Fade(rl.Pink, 0.4))
						rl.DrawCircleLines(drawScreenX-16, drawScreenY+4, 3, rl.Pink)

						rl.DrawCircle(drawScreenX+20, drawScreenY-14, 4, rl.Fade(rl.Pink, 0.4))
						rl.DrawCircleLines(drawScreenX+20, drawScreenY-14, 4, rl.Pink)
					}
					rl.DrawTextureRec(imgs, petMushroomIMG, activeSpecialV2, rl.White)
				case "trampolineTL":
					activeSpecialV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, trampolineIMG, activeSpecialV2, rl.White)
				case "cannonball":
					rl.DrawCircle(drawScreenX, drawScreenY+8, 10, rl.Black)
				case "cannonTL":
					activeSpecialV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, cannonIMG, activeSpecialV2, rl.White)
				case "iceballTL":
					activeSpecialV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, iceballIMG, activeSpecialV2, rl.White)
				case "poisonballTL":
					activeSpecialV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, poisonballIMG, activeSpecialV2, rl.White)
				case "sawbladeTL":
					activeSpecialV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, sawbladeIMG, activeSpecialV2, rl.White)
				case "fireballTL":
					activeSpecialV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, fireballLIMG, activeSpecialV2, rl.White)
				case "pigeonTL":
					activeSpecialV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawCircle(drawScreenX+19, drawScreenY+2, 30, rl.Fade(rl.White, 0.4))
					rl.DrawCircleLines(drawScreenX+19, drawScreenY+2, 30, rl.White)
					if pigeonShieldFlash1 {
						rl.DrawCircleLines(drawScreenX+19, drawScreenY+2, 32, rl.White)
					} else if pigeonShieldFlash2 {
						rl.DrawCircleLines(drawScreenX+19, drawScreenY+2, 34, rl.White)
					} else if pigeonShieldFlash3 {
						rl.DrawCircleLines(drawScreenX+19, drawScreenY+2, 36, rl.White)
					}

					rl.DrawTextureRec(imgs, petPigeonUpIMG, activeSpecialV2, rl.White)
				}

				switch checkPowerUpVehicleBlock {
				case "powerUpVehicleTL":

					powerUpVehicleV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, flyingRaddishIMG, powerUpVehicleV2, rl.White)
				}

				if playerDied == false {
					switch checkPlayerBlock {
					case "playerTL":
						dinoV2 = rl.NewVector2(float32(drawScreenX), float32(drawScreenY-2))
						shadowV2 = rl.NewVector2(float32(drawScreenX), float32(drawScreenY+16))
						if jumpActive == false && fallActive == false {
							rl.DrawTextureRec(imgs, shadowIMG, shadowV2, rl.White)
						}
						if exclamationOn {
							exclamationV2 := rl.NewVector2(float32(drawScreenX-4), float32(drawScreenY-20))
							rl.DrawTextureRec(imgs, exclamationIMG, exclamationV2, rl.White)
						}
						if propellorOn {
							propellorV2 := rl.NewVector2(float32(drawScreenX-4), float32(drawScreenY-10))
							rl.DrawTextureRec(imgs, propellorIMG, propellorV2, rl.White)
						}

						switch playerDirection {
						case "right":
							// draw player
							switch dinoType {
							case "greenDino":
								if invisibleOn == false {
									rl.DrawTextureRec(imgs, dinoGreenRIMG, dinoV2, rl.White)
								}
							case "redDino":
								if invisibleOn == false {
									rl.DrawTextureRec(imgs, dinoRedRIMG, dinoV2, rl.White)
								}
							case "yellowDino":
								if invisibleOn == false {
									rl.DrawTextureRec(imgs, dinoYellowRIMG, dinoV2, rl.White)
								}
							case "blueDino":
								if invisibleOn == false {
									rl.DrawTextureRec(imgs, dinoBlueRIMG, dinoV2, rl.White)
								}
							}
							// draw weapon
							switch currentPlayerWeapon {
							case "weapon1TL":
								weaponV2 := rl.NewVector2(float32(drawScreenX+14), float32(drawScreenY+2))
								if weapongameud {
									rl.DrawTextureRec(imgs, weapon1IMG, weaponV2, rl.White)
								} else {
									weaponV2.Y++
									rl.DrawTextureRec(imgs, weapon1IMG, weaponV2, rl.White)
								}
							case "weapon2TL":
								weaponV2 := rl.NewVector2(float32(drawScreenX+14), float32(drawScreenY+2))
								if weapongameud {
									rl.DrawTextureRec(imgs, weapon2IMG, weaponV2, rl.White)
								} else {
									weaponV2.Y++
									rl.DrawTextureRec(imgs, weapon2IMG, weaponV2, rl.White)
								}
							case "weapon3TL":
								weaponV2 := rl.NewVector2(float32(drawScreenX+14), float32(drawScreenY+2))
								if weapongameud {
									rl.DrawTextureRec(imgs, weapon3IMG, weaponV2, rl.White)
								} else {
									weaponV2.Y++
									rl.DrawTextureRec(imgs, weapon3IMG, weaponV2, rl.White)
								}
							case "weapon4TL":
								weaponV2 := rl.NewVector2(float32(drawScreenX+14), float32(drawScreenY+2))
								if weapongameud {
									rl.DrawTextureRec(imgs, weapon4IMG, weaponV2, rl.White)
								} else {
									weaponV2.Y++
									rl.DrawTextureRec(imgs, weapon4IMG, weaponV2, rl.White)
								}
							case "weapon5TL":
								weaponV2 := rl.NewVector2(float32(drawScreenX+14), float32(drawScreenY+2))
								if weapongameud {
									rl.DrawTextureRec(imgs, weapon5IMG, weaponV2, rl.White)
								} else {
									weaponV2.Y++
									rl.DrawTextureRec(imgs, weapon5IMG, weaponV2, rl.White)
								}
							case "weapon6TL":
								weaponV2 := rl.NewVector2(float32(drawScreenX+14), float32(drawScreenY+2))
								if weapongameud {
									rl.DrawTextureRec(imgs, weapon6IMG, weaponV2, rl.White)
								} else {
									weaponV2.Y++
									rl.DrawTextureRec(imgs, weapon6IMG, weaponV2, rl.White)
								}
							case "weapon7TL":
								weaponV2 := rl.NewVector2(float32(drawScreenX+14), float32(drawScreenY+2))
								if weapongameud {
									rl.DrawTextureRec(imgs, weapon7IMG, weaponV2, rl.White)
								} else {
									weaponV2.Y++
									rl.DrawTextureRec(imgs, weapon7IMG, weaponV2, rl.White)
								}
							}
						case "left":
							// draw player
							switch dinoType {
							case "greenDino":
								if invisibleOn == false {
									rl.DrawTextureRec(imgs, dinoGreenLIMG, dinoV2, rl.White)
								}
							case "redDino":
								if invisibleOn == false {
									rl.DrawTextureRec(imgs, dinoRedLIMG, dinoV2, rl.White)
								}
							case "yellowDino":
								if invisibleOn == false {
									rl.DrawTextureRec(imgs, dinoYellowLIMG, dinoV2, rl.White)
								}
							case "blueDino":
								if invisibleOn == false {
									rl.DrawTextureRec(imgs, dinoBlueLIMG, dinoV2, rl.White)
								}
							}
							// draw weapon

							switch currentPlayerWeapon {
							case "weapon1TL":
								weaponV2 := rl.NewVector2(float32(drawScreenX-15), float32(drawScreenY+2))
								if weapongameud {
									rl.DrawTextureRec(imgs, weapon1LIMG, weaponV2, rl.White)
								} else {
									weaponV2.Y++
									rl.DrawTextureRec(imgs, weapon1LIMG, weaponV2, rl.White)
								}
							case "weapon2TL":
								weaponV2 := rl.NewVector2(float32(drawScreenX-17), float32(drawScreenY+2))
								if weapongameud {
									rl.DrawTextureRec(imgs, weapon2LIMG, weaponV2, rl.White)
								} else {
									weaponV2.Y++
									rl.DrawTextureRec(imgs, weapon2LIMG, weaponV2, rl.White)
								}
							case "weapon3TL":
								weaponV2 := rl.NewVector2(float32(drawScreenX-15), float32(drawScreenY+2))
								if weapongameud {
									rl.DrawTextureRec(imgs, weapon3LIMG, weaponV2, rl.White)
								} else {
									weaponV2.Y++
									rl.DrawTextureRec(imgs, weapon3LIMG, weaponV2, rl.White)
								}
							case "weapon4TL":
								weaponV2 := rl.NewVector2(float32(drawScreenX-15), float32(drawScreenY+2))
								if weapongameud {
									rl.DrawTextureRec(imgs, weapon4LIMG, weaponV2, rl.White)
								} else {
									weaponV2.Y++
									rl.DrawTextureRec(imgs, weapon4LIMG, weaponV2, rl.White)
								}
							case "weapon5TL":
								weaponV2 := rl.NewVector2(float32(drawScreenX-15), float32(drawScreenY+2))
								if weapongameud {
									rl.DrawTextureRec(imgs, weapon5LIMG, weaponV2, rl.White)
								} else {
									weaponV2.Y++
									rl.DrawTextureRec(imgs, weapon5LIMG, weaponV2, rl.White)
								}
							case "weapon6TL":
								weaponV2 := rl.NewVector2(float32(drawScreenX-15), float32(drawScreenY+2))
								if weapongameud {
									rl.DrawTextureRec(imgs, weapon6LIMG, weaponV2, rl.White)
								} else {
									weaponV2.Y++
									rl.DrawTextureRec(imgs, weapon6LIMG, weaponV2, rl.White)
								}
							case "weapon7TL":
								weaponV2 := rl.NewVector2(float32(drawScreenX-15), float32(drawScreenY+2))
								if weapongameud {
									rl.DrawTextureRec(imgs, weapon7LIMG, weaponV2, rl.White)
								} else {
									weaponV2.Y++
									rl.DrawTextureRec(imgs, weapon7LIMG, weaponV2, rl.White)
								}
							}
						}
					}
				}

				switch checkTeleportBlock {

				case "teleportTL":
					teleportV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, teleportIMG, teleportV2, rl.White)
				}

				if cloudsOn {
					switch checkCloudsBlock {
					case "cloud":
						rl.DrawRectangle(drawScreenX+2, drawScreenY+2, 8, 8, rl.Fade(rl.DarkGray, 0.3))
						rl.DrawRectangle(drawScreenX, drawScreenY, 8, 8, rl.Fade(rl.White, 0.7))
					}
				}

				if tornadoOn {
					switch checkTornadoBlock {
					case "tornado":
						rl.DrawRectangle(drawScreenX+2, drawScreenY+2, 8, 8, rl.Fade(rl.DarkGray, 0.3))
						rl.DrawRectangle(drawScreenX, drawScreenY, 8, 8, rl.Fade(rl.White, 0.7))
					}
				}

				switch checkMeteorBlock {

				case "meteorTL":
					meteorV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, meteorIMG, meteorV2, rl.White)
				}

				switch checkEffectsBlock {
				case "bomb4":
					rl.DrawRectangle(drawScreenX+2, drawScreenY+2, 8, 8, rl.Fade(rl.DarkGray, 0.3))
					rl.DrawRectangle(drawScreenX, drawScreenY, 8, 8, rl.Fade(rl.White, 0.7))

				case "bomb3":
					rl.DrawRectangle(drawScreenX+2, drawScreenY+2, 6, 6, rl.Fade(rl.DarkGray, 0.3))
					rl.DrawRectangle(drawScreenX, drawScreenY, 6, 6, rl.Fade(rl.White, 0.7))

				case "bomb2":
					rl.DrawRectangle(drawScreenX+2, drawScreenY+2, 4, 4, rl.Fade(rl.DarkGray, 0.3))
					rl.DrawRectangle(drawScreenX, drawScreenY, 4, 4, rl.Fade(rl.White, 0.7))

				case "bomb1":
					rl.DrawRectangle(drawScreenX+2, drawScreenY+2, 2, 2, rl.Fade(rl.DarkGray, 0.3))
					rl.DrawRectangle(drawScreenX, drawScreenY, 2, 2, rl.Fade(rl.White, 0.7))
				}

				switch checkRainOfFireBlock {

				case "rainoffire":
					platformTileV2 := rl.NewVector2(float32(drawScreenX-8), float32(drawScreenY-8))
					rl.DrawTextureRec(imgs, flameIMG, platformTileV2, rl.White)

				}

				switch checkSnowBlock {
				case "snow1":
					snowV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, snow1IMG, snowV2, rl.Fade(rl.White, 0.8))
				case "snow2":
					snowV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, snow2IMG, snowV2, rl.Fade(rl.White, 0.8))
				case "snow3":
					snowV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, snow3IMG, snowV2, rl.Fade(rl.White, 0.8))
				case "snow4":
					snowV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, snow4IMG, snowV2, rl.Fade(rl.White, 0.8))
				case "snow5":
					snowV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, snow5IMG, snowV2, rl.Fade(rl.White, 0.8))
				case "snow6":
					snowV2 := rl.NewVector2(float32(drawScreenX), float32(drawScreenY))
					rl.DrawTextureRec(imgs, snow6IMG, snowV2, rl.Fade(rl.White, 0.8))

				}

				drawScreenCurrentBlock++
				drawScreenCurrentBlockWeather++
				drawScreenLineCount++
				drawScreenX += 8
				if drawScreenLineCount == 86 {
					drawScreenCurrentBlock += 8
					drawScreenCurrentBlockWeather += 9314
					drawScreenLineCount = 0
					drawScreenX = 0
					drawScreenY += 8
				}
			}
		} // end pauseOn
		rl.EndMode2D() // MARK: EndMode2D

		// MARK: draw screen layer 3 ZOOM OBJECTS

		if pauseOn == false {
			drawScreenCurrentBlock = 192
			drawScreenLineCount = 0
			drawScreenX = int32(0)
			drawScreenY = int32(0)

			for a := 0; a < 4128; a++ {

				checkMiniBossBlock := miniBossMAP[drawScreenCurrentBlock]

				switch checkMiniBossBlock {
				case "miniBossTL":

					currentMiniBossBlock = a
					//	rl.DrawRectangleLines(drawScreenX, drawScreenY, 8, 8, rl.Blue)
					switch miniBossType {
					case 1:
						drawscreenX2 := float32(drawScreenX)
						drawscreenX2 = drawscreenX2 / 1.5
						miniBossV2 := rl.NewVector2(drawscreenX2, float32(drawScreenY/2+40))
						rl.BeginMode2D(cameraMiniBoss)
						if miniBossRightLeft {
							rl.DrawTextureRec(imgs, miniBoss1RIMG, miniBossV2, rl.White)
						} else {
							rl.DrawTextureRec(imgs, miniBoss1IMG, miniBossV2, rl.White)
						}
						rl.EndMode2D()
					case 2:
						drawscreenX2 := float32(drawScreenX)
						drawscreenX2 = drawscreenX2 / 1.5
						miniBossV2 := rl.NewVector2(drawscreenX2, float32(drawScreenY/2+47))
						miniBossCirlceV2 := rl.NewVector2(drawscreenX2+13, float32(drawScreenY/2+63))
						rl.BeginMode2D(cameraMiniBoss)
						if miniBossRightLeft {
							if miniBossRollOn {
								rl.DrawCircleV(miniBossCirlceV2, 24, rl.Fade(rl.Magenta, 0.5))
								rl.DrawTextureRec(imgs, miniBoss2RRollIMG, miniBossV2, rl.White)
							} else {
								rl.DrawTextureRec(imgs, miniBoss2RIMG, miniBossV2, rl.White)
							}
						} else {
							if miniBossRollOn {
								rl.DrawCircleV(miniBossCirlceV2, 24, rl.Fade(rl.Magenta, 0.5))
								rl.DrawTextureRec(imgs, miniBoss2RollIMG, miniBossV2, rl.White)
							} else {
								rl.DrawTextureRec(imgs, miniBoss2IMG, miniBossV2, rl.White)
							}
						}
						rl.EndMode2D()
					case 3:
						drawscreenX2 := float32(drawScreenX)
						drawscreenX2 = drawscreenX2 / 1.5
						miniBossV2 := rl.NewVector2(drawscreenX2, float32(drawScreenY/2+49))
						rl.BeginMode2D(cameraMiniBoss)
						if miniBossRightLeft {
							if miniBossWallJumpOn {
								rl.DrawTextureRec(imgs, miniBoss3UpIMG, miniBossV2, rl.White)
							} else {
								rl.DrawTextureRec(imgs, miniBoss3RIMG, miniBossV2, rl.White)
							}
						} else {
							if miniBossWallJumpOn {
								rl.DrawTextureRec(imgs, miniBoss3RUpIMG, miniBossV2, rl.White)
							} else {
								rl.DrawTextureRec(imgs, miniBoss3IMG, miniBossV2, rl.White)
							}
						}
						rl.EndMode2D()
					case 4:
						drawscreenX2 := float32(drawScreenX)
						drawscreenX2 = drawscreenX2 / 1.5
						miniBossV2 := rl.NewVector2(drawscreenX2, float32(drawScreenY/2+48))
						rl.BeginMode2D(cameraMiniBoss)
						if miniBossRightLeft {
							rl.DrawTextureRec(imgs, miniBoss4RIMG, miniBossV2, rl.White)
						} else {
							rl.DrawTextureRec(imgs, miniBoss4IMG, miniBossV2, rl.White)
						}
						rl.EndMode2D()
					case 5:
						drawscreenX2 := float32(drawScreenX)
						drawscreenX2 = drawscreenX2 / 1.5
						miniBossV2 := rl.NewVector2(drawscreenX2, float32(drawScreenY/2+49))
						miniBossCirlceV2 := rl.NewVector2(drawscreenX2+13, float32(drawScreenY/2+63))
						rl.BeginMode2D(cameraMiniBoss)
						if miniBossRightLeft {
							if boss5RollOn {
								rl.DrawCircleV(miniBossCirlceV2, 24, rl.Fade(rl.Lime, 0.6))
								rl.DrawTextureRec(imgs, miniBoss5RRollIMG, miniBossV2, rl.White)
							} else {
								rl.DrawTextureRec(imgs, miniBoss5RIMG, miniBossV2, rl.White)
							}
						} else {
							if boss5RollOn {
								rl.DrawCircleV(miniBossCirlceV2, 24, rl.Fade(rl.Lime, 0.6))
								rl.DrawTextureRec(imgs, miniBoss5RollIMG, miniBossV2, rl.White)
							} else {
								rl.DrawTextureRec(imgs, miniBoss5IMG, miniBossV2, rl.White)
							}
						}
						rl.EndMode2D()
					case 6:
						drawscreenX2 := float32(drawScreenX)
						drawscreenX2 = drawscreenX2 / 1.5
						miniBossV2 := rl.NewVector2(drawscreenX2, float32(drawScreenY/2+36))
						rl.BeginMode2D(cameraMiniBoss)
						rl.DrawTextureRec(imgs, miniBoss6IMG, miniBossV2, rl.White)
						rl.EndMode2D()
					case 7:
						drawscreenX2 := float32(drawScreenX)
						drawscreenX2 = drawscreenX2 / 1.5
						miniBossV2 := rl.NewVector2(drawscreenX2, float32(drawScreenY/2+48))
						rl.BeginMode2D(cameraMiniBoss)
						if miniBossRightLeft {
							rl.DrawTextureRec(imgs, miniBoss7RIMG, miniBossV2, rl.White)
						} else {
							rl.DrawTextureRec(imgs, miniBoss7IMG, miniBossV2, rl.White)
						}
						rl.EndMode2D()
					}
				case "miniBoss":
					//	rl.DrawRectangleLines(drawScreenX, drawScreenY, 8, 8, rl.Magenta)
				}

				drawScreenCurrentBlock++
				drawScreenCurrentBlockWeather++
				drawScreenLineCount++
				drawScreenX += 8
				if drawScreenLineCount == 86 {
					drawScreenCurrentBlock += 8
					drawScreenCurrentBlockWeather += 9314
					drawScreenLineCount = 0
					drawScreenX = 0
					drawScreenY += 8
				}
			}
		} // end draw screen layer 3 PAUSE ON

		// MARK: drop new weapon
		if killzCount == dropWeaponCount {
			dropNewWeapon = true
			dropWeaponChange = rInt(3, 10)
			dropWeaponCount += dropWeaponChange

		}
		if dropNewWeapon {
			dropWeaponPosition := rInt(194, 272)

			weaponType := rInt(1, 8)

			weaponNameTL := "weapon" + strconv.Itoa(weaponType) + "TL"
			weaponName := "weapon" + strconv.Itoa(weaponType)

			weaponsMAP[dropWeaponPosition] = weaponNameTL
			weaponsMAP[dropWeaponPosition+1] = weaponName
			dropNewWeapon = false
		}

		// MARK: drop coin
		if killzCount == 5 && coinHasDropped1 == false {
			dropCoin = true
			coinHasDropped1 = true
		}
		if killzCount == 10 && coinHasDropped2 == false {
			dropCoin = true
			coinHasDropped2 = true
		}
		if killzCount == 20 && coinHasDropped3 == false {
			dropCoin = true
			coinHasDropped3 = true
		}
		if killzCount == 30 && coinHasDropped4 == false {
			dropCoin = true
			coinHasDropped4 = true
		}
		if killzCount == 40 && coinHasDropped5 == false {
			dropCoin = true
			coinHasDropped5 = true
		}
		if killzCount == 50 && coinHasDropped6 == false {
			dropCoin = true
			coinHasDropped6 = true
		}
		if killzCount == 55 && coinHasDropped7 == false {
			dropCoin = true
			coinHasDropped7 = true
		}
		if killzCount == 60 && coinHasDropped8 == false {
			dropCoin = true
			coinHasDropped8 = true
		}
		if killzCount == 65 && coinHasDropped9 == false {
			dropCoin = true
			coinHasDropped9 = true
		}
		if killzCount == 70 && coinHasDropped10 == false {
			dropCoin = true
			coinHasDropped10 = true
		}

		if dropCoin {
			coinBlockPosition := rInt(194, 272)
			objectsMAP[coinBlockPosition] = "coinTL"
			objectsMAP[coinBlockPosition+1] = "coin"
			objectsMAP[coinBlockPosition+94] = "coin"
			objectsMAP[coinBlockPosition+95] = "coin"
			dropCoinSound = true
			dropCoin = false
		}

		// MARK: draw screen foreground no zoom / no camera

		drawScreenCurrentBlock = 192
		drawScreenLineCount = 0
		drawScreenX = int32(0)
		drawScreenY = int32(0)

		if pauseOn == false {

			if frameCountGameStart%15 == 0 {
				if menuTextSpaceColor == rl.Red {
					menuTextSpaceColor = rl.Gold
				} else {
					menuTextSpaceColor = rl.Red
				}
			}

			rl.DrawText("f1 menu  |  f2 end level", 449, 27, 40, rl.White)
			rl.DrawText("f1 menu  |  f2 end level", 451, 25, 40, rl.Black)
			rl.DrawText("f1 menu  |  f2 end level", 453, 24, 40, menuTextSpaceColor)

			// draw fog
			if fogOn {
				rl.DrawTexture(fogImage, fogX, fogY, rl.Fade(rl.RayWhite, 0.2))
			}

			for a := 0; a < 4128; a++ {

				checkPlayerBlock := playerMAP[drawScreenCurrentBlock]
				checkDeadEnemiesBlock := deadEnemiesMAP[drawScreenCurrentBlock]
				checkRainBlock := rainMAP[drawScreenCurrentBlock]

				if rainOn {
					switch checkRainBlock {

					case "rain":
						rl.DrawCircle((drawScreenX*2)+4, (drawScreenY*2)+2, 2, rl.Fade(rl.DarkBlue, 0.4))
						rl.DrawCircle((drawScreenX*2)+4, (drawScreenY*2)+4, 3, rl.Fade(rl.DarkBlue, 0.4))
					}
				}

				if playerDied {
					switch checkPlayerBlock {
					case "deadPlayer":
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, 40, rl.Fade(rl.Red, 0.5))
					case "deadPlayer1":
						deadPlayerRadius := playerDiedCirclesMAP[0]
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Red, 0.6))
						rl.DrawCircleLines(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
					case "deadPlayer2":
						deadPlayerRadius := playerDiedCirclesMAP[1]
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Red, 0.6))
						rl.DrawCircleLines(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
					case "deadPlayer3":
						deadPlayerRadius := playerDiedCirclesMAP[2]
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Red, 0.6))
						rl.DrawCircleLines(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
					case "deadPlayer4":
						deadPlayerRadius := playerDiedCirclesMAP[3]
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Red, 0.6))
						rl.DrawCircleLines(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
					case "deadPlayer5":
						deadPlayerRadius := playerDiedCirclesMAP[4]
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Red, 0.6))
						rl.DrawCircleLines(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
					case "deadPlayer6":
						deadPlayerRadius := playerDiedCirclesMAP[5]
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Red, 0.6))
						rl.DrawCircleLines(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
					case "deadPlayer7":
						deadPlayerRadius := playerDiedCirclesMAP[6]
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Red, 0.6))
						rl.DrawCircleLines(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
					case "deadPlayer8":
						deadPlayerRadius := playerDiedCirclesMAP[7]
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Red, 0.6))
						rl.DrawCircleLines(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
					case "deadPlayer9":
						deadPlayerRadius := playerDiedCirclesMAP[8]
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Red, 0.6))
						rl.DrawCircleLines(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
					case "deadPlayer10":
						deadPlayerRadius := playerDiedCirclesMAP[9]
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))
						rl.DrawCircle(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Red, 0.6))
						rl.DrawCircleLines(drawScreenX*2, drawScreenY*2, deadPlayerRadius, rl.Fade(rl.Black, 0.4))

					}

				}

				// MARK: draw dead enemies

				switch checkDeadEnemiesBlock {
				case "deadEnemy1_4":

					deadEnemyCircles()

					deadEnemyRadius1 := deadEnemiesCirclesRadius[0]
					deadEnemies1V2 := deadEnemiesCirclesV2[0]
					deadEnemyRadius2 := deadEnemiesCirclesRadius[1]
					deadEnemies2V2 := deadEnemiesCirclesV2[1]
					deadEnemyRadius3 := deadEnemiesCirclesRadius[2]
					deadEnemies3V2 := deadEnemiesCirclesV2[2]
					deadEnemyRadius4 := deadEnemiesCirclesRadius[3]
					deadEnemies4V2 := deadEnemiesCirclesV2[3]
					deadEnemyRadius5 := deadEnemiesCirclesRadius[4]
					deadEnemies5V2 := deadEnemiesCirclesV2[4]
					deadEnemyRadius6 := deadEnemiesCirclesRadius[5]
					deadEnemies6V2 := deadEnemiesCirclesV2[5]
					deadEnemyRadius7 := deadEnemiesCirclesRadius[6]
					deadEnemies7V2 := deadEnemiesCirclesV2[6]
					deadEnemyRadius8 := deadEnemiesCirclesRadius[7]
					deadEnemies8V2 := deadEnemiesCirclesV2[7]
					deadEnemyRadius9 := deadEnemiesCirclesRadius[8]
					deadEnemies9V2 := deadEnemiesCirclesV2[8]
					deadEnemyRadius10 := deadEnemiesCirclesRadius[9]
					deadEnemies10V2 := deadEnemiesCirclesV2[9]
					deadEnemyRadius11 := deadEnemiesCirclesRadius[10]
					deadEnemies11V2 := deadEnemiesCirclesV2[10]
					deadEnemyRadius12 := deadEnemiesCirclesRadius[11]
					deadEnemies12V2 := deadEnemiesCirclesV2[11]
					deadEnemyRadius13 := deadEnemiesCirclesRadius[12]
					deadEnemies13V2 := deadEnemiesCirclesV2[12]
					deadEnemyRadius14 := deadEnemiesCirclesRadius[13]
					deadEnemies14V2 := deadEnemiesCirclesV2[13]
					deadEnemyRadius15 := deadEnemiesCirclesRadius[14]
					deadEnemies15V2 := deadEnemiesCirclesV2[14]

					rl.DrawCircleV(deadEnemies1V2, deadEnemyRadius1, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies2V2, deadEnemyRadius2, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies3V2, deadEnemyRadius3, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies4V2, deadEnemyRadius4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies5V2, deadEnemyRadius5, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies6V2, deadEnemyRadius6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies7V2, deadEnemyRadius7, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies8V2, deadEnemyRadius8, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies9V2, deadEnemyRadius9, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies10V2, deadEnemyRadius10, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies11V2, deadEnemyRadius11, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies12V2, deadEnemyRadius12, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies13V2, deadEnemyRadius13, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies14V2, deadEnemyRadius14, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies15V2, deadEnemyRadius15, rl.Fade(rl.Red, 0.8))

				case "deadEnemy1_3":

					deadEnemyRadius1 := deadEnemiesCirclesRadius[0]
					deadEnemies1V2 := deadEnemiesCirclesV2[0]
					deadEnemyRadius2 := deadEnemiesCirclesRadius[1]
					deadEnemies2V2 := deadEnemiesCirclesV2[1]
					deadEnemyRadius3 := deadEnemiesCirclesRadius[2]
					deadEnemies3V2 := deadEnemiesCirclesV2[2]
					deadEnemyRadius4 := deadEnemiesCirclesRadius[3]
					deadEnemies4V2 := deadEnemiesCirclesV2[3]
					deadEnemyRadius5 := deadEnemiesCirclesRadius[4]
					deadEnemies5V2 := deadEnemiesCirclesV2[4]
					deadEnemyRadius6 := deadEnemiesCirclesRadius[5]
					deadEnemies6V2 := deadEnemiesCirclesV2[5]
					deadEnemyRadius7 := deadEnemiesCirclesRadius[6]
					deadEnemies7V2 := deadEnemiesCirclesV2[6]
					deadEnemyRadius8 := deadEnemiesCirclesRadius[7]
					deadEnemies8V2 := deadEnemiesCirclesV2[7]
					deadEnemyRadius9 := deadEnemiesCirclesRadius[8]
					deadEnemies9V2 := deadEnemiesCirclesV2[8]
					deadEnemyRadius10 := deadEnemiesCirclesRadius[9]
					deadEnemies10V2 := deadEnemiesCirclesV2[9]
					deadEnemyRadius11 := deadEnemiesCirclesRadius[10]
					deadEnemies11V2 := deadEnemiesCirclesV2[10]
					deadEnemyRadius12 := deadEnemiesCirclesRadius[11]
					deadEnemies12V2 := deadEnemiesCirclesV2[11]
					deadEnemyRadius13 := deadEnemiesCirclesRadius[12]
					deadEnemies13V2 := deadEnemiesCirclesV2[12]
					deadEnemyRadius14 := deadEnemiesCirclesRadius[13]
					deadEnemies14V2 := deadEnemiesCirclesV2[13]
					deadEnemyRadius15 := deadEnemiesCirclesRadius[14]
					deadEnemies15V2 := deadEnemiesCirclesV2[14]

					rl.DrawCircleV(deadEnemies1V2, deadEnemyRadius1-2, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies2V2, deadEnemyRadius2-2, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies3V2, deadEnemyRadius3-2, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies4V2, deadEnemyRadius4-2, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies5V2, deadEnemyRadius5-2, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies6V2, deadEnemyRadius6-2, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies7V2, deadEnemyRadius7-2, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies8V2, deadEnemyRadius8-2, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies9V2, deadEnemyRadius9-2, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies10V2, deadEnemyRadius10-2, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies11V2, deadEnemyRadius11-2, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies12V2, deadEnemyRadius12-2, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies13V2, deadEnemyRadius13-2, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies14V2, deadEnemyRadius14-2, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies15V2, deadEnemyRadius15-2, rl.Fade(rl.Red, 0.8))

				case "deadEnemy1_2":

					deadEnemyRadius1 := deadEnemiesCirclesRadius[0]
					deadEnemies1V2 := deadEnemiesCirclesV2[0]
					deadEnemyRadius2 := deadEnemiesCirclesRadius[1]
					deadEnemies2V2 := deadEnemiesCirclesV2[1]
					deadEnemyRadius3 := deadEnemiesCirclesRadius[2]
					deadEnemies3V2 := deadEnemiesCirclesV2[2]
					deadEnemyRadius4 := deadEnemiesCirclesRadius[3]
					deadEnemies4V2 := deadEnemiesCirclesV2[3]
					deadEnemyRadius5 := deadEnemiesCirclesRadius[4]
					deadEnemies5V2 := deadEnemiesCirclesV2[4]
					deadEnemyRadius6 := deadEnemiesCirclesRadius[5]
					deadEnemies6V2 := deadEnemiesCirclesV2[5]
					deadEnemyRadius7 := deadEnemiesCirclesRadius[6]
					deadEnemies7V2 := deadEnemiesCirclesV2[6]
					deadEnemyRadius8 := deadEnemiesCirclesRadius[7]
					deadEnemies8V2 := deadEnemiesCirclesV2[7]
					deadEnemyRadius9 := deadEnemiesCirclesRadius[8]
					deadEnemies9V2 := deadEnemiesCirclesV2[8]
					deadEnemyRadius10 := deadEnemiesCirclesRadius[9]
					deadEnemies10V2 := deadEnemiesCirclesV2[9]
					deadEnemyRadius11 := deadEnemiesCirclesRadius[10]
					deadEnemies11V2 := deadEnemiesCirclesV2[10]
					deadEnemyRadius12 := deadEnemiesCirclesRadius[11]
					deadEnemies12V2 := deadEnemiesCirclesV2[11]
					deadEnemyRadius13 := deadEnemiesCirclesRadius[12]
					deadEnemies13V2 := deadEnemiesCirclesV2[12]
					deadEnemyRadius14 := deadEnemiesCirclesRadius[13]
					deadEnemies14V2 := deadEnemiesCirclesV2[13]
					deadEnemyRadius15 := deadEnemiesCirclesRadius[14]
					deadEnemies15V2 := deadEnemiesCirclesV2[14]

					rl.DrawCircleV(deadEnemies1V2, deadEnemyRadius1-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies1V2, deadEnemyRadius1-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies1V2.X), int32(deadEnemies1V2.Y), deadEnemyRadius1-4, rl.Black)
					rl.DrawCircleV(deadEnemies2V2, deadEnemyRadius2-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies2V2, deadEnemyRadius2-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies2V2.X), int32(deadEnemies2V2.Y), deadEnemyRadius2-4, rl.Black)
					rl.DrawCircleV(deadEnemies3V2, deadEnemyRadius3-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies3V2, deadEnemyRadius3-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies3V2.X), int32(deadEnemies3V2.Y), deadEnemyRadius3-4, rl.Black)
					rl.DrawCircleV(deadEnemies4V2, deadEnemyRadius4-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies4V2, deadEnemyRadius4-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies4V2.X), int32(deadEnemies4V2.Y), deadEnemyRadius4-4, rl.Black)
					rl.DrawCircleV(deadEnemies5V2, deadEnemyRadius5-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies5V2, deadEnemyRadius5-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies5V2.X), int32(deadEnemies5V2.Y), deadEnemyRadius5-4, rl.Black)
					rl.DrawCircleV(deadEnemies6V2, deadEnemyRadius6-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies6V2, deadEnemyRadius6-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies6V2.X), int32(deadEnemies6V2.Y), deadEnemyRadius6-4, rl.Black)
					rl.DrawCircleV(deadEnemies7V2, deadEnemyRadius7-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies7V2, deadEnemyRadius7-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies7V2.X), int32(deadEnemies7V2.Y), deadEnemyRadius7-4, rl.Black)
					rl.DrawCircleV(deadEnemies8V2, deadEnemyRadius8-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies8V2, deadEnemyRadius8-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies8V2.X), int32(deadEnemies8V2.Y), deadEnemyRadius8-4, rl.Black)
					rl.DrawCircleV(deadEnemies9V2, deadEnemyRadius9-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies9V2, deadEnemyRadius9-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies9V2.X), int32(deadEnemies9V2.Y), deadEnemyRadius9-4, rl.Black)
					rl.DrawCircleV(deadEnemies10V2, deadEnemyRadius10-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies10V2, deadEnemyRadius10-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies10V2.X), int32(deadEnemies10V2.Y), deadEnemyRadius10-4, rl.Black)
					rl.DrawCircleV(deadEnemies11V2, deadEnemyRadius11-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies11V2, deadEnemyRadius11-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies11V2.X), int32(deadEnemies11V2.Y), deadEnemyRadius11-4, rl.Black)
					rl.DrawCircleV(deadEnemies12V2, deadEnemyRadius12-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies12V2, deadEnemyRadius12-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies12V2.X), int32(deadEnemies12V2.Y), deadEnemyRadius12-4, rl.Black)
					rl.DrawCircleV(deadEnemies13V2, deadEnemyRadius13-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies13V2, deadEnemyRadius13-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies13V2.X), int32(deadEnemies13V2.Y), deadEnemyRadius13-4, rl.Black)
					rl.DrawCircleV(deadEnemies14V2, deadEnemyRadius14-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies14V2, deadEnemyRadius14-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies14V2.X), int32(deadEnemies14V2.Y), deadEnemyRadius14-4, rl.Black)
					rl.DrawCircleV(deadEnemies15V2, deadEnemyRadius15-4, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies15V2, deadEnemyRadius15-4, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies15V2.X), int32(deadEnemies15V2.Y), deadEnemyRadius15-4, rl.Black)

				case "deadEnemy1_1":

					deadEnemyRadius1 := deadEnemiesCirclesRadius[0]
					deadEnemies1V2 := deadEnemiesCirclesV2[0]
					deadEnemyRadius2 := deadEnemiesCirclesRadius[1]
					deadEnemies2V2 := deadEnemiesCirclesV2[1]
					deadEnemyRadius3 := deadEnemiesCirclesRadius[2]
					deadEnemies3V2 := deadEnemiesCirclesV2[2]
					deadEnemyRadius4 := deadEnemiesCirclesRadius[3]
					deadEnemies4V2 := deadEnemiesCirclesV2[3]
					deadEnemyRadius5 := deadEnemiesCirclesRadius[4]
					deadEnemies5V2 := deadEnemiesCirclesV2[4]
					deadEnemyRadius6 := deadEnemiesCirclesRadius[5]
					deadEnemies6V2 := deadEnemiesCirclesV2[5]
					deadEnemyRadius7 := deadEnemiesCirclesRadius[6]
					deadEnemies7V2 := deadEnemiesCirclesV2[6]
					deadEnemyRadius8 := deadEnemiesCirclesRadius[7]
					deadEnemies8V2 := deadEnemiesCirclesV2[7]
					deadEnemyRadius9 := deadEnemiesCirclesRadius[8]
					deadEnemies9V2 := deadEnemiesCirclesV2[8]
					deadEnemyRadius10 := deadEnemiesCirclesRadius[9]
					deadEnemies10V2 := deadEnemiesCirclesV2[9]
					deadEnemyRadius11 := deadEnemiesCirclesRadius[10]
					deadEnemies11V2 := deadEnemiesCirclesV2[10]
					deadEnemyRadius12 := deadEnemiesCirclesRadius[11]
					deadEnemies12V2 := deadEnemiesCirclesV2[11]
					deadEnemyRadius13 := deadEnemiesCirclesRadius[12]
					deadEnemies13V2 := deadEnemiesCirclesV2[12]
					deadEnemyRadius14 := deadEnemiesCirclesRadius[13]
					deadEnemies14V2 := deadEnemiesCirclesV2[13]
					deadEnemyRadius15 := deadEnemiesCirclesRadius[14]
					deadEnemies15V2 := deadEnemiesCirclesV2[14]

					rl.DrawCircleV(deadEnemies1V2, deadEnemyRadius1-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies1V2, deadEnemyRadius1-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies1V2.X), int32(deadEnemies1V2.Y), deadEnemyRadius1-6, rl.Black)
					rl.DrawCircleV(deadEnemies2V2, deadEnemyRadius2-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies2V2, deadEnemyRadius2-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies2V2.X), int32(deadEnemies2V2.Y), deadEnemyRadius2-6, rl.Black)
					rl.DrawCircleV(deadEnemies3V2, deadEnemyRadius3-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies3V2, deadEnemyRadius3-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies3V2.X), int32(deadEnemies3V2.Y), deadEnemyRadius3-6, rl.Black)
					rl.DrawCircleV(deadEnemies4V2, deadEnemyRadius4-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies4V2, deadEnemyRadius4-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies4V2.X), int32(deadEnemies4V2.Y), deadEnemyRadius4-6, rl.Black)
					rl.DrawCircleV(deadEnemies5V2, deadEnemyRadius5-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies5V2, deadEnemyRadius5-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies5V2.X), int32(deadEnemies5V2.Y), deadEnemyRadius5-6, rl.Black)
					rl.DrawCircleV(deadEnemies6V2, deadEnemyRadius6-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies6V2, deadEnemyRadius6-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies6V2.X), int32(deadEnemies6V2.Y), deadEnemyRadius6-6, rl.Black)
					rl.DrawCircleV(deadEnemies7V2, deadEnemyRadius7-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies7V2, deadEnemyRadius7-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies7V2.X), int32(deadEnemies7V2.Y), deadEnemyRadius7-6, rl.Black)
					rl.DrawCircleV(deadEnemies8V2, deadEnemyRadius8-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies8V2, deadEnemyRadius8-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies8V2.X), int32(deadEnemies8V2.Y), deadEnemyRadius8-6, rl.Black)
					rl.DrawCircleV(deadEnemies9V2, deadEnemyRadius9-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies9V2, deadEnemyRadius9-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies9V2.X), int32(deadEnemies9V2.Y), deadEnemyRadius9-6, rl.Black)
					rl.DrawCircleV(deadEnemies10V2, deadEnemyRadius10-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies10V2, deadEnemyRadius10-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies10V2.X), int32(deadEnemies10V2.Y), deadEnemyRadius10-6, rl.Black)
					rl.DrawCircleV(deadEnemies11V2, deadEnemyRadius11-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies11V2, deadEnemyRadius11-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies11V2.X), int32(deadEnemies11V2.Y), deadEnemyRadius11-6, rl.Black)
					rl.DrawCircleV(deadEnemies12V2, deadEnemyRadius12-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies12V2, deadEnemyRadius12-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies12V2.X), int32(deadEnemies12V2.Y), deadEnemyRadius12-6, rl.Black)
					rl.DrawCircleV(deadEnemies13V2, deadEnemyRadius13-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies13V2, deadEnemyRadius13-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies13V2.X), int32(deadEnemies13V2.Y), deadEnemyRadius13-6, rl.Black)
					rl.DrawCircleV(deadEnemies14V2, deadEnemyRadius14-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies14V2, deadEnemyRadius14-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies14V2.X), int32(deadEnemies14V2.Y), deadEnemyRadius14-6, rl.Black)
					rl.DrawCircleV(deadEnemies15V2, deadEnemyRadius15-6, rl.Fade(rl.Red, 0.8))
					rl.DrawCircleV(deadEnemies15V2, deadEnemyRadius15-6, rl.Fade(rl.Maroon, 0.1))
					rl.DrawCircleLines(int32(deadEnemies15V2.X), int32(deadEnemies15V2.Y), deadEnemyRadius15-6, rl.Black)

				}
				drawScreenCurrentBlock++
				drawScreenLineCount++
				drawScreenX += 8
				if drawScreenLineCount == 86 {
					drawScreenCurrentBlock += 8
					drawScreenLineCount = 0
					drawScreenX = 0
					drawScreenY += 8
				}
			}

			// MARK: ██████████████████████ current
			// MARK: update timers

			// miniboss hp pause timer
			if miniBossHPpause {
				if frameCount%120 == 0 {
					miniBossHPpause = false
				}
			}
			if miniBossHP <= 0 {
				minibossDie = true
			}
			if minibossDie {
				miniBossOn1 = false
				miniBossOn2 = false
				miniBossOn3 = false
				for a := 0; a < 5264; a++ {
					miniBossMAP[a] = ""
				}
				minibossDie = false
				miniBossHP = 5
				killMiniboss()
			}
			// MARK: update music
			if musicOn {
				rl.PlayMusicStream(levelTune)
				rl.UpdateMusicStream(levelTune)
			}
			if fxOn {
				if shootSound {
					sound := rl.LoadSound("thwip_1.mp3")
					rl.PlaySound(sound)
					shootSound = false
				}
				if coinCollectedSound {
					sound := rl.LoadSound("coin_pickup.mp3")
					rl.PlaySound(sound)
					coinCollectedSound = false
				}
				if dropCoinSound {
					sound := rl.LoadSound("zap_3.mp3")
					rl.PlaySound(sound)
					dropCoinSound = false
				}
				if hpLossSound {
					sound := rl.LoadSound("spring_board.mp3")
					rl.PlaySound(sound)
					hpLossSound = false
				}
				if enemyDieSound {
					sound := rl.LoadSound("crash_7.mp3")
					choose := rInt(1, 4)
					switch choose {
					case 1:
						sound = rl.LoadSound("crash_7.mp3")
					case 2:
						sound = rl.LoadSound("crash_4.mp3")
					case 3:
						sound = rl.LoadSound("crash_zap.mp3")
					}
					rl.PlaySound(sound)
					enemyDieSound = false
				}
			}

			// TESTING REMOVE
			// MARK: create powerup
			if rl.IsKeyPressed(rl.KeyZero) {
				powerUpDropped = true
			}
			if powerUpDropped {

				powerUpType := rInt(1, 11)

				powerUpType = 3 // TESTING REMOVE

				powerUpName := "powerUp" + strconv.Itoa(powerUpType)
				powerUpNameTL := powerUpName + "TL"

				powerUpPosition := 980

				powerUpsMAP[powerUpPosition] = powerUpNameTL
				powerUpsMAP[powerUpPosition+1] = powerUpName
				powerUpsMAP[powerUpPosition+94] = powerUpName
				powerUpsMAP[powerUpPosition+95] = powerUpName

				powerUpDropped = false
				powerUpHasDropped = true
			} // TESTING REMOVE

			// MARK: check difficulty on/off

			if easyDiffOn && averageDiffOn {
				easyDiffOn = true
				averageDiffOn = false
			}
			if averageDiffOn && difficultDiffOn {
				difficultDiffOn = true
				averageDiffOn = false
			}

			// MARK: block distortion
			if platformDistortionOn {
				if createDistortion == false {
					for a := 0; a < 4136; a++ {
						if platformsMAP[a] != "" {
							distort := rolldice()
							if distort == 6 {
								platformHOLDER := platformsMAP[a]
								platformsMAP[a] = ""
								levelMAP[a] = ""
								choose := flipcoin()
								if choose {
									platformsMAP[a-94] = platformHOLDER
									levelMAP[a-94] = "floor"
								} else {
									platformsMAP[a+94] = platformHOLDER
									levelMAP[a+94] = "floor"
								}
							}
						}
					}
					platformDistortionOn = false
				}
			}
			// MARK: shock blocks
			if shockBlocksOn {

				if frameCount%6 == 0 {
					shockBlockIMG.Y += 64
					if shockBlockIMG.Y >= 320 {
						shockBlockIMG.Y = 79
					}
				}

				if createShockBlocks == false {

					for a := 0; a < 5264; a++ {
						if platformsMAP[a] != "" {
							placeShockBlock := rolldice() + rolldice()

							if placeShockBlock == 12 {
								shockBlockMAP[a] = "shockBlock"
							}
						}
					}

					createShockBlocks = true

				}

			}

			// MARK: snow

			if snowOn {

				if frameCount%60 == 0 {
					snowTimer--
				}
				if snowTimer == 0 {
					createSnow = true
					snowTimer = 2
				}
			}

			if createSnow {

				snowBlock := rInt(4, 16) + 94
				snowNumber := rInt(8, 13)

				for a := 0; a < snowNumber; a++ {

					addUpDown := flipcoin()
					chooseSnowType := rInt(1, 7)

					if addUpDown {
						upDownChange := rInt(1, 4)

						switch chooseSnowType {
						case 1:
							snowMAP[snowBlock+(94*upDownChange)] = "snow1"
						case 2:
							snowMAP[snowBlock+(94*upDownChange)] = "snow2"
						case 3:
							snowMAP[snowBlock+(94*upDownChange)] = "snow3"
						case 4:
							snowMAP[snowBlock+(94*upDownChange)] = "snow4"
						case 5:
							snowMAP[snowBlock+(94*upDownChange)] = "snow5"
						case 6:
							snowMAP[snowBlock+(94*upDownChange)] = "snow6"
						}
					} else {
						switch chooseSnowType {
						case 1:
							snowMAP[snowBlock] = "snow1"
						case 2:
							snowMAP[snowBlock] = "snow2"
						case 3:
							snowMAP[snowBlock] = "snow3"
						case 4:
							snowMAP[snowBlock] = "snow4"
						case 5:
							snowMAP[snowBlock] = "snow5"
						case 6:
							snowMAP[snowBlock] = "snow6"
						}
					}
					snowBlock += rInt(6, 12)

					if snowBlock > 184 {
						break
					}
				}

				createSnow = false
			}

			// MARK: rain of fire
			if rainoffireOn {

				if frameCount%60 == 0 {
					rainoffireTimerFinal--
				}
				if rainoffireTimerFinal == 0 {
					rainoffireTimerFinal = 5
					rainoffireOn = false
					rainoffireTimer = 1
					rainoffireTimerOn = false
					for a := 0; a < 5264; a++ {
						rainoffireMAP[a] = ""
						platformsEffectsMAP[a] = ""
					}
				}

				if frameCountGameStart%6 == 0 {
					flameIMG.X += 25
					if flameIMG.X > 1240 {
						flameIMG.X = 1045
					}
				}

				if rainoffireTimerOn {
					if frameCount%60 == 0 {
						rainoffireTimer--
					}
					if rainoffireTimer == 0 {
						rainoffireTimer = 1
						rainoffireTimerOn = false
					}

				} else {
					rainoffireStartBlock := rInt(8, 16) + 94
					for a := 0; a < 10; a++ {

						rainoffireMAP[rainoffireStartBlock] = "rainoffire"
						rainoffireStartBlock += rInt(6, 12)
						if rainoffireStartBlock > 184 {
							break
						}
					}
					rainoffireTimerOn = true
				}
			}

			// MARK: meteor

			if meteorOn {

				if frameCountGameStart%6 == 0 {
					flameIMG.X += 25
					if flameIMG.X > 1240 {
						flameIMG.X = 1045
					}
				}

				if createMeteor == false {
					meteorLR = flipcoin()
					meteorStartBlock = rInt(20, 74)

					meteorMAP[meteorStartBlock] = "meteorTL"
					meteorMAP[meteorStartBlock+1] = "meteor"
					meteorMAP[meteorStartBlock+2] = "meteor"
					meteorMAP[meteorStartBlock+3] = "meteor"
					meteorMAP[meteorStartBlock+94] = "meteor"
					meteorMAP[meteorStartBlock+95] = "meteor"
					meteorMAP[meteorStartBlock+96] = "meteor"
					meteorMAP[meteorStartBlock+97] = "meteor"
					meteorMAP[meteorStartBlock+188] = "meteor"
					meteorMAP[meteorStartBlock+189] = "meteor"
					meteorMAP[meteorStartBlock+190] = "meteor"
					meteorMAP[meteorStartBlock+191] = "meteor"
					createMeteor = true
				}

				if frameCount%2 == 0 {
					moveMeteorDown()
				}

			}

			// MARK: tornado

			if tornadoOn {

				if frameCount%6 == 0 {
					createTornado = true
					moveTornado()
				}

				if createTornado == false {

					tornadoStartBlockHOLDER := tornadoStartBlock
					tornadoL := 1

					for {
						for a := 0; a < tornadoL; a++ {
							tornadoMAP[tornadoStartBlock] = "tornado"
							tornadoStartBlock++
						}
						tornadoStartBlock = tornadoStartBlockHOLDER
						tornadoStartBlock -= 94
						tornadoStartBlockHOLDER = tornadoStartBlock

						for a := 0; a < tornadoL; a++ {
							tornadoMAP[tornadoStartBlock] = "tornado"
							tornadoStartBlock++
						}

						tornadoStartBlock = tornadoStartBlockHOLDER
						tornadoLBlock = tornadoStartBlock
						tornadoRBlock = tornadoStartBlock + tornadoL

						tornadoStartBlock -= 94
						tornadoLChange := rInt(0, 3)
						tornadoStartBlock -= tornadoLChange
						tornadoL += (tornadoLChange * 2)

						tornadoStartBlockHOLDER = tornadoStartBlock

						tornadoHorizontal := tornadoStartBlock / 94

						if tornadoHorizontal <= 2 {

							break
						}

					}
					createTornado = true
					tornadoLR = flipcoin()

				}

			}

			// MARK: rain frog

			if frogRainOn {

				if rainFrogTimerOn {

					if frameCount%60 == 0 {
						rainFrogTimer--
						if rainFrogTimer == 0 {
							rainFrogTimerOn = false
							rainFrogTimer = 1
							rainFrogOn = true
						}
					}

				}

				if frameCount%6 == 0 {

					rainFrogIMG.X += 24
					if rainFrogIMG.X > 760 {
						rainFrogIMG.X = 668
					}
					rainFrogRIMG.X += 24
					if rainFrogRIMG.X > 865 {
						rainFrogRIMG.X = 776
					}

				}

				if rainFrogOn {

					rainFrogBlock := rInt(10, 20) + 94
					for a := 0; a < 10; a++ {
						chooseLR := flipcoin()
						if chooseLR {
							chooseDown := rInt(0, 3)
							rainFrogMAP[rainFrogBlock+(chooseDown*94)] = "rainFrog"
						} else {
							chooseDown := rInt(0, 3)
							rainFrogMAP[rainFrogBlock+(chooseDown*94)] = "rainFrogR"
						}
						rainFrogBlock += rInt(4, 9)
					}
					rainFrogOn = false
				}
			}

			// MARK: power ups

			if activeSpecialItem == "propellor" {
				if frameCountGameStart%5 == 0 {
					propellorIMG.X += 24
					if propellorIMG.X == 1509 {
						propellorIMG.X = 1413
					}
				}

			}

			// update power up images
			if frameCount%12 == 0 {
				bombIMG.X += 32
				if bombIMG.X >= 240 {
					bombIMG.X = 154
				}
				hpPowerUpIMG.X += 32
				if hpPowerUpIMG.X >= 90 {
					hpPowerUpIMG.X = 0
				}
				specialItemIMG.X += 32
				if specialItemIMG.X >= 60 {
					specialItemIMG.X = 0
				}

				randomPowerUpIMG.X += 32
				if randomPowerUpIMG.X >= 60 {
					randomPowerUpIMG.X = 0
				}
			}

			if powerUpCollected {

				switch powerUpCurrentActive {

				case 1:

					chooseRandomPowerUp := rolldice() + rolldice() + rolldice()

					chooseRandomPowerUp = 18 // TESTING REMOVE

					switch chooseRandomPowerUp {

					case 1, 2:
						camera.Zoom = 2.2
						camera.Target.Y = 10
						camera.Target.X = 40

						if powerUpRotation1 {
							if frameCount%2 == 0 {
								camera.Rotation -= 1.0
							}
							if camera.Rotation <= -10.0 {
								powerUpRotation1 = false
							}
						} else {
							if frameCount%2 == 0 {
								camera.Rotation += 1.0
							}
							if camera.Rotation >= 10.0 {
								powerUpRotation1 = true
							}
						}
						if frameCount%60 == 0 {
							powerUpTimer5--
						}
						if powerUpTimer5 == 0 {
							powerUpTimer5 = 5
							powerUpCollected = false
							powerUpCurrentActive = 0
							camera.Zoom = 2.0
							camera.Target.Y = 0
							camera.Target.X = 0
							camera.Rotation = 0.0
						}

					case 3, 4:
						playerHP = playerHPmax
						powerUpCollected = false
						powerUpCurrentActive = 0

					case 5, 6:
						chooseActiveSpecialItem := rInt(0, len(shopItemsMAP))
						activeSpecialItem = shopItemsMAP[chooseActiveSpecialItem]
						activeSpecialOn = true
						activeSpecialActive = true
						powerUpCollected = false
						powerUpCurrentActive = 0

					case 7, 8:
						camera.Zoom = 0.8
						camera.Target.Y = -250
						camera.Target.X = -500

						if frameCount%60 == 0 {
							powerUpTimer5--
						}
						if powerUpTimer5 == 0 {
							powerUpTimer5 = 5
							powerUpCollected = false
							powerUpCurrentActive = 0
							camera.Zoom = 2.0
							camera.Target.Y = 0
							camera.Target.X = 0
						}

					case 9, 10, 17:
						camera.Rotation = 180.0
						camera.Target.Y = 384
						camera.Target.X = 683

						if frameCount%60 == 0 {
							powerUpTimer5--
						}
						if powerUpTimer5 == 0 {
							powerUpTimer5 = 5
							powerUpCollected = false
							powerUpCurrentActive = 0
							camera.Target.Y = 0
							camera.Target.X = 0
							camera.Rotation = 0.0
						}
					case 11, 12, 16:
						invisibleOn = true
						if frameCount%60 == 0 {
							powerUpTimer5--
						}
						if powerUpTimer5 == 0 {
							invisibleOn = false
							powerUpTimer5 = 5
							powerUpCollected = false
							powerUpCurrentActive = 0
						}
					case 13, 14, 15:

						bombBlock = playerCurrentBlock + 2

						BombLength := rInt(10, 15)
						bombSize := rInt(6, 11)
						bombBlock -= BombLength / 2
						bombBlock -= 94 * (bombSize - 2)

						for a := 0; a < bombSize; a++ {

							for b := 0; b < BombLength; b++ {
								effectsMAP[bombBlock] = "bomb4"
								backgroundObjectsMAP[bombBlock] = ""

								if bombBlock < 4136 {
									if levelMAP[bombBlock] == "floor" {

										for c := 0; c < 43; c++ {
											platformType := rInt(1, 44)
											platformTypeName = "platform" + strconv.Itoa(platformType)
											platformTypeNameTL = platformTypeName + "TL"

											if platformsMAP[bombBlock] == platformTypeNameTL {

												levelMAP[bombBlock] = ""
												levelMAP[bombBlock+1] = ""
												levelMAP[bombBlock+94] = ""
												levelMAP[bombBlock+95] = ""

												platformsMAP[bombBlock] = ""
												platformsMAP[bombBlock+1] = ""
												platformsMAP[bombBlock+94] = ""
												platformsMAP[bombBlock+95] = ""

												backgroundObjectsMAP[bombBlock] = ""
												backgroundObjectsMAP[bombBlock+1] = ""
												backgroundObjectsMAP[bombBlock+94] = ""
												backgroundObjectsMAP[bombBlock+95] = ""

											} else if platformsMAP[bombBlock] == platformTypeName {

												if platformsMAP[bombBlock-1] == platformTypeNameTL {

													newBombBlock := bombBlock - 1

													levelMAP[newBombBlock] = ""
													levelMAP[newBombBlock+1] = ""
													levelMAP[newBombBlock+94] = ""
													levelMAP[newBombBlock+95] = ""

													platformsMAP[newBombBlock] = ""
													platformsMAP[newBombBlock+1] = ""
													platformsMAP[newBombBlock+94] = ""
													platformsMAP[newBombBlock+95] = ""

													backgroundObjectsMAP[newBombBlock] = ""
													backgroundObjectsMAP[newBombBlock+1] = ""
													backgroundObjectsMAP[newBombBlock+94] = ""
													backgroundObjectsMAP[newBombBlock+95] = ""

												} else if platformsMAP[bombBlock-95] == platformTypeNameTL {

													newBombBlock := bombBlock - 95

													levelMAP[newBombBlock] = ""
													levelMAP[newBombBlock+1] = ""
													levelMAP[newBombBlock+94] = ""
													levelMAP[newBombBlock+95] = ""

													platformsMAP[newBombBlock] = ""
													platformsMAP[newBombBlock+1] = ""
													platformsMAP[newBombBlock+94] = ""
													platformsMAP[newBombBlock+95] = ""

													backgroundObjectsMAP[newBombBlock] = ""
													backgroundObjectsMAP[newBombBlock+1] = ""
													backgroundObjectsMAP[newBombBlock+94] = ""
													backgroundObjectsMAP[newBombBlock+95] = ""

												} else if platformsMAP[bombBlock-94] == platformTypeNameTL {

													newBombBlock := bombBlock - 94

													levelMAP[newBombBlock] = ""
													levelMAP[newBombBlock+1] = ""
													levelMAP[newBombBlock+94] = ""
													levelMAP[newBombBlock+95] = ""

													platformsMAP[newBombBlock] = ""
													platformsMAP[newBombBlock+1] = ""
													platformsMAP[newBombBlock+94] = ""
													platformsMAP[newBombBlock+95] = ""

													backgroundObjectsMAP[newBombBlock] = ""
													backgroundObjectsMAP[newBombBlock+1] = ""
													backgroundObjectsMAP[newBombBlock+94] = ""
													backgroundObjectsMAP[newBombBlock+95] = ""
												}
											}
										}

									}

								}

								bombBlock++
							}
							bombBlock += 94
							bombBlock += rInt(-4, 5)
							BombLength = rInt(10, 15)

							for b := BombLength; b > 0; b-- {
								effectsMAP[bombBlock] = "bomb4"
								bombBlock--
							}
							bombBlock += 94
							bombBlock += rInt(-4, 5)
							BombLength = rInt(10, 15)
						}

						powerUpCollected = false
						powerUpCurrentActive = 0

					case 18:
						if playerHPmax < 6 {
							playerHPmax++
							powerUpCollected = false
							powerUpCurrentActive = 0
						} else if playerHPmax == 6 {
							rl.DrawText("player hp already max", 200, 200, 40, rl.White)
							if frameCount%60 == 0 {
								powerUpTimer5--
							}
							if powerUpTimer5 == 0 {
								powerUpTimer5 = 5
								powerUpCollected = false
								powerUpCurrentActive = 0
							}
						}

					} // end chooseRandomPowerUp
				case 2:
					playerHP = playerHPmax
					powerUpCollected = false
					powerUpCurrentActive = 0
				case 3:
					chooseActiveSpecialItem := rInt(0, len(shopItemsMAP))
					activeSpecialItem = shopItemsMAP[chooseActiveSpecialItem]
					activeSpecialOn = true
					activeSpecialActive = true
					powerUpCollected = false
					powerUpCurrentActive = 0
				case 4:
					bombBlock = playerCurrentBlock + 2

					BombLength := rInt(10, 15)
					bombSize := rInt(6, 11)
					bombBlock -= BombLength / 2
					bombBlock -= 94 * (bombSize - 2)

					for a := 0; a < bombSize; a++ {

						for b := 0; b < BombLength; b++ {
							effectsMAP[bombBlock] = "bomb4"

							if bombBlock < 4136 {
								if levelMAP[bombBlock] == "floor" {

									for c := 0; c < 43; c++ {
										platformType := rInt(1, 44)
										platformTypeName = "platform" + strconv.Itoa(platformType)
										platformTypeNameTL = platformTypeName + "TL"

										if platformsMAP[bombBlock] == platformTypeNameTL {

											levelMAP[bombBlock] = ""
											levelMAP[bombBlock+1] = ""
											levelMAP[bombBlock+94] = ""
											levelMAP[bombBlock+95] = ""

											platformsMAP[bombBlock] = ""
											platformsMAP[bombBlock+1] = ""
											platformsMAP[bombBlock+94] = ""
											platformsMAP[bombBlock+95] = ""

										} else if platformsMAP[bombBlock] == platformTypeName {

											if platformsMAP[bombBlock-1] == platformTypeNameTL {

												newBombBlock := bombBlock - 1

												levelMAP[newBombBlock] = ""
												levelMAP[newBombBlock+1] = ""
												levelMAP[newBombBlock+94] = ""
												levelMAP[newBombBlock+95] = ""

												platformsMAP[newBombBlock] = ""
												platformsMAP[newBombBlock+1] = ""
												platformsMAP[newBombBlock+94] = ""
												platformsMAP[newBombBlock+95] = ""
											} else if platformsMAP[bombBlock-95] == platformTypeNameTL {

												newBombBlock := bombBlock - 95

												levelMAP[newBombBlock] = ""
												levelMAP[newBombBlock+1] = ""
												levelMAP[newBombBlock+94] = ""
												levelMAP[newBombBlock+95] = ""

												platformsMAP[newBombBlock] = ""
												platformsMAP[newBombBlock+1] = ""
												platformsMAP[newBombBlock+94] = ""
												platformsMAP[newBombBlock+95] = ""
											} else if platformsMAP[bombBlock-94] == platformTypeNameTL {

												newBombBlock := bombBlock - 94

												levelMAP[newBombBlock] = ""
												levelMAP[newBombBlock+1] = ""
												levelMAP[newBombBlock+94] = ""
												levelMAP[newBombBlock+95] = ""

												platformsMAP[newBombBlock] = ""
												platformsMAP[newBombBlock+1] = ""
												platformsMAP[newBombBlock+94] = ""
												platformsMAP[newBombBlock+95] = ""
											}
										}
									}

								}

							}

							bombBlock++
						}
						bombBlock += 94
						bombBlock += rInt(-4, 5)
						BombLength = rInt(10, 15)

						for b := BombLength; b > 0; b-- {
							effectsMAP[bombBlock] = "bomb4"
							bombBlock--
						}
						bombBlock += 94
						bombBlock += rInt(-4, 5)
						BombLength = rInt(10, 15)
					}

					powerUpCollected = false
					powerUpCurrentActive = 0
				case 5:
					invisibleOn = true
					if frameCount%60 == 0 {
						powerUpTimer5--
					}
					if powerUpTimer5 == 0 {
						invisibleOn = false
						powerUpTimer5 = 5
						powerUpCollected = false
						powerUpCurrentActive = 0
					}
				case 6:
					playerHP = playerHPmax
					powerUpCollected = false
					powerUpCurrentActive = 0
				case 7:
					playerHP = playerHPmax
					powerUpCollected = false
					powerUpCurrentActive = 0
				case 8:
					playerHP = playerHPmax
					powerUpCollected = false
					powerUpCurrentActive = 0
				case 9:
					playerHP = playerHPmax
					powerUpCollected = false
					powerUpCurrentActive = 0
				case 10:
					playerHP = playerHPmax
					powerUpCollected = false
					powerUpCurrentActive = 0

				} //end powerUpCurrentActive
			} // end powerUpCollected

			// update enemy5 shields
			if frameCount%15 == 0 {

				if enemy5Shield1 {
					enemy5Shield2 = true
					enemy5Shield1 = false
				} else if enemy5Shield2 {
					enemy5Shield2 = false
					enemy5Shield3 = true
				} else if enemy5Shield3 {
					enemy5Shield3 = false
					enemy5Shield1 = true
				}

			}

			// MARK: update enemy images
			switch enemyTypeGenerate {

			case 1:
				if frameCount%9 == 0 {
					enemy1IMG.X += 32
					if enemy1IMG.X >= 1024 {
						enemy1IMG.X = 514
					}
				}
			case 2:
				if frameCount%9 == 0 {
					enemy2IMG.X += 32
					if enemy2IMG.X >= 2039 {
						enemy2IMG.X = 1591
					}
				}
			case 3:
				if frameCount%6 == 0 {
					enemy3IMG.X += 44
					if enemy3IMG.X >= 1592 {
						enemy3IMG.X = 1108
					}
				}
			case 4:
				if frameCount%6 == 0 {
					enemy4IMG.X += 32
					if enemy4IMG.X >= 2035 {
						enemy4IMG.X = 1587
					}
				}
			case 5:
				if frameCount%4 == 0 {
					enemy5IMG.X += 64
					if enemy5IMG.X >= 2048 {
						enemy5IMG.X = 1348
					}
				}
			case 6:
				if frameCount%3 == 0 {
					enemy6IMG.X += 46
					if enemy6IMG.X >= 1456 {
						enemy6IMG.X = 1134
					}
				}
			case 7:
				if frameCount%6 == 0 {
					enemy7URIMG.Y -= 44
					if enemy7URIMG.Y <= 654 {
						enemy7URIMG.Y = 970
					}
					enemy7ULIMG.Y -= 44
					if enemy7ULIMG.Y <= 654 {
						enemy7ULIMG.Y = 972
					}
					enemy7IMG.X += 44
					if enemy7IMG.X >= 1593 {
						enemy7IMG.X = 1241
					}
					enemy7UPIMG.X += 44
					if enemy7UPIMG.X >= 1560 {
						enemy7UPIMG.X = 1218
					}

				}

			case 8:
				if frameCount%4 == 0 {
					enemy8IMG.X += 44
					if enemy8IMG.X >= 2040 {
						enemy8IMG.X = 1600
					}
				}
			case 9:
				if frameCount%4 == 0 {
					enemy9IMG.X += 52
					if enemy9IMG.X >= 2041 {
						enemy9IMG.X = 1469
					}
				}
			}

			// MARK: check enemy5 shield  collision
			if hpLossPause == false {
				if enemyTypeGenerate == 5 {

					playerPlus4L := enemiesMAP[playerCurrentBlock-8]
					playerPlus4R := enemiesMAP[playerCurrentBlock+9]

					if playerPlus4L == "enemy" {
						playerHP--
						hpLossPause = true
						hpLossSound = true
					}
					if playerPlus4R == "enemy" || playerPlus4R == "enemyTL" {
						playerHP--
						hpLossPause = true
						hpLossSound = true
					}
				}
			}

			// MARK: check enemy kills vs enemy screen count
			onScreenEnemiesCount = 0

			if frameCount%60 == 0 {
				for a := 0; a < 5264; a++ {
					checkEnemyBlock := enemiesMAP[a]

					if checkEnemyBlock == "enemyTL" {
						onScreenEnemiesCount++
					}
				}
				if onScreenEnemiesCount < enemiesScreenCount {
					equalizeEnemyNumbers = true
				}
			}
			if equalizeEnemyNumbers {
				for {

					enemiesScreenCount--
					if enemiesScreenCount == onScreenEnemiesCount {
						equalizeEnemyNumbers = false
						break
					}
				}

			}

			// MARK: check stuck enemies

			// MARK: check enemies stuck on platform

			for a := 0; a < 2350; a++ {

				checkEnemyBlock := enemiesMAP[a]
				if checkEnemyBlock == "enemyTL" {
					if platformsMAP[a+(94*2)] == "movingPlatformDownTL" || platformsMAP[a+(94*2)] == "movingPlatformDown" || platformsMAP[a+(94*2)] == "movingPlatformUpTL" || platformsMAP[a+(94*2)] == "movingPlatformUp" || platformsMAP[a+((94*2)+1)] == "movingPlatformDownTL" || platformsMAP[a+((94*2)+1)] == "movingPlatformDown" || platformsMAP[a+((94*2)+1)] == "movingPlatformUpTL" || platformsMAP[a+((94*2)+1)] == "movingPlatformUp" {

						chooseJumpOffPlatform := flipcoin()

						if chooseJumpOffPlatform {
							chooseLR := flipcoin()
							chooseLR2 := flipcoin()

							if chooseLR {
								enemiesMovementMAP[a] = "jump4"
							} else {
								if chooseLR2 {
									clearEnemyBlock = a
									movementHOLDER := enemiesMovementMAP[a]
									clearEnemy()
									enemiesMovementMAP[a+1] = movementHOLDER
									moveEnemyRight()
								} else {
									clearEnemyBlock = a
									movementHOLDER := enemiesMovementMAP[a]
									clearEnemy()
									enemiesMovementMAP[a-1] = movementHOLDER
									moveEnemyLeft()

								}
							}
						}
					}
				}

			}

			if frameCount%6 == 0 {
				if checkEnemy1On {

					currentEnemyCheck := 0
					for a := 0; a < 5264; a++ {
						checkEnemyBlock := enemiesMAP[a]

						if checkEnemyBlock == "enemyTL" {

							enemyPosition1MAP[currentEnemyCheck] = a
							currentEnemyCheck++

						}
					}

					checkEnemy1On = false
					checkEnemy2On = true
					currentEnemyCheck = 0
				} else if checkEnemy2On {
					currentEnemyCheck := 0
					for a := 0; a < 5264; a++ {
						checkEnemyBlock := enemiesMAP[a]

						if checkEnemyBlock == "enemyTL" {

							enemyPosition2MAP[currentEnemyCheck] = a
							currentEnemyCheck++

						}
					}
					checkEnemy2On = false
					checkEnemy3On = true
					currentEnemyCheck = 0
				} else if checkEnemy3On {
					currentEnemyCheck := 0
					for a := 0; a < 5264; a++ {
						checkEnemyBlock := enemiesMAP[a]

						if checkEnemyBlock == "enemyTL" {
							enemyPosition3MAP[currentEnemyCheck] = a
							currentEnemyCheck++

						}
					}
					checkEnemy3On = false
					checkEnemy4On = true
					currentEnemyCheck = 0
				} else if checkEnemy4On {

					for a := 0; a < 20; a++ {

						if enemyPosition1MAP[a] > 376 {

							if enemyPosition1MAP[a] == enemyPosition2MAP[a] && enemyPosition2MAP[a] == enemyPosition3MAP[a] {

								moveEnemyBlock := enemyPosition1MAP[a]

								if enemiesMAP[moveEnemyBlock] == "enemyTL" {

									moveEnemyBlock -= rInt(2, 5) * 94
									moveEnemyBlock += rInt(-3, 4)

									if levelMAP[moveEnemyBlock] == "floor" {

										for {

											moveEnemyBlockHorizontal := moveEnemyBlock / 94

											if moveEnemyBlockHorizontal > 10 {
												moveEnemyBlock -= 94 * 2
											} else {
												moveEnemyBlock += 2
												moveEnemyBlock += 94
											}

											if levelMAP[moveEnemyBlock] != "floor" && levelMAP[moveEnemyBlock+94] != "floor" {
												break
											}
										}
									}

									oldMoveEnemyBlock := enemyPosition1MAP[a]

									clearEnemyBlock = oldMoveEnemyBlock
									clearEnemy()

									drawNewEnemyPosition(moveEnemyBlock)

									enemiesMovementMAP[moveEnemyBlock] = "enemy4Special2"

								}

							}
						}
					}
					checkEnemy4On = false
					checkEnemy5On = true

				} else if checkEnemy5On {

					for a := 0; a < 20; a++ {

						enemyPosition1MAP[a] = 0
						enemyPosition2MAP[a] = 0
						enemyPosition3MAP[a] = 0

					}

					checkEnemy5On = false
					checkEnemy1On = true

				}

			}

			// MARK: check enemy stuck on block / check no enemy movement position
			for a := 0; a < 5264; a++ {

				checkEnemyBlock := enemiesMAP[a]

				if checkEnemyBlock == "enemyTL" {

					if levelMAP[a] == "floor" || levelMAP[a+1] == "floor" || levelMAP[a+94] == "floor" || levelMAP[a+95] == "floor" {
						checkEnemyStuckOnBlock(a)
					}

					movementBlock := enemiesMovementMAP[a]

					if movementBlock == "" {
						enemiesMovementMAP[a] = "jump4"
					}

				}

			}

			// MARK: falling blocks timer
			if fallingBlocksOn {
				if frameCount%360 == 0 {
					cFALLINGBLOCKS()
				}
			}

			// MARK: create earthquake
			if earthquakesActive {
				if earthquakeTimerText1 == secondsCountdown || earthquakeTimerText2 == secondsCountdown || earthquakeTimer3 == secondsCountdown {
					earthquakeTextOn = true
					earthquakeCountdown = 3
					earthquakesOn = true
				}
				if earthquakesOn == true {

					if earthquakeTextOn {
						if frameCount%60 == 0 {
							earthquakeCountdown--
						}
						earthquakeCountdownTEXT := strconv.Itoa(earthquakeCountdown)
						rl.DrawText("earthquake in", screenW/4-3, screenH/6+3, 80, rl.Black)
						rl.DrawText("earthquake in", screenW/4-1, screenH/6+1, 80, rl.Yellow)
						rl.DrawText("earthquake in", screenW/4, screenH/6, 80, rl.Magenta)
						rl.DrawText(earthquakeCountdownTEXT, screenW/2+257, screenH/6+3, 80, rl.Black)
						rl.DrawText(earthquakeCountdownTEXT, screenW/2+259, screenH/6+1, 80, rl.Yellow)
						rl.DrawText(earthquakeCountdownTEXT, screenW/2+260, screenH/6, 80, rl.Magenta)
					}

					if earthquakeTimer1 == secondsCountdown || earthquakeTimer2 == secondsCountdown || earthquakeTimer3 == secondsCountdown {

						earthquakeTextOn = false

						if earthquakeCount != 0 {

							placeEarthquake := rInt(4234, 4320)

							earthquakeSize := rInt(4, 13)
							earthquakeArea := earthquakeSize * 8
							earthquakeLineCount := 0

							for a := 0; a < earthquakeArea; a++ {

								earthquakesMAP[placeEarthquake] = "earthquake"
								levelMAP[placeEarthquake] = "earthquake"
								placeEarthquake++
								earthquakeLineCount++
								if earthquakeLineCount == earthquakeSize {
									earthquakeLineCount = 0
									placeEarthquake += (94 - earthquakeSize) + (rInt(-1, 2))
								}

							}
							screenShake = true
							earthquakeCount--
							earthquakesOn = false
						} else {
							earthquakesOn = false
						}
					}
				}

				// earthquake check
				if playerGroundBlockL == "earthquake" && playerGroundBlockR == "earthquake" {
					earthquakeFallOn = true
					fallActive = true
				}
				if fallActive == true {

					if playerGroundBlockL != "floor" && playerGroundBlockR != "floor" && playerCurrentBlock < 5076 {
						clearPlayer()
						playerCurrentBlock += 94
					} else if playerCurrentBlock > 5076 {
						fallActive = false
					}

				}
			}

			if miniBossOn1 || miniBossOn2 || miniBossOn3 {
				// update miniboss image
				if frameCount%3 == 0 {
					miniBoss1IMG.X += 34
					if miniBoss1IMG.X >= 409 {
						miniBoss1IMG.X = 1
					}
					miniBoss1RIMG.X += 34
					if miniBoss1RIMG.X >= 409 {
						miniBoss1RIMG.X = 1
					}
					miniBoss2IMG.X += 32
					if miniBoss2IMG.X >= 374 {
						miniBoss2IMG.X = 0
					}
					miniBoss2RIMG.X += 32
					if miniBoss2RIMG.X >= 374 {
						miniBoss2RIMG.X = 0
					}
					miniBoss2RollIMG.X += 32
					if miniBoss2RollIMG.X >= 582 {
						miniBoss2RollIMG.X = 395
					}
					miniBoss2RRollIMG.X += 32
					if miniBoss2RRollIMG.X >= 582 {
						miniBoss2RRollIMG.X = 395
					}
					miniBoss3IMG.X += 32
					if miniBoss3IMG.X >= 374 {
						miniBoss3IMG.X = 0
					}
					miniBoss3RIMG.X += 32
					if miniBoss3RIMG.X >= 374 {
						miniBoss3RIMG.X = 0
					}
					miniBoss5IMG.X += 32
					if miniBoss5IMG.X >= 380 {
						miniBoss5IMG.X = 0
					}
					miniBoss5RIMG.X += 32
					if miniBoss5RIMG.X >= 380 {
						miniBoss5RIMG.X = 0
					}
					miniBoss5RollIMG.X += 32
					if miniBoss5RollIMG.X >= 420 {
						miniBoss5RollIMG.X = 238
					}
					miniBoss5RRollIMG.X += 32
					if miniBoss5RRollIMG.X >= 420 {
						miniBoss5RRollIMG.X = 238
					}
					miniBoss6IMG.X += 40
					if miniBoss6IMG.X >= 1370 {
						miniBoss6IMG.X = 1050
					}
					miniBoss7IMG.X += 36
					if miniBoss7IMG.X >= 990 {
						miniBoss7IMG.X = 640
					}
					miniBoss7RIMG.X += 36
					if miniBoss7RIMG.X >= 989 {
						miniBoss7RIMG.X = 639
					}
				}
				if frameCount%5 == 0 {
					miniBoss4IMG.X += 32
					if miniBoss4IMG.X >= 220 {
						miniBoss4IMG.X = 0
					}
					miniBoss4RIMG.X += 32
					if miniBoss4RIMG.X >= 220 {
						miniBoss4RIMG.X = 0
					}
				}
			}

			// MARK: frameCount%3 timer
			if frameCount%3 == 0 {

				// update rain
				upRAIN()

				// fog movement
				if fogRightLeft {
					fogX++
				} else {
					fogX--
				}
				if fogX < (0 - screenW*2) {
					fogRightLeft = true
				}

				// right down movement
				for a := 5263; a >= 0; a-- {

					checkEnemyBlock := enemiesMAP[a]
					checkPowerUpVehicleBlock := powerUpVehicleMAP[a]
					checkPowerUpBlock := powerUpsMAP[a]
					checkCoinBlock := objectsMAP[a]
					checkPlatformsBlock := platformsMAP[a]
					checkDropWeaponBlock := weaponsMAP[a]
					checkRainBlock := rainMAP[a]
					checkMiniBossBlock := miniBossMAP[a]
					checkMiniBossEffectsBlock := miniBossEffectsMAP[a]
					checkFallingBlock := levelMAP[a]
					checkActiveSpecialBlock := activeSpecialMAP[a]
					checkLevelBlock := levelMAP[a]
					checkPlatformEffectsBlock := platformsEffectsMAP[a]
					checkRainFrogBlock := rainFrogMAP[a]
					checkRainOfFireBlock := rainoffireMAP[a]
					checkSnowBlock := snowMAP[a]

					// update horizontal platforms
					if horizplatlr == false {
						if checkPlatformsBlock == "movingplatH8" {
							horizplatH = a / 94
							horizplatV = a - (horizplatH * 94)

							if horizplatV > 88 {
								horizplatlr = true
							} else {
								platformsMAP[a] = ""
								platformsMAP[a-1] = ""
								platformsMAP[a-2] = ""
								platformsMAP[a-3] = ""
								platformsMAP[a-4] = ""
								platformsMAP[a-5] = ""
								platformsMAP[a-6] = ""
								platformsMAP[a-7] = ""

								levelMAP[a] = ""
								levelMAP[a-1] = ""
								levelMAP[a-2] = ""
								levelMAP[a-3] = ""
								levelMAP[a-4] = ""
								levelMAP[a-5] = ""
								levelMAP[a-6] = ""
								levelMAP[a-7] = ""

								platformsMAP[a+1] = "movingplatH8"
								platformsMAP[a] = "movingplatH7"
								platformsMAP[a-1] = "movingplatH6"
								platformsMAP[a-2] = "movingplatH5"
								platformsMAP[a-3] = "movingplatH4"
								platformsMAP[a-4] = "movingplatH3"
								platformsMAP[a-5] = "movingplatH2"
								platformsMAP[a-6] = "movingplatH1"

								levelMAP[a+1] = "floor"
								levelMAP[a] = "floor"
								levelMAP[a-1] = "floor"
								levelMAP[a-2] = "floor"
								levelMAP[a-3] = "floor"
								levelMAP[a-4] = "floor"
								levelMAP[a-5] = "floor"
								levelMAP[a-6] = "floor"

							}
						}
					}
					// update snow
					if snowOn {
						if checkSnowBlock == "snow1" {

							snowHorizontal := a / 94

							if snowHorizontal < 44 {
								if frameCount%15 == 0 {
									snowMAP[a] = ""
									snowMAP[a+94*3] = "snow1"
								}
							} else {
								snowMAP[a] = ""
							}
						} else if checkSnowBlock == "snow2" {

							snowHorizontal := a / 94

							if snowHorizontal < 44 {
								if frameCount%15 == 0 {
									snowMAP[a] = ""
									snowMAP[a+94*3] = "snow2"
								}
							} else {
								snowMAP[a] = ""
							}
						} else if checkSnowBlock == "snow3" {

							snowHorizontal := a / 94

							if snowHorizontal < 44 {
								if frameCount%15 == 0 {
									snowMAP[a] = ""
									snowMAP[a+94*3] = "snow3"
								}
							} else {
								snowMAP[a] = ""
							}
						} else if checkSnowBlock == "snow4" {

							snowHorizontal := a / 94

							if snowHorizontal < 44 {
								if frameCount%15 == 0 {
									snowMAP[a] = ""
									snowMAP[a+94*2] = "snow4"
								}
							} else {
								snowMAP[a] = ""
							}
						} else if checkSnowBlock == "snow5" {

							snowHorizontal := a / 94

							if snowHorizontal < 44 {
								if frameCount%15 == 0 {
									snowMAP[a] = ""
									snowMAP[a+94*2] = "snow5"
								}
							} else {
								snowMAP[a] = ""
							}
						} else if checkSnowBlock == "snow6" {

							snowHorizontal := a / 94

							if snowHorizontal < 44 {
								if frameCount%15 == 0 {
									snowMAP[a] = ""
									snowMAP[a+94*3] = "snow6"
								}
							} else {
								snowMAP[a] = ""
							}
						}

					}

					// update rain of fire
					if rainoffireOn {
						if checkRainOfFireBlock == "rainoffire" {

							rainoffireHorizontal := a / 94

							if rainoffireHorizontal < 42 && levelMAP[a+94] != "floor" {
								if frameCount%6 == 0 {
									rainoffireMAP[a] = ""
									rainoffireMAP[a+94] = "rainoffire"
								}
							} else if rainoffireHorizontal < 42 && levelMAP[a+94] == "floor" {
								platformsEffectsMAP[a+94] = "fire6"

							} else if rainoffireHorizontal >= 42 {
								rainoffireMAP[a] = ""
							}
						}
					}

					// update rain frog
					if frogRainOn {
						if checkRainFrogBlock == "rainFrog" {
							rainFrogHorizontal := a / 94
							if rainFrogHorizontal < 42 {
								if levelMAP[a+94*3] != "floor" {
									rainFrogMAP[a] = ""
									rainFrogMAP[a+94] = "rainFrog"
								}
							}
						}
						if checkRainFrogBlock == "rainFrogR" {
							rainFrogHorizontal := a / 94
							rainFrogVertical := a - (94 * rainFrogHorizontal)
							if rainFrogHorizontal < 42 && levelMAP[a+94*3] != "floor" {
								rainFrogMAP[a] = ""
								rainFrogMAP[a+94] = "rainFrogR"
							} else if rainFrogVertical < 90 {
								rainFrogMAP[a] = ""
								rainFrogMAP[a+1] = "rainFrogR"
							} else if rainFrogVertical >= 90 {
								rainFrogMAP[a] = ""
								if rainFrogTimerOn == false {
									rainFrogOn = true
									rainFrogTimerOn = true
								}
							}
						}
					}

					// update active special effects

					if checkActiveSpecialBlock == "apple" && checkEnemyBlock == "enemy" && enemiesMAP[a-94] == "enemyTL" {
						clearEnemyBlock = a - 94
						clearEnemy()
						killEnemy()
					}
					if checkActiveSpecialBlock == "apple" && checkEnemyBlock == "enemyTL" {
						clearEnemyBlock = a
						clearEnemy()
						killEnemy()
					}
					if checkActiveSpecialBlock == "fireball" && checkEnemyBlock == "enemy" && enemiesMAP[a-1] == "enemyTL" {
						clearEnemyBlock = a - 1
						clearEnemy()
						killEnemy()
					}
					if checkActiveSpecialBlock == "fireballTL" && checkEnemyBlock == "enemy" && enemiesMAP[a-1] == "enemyTL" {
						clearEnemyBlock = a - 1
						clearEnemy()
						killEnemy()
					}
					if checkActiveSpecialBlock == "fireballTL" && checkEnemyBlock == "enemy" && enemiesMAP[a-95] == "enemyTL" {
						clearEnemyBlock = a - 95
						clearEnemy()
						killEnemy()
					}
					if checkActiveSpecialBlock == "poisonball" && checkEnemyBlock == "enemy" && enemiesMAP[a-94] == "enemyTL" {
						clearEnemyBlock = a - 94
						clearEnemy()
						killEnemy()
					}
					if checkActiveSpecialBlock == "poisonballTL" && checkEnemyBlock == "enemyTL" {
						clearEnemyBlock = a
						clearEnemy()
						killEnemy()
					}

					if frameCount%30 == 0 {
						if checkPlatformEffectsBlock == "fire" {
							platformsEffectsMAP[a] = "fire2"
						} else if checkPlatformEffectsBlock == "fire2" {
							platformsEffectsMAP[a] = "fire3"
						} else if checkPlatformEffectsBlock == "fire3" {
							platformsEffectsMAP[a] = "fire4"
						} else if checkPlatformEffectsBlock == "fire4" {
							platformsEffectsMAP[a] = "fire5"
						} else if checkPlatformEffectsBlock == "fire5" {
							platformsEffectsMAP[a] = ""
						}
						if checkPlatformEffectsBlock == "poison" {
							platformsEffectsMAP[a] = "poison2"
						} else if checkPlatformEffectsBlock == "poison2" {
							platformsEffectsMAP[a] = "poison3"
						} else if checkPlatformEffectsBlock == "poison3" {
							platformsEffectsMAP[a] = "poison4"
						} else if checkPlatformEffectsBlock == "poison4" {
							platformsEffectsMAP[a] = "poison5"
						} else if checkPlatformEffectsBlock == "poison5" {
							platformsEffectsMAP[a] = ""
						}
					}

					if checkPlatformEffectsBlock == "fire" || checkPlatformEffectsBlock == "fire2" || checkPlatformEffectsBlock == "fire3" || checkPlatformEffectsBlock == "fire4" || checkPlatformEffectsBlock == "fire5" || checkPlatformEffectsBlock == "fire6" {
						if enemiesMAP[a-(94*2)] == "enemy" && enemiesMAP[a-((94*2)+1)] == "enemyTL" {
							clearEnemyBlock = a - ((94 * 2) + 1)
							clearEnemy()
							killEnemy()

						} else if enemiesMAP[a-94] == "enemy" && enemiesMAP[a-95] == "enemyTL" {
							clearEnemyBlock = a - 95
							clearEnemy()
							killEnemy()

						} else if enemiesMAP[a-(94*2)] == "enemyTL" {
							clearEnemyBlock = a - (94 * 2)
							clearEnemy()
							killEnemy()

						} else if enemiesMAP[a-94] == "enemyTL" {
							clearEnemyBlock = a - 94
							clearEnemy()
							killEnemy()

						}
					}

					if checkPlatformEffectsBlock == "poison" || checkPlatformEffectsBlock == "poison2" || checkPlatformEffectsBlock == "poison3" || checkPlatformEffectsBlock == "poison4" || checkPlatformEffectsBlock == "poison5" {
						if enemiesMAP[a-(94*2)] == "enemy" && enemiesMAP[a-((94*2)+1)] == "enemyTL" {
							clearEnemyBlock = a - ((94 * 2) + 1)
							clearEnemy()
							killEnemy()

						} else if enemiesMAP[a-94] == "enemy" && enemiesMAP[a-95] == "enemyTL" {
							clearEnemyBlock = a - 95
							clearEnemy()
							killEnemy()

						} else if enemiesMAP[a-(94*2)] == "enemyTL" {
							clearEnemyBlock = a - (94 * 2)
							clearEnemy()
							killEnemy()

						} else if enemiesMAP[a-94] == "enemyTL" {
							clearEnemyBlock = a - 94
							clearEnemy()
							killEnemy()

						}
					}

					if activeSpecialItem == "fireball" {
						if checkLevelBlock == "floor" && checkActiveSpecialBlock != "" {
							platformsEffectsMAP[a] = "fire"
						}
					}
					if activeSpecialItem == "poisonball" {
						if checkLevelBlock == "floor" && checkActiveSpecialBlock != "" {
							platformsEffectsMAP[a] = "poison"
						}
					}

					if checkActiveSpecialBlock == "zap11" {
						if a < 4230 {
							activeSpecialMAP[a] = ""
							zapLR := flipcoin()
							if zapLR {
								activeSpecialMAP[a+95] = "zap11"
							} else {
								activeSpecialMAP[a+93] = "zap11"
							}
						} else {
							activeSpecialMAP[a] = ""
						}
					}
					if checkActiveSpecialBlock == "zap10" {
						if a < 4230 {
							activeSpecialMAP[a] = ""
							activeSpecialMAP[a+94] = "zap10"
						} else {
							activeSpecialMAP[a] = ""
						}
					}
					if checkActiveSpecialBlock == "zap9" {
						if a < 4230 {
							activeSpecialMAP[a] = ""
							activeSpecialMAP[a+94] = "zap9"
						} else {
							activeSpecialMAP[a] = ""
						}
					}

					if checkActiveSpecialBlock == "zap11" {
						if a < 4230 {
							activeSpecialMAP[a] = ""
							zapLR := flipcoin()
							if zapLR {
								activeSpecialMAP[a+95] = "zap11"
							} else {
								activeSpecialMAP[a+93] = "zap11"
							}
						} else {
							activeSpecialMAP[a] = ""
						}
					}
					if checkActiveSpecialBlock == "zap10" {
						if a < 4230 {
							activeSpecialMAP[a] = ""
							activeSpecialMAP[a+94] = "zap10"
						} else {
							activeSpecialMAP[a] = ""
						}
					}
					if checkActiveSpecialBlock == "zap9" {
						if a < 4230 {
							activeSpecialMAP[a] = ""
							activeSpecialMAP[a+94] = "zap9"
						} else {
							activeSpecialMAP[a] = ""
						}
					}

					if checkActiveSpecialBlock == "zap12" && checkEnemyBlock == "enemy" && enemiesMAP[a-95] == "enemyTL" {
						clearEnemyBlock = a - 95
						clearEnemy()
						killEnemy()
					}
					if checkActiveSpecialBlock == "zap12" && checkEnemyBlock == "enemy" && enemiesMAP[a-94] == "enemyTL" {
						clearEnemyBlock = a - 94
						clearEnemy()
						killEnemy()
					}
					if checkActiveSpecialBlock == "fireball" && checkEnemyBlock == "enemy" && enemiesMAP[a-1] == "enemyTL" {
						clearEnemyBlock = a - 1
						clearEnemy()
						killEnemy()
					}
					if checkActiveSpecialBlock == "fireballTL" && checkEnemyBlock == "enemy" && enemiesMAP[a-1] == "enemyTL" {
						clearEnemyBlock = a - 1
						clearEnemy()
						killEnemy()
					}
					if checkActiveSpecialBlock == "fireballTL" && checkEnemyBlock == "enemy" && enemiesMAP[a-95] == "enemyTL" {
						clearEnemyBlock = a - 95
						clearEnemy()
						killEnemy()
					}
					if checkActiveSpecialBlock == "zap11" && checkEnemyBlock == "enemyTL" {
						clearEnemyBlock = a
						clearEnemy()
						killEnemy()
					}
					if checkActiveSpecialBlock == "zap11" && checkEnemyBlock == "enemy" && enemiesMAP[a-1] == "enemyTL" {
						clearEnemyBlock = a - 1
						clearEnemy()
						killEnemy()
					}
					if checkActiveSpecialBlock == "zap10" && checkEnemyBlock == "enemyTL" {
						clearEnemyBlock = a
						clearEnemy()
						killEnemy()
					}
					if checkActiveSpecialBlock == "zap10" && checkEnemyBlock == "enemy" && enemiesMAP[a-1] == "enemyTL" {
						clearEnemyBlock = a - 1
						clearEnemy()
						killEnemy()
					}
					if checkActiveSpecialBlock == "zap9" && checkEnemyBlock == "enemyTL" {
						clearEnemyBlock = a
						clearEnemy()
						killEnemy()
					}
					if checkActiveSpecialBlock == "zap9" && checkEnemyBlock == "enemy" && enemiesMAP[a-1] == "enemyTL" {
						clearEnemyBlock = a - 1
						clearEnemy()
						killEnemy()
					}

					// update mini boss effects
					if checkMiniBossEffectsBlock == "zap" {
						miniBossEffectsMAP[a] = "zap2"
					} else if checkMiniBossEffectsBlock == "zap2" {
						miniBossEffectsMAP[a] = "zap3"
					} else if checkMiniBossEffectsBlock == "zap3" {
						miniBossEffectsMAP[a] = "zap4"
					} else if checkMiniBossEffectsBlock == "zap4" {
						miniBossEffectsMAP[a] = ""
					} else if checkMiniBossEffectsBlock == "zap5" {
						if a < 4136 {
							miniBossEffectsMAP[a] = ""
							miniBossEffectsMAP[a+95] = "zap5"
						} else {
							miniBossEffectsMAP[a] = ""
						}
					} else if checkMiniBossEffectsBlock == "zap6" {
						if a < 4136 {
							miniBossEffectsMAP[a] = ""
							miniBossEffectsMAP[a+93] = "zap6"
						} else {
							miniBossEffectsMAP[a] = ""
						}
					} else if checkMiniBossEffectsBlock == "zap8" {
						if a < 4136 {
							chooseLR := flipcoin()
							if chooseLR {
								miniBossEffectsMAP[a] = ""
								miniBossEffectsMAP[a+93] = "zap8"
							} else {
								miniBossEffectsMAP[a] = ""
								miniBossEffectsMAP[a+95] = "zap8"
							}
						} else {
							miniBossEffectsMAP[a] = ""
						}
					}

					// MARK: move falling blocks down
					if checkFallingBlock == "fallingBlock" && platformsMAP[a] != "" {
						if levelMAP[a+(94*2)] != "floor" && levelMAP[a+((94*2)+1)] != "floor" {

							fallingBlockTypeHOLDER = platformsMAP[a]
							currentFallingBlock = a
							clearFallingBlock()
							currentFallingBlock += 94
							moveFallingBlockDown()

							checkEnemyBelowBlock := enemiesMAP[a+(94*2)]
							checkEnemyBelowBlock2 := enemiesMAP[(a+1)+(94*2)]

							if checkEnemyBelowBlock != "" || checkEnemyBelowBlock2 != "" {

								if checkEnemyBelowBlock == "enemyTL" {
									clearEnemyBlock = a + (94 * 2)
									clearEnemy()
									killEnemy()
								}

								if checkEnemyBelowBlock2 == "enemyTL" {
									clearEnemyBlock = (a + 1) + (94 * 2)
									clearEnemy()
									killEnemy()
								}
								if checkEnemyBelowBlock == "enemy" {
									clearEnemyBlock = (a - 1) + (94 * 2)
									clearEnemy()
									killEnemy()
								}
								if checkEnemyBelowBlock2 == "enemy" {
									clearEnemyBlock = a + (94 * 2)
									clearEnemy()
									killEnemy()
								}

							}

						} else {
							levelMAP[a] = "floor"
							levelMAP[a+1] = "floor"
							levelMAP[a+94] = "floor"
							levelMAP[a+95] = "floor"
						}

					}

					// move mini boss right
					if checkMiniBossBlock == "miniBossTL" {
						if miniBossJumpOn == false && miniBossFallOn == false && miniBossWallJumpOn == false {
							if miniBossRightLeft {
								miniBossCurrentBlock = a
								moveMiniBossRight()
								if miniBossRollOn == true {
									miniBossCurrentBlock = a + 1
									moveMiniBossRight()
								}
							}

						}
						if miniBossJumpOn == false && miniBossWallJumpOn == false && miniBossType != 6 {
							if levelMAP[a+94*4] != "floor" {
								if a < 3760 {
									miniBossCurrentBlock = a
									moveMiniBossDown()
								}
							}
						}
					}

					// move rain down
					if checkRainBlock == "rain" {

						if a < 4136 {
							rainMAP[a] = ""
							rainMAP[a+94] = "rain"
						} else {
							rainMAP[a] = ""

						}

					}

					// move drop weapon down
					if checkDropWeaponBlock != "" && weaponsMAP[a+1] != "" {
						if levelMAP[a+(94*2)] != "floor" && levelMAP[(a+1)+(94*2)] != "floor" {
							dropWeaponHOLDER1 = checkDropWeaponBlock
							dropWeaponHOLDER2 = weaponsMAP[a+1]
							currentDropWeaponBlock = a
							clearWeapon()
							moveWeaponDown()
						}

					}

					// moving platforms down
					if checkPlatformsBlock == "movingPlatformDown" {

						movingPlatformHorizontal := a / 94

						if movingPlatformHorizontal < 43 {

							levelMAP[a] = ""
							levelMAP[a+94] = "floor"
							platformsMAP[a] = ""
							platformsMAP[a+94] = "movingPlatformDown"
							backgroundObjectsMAP[a] = ""
							backgroundObjectsMAP[a+94] = ""

						} else if movingPlatformHorizontal >= 43 {

							platformsMAP[a] = "movingPlatformUp"

						}

					}

					if checkPlatformsBlock == "movingPlatformDownTL" {

						movingPlatformHorizontal := a / 94

						if movingPlatformHorizontal < 43 {

							levelMAP[a] = ""
							levelMAP[a+94] = "floor"
							platformsMAP[a] = ""
							platformsMAP[a+94] = "movingPlatformDownTL"
							backgroundObjectsMAP[a] = ""
							backgroundObjectsMAP[a+94] = ""

						} else if movingPlatformHorizontal >= 43 {
							platformsMAP[a] = "movingPlatformUpTL"
						}

					}

					// move coin down
					if checkCoinBlock == "coinTL" {

						if levelMAP[a+(94*2)] != "floor" && levelMAP[(a+1)+(94*2)] != "floor" {
							coinCurrentBlock = a
							clearCoin()
							moveCoinDown()
						}

					}

					// MARK: collect power up

					if checkPowerUpBlock == "powerUp1TL" || checkPowerUpBlock == "powerUp2TL" || checkPowerUpBlock == "powerUp3TL" || checkPowerUpBlock == "powerUp4TL" || checkPowerUpBlock == "powerUp5TL" || checkPowerUpBlock == "powerUp6TL" || checkPowerUpBlock == "powerUp7TL" || checkPowerUpBlock == "powerUp8TL" || checkPowerUpBlock == "powerUp9TL" || checkPowerUpBlock == "powerUp10TL" {

						if levelMAP[a+(94*2)] != "floor" && levelMAP[(a+1)+(94*2)] != "floor" {

							if a > 5076 {
								clearPowerUp()
							} else {

								currentPowerUpBlock = a
								clearPowerUp()
								movePowerUpDown()

							}

						}

					}

					if checkPowerUpVehicleBlock == "powerUpVehicleTL" {
						powerUpVehicleCurrentBlock = a
						clearPowerUpVehicle()
						movePowerUpVehicleRight()
					}

					if checkEnemyBlock == "enemyTL" {

						if enemyTypeGenerate == 8 {

							currentEnemyMovement := enemiesMovementMAP[a]

							if secondsCountdown > 58 {

								if checkBelowEnemy(a) {

									chooseLR := flipcoin()

									if chooseLR {
										clearEnemyBlock = a
										movementHOLDER := enemiesMovementMAP[a]
										clearEnemy()
										enemiesMovementMAP[a+1] = movementHOLDER
										moveEnemyRight()
									} else {
										clearEnemyBlock = a
										movementHOLDER := enemiesMovementMAP[a]
										clearEnemy()
										enemiesMovementMAP[a-1] = movementHOLDER
										moveEnemyLeft()
									}

								} else {
									clearEnemyBlock = a
									movementHOLDER := enemiesMovementMAP[a]
									clearEnemy()
									enemiesMovementMAP[a+94] = movementHOLDER
									moveEnemyDown()
								}
							}
							if secondsCountdown > 50 && secondsCountdown < 53 {

								if checkBelowEnemy(a) {

									chooseLR := flipcoin()

									if chooseLR {
										clearEnemyBlock = a
										movementHOLDER := enemiesMovementMAP[a]
										clearEnemy()
										enemiesMovementMAP[a+1] = movementHOLDER
										moveEnemyRight()
									} else {
										clearEnemyBlock = a
										movementHOLDER := enemiesMovementMAP[a]
										clearEnemy()
										enemiesMovementMAP[a-1] = movementHOLDER
										moveEnemyLeft()
									}

								} else {
									clearEnemyBlock = a
									movementHOLDER := enemiesMovementMAP[a]
									clearEnemy()
									enemiesMovementMAP[a+94] = movementHOLDER
									moveEnemyDown()
								}
							}

							if secondsCountdown > 43 && secondsCountdown < 47 {

								if checkBelowEnemy(a) {

									chooseLR := flipcoin()

									if chooseLR {
										clearEnemyBlock = a
										movementHOLDER := enemiesMovementMAP[a]
										clearEnemy()
										enemiesMovementMAP[a+1] = movementHOLDER
										moveEnemyRight()
									} else {
										clearEnemyBlock = a
										movementHOLDER := enemiesMovementMAP[a]
										clearEnemy()
										enemiesMovementMAP[a-1] = movementHOLDER
										moveEnemyLeft()
									}

								} else {
									clearEnemyBlock = a
									movementHOLDER := enemiesMovementMAP[a]
									clearEnemy()
									enemiesMovementMAP[a+94] = movementHOLDER
									moveEnemyDown()
								}
							}

							if secondsCountdown > 35 && secondsCountdown < 38 {

								if checkBelowEnemy(a) {

									chooseLR := flipcoin()

									if chooseLR {
										clearEnemyBlock = a
										movementHOLDER := enemiesMovementMAP[a]
										clearEnemy()
										enemiesMovementMAP[a+1] = movementHOLDER
										moveEnemyRight()
									} else {
										clearEnemyBlock = a
										movementHOLDER := enemiesMovementMAP[a]
										clearEnemy()
										enemiesMovementMAP[a-1] = movementHOLDER
										moveEnemyLeft()
									}

								} else {
									clearEnemyBlock = a
									movementHOLDER := enemiesMovementMAP[a]
									clearEnemy()
									enemiesMovementMAP[a+94] = movementHOLDER
									moveEnemyDown()
								}
							}

							enemyMovementTEST = enemiesMovementMAP[a]

							enemyHorizontal := a / 94

							if enemyHorizontal < 8 {
								enemiesMovementMAP[a] = "ghostDown"
							}

							if currentEnemyMovement == "ghostDown" && levelMAP[a+(94*2)] == "floor" {

								movementHOLDER := enemiesMovementMAP[a]
								clearEnemy()
								enemiesMovementMAP[a+1] = movementHOLDER
								moveEnemyRight()

							} else if currentEnemyMovement == "ghostDown" && levelMAP[a+(94*2)] != "floor" {

								clearEnemyBlock = a
								movementHOLDER := enemiesMovementMAP[a]
								clearEnemy()
								enemiesMovementMAP[a+94] = movementHOLDER
								moveEnemyDown()
							}

							if enemyHorizontal < 10 {
								if currentEnemyMovement == "enemy4Special2" {
									enemiesMovementMAP[a] = "ghostDown"
								}
							}

						} else {

							if levelMAP[a+(94*2)] == "earthquake" && levelMAP[(a+1)+(94*2)] == "earthquake" {
								clearEnemyBlock = a
								movementHOLDER := enemiesMovementMAP[a]
								clearEnemy()
								enemiesMovementMAP[a+94] = movementHOLDER
								moveEnemyDown()

								if a > 4418 {
									clearEnemyBlock = a
									clearEnemy()
									enemiesScreenCount--
								}

							} else if a > 4134 { // boundary correction
								clearEnemyBlock = a
								movementHOLDER := enemiesMovementMAP[a]
								clearEnemy()
								enemiesMovementMAP[a-94] = movementHOLDER
								moveEnemyUp()
							} else if levelMAP[a+(94*2)] != "floor" && levelMAP[(a+1)+(94*2)] != "floor" {

								// MARK:  random enemy movements

								currentEnemyMovement := enemiesMovementMAP[a]

								if currentEnemyMovement != "jump4" && enemiesMovementMAP[a] != "jump3" && enemiesMovementMAP[a] != "jump2" && enemiesMovementMAP[a] != "jump1" && enemiesMovementMAP[a] != "jump5" && enemiesMovementMAP[a] != "jump6" && enemiesMovementMAP[a] != "roofR" && enemiesMovementMAP[a] != "roofL" && enemiesMovementMAP[a] != "rightUp" && enemiesMovementMAP[a] != "leftUp" {

									if checkBelowEnemy(a) {

										enemyHorizontal := a / 94

										if enemyHorizontal > 4 {
											clearEnemyBlock = a
											movementHOLDER := enemiesMovementMAP[a]
											clearEnemy()
											enemiesMovementMAP[a-94] = movementHOLDER
											moveEnemyUp()
										}

									} else {
										// MARK: enemy gravity
										if enemyTypeGenerate != 6 {
											clearEnemyBlock = a
											movementHOLDER := enemiesMovementMAP[a]
											clearEnemy()
											enemiesMovementMAP[a+94] = movementHOLDER
											moveEnemyDown()
										} else if enemyTypeGenerate == 6 {
											if levelMAP[a+(94*6)] != "floor" {
												clearEnemyBlock = a
												movementHOLDER := enemiesMovementMAP[a]
												clearEnemy()
												enemiesMovementMAP[a+94] = movementHOLDER
												moveEnemyDown()
											}

										}
									}

								} else if currentEnemyMovement == "jump4" {

									if checkAboveEnemy(a) {
										rightLeft := flipcoin()
										if rightLeft {
											enemiesMovementMAP[a] = "left"
										} else {
											enemiesMovementMAP[a] = "right"
										}
									} else {
										enemyHorizontal := a / 94
										if enemyHorizontal > 6 {

											clearEnemyBlock = a
											clearEnemy()
											enemiesMovementMAP[a-94] = "jump3"
											moveEnemyUp()
										} else {
											rightLeft := flipcoin()
											if rightLeft {
												enemiesMovementMAP[a] = "left"
											} else {
												enemiesMovementMAP[a] = "right"
											}
										}
									}
								} else if currentEnemyMovement == "jump3" {
									if checkAboveEnemy(a) {
										rightLeft := flipcoin()
										if rightLeft {
											enemiesMovementMAP[a] = "left"
										} else {
											enemiesMovementMAP[a] = "right"
										}
									} else {
										enemyHorizontal := a / 94
										if enemyHorizontal > 6 {

											clearEnemyBlock = a
											clearEnemy()
											enemiesMovementMAP[a-94] = "jump2"
											moveEnemyUp()
										} else {
											rightLeft := flipcoin()
											if rightLeft {
												enemiesMovementMAP[a] = "left"
											} else {
												enemiesMovementMAP[a] = "right"
											}
										}
									}
								} else if currentEnemyMovement == "jump2" {
									if checkAboveEnemy(a) {
										rightLeft := flipcoin()
										if rightLeft {
											enemiesMovementMAP[a] = "left"
										} else {
											enemiesMovementMAP[a] = "right"
										}
									} else {
										enemyHorizontal := a / 94
										if enemyHorizontal > 6 {

											clearEnemyBlock = a
											clearEnemy()
											enemiesMovementMAP[a-94] = "jump1"
											moveEnemyUp()
										} else {
											rightLeft := flipcoin()
											if rightLeft {
												enemiesMovementMAP[a] = "left"
											} else {
												enemiesMovementMAP[a] = "right"
											}
										}
									}
								} else if currentEnemyMovement == "jump1" {

									rightLeft := flipcoin()
									if rightLeft {
										enemiesMovementMAP[a] = "left"
									} else {
										enemiesMovementMAP[a] = "right"
									}
								}

							} else if levelMAP[a+(94*2)] == "floor" || levelMAP[(a+1)+(94*2)] == "floor" {
								// enemy movement right

								currentEnemyMovement := enemiesMovementMAP[a]

								enemyHorizontal := a / 94
								enemyVertical := a - (enemyHorizontal * 94)

								if enemyTypeGenerate == 7 && enemyVertical > 85 && enemiesMovementMAP[a] != "roofR" && enemiesMovementMAP[a] != "roofL" {

									enemiesMovementMAP[a] = "jump6"

								} else if enemyTypeGenerate == 8 && enemiesMovementMAP[a] == "left" {

									enemyHorizontal := a / 94

									if enemyHorizontal > 30 {
										enemiesMovementMAP[a] = "leftUp"
									}

								} else {

									switch currentEnemyMovement {

									case "right":

										if checkRightEnemy(a) {
											enemiesMovementMAP[a] = "left"
											enemiesMovementMAP[a+4] = "right"
										} else {
											clearEnemyBlock = a
											movementHOLDER := enemiesMovementMAP[a]
											clearEnemy()
											enemiesMovementMAP[a+1] = movementHOLDER
											moveEnemyRight()
										}
									case "jump4":
										clearEnemyBlock = a
										clearEnemy()
										enemiesMovementMAP[a-94] = "jump3"
										moveEnemyUp()
									}
								}
							}

							if enemyTypeGenerate == 6 {
								currentEnemyMovement := enemiesMovementMAP[a]

								switch currentEnemyMovement {

								case "right":

									if checkRightEnemy(a) {
										enemiesMovementMAP[a] = "left"
										enemiesMovementMAP[a+4] = "right"
									} else {
										clearEnemyBlock = a
										movementHOLDER := enemiesMovementMAP[a]
										clearEnemy()
										enemiesMovementMAP[a+1] = movementHOLDER
										moveEnemyRight()
									}

								}

							}

							if enemyTypeGenerate == 7 {
								currentEnemyMovement := enemiesMovementMAP[a]

								if currentEnemyMovement == "roofR" {

									enemyHorizontal := a / 94
									enemyVertical := a - (enemyHorizontal * 94)

									chooseFallBlock := rolldice()
									fallBlock := 40

									switch chooseFallBlock {
									case 1:
										fallBlock = enemy7RoofFall1
									case 2:
										fallBlock = enemy7RoofFall2
									case 3:
										fallBlock = enemy7RoofFall3
									case 4:
										fallBlock = enemy7RoofFall4
									case 5:
										fallBlock = enemy7RoofFall5
									case 6:
										fallBlock = enemy7RoofFall6
									}

									if enemyVertical < fallBlock {
										clearEnemyBlock = a
										clearEnemy()
										enemiesMovementMAP[a+1] = "roofR"
										moveEnemyRight()
									} else {
										enemiesMovementMAP[a] = "right"
									}
								}

							}

						}
					} // end checkEnemy enemyTL
				} // end  right down movemment for a := 5263; a >= 0; a--

				// left up movements

				for a := 0; a < 5264; a++ {
					checkEnemyBlock := enemiesMAP[a]
					checkDeadEnemiesBlock := deadEnemiesMAP[a]
					checkMiniBossBlock := miniBossMAP[a]
					checkMiniBossEffectsBlock := miniBossEffectsMAP[a]
					checkActiveSpecialBlock := activeSpecialMAP[a]
					checkRainFrogBlock := rainFrogMAP[a]
					checkPlatformsBlock := platformsMAP[a]

					// update horizontal platforms
					if horizplatlr == true {
						if checkPlatformsBlock == "movingplatH1" {
							horizplatH = a / 94
							horizplatV = a - (horizplatH * 94)

							if horizplatV < 6 {
								horizplatlr = false
							} else {
								platformsMAP[a] = ""
								platformsMAP[a+1] = ""
								platformsMAP[a+2] = ""
								platformsMAP[a+3] = ""
								platformsMAP[a+4] = ""
								platformsMAP[a+5] = ""
								platformsMAP[a+6] = ""
								platformsMAP[a+7] = ""

								levelMAP[a] = ""
								levelMAP[a+1] = ""
								levelMAP[a+2] = ""
								levelMAP[a+3] = ""
								levelMAP[a+4] = ""
								levelMAP[a+5] = ""
								levelMAP[a+6] = ""
								levelMAP[a+7] = ""

								platformsMAP[a-1] = "movingplatH1"
								platformsMAP[a] = "movingplatH2"
								platformsMAP[a+1] = "movingplatH3"
								platformsMAP[a+2] = "movingplatH4"
								platformsMAP[a+3] = "movingplatH5"
								platformsMAP[a+4] = "movingplatH6"
								platformsMAP[a+5] = "movingplatH7"
								platformsMAP[a+6] = "movingplatH8"

								levelMAP[a-1] = "floor"
								levelMAP[a] = "floor"
								levelMAP[a+1] = "floor"
								levelMAP[a+2] = "floor"
								levelMAP[a+3] = "floor"
								levelMAP[a+4] = "floor"
								levelMAP[a+5] = "floor"
								levelMAP[a+6] = "floor"

							}
						}
					}

					// update frog rain

					if checkRainFrogBlock == "rainFrog" {
						rainFrogHorizontal := a / 94
						rainFrogVertical := a - (94 * rainFrogHorizontal)

						if levelMAP[a+94*3] == "floor" && rainFrogVertical > 2 {
							rainFrogMAP[a] = ""
							rainFrogMAP[a-1] = "rainFrog"
						} else if rainFrogVertical <= 2 {
							rainFrogMAP[a] = ""
							if rainFrogTimerOn == false {
								rainFrogOn = true
								rainFrogTimerOn = true
							}
						}
					}

					// update active special
					if checkActiveSpecialBlock == "zap12" {
						if a > 282 {
							activeSpecialMAP[a] = ""
							zapLR := flipcoin()
							if zapLR {
								activeSpecialMAP[a-95] = "zap12"
							} else {
								activeSpecialMAP[a-93] = "zap12"
							}
						} else {
							activeSpecialMAP[a] = ""
						}
					}

					// update mini boss effects

					if checkMiniBossEffectsBlock == "zap7" {

						if a > 470 {
							miniBossEffectsMAP[a] = ""
							miniBossEffectsMAP[a-94] = "zap7"
						} else {
							miniBossEffectsMAP[a] = "zap8"
						}
					}

					if checkMiniBossBlock == "miniBossTL" {
						currentMiniBossBlock = a
					}

					// move mini boss left
					if miniBossJumpOn == false && miniBossFallOn == false && miniBossWallJumpOn == false {
						if miniBossRightLeft == false {
							if checkMiniBossBlock == "miniBossTL" {
								miniBossCurrentBlock = a
								moveMiniBossLeft()
								if miniBossRollOn {
									miniBossCurrentBlock = a - 1
									moveMiniBossLeft()
								}
							}
						}
					} else {
						if checkMiniBossBlock == "miniBossTL" {

							if miniBossCurrentBlock > 1504 {
								miniBossCurrentBlock = a
								moveMiniBossUp()
							}
						}

					}

					if checkDeadEnemiesBlock == "deadEnemy1_1" {
						deadEnemiesMAP[a] = ""
					} else if checkDeadEnemiesBlock == "deadEnemy1_2" {
						deadEnemiesMAP[a] = "deadEnemy1_1"
					} else if checkDeadEnemiesBlock == "deadEnemy1_3" {
						deadEnemiesMAP[a] = "deadEnemy1_2"
					} else if checkDeadEnemiesBlock == "deadEnemy1_4" {
						deadEnemiesMAP[a] = "deadEnemy1_3"
					}

					if checkEnemyBlock == "enemyTL" {
						if levelMAP[a+(94*2)] == "earthquake" && levelMAP[(a+1)+(94*2)] == "earthquake" {
							clearEnemyBlock = a
							movementHOLDER := enemiesMovementMAP[a]
							clearEnemy()
							enemiesMovementMAP[a+94] = movementHOLDER
							moveEnemyDown()
							if a > 4418 {
								clearEnemyBlock = a
								clearEnemy()
								enemiesScreenCount--
							}

						} else if a > 4134 { // boundary correction
							clearEnemyBlock = a
							movementHOLDER := enemiesMovementMAP[a]
							clearEnemy()
							enemiesMovementMAP[a-94] = movementHOLDER
							moveEnemyUp()
						} else if levelMAP[a+(94*2)] == "floor" || levelMAP[(a+1)+(94*2)] == "floor" {

							currentEnemyMovement := enemiesMovementMAP[a]
							// random enemy special move
							if currentEnemyMovement != "jump4" && enemiesMovementMAP[a] != "jump3" && enemiesMovementMAP[a] != "jump2" && enemiesMovementMAP[a] != "jump1" {

								if enemyTypeGenerate == 1 {
									chooseEnemySpecialMove := rInt(0, 1000)
									if chooseEnemySpecialMove > 950 {
										enemiesMovementMAP[a] = "jump4"
									}
								} else if enemyTypeGenerate == 2 {
									chooseEnemySpecialMove := rInt(0, 1000)
									if chooseEnemySpecialMove > 950 {
										enemiesMovementMAP[a] = "left2"
									}
								} else if enemyTypeGenerate == 3 {
									chooseEnemySpecialMove := rInt(0, 1000)
									if chooseEnemySpecialMove > 950 {
										enemyCurrentMovement := enemiesMovementMAP[a]
										if enemyCurrentMovement == "left" {
											bulletsMAP[a-2] = "bulletEnemyL"
										} else if enemyCurrentMovement == "right" {
											bulletsMAP[a+2] = "bulletEnemyR"
										}
									}
								} else if enemyTypeGenerate == 4 {
									chooseEnemySpecialMove := rInt(0, 1000)
									if chooseEnemySpecialMove > 950 {
										enemiesMovementMAP[a] = "enemy4Special"
									}
								} else if enemyTypeGenerate == 5 {
									chooseEnemySpecialMove := rInt(0, 1000)
									if chooseEnemySpecialMove > 950 {
										enemyCurrentMovement := enemiesMovementMAP[a]
										if enemyCurrentMovement == "left" {
											bulletsMAP[a-2] = "bulletEnemyL"
										} else if enemyCurrentMovement == "right" {
											bulletsMAP[a+2] = "bulletEnemyR"
										}
									} else if chooseEnemySpecialMove < 950 && chooseEnemySpecialMove > 900 {
										enemiesMovementMAP[a] = "jump4"

									}
								}

							}

							// enemy movement left

							enemyHorizontal := a / 94
							enemyVertical := a - (enemyHorizontal * 94)

							if enemyTypeGenerate == 7 && enemyVertical < 5 && enemiesMovementMAP[a] != "roofR" && enemiesMovementMAP[a] != "roofL" {
								enemiesMovementMAP[a] = "jump5"
							} else {

								switch currentEnemyMovement {

								case "enemy4Special":

									clearEnemyBlock = a
									clearEnemy()

									newEnemyBlock := moveEnemyTeleport()

									enemiesMovementMAP[newEnemyBlock] = "enemy4Special2"

								case "enemy4Special2":

									rightLeft := flipcoin()
									if rightLeft {
										enemiesMovementMAP[a] = "left"
									} else {
										enemiesMovementMAP[a] = "right"
									}

								case "left":

									if checkLeftEnemy(a) {
										enemiesMovementMAP[a-4] = "left"
										enemiesMovementMAP[a] = "right"
									} else {
										clearEnemyBlock = a
										movementHOLDER := enemiesMovementMAP[a]
										clearEnemy()
										enemiesMovementMAP[a-1] = movementHOLDER
										moveEnemyLeft()
									}
								case "left2":
									if checkLeftEnemy(a) {
										enemiesMovementMAP[a-4] = "left"
										enemiesMovementMAP[a] = "right"
									} else if checkLeftEnemy2(a) {
										enemiesMovementMAP[a-4] = "left"
										enemiesMovementMAP[a] = "right"
									} else {
										clearEnemyBlock = a
										movementHOLDER := enemiesMovementMAP[a]
										clearEnemy()
										enemiesMovementMAP[a-1] = movementHOLDER
										moveEnemyLeft()
										clearEnemyBlock = a - 1
										movementHOLDER = enemiesMovementMAP[a-1]
										clearEnemy()
										enemiesMovementMAP[a-2] = movementHOLDER
										moveEnemyLeft()
									}

								}
							}
						}

						// spike man movement up
						if enemyTypeGenerate == 7 {
							currentEnemyMovement := enemiesMovementMAP[a]

							if currentEnemyMovement == "jump5" {

								enemyHorizontal := a / 94

								if enemyHorizontal > 2 {

									clearEnemyBlock = a
									clearEnemy()
									enemiesMovementMAP[a-94] = "jump5"
									moveEnemyUp()

								} else {
									enemiesMovementMAP[a] = "roofR"
								}
							} else if currentEnemyMovement == "jump6" {

								enemyHorizontal := a / 94

								if enemyHorizontal > 2 {

									clearEnemyBlock = a
									clearEnemy()
									enemiesMovementMAP[a-94] = "jump6"
									moveEnemyUp()

								} else {
									enemiesMovementMAP[a] = "roofL"
								}
							} else if currentEnemyMovement == "roofL" {
								enemyHorizontal := a / 94
								enemyVertical := a - (enemyHorizontal * 94)

								chooseFallBlock := rolldice()
								fallBlock := 40

								switch chooseFallBlock {
								case 1:
									fallBlock = enemy7RoofFall1
								case 2:
									fallBlock = enemy7RoofFall2
								case 3:
									fallBlock = enemy7RoofFall3
								case 4:
									fallBlock = enemy7RoofFall4
								case 5:
									fallBlock = enemy7RoofFall5
								case 6:
									fallBlock = enemy7RoofFall6
								}
								if enemyVertical > fallBlock {
									clearEnemyBlock = a
									clearEnemy()
									enemiesMovementMAP[a-1] = "roofL"
									moveEnemyLeft()
								} else {
									enemiesMovementMAP[a] = "left"
								}

							}
						}

						if enemyTypeGenerate == 6 {

							enemyHorizontal := a / 94

							if enemyHorizontal > 10 {
								chooseEnemySpecialMove := rInt(0, 1000)
								if chooseEnemySpecialMove > 950 {
									enemiesMovementMAP[a] = "jump4"
								} else if chooseEnemySpecialMove < 950 && chooseEnemySpecialMove > 900 {
									bulletsMAP[a+(94*3)] = "bulletEnemyD"
								}
							}

							currentEnemyMovement := enemiesMovementMAP[a]

							switch currentEnemyMovement {

							case "left":

								if checkLeftEnemy(a) {
									enemiesMovementMAP[a-4] = "left"
									enemiesMovementMAP[a] = "right"
								} else {
									clearEnemyBlock = a
									movementHOLDER := enemiesMovementMAP[a]
									clearEnemy()
									enemiesMovementMAP[a-1] = movementHOLDER
									moveEnemyLeft()
								}
							case "jump4":
								if checkAboveEnemy(a) {
									rightLeft := flipcoin()
									if rightLeft {
										enemiesMovementMAP[a] = "left"
									} else {
										enemiesMovementMAP[a] = "right"
									}
								} else {
									clearEnemyBlock = a
									clearEnemy()
									enemiesMovementMAP[a-94] = "jump3"
									moveEnemyUp()
								}
							case "jump3":
								if checkAboveEnemy(a) {
									rightLeft := flipcoin()
									if rightLeft {
										enemiesMovementMAP[a] = "left"
									} else {
										enemiesMovementMAP[a] = "right"
									}
								} else {
									clearEnemyBlock = a
									clearEnemy()
									enemiesMovementMAP[a-94] = "jump2"
									moveEnemyUp()
								}
							case "jump2":
								if checkAboveEnemy(a) {
									rightLeft := flipcoin()
									if rightLeft {
										enemiesMovementMAP[a] = "left"
									} else {
										enemiesMovementMAP[a] = "right"
									}
								} else {
									clearEnemyBlock = a
									clearEnemy()
									enemiesMovementMAP[a-94] = "jump1"
									moveEnemyUp()
								}
							case "jump1":

								rightLeft := flipcoin()
								if rightLeft {
									enemiesMovementMAP[a] = "left"
								} else {
									enemiesMovementMAP[a] = "right"
								}

							}

						}
					}

				}
			}

			// MARK: update miniboss

			if miniBossOn1 == false {
				if killzCount == 10 {
					cMINIBOSS()
					miniBossOn1 = true
				}
			}
			if miniBossOn2 == false {
				if killzCount == 20 {
					cMINIBOSS()
					miniBossOn2 = true
				}
			}
			if miniBossOn3 == false {
				if killzCount == 30 {
					cMINIBOSS()
					miniBossOn3 = true
				}
			}

			if miniBossOn1 || miniBossOn2 || miniBossOn3 {

				bossTextX -= 2
				rl.DrawText("the boss is here...", bossTextX-3, 203, 60, rl.Black)
				rl.DrawText("the boss is here...", bossTextX-1, 201, 60, rl.White)
				rl.DrawText("the boss is here...", bossTextX, 200, 60, rl.Magenta)

				if bossTextX < -550 {
					bossTextX = screenW + 10
				}

				if miniBossType == 5 {
					if frameCount%120 == 0 {
						if boss5RollOn {
							boss5RollOn = false
						} else {
							boss5RollOn = true
						}
					}
				}

				if miniBossType == 3 {
					if secondsCountdown < 58 {
						if frameCount%60 == 0 {
							if miniBossJumpTimerOn {
								miniBossJumpTimerOn = false
							} else {
								miniBossJumpTimerOn = true
							}
						}
					}

					if miniBossWallJumpOn {
						if frameCount%120 == 0 {
							for a := 0; a < 5264; a++ {
								checkMiniBossBlock := miniBossMAP[a]
								if checkMiniBossBlock == "miniBossTL" {
									miniBossCurrentBlock = a
								}
							}

							zapBlock := miniBossCurrentBlock + 94*12

							miniBossEffectsMAP[zapBlock+4] = "zap5"
							miniBossEffectsMAP[zapBlock+1] = "zap6"
							miniBossEffectsMAP[zapBlock+192] = "zap5"
							miniBossEffectsMAP[zapBlock-189] = "zap6"
							miniBossEffectsMAP[zapBlock+282] = "zap5"
							miniBossEffectsMAP[zapBlock+470] = "zap6"
							miniBossWallJumpOn = false
						}

					}

				}

				if miniBossRollOn == true {
					if frameCount%60 == 0 {
						miniBossRollOn = false
					}
				}

				if frameCount%30 == 0 {

					if miniBossOn1 || miniBossOn2 || miniBossOn3 {

						switch miniBossType {
						case 1:
							if miniBossJumpOn == false {
								if miniBossRightLeft {
									zapBlock := currentMiniBossBlock + 94*2
									zapBlock -= 6
									for a := 0; a < 20; a++ {
										miniBossEffectsMAP[zapBlock+a] = "zap"
									}
									zapBlock += 11
									for a := 0; a < 6; a++ {
										newZapBlock := zapBlock - (a * 95)
										miniBossEffectsMAP[newZapBlock] = "zap"
										newZapBlock = zapBlock - (a * 93)
										miniBossEffectsMAP[newZapBlock] = "zap"
										newZapBlock = zapBlock + (a * 95)
										miniBossEffectsMAP[newZapBlock] = "zap"
										newZapBlock = zapBlock + (a * 93)
										miniBossEffectsMAP[newZapBlock] = "zap"
									}
								} else {
									zapBlock := currentMiniBossBlock + 94*2
									zapBlock -= 10
									for a := 0; a < 20; a++ {
										miniBossEffectsMAP[zapBlock+a] = "zap"
									}
									zapBlock += 11
									for a := 0; a < 6; a++ {
										newZapBlock := zapBlock - (a * 95)
										miniBossEffectsMAP[newZapBlock] = "zap"
										newZapBlock = zapBlock - (a * 93)
										miniBossEffectsMAP[newZapBlock] = "zap"
										newZapBlock = zapBlock + (a * 95)
										miniBossEffectsMAP[newZapBlock] = "zap"
										newZapBlock = zapBlock + (a * 93)
										miniBossEffectsMAP[newZapBlock] = "zap"
									}
								}
							}

						case 4:
							chooseZap := flipcoin()

							if chooseZap {
								if miniBossRightLeft {
									zapBlock := currentMiniBossBlock - 94*4
									zapBlock += 5
									miniBossEffectsMAP[zapBlock] = "zap7"
									zapBlock -= 94 * 2

									miniBossEffectsMAP[zapBlock] = "zap7"
								} else {
									zapBlock := currentMiniBossBlock - 94*4
									zapBlock++
									miniBossEffectsMAP[zapBlock] = "zap7"
									zapBlock -= 94 * 2

									miniBossEffectsMAP[zapBlock] = "zap7"

								}
							}

						case 6:

							chooseZap := flipcoin()

							if chooseZap {

								if miniBossRightLeft {
									zapBlock := currentMiniBossBlock + 94*6
									zapBlock += 5
									miniBossEffectsMAP[zapBlock] = "zap7"
									zapBlock -= 94 * 2
									miniBossEffectsMAP[zapBlock] = "zap7"
								} else {
									zapBlock := currentMiniBossBlock + 94*6
									zapBlock++
									miniBossEffectsMAP[zapBlock] = "zap7"
									zapBlock -= 94 * 2
									miniBossEffectsMAP[zapBlock] = "zap7"

								}
							}

						}

					}
				}

			} // end update mini boss

			// MARK: frameCount%6 timer

			if frameCount%6 == 0 {

				if miniBossOn1 || miniBossOn2 || miniBossOn3 {

					switch miniBossType {
					case 1:
						if miniBossJumpOn == false {
							bossSpecial := rolldice() + rolldice()
							if bossSpecial == 12 {
								miniBossJumpOn = true
							}
						}
					case 2:
						if miniBossRollOn == false {
							bossSpecial := rolldice() + rolldice()
							if bossSpecial == 12 {
								miniBossRollOn = true
							}
						}

					case 3:
						if miniBossWallJumpOn == false {
							miniBossHorizontal := miniBossCurrentBlock / 94
							miniBossVertical := miniBossCurrentBlock - (miniBossHorizontal * 94)
							if miniBossJumpTimerOn {
								if miniBossVertical > 6 || miniBossVertical < 80 {
									chooseJump := flipcoin()
									if chooseJump {
										miniBossWallJumpOn = true
									}

								}
							}
						}

					}

				}

				// weapon text color change timer
				weaponTextColorChange++

				if weaponTextColorChange == 7 {
					weaponTextColorChange = 0
				}

				switch weaponTextColorChange {

				case 0:
					weaponTextColor = rl.Red
				case 1:
					weaponTextColor = rl.Yellow
				case 2:
					weaponTextColor = rl.Pink
				case 3:
					weaponTextColor = rl.Blue
				case 4:
					weaponTextColor = rl.Violet
				case 5:
					weaponTextColor = rl.Lime
				case 6:
					weaponTextColor = rl.Orange
				case 7:
					weaponTextColor = rl.Maroon

				}

			} // end frameCount%6 timer

			// MARK: frameCount%4 timer
			if frameCount%4 == 0 {

				for a := 0; a < 5264; a++ {
					checkPlatformsBlock := platformsMAP[a]

					checkEffectsBlock := effectsMAP[a]

					if checkEffectsBlock == "bomb4" {
						effectsMAP[a] = "bomb3"
					} else if checkEffectsBlock == "bomb3" {
						effectsMAP[a] = "bomb2"
					} else if checkEffectsBlock == "bomb2" {
						effectsMAP[a] = "bomb1"
					} else if checkEffectsBlock == "bomb1" {
						effectsMAP[a] = ""
					}

					// moving platforms up
					if checkPlatformsBlock == "movingPlatformUp" {

						movingPlatformHorizontal := a / 94

						if movingPlatformHorizontal > 8 {

							levelMAP[a] = ""
							platformsMAP[a] = ""
							levelMAP[a-94] = "floor"
							platformsMAP[a-94] = "movingPlatformUp"
							backgroundObjectsMAP[a] = ""
							backgroundObjectsMAP[a-94] = ""
							backgroundObjectsMAP[a-(94*2)] = ""
							backgroundObjectsMAP[a-(94*3)] = ""

						} else if movingPlatformHorizontal <= 8 {

							platformsMAP[a] = "movingPlatformDown"

						}

					}

					if checkPlatformsBlock == "movingPlatformUpTL" {

						movingPlatformHorizontal := a / 94

						if movingPlatformHorizontal > 8 {

							levelMAP[a] = ""
							platformsMAP[a] = ""
							levelMAP[a-94] = "floor"
							platformsMAP[a-94] = "movingPlatformUpTL"
							backgroundObjectsMAP[a] = ""
							backgroundObjectsMAP[a-94] = ""
							backgroundObjectsMAP[a-(94*2)] = ""
							backgroundObjectsMAP[a-(94*3)] = ""

						} else if movingPlatformHorizontal <= 8 {

							platformsMAP[a] = "movingPlatformDownTL"

						}

					}

				}

			}
			// MARK: frameCount%9 timer
			if frameCount%9 == 0 {

				// weapon up down
				if weapongameud {
					weapongameud = false
				} else {
					weapongameud = true
				}

				// update teleport image
				teleportIMG.X += 64
				if teleportIMG.X == 520 {
					teleportIMG.X = 8
				}

				backObjGround43IMG.X += 48

				if backObjGround43IMG.X > 985 {
					backObjGround43IMG.X = 841
				}

				// update weather
				if cloudsOn {
					// move clouds
					drawScreenCurrentBlockWeatherNEXT++
				}

				// update coin img
				coinIMG.X += 16
				if coinIMG.X == 414 {
					coinIMG.X = 334
				}
				// update popup image
				getArrowIMG.X += 17
				if getArrowIMG.X == 51 {
					getArrowIMG.X = 0
				}

				if enemiesScreenCount == 0 && secondsCountdown < 59 && secondsCountdown > 0 {
					cENEMIES()
				}

				// update dino idle
				switch playerDirection {
				case "right":
					switch dinoType {
					case "greenDino":
						dinoGreenRIMG.X += 24
						if dinoGreenRIMG.X > 52 {
							dinoGreenRIMG.X = 4
						}
					case "redDino":
						dinoRedRIMG.X += 24
						if dinoRedRIMG.X > 53 {
							dinoRedRIMG.X = 5
						}
					case "yellowDino":
						dinoYellowRIMG.X += 24
						if dinoYellowRIMG.X > 52 {
							dinoYellowRIMG.X = 4
						}
					case "blueDino":
						dinoBlueRIMG.X += 24
						if dinoBlueRIMG.X > 52 {
							dinoBlueRIMG.X = 4
						}
					}

				case "left":
					switch dinoType {
					case "greenDino":
						dinoGreenLIMG.X -= 24
						if dinoGreenLIMG.X < 269 {
							dinoGreenLIMG.X = 317
						}
					case "redDino":
						dinoRedLIMG.X -= 24
						if dinoRedLIMG.X < 267 {
							dinoRedLIMG.X = 315
						}
					case "yellowDino":
						dinoYellowLIMG.X -= 24
						if dinoYellowLIMG.X < 269 {
							dinoYellowLIMG.X = 317
						}
					case "blueDino":
						dinoBlueLIMG.X -= 24
						if dinoBlueLIMG.X < 269 {
							dinoBlueLIMG.X = 317
						}
					}

				}

			} // end framecount%9 timer

			// player enemies collision timer
			if hpLossPause {
				collisionCount++
				if collisionCount == 60 {
					collisionCount = 0
					hpLossPause = false
				}
			}

			// MARK: move bullets
			// left movement
			for a := 0; a < 5264; a++ {

				checkBulletsBlock := bulletsMAP[a]

				// MARK: bullet collisions
				if checkBulletsBlock == "bulletL" {

					bulletHorizontal := a / 94
					bulletVertical := a - (bulletHorizontal * 94)

					if bulletVertical <= 2 {
						bulletsMAP[a] = ""
					} else if levelMAP[a-1] == "floor" {
						bulletsMAP[a] = ""
					} else if enemiesMAP[a-1] == "enemy" && enemiesMAP[a-95] == "enemy" {
						bulletsMAP[a] = ""
						clearEnemyBlock = a - 96
						clearEnemy()
						killEnemy()
					} else if enemiesMAP[a-1] == "enemy" && enemiesMAP[a-95] == "boundary" {
						bulletsMAP[a] = ""
						clearEnemyBlock = a - 2
						clearEnemy()
						killEnemy()
					} else if miniBossMAP[a-1] != "" {
						if miniBossHPpause == false {
							miniBossHPpause = true
							miniBossHP--
							bulletsMAP[a] = ""
						}
					} else {
						bulletsMAP[a] = ""
						bulletsMAP[a-1] = "bulletL"
					}
				}

				if checkBulletsBlock == "bulletEnemyL" {

					bulletHorizontal := a / 94
					bulletVertical := a - (bulletHorizontal * 94)

					if bulletVertical <= 2 {
						bulletsMAP[a] = ""
					} else if levelMAP[a-1] == "floor" {
						bulletsMAP[a] = ""
					} else if bulletsMAP[a-1] == "bulletR" {
						bulletsMAP[a-1] = ""
						bulletsMAP[a] = ""
						bulletsMAP[a-1] = "bulletEnemyL"
					} else if playerMAP[a-1] == "player" {
						if hpLossPause == false {
							bulletsMAP[a] = ""
							playerHP--
							hpLossPause = true
							hpLossSound = true
						} else {
							bulletsMAP[a] = ""
						}

					} else {
						bulletsMAP[a] = ""
						bulletsMAP[a-1] = "bulletEnemyL"
					}
				}

				if checkBulletsBlock == "bulletU" {

					bulletHorizontal := a / 94

					if bulletHorizontal <= 2 {
						bulletsMAP[a] = ""
					} else if levelMAP[a-94] == "floor" {
						bulletsMAP[a] = ""
					} else if enemiesMAP[a-94] == "enemy" && enemiesMAP[a-93] == "enemy" {
						bulletsMAP[a] = ""
						clearEnemyBlock = a - 188
						clearEnemy()
						killEnemy()
					} else if enemiesMAP[a-94] == "enemy" && enemiesMAP[a-95] == "enemy" {
						bulletsMAP[a] = ""
						clearEnemyBlock = a - 189
						clearEnemy()
						killEnemy()
					} else if miniBossMAP[a-94] != "" {
						if miniBossHPpause == false {
							miniBossHPpause = true
							miniBossHP--
							bulletsMAP[a] = ""
						}
					} else {
						bulletsMAP[a] = ""
						bulletsMAP[a-94] = "bulletU"
					}
				}

			}

			// right movement
			for a := 5263; a >= 0; a-- {
				checkBulletsBlock := bulletsMAP[a]

				checkActiveSpecialBlock := activeSpecialMAP[a]

				if checkActiveSpecialBlock == "cannonball" {

					cannoballHorizontal := a / 94
					cannonballVertical := a - (cannoballHorizontal * 94)

					if cannonballVertical > 90 {
						activeSpecialMAP[a] = ""
					} else {

						activeSpecialMAP[a] = ""
						activeSpecialMAP[a+1] = "cannonball"
					}
				}

				if checkBulletsBlock == "bulletEnemyD" {

					bulletHorizontal := a / 94
					bulletVertical := a - (bulletHorizontal * 94)

					if bulletVertical <= 2 {
						bulletsMAP[a] = ""
					} else if bulletHorizontal >= 50 {
						bulletsMAP[a] = ""
					} else if levelMAP[a+94] == "floor" {
						bulletsMAP[a] = ""
					} else if playerMAP[a+94] == "player" || playerMAP[a+94] == "playerTL" || playerMAP[a+95] == "player" || playerMAP[a+95] == "playerTL" {
						if hpLossPause == false {
							bulletsMAP[a] = ""
							playerHP--
							hpLossPause = true
							hpLossSound = true
						} else {
							bulletsMAP[a] = ""
						}
					} else {
						bulletsMAP[a] = ""
						bulletsMAP[a+94] = "bulletEnemyD"
					}
				}

				if checkBulletsBlock == "bulletR" {
					bulletHorizontal := a / 94
					bulletVertical := a - (bulletHorizontal * 94)

					if bulletVertical >= 91 {
						bulletsMAP[a] = ""
					} else if levelMAP[a+1] == "floor" {
						bulletsMAP[a] = ""
					} else if enemiesMAP[a+1] == "enemyTL" && enemiesMAP[a+95] == "enemy" {
						bulletsMAP[a] = ""
						clearEnemyBlock = a + 1
						clearEnemy()
						killEnemy()
					} else if enemiesMAP[a+1] == "enemy" && enemiesMAP[a+95] == "boundary" {
						bulletsMAP[a] = ""
						clearEnemyBlock = a - 93
						clearEnemy()
						killEnemy()
					} else if miniBossMAP[a+1] != "" {
						if miniBossHPpause == false {
							miniBossHPpause = true
							miniBossHP--
							bulletsMAP[a] = ""
						}
					} else {
						bulletsMAP[a] = ""
						bulletsMAP[a+1] = "bulletR"
					}
				}
				if checkBulletsBlock == "bulletEnemyR" {
					bulletHorizontal := a / 94
					bulletVertical := a - (bulletHorizontal * 94)

					if bulletVertical >= 91 {
						bulletsMAP[a] = ""
					} else if levelMAP[a+1] == "floor" {
						bulletsMAP[a] = ""
					} else if playerMAP[a+1] == "player" || playerMAP[a+1] == "playerTL" {
						if hpLossPause == false {
							bulletsMAP[a] = ""
							playerHP--
							hpLossPause = true
							hpLossSound = true
						} else {
							bulletsMAP[a] = ""
						}
					} else {
						bulletsMAP[a] = ""
						bulletsMAP[a+1] = "bulletEnemyR"
					}
				}
			}

			// MARK: teleport
			if teleportActive {
				if frameCount%60 == 0 {
					teleportTimer--
					teleportTimer2--
				}
				if teleportTimer == 0 {
					teleportOn = true
					teleportTimer--
				}
				if teleportTimer2 == 0 {
					for a := 0; a < 5264; a++ {
						teleportsMAP[a] = ""
					}
					teleportTimer2--

				}
			}

			if teleportOn {
				placeTeleport := rInt(3770, 3850)
				teleportPostion1 = placeTeleport
				teleportsMAP[placeTeleport] = "teleportTL"
				for a := 0; a < 5; a++ {
					placeTeleport += 94
					teleportsMAP[placeTeleport] = "teleport"
				}

				placeTeleport = rInt(1890, 1970)
				placeTeleport += (rInt(-8, 9) * 94)
				teleportPostion2 = placeTeleport
				teleportsMAP[placeTeleport] = "teleportTL"
				for a := 0; a < 5; a++ {
					placeTeleport += 94
					teleportsMAP[placeTeleport] = "teleport"
				}
				teleportOn = false
			}

			// MARK: screenShake

			if screenShake {

				camera.Zoom = 2.3

				if screenShake2 == false {

					camera.Target.X = 60.0
					camera.Rotation = -3.0
					camera.Target.Y = 60.0
					screenShake2 = true
					screenShake3 = true
				}
				if screenShake3 {
					screenShakeTimer--

					if screenShakeTimer == 0 {
						camera.Target.X = 5.0
						camera.Rotation = 3.0
						screenShakeTimer = 4
						screenShake3 = false
						screenShake4 = true

					}

				}
				if screenShake4 {
					screenShakeTimer--
					if screenShakeTimer == 0 {
						camera.Target.X = 0.0
						camera.Target.Y = 0.0
						camera.Rotation = 0.0
						screenShakeTimer = 4
						camera.Zoom = 2.0
						screenShake4 = false
						screenShake = false
						screenShake2 = false

					}

				}

			}

			// MARK: info display

			// background rectangle left top
			if playerHPmax == 3 {
				rl.DrawRectangle(0, 12, 220, 70, rl.Fade(rl.LightGray, 0.6))
				rl.DrawRectangle(0, 82, 217, 3, rl.Fade(rl.Gray, 0.5))
			} else if playerHPmax == 4 {
				rl.DrawRectangle(0, 12, 280, 70, rl.Fade(rl.LightGray, 0.6))
				rl.DrawRectangle(0, 82, 277, 3, rl.Fade(rl.Gray, 0.5))
			} else if playerHPmax == 5 {
				rl.DrawRectangle(0, 12, 340, 70, rl.Fade(rl.LightGray, 0.6))
				rl.DrawRectangle(0, 82, 337, 3, rl.Fade(rl.Gray, 0.5))
			} else if playerHPmax == 6 {
				rl.DrawRectangle(0, 12, 400, 70, rl.Fade(rl.LightGray, 0.6))
				rl.DrawRectangle(0, 82, 397, 3, rl.Fade(rl.Gray, 0.5))
			}

			// background rectangle right top
			rl.DrawRectangle(screenW-365, 12, 365, 100, rl.Fade(rl.LightGray, 0.6))
			rl.DrawRectangle(screenW-362, 112, 362, 3, rl.Fade(rl.Gray, 0.5))

			// hp

			heartX := float32(20)
			heartV2 := rl.NewVector2(heartX, 20)
			heartV2back := rl.NewVector2(heartX-4, 24)
			for a := 0; a < playerHPmax; a++ {
				rl.DrawTextureRec(imgs, heartIMG, heartV2back, rl.Black)
				rl.DrawTextureRec(imgs, heartIMG, heartV2, rl.Fade(rl.Yellow, 0.7))
				heartV2.X += 60
				heartV2back.X += 60
			}
			heartX = float32(20)
			heartV2 = rl.NewVector2(heartX, 20)
			for a := 0; a < playerHP; a++ {
				rl.DrawTextureRec(imgs, heartIMG, heartV2, rl.Red)
				heartV2.X += 60
			}

			// countdown timer
			secondsCountdownTEXT := strconv.Itoa(secondsCountdown)

			rl.DrawText("level ends in", screenW-353, 23, 40, rl.Black)
			rl.DrawText("level ends in", screenW-351, 21, 40, rl.White)
			rl.DrawText("level ends in", screenW-350, 20, 40, rl.Black)
			rl.DrawText(secondsCountdownTEXT, screenW-83, 23, 40, rl.Black)
			rl.DrawText(secondsCountdownTEXT, screenW-81, 21, 40, rl.White)
			rl.DrawText(secondsCountdownTEXT, screenW-80, 20, 40, rl.Blue)

			// kills count text
			killzCountTEXT := strconv.Itoa(killzCount)

			rl.DrawText("killz", screenW-171, 63, 40, rl.Black)
			rl.DrawText("killz", screenW-173, 61, 40, rl.White)
			rl.DrawText("killz", screenW-174, 60, 40, rl.Black)

			rl.DrawText(killzCountTEXT, screenW-79, 63, 40, rl.Black)
			rl.DrawText(killzCountTEXT, screenW-81, 61, 40, rl.White)
			rl.DrawText(killzCountTEXT, screenW-82, 60, 40, rl.Red)

			// coins count text
			playerCoinsTEXT := strconv.Itoa(playerCoins)

			rl.DrawText(playerCoinsTEXT, screenW-232, 63, 40, rl.Black)
			rl.DrawText(playerCoinsTEXT, screenW-234, 61, 40, rl.White)
			rl.DrawText(playerCoinsTEXT, screenW-235, 60, 40, rl.Gold)

			rl.BeginMode2D(camera)
			coinInfoV2 := rl.NewVector2(546, 31)
			rl.DrawTextureRec(imgs, coinIMG, coinInfoV2, rl.White)
			rl.EndMode2D()

			// MARK: scan lines pixel noise
			// noise lines
			if noiseLinesOn {
				if frameCountGameStart%60 == 0 {
					if noiseLinesScreenOn {
						noiseLinesScreenOn = false
					} else {
						noiseLinesScreenOn = true
					}
				}

				if noiseLinesScreenOn {
					for a := 0; a < noiseLineDistance1; a++ {
						noiseLineX1Change++
					}
					for a := 0; a < noiseLineDistance2; a++ {
						noiseLineX2Change++
					}
					for a := 0; a < noiseLineDistance3; a++ {
						noiseLineX3Change++
					}
					for a := 0; a < noiseLineDistance4; a++ {
						noiseLineX4Change++
					}
					if noiseLineLR1 {
						rl.DrawLine(noiseLineX1+noiseLineX1Change, 0, noiseLineX1+noiseLineX1Change, screenH, rl.Fade(rl.Black, 0.5))
					} else {
						rl.DrawLine(noiseLineX1-noiseLineX1Change, 0, noiseLineX1-noiseLineX1Change, screenH, rl.Fade(rl.Black, 0.5))
					}
					if noiseLineLR2 {
						rl.DrawLine(noiseLineX2+noiseLineX2Change, 0, noiseLineX1+noiseLineX2Change, screenH, rl.Fade(rl.Black, 0.5))
					} else {
						rl.DrawLine(noiseLineX2-noiseLineX2Change, 0, noiseLineX2-noiseLineX2Change, screenH, rl.Fade(rl.Black, 0.5))
					}
					if noiseLineLR3 {
						rl.DrawLine(noiseLineX3+noiseLineX3Change, 0, noiseLineX3+noiseLineX3Change, screenH, rl.Fade(rl.Black, 0.5))
						rl.DrawLine(noiseLineX3-noiseLineX3Change, 0, noiseLineX3-noiseLineX3Change, screenH, rl.Fade(rl.Black, 0.5))
					}
					if noiseLineLR4 {
						rl.DrawLine(noiseLineX4+noiseLineX4Change, 0, noiseLineX4+noiseLineX4Change, screenH, rl.Fade(rl.Black, 0.5))
					} else {
						rl.DrawLine(noiseLineX4-noiseLineX4Change, 0, noiseLineX4-noiseLineX4Change, screenH, rl.Fade(rl.Black, 0.5))
					}
				} else {
					cNOISELINES()

					noiseLineX1Change = 0
					noiseLineX2Change = 0
					noiseLineX3Change = 0
					noiseLineX4Change = 0
					noiseLineX1 = noiseLinesMAP[0]
					noiseLineX2 = noiseLinesMAP[1]
					noiseLineX3 = noiseLinesMAP[2]
					noiseLineX4 = noiseLinesMAP[3]
					noiseLineDistance1 = noiseLinesDistanceMAP[0]
					noiseLineDistance2 = noiseLinesDistanceMAP[1]
					noiseLineDistance3 = noiseLinesDistanceMAP[2]
					noiseLineDistance4 = noiseLinesDistanceMAP[3]
					noiseLineLR1 = noiseLinesLRMAP[0]
					noiseLineLR2 = noiseLinesLRMAP[1]
					noiseLineLR3 = noiseLinesLRMAP[2]
					noiseLineLR4 = noiseLinesLRMAP[3]

				}

			} // end noise lines
			// draw pixel noise
			if pixelNoiseOn {
				if frameCountGameStart%2 == 0 {
					cPIXELNOISE()
				}

				lineCountPixelNoise := 0
				pixelNoiseY := int32(0)
				pixelNoiseX := int32(0)
				for a := 0; a < 880; a++ {

					if pixelNoiseMAP[a] == true {
						rl.DrawRectangle(pixelNoiseX, pixelNoiseY, 2, 2, rl.Black)
					}

					lineCountPixelNoise += 34
					pixelNoiseX += 34
					if lineCountPixelNoise > 1350 {
						lineCountPixelNoise = 0
						pixelNoiseX = 0
						pixelNoiseY += 34

					}
				}

			}

			// draw scan lines main game window
			if scanLinesOn {
				if switchScanLines {
					linesY := int32(0)
					for a := 0; a < int(screenH); a++ {
						rl.DrawLine(0, linesY, screenW, linesY, rl.Fade(rl.Black, 0.3))
						linesY += 2
						a++
					}
				} else {
					linesY := int32(1)
					for a := 0; a < int(screenH); a++ {
						rl.DrawLine(0, linesY, screenW, linesY, rl.Fade(rl.Black, 0.3))
						linesY += 2
						a++
					}

				}
			}

		} // end pauseOn

		// MARK: active special item

		if activeSpecialOn {
			if rl.IsKeyPressed(rl.KeySpace) {
				activeSpecialActive = true
			}

			if activeSpecialActive {
				switch activeSpecialItem {

				case "apple":

					if activeSpecialComplete == false {

						appleStartBlock = 0

						for {
							changeAppleBlock := (rInt(4, 9) * 94)
							changeAppleBlockHorizontal := rInt(0, 4)
							appleStartBlock += changeAppleBlock + changeAppleBlockHorizontal
							if appleStartBlock > 4042 {
								break
							}
							activeSpecialMAP[appleStartBlock] = "appleTL"
							activeSpecialMAP[appleStartBlock+1] = "apple"
							activeSpecialMAP[appleStartBlock+94] = "apple"
							activeSpecialMAP[appleStartBlock+95] = "apple"

							appleStartBlock -= changeAppleBlockHorizontal

						}

						activeSpecialComplete = true

					}

					if frameCount%400 == 0 {
						if activeSpecialComplete {
							appleStartBlock = 0
							activeSpecialComplete = false
						}
					}

				case "petRedSlime":
					if activeSpecialComplete == false {

						petRedSlimeStartBlock = rInt(3964, 4044)
						activeSpecialMAP[petRedSlimeStartBlock] = "petRedSlimeTL"
						activeSpecialMAP[petRedSlimeStartBlock+1] = "petRedSlime"
						activeSpecialMAP[petRedSlimeStartBlock+94] = "petRedSlime"
						activeSpecialMAP[petRedSlimeStartBlock+95] = "petRedSlime"

						activeSpecialComplete = true
					}

				case "petSlime":
					if activeSpecialComplete == false {

						petSlimeStartBlock = rInt(3964, 4044)
						activeSpecialMAP[petSlimeStartBlock] = "petSlimeTL"
						activeSpecialMAP[petSlimeStartBlock+1] = "petSlime"
						activeSpecialMAP[petSlimeStartBlock+94] = "petSlime"
						activeSpecialMAP[petSlimeStartBlock+95] = "petSlime"

						activeSpecialComplete = true
					}

				case "watermelon":
					if activeSpecialComplete == false {

						watermelonStartBlock := rInt(10, 25)

						for {
							changeWatermelonBlock := rInt(8, 15)
							watermelonStartBlock += changeWatermelonBlock

							activeSpecialMAP[watermelonStartBlock] = "watermelonTL"
							activeSpecialMAP[watermelonStartBlock+1] = "watermelon"
							activeSpecialMAP[watermelonStartBlock+94] = "watermelon"
							activeSpecialMAP[watermelonStartBlock+95] = "watermelon"

							if watermelonStartBlock > 78 {
								break
							}

						}
						activeSpecialComplete = true
					}

					if frameCount%121 == 0 {
						if activeSpecialComplete {
							activeSpecialComplete = false
						}
					}
				case "cherries":
					if activeSpecialComplete == false {

						cherriesStartBlock := rInt(10, 25)

						for {
							changeCherriesBlock := rInt(8, 15)
							cherriesStartBlock += changeCherriesBlock

							activeSpecialMAP[cherriesStartBlock] = "cherriesTL"
							activeSpecialMAP[cherriesStartBlock+1] = "cherries"
							activeSpecialMAP[cherriesStartBlock+94] = "cherries"
							activeSpecialMAP[cherriesStartBlock+95] = "cherries"

							if cherriesStartBlock > 78 {
								break
							}

						}
						activeSpecialComplete = true
					}

					if frameCount%121 == 0 {
						if activeSpecialComplete {
							activeSpecialComplete = false
						}
					}
				case "propellor":
					if activeSpecialComplete == false {
						propellorOn = true
						activeSpecialComplete = true
					}
				case "exclamation":
					if activeSpecialComplete == false {
						exclamationOn = true
						activeSpecialComplete = true
					}

				case "trampoline":
					if activeSpecialComplete == false {
						trampolineOn = true
						trampolineNumber := rInt(2, 6)

						for {
							trampolineBlock := rInt(3954, 4034)

							if activeSpecialMAP[trampolineBlock] != "trampolineTL" && activeSpecialMAP[trampolineBlock+1] != "trampolineTL" && activeSpecialMAP[trampolineBlock+2] != "trampolineTL" && activeSpecialMAP[trampolineBlock+3] != "trampolineTL" && activeSpecialMAP[trampolineBlock+4] != "trampolineTL" && activeSpecialMAP[trampolineBlock-1] != "trampolineTL" && activeSpecialMAP[trampolineBlock-2] != "trampolineTL" && activeSpecialMAP[trampolineBlock-3] != "trampolineTL" && activeSpecialMAP[trampolineBlock-4] != "trampolineTL" {

								activeSpecialMAP[trampolineBlock] = "trampolineTL"
								activeSpecialMAP[trampolineBlock+1] = "trampoline"
								activeSpecialMAP[trampolineBlock+94] = "trampoline"
								activeSpecialMAP[trampolineBlock+95] = "trampoline"
								activeSpecialMAP[trampolineBlock-94] = "trampoline"
								activeSpecialMAP[trampolineBlock-95] = "trampoline"
								trampolineNumber--

							}
							if trampolineNumber == 0 {
								activeSpecialComplete = true
								break
							}

						}
					}

				case "cannon":
					if activeSpecialComplete == false {
						cannonTimer = 5
						cannonStartBlock = 4048
						activeSpecialMAP[cannonStartBlock] = "cannonTL"
						activeSpecialComplete = true
					}
				case "iceball":
					if activeSpecialComplete == false {
						iceballStartBlock = 376
						activeSpecialMAP[iceballStartBlock] = "iceballTL"
						activeSpecialComplete = true
					}
				case "poisonball":
					if activeSpecialComplete == false {
						poisonballStartBlock = 942
						activeSpecialMAP[poisonballStartBlock] = "poisonballTL"
						activeSpecialComplete = true
					}
				case "fireball":
					if activeSpecialComplete == false {
						fireballStartBlock = 4042
						activeSpecialMAP[fireballStartBlock] = "fireballTL"
						activeSpecialMAP[fireballStartBlock+94] = "fireball"
						activeSpecialComplete = true
					}
				case "sawblade":
					if activeSpecialComplete == false {
						sawbladeStartBlock = (rInt(5, 40) * 94)
						activeSpecialMAP[sawbladeStartBlock] = "sawbladeTL"
						activeSpecialMAP[sawbladeStartBlock+1] = "sawblade"
						activeSpecialMAP[sawbladeStartBlock+94] = "sawblade"
						activeSpecialMAP[sawbladeStartBlock+95] = "sawblade"
						activeSpecialDirection = rInt(1, 5)
						activeSpecialComplete = true
					}

				}

			}
		}

		// MARK: move/animate active special items

		if activeSpecialActive {

			switch activeSpecialItem {

			case "cherries":

				if frameCount%3 == 0 {

					for a := 5263; a >= 0; a-- {

						checkCherriesBlock := activeSpecialMAP[a]

						if checkCherriesBlock == "cherriesTL" {

							if levelMAP[a+(94*2)] != "floor" && levelMAP[a+((94*2)+1)] != "floor" {

								activeSpecialMAP[a] = ""
								activeSpecialMAP[a+1] = ""
								activeSpecialMAP[a+94] = ""
								activeSpecialMAP[a+95] = ""

								activeSpecialMAP[a+94] = "cherriesTL"
								activeSpecialMAP[a+95] = "cherries"
								activeSpecialMAP[a+(94*2)] = "cherries"
								activeSpecialMAP[a+((94*2)+1)] = "cherries"

								if enemiesMAP[a+94] == "enemy" || enemiesMAP[a+94] == "enemyTL" {

									if enemiesMAP[a+94] == "enemyTL" {
										clearEnemyBlock = findEnemyTL(a + 94)
										clearEnemy()
										killEnemy()
									} else if enemiesMAP[a+94] == "enemy" {
										if enemiesMAP[a] == "boundary" {
											clearEnemyBlock = findEnemyTL(a + 93)
											clearEnemy()
											killEnemy()
										} else if enemiesMAP[a] == "enemy" {
											clearEnemyBlock = findEnemyTL(a - 1)
											clearEnemy()
											killEnemy()
										} else if enemiesMAP[a] == "enemyTL" {
											clearEnemyBlock = findEnemyTL(a)
											clearEnemy()
											killEnemy()
										}
									}
								}
								if enemiesMAP[a+95] == "enemy" || enemiesMAP[a+95] == "enemyTL" {

									if enemiesMAP[a+95] == "enemyTL" {
										clearEnemyBlock = findEnemyTL(a + 95)
										clearEnemy()
										killEnemy()
									} else if enemiesMAP[a+95] == "enemy" {
										if enemiesMAP[a+1] == "boundary" {
											clearEnemyBlock = findEnemyTL(a + 94)
											clearEnemy()
											killEnemy()
										} else if enemiesMAP[a+1] == "enemy" {
											clearEnemyBlock = findEnemyTL(a)
											clearEnemy()
											killEnemy()
										} else if enemiesMAP[a+1] == "enemyTL" {
											clearEnemyBlock = findEnemyTL(a + 1)
											clearEnemy()
											killEnemy()
										}
									}

								}
								if enemiesMAP[a+(94*2)] == "enemy" || enemiesMAP[a+(94*2)] == "enemyTL" {

									if enemiesMAP[a+(94*2)] == "enemyTL" {
										clearEnemyBlock = findEnemyTL(a + (94 * 2))
										clearEnemy()
										killEnemy()
									} else if enemiesMAP[a+(94*2)] == "enemy" {
										if enemiesMAP[a+94] == "boundary" {
											clearEnemyBlock = findEnemyTL((a + (94 * 2)) - 1)
											clearEnemy()
											killEnemy()
										} else if enemiesMAP[a+94] == "enemy" {
											clearEnemyBlock = findEnemyTL(a + 93)
											clearEnemy()
											killEnemy()
										} else if enemiesMAP[a+95] == "enemyTL" {
											clearEnemyBlock = findEnemyTL(a + 95)
											clearEnemy()
											killEnemy()
										}
									}

								}
								if enemiesMAP[a+((94*2)+1)] == "enemy" || enemiesMAP[a+((94*2)+1)] == "enemyTL" {

									if enemiesMAP[a+((94*2)+1)] == "enemyTL" {
										clearEnemyBlock = findEnemyTL(a + ((94 * 2) + 1))
										clearEnemy()
										killEnemy()
									} else if enemiesMAP[a+((94*2)+1)] == "enemy" {
										if enemiesMAP[a+95] == "boundary" {
											clearEnemyBlock = findEnemyTL(a + (94 * 2))
											clearEnemy()
											killEnemy()
										} else if enemiesMAP[a+95] == "enemy" {
											clearEnemyBlock = findEnemyTL(a + 94)
											clearEnemy()
											killEnemy()
										} else if enemiesMAP[a+95] == "enemyTL" {
											clearEnemyBlock = findEnemyTL(a + 95)
											clearEnemy()
											killEnemy()
										}
									}

								}

							}
						}

					}

				}

			case "apple":

				if frameCount%3 == 0 {

					for a := 5263; a >= 0; a-- {
						checkAppleBlock := activeSpecialMAP[a]

						if checkAppleBlock == "appleTL" {

							appleHorizontal := a / 94
							appleVertical := a - (appleHorizontal * 94)

							if appleVertical > 90 {
								activeSpecialMAP[a] = ""
								activeSpecialMAP[a+1] = ""
								activeSpecialMAP[a+94] = ""
								activeSpecialMAP[a+95] = ""
							} else {

								activeSpecialMAP[a] = ""
								activeSpecialMAP[a+1] = ""
								activeSpecialMAP[a+94] = ""
								activeSpecialMAP[a+95] = ""

								activeSpecialMAP[a+1] = "appleTL"
								activeSpecialMAP[a+2] = "apple"
								activeSpecialMAP[a+95] = "apple"
								activeSpecialMAP[a+96] = "apple"
							}

						}

					}
				}
			case "petBat":
				if frameCount%6 == 0 {
					petBatHorizontal := petBatStartBlock / 94
					petBatVertical := petBatStartBlock - (petBatHorizontal * 94)

					if petBatLeftRight {
						if petBatVertical < 10 {
							petBatLeftRight = false
						}
						activeSpecialMAP[petBatStartBlock] = ""
						activeSpecialMAP[petBatStartBlock+1] = ""
						activeSpecialMAP[petBatStartBlock+94] = ""
						activeSpecialMAP[petBatStartBlock+95] = ""
						petBatStartBlock--
						activeSpecialMAP[petBatStartBlock] = "petBatTL"
						activeSpecialMAP[petBatStartBlock+1] = "petBat"
						activeSpecialMAP[petBatStartBlock+94] = "petBat"
						activeSpecialMAP[petBatStartBlock+95] = "petBat"
					} else {
						if petBatVertical > 84 {
							petBatLeftRight = true
						}
						activeSpecialMAP[petBatStartBlock] = ""
						activeSpecialMAP[petBatStartBlock+1] = ""
						activeSpecialMAP[petBatStartBlock+94] = ""
						activeSpecialMAP[petBatStartBlock+95] = ""
						petBatStartBlock++
						activeSpecialMAP[petBatStartBlock] = "petBatTL"
						activeSpecialMAP[petBatStartBlock+1] = "petBat"
						activeSpecialMAP[petBatStartBlock+94] = "petBat"
						activeSpecialMAP[petBatStartBlock+95] = "petBat"
					}
				}

				if frameCount%6 == 0 {
					batShoot := rolldice()
					if batShoot == 6 {
						activeSpecialMAP[petBatStartBlock+((94*4)+1)] = "zap9"
					}
				}

			case "petRedSlime":

				if frameCount%15 == 0 {
					exclamationIMG.X += 34
					if exclamationIMG.X == 1397 {
						exclamationIMG.X = 1295
					}
				}

				if frameCount%6 == 0 {
					redSlimeHorizontal := petRedSlimeStartBlock / 94
					redSlimeVertical := petRedSlimeStartBlock - (redSlimeHorizontal * 94)

					if petRedSlimeLeftRight {
						if redSlimeVertical < 10 {
							petRedSlimeLeftRight = false
						}
						activeSpecialMAP[petRedSlimeStartBlock] = ""
						activeSpecialMAP[petRedSlimeStartBlock+1] = ""
						activeSpecialMAP[petRedSlimeStartBlock+94] = ""
						activeSpecialMAP[petRedSlimeStartBlock+95] = ""
						petRedSlimeStartBlock--
						activeSpecialMAP[petRedSlimeStartBlock] = "petRedSlimeTL"
						activeSpecialMAP[petRedSlimeStartBlock+1] = "petRedSlime"
						activeSpecialMAP[petRedSlimeStartBlock+94] = "petRedSlime"
						activeSpecialMAP[petRedSlimeStartBlock+95] = "petRedSlime"
					} else {
						if redSlimeVertical > 84 {
							petRedSlimeLeftRight = true
						}
						activeSpecialMAP[petRedSlimeStartBlock] = ""
						activeSpecialMAP[petRedSlimeStartBlock+1] = ""
						activeSpecialMAP[petRedSlimeStartBlock+94] = ""
						activeSpecialMAP[petRedSlimeStartBlock+95] = ""
						petRedSlimeStartBlock++
						activeSpecialMAP[petRedSlimeStartBlock] = "petRedSlimeTL"
						activeSpecialMAP[petRedSlimeStartBlock+1] = "petRedSlime"
						activeSpecialMAP[petRedSlimeStartBlock+94] = "petRedSlime"
						activeSpecialMAP[petRedSlimeStartBlock+95] = "petRedSlime"
					}
				}

			case "petKnight":

				if frameCount%30 == 0 {
					knightShoot := rolldice()
					if knightShoot == 6 {
						activeSpecialMAP[petKnightStartBlock+6] = "zap11"
						activeSpecialMAP[petKnightStartBlock-6] = "zap11"
						activeSpecialMAP[petKnightStartBlock+((94*2)+3)] = "zap11"
						activeSpecialMAP[petKnightStartBlock+((94*2)-3)] = "zap11"
						activeSpecialMAP[petKnightStartBlock+((94*4)+5)] = "zap11"
						activeSpecialMAP[petKnightStartBlock+((94*4)-5)] = "zap11"
					}

				}

				if frameCount%6 == 0 {
					if knightBackFadeOn {
						if knightBackFade < 0.7 {
							knightBackFade += 0.1
						} else if knightBackFade >= 0.7 {
							knightBackFadeOn = false
						}
					} else {
						if knightBackFade > 0.1 {
							knightBackFade -= 0.1
						} else if knightBackFade <= 0.1 {
							knightBackFadeOn = true
						}
					}

					knightHorizontal := petKnightStartBlock / 94
					knightVertical := petKnightStartBlock - (knightHorizontal * 94)

					if petKnightLeftRight {
						if knightVertical < 10 {
							petSlimeLeftRight = false
						}
						activeSpecialMAP[petKnightStartBlock] = ""
						activeSpecialMAP[petKnightStartBlock+1] = ""
						activeSpecialMAP[petKnightStartBlock+94] = ""
						activeSpecialMAP[petKnightStartBlock+95] = ""
						petKnightStartBlock--
						activeSpecialMAP[petKnightStartBlock] = "petKnightTL"
						activeSpecialMAP[petKnightStartBlock+1] = "petKnight"
						activeSpecialMAP[petKnightStartBlock+94] = "petKnight"
						activeSpecialMAP[petKnightStartBlock+95] = "petKnight"
					} else {
						if knightVertical > 84 {
							petKnightLeftRight = true
						}
						activeSpecialMAP[petKnightStartBlock] = ""
						activeSpecialMAP[petKnightStartBlock+1] = ""
						activeSpecialMAP[petKnightStartBlock+94] = ""
						activeSpecialMAP[petKnightStartBlock+95] = ""
						petKnightStartBlock++
						activeSpecialMAP[petKnightStartBlock] = "petKnightTL"
						activeSpecialMAP[petKnightStartBlock+1] = "petKnight"
						activeSpecialMAP[petKnightStartBlock+94] = "petKnight"
						activeSpecialMAP[petKnightStartBlock+95] = "petKnight"
					}

					if knightJumpOn == false {
						petKnightGroundBlockL := levelMAP[petKnightStartBlock+(94*2)]
						petKnightGroundBlockR := levelMAP[petKnightStartBlock+((94*2)+1)]

						if petKnightGroundBlockL != "floor" && petKnightGroundBlockR != "floor" {
							activeSpecialMAP[petKnightStartBlock] = ""
							activeSpecialMAP[petKnightStartBlock+1] = ""
							activeSpecialMAP[petKnightStartBlock+94] = ""
							activeSpecialMAP[petKnightStartBlock+95] = ""
							petKnightStartBlock += 94
							activeSpecialMAP[petKnightStartBlock] = "petKnightTL"
							activeSpecialMAP[petKnightStartBlock+1] = "petKnight"
							activeSpecialMAP[petKnightStartBlock+94] = "petKnight"
							activeSpecialMAP[petKnightStartBlock+95] = "petKnight"

						}
					}

					if knightJumpOn == true {

						if knightHorizontal > 6 {

							if knightJumpHeight != 0 {
								activeSpecialMAP[petKnightStartBlock] = ""
								activeSpecialMAP[petKnightStartBlock+1] = ""
								activeSpecialMAP[petKnightStartBlock+94] = ""
								activeSpecialMAP[petKnightStartBlock+95] = ""
								petKnightStartBlock -= 94
								activeSpecialMAP[petKnightStartBlock] = "petKnightTL"
								activeSpecialMAP[petKnightStartBlock+1] = "petKnight"
								activeSpecialMAP[petKnightStartBlock+94] = "petKnight"
								activeSpecialMAP[petKnightStartBlock+95] = "petKnight"
								knightJumpHeight--
							} else if knightJumpHeight == 0 {
								knightJumpOn = false
							}

						} else {
							knightJumpHeight = 0
							knightJumpOn = false
						}

					}

				}

				if knightJumpOn == false {
					if frameCount%15 == 0 {
						knightJump := rolldice()
						if knightJump == 6 {
							knightJumpHeight = 20
							knightJumpOn = true
						}
					}
				}

			case "petSlime":
				if frameCount%6 == 0 {
					slimeHorizontal := petSlimeStartBlock / 94
					slimeVertical := petSlimeStartBlock - (slimeHorizontal * 94)

					if petSlimeLeftRight {
						if slimeVertical < 10 {
							petSlimeLeftRight = false
						}
						activeSpecialMAP[petSlimeStartBlock] = ""
						activeSpecialMAP[petSlimeStartBlock+1] = ""
						activeSpecialMAP[petSlimeStartBlock+94] = ""
						activeSpecialMAP[petSlimeStartBlock+95] = ""
						petSlimeStartBlock--
						activeSpecialMAP[petSlimeStartBlock] = "petSlimeTL"
						activeSpecialMAP[petSlimeStartBlock+1] = "petSlime"
						activeSpecialMAP[petSlimeStartBlock+94] = "petSlime"
						activeSpecialMAP[petSlimeStartBlock+95] = "petSlime"
					} else {
						if slimeVertical > 84 {
							petSlimeLeftRight = true
						}
						activeSpecialMAP[petSlimeStartBlock] = ""
						activeSpecialMAP[petSlimeStartBlock+1] = ""
						activeSpecialMAP[petSlimeStartBlock+94] = ""
						activeSpecialMAP[petSlimeStartBlock+95] = ""
						petSlimeStartBlock++
						activeSpecialMAP[petSlimeStartBlock] = "petSlimeTL"
						activeSpecialMAP[petSlimeStartBlock+1] = "petSlime"
						activeSpecialMAP[petSlimeStartBlock+94] = "petSlime"
						activeSpecialMAP[petSlimeStartBlock+95] = "petSlime"
					}

					petSlimeGroundBlockL := levelMAP[petSlimeStartBlock+(94*2)]
					petSlimeGroundBlockR := levelMAP[petSlimeStartBlock+((94*2)+1)]

					if petSlimeGroundBlockL != "floor" && petSlimeGroundBlockR != "floor" {
						activeSpecialMAP[petSlimeStartBlock] = ""
						activeSpecialMAP[petSlimeStartBlock+1] = ""
						activeSpecialMAP[petSlimeStartBlock+94] = ""
						activeSpecialMAP[petSlimeStartBlock+95] = ""
						petSlimeStartBlock += 94
						activeSpecialMAP[petSlimeStartBlock] = "petSlimeTL"
						activeSpecialMAP[petSlimeStartBlock+1] = "petSlime"
						activeSpecialMAP[petSlimeStartBlock+94] = "petSlime"
						activeSpecialMAP[petSlimeStartBlock+95] = "petSlime"

					}
				}
				if frameCount%15 == 0 {

					slimeShoot := rolldice()

					if slimeShoot == 6 {

						zapChange := rInt(3, 10)
						zapChange1 := rInt(3, 10)
						zapChange2 := rInt(3, 10)
						zapChange3 := rInt(-12, 12)
						zapChange4 := rInt(-12, 12)
						zapChange5 := rInt(-12, 12)

						activeSpecialMAP[petSlimeStartBlock-((94*zapChange)+zapChange3)] = "zap12"
						activeSpecialMAP[petSlimeStartBlock-((94*zapChange1)+zapChange4)] = "zap12"
						activeSpecialMAP[petSlimeStartBlock-((94*zapChange2)+zapChange5)] = "zap12"

					}

				}

			case "petSkeleton":

				if skeletonHasDropped1 == false {
					if skeletonDropTimer1 != 0 && skeletonDropTimer1 == secondsCountdown {
						dropNewWeapon = true
						skeletonHasDropped1 = true
					}
				}
				if skeletonHasDropped2 == false {
					if skeletonDropTimer2 != 0 && skeletonDropTimer2 == secondsCountdown {
						dropNewWeapon = true
						skeletonHasDropped2 = true
					}
				}
				if skeletonHasDropped3 == false {
					if skeletonDropTimer3 != 0 && skeletonDropTimer3 == secondsCountdown {
						dropNewWeapon = true
						skeletonHasDropped3 = true
					}
				}

				if frameCount%8 == 0 {
					if skeletonShield1 == true {
						skeletonShield1 = false
						skeletonShield2 = true
					} else if skeletonShield2 == true {
						skeletonShield2 = false
						skeletonShield3 = true
					} else if skeletonShield3 == true {
						skeletonShield3 = false
						skeletonShield4 = true
					} else if skeletonShield4 == true {
						skeletonShield4 = false
						skeletonShield1 = true
					}
				}

				if frameCount%6 == 0 {
					skeletonHorizontal := petSkeletonStartBlock / 94
					skeletonVertical := petSkeletonStartBlock - (skeletonHorizontal * 94)

					if petSkeletonLeftRight {
						if skeletonVertical < 10 {
							petSkeletonLeftRight = false
						}
						activeSpecialMAP[petSkeletonStartBlock] = ""
						activeSpecialMAP[petSkeletonStartBlock+1] = ""
						activeSpecialMAP[petSkeletonStartBlock+94] = ""
						activeSpecialMAP[petSkeletonStartBlock+95] = ""
						petSkeletonStartBlock--
						activeSpecialMAP[petSkeletonStartBlock] = "petSkeletonTL"
						activeSpecialMAP[petSkeletonStartBlock+1] = "petSkeleton"
						activeSpecialMAP[petSkeletonStartBlock+94] = "petSkeleton"
						activeSpecialMAP[petSkeletonStartBlock+95] = "petSkeleton"
					} else {
						if skeletonVertical > 84 {
							petSkeletonLeftRight = true
						}
						activeSpecialMAP[petSkeletonStartBlock] = ""
						activeSpecialMAP[petSkeletonStartBlock+1] = ""
						activeSpecialMAP[petSkeletonStartBlock+94] = ""
						activeSpecialMAP[petSkeletonStartBlock+95] = ""
						petSkeletonStartBlock++
						activeSpecialMAP[petSkeletonStartBlock] = "petSkeletonTL"
						activeSpecialMAP[petSkeletonStartBlock+1] = "petSkeleton"
						activeSpecialMAP[petSkeletonStartBlock+94] = "petSkeleton"
						activeSpecialMAP[petSkeletonStartBlock+95] = "petSkeleton"
					}

					petSkeletonGroundBlockL := levelMAP[petSkeletonStartBlock+(94*2)]
					petSkeletonGroundBlockR := levelMAP[petSkeletonStartBlock+((94*2)+1)]

					if petSkeletonGroundBlockL != "floor" && petSkeletonGroundBlockR != "floor" {
						activeSpecialMAP[petSkeletonStartBlock] = ""
						activeSpecialMAP[petSkeletonStartBlock+1] = ""
						activeSpecialMAP[petSkeletonStartBlock+94] = ""
						activeSpecialMAP[petSkeletonStartBlock+95] = ""
						petSkeletonStartBlock += 94
						activeSpecialMAP[petSkeletonStartBlock] = "petSkeletonTL"
						activeSpecialMAP[petSkeletonStartBlock+1] = "petSkeleton"
						activeSpecialMAP[petSkeletonStartBlock+94] = "petSkeleton"
						activeSpecialMAP[petSkeletonStartBlock+95] = "petSkeleton"

					}
				}

			case "watermelon":

				if frameCount%3 == 0 {

					for a := 5263; a >= 0; a-- {

						checkWaterMelonBlock := activeSpecialMAP[a]

						if checkWaterMelonBlock == "watermelonTL" {

							if levelMAP[a+(94*2)] != "floor" && levelMAP[a+((94*2)+1)] != "floor" {

								activeSpecialMAP[a] = ""
								activeSpecialMAP[a+1] = ""
								activeSpecialMAP[a+94] = ""
								activeSpecialMAP[a+95] = ""

								activeSpecialMAP[a+94] = "watermelonTL"
								activeSpecialMAP[a+95] = "watermelon"
								activeSpecialMAP[a+(94*2)] = "watermelon"
								activeSpecialMAP[a+((94*2)+1)] = "watermelon"

								if enemiesMAP[a+94] == "enemy" || enemiesMAP[a+94] == "enemyTL" {

									if enemiesMAP[a+94] == "enemyTL" {
										clearEnemyBlock = findEnemyTL(a + 94)
										clearEnemy()
										killEnemy()
									} else if enemiesMAP[a+94] == "enemy" {
										if enemiesMAP[a] == "boundary" {
											clearEnemyBlock = findEnemyTL(a + 93)
											clearEnemy()
											killEnemy()
										} else if enemiesMAP[a] == "enemy" {
											clearEnemyBlock = findEnemyTL(a - 1)
											clearEnemy()
											killEnemy()
										} else if enemiesMAP[a] == "enemyTL" {
											clearEnemyBlock = findEnemyTL(a)
											clearEnemy()
											killEnemy()
										}
									}
								}
								if enemiesMAP[a+95] == "enemy" || enemiesMAP[a+95] == "enemyTL" {

									if enemiesMAP[a+95] == "enemyTL" {
										clearEnemyBlock = findEnemyTL(a + 95)
										clearEnemy()
										killEnemy()
									} else if enemiesMAP[a+95] == "enemy" {
										if enemiesMAP[a+1] == "boundary" {
											clearEnemyBlock = findEnemyTL(a + 94)
											clearEnemy()
											killEnemy()
										} else if enemiesMAP[a+1] == "enemy" {
											clearEnemyBlock = findEnemyTL(a)
											clearEnemy()
											killEnemy()
										} else if enemiesMAP[a+1] == "enemyTL" {
											clearEnemyBlock = findEnemyTL(a + 1)
											clearEnemy()
											killEnemy()
										}
									}

								}
								if enemiesMAP[a+(94*2)] == "enemy" || enemiesMAP[a+(94*2)] == "enemyTL" {

									if enemiesMAP[a+(94*2)] == "enemyTL" {
										clearEnemyBlock = findEnemyTL(a + (94 * 2))
										clearEnemy()
										killEnemy()
									} else if enemiesMAP[a+(94*2)] == "enemy" {
										if enemiesMAP[a+94] == "boundary" {
											clearEnemyBlock = findEnemyTL((a + (94 * 2)) - 1)
											clearEnemy()
											killEnemy()
										} else if enemiesMAP[a+94] == "enemy" {
											clearEnemyBlock = findEnemyTL(a + 93)
											clearEnemy()
											killEnemy()
										} else if enemiesMAP[a+95] == "enemyTL" {
											clearEnemyBlock = findEnemyTL(a + 95)
											clearEnemy()
											killEnemy()
										}
									}

								}
								if enemiesMAP[a+((94*2)+1)] == "enemy" || enemiesMAP[a+((94*2)+1)] == "enemyTL" {

									if enemiesMAP[a+((94*2)+1)] == "enemyTL" {
										clearEnemyBlock = findEnemyTL(a + ((94 * 2) + 1))
										clearEnemy()
										killEnemy()
									} else if enemiesMAP[a+((94*2)+1)] == "enemy" {
										if enemiesMAP[a+95] == "boundary" {
											clearEnemyBlock = findEnemyTL(a + (94 * 2))
											clearEnemy()
											killEnemy()
										} else if enemiesMAP[a+95] == "enemy" {
											clearEnemyBlock = findEnemyTL(a + 94)
											clearEnemy()
											killEnemy()
										} else if enemiesMAP[a+95] == "enemyTL" {
											clearEnemyBlock = findEnemyTL(a + 95)
											clearEnemy()
											killEnemy()
										}
									}

								}

							}
						}

					}

				}

			case "petGreenPig":

				chooseCoins := rInt(1, 4)

				if activeSpecialComplete == false {
					if chooseCoins == 1 {
						pigCoinTimer1 = rInt(10, 40)
					} else if chooseCoins == 2 {
						pigCoinTimer1 = rInt(10, 20)
						pigCoinTimer2 = rInt(20, 40)
					} else if chooseCoins == 3 {
						pigCoinTimer1 = rInt(10, 20)
						pigCoinTimer2 = rInt(20, 30)
						pigCoinTimer3 = rInt(30, 40)
					}
					activeSpecialComplete = true
				}

				if pigCoinTimer1 == secondsCountdown && pigCoinTimer1 != 0 {
					if pigCoinHasDropped == false {
						dropCoin = true
						pigCoinHasDropped = true
						pigShieldOn = true
					}
				}
				if pigCoinHasDropped == true && pigCoinTimer1-1 == secondsCountdown {
					pigCoinHasDropped = false
					pigShieldOn = false
				}
				if pigCoinTimer2 == secondsCountdown && pigCoinTimer2 != 0 {
					if pigCoinHasDropped == false {
						dropCoin = true
						pigCoinHasDropped = true
						pigShieldOn = true
					}
				}
				if pigCoinHasDropped == true && pigCoinTimer2-1 == secondsCountdown {
					pigCoinHasDropped = false
					pigShieldOn = false
				}
				if pigCoinTimer3 == secondsCountdown && pigCoinTimer3 != 0 {
					if pigCoinHasDropped == false {
						dropCoin = true
						pigCoinHasDropped = true
						pigShieldOn = true
					}
				}
				if pigCoinHasDropped == true && pigCoinTimer3-1 == secondsCountdown {
					pigCoinHasDropped = false
					pigShieldOn = false
				}

				if frameCount%6 == 0 {
					greenPigHorizontal := petGreenPigStartBlock / 94
					greenPigVertical := petGreenPigStartBlock - (greenPigHorizontal * 94)

					if petGreenPigLeftRight {
						if greenPigVertical < 10 {
							petGreenPigLeftRight = false
						}
						activeSpecialMAP[petGreenPigStartBlock] = ""
						activeSpecialMAP[petGreenPigStartBlock+1] = ""
						activeSpecialMAP[petGreenPigStartBlock+94] = ""
						activeSpecialMAP[petGreenPigStartBlock+95] = ""
						petGreenPigStartBlock--
						activeSpecialMAP[petGreenPigStartBlock] = "petGreenPigTL"
						activeSpecialMAP[petGreenPigStartBlock+1] = "petGreenPig"
						activeSpecialMAP[petGreenPigStartBlock+94] = "petGreenPig"
						activeSpecialMAP[petGreenPigStartBlock+95] = "petGreenPig"
					} else {
						if greenPigVertical > 84 {
							petGreenPigLeftRight = true
						}
						activeSpecialMAP[petGreenPigStartBlock] = ""
						activeSpecialMAP[petGreenPigStartBlock+1] = ""
						activeSpecialMAP[petGreenPigStartBlock+94] = ""
						activeSpecialMAP[petGreenPigStartBlock+95] = ""
						petGreenPigStartBlock++
						activeSpecialMAP[petGreenPigStartBlock] = "petGreenPigTL"
						activeSpecialMAP[petGreenPigStartBlock+1] = "petGreenPig"
						activeSpecialMAP[petGreenPigStartBlock+94] = "petGreenPig"
						activeSpecialMAP[petGreenPigStartBlock+95] = "petGreenPig"
					}
				}

				petGreenPigGroundBlockL := levelMAP[petGreenPigStartBlock+(94*2)]
				petGreenPigGroundBlockR := levelMAP[petGreenPigStartBlock+((94*2)+1)]

				if petGreenPigGroundBlockL != "floor" && petGreenPigGroundBlockR != "floor" {
					activeSpecialMAP[petGreenPigStartBlock] = ""
					activeSpecialMAP[petGreenPigStartBlock+1] = ""
					activeSpecialMAP[petGreenPigStartBlock+94] = ""
					activeSpecialMAP[petGreenPigStartBlock+95] = ""
					petGreenPigStartBlock += 94
					activeSpecialMAP[petGreenPigStartBlock] = "petGreenPigTL"
					activeSpecialMAP[petGreenPigStartBlock+1] = "petGreenPig"
					activeSpecialMAP[petGreenPigStartBlock+94] = "petGreenPig"
					activeSpecialMAP[petGreenPigStartBlock+95] = "petGreenPig"

				}

			case "petMushroom":

				if frameCount%15 == 0 {
					if mushroomShieldOn {
						mushroomShieldOn = false
					} else {
						mushroomShieldOn = true
					}
				}

				mushroomBomb := rolldice() + rolldice()

				if mushroomBomb == 12 {
					mushroomBomb := rolldice()

					if mushroomBomb == 6 {
						bombBlock = petMushroomStartBlock

						BombLength := rInt(10, 15)
						bombSize := rInt(6, 11)
						bombBlock -= BombLength / 2
						bombBlock -= 94 * (bombSize - 2)

						for a := 0; a < bombSize; a++ {

							for b := 0; b < BombLength; b++ {
								effectsMAP[bombBlock] = "bomb4"
								backgroundObjectsMAP[bombBlock] = ""

								if bombBlock < 4136 {
									if levelMAP[bombBlock] == "floor" {

										for c := 0; c < 43; c++ {
											platformType := rInt(1, 44)
											platformTypeName = "platform" + strconv.Itoa(platformType)
											platformTypeNameTL = platformTypeName + "TL"

											if platformsMAP[bombBlock] == platformTypeNameTL {

												levelMAP[bombBlock] = ""
												levelMAP[bombBlock+1] = ""
												levelMAP[bombBlock+94] = ""
												levelMAP[bombBlock+95] = ""

												platformsMAP[bombBlock] = ""
												platformsMAP[bombBlock+1] = ""
												platformsMAP[bombBlock+94] = ""
												platformsMAP[bombBlock+95] = ""

												backgroundObjectsMAP[bombBlock] = ""
												backgroundObjectsMAP[bombBlock+1] = ""
												backgroundObjectsMAP[bombBlock+94] = ""
												backgroundObjectsMAP[bombBlock+95] = ""

											} else if platformsMAP[bombBlock] == platformTypeName {

												if platformsMAP[bombBlock-1] == platformTypeNameTL {

													newBombBlock := bombBlock - 1

													levelMAP[newBombBlock] = ""
													levelMAP[newBombBlock+1] = ""
													levelMAP[newBombBlock+94] = ""
													levelMAP[newBombBlock+95] = ""

													platformsMAP[newBombBlock] = ""
													platformsMAP[newBombBlock+1] = ""
													platformsMAP[newBombBlock+94] = ""
													platformsMAP[newBombBlock+95] = ""

													backgroundObjectsMAP[newBombBlock] = ""
													backgroundObjectsMAP[newBombBlock+1] = ""
													backgroundObjectsMAP[newBombBlock+94] = ""
													backgroundObjectsMAP[newBombBlock+95] = ""

												} else if platformsMAP[bombBlock-95] == platformTypeNameTL {

													newBombBlock := bombBlock - 95

													levelMAP[newBombBlock] = ""
													levelMAP[newBombBlock+1] = ""
													levelMAP[newBombBlock+94] = ""
													levelMAP[newBombBlock+95] = ""

													platformsMAP[newBombBlock] = ""
													platformsMAP[newBombBlock+1] = ""
													platformsMAP[newBombBlock+94] = ""
													platformsMAP[newBombBlock+95] = ""

													backgroundObjectsMAP[newBombBlock] = ""
													backgroundObjectsMAP[newBombBlock+1] = ""
													backgroundObjectsMAP[newBombBlock+94] = ""
													backgroundObjectsMAP[newBombBlock+95] = ""

												} else if platformsMAP[bombBlock-94] == platformTypeNameTL {

													newBombBlock := bombBlock - 94

													levelMAP[newBombBlock] = ""
													levelMAP[newBombBlock+1] = ""
													levelMAP[newBombBlock+94] = ""
													levelMAP[newBombBlock+95] = ""

													platformsMAP[newBombBlock] = ""
													platformsMAP[newBombBlock+1] = ""
													platformsMAP[newBombBlock+94] = ""
													platformsMAP[newBombBlock+95] = ""

													backgroundObjectsMAP[newBombBlock] = ""
													backgroundObjectsMAP[newBombBlock+1] = ""
													backgroundObjectsMAP[newBombBlock+94] = ""
													backgroundObjectsMAP[newBombBlock+95] = ""
												}
											}
										}

									}

								}

								bombBlock++
							}
							bombBlock += 94
							bombBlock += rInt(-4, 5)
							BombLength = rInt(10, 15)

							for b := BombLength; b > 0; b-- {
								effectsMAP[bombBlock] = "bomb4"
								bombBlock--
							}
							bombBlock += 94
							bombBlock += rInt(-4, 5)
							BombLength = rInt(10, 15)
						}

					}
				}

				if frameCount%2 == 0 {
					if mushroomJumpOn == false {
						petMushroomGroundBlockL := levelMAP[petMushroomStartBlock+(94*2)]
						petMushroomGroundBlockR := levelMAP[petMushroomStartBlock+((94*2)+1)]

						if petMushroomGroundBlockL != "floor" && petMushroomGroundBlockR != "floor" {
							activeSpecialMAP[petMushroomStartBlock] = ""
							activeSpecialMAP[petMushroomStartBlock+1] = ""
							activeSpecialMAP[petMushroomStartBlock+94] = ""
							activeSpecialMAP[petMushroomStartBlock+95] = ""
							petMushroomStartBlock += 94
							activeSpecialMAP[petMushroomStartBlock] = "petMushroomTL"
							activeSpecialMAP[petMushroomStartBlock+1] = "petMushroom"
							activeSpecialMAP[petMushroomStartBlock+94] = "petMushroom"
							activeSpecialMAP[petMushroomStartBlock+95] = "petMushroom"

						}
					}
				}

				if frameCount%6 == 0 {

					petMushroomHorizontal := petMushroomStartBlock / 94
					petMushroomVertical := petMushroomStartBlock - (petMushroomHorizontal * 94)

					if petMushroomHorizontal < 8 {
						activeSpecialMAP[petMushroomStartBlock] = ""
						activeSpecialMAP[petMushroomStartBlock+1] = ""
						activeSpecialMAP[petMushroomStartBlock+94] = ""
						activeSpecialMAP[petMushroomStartBlock+95] = ""
						petMushroomStartBlock += 94
						activeSpecialMAP[petMushroomStartBlock] = "petMushroomTL"
						activeSpecialMAP[petMushroomStartBlock+1] = "petMushroom"
						activeSpecialMAP[petMushroomStartBlock+94] = "petMushroom"
						activeSpecialMAP[petMushroomStartBlock+95] = "petMushroom"
					}

					if petMushroomLeftRight {
						if petMushroomVertical < 10 {
							petMushroomLeftRight = false
						}
						activeSpecialMAP[petMushroomStartBlock] = ""
						activeSpecialMAP[petMushroomStartBlock+1] = ""
						activeSpecialMAP[petMushroomStartBlock+94] = ""
						activeSpecialMAP[petMushroomStartBlock+95] = ""
						petMushroomStartBlock--
						activeSpecialMAP[petMushroomStartBlock] = "petMushroomTL"
						activeSpecialMAP[petMushroomStartBlock+1] = "petMushroom"
						activeSpecialMAP[petMushroomStartBlock+94] = "petMushroom"
						activeSpecialMAP[petMushroomStartBlock+95] = "petMushroom"
					} else {
						if petMushroomVertical > 84 {
							petMushroomLeftRight = true
						}
						activeSpecialMAP[petMushroomStartBlock] = ""
						activeSpecialMAP[petMushroomStartBlock+1] = ""
						activeSpecialMAP[petMushroomStartBlock+94] = ""
						activeSpecialMAP[petMushroomStartBlock+95] = ""
						petMushroomStartBlock++
						activeSpecialMAP[petMushroomStartBlock] = "petMushroomTL"
						activeSpecialMAP[petMushroomStartBlock+1] = "petMushroom"
						activeSpecialMAP[petMushroomStartBlock+94] = "petMushroom"
						activeSpecialMAP[petMushroomStartBlock+95] = "petMushroom"
					}

					if mushroomJumpOn == false {
						mushroomHorizontal := petMushroomStartBlock / 94
						if mushroomHorizontal > 14 {
							mushroomJump := rolldice() + rolldice()
							if mushroomJump > 10 {
								mushroomJumpOn = true
								mushroomJumpHeight = rInt(5, 10)
							}
						}
					}

					if mushroomJumpOn {
						for {

							mushroomHorizontal := petMushroomStartBlock / 94
							if mushroomHorizontal > 8 {

								activeSpecialMAP[petMushroomStartBlock] = ""
								activeSpecialMAP[petMushroomStartBlock+1] = ""
								activeSpecialMAP[petMushroomStartBlock+94] = ""
								activeSpecialMAP[petMushroomStartBlock+95] = ""
								petMushroomStartBlock -= 94
								activeSpecialMAP[petMushroomStartBlock] = "petMushroomTL"
								activeSpecialMAP[petMushroomStartBlock+1] = "petMushroom"
								activeSpecialMAP[petMushroomStartBlock+94] = "petMushroom"
								activeSpecialMAP[petMushroomStartBlock+95] = "petMushroom"

								mushroomJumpHeight--

								if mushroomJumpHeight == 0 {
									mushroomJumpOn = false
									break
								}
							} else {
								mushroomJumpHeight = 0
								mushroomJumpOn = false
							}

						}

					}

				}

			case "cannon":
				if frameCount%60 == 0 {
					cannonTimer--
				}
				if cannonTimer == 0 {
					activeSpecialItem = ""
					activeSpecialOn = false
					activeSpecialActive = false
					activeSpecialMAP[cannonStartBlock] = ""
				}

				if frameCount%9 == 0 {
					activeSpecialMAP[cannonStartBlock+8] = "cannonball"
					bulletsMAP[cannonStartBlock+8] = "bulletR"
					bulletsMAP[cannonStartBlock+102] = "bulletR"
				}
			case "iceball":
				if frameCount%3 == 0 {
					if iceballStartBlock < 4230 {
						activeSpecialMAP[iceballStartBlock] = ""
						iceballStartBlock++
						activeSpecialMAP[iceballStartBlock] = "iceballTL"
					} else {
						activeSpecialMAP[iceballStartBlock] = ""
					}
				}
			case "poisonball":
				if frameCount%3 == 0 {
					if fireballStartBlock < 4230 {
						activeSpecialMAP[poisonballStartBlock] = ""
						poisonballStartBlock++
						activeSpecialMAP[poisonballStartBlock] = "poisonballTL"
					} else {
						activeSpecialMAP[poisonballStartBlock] = ""
					}
				}
			case "fireball":
				if frameCount%3 == 0 {
					if fireballStartBlock > 188 {
						activeSpecialMAP[fireballStartBlock] = ""
						activeSpecialMAP[fireballStartBlock+94] = ""
						fireballStartBlock--
						activeSpecialMAP[fireballStartBlock] = "fireballTL"
						activeSpecialMAP[fireballStartBlock+94] = "fireball"
					} else {
						activeSpecialMAP[fireballStartBlock] = ""
						activeSpecialMAP[fireballStartBlock+94] = ""
					}
				}
			case "petPigeon":

				if frameCount%15 == 0 {
					if pigeonShieldFlash1 {
						pigeonShieldFlash2 = true
						pigeonShieldFlash1 = false
					} else if pigeonShieldFlash2 {
						pigeonShieldFlash3 = true
						pigeonShieldFlash2 = false

					} else if pigeonShieldFlash3 {
						pigeonShieldFlash1 = true
						pigeonShieldFlash3 = false
					}

					pigeonShoot := rolldice()

					if pigeonShoot == 6 {
						activeSpecialMAP[pigeonStartBlock+94*4] = "zap10"
						activeSpecialMAP[pigeonStartBlock+94*6] = "zap10"
					}

				}

				if frameCount%9 == 0 {

					if pigeonStartBlock > 4044 {
						activeSpecialMAP[pigeonStartBlock] = ""
						activeSpecialMAP[pigeonStartBlock+1] = ""
						activeSpecialMAP[pigeonStartBlock+94] = ""
						activeSpecialMAP[pigeonStartBlock+95] = ""
						pigeonStartBlock -= 94
						activeSpecialMAP[pigeonStartBlock] = "pigeonTL"
						activeSpecialMAP[pigeonStartBlock+1] = "pigeon"
						activeSpecialMAP[pigeonStartBlock+94] = "pigeon"
						activeSpecialMAP[pigeonStartBlock+95] = "pigeon"
					}

					activeSpecialMAP[pigeonStartBlock] = ""
					activeSpecialMAP[pigeonStartBlock+1] = ""
					activeSpecialMAP[pigeonStartBlock+94] = ""
					activeSpecialMAP[pigeonStartBlock+95] = ""
					pigeonStartBlock--

					if pigeonStartBlock < 184 {
						pigeonStartBlock = 276
					}
					activeSpecialMAP[pigeonStartBlock] = "pigeonTL"
					activeSpecialMAP[pigeonStartBlock+1] = "pigeon"
					activeSpecialMAP[pigeonStartBlock+94] = "pigeon"
					activeSpecialMAP[pigeonStartBlock+95] = "pigeon"

				}
			case "sawblade":
				if frameCount%3 == 0 {
					checkEnemyBlock1 := enemiesMAP[sawbladeStartBlock]
					checkEnemyBlock2 := enemiesMAP[sawbladeStartBlock+1]
					checkEnemyBlock3 := enemiesMAP[sawbladeStartBlock+94]
					checkEnemyBlock4 := enemiesMAP[sawbladeStartBlock+95]

					if checkEnemyBlock1 != "" && checkEnemyBlock1 != "boundary" {
						clearEnemyBlock = findEnemyTL(sawbladeStartBlock)
						clearEnemy()
						killEnemy()

					} else if checkEnemyBlock2 != "" && checkEnemyBlock1 != "boundary" {
						clearEnemyBlock = findEnemyTL(sawbladeStartBlock + 1)
						clearEnemy()
						killEnemy()

					} else if checkEnemyBlock3 != "" && checkEnemyBlock1 != "boundary" {
						clearEnemyBlock = findEnemyTL(sawbladeStartBlock + 94)
						clearEnemy()
						killEnemy()
					} else if checkEnemyBlock4 != "" && checkEnemyBlock1 != "boundary" {

						clearEnemyBlock = findEnemyTL(sawbladeStartBlock + 95)
						clearEnemy()
						killEnemy()
					}

					sawbladeHorizontal := sawbladeStartBlock / 94
					sawbladeVertical := sawbladeStartBlock - (sawbladeHorizontal * 94)

					if sawbladeLR == false {
						if sawbladeVertical < 84 {
							activeSpecialMAP[sawbladeStartBlock] = ""
							activeSpecialMAP[sawbladeStartBlock+1] = ""
							activeSpecialMAP[sawbladeStartBlock+94] = ""
							activeSpecialMAP[sawbladeStartBlock+95] = ""
							sawbladeStartBlock++
							activeSpecialMAP[sawbladeStartBlock] = "sawbladeTL"
							activeSpecialMAP[sawbladeStartBlock+1] = "sawblade"
							activeSpecialMAP[sawbladeStartBlock+94] = "sawblade"
							activeSpecialMAP[sawbladeStartBlock+95] = "sawblade"
						} else if sawbladeVertical >= 84 {
							if sawbladeHorizontal > 30 {
								activeSpecialMAP[sawbladeStartBlock] = ""
								activeSpecialMAP[sawbladeStartBlock+1] = ""
								activeSpecialMAP[sawbladeStartBlock+94] = ""
								activeSpecialMAP[sawbladeStartBlock+95] = ""

								sawbladeChange := rInt(2, 5)
								sawbladeStartBlock -= (94 * sawbladeChange)

								activeSpecialMAP[sawbladeStartBlock] = "sawbladeTL"
								activeSpecialMAP[sawbladeStartBlock+1] = "sawblade"
								activeSpecialMAP[sawbladeStartBlock+94] = "sawblade"
								activeSpecialMAP[sawbladeStartBlock+95] = "sawblade"

							} else if sawbladeHorizontal < 30 && sawbladeHorizontal > 10 {
								activeSpecialMAP[sawbladeStartBlock] = ""
								activeSpecialMAP[sawbladeStartBlock+1] = ""
								activeSpecialMAP[sawbladeStartBlock+94] = ""
								activeSpecialMAP[sawbladeStartBlock+95] = ""

								sawbladeUpDown := flipcoin()
								if sawbladeUpDown {
									sawbladeChange := rInt(2, 5)
									sawbladeStartBlock -= (94 * sawbladeChange)
								} else {
									sawbladeChange := rInt(2, 5)
									sawbladeStartBlock += (94 * sawbladeChange)
								}

								activeSpecialMAP[sawbladeStartBlock] = "sawbladeTL"
								activeSpecialMAP[sawbladeStartBlock+1] = "sawblade"
								activeSpecialMAP[sawbladeStartBlock+94] = "sawblade"
								activeSpecialMAP[sawbladeStartBlock+95] = "sawblade"
							} else if sawbladeVertical < 10 {
								activeSpecialMAP[sawbladeStartBlock] = ""
								activeSpecialMAP[sawbladeStartBlock+1] = ""
								activeSpecialMAP[sawbladeStartBlock+94] = ""
								activeSpecialMAP[sawbladeStartBlock+95] = ""

								sawbladeChange := rInt(2, 5)
								sawbladeStartBlock += (94 * sawbladeChange)

								activeSpecialMAP[sawbladeStartBlock] = "sawbladeTL"
								activeSpecialMAP[sawbladeStartBlock+1] = "sawblade"
								activeSpecialMAP[sawbladeStartBlock+94] = "sawblade"
								activeSpecialMAP[sawbladeStartBlock+95] = "sawblade"
							}
							sawbladeLR = true
						}
					}
					if sawbladeLR == true {
						if sawbladeVertical > 4 {
							activeSpecialMAP[sawbladeStartBlock] = ""
							activeSpecialMAP[sawbladeStartBlock+1] = ""
							activeSpecialMAP[sawbladeStartBlock+94] = ""
							activeSpecialMAP[sawbladeStartBlock+95] = ""
							sawbladeStartBlock--
							activeSpecialMAP[sawbladeStartBlock] = "sawbladeTL"
							activeSpecialMAP[sawbladeStartBlock+1] = "sawblade"
							activeSpecialMAP[sawbladeStartBlock+94] = "sawblade"
							activeSpecialMAP[sawbladeStartBlock+95] = "sawblade"
						} else if sawbladeVertical <= 4 {
							if sawbladeHorizontal > 30 {
								activeSpecialMAP[sawbladeStartBlock] = ""
								activeSpecialMAP[sawbladeStartBlock+1] = ""
								activeSpecialMAP[sawbladeStartBlock+94] = ""
								activeSpecialMAP[sawbladeStartBlock+95] = ""

								sawbladeChange := rInt(2, 5)
								sawbladeStartBlock -= (94 * sawbladeChange)

								activeSpecialMAP[sawbladeStartBlock] = "sawbladeTL"
								activeSpecialMAP[sawbladeStartBlock+1] = "sawblade"
								activeSpecialMAP[sawbladeStartBlock+94] = "sawblade"
								activeSpecialMAP[sawbladeStartBlock+95] = "sawblade"

							} else if sawbladeHorizontal < 30 && sawbladeHorizontal > 10 {
								activeSpecialMAP[sawbladeStartBlock] = ""
								activeSpecialMAP[sawbladeStartBlock+1] = ""
								activeSpecialMAP[sawbladeStartBlock+94] = ""
								activeSpecialMAP[sawbladeStartBlock+95] = ""

								sawbladeUpDown := flipcoin()
								if sawbladeUpDown {
									sawbladeChange := rInt(2, 5)
									sawbladeStartBlock -= (94 * sawbladeChange)
								} else {
									sawbladeChange := rInt(2, 5)
									sawbladeStartBlock += (94 * sawbladeChange)
								}

								activeSpecialMAP[sawbladeStartBlock] = "sawbladeTL"
								activeSpecialMAP[sawbladeStartBlock+1] = "sawblade"
								activeSpecialMAP[sawbladeStartBlock+94] = "sawblade"
								activeSpecialMAP[sawbladeStartBlock+95] = "sawblade"
							} else if sawbladeVertical < 10 {
								activeSpecialMAP[sawbladeStartBlock] = ""
								activeSpecialMAP[sawbladeStartBlock+1] = ""
								activeSpecialMAP[sawbladeStartBlock+94] = ""
								activeSpecialMAP[sawbladeStartBlock+95] = ""

								sawbladeChange := rInt(2, 5)
								sawbladeStartBlock += (94 * sawbladeChange)

								activeSpecialMAP[sawbladeStartBlock] = "sawbladeTL"
								activeSpecialMAP[sawbladeStartBlock+1] = "sawblade"
								activeSpecialMAP[sawbladeStartBlock+94] = "sawblade"
								activeSpecialMAP[sawbladeStartBlock+95] = "sawblade"
							}

							sawbladeLR = false
						}
					}

				}
			}

		}

		// MARK: active special animations timers
		switch activeSpecialItem {
		case "fireball":
			fireballIMG.X -= 68
			if fireballIMG.X == 1368 {
				fireballIMG.X = 1980
			}
			fireballLIMG.X += 68
			if fireballLIMG.X > 2040 {
				fireballLIMG.X = 1374
			}
			if frameCountGameStart%6 == 0 {
				flameIMG.X += 25
				if flameIMG.X > 1240 {
					flameIMG.X = 1045
				}
			}
		case "sawblade":
			if frameCountGameStart%2 == 0 {
				sawbladeIMG.X += 38
				if sawbladeIMG.X == 2044 {
					sawbladeIMG.X = 1740
				}
			}
		case "petPigeon":
			if frameCountGameStart%8 == 0 {
				petPigeonIMG.X += 48
				if petPigeonIMG.X == 1674 {
					petPigeonIMG.X = 1482
				}
				petPigeonUpIMG.X += 48
				if petPigeonUpIMG.X > 1452 {
					petPigeonUpIMG.X = 1264
				}
			}
		case "petMushroom":
			if frameCountGameStart%9 == 0 {
				petMushroomIMG.X += 32
				if petMushroomIMG.X == 1473 {
					petMushroomIMG.X = 1281
				}
			}
		case "petBat":
			if frameCountGameStart%6 == 0 {
				petBatIMG.X += 16
				if petBatIMG.X == 1483 {
					petBatIMG.X = 1403
				}
			}
		case "petGreenPig":
			if frameCountGameStart%4 == 0 {
				petGreenPigIMG.X += 34
				if petGreenPigIMG.X == 2018 {
					petGreenPigIMG.X = 1678
				}
			}
		case "petRedSlime":
			if frameCountGameStart%6 == 0 {
				petRedSlimeIMG.Y += 16
				if petRedSlimeIMG.Y == 215 {
					petRedSlimeIMG.Y = 87
				}
			}
		case "petKnight":
			if frameCountGameStart%6 == 0 {
				petKnightIMG.X += 16
				if petKnightIMG.X == 1509 {
					petKnightIMG.X = 1413
				}
				petKnightLIMG.X += 16
				if petKnightLIMG.X > 1260 {
					petKnightLIMG.X = 1172
				}
			}
		case "exclamation":
			if frameCountGameStart%15 == 0 {
				exclamationIMG.X += 34
				if exclamationIMG.X == 1397 {
					exclamationIMG.X = 1295
				}
			}
		case "propellor":
			if frameCountGameStart%5 == 0 {
				propellorIMG.X += 24
				if propellorIMG.X == 1509 {
					propellorIMG.X = 1413
				}
			}
		case "trampoline":
			if frameCountGameStart%10 == 0 {
				trampolineIMG.X += 28
				if trampolineIMG.X == 1507 {
					trampolineIMG.X = 1283
				}
			}
		case "greenHeart":
			if frameCountGameStart%10 == 0 {
				greenHeartIMG.X += 16
				if greenHeartIMG.X == 1527 {
					greenHeartIMG.X = 1399
				}
			}
		case "petSkeleton":
			if frameCountGameStart%10 == 0 {
				petSkeletonIMG.X += 16
				if petSkeletonIMG.X == 1582 {
					petSkeletonIMG.X = 1518
				}
				petSkeletonLIMG.X += 16
				if petSkeletonLIMG.X > 1308 {
					petSkeletonLIMG.X = 1252
				}
			}
		case "petSlime":
			if frameCountGameStart%7 == 0 {
				petSlimeIMG.X += 16
				if petSlimeIMG.X == 1503 {
					petSlimeIMG.X = 1407
				}
			}
		case "cannon":
			if frameCountGameStart%10 == 0 {
				cannonIMG.X -= 44
				if cannonIMG.X == 1531 {
					cannonIMG.X = 1663
				}
			}
		case "watermelon":
			if frameCountGameStart%3 == 0 {
				watermelonIMG.X += 32
				if watermelonIMG.X == 2048 {
					watermelonIMG.X = 1504
				}
			}
		case "poisonball":
			poisonballIMG.X -= 65
			if poisonballIMG.X == 1400 {
				poisonballIMG.X = 1985
			}
			if frameCountGameStart%6 == 0 {
				poisonFlameIMG.X += 25
				if poisonFlameIMG.X > 1240 {
					poisonFlameIMG.X = 1045
				}
			}
		case "iceball":
			iceballIMG.X -= 84
			if iceballIMG.X == 1230 {
				iceballIMG.X = 1986
			}
		case "apple":
			if frameCountGameStart%3 == 0 {
				appleIMG.X += 32
				if appleIMG.X == 2048 {
					appleIMG.X = 1504
				}
			}
		case "cherries":
			if frameCountGameStart%3 == 0 {
				cherriesIMG.X += 32
				if cherriesIMG.X == 2048 {
					cherriesIMG.X = 1504
				}
			}
		case "heart":
			if frameCountGameStart%6 == 0 {
				heartShopIMG.X += 18
				if heartShopIMG.X == 1739 {
					heartShopIMG.X = 1595
				}
			}

		}

		// MARK: shop
		if shopOn {
			pauseOn = true

			if rl.IsKeyPressed(rl.KeyF3) {
				shopMenuCount = 1
				for {

					chooseShopItem1 := rInt(0, len(shopItemsMAP))
					chooseShopItem2 := rInt(0, len(shopItemsMAP))
					chooseShopItem3 := rInt(0, len(shopItemsMAP))

					if chooseShopItem1 != chooseShopItem2 && chooseShopItem2 != chooseShopItem3 && chooseShopItem1 != chooseShopItem3 {

						shopItem1 = shopItemsMAP[chooseShopItem1]
						shopItem2 = shopItemsMAP[chooseShopItem2]
						shopItem3 = shopItemsMAP[chooseShopItem3]
						break
					}
				}

				if playerCoins == 1 {
					shopPrice1 = 1
					shopPrice2 = 1
					shopPrice3 = 1
				} else if playerCoins < 5 {
					shopPrice1 = rInt(2, 5)
					shopPrice2 = rInt(2, 5)
					shopPrice3 = rInt(2, 5)
				} else if playerCoins < 10 && playerCoins >= 5 {
					shopPrice1 = rInt(5, 10)
					shopPrice2 = rInt(5, 10)
					shopPrice3 = rInt(5, 10)
				} else if playerCoins >= 10 && playerCoins < 20 {
					shopPrice1 = rInt(10, 25)
					shopPrice2 = rInt(10, 25)
					shopPrice3 = rInt(10, 25)
				} else if playerCoins >= 20 && playerCoins < 40 {
					shopPrice1 = rInt(20, 50)
					shopPrice2 = rInt(20, 50)
					shopPrice3 = rInt(20, 50)
				} else if playerCoins > 40 {
					shopPrice1 = rInt(40, 200)
					shopPrice2 = rInt(40, 200)
					shopPrice3 = rInt(40, 200)
				}
			}

			if rl.IsKeyPressed(rl.KeyLeft) {
				shopMenuCount--
				if shopMenuCount == 0 {
					shopMenuCount = 3
				}
			}
			if rl.IsKeyPressed(rl.KeyRight) {
				shopMenuCount++
				if shopMenuCount == 4 {
					shopMenuCount = 1
				}
			}

			if shuffleDone == false && playerCoins > 1 {
				if rl.IsKeyPressed(rl.KeyUp) {
					shopMenuCount = 5
				}
				if rl.IsKeyPressed(rl.KeyDown) {
					shopMenuCount = 1
				}
			}

			if rl.IsKeyPressed(rl.KeySpace) {

				switch shopMenuCount {

				case 1:
					if playerCoins-shopPrice1 >= 0 {
						activeSpecialItem = shopItem1
						activeSpecialOn = true
						shopOn = false
						playerCoins -= shopPrice1
						startNewLevel()
					}
				case 2:
					if playerCoins-shopPrice2 >= 0 {
						activeSpecialItem = shopItem2
						activeSpecialOn = true
						shopOn = false
						playerCoins -= shopPrice2
						startNewLevel()
					}
				case 3:
					if playerCoins-shopPrice3 >= 0 {
						activeSpecialItem = shopItem3
						activeSpecialOn = true
						shopOn = false
						playerCoins -= shopPrice3
						startNewLevel()
					}
				case 5:
					playerCoins--
					for {

						chooseShopItem1 := rInt(0, len(shopItemsMAP))
						chooseShopItem2 := rInt(0, len(shopItemsMAP))
						chooseShopItem3 := rInt(0, len(shopItemsMAP))

						if chooseShopItem1 != chooseShopItem2 && chooseShopItem2 != chooseShopItem3 && chooseShopItem1 != chooseShopItem3 {

							shopItem1 = shopItemsMAP[chooseShopItem1]
							shopItem2 = shopItemsMAP[chooseShopItem2]
							shopItem3 = shopItemsMAP[chooseShopItem3]
							break
						}
					}
					if playerCoins == 1 {
						shopPrice1 = 1
						shopPrice2 = 1
						shopPrice3 = 1
					} else if playerCoins < 5 {
						shopPrice1 = rInt(2, 5)
						shopPrice2 = rInt(2, 5)
						shopPrice3 = rInt(2, 5)
					} else if playerCoins < 10 && playerCoins >= 5 {
						shopPrice1 = rInt(5, 10)
						shopPrice2 = rInt(5, 10)
						shopPrice3 = rInt(5, 10)
					} else if playerCoins >= 10 && playerCoins < 20 {
						shopPrice1 = rInt(10, 25)
						shopPrice2 = rInt(10, 25)
						shopPrice3 = rInt(10, 25)
					} else if playerCoins >= 20 && playerCoins < 40 {
						shopPrice1 = rInt(20, 50)
						shopPrice2 = rInt(20, 50)
						shopPrice3 = rInt(20, 50)
					} else if playerCoins > 40 {
						shopPrice1 = rInt(40, 200)
						shopPrice2 = rInt(40, 200)
						shopPrice3 = rInt(40, 200)
					}

					shopMenuCount = 1
					shuffleDone = true
				}
			}
			// black background
			rl.DrawRectangle(0, 0, screenW, screenH, rl.Black)

			// draw enemies
			if frameCountGameStart%6 == 0 {
				enemy1IMG.X += 32
				if enemy1IMG.X >= 1024 {
					enemy1IMG.X = 514
				}
			}
			if frameCountGameStart%3 == 0 {
				enemy7UPIMG.X += 44
				if enemy7UPIMG.X >= 1560 {
					enemy7UPIMG.X = 1218
				}
			}
			rl.BeginMode2D(camera)

			enemyshopv2.X--
			if enemyshopv2.X <= -64 {
				enemyshopv2.X = 700
			}
			enemyshop2v2.X++
			if enemyshop2v2.X >= 700 {
				enemyshop2v2.X = -64
			}
			rl.DrawTextureRec(imgs, enemy1IMG, enemyshopv2, rl.Fade(rl.White, 0.4))
			rl.DrawTextureRec(imgs, enemy7UPIMG, enemyshop2v2, rl.Fade(rl.White, 0.4))
			rl.EndMode2D()

			// moving rectangles
			rl.DrawRectangle(0, backShopRecY1, screenW, backShopRecH1, rl.Fade(rl.Maroon, 0.1))
			rl.DrawRectangle(0, backShopRecY1, screenW, backShopRecH1, rl.Fade(rl.Black, 0.2))

			rl.DrawRectangle(0, backShopRecY2, screenW, backShopRecH2, rl.Fade(rl.Maroon, 0.1))
			rl.DrawRectangle(0, backShopRecY2, screenW, backShopRecH2, rl.Fade(rl.Black, 0.2))

			rl.DrawRectangle(0, backShopRecY3, screenW, backShopRecH3, rl.Fade(rl.Maroon, 0.1))
			rl.DrawRectangle(0, backShopRecY3, screenW, backShopRecH3, rl.Fade(rl.Black, 0.2))

			rl.DrawRectangle(0, backShopRecY4, screenW, backShopRecH4, rl.Fade(rl.Maroon, 0.1))
			rl.DrawRectangle(0, backShopRecY4, screenW, backShopRecH4, rl.Fade(rl.Black, 0.2))

			if backShopRec1UpDown {
				backShopRecY1++
			} else {
				backShopRecY1--
			}
			if backShopRec2UpDown {
				backShopRecY2++
			} else {
				backShopRecY2--
			}
			if backShopRec3UpDown {
				backShopRecY3++
			} else {
				backShopRecY3--
			}
			if backShopRec4UpDown {
				backShopRecY4++
			} else {
				backShopRecY4--
			}

			if backShopRecY1 <= -100 {
				backShopRec1UpDown = true
			} else if backShopRecY1 >= screenH+100 {
				backShopRec1UpDown = false
			}
			if backShopRecY2 <= -100 {
				backShopRec2UpDown = true
			} else if backShopRecY2 >= screenH+100 {
				backShopRec2UpDown = false
			}
			if backShopRecY3 <= -100 {
				backShopRec3UpDown = true
			} else if backShopRecY3 >= screenH+100 {
				backShopRec3UpDown = false
			}
			if backShopRecY4 <= -100 {
				backShopRec4UpDown = true
			} else if backShopRecY4 >= screenH+100 {
				backShopRec4UpDown = false
			}

			rl.DrawText("shop", 50-3, 50+3, 80, rl.DarkGray)
			rl.DrawText("shop", 50-1, 50+1, 80, rl.Black)
			rl.DrawText("shop", 50, 50, 80, rl.White)
			rl.DrawText("you have    coin[s]", 100-2, 150+2, 40, rl.DarkGray)
			rl.DrawText("you have    coin[s]", 100-1, 150+1, 40, rl.Black)
			rl.DrawText("you have    coin[s]", 100, 150, 40, rl.White)

			playerCoinsTEXT := strconv.Itoa(playerCoins)
			if playerCoins < 10 {
				rl.DrawText(playerCoinsTEXT, 302, 156, 40, rl.DarkGray)
				rl.DrawText(playerCoinsTEXT, 304, 154, 40, rl.Black)
				rl.DrawText(playerCoinsTEXT, 305, 153, 40, menuTextSpaceColor)
			} else {
				rl.DrawText(playerCoinsTEXT, 295, 155, 40, rl.DarkGray)
				rl.DrawText(playerCoinsTEXT, 297, 154, 40, rl.Black)
				rl.DrawText(playerCoinsTEXT, 298, 153, 40, menuTextSpaceColor)

			}

			if shuffleDone == false && playerCoins > 1 {
				// reshuffle button
				rl.DrawRectangle(screenW-458, 53, 260, 80, rl.Black)
				rl.DrawRectangle(screenW-458, 53, 260, 80, rl.Fade(rl.White, 0.4))
				rl.DrawRectangle(screenW-456, 50, 260, 80, rl.Black)

				if shopMenuCount == 5 {

					if frameCountGameStart%15 == 0 {

						if flashButtons {
							flashButtons = false
						} else {
							flashButtons = true
						}
					}

					if flashButtons {
						rl.DrawRectangle(screenW-455, 49, 260, 80, rl.Gold)
					} else {
						rl.DrawRectangle(screenW-455, 49, 260, 80, rl.Maroon)
					}
				} else {
					rl.DrawRectangle(screenW-455, 49, 260, 80, rl.Gold)
				}

				rl.DrawText("reshuffle", screenW-423, 71, 40, rl.White)
				rl.DrawText("reshuffle", screenW-421, 69, 40, rl.Red)
				rl.DrawText("reshuffle", screenW-420, 68, 40, rl.Black)

				rl.DrawText("1 coin | only once", screenW-501, 158, 40, descTxtColor)
				rl.DrawText("1 coin | only once", screenW-499, 156, 40, rl.Black)
				rl.DrawText("1 coin | only once", screenW-498, 155, 40, rl.White)

				if frameCountGameStart%30 == 0 {
					if coinsTextColor == rl.Gold {
						coinsTextColor = rl.Yellow
					} else {
						coinsTextColor = rl.Gold
					}
					if shopMenuBorder == rl.Fade(rl.Red, 0.6) {
						shopMenuBorder = rl.Fade(rl.Maroon, 0.6)
					} else {
						shopMenuBorder = rl.Fade(rl.Red, 0.6)
					}
				}
			}

			// item backgrounds
			rl.DrawRectangle(175, 280, 235, 235, rl.Fade(rl.DarkPurple, 0.1))
			rl.DrawRectangle(195, 300, 195, 195, rl.Black)
			rl.DrawRectangle(195, 300, 195, 195, rl.Fade(rl.DarkGray, 0.2))

			rl.DrawRectangle(565, 280, 235, 235, rl.Fade(rl.DarkPurple, 0.1))
			rl.DrawRectangle(585, 300, 195, 195, rl.Black)
			rl.DrawRectangle(585, 300, 195, 195, rl.Fade(rl.DarkGray, 0.2))

			rl.DrawRectangle(955, 280, 235, 235, rl.Fade(rl.DarkPurple, 0.1))
			rl.DrawRectangle(975, 300, 195, 195, rl.Black)
			rl.DrawRectangle(975, 300, 195, 195, rl.Fade(rl.DarkGray, 0.2))

			// item select border
			switch shopMenuCount {
			case 1:
				rl.DrawRectangle(175, 280, 235, 20, shopMenuBorder)
				rl.DrawRectangle(175, 495, 235, 20, shopMenuBorder)
				rl.DrawRectangle(175, 300, 20, 195, shopMenuBorder)
				rl.DrawRectangle(390, 300, 20, 195, shopMenuBorder)
				rl.DrawRectangle(585, 300, 195, 195, rl.Fade(rl.Black, 0.4))
				rl.DrawRectangle(975, 300, 195, 195, rl.Fade(rl.Black, 0.4))

			case 2:
				rl.DrawRectangle(565, 280, 235, 20, shopMenuBorder)
				rl.DrawRectangle(565, 495, 235, 20, shopMenuBorder)
				rl.DrawRectangle(565, 300, 20, 195, shopMenuBorder)
				rl.DrawRectangle(780, 300, 20, 195, shopMenuBorder)
				rl.DrawRectangle(195, 300, 195, 195, rl.Fade(rl.Black, 0.4))
				rl.DrawRectangle(975, 300, 195, 195, rl.Fade(rl.Black, 0.4))
			case 3:
				rl.DrawRectangle(955, 280, 235, 20, shopMenuBorder)
				rl.DrawRectangle(955, 495, 235, 20, shopMenuBorder)
				rl.DrawRectangle(955, 300, 20, 195, shopMenuBorder)
				rl.DrawRectangle(1170, 300, 20, 195, shopMenuBorder)
				rl.DrawRectangle(195, 300, 195, 195, rl.Fade(rl.Black, 0.4))
				rl.DrawRectangle(585, 300, 195, 195, rl.Fade(rl.Black, 0.4))

			case 5:
				rl.DrawRectangle(195, 300, 195, 195, rl.Fade(rl.Black, 0.4))
				rl.DrawRectangle(585, 300, 195, 195, rl.Fade(rl.Black, 0.4))
				rl.DrawRectangle(975, 300, 195, 195, rl.Fade(rl.Black, 0.4))
			}

			// flash shop item text
			if frameCountGameStart%15 == 0 {
				if menuTextSpaceColor == rl.Yellow {
					menuTextSpaceColor = rl.Gold
				} else {
					menuTextSpaceColor = rl.Yellow
				}
				if descTxtColor == rl.Green {
					descTxtColor = rl.DarkGreen
				} else {
					descTxtColor = rl.Green
				}
			}

			// MARK: shop prices
			shopPrice1TEXT := strconv.Itoa(shopPrice1)
			shopPrice2TEXT := strconv.Itoa(shopPrice2)
			shopPrice3TEXT := strconv.Itoa(shopPrice3)

			if playerCoins < 10 {
				rl.DrawText(shopPrice1TEXT, 287, 643, 40, rl.DarkGray)
				rl.DrawText(shopPrice1TEXT, 289, 641, 40, rl.Black)
				rl.DrawText(shopPrice1TEXT, 290, 640, 40, menuTextSpaceColor)

				rl.DrawText(shopPrice2TEXT, 683, 643, 40, rl.DarkGray)
				rl.DrawText(shopPrice2TEXT, 685, 641, 40, rl.Black)
				rl.DrawText(shopPrice2TEXT, 686, 640, 40, menuTextSpaceColor)

				rl.DrawText(shopPrice3TEXT, 1075, 643, 40, rl.DarkGray)
				rl.DrawText(shopPrice3TEXT, 1077, 641, 40, rl.Black)
				rl.DrawText(shopPrice3TEXT, 1078, 640, 40, menuTextSpaceColor)
			} else {
				rl.DrawText(shopPrice1TEXT, 277, 643, 40, rl.DarkGray)
				rl.DrawText(shopPrice1TEXT, 279, 641, 40, rl.Black)
				rl.DrawText(shopPrice1TEXT, 280, 640, 40, menuTextSpaceColor)

				rl.DrawText(shopPrice2TEXT, 673, 643, 40, rl.DarkGray)
				rl.DrawText(shopPrice2TEXT, 675, 641, 40, rl.Black)
				rl.DrawText(shopPrice2TEXT, 676, 640, 40, menuTextSpaceColor)

				rl.DrawText(shopPrice3TEXT, 1065, 643, 40, rl.DarkGray)
				rl.DrawText(shopPrice3TEXT, 1067, 641, 40, rl.Black)
				rl.DrawText(shopPrice3TEXT, 1068, 640, 40, menuTextSpaceColor)

			}

			// draw coin
			rl.BeginMode2D(camera) // 2X zoom
			if playerCoins < 10 {
				coinIMGV21 := rl.NewVector2(124, 322)
				coinIMGV22 := rl.NewVector2(322, 322)
				coinIMGV23 := rl.NewVector2(519, 322)
				rl.DrawTextureRec(imgs, coinIMG, coinIMGV21, rl.White)
				rl.DrawTextureRec(imgs, coinIMG, coinIMGV22, rl.White)
				rl.DrawTextureRec(imgs, coinIMG, coinIMGV23, rl.White)
			} else {
				coinIMGV21 := rl.NewVector2(121, 322)
				coinIMGV22 := rl.NewVector2(319, 322)
				coinIMGV23 := rl.NewVector2(516, 322)
				rl.DrawTextureRec(imgs, coinIMG, coinIMGV21, rl.White)
				rl.DrawTextureRec(imgs, coinIMG, coinIMGV22, rl.White)
				rl.DrawTextureRec(imgs, coinIMG, coinIMGV23, rl.White)
			}

			if frameCountGameStart%9 == 0 {
				// update coin img
				coinIMG.X += 16
				if coinIMG.X == 414 {
					coinIMG.X = 334
				}
			}

			rl.EndMode2D() // end 2X zoom

			// MARK: shop items list
			switch shopItem1 {
			case "sawblade":

				rl.DrawText("sore", 237, 533, 40, rl.DarkGray)
				rl.DrawText("sore", 239, 531, 40, rl.Black)
				rl.DrawText("sore", 240, 530, 40, menuTextSpaceColor)

				rl.DrawText("cutz", 237, 583, 40, descTxtColor)
				rl.DrawText("cutz", 239, 581, 40, rl.Black)
				rl.DrawText("cutz", 240, 580, 40, rl.White)
			case "petPigeon":
				rl.DrawText("pijin", 247, 533, 40, rl.DarkGray)
				rl.DrawText("pijin", 249, 531, 40, rl.Black)
				rl.DrawText("pijin", 250, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 253, 583, 40, descTxtColor)
				rl.DrawText("pet", 254, 581, 40, rl.Black)
				rl.DrawText("pet", 255, 580, 40, rl.White)
			case "fireball":
				rl.DrawText("fireballz", 207, 533, 40, rl.DarkGray)
				rl.DrawText("fireballz", 209, 531, 40, rl.Black)
				rl.DrawText("fireballz", 210, 530, 40, menuTextSpaceColor)

				rl.DrawText("burns", 227, 583, 40, descTxtColor)
				rl.DrawText("burns", 229, 581, 40, rl.Black)
				rl.DrawText("burns", 230, 580, 40, rl.White)
			case "poisonball":
				rl.DrawText("poizinballz", 189, 533, 40, rl.DarkGray)
				rl.DrawText("poizinballz", 191, 531, 40, rl.Black)
				rl.DrawText("poizinballz", 192, 530, 40, menuTextSpaceColor)

				rl.DrawText("poisons", 217, 583, 40, descTxtColor)
				rl.DrawText("poisons", 219, 581, 40, rl.Black)
				rl.DrawText("poisons", 220, 580, 40, rl.White)
			case "iceball":
				rl.DrawText("eyzballz", 205, 533, 40, rl.DarkGray)
				rl.DrawText("eyzballz", 207, 531, 40, rl.Black)
				rl.DrawText("eyzballz", 208, 530, 40, menuTextSpaceColor)

				rl.DrawText("freezes", 205, 583, 40, descTxtColor)
				rl.DrawText("freezes", 207, 581, 40, rl.Black)
				rl.DrawText("freezes", 208, 580, 40, rl.White)
			case "cannon":
				rl.DrawText("cannin", 227, 533, 40, rl.DarkGray)
				rl.DrawText("cannin", 229, 531, 40, rl.Black)
				rl.DrawText("cannin", 230, 530, 40, menuTextSpaceColor)

				rl.DrawText("blasts", 225, 583, 40, descTxtColor)
				rl.DrawText("blasts", 227, 581, 40, rl.Black)
				rl.DrawText("blasts", 228, 580, 40, rl.White)
			case "trampoline":
				rl.DrawText("trampoleen", 179, 533, 40, rl.DarkGray)
				rl.DrawText("trampoleen", 181, 531, 40, rl.Black)
				rl.DrawText("trampoleen", 182, 530, 40, menuTextSpaceColor)

				rl.DrawText("bounce", 217, 583, 40, descTxtColor)
				rl.DrawText("bounce", 219, 581, 40, rl.Black)
				rl.DrawText("bounce", 220, 580, 40, rl.White)
			case "petMushroom":
				rl.DrawText("mushrum", 207, 533, 40, rl.DarkGray)
				rl.DrawText("mushrum", 209, 531, 40, rl.Black)
				rl.DrawText("mushrum", 210, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 253, 583, 40, descTxtColor)
				rl.DrawText("pet", 254, 581, 40, rl.Black)
				rl.DrawText("pet", 255, 580, 40, rl.White)
			case "exclamation":
				rl.DrawText("#!?", 252, 533, 40, rl.DarkGray)
				rl.DrawText("#!?", 254, 531, 40, rl.Black)
				rl.DrawText("#!?", 255, 530, 40, menuTextSpaceColor)

				rl.DrawText("talk", 242, 583, 40, descTxtColor)
				rl.DrawText("talk", 244, 581, 40, rl.Black)
				rl.DrawText("talk", 245, 580, 40, rl.White)
			case "propellor":
				rl.DrawText("propella", 206, 533, 40, rl.DarkGray)
				rl.DrawText("propella", 208, 531, 40, rl.Black)
				rl.DrawText("propella", 209, 530, 40, menuTextSpaceColor)

				rl.DrawText("higha", 237, 583, 40, descTxtColor)
				rl.DrawText("higha", 239, 581, 40, rl.Black)
				rl.DrawText("higha", 240, 580, 40, rl.White)
			case "petGreenPig":
				rl.DrawText("litl pig", 225, 533, 40, rl.DarkGray)
				rl.DrawText("litl pig", 227, 531, 40, rl.Black)
				rl.DrawText("litl pig", 228, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 253, 583, 40, descTxtColor)
				rl.DrawText("pet", 254, 581, 40, rl.Black)
				rl.DrawText("pet", 255, 580, 40, rl.White)
			case "watermelon":
				rl.DrawText("wart-a-melin", 170, 533, 40, rl.DarkGray)
				rl.DrawText("wart-a-melin", 172, 531, 40, rl.Black)
				rl.DrawText("wart-a-melin", 173, 530, 40, menuTextSpaceColor)

				rl.DrawText("rain", 247, 583, 40, descTxtColor)
				rl.DrawText("rain", 249, 581, 40, rl.Black)
				rl.DrawText("rain", 250, 580, 40, rl.White)
			case "heart":
				rl.DrawText("hart", 249, 533, 40, rl.DarkGray)
				rl.DrawText("hart", 251, 531, 40, rl.Black)
				rl.DrawText("hart", 252, 530, 40, menuTextSpaceColor)

				rl.DrawText("hp", 267, 583, 40, descTxtColor)
				rl.DrawText("hp", 269, 581, 40, rl.Black)
				rl.DrawText("hp", 270, 580, 40, rl.White)
			case "cherries":
				rl.DrawText("chairees", 202, 533, 40, rl.DarkGray)
				rl.DrawText("chairees", 204, 531, 40, rl.Black)
				rl.DrawText("chairees", 205, 530, 40, menuTextSpaceColor)

				rl.DrawText("rain", 247, 583, 40, descTxtColor)
				rl.DrawText("rain", 249, 581, 40, rl.Black)
				rl.DrawText("rain", 250, 580, 40, rl.White)
			case "petSkeleton":
				rl.DrawText("skelitin", 213, 533, 40, rl.DarkGray)
				rl.DrawText("skelitin", 215, 531, 40, rl.Black)
				rl.DrawText("skelitin", 216, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 253, 583, 40, descTxtColor)
				rl.DrawText("pet", 254, 581, 40, rl.Black)
				rl.DrawText("pet", 255, 580, 40, rl.White)
			case "petSlime":
				rl.DrawText("zlime", 241, 533, 40, rl.DarkGray)
				rl.DrawText("zlime", 243, 531, 40, rl.Black)
				rl.DrawText("zlime", 244, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 253, 583, 40, descTxtColor)
				rl.DrawText("pet", 254, 581, 40, rl.Black)
				rl.DrawText("pet", 255, 580, 40, rl.White)
			case "greenHeart":
				rl.DrawText("gween hart", 177, 533, 40, rl.DarkGray)
				rl.DrawText("gween hart", 179, 531, 40, rl.Black)
				rl.DrawText("gween hart", 180, 530, 40, menuTextSpaceColor)

				rl.DrawText("randumb", 208, 583, 40, descTxtColor)
				rl.DrawText("randumb", 209, 581, 40, rl.Black)
				rl.DrawText("randumb", 210, 580, 40, rl.White)
			case "apple":
				rl.DrawText("toemate", 209, 533, 40, rl.DarkGray)
				rl.DrawText("toemate", 211, 531, 40, rl.Black)
				rl.DrawText("toemate", 212, 530, 40, menuTextSpaceColor)

				rl.DrawText("rotin", 237, 583, 40, descTxtColor)
				rl.DrawText("rotin", 239, 581, 40, rl.Black)
				rl.DrawText("rotin", 240, 580, 40, rl.White)
			case "petKnight":
				rl.DrawText("knite", 237, 533, 40, rl.DarkGray)
				rl.DrawText("knite", 239, 531, 40, rl.Black)
				rl.DrawText("knite", 240, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 253, 583, 40, descTxtColor)
				rl.DrawText("pet", 254, 581, 40, rl.Black)
				rl.DrawText("pet", 255, 580, 40, rl.White)
			case "petRedSlime":
				rl.DrawText("fred zlime", 192, 533, 40, rl.DarkGray)
				rl.DrawText("fred zlime", 194, 531, 40, rl.Black)
				rl.DrawText("fred zlime", 195, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 253, 583, 40, descTxtColor)
				rl.DrawText("pet", 254, 581, 40, rl.Black)
				rl.DrawText("pet", 255, 580, 40, rl.White)
			case "petBat":
				rl.DrawText("bart the bat", 165, 533, 40, rl.DarkGray)
				rl.DrawText("bart the bat", 167, 531, 40, rl.Black)
				rl.DrawText("bart the bat", 168, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 253, 583, 40, descTxtColor)
				rl.DrawText("pet", 254, 581, 40, rl.Black)
				rl.DrawText("pet", 255, 580, 40, rl.White)
			}
			switch shopItem2 {
			case "sawblade":
				rl.DrawText("sore", 627, 533, 40, rl.DarkGray)
				rl.DrawText("sore", 629, 531, 40, rl.Black)
				rl.DrawText("sore", 630, 530, 40, menuTextSpaceColor)

				rl.DrawText("cutz", 630, 583, 40, descTxtColor)
				rl.DrawText("cutz", 632, 581, 40, rl.Black)
				rl.DrawText("cutz", 633, 580, 40, rl.White)
			case "petPigeon":
				rl.DrawText("pijin", 642, 533, 40, rl.DarkGray)
				rl.DrawText("pijin", 644, 531, 40, rl.Black)
				rl.DrawText("pijin", 645, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 650, 583, 40, descTxtColor)
				rl.DrawText("pet", 652, 581, 40, rl.Black)
				rl.DrawText("pet", 653, 580, 40, rl.White)
			case "fireball":
				rl.DrawText("fireballz", 597, 533, 40, rl.DarkGray)
				rl.DrawText("fireballz", 599, 531, 40, rl.Black)
				rl.DrawText("fireballz", 600, 530, 40, menuTextSpaceColor)

				rl.DrawText("burns", 628, 583, 40, descTxtColor)
				rl.DrawText("burns", 630, 581, 40, rl.Black)
				rl.DrawText("burns", 631, 580, 40, rl.White)
			case "poisonball":
				rl.DrawText("poizinballz", 577, 533, 40, rl.DarkGray)
				rl.DrawText("poizinballz", 579, 531, 40, rl.Black)
				rl.DrawText("poizinballz", 580, 530, 40, menuTextSpaceColor)

				rl.DrawText("poisins", 607, 583, 40, descTxtColor)
				rl.DrawText("poisins", 609, 581, 40, rl.Black)
				rl.DrawText("poisins", 610, 580, 40, rl.White)
			case "iceball":
				rl.DrawText("eyzballz", 599, 533, 40, rl.DarkGray)
				rl.DrawText("eyzballz", 601, 531, 40, rl.Black)
				rl.DrawText("eyzballz", 602, 530, 40, menuTextSpaceColor)

				rl.DrawText("freezes", 600, 583, 40, descTxtColor)
				rl.DrawText("freezes", 602, 581, 40, rl.Black)
				rl.DrawText("freezes", 603, 580, 40, rl.White)
			case "cannon":
				rl.DrawText("cannin", 617, 533, 40, rl.DarkGray)
				rl.DrawText("cannin", 619, 531, 40, rl.Black)
				rl.DrawText("cannin", 620, 530, 40, menuTextSpaceColor)

				rl.DrawText("blasts", 617, 583, 40, descTxtColor)
				rl.DrawText("blasts", 619, 581, 40, rl.Black)
				rl.DrawText("blasts", 620, 580, 40, rl.White)
			case "trampoline":
				rl.DrawText("trampoleen", 569, 533, 40, rl.DarkGray)
				rl.DrawText("trampoleen", 571, 531, 40, rl.Black)
				rl.DrawText("trampoleen", 572, 530, 40, menuTextSpaceColor)

				rl.DrawText("bounce", 617, 583, 40, descTxtColor)
				rl.DrawText("bounce", 619, 581, 40, rl.Black)
				rl.DrawText("bounce", 620, 580, 40, rl.White)
			case "petMushroom":
				rl.DrawText("mushrum", 595, 533, 40, rl.DarkGray)
				rl.DrawText("mushrum", 597, 531, 40, rl.Black)
				rl.DrawText("mushrum", 598, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 650, 583, 40, descTxtColor)
				rl.DrawText("pet", 652, 581, 40, rl.Black)
				rl.DrawText("pet", 653, 580, 40, rl.White)
			case "exclamation":
				rl.DrawText("#!?", 649, 533, 40, rl.DarkGray)
				rl.DrawText("#!?", 651, 531, 40, rl.Black)
				rl.DrawText("#!?", 652, 530, 40, menuTextSpaceColor)

				rl.DrawText("talk", 642, 583, 40, descTxtColor)
				rl.DrawText("talk", 644, 581, 40, rl.Black)
				rl.DrawText("talk", 645, 580, 40, rl.White)
			case "propellor":
				rl.DrawText("propella", 596, 533, 40, rl.DarkGray)
				rl.DrawText("propella", 598, 531, 40, rl.Black)
				rl.DrawText("propella", 599, 530, 40, menuTextSpaceColor)

				rl.DrawText("higha", 632, 583, 40, descTxtColor)
				rl.DrawText("higha", 634, 581, 40, rl.Black)
				rl.DrawText("higha", 635, 580, 40, rl.White)
			case "petGreenPig":
				rl.DrawText("litl pig", 617, 533, 40, rl.DarkGray)
				rl.DrawText("litl pig", 619, 531, 40, rl.Black)
				rl.DrawText("litl pig", 620, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 650, 583, 40, descTxtColor)
				rl.DrawText("pet", 652, 581, 40, rl.Black)
				rl.DrawText("pet", 653, 580, 40, rl.White)
			case "watermelon":
				rl.DrawText("wart-a-melin", 558, 533, 40, rl.DarkGray)
				rl.DrawText("wart-a-melin", 560, 531, 40, rl.Black)
				rl.DrawText("wart-a-melin", 561, 530, 40, menuTextSpaceColor)

				rl.DrawText("rain", 640, 583, 40, descTxtColor)
				rl.DrawText("rain", 641, 581, 40, rl.Black)
				rl.DrawText("rain", 642, 580, 40, rl.White)
			case "heart":
				rl.DrawText("hart", 639, 533, 40, rl.DarkGray)
				rl.DrawText("hart", 641, 531, 40, rl.Black)
				rl.DrawText("hart", 642, 530, 40, menuTextSpaceColor)

				rl.DrawText("hp", 659, 583, 40, descTxtColor)
				rl.DrawText("hp", 661, 581, 40, rl.Black)
				rl.DrawText("hp", 662, 580, 40, rl.White)
			case "cherries":
				rl.DrawText("chairees", 595, 533, 40, rl.DarkGray)
				rl.DrawText("chairees", 597, 531, 40, rl.Black)
				rl.DrawText("chairees", 598, 530, 40, menuTextSpaceColor)

				rl.DrawText("rain", 640, 583, 40, descTxtColor)
				rl.DrawText("rain", 641, 581, 40, rl.Black)
				rl.DrawText("rain", 642, 580, 40, rl.White)
			case "petSkeleton":
				rl.DrawText("skelitin", 611, 533, 40, rl.DarkGray)
				rl.DrawText("skelitin", 613, 531, 40, rl.Black)
				rl.DrawText("skelitin", 614, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 650, 583, 40, descTxtColor)
				rl.DrawText("pet", 652, 581, 40, rl.Black)
				rl.DrawText("pet", 653, 580, 40, rl.White)
			case "petSlime":
				rl.DrawText("zlime", 635, 533, 40, rl.DarkGray)
				rl.DrawText("zlime", 637, 531, 40, rl.Black)
				rl.DrawText("zlime", 638, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 650, 583, 40, descTxtColor)
				rl.DrawText("pet", 652, 581, 40, rl.Black)
				rl.DrawText("pet", 653, 580, 40, rl.White)
			case "greenHeart":
				rl.DrawText("gween hart", 568, 533, 40, rl.DarkGray)
				rl.DrawText("gween hart", 570, 531, 40, rl.Black)
				rl.DrawText("gween hart", 571, 530, 40, menuTextSpaceColor)

				rl.DrawText("randumb", 607, 583, 40, descTxtColor)
				rl.DrawText("randumb", 609, 581, 40, rl.Black)
				rl.DrawText("randumb", 610, 580, 40, rl.White)
			case "apple":
				rl.DrawText("toemate", 607, 533, 40, rl.DarkGray)
				rl.DrawText("toemate", 609, 531, 40, rl.Black)
				rl.DrawText("toemate", 610, 530, 40, menuTextSpaceColor)

				rl.DrawText("rotin", 635, 583, 40, descTxtColor)
				rl.DrawText("rotin", 637, 581, 40, rl.Black)
				rl.DrawText("rotin", 639, 580, 40, rl.White)
			case "petKnight":
				rl.DrawText("knite", 631, 533, 40, rl.DarkGray)
				rl.DrawText("knite", 633, 531, 40, rl.Black)
				rl.DrawText("knite", 634, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 650, 583, 40, descTxtColor)
				rl.DrawText("pet", 652, 581, 40, rl.Black)
				rl.DrawText("pet", 653, 580, 40, rl.White)
			case "petRedSlime":
				rl.DrawText("fred zlime", 580, 533, 40, rl.DarkGray)
				rl.DrawText("fred zlime", 582, 531, 40, rl.Black)
				rl.DrawText("fred zlime", 583, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 650, 583, 40, descTxtColor)
				rl.DrawText("pet", 652, 581, 40, rl.Black)
				rl.DrawText("pet", 653, 580, 40, rl.White)
			case "petBat":
				rl.DrawText("bart the bat", 555, 533, 40, rl.DarkGray)
				rl.DrawText("bart the bat", 557, 531, 40, rl.Black)
				rl.DrawText("bart the bat", 558, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 650, 583, 40, descTxtColor)
				rl.DrawText("pet", 652, 581, 40, rl.Black)
				rl.DrawText("pet", 653, 580, 40, rl.White)
			}
			switch shopItem3 {
			case "sawblade":
				rl.DrawText("sore", 1017, 533, 40, rl.DarkGray)
				rl.DrawText("sore", 1019, 531, 40, rl.Black)
				rl.DrawText("sore", 1020, 530, 40, menuTextSpaceColor)

				rl.DrawText("cutz", 1020, 583, 40, descTxtColor)
				rl.DrawText("cutz", 1022, 581, 40, rl.Black)
				rl.DrawText("cutz", 1023, 580, 40, rl.White)
			case "petPigeon":
				rl.DrawText("pijin", 1031, 533, 40, rl.DarkGray)
				rl.DrawText("pijin", 1033, 531, 40, rl.Black)
				rl.DrawText("pijin", 1034, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 1040, 583, 40, descTxtColor)
				rl.DrawText("pet", 1042, 581, 40, rl.Black)
				rl.DrawText("pet", 1043, 580, 40, rl.White)
			case "fireball":
				rl.DrawText("fireballz", 987, 533, 40, rl.DarkGray)
				rl.DrawText("fireballz", 989, 531, 40, rl.Black)
				rl.DrawText("fireballz", 990, 530, 40, menuTextSpaceColor)

				rl.DrawText("burns", 1017, 583, 40, descTxtColor)
				rl.DrawText("burns", 1019, 581, 40, rl.Black)
				rl.DrawText("burns", 1020, 580, 40, rl.White)
			case "poisonball":
				rl.DrawText("poizinballz", 969, 533, 40, rl.DarkGray)
				rl.DrawText("poizinballz", 971, 531, 40, rl.Black)
				rl.DrawText("poizinballz", 972, 530, 40, menuTextSpaceColor)

				rl.DrawText("poisins", 1007, 583, 40, descTxtColor)
				rl.DrawText("poisins", 1009, 581, 40, rl.Black)
				rl.DrawText("poisins", 1010, 580, 40, rl.White)
			case "iceball":
				rl.DrawText("eyzballz", 989, 533, 40, rl.DarkGray)
				rl.DrawText("eyzballz", 991, 531, 40, rl.Black)
				rl.DrawText("eyzballz", 992, 530, 40, menuTextSpaceColor)

				rl.DrawText("freezes", 991, 583, 40, descTxtColor)
				rl.DrawText("freezes", 993, 581, 40, rl.Black)
				rl.DrawText("freezes", 994, 580, 40, rl.White)
			case "cannon":
				rl.DrawText("cannin", 1005, 533, 40, rl.DarkGray)
				rl.DrawText("cannin", 1007, 531, 40, rl.Black)
				rl.DrawText("cannin", 1008, 530, 40, menuTextSpaceColor)

				rl.DrawText("blasts", 1005, 583, 40, descTxtColor)
				rl.DrawText("blasts", 1007, 581, 40, rl.Black)
				rl.DrawText("blasts", 1008, 580, 40, rl.White)
			case "trampoline":
				rl.DrawText("trampoleen", 959, 533, 40, rl.DarkGray)
				rl.DrawText("trampoleen", 961, 531, 40, rl.Black)
				rl.DrawText("trampoleen", 962, 530, 40, menuTextSpaceColor)

				rl.DrawText("bounce", 1007, 583, 40, descTxtColor)
				rl.DrawText("bounce", 1009, 581, 40, rl.Black)
				rl.DrawText("bounce", 1010, 580, 40, rl.White)
			case "petMushroom":
				rl.DrawText("mushrum", 987, 533, 40, rl.DarkGray)
				rl.DrawText("mushrum", 989, 531, 40, rl.Black)
				rl.DrawText("mushrum", 990, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 1040, 583, 40, descTxtColor)
				rl.DrawText("pet", 1042, 581, 40, rl.Black)
				rl.DrawText("pet", 1043, 580, 40, rl.White)
			case "exclamation":
				rl.DrawText("#!?", 1035, 533, 40, rl.DarkGray)
				rl.DrawText("#!?", 1037, 531, 40, rl.Black)
				rl.DrawText("#!?", 1038, 530, 40, menuTextSpaceColor)

				rl.DrawText("talk", 1026, 583, 40, descTxtColor)
				rl.DrawText("talk", 1028, 581, 40, rl.Black)
				rl.DrawText("talk", 1029, 580, 40, rl.White)
			case "propellor":
				rl.DrawText("propella", 987, 533, 40, rl.DarkGray)
				rl.DrawText("propella", 989, 531, 40, rl.Black)
				rl.DrawText("propella", 990, 530, 40, menuTextSpaceColor)

				rl.DrawText("higha", 1022, 583, 40, descTxtColor)
				rl.DrawText("higha", 1024, 581, 40, rl.Black)
				rl.DrawText("higha", 1025, 580, 40, rl.White)
			case "petGreenPig":
				rl.DrawText("litl pig", 1011, 533, 40, rl.DarkGray)
				rl.DrawText("litl pig", 1013, 531, 40, rl.Black)
				rl.DrawText("litl pig", 1014, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 1040, 583, 40, descTxtColor)
				rl.DrawText("pet", 1042, 581, 40, rl.Black)
				rl.DrawText("pet", 1043, 580, 40, rl.White)
			case "watermelon":
				rl.DrawText("wart-a-melin", 946, 533, 40, rl.DarkGray)
				rl.DrawText("wart-a-melin", 948, 531, 40, rl.Black)
				rl.DrawText("wart-a-melin", 949, 530, 40, menuTextSpaceColor)

				rl.DrawText("rain", 1031, 583, 40, descTxtColor)
				rl.DrawText("rain", 1033, 581, 40, rl.Black)
				rl.DrawText("rain", 1034, 580, 40, rl.White)
			case "heart":
				rl.DrawText("hart", 1027, 533, 40, rl.DarkGray)
				rl.DrawText("hart", 1029, 531, 40, rl.Black)
				rl.DrawText("hart", 1030, 530, 40, menuTextSpaceColor)

				rl.DrawText("hp", 1047, 583, 40, descTxtColor)
				rl.DrawText("hp", 1049, 581, 40, rl.Black)
				rl.DrawText("hp", 1050, 580, 40, rl.White)
			case "cherries":
				rl.DrawText("chairees", 985, 533, 40, rl.DarkGray)
				rl.DrawText("chairees", 987, 531, 40, rl.Black)
				rl.DrawText("chairees", 988, 530, 40, menuTextSpaceColor)

				rl.DrawText("rain", 1031, 583, 40, descTxtColor)
				rl.DrawText("rain", 1033, 581, 40, rl.Black)
				rl.DrawText("rain", 1034, 580, 40, rl.White)
			case "petSkeleton":
				rl.DrawText("skelitin", 999, 533, 40, rl.DarkGray)
				rl.DrawText("skelitin", 1001, 531, 40, rl.Black)
				rl.DrawText("skelitin", 1002, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 1040, 583, 40, descTxtColor)
				rl.DrawText("pet", 1042, 581, 40, rl.Black)
				rl.DrawText("pet", 1043, 580, 40, rl.White)
			case "petSlime":
				rl.DrawText("zlime", 1025, 533, 40, rl.DarkGray)
				rl.DrawText("zlime", 1027, 531, 40, rl.Black)
				rl.DrawText("zlime", 1028, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 1040, 583, 40, descTxtColor)
				rl.DrawText("pet", 1042, 581, 40, rl.Black)
				rl.DrawText("pet", 1043, 580, 40, rl.White)
			case "greenHeart":
				rl.DrawText("gween hart", 958, 533, 40, rl.DarkGray)
				rl.DrawText("gween hart", 960, 531, 40, rl.Black)
				rl.DrawText("gween hart", 961, 530, 40, menuTextSpaceColor)

				rl.DrawText("randumb", 992, 583, 40, descTxtColor)
				rl.DrawText("randumb", 994, 581, 40, rl.Black)
				rl.DrawText("randumb", 995, 580, 40, rl.White)
			case "apple":
				rl.DrawText("toemate", 1002, 533, 40, rl.DarkGray)
				rl.DrawText("toemate", 1004, 531, 40, rl.Black)
				rl.DrawText("toemate", 1005, 530, 40, menuTextSpaceColor)

				rl.DrawText("rotin", 1031, 583, 40, descTxtColor)
				rl.DrawText("rotin", 1033, 581, 40, rl.Black)
				rl.DrawText("rotin", 1034, 580, 40, rl.White)
			case "petKnight":
				rl.DrawText("knite", 1026, 533, 40, rl.DarkGray)
				rl.DrawText("knite", 1028, 531, 40, rl.Black)
				rl.DrawText("knite", 1029, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 1040, 583, 40, descTxtColor)
				rl.DrawText("pet", 1042, 581, 40, rl.Black)
				rl.DrawText("pet", 1043, 580, 40, rl.White)
			case "petRedSlime":
				rl.DrawText("fred zlime", 973, 533, 40, rl.DarkGray)
				rl.DrawText("fred zlime", 975, 531, 40, rl.Black)
				rl.DrawText("fred zlime", 976, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 1040, 583, 40, descTxtColor)
				rl.DrawText("pet", 1042, 581, 40, rl.Black)
				rl.DrawText("pet", 1043, 580, 40, rl.White)
			case "petBat":
				rl.DrawText("bart the bat", 945, 533, 40, rl.DarkGray)
				rl.DrawText("bart the bat", 947, 531, 40, rl.Black)
				rl.DrawText("bart the bat", 948, 530, 40, menuTextSpaceColor)

				rl.DrawText("pet", 1040, 583, 40, descTxtColor)
				rl.DrawText("pet", 1042, 581, 40, rl.Black)
				rl.DrawText("pet", 1043, 580, 40, rl.White)
			}
			// draw coin
			rl.BeginMode2D(cameraShop) // 5X zoom

			coinIMGV2 := rl.NewVector2(102, 26)
			rl.DrawTextureRec(imgs, coinIMG, coinIMGV2, rl.White)

			if frameCountGameStart%9 == 0 {
				// update coin img
				coinIMG.X += 16
				if coinIMG.X == 414 {
					coinIMG.X = 334
				}
			}

			rl.EndMode2D() // end 5X zoom

			// draw shop items

			rl.BeginMode2D(cameraShop3) // 2.4X zoom

			if shopItem1 == "sawblade" {
				boxImageV2 := rl.NewVector2(102, 147)
				rl.DrawTextureRec(imgs, sawbladeIMG, boxImageV2, rl.White)
			} else if shopItem1 == "petPigeon" {
				boxImageV2 := rl.NewVector2(102, 149)
				rl.DrawTextureRec(imgs, petPigeonIMG, boxImageV2, rl.White)
			}

			if shopItem2 == "sawblade" {
				boxImageV2 := rl.NewVector2(265, 147)
				rl.DrawTextureRec(imgs, sawbladeIMG, boxImageV2, rl.White)
			} else if shopItem2 == "petPigeon" {
				boxImageV2 := rl.NewVector2(265, 149)
				rl.DrawTextureRec(imgs, petPigeonIMG, boxImageV2, rl.White)
			}

			if shopItem3 == "sawblade" {
				boxImageV2 := rl.NewVector2(427, 147)
				rl.DrawTextureRec(imgs, sawbladeIMG, boxImageV2, rl.White)
			} else if shopItem3 == "petPigeon" {
				boxImageV2 := rl.NewVector2(427, 149)
				rl.DrawTextureRec(imgs, petPigeonIMG, boxImageV2, rl.White)
			}

			rl.EndMode2D() // end 2.4X zoom

			rl.BeginMode2D(cameraShop2) // 3.4X zoom

			if shopItem1 == "fireball" {
				boxImageV2 := rl.NewVector2(45, 112)
				rl.DrawTextureRec(imgs, fireballIMG, boxImageV2, rl.White)
			} else if shopItem1 == "poisonball" {
				boxImageV2 := rl.NewVector2(50, 112)
				rl.DrawTextureRec(imgs, poisonballIMG, boxImageV2, rl.White)
			} else if shopItem1 == "iceball" {
				boxImageV2 := rl.NewVector2(47, 112)
				rl.DrawTextureRec(imgs, iceballIMG, boxImageV2, rl.White)
			} else if shopItem1 == "cannon" {
				boxImageV2 := rl.NewVector2(64, 107)
				rl.DrawTextureRec(imgs, cannonIMG, boxImageV2, rl.White)
			} else if shopItem1 == "trampoline" {
				boxImageV2 := rl.NewVector2(72, 102)
				rl.DrawTextureRec(imgs, trampolineIMG, boxImageV2, rl.White)
			} else if shopItem1 == "petMushroom" {
				boxImageV2 := rl.NewVector2(72, 102)
				rl.DrawTextureRec(imgs, petMushroomIMG, boxImageV2, rl.White)
			}

			if shopItem2 == "fireball" {
				boxImageV2 := rl.NewVector2(160, 112)
				rl.DrawTextureRec(imgs, fireballIMG, boxImageV2, rl.White)
			} else if shopItem2 == "poisonball" {
				boxImageV2 := rl.NewVector2(165, 112)
				rl.DrawTextureRec(imgs, poisonballIMG, boxImageV2, rl.White)
			} else if shopItem2 == "iceball" {
				boxImageV2 := rl.NewVector2(162, 112)
				rl.DrawTextureRec(imgs, iceballIMG, boxImageV2, rl.White)
			} else if shopItem2 == "cannon" {
				boxImageV2 := rl.NewVector2(179, 107)
				rl.DrawTextureRec(imgs, cannonIMG, boxImageV2, rl.White)
			} else if shopItem2 == "trampoline" {
				boxImageV2 := rl.NewVector2(187, 102)
				rl.DrawTextureRec(imgs, trampolineIMG, boxImageV2, rl.White)
			} else if shopItem2 == "petMushroom" {
				boxImageV2 := rl.NewVector2(187, 102)
				rl.DrawTextureRec(imgs, petMushroomIMG, boxImageV2, rl.White)
			}

			if shopItem3 == "fireball" {
				boxImageV2 := rl.NewVector2(274, 112)
				rl.DrawTextureRec(imgs, fireballIMG, boxImageV2, rl.White)
			} else if shopItem3 == "poisonball" {
				boxImageV2 := rl.NewVector2(279, 112)
				rl.DrawTextureRec(imgs, poisonballIMG, boxImageV2, rl.White)
			} else if shopItem3 == "iceball" {
				boxImageV2 := rl.NewVector2(277, 112)
				rl.DrawTextureRec(imgs, iceballIMG, boxImageV2, rl.White)
			} else if shopItem3 == "cannon" {
				boxImageV2 := rl.NewVector2(294, 107)
				rl.DrawTextureRec(imgs, cannonIMG, boxImageV2, rl.White)
			} else if shopItem3 == "trampoline" {
				boxImageV2 := rl.NewVector2(302, 102)
				rl.DrawTextureRec(imgs, trampolineIMG, boxImageV2, rl.White)
			} else if shopItem3 == "petMushroom" {
				boxImageV2 := rl.NewVector2(302, 102)
				rl.DrawTextureRec(imgs, petMushroomIMG, boxImageV2, rl.White)
			}

			rl.EndMode2D() // end 3.4X zoom

			rl.BeginMode2D(cameraShop5) // 4.2X zoom

			if shopItem1 == "exclamation" {
				boxImageV2 := rl.NewVector2(55, 86)
				rl.DrawTextureRec(imgs, exclamationIMG, boxImageV2, rl.White)
			} else if shopItem1 == "propellor" {
				boxImageV2 := rl.NewVector2(57, 89)
				rl.DrawTextureRec(imgs, propellorIMG, boxImageV2, rl.White)
			} else if shopItem1 == "petGreenPig" {
				boxImageV2 := rl.NewVector2(57, 83)
				rl.DrawTextureRec(imgs, petGreenPigIMG, boxImageV2, rl.White)
			}

			if shopItem2 == "propellor" {
				boxImageV2 := rl.NewVector2(151, 89)
				rl.DrawTextureRec(imgs, propellorIMG, boxImageV2, rl.White)
			} else if shopItem2 == "exclamation" {
				boxImageV2 := rl.NewVector2(149, 86)
				rl.DrawTextureRec(imgs, exclamationIMG, boxImageV2, rl.White)
			} else if shopItem2 == "petGreenPig" {
				boxImageV2 := rl.NewVector2(151, 83)
				rl.DrawTextureRec(imgs, petGreenPigIMG, boxImageV2, rl.White)
			}

			if shopItem3 == "exclamation" {
				boxImageV2 := rl.NewVector2(241, 86)
				rl.DrawTextureRec(imgs, exclamationIMG, boxImageV2, rl.White)
			} else if shopItem3 == "propellor" {
				boxImageV2 := rl.NewVector2(244, 89)
				rl.DrawTextureRec(imgs, propellorIMG, boxImageV2, rl.White)
			} else if shopItem3 == "petGreenPig" {
				boxImageV2 := rl.NewVector2(244, 83)
				rl.DrawTextureRec(imgs, petGreenPigIMG, boxImageV2, rl.White)
			}

			rl.EndMode2D() // end 4.2X zoom

			rl.BeginMode2D(cameraShop4) // 6X zoom

			if shopItem1 == "watermelon" {
				boxImageV2 := rl.NewVector2(33, 57)
				rl.DrawTextureRec(imgs, watermelonIMG, boxImageV2, rl.White)
			} else if shopItem1 == "heart" {
				boxImageV2 := rl.NewVector2(43, 60)
				rl.DrawTextureRec(imgs, heartShopIMG, boxImageV2, rl.White)
			} else if shopItem1 == "cherries" {
				boxImageV2 := rl.NewVector2(33, 54)
				rl.DrawTextureRec(imgs, cherriesIMG, boxImageV2, rl.White)
			} else if shopItem1 == "petSkeleton" {
				boxImageV2 := rl.NewVector2(42, 58)
				rl.DrawTextureRec(imgs, petSkeletonIMG, boxImageV2, rl.White)
			} else if shopItem1 == "petSlime" {
				boxImageV2 := rl.NewVector2(40, 60)
				rl.DrawTextureRec(imgs, petSlimeIMG, boxImageV2, rl.White)
			} else if shopItem1 == "greenHeart" {
				boxImageV2 := rl.NewVector2(40, 60)
				rl.DrawTextureRec(imgs, greenHeartIMG, boxImageV2, rl.White)
			} else if shopItem1 == "apple" {
				boxImageV2 := rl.NewVector2(34, 55)
				rl.DrawTextureRec(imgs, appleIMG, boxImageV2, rl.White)
			} else if shopItem1 == "petKnight" {
				boxImageV2 := rl.NewVector2(40, 58)
				rl.DrawTextureRec(imgs, petKnightIMG, boxImageV2, rl.White)
			} else if shopItem1 == "petRedSlime" {
				boxImageV2 := rl.NewVector2(42, 59)
				rl.DrawTextureRec(imgs, petRedSlimeIMG, boxImageV2, rl.White)
			} else if shopItem1 == "petBat" {
				boxImageV2 := rl.NewVector2(40, 59)
				rl.DrawTextureRec(imgs, petBatIMG, boxImageV2, rl.White)
			}

			if shopItem2 == "heart" {
				boxImageV2 := rl.NewVector2(108, 60)
				rl.DrawTextureRec(imgs, heartShopIMG, boxImageV2, rl.White)
			} else if shopItem2 == "watermelon" {
				boxImageV2 := rl.NewVector2(98, 57)
				rl.DrawTextureRec(imgs, watermelonIMG, boxImageV2, rl.White)
			} else if shopItem2 == "apple" {
				boxImageV2 := rl.NewVector2(99, 55)
				rl.DrawTextureRec(imgs, appleIMG, boxImageV2, rl.White)
			} else if shopItem2 == "cherries" {
				boxImageV2 := rl.NewVector2(98, 54)
				rl.DrawTextureRec(imgs, cherriesIMG, boxImageV2, rl.White)
			} else if shopItem2 == "petSlime" {
				boxImageV2 := rl.NewVector2(106, 60)
				rl.DrawTextureRec(imgs, petSlimeIMG, boxImageV2, rl.White)
			} else if shopItem2 == "petSkeleton" {
				boxImageV2 := rl.NewVector2(107, 58)
				rl.DrawTextureRec(imgs, petSkeletonIMG, boxImageV2, rl.White)
			} else if shopItem2 == "greenHeart" {
				boxImageV2 := rl.NewVector2(106, 60)
				rl.DrawTextureRec(imgs, greenHeartIMG, boxImageV2, rl.White)
			} else if shopItem2 == "petKnight" {
				boxImageV2 := rl.NewVector2(105, 58)
				rl.DrawTextureRec(imgs, petKnightIMG, boxImageV2, rl.White)
			} else if shopItem2 == "petRedSlime" {
				boxImageV2 := rl.NewVector2(107, 59)
				rl.DrawTextureRec(imgs, petRedSlimeIMG, boxImageV2, rl.White)
			} else if shopItem2 == "petBat" {
				boxImageV2 := rl.NewVector2(106, 59)
				rl.DrawTextureRec(imgs, petBatIMG, boxImageV2, rl.White)
			}

			if shopItem3 == "apple" {
				boxImageV2 := rl.NewVector2(163, 55)
				rl.DrawTextureRec(imgs, appleIMG, boxImageV2, rl.White)
			} else if shopItem3 == "cherries" {
				boxImageV2 := rl.NewVector2(163, 54)
				rl.DrawTextureRec(imgs, cherriesIMG, boxImageV2, rl.White)
			} else if shopItem3 == "heart" {
				boxImageV2 := rl.NewVector2(173, 60)
				rl.DrawTextureRec(imgs, heartShopIMG, boxImageV2, rl.White)
			} else if shopItem3 == "watermelon" {
				boxImageV2 := rl.NewVector2(163, 57)
				rl.DrawTextureRec(imgs, watermelonIMG, boxImageV2, rl.White)
			} else if shopItem3 == "petSkeleton" {
				boxImageV2 := rl.NewVector2(172, 58)
				rl.DrawTextureRec(imgs, petSkeletonIMG, boxImageV2, rl.White)
			} else if shopItem3 == "petSlime" {
				boxImageV2 := rl.NewVector2(171, 60)
				rl.DrawTextureRec(imgs, petSlimeIMG, boxImageV2, rl.White)
			} else if shopItem3 == "greenHeart" {
				boxImageV2 := rl.NewVector2(171, 60)
				rl.DrawTextureRec(imgs, greenHeartIMG, boxImageV2, rl.White)
			} else if shopItem3 == "petKnight" {
				boxImageV2 := rl.NewVector2(171, 58)
				rl.DrawTextureRec(imgs, petKnightIMG, boxImageV2, rl.White)
			} else if shopItem3 == "petRedSlime" {
				boxImageV2 := rl.NewVector2(172, 59)
				rl.DrawTextureRec(imgs, petRedSlimeIMG, boxImageV2, rl.White)
			} else if shopItem3 == "petBat" {
				boxImageV2 := rl.NewVector2(172, 59)
				rl.DrawTextureRec(imgs, petBatIMG, boxImageV2, rl.White)
			}

			rl.EndMode2D() // end 6X zoom

			// MARK: shop animations
			switch shopMenuCount {
			case 1:
				switch shopItem1 {
				case "petMushroom":
					if frameCountGameStart%9 == 0 {
						petMushroomIMG.X += 32
						if petMushroomIMG.X == 1473 {
							petMushroomIMG.X = 1281
						}
					}
				case "petBat":
					if frameCountGameStart%6 == 0 {
						petBatIMG.X += 16
						if petBatIMG.X == 1483 {
							petBatIMG.X = 1403
						}
					}
				case "petPigeon":
					if frameCountGameStart%10 == 0 {
						petPigeonIMG.X += 48
						if petPigeonIMG.X == 1674 {
							petPigeonIMG.X = 1482
						}
					}
				case "petGreenPig":
					if frameCountGameStart%4 == 0 {
						petGreenPigIMG.X += 34
						if petGreenPigIMG.X == 2018 {
							petGreenPigIMG.X = 1678
						}
					}
				case "petRedSlime":
					if frameCountGameStart%6 == 0 {
						petRedSlimeIMG.Y += 16
						if petRedSlimeIMG.Y == 215 {
							petRedSlimeIMG.Y = 87
						}
					}
				case "petKnight":
					if frameCountGameStart%6 == 0 {
						petKnightIMG.X += 16
						if petKnightIMG.X == 1509 {
							petKnightIMG.X = 1413
						}
					}
				case "exclamation":
					if frameCountGameStart%15 == 0 {
						exclamationIMG.X += 34
						if exclamationIMG.X == 1397 {
							exclamationIMG.X = 1295
						}
					}
				case "propellor":
					if frameCountGameStart%5 == 0 {
						propellorIMG.X += 24
						if propellorIMG.X == 1509 {
							propellorIMG.X = 1413
						}
					}
				case "trampoline":
					if frameCountGameStart%10 == 0 {
						trampolineIMG.X += 28
						if trampolineIMG.X == 1507 {
							trampolineIMG.X = 1283
						}
					}
				case "greenHeart":
					if frameCountGameStart%10 == 0 {
						greenHeartIMG.X += 16
						if greenHeartIMG.X == 1527 {
							greenHeartIMG.X = 1399
						}
					}
				case "petSkeleton":
					if frameCountGameStart%10 == 0 {
						petSkeletonIMG.X += 16
						if petSkeletonIMG.X == 1582 {
							petSkeletonIMG.X = 1518
						}
					}
				case "petSlime":
					if frameCountGameStart%7 == 0 {
						petSlimeIMG.X += 16
						if petSlimeIMG.X == 1503 {
							petSlimeIMG.X = 1407
						}
					}
				case "cannon":
					if frameCountGameStart%10 == 0 {
						cannonIMG.X -= 44
						if cannonIMG.X == 1531 {
							cannonIMG.X = 1663
						}
					}
				case "fireball":
					fireballIMG.X -= 68
					if fireballIMG.X == 1368 {
						fireballIMG.X = 1980
					}
				case "watermelon":
					if frameCountGameStart%3 == 0 {
						watermelonIMG.X += 32
						if watermelonIMG.X == 2048 {
							watermelonIMG.X = 1504
						}
					}
				case "poisonball":
					poisonballIMG.X -= 65
					if poisonballIMG.X == 1400 {
						poisonballIMG.X = 1985
					}
				case "iceball":
					iceballIMG.X -= 84
					if iceballIMG.X == 1230 {
						iceballIMG.X = 1986
					}
				case "sawblade":
					if frameCountGameStart%2 == 0 {
						sawbladeIMG.X += 38
						if sawbladeIMG.X == 2044 {
							sawbladeIMG.X = 1740
						}
					}
				case "apple":
					if frameCountGameStart%3 == 0 {
						appleIMG.X += 32
						if appleIMG.X == 2048 {
							appleIMG.X = 1504
						}
					}
				case "cherries":
					if frameCountGameStart%3 == 0 {
						cherriesIMG.X += 32
						if cherriesIMG.X == 2048 {
							cherriesIMG.X = 1504
						}
					}
				case "heart":
					if frameCountGameStart%6 == 0 {
						heartShopIMG.X += 18
						if heartShopIMG.X == 1739 {
							heartShopIMG.X = 1595
						}
					}
				}
			case 2:
				switch shopItem2 {
				case "petRedSlime":
					if frameCountGameStart%6 == 0 {
						petRedSlimeIMG.Y += 16
						if petRedSlimeIMG.Y == 215 {
							petRedSlimeIMG.Y = 87
						}
					}
				case "petMushroom":
					if frameCountGameStart%9 == 0 {
						petMushroomIMG.X += 32
						if petMushroomIMG.X == 1473 {
							petMushroomIMG.X = 1281
						}
					}
				case "petBat":
					if frameCountGameStart%6 == 0 {
						petBatIMG.X += 16
						if petBatIMG.X == 1483 {
							petBatIMG.X = 1403
						}
					}
				case "petPigeon":
					if frameCountGameStart%10 == 0 {
						petPigeonIMG.X += 48
						if petPigeonIMG.X == 1674 {
							petPigeonIMG.X = 1482
						}
					}
				case "petGreenPig":
					if frameCountGameStart%4 == 0 {
						petGreenPigIMG.X += 34
						if petGreenPigIMG.X == 2018 {
							petGreenPigIMG.X = 1678
						}
					}
				case "petKnight":
					if frameCountGameStart%6 == 0 {
						petKnightIMG.X += 16
						if petKnightIMG.X == 1509 {
							petKnightIMG.X = 1413
						}
					}
				case "exclamation":
					if frameCountGameStart%15 == 0 {
						exclamationIMG.X += 34
						if exclamationIMG.X == 1397 {
							exclamationIMG.X = 1295
						}
					}
				case "propellor":
					if frameCountGameStart%5 == 0 {
						propellorIMG.X += 24
						if propellorIMG.X == 1509 {
							propellorIMG.X = 1413
						}
					}
				case "trampoline":
					if frameCountGameStart%10 == 0 {
						trampolineIMG.X += 28
						if trampolineIMG.X == 1507 {
							trampolineIMG.X = 1283
						}
					}
				case "greenHeart":
					if frameCountGameStart%10 == 0 {
						greenHeartIMG.X += 16
						if greenHeartIMG.X == 1527 {
							greenHeartIMG.X = 1399
						}
					}
				case "petSkeleton":
					if frameCountGameStart%10 == 0 {
						petSkeletonIMG.X += 16
						if petSkeletonIMG.X == 1582 {
							petSkeletonIMG.X = 1518
						}
					}
				case "petSlime":
					if frameCountGameStart%7 == 0 {
						petSlimeIMG.X += 16
						if petSlimeIMG.X == 1503 {
							petSlimeIMG.X = 1407
						}
					}
				case "heart":
					if frameCountGameStart%6 == 0 {
						heartShopIMG.X += 18
						if heartShopIMG.X == 1739 {
							heartShopIMG.X = 1595
						}
					}
				case "cannon":
					if frameCountGameStart%10 == 0 {
						cannonIMG.X -= 44
						if cannonIMG.X == 1531 {
							cannonIMG.X = 1663
						}
					}
				case "fireball":
					fireballIMG.X -= 68
					if fireballIMG.X == 1368 {
						fireballIMG.X = 1980
					}
				case "watermelon":
					if frameCountGameStart%3 == 0 {
						watermelonIMG.X += 32
						if watermelonIMG.X == 2048 {
							watermelonIMG.X = 1504
						}
					}
				case "poisonball":
					poisonballIMG.X -= 65
					if poisonballIMG.X == 1400 {
						poisonballIMG.X = 1985
					}
				case "iceball":
					iceballIMG.X -= 84
					if iceballIMG.X == 1230 {
						iceballIMG.X = 1986
					}
				case "sawblade":
					if frameCountGameStart%2 == 0 {
						sawbladeIMG.X += 38
						if sawbladeIMG.X == 2044 {
							sawbladeIMG.X = 1740
						}
					}
				case "apple":
					if frameCountGameStart%3 == 0 {
						appleIMG.X += 32
						if appleIMG.X == 2048 {
							appleIMG.X = 1504
						}
					}
				case "cherries":
					if frameCountGameStart%3 == 0 {
						cherriesIMG.X += 32
						if cherriesIMG.X == 2048 {
							cherriesIMG.X = 1504
						}
					}
				}
			case 3:
				switch shopItem3 {
				case "petRedSlime":
					if frameCountGameStart%6 == 0 {
						petRedSlimeIMG.Y += 16
						if petRedSlimeIMG.Y == 215 {
							petRedSlimeIMG.Y = 87
						}
					}
				case "petMushroom":
					if frameCountGameStart%9 == 0 {
						petMushroomIMG.X += 32
						if petMushroomIMG.X == 1473 {
							petMushroomIMG.X = 1281
						}
					}
				case "petBat":
					if frameCountGameStart%6 == 0 {
						petBatIMG.X += 16
						if petBatIMG.X == 1483 {
							petBatIMG.X = 1403
						}
					}
				case "petPigeon":
					if frameCountGameStart%10 == 0 {
						petPigeonIMG.X += 48
						if petPigeonIMG.X == 1674 {
							petPigeonIMG.X = 1482
						}
					}
				case "petGreenPig":
					if frameCountGameStart%4 == 0 {
						petGreenPigIMG.X += 34
						if petGreenPigIMG.X == 2018 {
							petGreenPigIMG.X = 1678
						}
					}
				case "petKnight":
					if frameCountGameStart%6 == 0 {
						petKnightIMG.X += 16
						if petKnightIMG.X == 1509 {
							petKnightIMG.X = 1413
						}
					}
				case "heart":
					if frameCountGameStart%6 == 0 {
						heartShopIMG.X += 18
						if heartShopIMG.X == 1739 {
							heartShopIMG.X = 1595
						}
					}
				case "exclamation":
					if frameCountGameStart%15 == 0 {
						exclamationIMG.X += 34
						if exclamationIMG.X == 1397 {
							exclamationIMG.X = 1295
						}
					}
				case "propellor":
					if frameCountGameStart%5 == 0 {
						propellorIMG.X += 24
						if propellorIMG.X == 1509 {
							propellorIMG.X = 1413
						}
					}
				case "trampoline":
					if frameCountGameStart%10 == 0 {
						trampolineIMG.X += 28
						if trampolineIMG.X == 1507 {
							trampolineIMG.X = 1283
						}
					}
				case "greenHeart":
					if frameCountGameStart%10 == 0 {
						greenHeartIMG.X += 16
						if greenHeartIMG.X == 1527 {
							greenHeartIMG.X = 1399
						}
					}
				case "petSkeleton":
					if frameCountGameStart%10 == 0 {
						petSkeletonIMG.X += 16
						if petSkeletonIMG.X == 1582 {
							petSkeletonIMG.X = 1518
						}
					}
				case "petSlime":
					if frameCountGameStart%7 == 0 {
						petSlimeIMG.X += 16
						if petSlimeIMG.X == 1503 {
							petSlimeIMG.X = 1407
						}
					}
				case "cherries":
					if frameCountGameStart%3 == 0 {
						cherriesIMG.X += 32
						if cherriesIMG.X == 2048 {
							cherriesIMG.X = 1504
						}
					}
				case "cannon":
					if frameCountGameStart%10 == 0 {
						cannonIMG.X -= 44
						if cannonIMG.X == 1531 {
							cannonIMG.X = 1663
						}
					}
				case "fireball":
					fireballIMG.X -= 68
					if fireballIMG.X == 1368 {
						fireballIMG.X = 1980
					}
				case "watermelon":
					if frameCountGameStart%3 == 0 {
						watermelonIMG.X += 32
						if watermelonIMG.X == 2048 {
							watermelonIMG.X = 1504
						}
					}
				case "poisonball":
					poisonballIMG.X -= 65
					if poisonballIMG.X == 1400 {
						poisonballIMG.X = 1985
					}
				case "iceball":
					iceballIMG.X -= 84
					if iceballIMG.X == 1230 {
						iceballIMG.X = 1986
					}
				case "sawblade":
					if frameCountGameStart%2 == 0 {
						sawbladeIMG.X += 38
						if sawbladeIMG.X == 2044 {
							sawbladeIMG.X = 1740
						}
					}
				case "apple":
					if frameCountGameStart%3 == 0 {
						appleIMG.X += 32
						if appleIMG.X == 2048 {
							appleIMG.X = 1504
						}
					}
				}
			}

			// draw pixel noise
			if pixelNoiseOn {
				if frameCountGameStart%2 == 0 {
					cPIXELNOISE()
				}

				lineCountPixelNoise := 0
				pixelNoiseY := int32(0)
				pixelNoiseX := int32(0)
				for a := 0; a < 880; a++ {

					if pixelNoiseMAP[a] == true {
						rl.DrawRectangle(pixelNoiseX, pixelNoiseY, 2, 2, rl.Black)
					}

					lineCountPixelNoise += 34
					pixelNoiseX += 34
					if lineCountPixelNoise > 1350 {
						lineCountPixelNoise = 0
						pixelNoiseX = 0
						pixelNoiseY += 34

					}
				}

			}

			// draw scan lines
			if scanLinesOn {
				if switchScanLines {
					linesY := int32(0)
					for a := 0; a < int(screenH); a++ {
						rl.DrawLine(0, linesY, screenW, linesY, rl.Fade(rl.Black, 0.5))
						linesY += 2
						a++
					}
				} else {
					linesY := int32(1)
					for a := 0; a < int(screenH); a++ {
						rl.DrawLine(0, linesY, screenW, linesY, rl.Fade(rl.Black, 0.5))
						linesY += 2
						a++
					}

				}
			}

		}

		// MARK: level end
		if levelEnd {
			rl.DrawRectangle(0, 0, screenW, screenH, rl.Black)
			levelMusic = false
			levelEndMusic = true

			// flash shop item text
			if frameCountGameStart%15 == 0 {
				if menuTextSpaceColor == rl.Yellow {
					menuTextSpaceColor = rl.Gold
				} else {
					menuTextSpaceColor = rl.Yellow
				}
				if descTxtColor == rl.Green {
					descTxtColor = rl.DarkGreen
				} else {
					descTxtColor = rl.Green
				}
			}

			if playerCoins == 0 {
				levelEndMenuCount = 2
			}

			// draw splatter circles
			for a := 0; a < 30; a++ {
				introCircleX := introScreenCircleXMAP[a]
				introCircleY := introScreenCircleYMAP[a]
				introCircleRadius := introScreenCircleRadius[a]

				rl.DrawCircle(introCircleX, introCircleY, introCircleRadius, rl.Fade(rl.Red, 0.4))

			}
			if frameCountGameStart%3 == 0 {
				for a := 0; a < 30; a++ {
					introCircleRadius := introScreenCircleRadius[a]
					introCircleRadius--
					introScreenCircleRadius[a] = introCircleRadius
				}
			}

			sumIntroRadius := float32(0)
			for _, introRadius := range introScreenCircleRadius {
				sumIntroRadius += introRadius
			}

			if sumIntroRadius <= 0 {
				createIntroCircles = false
			}

			// keys level end
			if rl.IsKeyPressed(rl.KeyLeft) {
				if playerCoins > 0 {
					levelEndMenuCount--
					if levelEndMenuCount == 0 {
						levelEndMenuCount = 2
					}
				} else {
					noCoinsTextOn = true
				}
			}
			if rl.IsKeyPressed(rl.KeyRight) {
				if playerCoins > 0 {
					levelEndMenuCount++
					if levelEndMenuCount == 3 {
						levelEndMenuCount = 1
					}
				} else {
					noCoinsTextOn = true
				}
			}
			if rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyKpEnter) {
				if levelEndMenuCount == 1 {
					levelEnd = false
					shopMenuCount = 1
					for {

						chooseShopItem1 := rInt(0, len(shopItemsMAP))
						chooseShopItem2 := rInt(0, len(shopItemsMAP))
						chooseShopItem3 := rInt(0, len(shopItemsMAP))

						if chooseShopItem1 != chooseShopItem2 && chooseShopItem2 != chooseShopItem3 && chooseShopItem1 != chooseShopItem3 {

							shopItem1 = shopItemsMAP[chooseShopItem1]
							shopItem2 = shopItemsMAP[chooseShopItem2]
							shopItem3 = shopItemsMAP[chooseShopItem3]
							break
						}
					}
					if playerCoins == 1 {
						shopPrice1 = 1
						shopPrice2 = 1
						shopPrice3 = 1
					} else if playerCoins < 5 {
						shopPrice1 = rInt(2, 5)
						shopPrice2 = rInt(2, 5)
						shopPrice3 = rInt(2, 5)
					} else if playerCoins < 10 && playerCoins >= 5 {
						shopPrice1 = rInt(5, 10)
						shopPrice2 = rInt(5, 10)
						shopPrice3 = rInt(5, 10)
					} else if playerCoins >= 10 && playerCoins < 20 {
						shopPrice1 = rInt(10, 25)
						shopPrice2 = rInt(10, 25)
						shopPrice3 = rInt(10, 25)
					} else if playerCoins >= 20 && playerCoins < 40 {
						shopPrice1 = rInt(20, 50)
						shopPrice2 = rInt(20, 50)
						shopPrice3 = rInt(20, 50)
					} else if playerCoins > 40 {
						shopPrice1 = rInt(40, 200)
						shopPrice2 = rInt(40, 200)
						shopPrice3 = rInt(40, 200)
					}

					shopOn = true

				} else if levelEndMenuCount == 2 {
					startNewLevel()
				}
			}

			// draw text
			killzCountTEXT := strconv.Itoa(killzCount)
			totalKillzTEXT := strconv.Itoa(totalKillz)
			playerCoinsTEXT := strconv.Itoa(playerCoins)
			levelNumberTEXT := strconv.Itoa(totalLevels)
			minibosskillsTEXT := strconv.Itoa(minibosskills)

			// text flicker
			if frameCountGameStart%24 == 0 {
				if flickerText {
					flickerText = false
				} else {
					flickerText = true
				}
			}

			rl.DrawText("the end of level", 47, 53, 80, rl.DarkGray)
			rl.DrawText("the end of level", 49, 51, 80, rl.Black)
			rl.DrawText("the end of level", 50, 50, 80, rl.White)

			rl.DrawText(levelNumberTEXT, 737, 53, 80, rl.DarkGray)
			rl.DrawText(levelNumberTEXT, 739, 51, 80, rl.Black)
			rl.DrawText(levelNumberTEXT, 740, 50, 80, rl.White)

			rl.DrawText("level killz", 197, 203, 40, rl.DarkGray)
			rl.DrawText("level killz", 199, 201, 40, rl.Black)
			rl.DrawText("level killz", 200, 200, 40, rl.White)

			rl.DrawText(killzCountTEXT, 447, 202, 40, descTxtColor)
			rl.DrawText(killzCountTEXT, 449, 201, 40, rl.Black)
			rl.DrawText(killzCountTEXT, 450, 200, 40, rl.White)

			rl.DrawText("game killz", 197, 253, 40, rl.DarkGray)
			rl.DrawText("game killz", 199, 251, 40, rl.Black)
			rl.DrawText("game killz", 200, 250, 40, rl.White)

			rl.DrawText(totalKillzTEXT, 447, 253, 40, descTxtColor)
			rl.DrawText(totalKillzTEXT, 449, 251, 40, rl.Black)
			rl.DrawText(totalKillzTEXT, 450, 250, 40, rl.White)

			rl.DrawText("boss killz", 197, 303, 40, rl.DarkGray)
			rl.DrawText("boss killz", 199, 301, 40, rl.Black)
			rl.DrawText("boss killz", 200, 300, 40, rl.White)

			rl.DrawText(minibosskillsTEXT, 447, 303, 40, descTxtColor)
			rl.DrawText(minibosskillsTEXT, 449, 301, 40, rl.Black)
			rl.DrawText(minibosskillsTEXT, 450, 300, 40, rl.White)

			rl.DrawText("coins", 284, 353, 40, rl.DarkGray)
			rl.DrawText("coins", 286, 351, 40, rl.Black)
			rl.DrawText("coins", 287, 350, 40, rl.White)

			rl.DrawText(playerCoinsTEXT, 447, 353, 40, descTxtColor)
			rl.DrawText(playerCoinsTEXT, 449, 351, 40, rl.Black)
			rl.DrawText(playerCoinsTEXT, 450, 350, 40, rl.White)

			// draw buttons
			rl.DrawRectangle(screenW/5-4, screenH-236, 260, 80, rl.Black)
			rl.DrawRectangle(screenW/5-4, screenH-236, 260, 80, rl.Fade(rl.White, 0.4))
			rl.DrawRectangle(screenW/5-1, screenH-239, 260, 80, rl.Black)
			if levelEndMenuCount == 1 {
				if flashButtons {
					rl.DrawRectangle(screenW/5, screenH-240, 260, 80, rl.Maroon)
				} else {
					rl.DrawRectangle(screenW/5, screenH-240, 260, 80, rl.Gold)
				}
			} else {
				rl.DrawRectangle(screenW/5, screenH-240, 260, 80, rl.Gold)
			}
			rl.DrawText("shop", screenW/5+77, screenH-217, 40, rl.White)
			rl.DrawText("shop", screenW/5+79, screenH-219, 40, rl.Red)
			rl.DrawText("shop", screenW/5+80, screenH-220, 40, rl.Black)

			if noCoinsTextOn {
				rl.DrawText("no coins! collect more", screenW/8+17, screenH-137, 40, descTxtColor)
				rl.DrawText("no coins! collect more", screenW/8+19, screenH-139, 40, rl.Black)
				rl.DrawText("no coins! collect more", screenW/8+20, screenH-140, 40, rl.White)
			}

			rl.DrawRectangle(screenW/2+196, screenH-236, 260, 80, rl.Black)
			rl.DrawRectangle(screenW/2+196, screenH-236, 260, 80, rl.Fade(rl.White, 0.4))
			rl.DrawRectangle(screenW/2+199, screenH-239, 260, 80, rl.Black)
			if levelEndMenuCount == 2 {
				if flashButtons {
					rl.DrawRectangle(screenW/2+200, screenH-240, 260, 80, rl.Maroon)
				} else {
					rl.DrawRectangle(screenW/2+200, screenH-240, 260, 80, rl.Gold)
				}
			} else {
				rl.DrawRectangle(screenW/2+200, screenH-240, 260, 80, rl.Gold)
			}
			rl.DrawText("next level", screenW/2+227, screenH-217, 40, rl.White)
			rl.DrawText("next level", screenW/2+229, screenH-219, 40, rl.Red)
			rl.DrawText("next level", screenW/2+230, screenH-220, 40, rl.Black)

			// draw weapon
			rl.BeginMode2D(cameraShop4)
			switch currentPlayerWeapon {
			case "weapon1TL":
				rl.DrawTextureRec(imgs, weapon1LIMG, weaponLevelEndV2, rl.White)
			case "weapon2TL":
				rl.DrawTextureRec(imgs, weapon2LIMG, weaponLevelEndV2, rl.White)
			case "weapon3TL":
				rl.DrawTextureRec(imgs, weapon3LIMG, weaponLevelEndV2, rl.White)
			case "weapon4TL":
				rl.DrawTextureRec(imgs, weapon4LIMG, weaponLevelEndV2, rl.White)
			case "weapon5TL":
				rl.DrawTextureRec(imgs, weapon5LIMG, weaponLevelEndV2, rl.White)
			case "weapon6TL":
				rl.DrawTextureRec(imgs, weapon6LIMG, weaponLevelEndV2, rl.White)
			case "weapon7TL":
				rl.DrawTextureRec(imgs, weapon7LIMG, weaponLevelEndV2, rl.White)
			}

			// draw dino
			switch dinoType {
			case "greenDino":
				dinoV2 := rl.NewVector2(160, 40)
				rl.DrawTextureRec(imgs, dinoGreenLIMG, dinoV2, rl.White)
			case "redDino":
				dinoV2 := rl.NewVector2(160, 40)
				rl.DrawTextureRec(imgs, dinoRedLIMG, dinoV2, rl.White)
			case "yellowDino":
				dinoV2 := rl.NewVector2(160, 40)
				rl.DrawTextureRec(imgs, dinoYellowLIMG, dinoV2, rl.White)
			case "blueDino":
				dinoV2 := rl.NewVector2(160, 40)
				rl.DrawTextureRec(imgs, dinoBlueLIMG, dinoV2, rl.White)

			}

			rl.EndMode2D()

			// draw enemies
			if frameCountGameStart%6 == 0 {
				enemy1IMG.X += 32
				if enemy1IMG.X >= 1024 {
					enemy1IMG.X = 514
				}
			}
			if frameCountGameStart%3 == 0 {
				enemy7UPIMG.X += 44
				if enemy7UPIMG.X >= 1560 {
					enemy7UPIMG.X = 1218
				}
			}
			rl.BeginMode2D(camera)

			enemyendv2.X--
			if enemyendv2.X <= -64 {
				enemyendv2.X = 700
			}
			enemyend2v2.X++
			if enemyend2v2.X >= 700 {
				enemyend2v2.X = -64
			}
			rl.DrawTextureRec(imgs, enemy1IMG, enemyendv2, rl.Fade(rl.White, 0.4))
			rl.DrawTextureRec(imgs, enemy7UPIMG, enemyend2v2, rl.Fade(rl.White, 0.4))
			rl.EndMode2D()
			// dino text

			switch dinoType {
			case "greenDino":
				rl.DrawText("Genie", 951, 373, 40, descTxtColor)
				rl.DrawText("Genie", 953, 371, 40, rl.Black)
				rl.DrawText("Genie", 954, 370, 40, rl.White)
			case "redDino":
				rl.DrawText("Fred", 962, 373, 40, descTxtColor)
				rl.DrawText("Fred", 964, 371, 40, rl.Black)
				rl.DrawText("Fred", 965, 370, 40, rl.White)
			case "yellowDino":
				rl.DrawText("Othello", 929, 373, 40, descTxtColor)
				rl.DrawText("Othello", 931, 371, 40, rl.Black)
				rl.DrawText("Othello", 932, 370, 40, rl.White)
			case "blueDino":
				rl.DrawText("Louie", 951, 373, 40, descTxtColor)
				rl.DrawText("Louie", 953, 371, 40, rl.Black)
				rl.DrawText("Louie", 954, 370, 40, rl.White)
			}

			// move weapon
			if frameCountGameStart%16 == 0 {
				if weaponUpDown {
					weaponLevelEndV2.Y -= 2
					weaponUpDown = false
				} else {
					weaponLevelEndV2.Y += 2
					weaponUpDown = true
				}
			}

			// flash buttons
			if frameCountGameStart%15 == 0 {

				if flashButtons {
					flashButtons = false
				} else {
					flashButtons = true
				}

				// animate dino
				switch dinoType {
				case "greenDino":
					dinoGreenLIMG.X -= 24
					if dinoGreenLIMG.X < 269 {
						dinoGreenLIMG.X = 317
					}
				case "redDino":
					dinoRedLIMG.X -= 24
					if dinoRedLIMG.X < 267 {
						dinoRedLIMG.X = 315
					}
				case "yellowDino":
					dinoYellowLIMG.X -= 24
					if dinoYellowLIMG.X < 269 {
						dinoYellowLIMG.X = 317
					}
				case "blueDino":
					dinoBlueLIMG.X -= 24
					if dinoBlueLIMG.X < 269 {
						dinoBlueLIMG.X = 317
					}
				}

			}

			// draw pixel noise
			if pixelNoiseOn {
				if frameCountGameStart%2 == 0 {
					cPIXELNOISE()
				}

				lineCountPixelNoise := 0
				pixelNoiseY := int32(0)
				pixelNoiseX := int32(0)
				for a := 0; a < 880; a++ {

					if pixelNoiseMAP[a] == true {
						rl.DrawRectangle(pixelNoiseX, pixelNoiseY, 2, 2, rl.Black)
					}

					lineCountPixelNoise += 34
					pixelNoiseX += 34
					if lineCountPixelNoise > 1350 {
						lineCountPixelNoise = 0
						pixelNoiseX = 0
						pixelNoiseY += 34

					}
				}

			}

			// draw scan lines
			if scanLinesOn {
				if switchScanLines {
					linesY := int32(0)
					for a := 0; a < int(screenH); a++ {
						rl.DrawLine(0, linesY, screenW, linesY, rl.Fade(rl.Black, 0.5))
						linesY += 2
						a++
					}
				} else {
					linesY := int32(1)
					for a := 0; a < int(screenH); a++ {
						rl.DrawLine(0, linesY, screenW, linesY, rl.Fade(rl.Black, 0.5))
						linesY += 2
						a++
					}

				}
			}

			// MARK: clear active special items
			activeSpecialOn = false
			activeSpecialActive = false
			activeSpecialComplete = false
			trampolineOn = false
			pigeonOn = false
			mushroomJumpOn = false
			exclamationOn = false
			propellorOn = false

		}

		// MARK: debugging
		if debuggingOn {
			rl.DrawRectangle(screenW-300, 0, 500, screenW, rl.Fade(rl.Black, 0.7))

			playerCurrentBlockTEXT := strconv.Itoa(playerCurrentBlock)
			playerHorizontalTEXT := strconv.Itoa(playerHorizontal)
			playerVerticalTEXT := strconv.Itoa(playerVertical)
			cameraZoomTEXT := fmt.Sprintf("%0f", camera.Zoom)
			chooseCharacterTEXT := strconv.Itoa(chooseCharacter)
			chooseMenuOptionTEXT := strconv.Itoa(chooseMenuOption)
			playerCoinsTEXT := strconv.Itoa(playerCoins)
			killzCountTEXT := strconv.Itoa(killzCount)
			totalKillzTEXT := strconv.Itoa(totalKillz)
			introScreenOnTEXT := strconv.FormatBool(introScreenOn)
			menuOnTEXT := strconv.FormatBool(menuOn)
			gamestartTEXT := strconv.FormatBool(gameStartOn)
			menuSelectBoxYTEXT := fmt.Sprint(menuSelectBoxY)
			teleportPostion1TEXT := strconv.Itoa(teleportPostion1)
			teleportPostion2TEXT := strconv.Itoa(teleportPostion2)
			levelEndMenuCountTEXT := strconv.Itoa(levelEndMenuCount)
			menuSelectNumberTEXT := strconv.Itoa(menuSelectNumber)
			fallingBlocksOnTEXT := strconv.FormatBool(fallingBlocksOn)
			teleportActiveTEXT := strconv.FormatBool(teleportActive)
			introStoryScreenOnTEXT := strconv.FormatBool(introStoryScreenOn)
			enemy4SpecialOnTEXT := strconv.FormatBool(enemy4SpecialOn)
			noiseLinesScreenOnTEXT := strconv.FormatBool(noiseLinesScreenOn)
			noiseLineLR1TEXT := strconv.FormatBool(noiseLineLR1)
			powerUpCollectedTEXT := strconv.FormatBool(powerUpCollected)
			powerUpCurrentActiveTEXT := strconv.Itoa(powerUpCurrentActive)
			playerHPTEXT := strconv.Itoa(playerHP)
			playerHPmaxTEXT := strconv.Itoa(playerHPmax)
			enemyPositionCheck := enemyPosition1MAP[0]
			enemyPositionCheckTEXT := strconv.Itoa(enemyPositionCheck)
			enemyPositionCheck2 := enemyPosition2MAP[0]
			enemyPositionCheck2TEXT := strconv.Itoa(enemyPositionCheck2)
			currentMiniBossBlockTEXT := strconv.Itoa(currentMiniBossBlock)
			shopMenuCountTEXT := strconv.Itoa(shopMenuCount)
			activeSpecialCompleteTEXT := strconv.FormatBool(activeSpecialComplete)
			tornadoLBlockTEXT := strconv.Itoa(tornadoLBlock)
			tornadoRBlockTEXT := strconv.Itoa(tornadoRBlock)
			horizplatlrTEXT := strconv.FormatBool(horizplatlr)
			horizplatHTEXT := strconv.Itoa(horizplatH)
			horizplatVTEXT := strconv.Itoa(horizplatV)
			miniBossHPTEXT := strconv.Itoa(miniBossHP)

			rl.DrawText(playerCurrentBlockTEXT, screenW-290, 10, 10, rl.White)
			rl.DrawText("playerCurrentBlock", screenW-200, 10, 10, rl.White)
			rl.DrawText(playerHorizontalTEXT, screenW-290, 20, 10, rl.White)
			rl.DrawText("playerHorizontal", screenW-200, 20, 10, rl.White)
			rl.DrawText(playerVerticalTEXT, screenW-290, 30, 10, rl.White)
			rl.DrawText("playerVertical", screenW-200, 30, 10, rl.White)
			rl.DrawText(cameraZoomTEXT, screenW-290, 40, 10, rl.White)
			rl.DrawText("cameraZoom", screenW-200, 40, 10, rl.White)
			rl.DrawText(chooseCharacterTEXT, screenW-290, 50, 10, rl.White)
			rl.DrawText("chooseCharacter", screenW-200, 50, 10, rl.White)
			rl.DrawText(dinoType, screenW-290, 60, 10, rl.White)
			rl.DrawText("chooseCharacter", screenW-200, 60, 10, rl.White)
			rl.DrawText(chooseMenuOptionTEXT, screenW-290, 70, 10, rl.White)
			rl.DrawText("chooseMenuOption", screenW-200, 70, 10, rl.White)
			rl.DrawText(playerCoinsTEXT, screenW-290, 80, 10, rl.White)
			rl.DrawText("playerCoins", screenW-200, 80, 10, rl.White)
			rl.DrawText(killzCountTEXT, screenW-290, 90, 10, rl.White)
			rl.DrawText("killzCount", screenW-200, 90, 10, rl.White)
			rl.DrawText(totalKillzTEXT, screenW-290, 100, 10, rl.White)
			rl.DrawText("totalKillz", screenW-200, 100, 10, rl.White)
			rl.DrawText(introScreenOnTEXT, screenW-290, 110, 10, rl.White)
			rl.DrawText("introScreenOn", screenW-200, 110, 10, rl.White)
			rl.DrawText(menuOnTEXT, screenW-290, 120, 10, rl.White)
			rl.DrawText("menuOn", screenW-200, 120, 10, rl.White)
			rl.DrawText(gamestartTEXT, screenW-290, 130, 10, rl.White)
			rl.DrawText("gameStart", screenW-200, 130, 10, rl.White)
			rl.DrawText(menuSelectBoxYTEXT, screenW-290, 140, 10, rl.White)
			rl.DrawText("menuSelectBoxY", screenW-200, 140, 10, rl.White)
			rl.DrawText(teleportPostion1TEXT, screenW-290, 150, 10, rl.White)
			rl.DrawText("teleportPostion1", screenW-200, 150, 10, rl.White)
			rl.DrawText(teleportPostion2TEXT, screenW-290, 160, 10, rl.White)
			rl.DrawText("teleportPostion2", screenW-200, 160, 10, rl.White)
			rl.DrawText(levelEndMenuCountTEXT, screenW-290, 170, 10, rl.White)
			rl.DrawText("levelEndMenuCount", screenW-200, 170, 10, rl.White)
			rl.DrawText(menuSelectNumberTEXT, screenW-290, 180, 10, rl.White)
			rl.DrawText("menuSelectNumber", screenW-200, 180, 10, rl.White)
			rl.DrawText(fallingBlocksOnTEXT, screenW-290, 190, 10, rl.White)
			rl.DrawText("fallingBlocksOn", screenW-200, 190, 10, rl.White)
			rl.DrawText(teleportActiveTEXT, screenW-290, 200, 10, rl.White)
			rl.DrawText("teleportActive", screenW-200, 200, 10, rl.White)
			rl.DrawText(introStoryScreenOnTEXT, screenW-290, 210, 10, rl.White)
			rl.DrawText("introStoryScreenOn", screenW-200, 210, 10, rl.White)
			rl.DrawText(enemy4SpecialOnTEXT, screenW-290, 220, 10, rl.White)
			rl.DrawText("enemy4SpecialOn", screenW-200, 220, 10, rl.White)
			rl.DrawText(noiseLinesScreenOnTEXT, screenW-290, 230, 10, rl.White)
			rl.DrawText("noiseLinesScreenOn", screenW-200, 230, 10, rl.White)
			rl.DrawText(noiseLineLR1TEXT, screenW-290, 240, 10, rl.White)
			rl.DrawText("noiseLineLR1", screenW-200, 240, 10, rl.White)
			rl.DrawText(powerUpCurrentActiveTEXT, screenW-290, 250, 10, rl.White)
			rl.DrawText("powerUpCurrentActive", screenW-200, 250, 10, rl.White)
			rl.DrawText(powerUpCollectedTEXT, screenW-290, 260, 10, rl.White)
			rl.DrawText("powerUpCollected", screenW-200, 260, 10, rl.White)
			rl.DrawText(playerHPTEXT, screenW-290, 270, 10, rl.White)
			rl.DrawText("playerHP", screenW-200, 270, 10, rl.White)
			rl.DrawText(playerHPmaxTEXT, screenW-290, 280, 10, rl.White)
			rl.DrawText("playerHPmax", screenW-200, 280, 10, rl.White)
			rl.DrawText(enemyPositionCheckTEXT, screenW-290, 290, 10, rl.White)
			rl.DrawText("enemyPositionCheck", screenW-200, 290, 10, rl.White)
			rl.DrawText(enemyPositionCheck2TEXT, screenW-290, 300, 10, rl.White)
			rl.DrawText("enemyPositionCheck2", screenW-200, 300, 10, rl.White)
			rl.DrawText(enemyMovementTEST, screenW-290, 310, 10, rl.White)
			rl.DrawText("enemyMovementTEST", screenW-200, 310, 10, rl.White)
			rl.DrawText(currentMiniBossBlockTEXT, screenW-290, 320, 10, rl.White)
			rl.DrawText("currentMiniBossBlock", screenW-200, 320, 10, rl.White)
			rl.DrawText(shopMenuCountTEXT, screenW-290, 330, 10, rl.White)
			rl.DrawText("shopMenuCount", screenW-200, 330, 10, rl.White)
			rl.DrawText(activeSpecialItem, screenW-290, 340, 10, rl.White)
			rl.DrawText("activeSpecialItem", screenW-200, 340, 10, rl.White)
			rl.DrawText(activeSpecialCompleteTEXT, screenW-290, 350, 10, rl.White)
			rl.DrawText("activeSpecialComplete", screenW-200, 350, 10, rl.White)
			rl.DrawText(tornadoLBlockTEXT, screenW-290, 360, 10, rl.White)
			rl.DrawText("tornadoLBlock", screenW-200, 360, 10, rl.White)
			rl.DrawText(tornadoRBlockTEXT, screenW-290, 370, 10, rl.White)
			rl.DrawText("tornadoRBlock", screenW-200, 370, 10, rl.White)
			rl.DrawText(horizplatlrTEXT, screenW-290, 380, 10, rl.White)
			rl.DrawText("horizplatlr", screenW-200, 380, 10, rl.White)
			rl.DrawText(horizplatHTEXT, screenW-290, 390, 10, rl.White)
			rl.DrawText("horizplatH", screenW-200, 390, 10, rl.White)
			rl.DrawText(horizplatVTEXT, screenW-290, 400, 10, rl.White)
			rl.DrawText("horizplatV", screenW-200, 400, 10, rl.White)
			rl.DrawText(miniBossHPTEXT, screenW-290, 410, 10, rl.White)
			rl.DrawText("miniBossHP", screenW-200, 410, 10, rl.White)

			// fps
			rl.DrawFPS(screenW-100, screenH-40)

		}

		rl.EndDrawing()
	} // end WindowShouldClose

	rl.CloseWindow()

	// MARK: unload textures audio
	rl.UnloadTexture(imgs)
	rl.UnloadTexture(backgroundTexture)
	rl.UnloadTexture(fogImage)
	rl.UnloadMusicStream(introTune)
	rl.UnloadMusicStream(levelTune)

} // end func main

// random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}

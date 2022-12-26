package wiseoldman

import "time"

type PlayerType string

const (
	PlayerTypeUnknown    PlayerType = "unknown"
	PlayerTypeRegular    PlayerType = "regular"
	PlayerTypeIronman    PlayerType = "ironman"
	PlayerTypeHardcore   PlayerType = "hardcore"
	PlayerTypeUltimate   PlayerType = "ultimate"
	PlayerTypeFreshStart PlayerType = "fresh_start"
)

type PlayerBuild string

const (
	PlayerBuildMain   PlayerBuild = "main"
	PlayerBuildF2P    PlayerBuild = "f2p"
	PlayerBuildLvl3   PlayerBuild = "lvl3"
	PlayerBuildZerker PlayerBuild = "zerker"
	PlayerBuildDef1   PlayerBuild = "def1"
	PlayerBuildHP10   PlayerBuild = "hp10"
)

type Skill struct {
	Metric     string  `json:"metric"`
	Experience int     `json:"experience"`
	Rank       int     `json:"rank"`
	Level      int     `json:"level"`
	EHP        float64 `json:"ehp"`
}

type Skills struct {
	Overall      Skill `json:"overall"`
	Attack       Skill `json:"attack"`
	Defence      Skill `json:"defence"`
	Strength     Skill `json:"strength"`
	Hitpoints    Skill `json:"hitpoints"`
	Ranged       Skill `json:"ranged"`
	Prayer       Skill `json:"prayer"`
	Magic        Skill `json:"magic"`
	Cooking      Skill `json:"cooking"`
	Woodcutting  Skill `json:"woodcutting"`
	Fletching    Skill `json:"fletching"`
	Fishing      Skill `json:"fishing"`
	Firemaking   Skill `json:"firemaking"`
	Crafting     Skill `json:"crafting"`
	Smithing     Skill `json:"smithing"`
	Mining       Skill `json:"mining"`
	Herblore     Skill `json:"herblore"`
	Agility      Skill `json:"agility"`
	Thieving     Skill `json:"thieving"`
	Slayer       Skill `json:"slayer"`
	Farming      Skill `json:"farming"`
	Runecrafting Skill `json:"runecrafting"`
	Hunter       Skill `json:"hunter"`
	Construction Skill `json:"construction"`
}

type Boss struct {
	Metric string  `json:"metric"`
	Kills  int     `json:"kills"`
	Rank   int     `json:"rank"`
	EHB    float64 `json:"ehb"`
}

type Bosses struct {
	AbyssalSire                  Boss `json:"abyssal_sire"`
	AlchemicalHydra              Boss `json:"alchemical_hydra"`
	BarrowsChests                Boss `json:"barrows_chests"`
	Bryophyta                    Boss `json:"bryophyta"`
	Callisto                     Boss `json:"callisto"`
	Cerberus                     Boss `json:"cerberus"`
	ChambersOfXeric              Boss `json:"chambers_of_xeric"`
	ChambersOfXericChallengeMode Boss `json:"chambers_of_xeric_challenge_mode"`
	ChaosElemental               Boss `json:"chaos_elemental"`
	ChaosFanatic                 Boss `json:"chaos_fanatic"`
	CommanderZilyana             Boss `json:"commander_zilyana"`
	CorporealBeast               Boss `json:"corporeal_beast"`
	CrazyArchaeologist           Boss `json:"crazy_archaeologist"`
	DagannothPrime               Boss `json:"dagannoth_prime"`
	DagannothRex                 Boss `json:"dagannoth_rex"`
	DagannothSupreme             Boss `json:"dagannoth_supreme"`
	DerangedArchaeologist        Boss `json:"deranged_archaeologist"`
	GeneralGraardor              Boss `json:"general_graardor"`
	GiantMole                    Boss `json:"giant_mole"`
	GrotesqueGuardians           Boss `json:"grotesque_guardians"`
	Hespori                      Boss `json:"hespori"`
	KalphiteQueen                Boss `json:"kalphite_queen"`
	KingBlackDragon              Boss `json:"king_black_dragon"`
	Kraken                       Boss `json:"kraken"`
	Kreearra                     Boss `json:"kreearra"`
	KrilTsutsaroth               Boss `json:"kril_tsutsaroth"`
	Mimic                        Boss `json:"mimic"`
	Nex                          Boss `json:"nex"`
	Nightmare                    Boss `json:"nightmare"`
	PhosanisNightmare            Boss `json:"phosanis_nightmare"`
	Obor                         Boss `json:"obor"`
	Sarachnis                    Boss `json:"sarachnis"`
	Scorpia                      Boss `json:"scorpia"`
	Skotizo                      Boss `json:"skotizo"`
	Tempoross                    Boss `json:"tempoross"`
	TheGauntlet                  Boss `json:"the_gauntlet"`
	TheCorruptedGauntlet         Boss `json:"the_corrupted_gauntlet"`
	TheatreOfBlood               Boss `json:"theatre_of_blood"`
	TheatreOfBloodHardMode       Boss `json:"theatre_of_blood_hard_mode"`
	ThermonuclearSmokeDevil      Boss `json:"thermonuclear_smoke_devil"`
	TombsOfAmascut               Boss `json:"tombs_of_amascut"`
	TombsOfAmascutExpert         Boss `json:"tombs_of_amascut_expert"`
	TzKalZuk                     Boss `json:"tzkal_zuk"`
	TzTokJad                     Boss `json:"tztok_jad"`
	Venenatis                    Boss `json:"venenatis"`
	Vetion                       Boss `json:"vetion"`
	Vorkath                      Boss `json:"vorkath"`
	Wintertodt                   Boss `json:"wintertodt"`
	Zalcano                      Boss `json:"zalcano"`
	Zulrah                       Boss `json:"zulrah"`
}

type Activity struct {
	Metric string `json:"metric"`
	Score  int    `json:"score"`
	Rank   int    `json:"rank"`
}

type Activities struct {
	LeaguePoints        Activity `json:"league_points"`
	BountyHunterHunter  Activity `json:"bounty_hunter_hunter"`
	BountyHunterRogue   Activity `json:"bounty_hunter_rogue"`
	ClueScrollsAll      Activity `json:"clue_scrolls_all"`
	ClueScrollsBeginner Activity `json:"clue_scrolls_beginner"`
	ClueScrollsEasy     Activity `json:"clue_scrolls_easy"`
	ClueScrollsMedium   Activity `json:"clue_scrolls_medium"`
	ClueScrollsHard     Activity `json:"clue_scrolls_hard"`
	ClueScrollsElite    Activity `json:"clue_scrolls_elite"`
	ClueScrollsMaster   Activity `json:"clue_scrolls_master"`
	LastManStanding     Activity `json:"last_man_standing"`
	PVPArena            Activity `json:"pvp_arena"`
	SoulWarsZeal        Activity `json:"soul_wars_zeal"`
	GuardiansOfTheRift  Activity `json:"guardians_of_the_rift"`
}

type ComputedEHP struct {
	Metric string  `json:"metric"`
	Value  float64 `json:"value"`
	Rank   int     `json:"rank"`
}

type ComputedEHB struct {
	Metric string  `json:"metric"`
	Value  float64 `json:"value"`
	Rank   int     `json:"rank"`
}

type Computed struct {
	EHP ComputedEHP `json:"ehp"`
	EHB ComputedEHB `json:"ehb"`
}

type SnapshotDataValues struct {
	Skills     Skills     `json:"skills"`
	Bosses     Bosses     `json:"bosses"`
	Activities Activities `json:"activities"`
	Computed   Computed   `json:"computed"`
}

type Snapshot struct {
	ID         int                `json:"id"`
	PlayerID   int                `json:"playerId"`
	CreatedAt  time.Time          `json:"createdAt"`
	ImportedAt time.Time          `json:"importedAt"`
	Data       SnapshotDataValues `json:"data"`
}

type Player struct {
	ID             int         `json:"id"`
	Username       string      `json:"username"`
	DisplayName    string      `json:"displayName"`
	Type           PlayerType  `json:"type"`
	Build          PlayerBuild `json:"build"`
	Country        string      `json:"country"`
	Flagged        bool        `json:"flagged"`
	EXP            int         `json:"exp"`
	EHP            float64     `json:"ehp"`
	EHB            float64     `json:"ehb"`
	TTM            float64     `json:"ttm"`
	TT200M         float64     `json:"tt200m"`
	RegisteredAt   time.Time   `json:"registeredAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
	LastChangedAt  time.Time   `json:"lastChangedAt"`
	LastImportedAt time.Time   `json:"lastImportedAt"`
}

type PlayerDetails struct {
	*Player
	CombatLevel    int      `json:"combatLevel"`
	LatestSnapshot Snapshot `json:"latestSnapshot"`
}

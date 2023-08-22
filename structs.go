package rplpa

import (
	"time"
)

// Replay is the Parsed replay.
type Replay struct {
	PlayMode     int8
	OsuVersion   int32
	BeatmapMD5   string
	Username     string
	ReplayMD5    string
	Count300     uint16
	Count100     uint16
	Count50      uint16
	CountGeki    uint16
	CountKatu    uint16
	CountMiss    uint16
	Score        int32
	MaxCombo     uint16
	Fullcombo    bool
	Mods         uint32
	LifebarGraph []LifeBarGraph
	Timestamp    time.Time
	ReplayData   []*ReplayData
	ScoreID      int64 // idk if it's the scoreid, maybe it is maybe not.
	ScoreInfo    ScoreInfo
}

type ScoreInfo struct {
	Mods              []*ModInfo
	Statistics        *Statistics
	MaximumStatistics []*MaximumStatistics
}

type ModInfo struct {
	Acronym         string                 `json:"acronym"`
	Settings        map[string]interface{} `json:"settings,omitempty"`
	SpeedChange     string                 `json:"speed_change,omitempty"`
	ApproachRate    string                 `json:"approach_rate,omitempty"`
	ExtendedLimits  string                 `json:"extended_limits,omitempty"`
	ClassicNoteLock string                 `json:"classic_note_lock,omitempty"`
}

type Statistics struct {
	Miss          float64
	Great         float64
	SmallTickHit  float64
	LargeTickMiss float64
	SmallBonus    float64
	Ok            float64
	SmallTickMiss float64
	LargeTickHit  float64
	IgnoreMiss    float64
}

type MaximumStatistics struct {
	*Statistics
	LargeBonus float64
}

// ReplayData is the Parsed Compressed Replay data.
type ReplayData struct {
	Time       int64
	MouseX     float32
	MouseY     float32
	KeyPressed *KeyPressed
}

// KeyPressed is the Parsed Compressed KeyPressed.
type KeyPressed struct {
	LeftClick  bool
	RightClick bool
	Key1       bool
	Key2       bool
	Smoke      bool
}

// LifeBarGraph is the Bar under the Score stuff.
type LifeBarGraph struct {
	Time int32
	HP   float32
}

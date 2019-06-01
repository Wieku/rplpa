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

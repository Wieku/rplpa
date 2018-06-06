package rplpa

import (
	"io/ioutil"
	"testing"
)

func TestParseReplay(t *testing.T) {
	b, err := ioutil.ReadFile("data/replay1.osr")
	if err != nil {
		t.Error("Could not read replay, Doesn't exists?")
	}
	p, err := ParseReplay(b)
	if err != nil {
		t.Error("Could not parse replay", err)
	}
	if p != nil {
		t.Log("Playmode: ", p.PlayMode)
		t.Log("OsuVersion: ", p.OsuVersion)
		t.Log("BeatmapMD5: ", p.BeatmapMD5)
		t.Log("Username: ", p.Username)
		t.Log("ReplayMD5: ", p.ReplayMD5)
		t.Log("Count300: ", p.Count300)
		t.Log("Count100: ", p.Count100)
		t.Log("Count50: ", p.Count50)
		t.Log("CountGeki: ", p.CountGeki)
		t.Log("CountKatu: ", p.CountKatu)
		t.Log("CountMiss: ", p.CountMiss)
		t.Log("Score: ", p.Score)
		t.Log("MaxCombo: ", p.MaxCombo)
		t.Log("Fullcombo: ", p.Fullcombo)
		t.Log("Mods: ", p.Mods)
		t.Log("LifebarGraph: ", p.LifebarGraph)
		t.Log("Timestamp: ", p.Timestamp)
		for i := 0; i < len(p.ReplayData); i++ {
			t.Log("ReplayData: ", p.ReplayData[i])
			t.Log("KeyPressed: ", p.ReplayData[i].KeyPressed)
		}
	}
}

func TestParseCompressed(t *testing.T) {
	b, err := ioutil.ReadFile("data/replay3_raw.bin")
	if err != nil {
		t.Error("Could not read replaydata, Doesn't exists?")
	}
	p, err := ParseCompressed(b)
	if err != nil {
		t.Error("Could not parse replaydata", err)
	}
	for i := 0; i < len(p); i++ {
		t.Log("ReplayData: ", p[i])
		t.Log("KeyPressed: ", p[i].KeyPressed)
	}
}

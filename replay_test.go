package rplpa

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParseStableReplay(t *testing.T) {
	parseTest(t, "data/replay1.osr")
}

func TestParseLazerReplay(t *testing.T) {
	parseTest(t, "data/lazer.osr")
}

func parseTest(t *testing.T, filename string) {
	b, err := os.ReadFile(filename)
	if err != nil {
		t.Error("Could not read replay, Doesn't exist?")
	}

	p, err := ParseReplay(b)
	if err != nil {
		t.Error("Could not parse replay", err)
	}

	if p != nil {
		t.Log("PlayMode: ", p.PlayMode)
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
		t.Log("ScoreID: ", p.ScoreID)
		t.Log("InputEvents", len(p.ReplayData))

		firstTime := p.ReplayData[0].Time
		lastTime := p.ReplayData[0].Time

		for i, d := range p.ReplayData {
			if i == 0 {
				continue
			}

			lastTime += d.Time

			if lastTime < firstTime {
				firstTime = lastTime
			}
		}

		t.Log("First delta ms:", p.ReplayData[0].Time)
		t.Log("Lowest input time ms:", firstTime)
		t.Log("Last input time ms:", lastTime)
		t.Log("Replay duration:", lastTime-firstTime)

		if p.ScoreInfo != nil {
			t.Log("ScoreInfo Mods: ", len(p.ScoreInfo.Mods))

			if len(p.ScoreInfo.Mods) > 0 {
				for _, m := range p.ScoreInfo.Mods {
					t.Log("ScoreInfo Mod: ", *m)
				}
			}
			if p.ScoreInfo.Statistics != nil {
				t.Log("ScoreInfo Statistics: ", p.ScoreInfo.Statistics)
			}

			if p.ScoreInfo.MaximumStatistics != nil {
				t.Log("ScoreInfo MaximumStatistics: ", p.ScoreInfo.MaximumStatistics)
			}
		} else {
			t.Log("ScoreInfo is nil due to EOF (stable play)")
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
	if len(p) > 20 {
		t.Log("ReplayData: true")
	} else {
		t.Error("Error while parsing compressed.")
	}
}

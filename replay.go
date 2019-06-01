package rplpa

import (
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/itchio/lzma"
)

// NewReplay returns an Empty Replay
func NewReplay() *Replay {
	return &Replay{}
}

// ParseReplay parses a Replay and returns a *Replay
func ParseReplay(file []byte) (r *Replay, err error) {
	var LifeBarRaw string
	var ts int64
	var slength int32
	var compressedReplay []byte

	b := bytes.NewBuffer(file)
	r = NewReplay()
	r.PlayMode, err = rInt8(b)
	if err != nil {
		return
	}
	r.OsuVersion, err = rInt32(b)
	if err != nil {
		return
	}
	r.BeatmapMD5, err = rBString(b)
	if err != nil {
		return
	}
	r.Username, err = rBString(b)
	if err != nil {
		return
	}
	r.ReplayMD5, err = rBString(b)
	if err != nil {
		return
	}
	r.Count300, err = rUInt16(b)
	if err != nil {
		return
	}
	r.Count100, err = rUInt16(b)
	if err != nil {
		return
	}
	r.Count50, err = rUInt16(b)
	if err != nil {
		return
	}
	r.CountGeki, err = rUInt16(b)
	if err != nil {
		return
	}
	r.CountKatu, err = rUInt16(b)
	if err != nil {
		return
	}
	r.CountMiss, err = rUInt16(b)
	if err != nil {
		return
	}
	r.Score, err = rInt32(b)
	if err != nil {
		return
	}
	r.MaxCombo, err = rUInt16(b)
	if err != nil {
		return
	}
	r.Fullcombo, err = rBool(b)
	if err != nil {
		return
	}
	r.Mods, err = rUInt32(b)
	if err != nil {
		return
	}
	LifeBarRaw, err = rBString(b)
	if err != nil {
		return
	}
	r.LifebarGraph = parseLifebar(LifeBarRaw)
	ts, err = rInt64(b)
	if err != nil {
		return
	}
	r.Timestamp = timeFromTicks(ts)
	slength, err = rInt32(b)
	if err != nil {
		return
	}
	compressedReplay, err = rSlice(b, slength)
	if err != nil {
		return
	}
	r.ReplayData, err = ParseCompressed(compressedReplay)
	if err != nil {
		return
	}
	return
}

// https://stackoverflow.com/questions/33144967/what-is-the-c-sharp-datetimeoffset-equivalent-in-go/33161703#33161703

func timeFromTicks(ticks int64) time.Time {
	base := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	return time.Unix(ticks/10000000+base, ticks%10000000).UTC()
}

func parseLifebar(s string) []LifeBarGraph {
	var o []LifeBarGraph
	s = strings.Trim(s, ",")
	life := strings.Split(s, ",")
	for i := 0; i < len(life); i++ {
		y := strings.Split(life[i], "|")
		if len(y) < 2 {
			continue
		}
		f, err := strconv.ParseFloat(y[1], 32)
		if err != nil {
			continue
		}
		v, err := strconv.Atoi(y[0])
		o = append(o, LifeBarGraph{Time: int32(v), HP: float32(f)})
	}
	return o
}

// ParseCompressed parses a compressed replay, (ReplayData)
func ParseCompressed(file []byte) (d []*ReplayData, err error) {
	b := bytes.NewBuffer(file)
	r := lzma.NewReader(b)
	defer r.Close()

	x, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	s := strings.Trim(string(x), ",")

	sa := strings.Split(s, ",")

	for i := 0; i < len(sa); i++ {
		rd := sa[i]
		xd := strings.Split(rd, "|")
		if len(xd) < 4 {
			continue
		}
		var Time int
		var MouseX float64
		var MouseY float64
		var KPA int
		Time, err = strconv.Atoi(xd[0])
		if err != nil {
			return
		}
		MouseX, err = strconv.ParseFloat(xd[1], 32)
		if err != nil {
			return
		}
		MouseY, err = strconv.ParseFloat(xd[2], 32)
		if err != nil {
			return
		}
		KPA, err = strconv.Atoi(xd[3])
		if err != nil {
			return
		}
		KP := KeyPressed{
			LeftClick:  KPA&LEFTCLICK > 0,
			RightClick: KPA&RIGHTCLICK > 0,
			Key1:       KPA&KEY1 > 0,
			Key2:       KPA&KEY2 > 0,
			Smoke:      KPA&SMOKE > 0,
		}
		rdata := ReplayData{
			Time:       int64(Time),
			MouseX:     float32(MouseX),
			MouseY:     float32(MouseY),
			KeyPressed: &KP,
		}
		d = append(d, &rdata)
	}

	return
}

package mlib

import (
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("**NewMusicManager** failed!")
	}

	if mm.Len() != 0 {
		t.Error("**NewMusicManager** failed! ** not empty**")
	}

	m0 := &MusicEntry{
		"1", "My Heart Will Go On", "Celion Dion", "http://yinyueshiting.baidu.com/data2/music/199a6d5133bb1b830e98405400e34331/262190481/262189007255600128.mp3?xcode=ce70d72c84e37fa426bdb4a8ac0cec0f", "MP3"}

	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("**MusicManager.Add()** failed!")
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("**MusicManager.Find()** failed.")
	}
	if m.Id != m0.Id || m.Artist != m0.Artist ||
		m.Name != m0.Name ||
		m.Source != m0.Source || m.Type != m0.Type {
		t.Error("**MusicManager.Find()** failed. **Found item mismatch.**")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("**MusicManager.Get()** failed.", err)
	}

	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("**MusicManager.Remove()** failed.", err)
	}
}

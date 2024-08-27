package main

import (
	"slices"
)

type MangaStatus int

var available_manga_id uint = 0

const (
	Dropped MangaStatus = iota
	Acquiring
	Completed
)

// A Manga contains all the information of a Manga Series.
type Manga struct {
	Id               uint
	Series           string
	NVolumes         int
	ShelfNo          int
	CollectionStatus MangaStatus
	VolumesInv       []int
	Authors          []string
}

// Create a Manga struct.
// You only need to provide series, n_volumes, and authors to initialize the struct
func CreateManga(series string, n_volumes int, authors []string) Manga {
	available_manga_id += 1
	return Manga{
		Id:               available_manga_id - 1,
		Series:           series,
		NVolumes:         n_volumes,
		Authors:          authors,
		ShelfNo:          0,
		CollectionStatus: Acquiring,
		VolumesInv:       []int{},
	}
}

func (m *Manga) AddVolume(vol_no int) {
	// Base cases
	// Empty slice
	if len(m.VolumesInv) < 1 {
		m.VolumesInv = append(m.VolumesInv, vol_no)
		return
	}

	// We have 1 elements
	if len(m.VolumesInv) < 2 {
		if m.VolumesInv[0] < vol_no {
			m.VolumesInv = append(m.VolumesInv, vol_no)
		} else {
			// Append at the start
			m.VolumesInv = append([]int{vol_no}, m.VolumesInv...)
		}
		return
	}

	// We have 2 elements
	if len(m.VolumesInv) < 3 {
		if m.VolumesInv[0] > vol_no {
			m.VolumesInv = slices.Insert(m.VolumesInv, 0, vol_no)
		} else if m.VolumesInv[1] < vol_no {
			m.VolumesInv = slices.Insert(m.VolumesInv, 2, vol_no)
		} else {
			m.VolumesInv = slices.Insert(m.VolumesInv, 1, vol_no)
		}
		return
	}

	// Assuming that m.VolumesInv has n > 2 elements
	for i := 1; i < len(m.VolumesInv)-1; i++ {
		if vol_no < m.VolumesInv[i+1] && vol_no > m.VolumesInv[i-1] {
			m.VolumesInv = slices.Insert(m.VolumesInv, i+1, vol_no)
			return
		}
	}
	m.VolumesInv = append(m.VolumesInv, vol_no)
}

// Used when you Drop a Manga series
func (m *Manga) Drop() {
	m.CollectionStatus = Dropped
}

// Used when you decide to start collecting a series. Planning on using this
// when you start collecting a manga once it was originally dropped
func (m *Manga) Collect() {
	m.CollectionStatus = Acquiring
}

// Set our Manga.CollectionStatus to Complete. Note that if Manga.VolumesInv
// is empty, we instead fill up the remaining slots in Manga.VolumesInv.
func (m *Manga) Complete() {
	if len(m.VolumesInv) < m.NVolumes {
		// Cheap and Dirty
		// We clear the damn thing instead of constructing an algorithm to
		// fill up the unused bits and pieces
		// TODO: Fix this
		m.VolumesInv = m.VolumesInv[:0]
		for i := 1; i <= m.NVolumes; i++ {
			m.VolumesInv = append(m.VolumesInv, i)
		}
	}
	m.CollectionStatus = Completed
}

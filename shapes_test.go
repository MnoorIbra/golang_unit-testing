package main

import (
	"math"
	"testing"
)

func TestHitungLuasPersegi(t *testing.T) {
	sisi := 5.0
	hasilSeharusnya := 25.0

	hasil := HitungLuasPersegi(sisi)

	if hasil != hasilSeharusnya {
		t.Errorf("Hasilnya salah. Harusnya %.2f, got %.2f", hasilSeharusnya, hasil)
	}
}

func TestHitungLuasLingkaran(t *testing.T) {
	radius := 3.0
	hasilSeharusnya := math.Pi * radius * radius

	hasil := HitungLuasLingkaran(radius)

	if hasil != hasilSeharusnya {
		t.Errorf("Hasilnya salah. Harusnya %.2f, got %.2f", hasilSeharusnya, hasil)
	}
}

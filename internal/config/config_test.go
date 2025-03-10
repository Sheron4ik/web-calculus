package config

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("Default values", func(t *testing.T) {
		os.Unsetenv("PORT")
		os.Unsetenv("TIME_ADDITION_MS")
		os.Unsetenv("TIME_SUBTRACTION_MS")
		os.Unsetenv("TIME_MULTIPLICATIONS_MS")
		os.Unsetenv("TIME_DIVISIONS_MS")
		os.Unsetenv("COMPUTING_POWER")

		cfg := New()

		if cfg.Port != "8080" {
			t.Errorf("Ожидалось Port=8080, получено %s", cfg.Port)
		}
		if cfg.TimeAddMs != "1000" {
			t.Errorf("Ожидалось TimeAddMs=1000, получено %s", cfg.TimeAddMs)
		}
		if cfg.TimeSubMs != "1000" {
			t.Errorf("Ожидалось TimeSubMs=1000, получено %s", cfg.TimeSubMs)
		}
		if cfg.TimeMulMs != "2000" {
			t.Errorf("Ожидалось TimeMulMs=2000, получено %s", cfg.TimeMulMs)
		}
		if cfg.TimeDivMs != "2500" {
			t.Errorf("Ожидалось TimeDivMs=2500, получено %s", cfg.TimeDivMs)
		}
		if cfg.ComputingPower != "3" {
			t.Errorf("Ожидалось ComputingPower=3, получено %s", cfg.ComputingPower)
		}
	})

	t.Run("Environment variables set", func(t *testing.T) {
		os.Setenv("PORT", "9090")
		os.Setenv("TIME_ADDITION_MS", "500")
		os.Setenv("TIME_SUBTRACTION_MS", "600")
		os.Setenv("TIME_MULTIPLICATIONS_MS", "1000")
		os.Setenv("TIME_DIVISIONS_MS", "1500")
		os.Setenv("COMPUTING_POWER", "5")

		cfg := New()

		if cfg.Port != "9090" {
			t.Errorf("Ожидалось Port=9090, получено %s", cfg.Port)
		}
		if cfg.TimeAddMs != "500" {
			t.Errorf("Ожидалось TimeAddMs=500, получено %s", cfg.TimeAddMs)
		}
		if cfg.TimeSubMs != "600" {
			t.Errorf("Ожидалось TimeSubMs=600, получено %s", cfg.TimeSubMs)
		}
		if cfg.TimeMulMs != "1000" {
			t.Errorf("Ожидалось TimeMulMs=1000, получено %s", cfg.TimeMulMs)
		}
		if cfg.TimeDivMs != "1500" {
			t.Errorf("Ожидалось TimeDivMs=1500, получено %s", cfg.TimeDivMs)
		}
		if cfg.ComputingPower != "5" {
			t.Errorf("Ожидалось ComputingPower=5, получено %s", cfg.ComputingPower)
		}

		os.Unsetenv("PORT")
		os.Unsetenv("TIME_ADDITION_MS")
		os.Unsetenv("TIME_SUBTRACTION_MS")
		os.Unsetenv("TIME_MULTIPLICATIONS_MS")
		os.Unsetenv("TIME_DIVISIONS_MS")
		os.Unsetenv("COMPUTING_POWER")
	})

	t.Run("Mixed values", func(t *testing.T) {
		os.Setenv("PORT", "9090")
		os.Setenv("TIME_ADDITION_MS", "500")

		cfg := New()

		if cfg.Port != "9090" {
			t.Errorf("Ожидалось Port=9090, получено %s", cfg.Port)
		}
		if cfg.TimeAddMs != "500" {
			t.Errorf("Ожидалось TimeAddMs=500, получено %s", cfg.TimeAddMs)
		}
		if cfg.TimeSubMs != "1000" {
			t.Errorf("Ожидалось TimeSubMs=1000, получено %s", cfg.TimeSubMs)
		}
		if cfg.TimeMulMs != "2000" {
			t.Errorf("Ожидалось TimeMulMs=2000, получено %s", cfg.TimeMulMs)
		}
		if cfg.TimeDivMs != "2500" {
			t.Errorf("Ожидалось TimeDivMs=2500, получено %s", cfg.TimeDivMs)
		}
		if cfg.ComputingPower != "3" {
			t.Errorf("Ожидалось ComputingPower=3, получено %s", cfg.ComputingPower)
		}

		os.Unsetenv("PORT")
		os.Unsetenv("TIME_ADDITION_MS")
	})
}

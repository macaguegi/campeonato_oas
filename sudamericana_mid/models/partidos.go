package models

import ("time")

type Partido struct {
	Id      int
	Fecha   time.Time
	Ciudad	int
}
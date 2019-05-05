package models

import ()

type DetallePartido struct {
	Id             int
	GolesLocal     int
	GolesVisitante int
	Equipo1Id      int
	Equipo2Id      int
	IdPartido      *Partido
}
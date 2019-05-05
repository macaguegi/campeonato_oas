
package models

import()

type Ciudad struct {
	Id     int
	Nombre string
	PaisId *Pais
}
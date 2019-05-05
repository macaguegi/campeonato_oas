package controllers

import (
	"github.com/macaguegi/sudamericana_mid/models"
	"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
	"strconv"
)

// EquiposController operaciones para todos los equipos
type EquiposController struct {
	beego.Controller
}


// ConsultarEquipos ...
// @Title ConsultarEquipos
// @Description Consultar los equipos 
// @Success 200 {object} models.Equipo
// @Failure 403 body is empty
// @router /consultarEquipos
func (c *EquiposController) ConsultarEquipos() {
	var equipos []models.Equipo
	if err := getJson("http://localhost:8084/v1/equipo", &equipos); err == nil {
		c.Data["json"] = equipos
	} else {
		fmt.Println(err)	
	}
	c.ServeJSON()
}


// InsertarEquipo ...
// @Title InsertarEquipo
// @Description Create equipo
// @Param	body		body 	models.Equipo	true		"body for Equipo content"
// @Success 201 {int} models.Equipo
// @Failure 403 body is empty
// @router /InsertarEquipo [post]
func (c *EquiposController) InsertarEquipo()(idRespuesta int, err error) {
	var equipo_temp models.Equipo
	var equipo_res models.Equipo
	var ciudad_temp models.Ciudad

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &equipo_temp)

		if validate := getJson("http://localhost:8081/v1/ciudad/"+strconv.Itoa(equipo_temp.CiudadId), &ciudad_temp); validate == nil {
			if err = sendJson("http://localhost:8084/v1/equipo/","POST",&equipo_res,&equipo_temp) ; err==nil{
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = equipo_res
			} else {
				c.Data["json"] = err.Error()
				fmt.Println("error es: ", err)
			}
		} else {
			fmt.Println(validate)	
		}

		if err != nil {
			beego.Error(err)
			return idRespuesta, err
		}

		c.ServeJSON()
		return idRespuesta, err
	}
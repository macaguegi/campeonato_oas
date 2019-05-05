package controllers

import (
	"github.com/macaguegi/sudamericana_mid/models"
	"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
	"strconv"
	//"time"
)

// PartidosController operaciones para todos los partidos
type PartidosController struct {
	beego.Controller
}


// ConsultarPartidos ...
// @Title ConsultarPartidos
// @Description Consultar los partidos 
// @Success 200 {object} models.DetallePartido
// @Failure 403 body is empty
// @router /consultarPartidos
func (c *PartidosController) ConsultarPartidos() {
	var partidos []models.DetallePartido
	if err := getJson("http://localhost:8085/v1/detalle_partido", &partidos); err == nil {
		c.Data["json"] = partidos
	} else {
		fmt.Println(err)	
	}
	c.ServeJSON()
}

// InsertarPartido ...
// @Title InsertarPartido
// @Description Create partido
// @Param	body		body 	models.Partido	true		"body for Partido content"
// @Success 201 {int} models.Partido
// @Failure 403 body is empty
// @router /InsertarPartido [post]
func (c *PartidosController) InsertarPartido()(idRespuesta int, err error) {
	var partido_temp models.Partido
	var partido_res models.Partido
	var ciudad_temp models.Ciudad

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &partido_temp)

		if validate := getJson("http://localhost:8081/v1/ciudad/"+strconv.Itoa(partido_temp.Ciudad), &ciudad_temp); validate == nil {
			if err = sendJson("http://localhost:8085/v1/equipo/","POST",&partido_res,&partido_temp) ; err==nil{
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = partido_res
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
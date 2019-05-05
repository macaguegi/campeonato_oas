package controllers

import (
	"github.com/macaguegi/sudamericana_mid/models"
	"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
)

// LugaresController operaciones para todos los lugares (Paises, Ciudades)
type LugaresController struct {
	beego.Controller
}


// ConsultarPaises ...
// @Title ConsultarPaises
// @Description Consultar los paises 
// @Success 200 {object} models.Pais
// @Failure 403 body is empty
// @router /consultarPaises
func (c *LugaresController) ConsultarPaises() {
	var paises []models.Pais
	if err := getJson("http://localhost:8081/v1/pais", &paises); err == nil {
		c.Data["json"] = paises
	} else {
		fmt.Println(err)	
	}
	c.ServeJSON()
}


// ConsultarCiudades ...
// @Title ConsultarCiudades
// @Description Consultar las ciudades 
// @Success 200 {object} models.Ciudad
// @Failure 403 body is empty
// @router /consultarCiudades
func (c *LugaresController) ConsultarCiudades() {
	var ciudades []models.Ciudad
	if err := getJson("http://localhost:8081/v1/ciudad", &ciudades); err == nil {
		c.Data["json"] = ciudades
	} else {
		fmt.Println(err)	
	}
	c.ServeJSON()
}


// InsertarPais ...
// @Title InsertarPais
// @Description Create pais
// @Param	body		body 	models.Pais	true		"body for Pais content"
// @Success 201 {int} models.Pais
// @Failure 403 body is empty
// @router /insertarPais [post]
func (c *LugaresController) InsertarPais()(idRespuesta int, err error) {
	var pais_temp models.Pais
	var pais_res models.Pais
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &pais_temp)

	if err != nil {
		beego.Error(err)
		return idRespuesta, err
	}

	if err = sendJson("http://localhost:8081/v1/pais/","POST",&pais_res,&pais_temp) ; err==nil{
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = pais_res
	} else {
		c.Data["json"] = err.Error()
		fmt.Println("error es: ", err)
	}

	if err != nil {
		beego.Error(err)
		return idRespuesta, err
	}

	c.ServeJSON()
	return idRespuesta, err
}

// InsertarCiudad ...
// @Title InsertarCiudad
// @Description Create ciudad
// @Param	body		body 	models.Ciudad	true		"body for Ciudad content"
// @Success 201 {int} models.Ciudad
// @Failure 403 body is empty
// @router /insertarCiudad [post]
func (c *LugaresController) InsertarCiudad()(idRespuesta int, err error) {
	var ciudad_temp models.Ciudad
	var ciudad_res models.Ciudad
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &ciudad_temp)


	if err = sendJson("http://localhost:8081/v1/ciudad/","POST",&ciudad_res,&ciudad_temp) ; err==nil{
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = ciudad_res
	} else {
		c.Data["json"] = err.Error()
		fmt.Println("error es: ", err)
	}

	if err != nil {
		beego.Error(err)
		return idRespuesta, err
	}

	c.ServeJSON()
	return idRespuesta, err
}
-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.1
-- PostgreSQL version: 10.0
-- Project Site: pgmodeler.io
-- Model Author: ---


-- Database creation must be done outside a multicommand file.
-- These commands were put in this file only as a convenience.
-- -- object: sudamericana_db | type: DATABASE --
-- -- DROP DATABASE IF EXISTS sudamericana_db;
-- CREATE DATABASE sudamericana_db
-- 	ENCODING = 'UTF8'
-- 	LC_COLLATE = 'Spanish_Spain.1252'
-- 	LC_CTYPE = 'Spanish_Spain.1252'
-- 	TABLESPACE = pg_default
-- 	OWNER = postgres;
-- -- ddl-end --
-- 

-- object: paises | type: SCHEMA --
-- DROP SCHEMA IF EXISTS paises CASCADE;
CREATE SCHEMA paises;
-- ddl-end --
ALTER SCHEMA paises OWNER TO postgres;
-- ddl-end --

-- object: partidos | type: SCHEMA --
-- DROP SCHEMA IF EXISTS partidos CASCADE;
CREATE SCHEMA partidos;
-- ddl-end --
ALTER SCHEMA partidos OWNER TO postgres;
-- ddl-end --

-- object: liga_sudamericana | type: SCHEMA --
-- DROP SCHEMA IF EXISTS liga_sudamericana CASCADE;
CREATE SCHEMA liga_sudamericana;
-- ddl-end --
ALTER SCHEMA liga_sudamericana OWNER TO postgres;
-- ddl-end --

-- object: equipos | type: SCHEMA --
-- DROP SCHEMA IF EXISTS equipos CASCADE;
CREATE SCHEMA equipos;
-- ddl-end --
ALTER SCHEMA equipos OWNER TO postgres;
-- ddl-end --

SET search_path TO pg_catalog,public,paises,partidos,liga_sudamericana,equipos;
-- ddl-end --

-- object: paises."Pais_id_seq" | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS paises."Pais_id_seq" CASCADE;
CREATE SEQUENCE paises."Pais_id_seq"
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
ALTER SEQUENCE paises."Pais_id_seq" OWNER TO postgres;
-- ddl-end --

-- object: paises.pais | type: TABLE --
-- DROP TABLE IF EXISTS paises.pais CASCADE;
CREATE TABLE paises.pais(
	id integer NOT NULL DEFAULT nextval('paises."Pais_id_seq"'::regclass),
	nombre character varying(20) NOT NULL,
	CONSTRAINT pk_pais PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON COLUMN paises.pais.id IS 'Identificador para cada país registrado en la Liga Sudamericana';
-- ddl-end --
COMMENT ON COLUMN paises.pais.nombre IS 'Nombre del País ';
-- ddl-end --
ALTER TABLE paises.pais OWNER TO postgres;
-- ddl-end --

-- object: paises."Ciudad_id_seq" | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS paises."Ciudad_id_seq" CASCADE;
CREATE SEQUENCE paises."Ciudad_id_seq"
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
ALTER SEQUENCE paises."Ciudad_id_seq" OWNER TO postgres;
-- ddl-end --

-- object: paises.ciudad | type: TABLE --
-- DROP TABLE IF EXISTS paises.ciudad CASCADE;
CREATE TABLE paises.ciudad(
	id integer NOT NULL DEFAULT nextval('paises."Ciudad_id_seq"'::regclass),
	nombre character varying(20) NOT NULL,
	pais_id integer NOT NULL,
	CONSTRAINT pk_ciudad PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE paises.ciudad IS 'Esta tabla contiene todas las ciudades pertenecientes a los paises de la liga sudamericana';
-- ddl-end --
COMMENT ON COLUMN paises.ciudad.id IS 'Identificador de la  ciudad';
-- ddl-end --
COMMENT ON COLUMN paises.ciudad.nombre IS 'Nombre de la ciudad';
-- ddl-end --
COMMENT ON CONSTRAINT pk_ciudad ON paises.ciudad  IS 'Llave primaria de la tabla Ciudad';
-- ddl-end --
ALTER TABLE paises.ciudad OWNER TO postgres;
-- ddl-end --

-- object: paises."Equipo_id_seq" | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS paises."Equipo_id_seq" CASCADE;
CREATE SEQUENCE paises."Equipo_id_seq"
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
ALTER SEQUENCE paises."Equipo_id_seq" OWNER TO postgres;
-- ddl-end --

-- object: equipos.equipo | type: TABLE --
-- DROP TABLE IF EXISTS equipos.equipo CASCADE;
CREATE TABLE equipos.equipo(
	id integer NOT NULL DEFAULT nextval('"Equipo_id_seq"'::regclass),
	nombre character varying(20),
	director_tecnico character varying(25) NOT NULL,
	ciudad_id integer NOT NULL,
	CONSTRAINT pk_equipo PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE equipos.equipo IS 'Tabla con la información de los equipos participantes';
-- ddl-end --
COMMENT ON COLUMN equipos.equipo.id IS 'Identificador del equipo ';
-- ddl-end --
COMMENT ON COLUMN equipos.equipo.nombre IS 'Nombre del equipo participante';
-- ddl-end --
COMMENT ON COLUMN equipos.equipo.director_tecnico IS 'Nombre del director tecnico del equipo';
-- ddl-end --
ALTER TABLE equipos.equipo OWNER TO postgres;
-- ddl-end --

-- object: partidos."Partido_id_seq" | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS partidos."Partido_id_seq" CASCADE;
CREATE SEQUENCE partidos."Partido_id_seq"
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
ALTER SEQUENCE partidos."Partido_id_seq" OWNER TO postgres;
-- ddl-end --

-- object: partidos.partido | type: TABLE --
-- DROP TABLE IF EXISTS partidos.partido CASCADE;
CREATE TABLE partidos.partido(
	id integer NOT NULL DEFAULT nextval('partidos."Partido_id_seq"'::regclass),
	estadio character varying(20) NOT NULL,
	fecha date NOT NULL,
	CONSTRAINT pk_partido PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE partidos.partido OWNER TO postgres;
-- ddl-end --

-- object: partidos.detalle_partido_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS partidos.detalle_partido_id_seq CASCADE;
CREATE SEQUENCE partidos.detalle_partido_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
ALTER SEQUENCE partidos.detalle_partido_id_seq OWNER TO postgres;
-- ddl-end --

-- object: partidos.detalle_partido | type: TABLE --
-- DROP TABLE IF EXISTS partidos.detalle_partido CASCADE;
CREATE TABLE partidos.detalle_partido(
	id integer NOT NULL DEFAULT nextval('partidos.detalle_partido_id_seq'::regclass),
	goles_local integer NOT NULL,
	goles_visitante integer NOT NULL,
	equipo1_id integer NOT NULL,
	equipo2_id integer NOT NULL,
	id_partido integer,
	CONSTRAINT pk_detalle_partido PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE partidos.detalle_partido OWNER TO postgres;
-- ddl-end --

-- object: liga_sudamericana."Posicion_id_seq" | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS liga_sudamericana."Posicion_id_seq" CASCADE;
CREATE SEQUENCE liga_sudamericana."Posicion_id_seq"
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
ALTER SEQUENCE liga_sudamericana."Posicion_id_seq" OWNER TO postgres;
-- ddl-end --

-- object: liga_sudamericana.posicion | type: TABLE --
-- DROP TABLE IF EXISTS liga_sudamericana.posicion CASCADE;
CREATE TABLE liga_sudamericana.posicion(
	id integer NOT NULL DEFAULT nextval('liga_sudamericana."Posicion_id_seq"'::regclass),
	partidos_jugados integer,
	partidos_ganados integer,
	partidos_empatados integer,
	partidos_perdidos integer,
	total_puntos integer NOT NULL,
	posicion integer NOT NULL,
	equipo_id integer NOT NULL,
	CONSTRAINT pk_posicion PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE liga_sudamericana.posicion OWNER TO postgres;
-- ddl-end --

-- object: fk__partido | type: CONSTRAINT --
-- ALTER TABLE partidos.detalle_partido DROP CONSTRAINT IF EXISTS fk__partido CASCADE;
ALTER TABLE partidos.detalle_partido ADD CONSTRAINT fk__partido FOREIGN KEY (id_partido)
REFERENCES partidos.partido (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: uq_detalle_partido | type: CONSTRAINT --
-- ALTER TABLE partidos.detalle_partido DROP CONSTRAINT IF EXISTS uq_detalle_partido CASCADE;
ALTER TABLE partidos.detalle_partido ADD CONSTRAINT uq_detalle_partido UNIQUE (id_partido);
-- ddl-end --

-- object: fk_ciudad_pais | type: CONSTRAINT --
-- ALTER TABLE paises.ciudad DROP CONSTRAINT IF EXISTS fk_ciudad_pais CASCADE;
ALTER TABLE paises.ciudad ADD CONSTRAINT fk_ciudad_pais FOREIGN KEY (pais_id)
REFERENCES paises.pais (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --



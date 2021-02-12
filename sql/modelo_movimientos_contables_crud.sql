

-- DROP SCHEMA IF EXISTS movimientos_contables CASCADE;
CREATE SCHEMA movimientos_contables;
-- ddl-end --


set SEARCH_PATH to pg_catalog,public,movimientos_contables;
-- ddl-end --

-- object: movimientos_contables.movimiento | type: TABLE --
-- DROP TABLE IF EXISTS movimientos_contables.movimiento CASCADE;
CREATE TABLE movimientos_contables.movimiento (
	id serial NOT NULL,
	tercero_id integer,
	cuenta_id text NOT NULL,
	nombre_cuenta text NOT NULL,
	tipo_movimiento_id integer NOT NULL,
	valor numeric(20,7) NOT NULL,
	descripcion character varying(250),
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	transaccion_id integer NOT NULL,
	CONSTRAINT pk_movimiento PRIMARY KEY (id)

);

-- ddl-end --
COMMENT ON TABLE movimientos_contables.movimiento IS 'Tabla que almacena la informacion de los movimientos contables asociados a una transaccion';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.movimiento.id IS 'Identificador del movimiento';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.movimiento.tercero_id IS 'Campo opcional que referencia al esquema terceros. Hace referencia al id  de la persona natural o juridica vinculada al movimiento';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.movimiento.cuenta_id IS 'Campo obligatorio que referencia al api plan_cuentas_mongo_crud. Hace referencia al id  o numero de cuenta vinculado al movimiento';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.movimiento.nombre_cuenta IS 'Campo obligatorio que indica el nombre de la cuenta vinculada al movimiento';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.movimiento.tipo_movimiento_id IS 'Campo obligatorio que referencia al esquema parametros a la tabla parametro. Hace referencia al id del tipo de movimiento asociado (debito o credita)';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.movimiento.valor IS 'Campo obligatorio que relaciona el valor asociado al movimiento';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.movimiento.descripcion IS 'Campo opcional que indica la descripcion del movimiento';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.movimiento.activo IS 'Campo para indicar el estado del registro';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.movimiento.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.movimiento.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.movimiento.transaccion_id IS 'Referencia foranea a la tabla transaccion';
-- ddl-end --
COMMENT ON CONSTRAINT pk_movimiento ON movimientos_contables.movimiento  IS 'Llave primaria de la tabla movimiento';
-- ddl-end --





-- object: movimientos_contables.transaccion | type: TABLE --
-- DROP TABLE IF EXISTS movimientos_contables.transaccion CASCADE;
CREATE TABLE movimientos_contables.transaccion (
	id serial NOT NULL,
	consecutivo_id integer NOT NULL,
	etiquetas json,
	descripcion character varying(250),
	error_transaccion character varying(250),
	estado_id integer NOT NULL,
	fecha_transaccion timestamp NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_transaccion PRIMARY KEY (id)

);

-- ddl-end --
COMMENT ON TABLE movimientos_contables.transaccion IS 'Tabla que almacena la informacion general de una transaccion';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.transaccion.id IS 'Identificador de la transaccion';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.transaccion.consecutivo_id IS 'Campo obligatorio que referencia al esquema consecutivos. Hace referencia al id  del consecutivo generado para la transaccion';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.transaccion.etiquetas IS 'Campo opcional de tipo json que almacena los diferentes campos que pueden estar asociados al encabezado de una transaccion ';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.transaccion.descripcion IS 'Campo opcional que indica la descripcion de la transaccion';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.transaccion.error_transaccion IS 'Campo opcional que indica un posible error que puede ocurrir durante la transaccion';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.transaccion.estado_id IS 'Campo obligatorio que referencia al esquema parametros a la tabla parametro. Hace referencia al id del estado asociado a la transaccion';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.transaccion.fecha_transaccion IS 'Fecha de realizacion de la transaccion';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.transaccion.activo IS 'Campo para indicar el estado del registro';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.transaccion.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.transaccion.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_transaccion ON movimientos_contables.transaccion  IS 'Llave primaria de la tabla transaccion';
-- ddl-end --






-- object: fk_transaccion_movimiento | type: CONSTRAINT --
-- ALTER TABLE movimientos_contables.movimiento DROP CONSTRAINT IF EXISTS fk_transaccion_movimiento CASCADE;
ALTER TABLE movimientos_contables.movimiento ADD CONSTRAINT fk_transaccion_movimiento FOREIGN KEY (transaccion_id)
REFERENCES movimientos_contables.transaccion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --


-- Permisos de usuario
GRANT USAGE ON SCHEMA movimientos_contables TO desarrollooas;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA movimientos_contables TO desarrollooas;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA movimientos_contables TO desarrollooas;

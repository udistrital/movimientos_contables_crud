CREATE TABLE movimientos_contables.saldo (
	id serial NOT NULL,
	cuenta_id text NOT NULL UNIQUE,
	debito numeric(20,7) NOT NULL,
	credito numeric(20,7) NOT NULL,
	saldo numeric(20,7) NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_saldo PRIMARY KEY (id)
);

-- ddl-end --
COMMENT ON COLUMN movimientos_contables.saldo.cuenta_id IS 'Identificador cuenta contable';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.saldo.debito IS 'Sumatoria de los valores debitos asociados a la cuenta contable';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.saldo.credito IS 'Sumatoria de los valores creditos asociados a la cuenta contable';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.saldo.fecha_creacion IS 'Fehca de creacion del registro';
-- ddl-end --
COMMENT ON COLUMN movimientos_contables.saldo.fecha_modificacion IS 'Fehca de modificacion del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_saldo ON movimientos_contables.saldo  IS 'Llave primaria de la tabla saldo';
-- ddl-end --

ALTER TABLE movimientos_contables.movimiento 
ADD COLUMN saldo_anterior numeric(20,7) NOT NULL DEFAULT 0;
-- ddl-end --

ALTER TABLE movimientos_contables.movimiento 
ADD COLUMN nuevo_saldo numeric(20,7) NOT NULL DEFAULT 0;
-- ddl-end --

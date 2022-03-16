DROP TABLE IF EXISTS movimientos_contables.saldo CASCADE;
-- ddl-end --

ALTER TABLE movimientos_contables.movimiento 
DROP COLUMN saldo_anterior;
-- ddl-end --

ALTER TABLE movimientos_contables.movimiento 
DROP COLUMN nuevo_saldo;
-- ddl-end --

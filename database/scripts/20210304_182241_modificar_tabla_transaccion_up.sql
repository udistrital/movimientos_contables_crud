-- se cambia tipo de dato de la columna error_transaccion
ALTER TABLE movimientos_contables.transaccion 
ALTER COLUMN error_transaccion type text;
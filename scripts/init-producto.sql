CREATE TABLE Producto (
    IDProducto SERIAL PRIMARY KEY,
    Nombre VARCHAR(255) NOT NULL,
    Descripsion TEXT,
    Imagen VARCHAR(255),
    Unidad BIGINT,
    Creado TIMESTAMP DEFAULT NOW(),
    Actualizado TIMESTAMP DEFAULT NOW()
);

CREATE OR REPLACE PROCEDURE RegistrarProducto(
    p_Nombre VARCHAR(255),
    p_Descripsion TEXT,
    p_Imagen VARCHAR(255),
    p_Unidad BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO Producto (Nombre, Descripsion, Imagen, Unidad, Creado, Actualizado)
    VALUES (p_Nombre, p_Descripsion, p_Imagen, p_Unidad, NOW(), NOW());
END;
$$;

CREATE OR REPLACE PROCEDURE ActualizarProducto(
    p_ID INT,
    p_Nombre VARCHAR(255),
    p_Descripsion TEXT,
    p_Imagen VARCHAR(255),
    p_Unidad BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE Producto
    SET Nombre = p_Nombre,
        Descripsion = p_Descripsion,
        Imagen = p_Imagen,
        Unidad = p_Unidad,
        Actualizado = NOW()
    WHERE IDProducto = p_ID;
END;
$$;


CREATE OR REPLACE FUNCTION EliminarProducto(
    p_ID INT
)
RETURNS INT
LANGUAGE plpgsql
AS $$
DECLARE
    filas_eliminadas INT;
BEGIN
    DELETE FROM Producto
    WHERE IDProducto = p_ID;
    GET DIAGNOSTICS filas_eliminadas = ROW_COUNT;
    RETURN filas_eliminadas;
END;
$$;

CREATE OR REPLACE FUNCTION ObtenerProducto(id_producto INTEGER)
RETURNS TABLE (
    IDProducto INTEGER,
    Nombre VARCHAR(255),
    Descripsion TEXT,
    Imagen VARCHAR(255),
    Unidad BIGINT,
    Creado TIMESTAMP,
    Actualizado TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY SELECT Producto.IDProducto, Producto.Nombre, Producto.Descripsion, Producto.Imagen, Producto.Unidad, Producto.Creado, Producto.Actualizado
                 FROM Producto
                 WHERE Producto.IDProducto = id_producto;
END;
$$ LANGUAGE plpgsql;
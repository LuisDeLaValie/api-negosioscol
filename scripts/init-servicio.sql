
CREATE TABLE Servisio (
    IDServicio SERIAL PRIMARY KEY,
    Nombre VARCHAR(255) NOT NULL,
    Descripcion TEXT,
    Imagen VARCHAR(255),
    Unidad BIGINT,
    Creado TIMESTAMP DEFAULT NOW(),
    Actualizado TIMESTAMP DEFAULT NOW()
);


CREATE OR REPLACE PROCEDURE RegistrarServisio(
    p_Nombre VARCHAR(255),
    p_Descripcion TEXT,
    p_Imagen VARCHAR(255),
    p_Unidad BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO Servisio (Nombre, Descripcion, Imagen, Unidad, Creado, Actualizado)
    VALUES (p_Nombre, p_Descripcion, p_Imagen, p_Unidad, NOW(), NOW());
END;
$$;


CREATE OR REPLACE PROCEDURE ActualizarServisio(
    p_ID INT,
    p_Nombre VARCHAR(255),
    p_Descripcion TEXT,
    p_Imagen VARCHAR(255),
    p_Unidad BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE Servisio
    SET Nombre = p_Nombre,
        Descripcion = p_Descripcion,
        Imagen = p_Imagen,
        Unidad = p_Unidad,
        Actualizado = NOW()
    WHERE IDServicio = p_ID;
END;
$$;

CREATE OR REPLACE FUNCTION EliminarServisio(
    p_ID INT
)
RETURNS INT
LANGUAGE plpgsql
AS $$
DECLARE
    filas_eliminadas INT;
BEGIN
    DELETE FROM Servisio
    WHERE IDServicio = p_ID;
    GET DIAGNOSTICS filas_eliminadas = ROW_COUNT;
    RETURN filas_eliminadas;
END;
$$;

CREATE OR REPLACE FUNCTION ObtenerServisio(id_servisio INTEGER)
RETURNS TABLE (
    IDServicio INTEGER,
    Nombre VARCHAR(255),
    Descripcion TEXT,
    Imagen VARCHAR(255),
    Unidad BIGINT,
    Creado TIMESTAMP,
    Actualizado TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY SELECT Servisio.IDServicio, Servisio.Nombre, Servisio.Descripcion, Servisio.Imagen, Servisio.Unidad, Servisio.Creado, Servisio.Actualizado
                 FROM Servisio
                 WHERE Servisio.IDServicio = id_servisio;
END;
$$ LANGUAGE plpgsql;
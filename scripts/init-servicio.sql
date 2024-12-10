
CREATE TABLE Servisio (
    IDServicio SERIAL PRIMARY KEY,
    Nombre VARCHAR(255) NOT NULL,
    Descripcion TEXT,
    Imagen VARCHAR(255),
    Unidad BIGINT,
    IDNegocio INTEGER NOT NULL,
    Precio INTEGER NOT NULL,
    Creado TIMESTAMP DEFAULT NOW(),
    Actualizado TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_negocio FOREIGN KEY (IDNegocio) REFERENCES Negocio(IDNegocio) ON DELETE CASCADE

);


CREATE OR REPLACE PROCEDURE RegistrarServisio(
    p_Nombre VARCHAR(255),
    p_Descripcion TEXT,
    p_Imagen VARCHAR(255),
    p_Unidad BIGINT,
    p_IDNegocio INT,
    p_Precio INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO Servisio (Nombre, Descripcion, Imagen, Unidad, IDNegocio, Precio, Creado, Actualizado)
    VALUES (p_Nombre, p_Descripcion, p_Imagen, p_Unidad, p_IDNegocio, p_Precio, NOW(), NOW());
END;
$$;


CREATE OR REPLACE PROCEDURE ActualizarServisio(
    p_ID INT,
    p_Nombre VARCHAR(255),
    p_Descripcion TEXT,
    p_Imagen VARCHAR(255),
    p_Unidad BIGINT,
    p_Precio INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE Servisio
    SET Nombre = p_Nombre,
        Descripcion = p_Descripcion,
        Unidad = p_Unidad,
        Precio= p_Precio,
        Actualizado = NOW()
    WHERE IDServicio = p_ID;

     -- Actualizar Imagen solo si se proporciona
    IF p_Imagen IS NOT NULL THEN
        UPDATE Servisio
        SET Imagen = p_Imagen
        WHERE IDNegocio = p_ID;
    END IF;


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
    idnegocio INTEGER,
    Precio INTEGER,
    Creado TIMESTAMP,
    Actualizado TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY SELECT Servisio.IDServicio, Servisio.Nombre, Servisio.Descripcion, Servisio.Imagen, Servisio.Unidad,Servisio.idnegocio, Servisio.Precio, Servisio.Creado, Servisio.Actualizado
                 FROM Servisio
                 WHERE Servisio.IDServicio = id_servisio;
END;
$$ LANGUAGE plpgsql;
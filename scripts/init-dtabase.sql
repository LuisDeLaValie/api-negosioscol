
/*
 * CREAR BASE DE DATOS
 */

CREATE DATABASE NegociosCol;

\c NegociosCol

CREATE TABLE Usuario (
    ID SERIAL PRIMARY KEY,
    Nombre VARCHAR(255),
    Apellidos VARCHAR(255),
    Creado TIMESTAMP,
    Actualizado TIMESTAMP,
    Cumpleanos TIMESTAMP,
    Imagen VARCHAR(255)
);


CREATE TABLE Servisio (
    IDProducto SERIAL PRIMARY KEY,
    Nombre VARCHAR(255) NOT NULL,
    Descripcion TEXT,
    Unidad BIGINT,
    Creado TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Actualizado TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



/*
 * Procedure
 */

CREATE OR REPLACE PROCEDURE RegistrarUsuario(
    p_Nombre VARCHAR(255),
    p_Apellidos VARCHAR(255),
    p_Cumpleanos TIMESTAMP,
    p_Imagen VARCHAR(255)
)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO Usuario (Nombre, Apellidos, Creado, Actualizado, Cumpleanos, Imagen)
    VALUES (p_Nombre, p_Apellidos, NOW(), NOW(), p_Cumpleanos, p_Imagen);
END;
$$;


CREATE OR REPLACE PROCEDURE InsertarServisio(    
    IN p_Nombre VARCHAR(255),
    IN p_Descripcion TEXT,
    IN p_Unidad BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO Servisio ( Nombre, Descripcion, Unidad)
    VALUES ( p_Nombre, p_Descripcion, p_Unidad);
END;
$$;


CREATE OR REPLACE PROCEDURE ActualizarUsuario(
    p_ID INT,
    p_Nombre VARCHAR(255),
    p_Apellidos VARCHAR(255),
    p_Cumpleanos TIMESTAMP,
    p_Imagen VARCHAR(255)
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE Usuario
    SET Nombre = p_Nombre,
        Apellidos = p_Apellidos,
        Actualizado = NOW(),
        Cumpleanos = p_Cumpleanos,
        Imagen = p_Imagen
    WHERE ID = p_ID;
END;
$$;

CREATE OR REPLACE PROCEDURE ActualizarServisio(
    IN p_IDProducto int,
    IN p_Nombre VARCHAR(255),
    IN p_Descripcion TEXT,
    IN p_Unidad BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE Servisio
    SET Nombre = p_Nombre,
        Descripcion = p_Descripcion,
        Unidad = p_Unidad,
        Actualizado = NOW()
    WHERE IDProducto = p_IDProducto;
END;
$$;



/**
 * Funciones
 */

CREATE OR REPLACE FUNCTION EliminarUsuario(
    p_ID INT
)
RETURNS INT
LANGUAGE plpgsql
AS $$
DECLARE
    filas_eliminadas INT;
BEGIN
    DELETE FROM Usuario
    WHERE ID = p_ID;
    GET DIAGNOSTICS filas_eliminadas = ROW_COUNT;
   	return filas_eliminadas;
END;
$$;


CREATE OR REPLACE FUNCTION EliminarServisio(p_IDProducto int)
RETURNS INT
LANGUAGE plpgsql
AS $$
DECLARE
    filas_eliminadas INT;
BEGIN
    DELETE FROM Servisio 
    WHERE IDProducto = p_IDProducto;
    GET DIAGNOSTICS filas_eliminadas = ROW_COUNT;
   	return filas_eliminadas;
END;
$$ ;



CREATE OR replace FUNCTION ObtenerUsuario(id_usuario INTEGER)
RETURNS TABLE (
    ID INTEGER,
    Nombre VARCHAR(255),
    Apellidos VARCHAR(255),
    Creado TIMESTAMP,
    Actualizado TIMESTAMP,
    Cumpleanos TIMESTAMP,
    Imagen VARCHAR(255)
) AS $$
BEGIN
    RETURN QUERY SELECT Usuario.ID, Usuario.Nombre, Usuario.Apellidos, Usuario.Creado, Usuario.Actualizado, Usuario.Cumpleanos, Usuario.Imagen
                 FROM Usuario
                 WHERE Usuario.ID = id_usuario;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION ObtenerServisio(p_IDProducto INTEGER)
RETURNS TABLE (
    IDProducto INTEGER,
    Nombre VARCHAR,
    Descripcion TEXT,
    Unidad BIGINT,
    Creado TIMESTAMP,
    Actualizado TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY SELECT Servisio.IDProducto, Servisio.Nombre, Servisio.Descripcion, Servisio.Unidad, Servisio.Creado, Servisio.Actualizado
     FROM Servisio WHERE Servisio.IDProducto = p_IDProducto;
END;
$$ LANGUAGE plpgsql;



CREATE OR REPLACE FUNCTION ListarServisio()
RETURNS TABLE (
    IDProducto INTEGER,
    Nombre VARCHAR,
    Descripcion TEXT,
    Unidad BIGINT,
    Creado TIMESTAMP,
    Actualizado TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY SELECT * FROM Servisio;
END;
$$ LANGUAGE plpgsql;





/**
 * PUREBAS
 */

CALL RegistrarUsuario('Juan', 'Pérez', '1990-01-01', 'https://bestprofilepictures.com/wp-content/uploads/2021/08/Amazing-Profile-Picture.jpg');

CALL ActualizarUsuario(1, 'NuevoNombre', 'NuevosApellidos', '1990-01-01', 'https://bestprofilepictures.com/wp-content/uploads/2021/08/Amazing-Profile-Picture.jpg');

SELECT EliminarUsuario(1);

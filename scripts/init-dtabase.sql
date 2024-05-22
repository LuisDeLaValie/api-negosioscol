
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



/**
 * PUREBAS
 */

CALL RegistrarUsuario('Juan', 'Pérez', '1990-01-01', 'https://bestprofilepictures.com/wp-content/uploads/2021/08/Amazing-Profile-Picture.jpg');

CALL ActualizarUsuario(1, 'NuevoNombre', 'NuevosApellidos', '1990-01-01', 'https://bestprofilepictures.com/wp-content/uploads/2021/08/Amazing-Profile-Picture.jpg');

SELECT EliminarUsuario(1);

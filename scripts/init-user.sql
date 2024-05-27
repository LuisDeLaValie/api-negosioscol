
CREATE TABLE Usuario (
    ID SERIAL PRIMARY KEY,
    Nombre VARCHAR(255) NOT NULL,
    Apellidos VARCHAR(255) NOT NULL,
    Correo VARCHAR(255) NOT NULL UNIQUE,
    Password VARCHAR(255) NOT NULL,
    Creado TIMESTAMP DEFAULT NOW(),
    Actualizado TIMESTAMP DEFAULT NOW(),
    Cumpleanos TIMESTAMP,
    Imagen VARCHAR(255)
);

CREATE OR REPLACE PROCEDURE RegistrarUsuario(
    p_Nombre VARCHAR(255),
    p_Apellidos VARCHAR(255),
    p_Correo VARCHAR(255),
    p_Password VARCHAR(255),
    p_Cumpleanos TIMESTAMP,
    p_Imagen VARCHAR(255)
)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO Usuario (Nombre, Apellidos, Correo, Password, Creado, Actualizado, Cumpleanos, Imagen)
    VALUES (p_Nombre, p_Apellidos, p_Correo, p_Password, NOW(), NOW(), p_Cumpleanos, p_Imagen);
END;
$$;


CREATE OR REPLACE PROCEDURE ActualizarUsuario(
    p_ID INT,
    p_Nombre VARCHAR(255),
    p_Apellidos VARCHAR(255),
    p_Correo VARCHAR(255),
    p_Password VARCHAR(255),
    p_Cumpleanos TIMESTAMP,
    p_Imagen VARCHAR(255)
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE Usuario
    SET Nombre = p_Nombre,
        Apellidos = p_Apellidos,
        Correo = p_Correo,
        Password = p_Password,
        Actualizado = NOW(),
        Cumpleanos = p_Cumpleanos,
        Imagen = p_Imagen
    WHERE ID = p_ID;
END;
$$;

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
    RETURN filas_eliminadas;
END;
$$;


CREATE OR REPLACE FUNCTION ObtenerUsuario(id_usuario INTEGER)
RETURNS TABLE (
    ID INTEGER,
    Nombre VARCHAR(255),
    Apellidos VARCHAR(255),
    Correo VARCHAR(255),
    Password VARCHAR(255),
    Creado TIMESTAMP,
    Actualizado TIMESTAMP,
    Cumpleanos TIMESTAMP,
    Imagen VARCHAR(255)
) AS $$
BEGIN
    RETURN QUERY SELECT Usuario.ID, Usuario.Nombre, Usuario.Apellidos, Usuario.Correo, Usuario.Password, Usuario.Creado, Usuario.Actualizado, Usuario.Cumpleanos, Usuario.Imagen
                 FROM Usuario
                 WHERE Usuario.ID = id_usuario;
END;
$$ LANGUAGE plpgsql;
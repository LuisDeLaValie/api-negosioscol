
CREATE TABLE Negocio (
    IDNegocio SERIAL PRIMARY KEY,
    Nombre VARCHAR(255) NOT NULL,
    Descripsion TEXT,
    Direccion VARCHAR(255),
    Telefono VARCHAR(50),
    Correo VARCHAR(255),
    Imagen VARCHAR(255),
    Latitude FLOAT,
    Longitude FLOAT,
    Facebook VARCHAR(255),
    Twitter VARCHAR(255),
    Instagram VARCHAR(255),
    Website VARCHAR(255),
    Creado TIMESTAMP DEFAULT NOW(),
    Actualizado TIMESTAMP DEFAULT NOW()
);


CREATE OR REPLACE PROCEDURE RegistrarNegocio(
    p_Nombre VARCHAR(255),
    p_Descripsion TEXT,
    p_Direccion VARCHAR(255),
    p_Telefono VARCHAR(50),
    p_Correo VARCHAR(255),
    p_Imagen VARCHAR(255),
    p_Latitude FLOAT,
    p_Longitude FLOAT,
    p_Facebook VARCHAR(255),
    p_Twitter VARCHAR(255),
    p_Instagram VARCHAR(255),
    p_Website VARCHAR(255)
)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO Negocio (Nombre, Descripsion, Direccion, Telefono, Correo, Imagen, Latitude, Longitude, Facebook, Twitter, Instagram, Website, Creado, Actualizado)
    VALUES (p_Nombre, p_Descripsion, p_Direccion, p_Telefono, p_Correo, p_Imagen, p_Latitude, p_Longitude, p_Facebook, p_Twitter, p_Instagram, p_Website, NOW(), NOW());
END;
$$;

CREATE OR REPLACE PROCEDURE ActualizarNegocio(
    p_ID INT,
    p_Nombre VARCHAR(255),
    p_Descripsion TEXT,
    p_Direccion VARCHAR(255),
    p_Telefono VARCHAR(50),
    p_Correo VARCHAR(255),
    p_Imagen VARCHAR(255),
    p_Latitude FLOAT,
    p_Longitude FLOAT,
    p_Facebook VARCHAR(255),
    p_Twitter VARCHAR(255),
    p_Instagram VARCHAR(255),
    p_Website VARCHAR(255)
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE Negocio
    SET Nombre = p_Nombre,
        Descripsion = p_Descripsion,
        Direccion = p_Direccion,
        Telefono = p_Telefono,
        Correo = p_Correo,
        Imagen = p_Imagen,
        Latitude = p_Latitude,
        Longitude = p_Longitude,
        Facebook = p_Facebook,
        Twitter = p_Twitter,
        Instagram = p_Instagram,
        Website = p_Website,
        Actualizado = NOW()
    WHERE ID = p_ID;
END;
$$;


CREATE OR REPLACE FUNCTION EliminarNegocio(
    p_ID INT
)
RETURNS INT
LANGUAGE plpgsql
AS $$
DECLARE
    filas_eliminadas INT;
BEGIN
    DELETE FROM Negocio
    WHERE ID = p_ID;
    GET DIAGNOSTICS filas_eliminadas = ROW_COUNT;
    RETURN filas_eliminadas;
END;
$$;

CREATE OR REPLACE FUNCTION ObtenerNegocio(id_negocio INTEGER)
RETURNS TABLE (
    ID INTEGER,
    Nombre VARCHAR(255),
    Descripsion TEXT,
    Direccion VARCHAR(255),
    Telefono VARCHAR(50),
    Correo VARCHAR(255),
    Imagen VARCHAR(255),
    Latitude FLOAT,
    Longitude FLOAT,
    Facebook VARCHAR(255),
    Twitter VARCHAR(255),
    Instagram VARCHAR(255),
    Website VARCHAR(255),
    Creado TIMESTAMP,
    Actualizado TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY SELECT Negocio.ID, Negocio.Nombre, Negocio.Descripsion, Negocio.Direccion, Negocio.Telefono, Negocio.Correo, Negocio.Imagen, Negocio.Latitude, Negocio.Longitude, Negocio.Facebook, Negocio.Twitter, Negocio.Instagram, Negocio.Website, Negocio.Creado, Negocio.Actualizado
                 FROM Negocio
                 WHERE Negocio.ID = id_negocio;
END;
$$ LANGUAGE plpgsql;

CALL public.registrarnegocio('carpinteria', 'tienda de carpinteria', ':p_direccion', ':p_telefono', ':p_correo', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQYkdI_Bbm2wGjEXStSOCSw0-zPlLVzU4O9W8RGewADkg&s', 2.2, 365.2, ':p_facebook', ':p_twitter', ':p_instagram', ':p_website');
CALL public.registrarnegocio('ferreteria', 'tienda de fereteria', ':p_direccion', ':p_telefono', ':p_correo', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQYkdI_Bbm2wGjEXStSOCSw0-zPlLVzU4O9W8RGewADkg&s', 2.2, 365.2, ':p_facebook', ':p_twitter', ':p_instagram', ':p_website');
CALL public.registrarnegocio('piezeria', 'tienda de pizza', ':p_direccion', ':p_telefono', ':p_correo', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQYkdI_Bbm2wGjEXStSOCSw0-zPlLVzU4O9W8RGewADkg&s', 2.2, 365.2, ':p_facebook', ':p_twitter', ':p_instagram', ':p_website');

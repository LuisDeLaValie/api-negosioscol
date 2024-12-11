
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



-- Insertando servicios para los negocios

INSERT INTO Servisio (Nombre, Descripcion, Imagen, Unidad, IDNegocio, Precio, Creado, Actualizado) 
VALUES 
-- La Casa de la Pizza
('Delivery a Domicilio', 'Servicio de entrega de pizza a domicilio en un radio de 10 km.', 'delivery.jpg', 1, 1, 50, '2024-07-01', '2024-12-31'),
('Servicio para Eventos', 'Catering para eventos, con pizzas personalizadas y bebidas.', 'evento.jpg', 1, 1, 1500, '2024-07-01', '2024-12-31'),
-- Restaurante El Sol
('Comida a Domicilio', 'Envío de tus platillos favoritos hasta tu hogar.', 'comida_domicilio.jpg', 1, 2, 60, '2024-07-01', '2024-12-31'),
('Banquetes para Bodas', 'Servicio de banquetes con opciones gourmet para tu boda.', 'banquetes_boda.jpg', 1, 2, 5000, '2024-07-01', '2024-12-31'),
-- Tecnología y Más
('Reparación de Equipos', 'Servicio técnico para reparación de teléfonos, computadoras y más.', 'reparacion_equipos.jpg', 1, 3, 700, '2024-07-01', '2024-12-31'),
('Asesoría Técnica', 'Asesoría en la compra de equipos tecnológicos y configuración.', 'asesoria_tecnica.jpg', 1, 3, 350, '2024-07-01', '2024-12-31'),
-- Librería El Saber
('Venta Online de Libros', 'Compra libros desde la comodidad de tu hogar.', 'venta_online.jpg', 1, 4, 0, '2024-07-01', '2024-12-31'), 
('Encuadernación Personalizada', 'Encuadernación a medida de tus libros y documentos.', 'encuadernacion.jpg', 1, 4, 200, '2024-07-01', '2024-12-31'),
-- Supermercado La Esperanza
('Servicio de Entrega a Domicilio', 'Servicio de entrega de productos a tu domicilio en menos de 24 horas.', 'entrega_domicilio.jpg', 1, 5, 100, '2024-07-01', '2024-12-31'),
('Recogida en Tienda', 'Recoge tus compras online en la tienda sin costo adicional.', 'recogida_tienda.jpg', 1, 5, 0, '2024-07-07', '2024-12-31'),
-- Centro Estético Bella
('Tratamientos Faciales', 'Tratamientos para rejuvenecimiento y cuidado de la piel.', 'tratamiento_facial.jpg', 1, 6, 800, '2024-08-01', '2024-12-31'),
('Masajes Relajantes', 'Masajes terapéuticos para liberar estrés y tensiones.', 'masaje_relajante.jpg', 1, 6, 500, '2024-08-01', '2024-12-31'),
-- Cafetería Dulce Aroma
('Café para Empresas', 'Servicio de café y pasteles para reuniones de empresa.', 'cafe_empresas.jpg', 1, 7, 500, '2024-08-01', '2024-12-31'),
('Repostería Personalizada', 'Repostería hecha a medida para celebraciones especiales.', 'reposteria_personalizada.jpg', 1, 7, 800, '2024-08-01', '2024-12-31'),
-- Gimnasio PowerFit
('Clases de Yoga', 'Clases de yoga para mejorar flexibilidad y relajación.', 'clases_yoga.jpg', 1, 8, 250, '2024-09-01', '2024-12-31'),
('Clases de Zumba', 'Clases divertidas de zumba para quemar calorías y tonificar.', 'clases_zumba.jpg', 1, 8, 150, '2024-09-01', '2024-12-31');

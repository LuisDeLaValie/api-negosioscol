CREATE TABLE Producto (
    IDProducto SERIAL PRIMARY KEY,
    Nombre VARCHAR(255) NOT NULL,
    Descripsion TEXT,
    Imagen VARCHAR(255),
    Unidad BIGINT,
    IDNegocio INTEGER NOT NULL,
    Precio INTEGER NOT NULL,
    Creado TIMESTAMP DEFAULT NOW(),
    Actualizado TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_negocio FOREIGN KEY (IDNegocio) REFERENCES Negocio(IDNegocio) ON DELETE CASCADE

);

CREATE OR REPLACE PROCEDURE RegistrarProducto(
    p_Nombre VARCHAR(255),
    p_Descripsion TEXT,
    p_Imagen VARCHAR(255),
    p_Unidad BIGINT,
    p_IDNegocio INT,
    p_Precio INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO Producto (Nombre, Descripsion, Imagen, Unidad, IDNegocio,Precio , Creado, Actualizado)
    VALUES (p_Nombre, p_Descripsion, p_Imagen, p_Unidad, p_IDNegocio,p_Precio, NOW(), NOW());
END;
$$;

CREATE OR REPLACE PROCEDURE ActualizarProducto(
    p_ID INT,
    p_Nombre VARCHAR(255),
    p_Descripsion TEXT,
    p_Imagen VARCHAR(255),
    p_Unidad BIGINT,
    p_Precio INTEGER

)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE Producto
    SET Nombre = p_Nombre,
        Descripsion = p_Descripsion,
        Unidad = p_Unidad,
        Precio= p_Precio,
        Actualizado = NOW()
    WHERE IDProducto = p_ID;

    -- Actualizar Imagen solo si se proporciona
    IF p_Imagen IS NOT NULL THEN
        UPDATE Producto
        SET Imagen = p_Imagen
        WHERE IDNegocio = p_ID;
    END IF;
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
    idnegocio INTEGER,
    Precio INTEGER,
    Creado TIMESTAMP,
    Actualizado TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY SELECT Producto.IDProducto, Producto.Nombre, Producto.Descripsion, Producto.Imagen, Producto.Unidad, Producto.idnegocio, Producto.Precio, Producto.Creado, Producto.Actualizado
                 FROM Producto
                 WHERE Producto.IDProducto = id_producto;
END;
$$ LANGUAGE plpgsql;



-- Insertando productos para los negocios

-- Insertando productos en la tabla Producto con fechas
INSERT INTO Producto (Nombre, Descripsion, Imagen, Unidad, IDNegocio, Precio, Creado, Actualizado) 
VALUES 
-- La Casa de la Pizza
('Pizza Margarita', 'Pizza con salsa de tomate, queso mozzarella y albahaca.', 'pizza_margarita.jpg', 1, 1, 150, '2024-08-01 10:00:00', '2024-08-01 10:00:00'),
('Pizza Pepperoni', 'Pizza con salsa de tomate, queso mozzarella y pepperoni.', 'pizza_pepperoni.jpg', 1, 1, 180, '2024-08-02 11:00:00', '2024-08-02 11:00:00'),
('Pizza Hawaiana', 'Pizza con salsa de tomate, queso mozzarella, jamón y piña.', 'pizza_hawaiana.jpg', 1, 1, 190, '2024-08-03 12:00:00', '2024-08-03 12:00:00'),
('Pizza 4 Quesos', 'Pizza con salsa de tomate, queso mozzarella, cheddar, gorgonzola y parmesano.', 'pizza_4quesos.jpg', 1, 1, 220, '2024-08-04 13:00:00', '2024-08-04 13:00:00'),
('Alitas Picantes', 'Alitas de pollo cubiertas con salsa picante.', 'alitas_picantes.jpg', 1, 1, 100, '2024-08-05 14:00:00', '2024-08-05 14:00:00'),
('Pan de Ajo', 'Pan recién horneado con mantequilla y ajo.', 'pan_de_ajo.jpg', 1, 1, 50, '2024-08-06 15:00:00', '2024-08-06 15:00:00'),
-- Restaurante El Sol
('Tacos al Pastor', 'Tacos con carne de cerdo marinada, piña y salsa al pastor.', 'tacos_al_pastor.jpg', 1, 2, 60, '2024-08-07 16:00:00', '2024-08-07 16:00:00'),
('Burritos de Carne', 'Burritos rellenos de carne asada, frijoles y arroz.', 'burritos_carne.jpg', 1, 2, 80, '2024-08-08 17:00:00', '2024-08-08 17:00:00'),
('Pozole', 'Sopa de maíz con carne de cerdo y condimentos.', 'pozole.jpg', 1, 2, 120, '2024-08-09 18:00:00', '2024-08-09 18:00:00'),
-- Tecnología y Más
('Smartphone Modelo X', 'Teléfono inteligente con pantalla 6.5", 128GB de almacenamiento.', 'smartphone_modelo_x.jpg', 1, 3, 5000, '2024-08-10 19:00:00', '2024-08-10 19:00:00'),
('Laptop Ultra Rápida', 'Laptop con procesador i7, 16GB RAM, 512GB SSD.', 'laptop_rapida.jpg', 1, 3, 15000, '2024-09-11 20:00:00', '2024-09-11 20:00:00'),
('Smartwatch', 'Reloj inteligente con monitoreo de actividad y notificaciones.', 'smartwatch.jpg', 1, 3, 2000, '2024-09-12 21:00:00', '2024-09-12 21:00:00'),
-- Librería El Saber
('Diccionario Inglés-Español', 'Diccionario de inglés a español para estudiantes y profesionales.', 'diccionario_ingles.jpg', 1, 4, 350, '2024-09-13 22:00:00', '2024-09-13 22:00:00'),
('Libro de Cocina Mexicana', 'Recetas tradicionales de la cocina mexicana.', 'cocina_mexicana.jpg', 1, 4, 500, '2024-09-14 23:00:00', '2024-09-14 23:00:00'),
('Novela "1984"', 'Novela distópica de George Orwell.', '1984.jpg', 1, 4, 300, '2024-09-15 00:00:00', '2024-09-15 00:00:00'),
-- Supermercado La Esperanza
('Leche Entera', 'Leche fresca entera, ideal para desayunos.', 'leche_entera.jpg', 1, 5, 25, '2024-09-16 01:00:00', '2024-09-16 01:00:00'),
('Pan Integral', 'Pan de trigo integral, rico en fibra.', 'pan_integral.jpg', 1, 5, 30, '2024-09-17 02:00:00', '2024-09-17 02:00:00'),
('Aceite de Oliva', 'Aceite extra virgen, ideal para ensaladas y cocina.', 'aceite_oliva.jpg', 1, 5, 100, '2024-09-18 03:00:00', '2024-09-18 03:00:00'),
-- Centro Estético Bella
('Crema Facial Antiarrugas', 'Crema hidratante que combate los signos de envejecimiento.', 'crema_antiarrugas.jpg', 1, 6, 450, '2024-07-19 04:00:00', '2024-07-19 04:00:00'),
('Mascarilla Facial', 'Mascarilla que limpia y rejuvenece la piel.', 'mascarilla_facial.jpg', 1, 6, 250, '2024-07-20 05:00:00', '2024-07-20 05:00:00'),
('Perfume Floral', 'Perfume con fragancia floral de alta calidad.', 'perfume_floral.jpg', 1, 6, 300, '2024-07-21 06:00:00', '2024-07-21 06:00:00'),
-- Cafetería Dulce Aroma
('Café Americano', 'Café preparado con granos de alta calidad.', 'cafe_americano.jpg', 1, 7, 40, '2024-07-22 07:00:00', '2024-07-22 07:00:00'),
('Croissant de Mantequilla', 'Croissant recién horneado, suave y mantecoso.', 'croissant_mantequilla.jpg', 1, 7, 35, '2024-07-23 08:00:00', '2024-07-23 08:00:00'),
('Donut Glaseado', 'Dulce donut cubierto con glaseado de vainilla.', 'donut_glaseado.jpg', 1, 7, 25, '2024-07-24 09:00:00', '2024-07-24 09:00:00'),
-- Gimnasio PowerFit
('Proteína en Polvo', 'Proteína para recuperación muscular, sabor chocolate.', 'proteina_chocolate.jpg', 1, 8, 600, '2024-07-25 10:00:00', '2024-07-25 10:00:00'),
('BCAA', 'Aminoácidos de cadena ramificada para mejorar el rendimiento.', 'bcaa.jpg', 1, 8, 350, '2024-07-26 11:00:00', '2024-07-26 11:00:00'),
('Mancuernas', 'Mancuernas de 5kg para ejercicios en casa.', 'mancuernas.jpg', 1, 8, 500, '2024-07-27 12:00:00', '2024-07-27 12:00:00'),
-- Bar El Refugio
('Cerveza Artesanal', 'Cerveza artesanal de sabor suave.', 'cerveza_artesanal.jpg', 1, 9, 80, '2024-07-28 13:00:00', '2024-07-28 13:00:00'),
('Tequila Reposado', 'Tequila reposado premium.', 'tequila_reposado.jpg', 1, 9, 350, '2024-07-29 14:00:00', '2024-07-29 14:00:00'),
('Cócteles Especiales', 'Cócteles mixtos con frutas tropicales.', 'cocteles.jpg', 1, 9, 150, '2024-07-30 15:00:00', '2024-07-30 15:00:00');



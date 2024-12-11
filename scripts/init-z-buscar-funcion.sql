CREATE TABLE Busquedas(
    Consulta VARCHAR(100),
    IDNegocio INTEGER,
    IDProducto INTEGER,
    IDServicio INTEGER,
    Creado TIMESTAMP DEFAULT NOW()

);

CREATE OR REPLACE FUNCTION BuscarElementos(p_termino VARCHAR)
RETURNS TABLE (
    IDNegocio INTEGER,
    IDProducto INTEGER,
    IDServicio INTEGER,
    Negocio VARCHAR,
    Nombre VARCHAR,
    Descripcion TEXT,
    Imagen VARCHAR
) AS $$
BEGIN
    RETURN QUERY
    -- Buscar negocios
    SELECT 
        n.idnegocio,
        NULL::integer AS idproducto,
        NULL::integer AS idservicio,
        n.nombre AS negocio,
        n.nombre,
        n.descripsion AS descripcion,
        n.imagen
    FROM 
        Negocio n
    WHERE
        n.nombre ILIKE '%' || p_termino || '%'
        OR n.descripsion ILIKE '%' || p_termino || '%'
    
    UNION ALL
    
    -- Buscar productos
    SELECT 
        p.idnegocio  AS idnegocio,
        p.idproducto,
        NULL::integer AS idservicio,
        n.nombre  AS negocio,
        p.nombre,
        p.descripsion AS descripcion,
        p.imagen
    FROM 
        Producto p
	JOIN 
        negocio n ON p.idnegocio  = n.idnegocio 
    WHERE
        p.nombre ILIKE '%' || p_termino || '%'
        OR p.descripsion ILIKE '%' || p_termino || '%'
    
    UNION ALL
    
    -- Buscar servicios
    SELECT 
        s.idnegocio  AS idnegocio,
        NULL::integer AS idproducto,
        s.idservicio,
        n.nombre AS negocio,
        s.nombre,
        s.descripcion AS descripcion,
        s.imagen
    FROM 
        Servisio s
	JOIN 
        negocio n ON s.idnegocio  = n.idnegocio 
    WHERE
        s.nombre ILIKE '%' || p_termino || '%'
        OR s.descripcion ILIKE '%' || p_termino || '%';
END;
$$ LANGUAGE plpgsql;


-- Insertando registros en la tabla Busquedas
INSERT INTO Busquedas (Consulta, IDNegocio, IDProducto, IDServicio, Creado) 
VALUES
('Pizza cerca de mí', 1, 1, NULL, '2024-11-10 14:30:00'),  -- La Casa de la Pizza, Producto
('Comida mexicana', 2, 3, NULL, '2024-11-11 09:00:00'),  -- Restaurante El Sol, Producto
('Equipos de tecnología', 3, 5, NULL, '2024-11-12 11:15:00'),  -- Tecnología y Más, Producto
('Libros de historia', 4, 7, NULL, '2024-11-13 16:20:00'),  -- Librería El Saber, Producto
('Supermercado con envío a domicilio', 5, 9, NULL, '2024-11-14 10:30:00'),  -- Supermercado La Esperanza, Producto
('Corte de cabello', 6, NULL, 1, '2024-11-15 14:00:00'),  -- Centro Estético Bella, Servicio
('Café y pasteles', 7, 10, NULL, '2024-11-16 13:45:00'),  -- Cafetería Dulce Aroma, Producto
('Gimnasio para mujeres', 8, NULL, 2, '2024-11-17 09:30:00'),  -- Gimnasio PowerFit, Servicio
('Fiesta de cumpleaños', 9, NULL, 3, '2024-11-18 12:00:00'),  -- Bar El Refugio, Servicio
('Herramientas para construcción', 10, 11, NULL, '2024-11-19 15:10:00'),  -- Ferretería El Martillo, Producto
('Zapatos deportivos', 11, 13, NULL, '2024-11-20 08:30:00'),  -- Zapatería Rápido y Cómodo, Producto
('Comida en el tianguis', 12, 15, NULL, '2024-11-21 17:30:00'),  -- Tianguis El Mercado, Producto
('Guardería para niños', 13, NULL, 4, '2024-11-22 10:00:00'),  -- Escuela Infantil Pequeños Genios, Servicio
('Mariscos frescos', 14, 17, NULL, '2024-11-23 14:45:00'),  -- Restaurante Mariscos El Goloso, Producto
('Panadería artesanal', 15, 19, NULL, '2024-11-24 09:30:00'),  -- Panadería La Esperanza, Producto
('Medicamentos genéricos', 16, 21, NULL, '2024-11-25 12:00:00'),  -- Farmacia Salud y Bienestar, Producto
('Examen de vista', 17, NULL, 5, '2024-11-26 10:15:00'),  -- Óptica Visión Clara, Servicio
('Ropa de temporada', 18, 23, NULL, '2024-11-27 13:00:00'),  -- Bazar El Tesoro, Producto
('Clases de inglés', 19, NULL, 6, '2024-11-28 15:45:00'),  -- Centro de Idiomas Global, Servicio
('Moda femenina', 20, 25, NULL, '2024-11-29 08:20:00'),  -- Tienda de Ropa Chic, Producto
('Lentes para prescripción', 21, 27, NULL, '2024-11-30 10:00:00'),  -- Óptica Salud Visual, Producto
('Juguetes educativos', 22, 29, NULL, '2024-12-01 14:30:00'),  -- Juguetería Mundo Infantil, Producto
('Cuidado para mascotas', 23, NULL, 7, '2024-12-02 12:15:00');  -- Veterinaria La Mascota Feliz, Servicio



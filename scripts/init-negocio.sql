
CREATE TABLE Negocio (
    IDNegocio SERIAL PRIMARY KEY,
    Nombre VARCHAR(255) NOT NULL,
    Password VARCHAR(255) NOT NULL,
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
    p_Password VARCHAR(255),
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
    INSERT INTO Negocio (Nombre, Password, Descripsion, Direccion, Telefono, Correo, Imagen, Latitude, Longitude, Facebook, Twitter, Instagram, Website, Creado, Actualizado)
    VALUES (p_Nombre, p_Password, p_Descripsion, p_Direccion, p_Telefono, p_Correo, p_Imagen, p_Latitude, p_Longitude, p_Facebook, p_Twitter, p_Instagram, p_Website, NOW(), NOW());
END;
$$;

CREATE OR REPLACE PROCEDURE ActualizarNegocio(
    p_ID INT,
    p_Nombre VARCHAR(255),
    p_Password VARCHAR(255),
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
    SET 
        Nombre = p_Nombre,
        Descripsion = p_Descripsion,
        Password = p_Password,
        Direccion = p_Direccion,
        Telefono = p_Telefono,
        Correo = p_Correo,
        Latitude = p_Latitude,
        Longitude = p_Longitude,
        Website = p_Website,
        Actualizado = NOW()
    WHERE IDNegocio = p_ID;

    -- Actualizar Imagen solo si se proporciona
    IF p_Imagen IS NOT NULL THEN
        UPDATE Negocio
        SET Imagen = p_Imagen
        WHERE IDNegocio = p_ID;
    END IF;

    -- Actualizar redes sociales solo si se proporcionan
    IF p_Facebook IS NOT NULL THEN
        UPDATE Negocio
        SET Facebook = p_Facebook
        WHERE IDNegocio = p_ID;
    END IF;

    IF p_Twitter IS NOT NULL THEN
        UPDATE Negocio
        SET Twitter = p_Twitter
        WHERE IDNegocio = p_ID;
    END IF;

    IF p_Instagram IS NOT NULL THEN
        UPDATE Negocio
        SET Instagram = p_Instagram
        WHERE IDNegocio = p_ID;
    END IF;
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
    WHERE IDNegocio = p_ID;
    GET DIAGNOSTICS filas_eliminadas = ROW_COUNT;
    RETURN filas_eliminadas;
END;
$$;

CREATE OR REPLACE FUNCTION ObtenerNegocio(id_negocio INTEGER)
RETURNS TABLE (
    IDNegocio INTEGER,
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
    RETURN QUERY SELECT Negocio.IDNegocio, Negocio.Nombre, Negocio.Descripsion, Negocio.Direccion, Negocio.Telefono, Negocio.Correo, Negocio.Imagen, Negocio.Latitude, Negocio.Longitude, Negocio.Facebook, Negocio.Twitter, Negocio.Instagram, Negocio.Website, Negocio.Creado, Negocio.Actualizado
                 FROM Negocio
                 WHERE Negocio.IDNegocio = id_negocio;
END;
$$ LANGUAGE plpgsql;




INSERT INTO Negocio 
(Nombre, Password, Descripsion, Direccion, Telefono, Correo, Imagen, Latitude, Longitude, Facebook, Twitter, Instagram, Website) 
VALUES
('La Casa de la Pizza', 'contraseña123', 'Pizzería tradicional con una variedad de pizzas artesanales y pastas.', 'Av. Juárez 123, Colima', '312-123-4567', 'lacasa@pizza.com', 'pizzeria.jpg', 19.2435, -103.7222, 'https://facebook.com/lacasadelapizza', 'https://twitter.com/lacasapizza', 'https://instagram.com/lacasapizza', 'https://lacasa-pizza.com'),
('Restaurante El Sol', 'claveSecreta456', 'Comida mexicana típica, especialidad en tacos al pastor y pozole.', 'Calle Madero 45, Colima', '312-987-6543', 'elsol@restaurant.com', 'restaurante_sol.jpg', 19.2310, -103.7185, 'https://facebook.com/restauranteelsol', 'https://twitter.com/restauranteelsol', 'https://instagram.com/restauranteelsol', 'https://elsol-restaurant.com'),
('Tecnología y Más', 'clave987', 'Venta de productos electrónicos y asesoría en tecnología.', 'Avenida López Mateos 25, Colima', '312-234-5678', 'tecnologiaymas@store.com', 'tecnologiaymas.jpg', 19.2378, -103.7150, 'https://facebook.com/tecnologiaymas', 'https://twitter.com/tecnologiaymas', 'https://instagram.com/tecnologiaymas', 'https://tecnologiaymas.com'),
('Librería El Saber', 'libro1234', 'Librería con una amplia variedad de libros para todos los gustos y edades.', 'Calle de los Libros 11, Colima', '312-345-6789', 'elsaber@libreria.com', 'libreria_saber.jpg', 19.2500, -103.7350, 'https://facebook.com/libreriaelsaber', 'https://twitter.com/libreriaelsaber', 'https://instagram.com/libreriaelsaber', 'https://elsaber-libreria.com'),
('Supermercado La Esperanza', 'super2023', 'Supermercado de productos frescos y comestibles a precios bajos.', 'Calle 5 de Febrero 33, Colima', '312-432-1234', 'esperanza@supermercado.com', 'supermercado_esperanza.jpg', 19.2433, -103.7195, 'https://facebook.com/supermercadolaesperanza', 'https://twitter.com/supermercadolaesperanza', 'https://instagram.com/supermercadolaesperanza', 'https://laesperanza-supermercado.com'),
('Centro Estético Bella', 'estetica2024', 'Servicios de estética, masajes, y tratamientos de belleza.', 'Av. Hidalgo 22, Colima', '312-567-8901', 'bella@estetica.com', 'estetica_bella.jpg', 19.2550, -103.7100, 'https://facebook.com/centroesteticobella', 'https://twitter.com/centroesteticobella', 'https://instagram.com/centroesteticobella', 'https://centroesteticobella.com'),
('Cafetería Dulce Aroma', 'cafeteria123', 'Cafés y postres artesanales para disfrutar en un ambiente acogedor.', 'Calle Reforma 55, Colima', '312-678-1234', 'dulcearoma@cafeteria.com', 'cafeteria_dulcearoma.jpg', 19.2400, -103.7210, 'https://facebook.com/cafeteriadulcearoma', 'https://twitter.com/cafeteriadulcearoma', 'https://instagram.com/cafeteriadulcearoma', 'https://dulcearoma-cafe.com'),
('Gimnasio PowerFit', 'gimnasio2024', 'Gimnasio con equipos de última tecnología para entrenamiento físico.', 'Calle San Fernando 25, Colima', '312-432-2345', 'powerfit@gimnasio.com', 'gimnasio_powerfit.jpg', 19.2485, -103.7280, 'https://facebook.com/gimnasiopowerfit', 'https://twitter.com/gimnasiopowerfit', 'https://instagram.com/gimnasiopowerfit', 'https://powerfit-gimnasio.com'),
('Bar El Refugio', 'barrefugio123', 'Bar de estilo moderno con una gran variedad de cócteles y música en vivo.', 'Calle 10 de Abril 78, Colima', '312-345-8765', 'refugio@bar.com', 'bar_refugio.jpg', 19.2520, -103.7175, 'https://facebook.com/barelrefugio', 'https://twitter.com/barelrefugio', 'https://instagram.com/barelrefugio', 'https://barelrefugio.com'),
('Ferretería El Martillo', 'ferreteria1234', 'Venta de materiales de construcción, herramientas y artículos para el hogar.', 'Calle Pino 34, Colima', '312-567-2345', 'martillo@ferreteria.com', 'ferreteria_martillo.jpg', 19.2395, -103.7165, 'https://facebook.com/ferreteriaelmartillo', 'https://twitter.com/ferreteriaelmartillo', 'https://instagram.com/ferreteriaelmartillo', 'https://elmartillo-ferreteria.com'),
('Zapatería Rápido y Cómodo', 'zapateria123', 'Zapatos cómodos y modernos para toda la familia.', 'Calle 16 de Septiembre 30, Colima', '312-678-2345', 'rapidoycomodo@zapateria.com', 'zapateria_rapidoycomodo.jpg', 19.2330, -103.7205, 'https://facebook.com/zapateriarapidoycomodo', 'https://twitter.com/zapateriarapidoycomodo', 'https://instagram.com/zapateriarapidoycomodo', 'https://rapidoycomodo-zapateria.com'),
('Tianguis El Mercado', 'tianguis123', 'Venta de productos frescos y artesanales en un ambiente tradicional.', 'Calle Hidalgo 99, Colima', '312-876-5432', 'mercado@tianguis.com', 'tianguis_mercado.jpg', 19.2405, -103.7190, 'https://facebook.com/tianguiselmercado', 'https://twitter.com/tianguiselmercado', 'https://instagram.com/tianguiselmercado', 'https://tianguiselmercado.com'),
('Escuela Infantil Pequeños Genios', 'pequenosgenios2024', 'Escuela preescolar para niños, con un enfoque en el desarrollo integral.', 'Calle San Pablo 22, Colima', '312-543-9876', 'pequenosgenios@escuela.com', 'escuela_pequenosgenios.jpg', 19.2505, -103.7180, 'https://facebook.com/escuelapequenosgenios', 'https://twitter.com/escuelapequenosgenios', 'https://instagram.com/escuelapequenosgenios', 'https://pequenosgenios-escuela.com'),
('Restaurante Mariscos El Goloso', 'mariscos2024', 'Mariscos frescos y platillos del mar preparados al estilo tradicional.', 'Calle Las Palmas 12, Colima', '312-123-7890', 'mariscos@elgoloso.com', 'mariscos_elgoloso.jpg', 19.2455, -103.7265, 'https://facebook.com/mariscoselgoloso', 'https://twitter.com/mariscoselgoloso', 'https://instagram.com/mariscoselgoloso', 'https://elgoloso-mariscos.com'),
('Panadería La Esperanza', 'panaderia123', 'Panadería con productos frescos, pasteles y panes artesanales.', 'Calle del Sol 100, Colima', '312-234-6789', 'esperanza@panaderia.com', 'panaderia_esperanza.jpg', 19.2385, -103.7200, 'https://facebook.com/panaderiaesperanza', 'https://twitter.com/panaderiaesperanza', 'https://instagram.com/panaderiaesperanza', 'https://laesperanza-panaderia.com'),
('Farmacia Salud y Bienestar', 'farmacia123', 'Venta de medicamentos y productos de salud de alta calidad.', 'Av. Lázaro Cárdenas 80, Colima', '312-234-7654', 'salud@farmacia.com', 'farmacia_saludbienestar.jpg', 19.2420, -103.7180, 'https://facebook.com/farmaciasaludbienestar', 'https://twitter.com/farmaciasaludbienestar', 'https://instagram.com/farmaciasaludbienestar', 'https://saludbienestar-farmacia.com'),
('Óptica Visión Clara', 'optica2024', 'Servicios ópticos y venta de lentes de contacto y gafas de sol.', 'Calle 5 de Febrero 45, Colima', '312-876-4321', 'visionclara@optica.com', 'optica_visionclara.jpg', 19.2365, -103.7200, 'https://facebook.com/opticavisionclara', 'https://twitter.com/opticavisionclara', 'https://instagram.com/opticavisionclara', 'https://visionclara-optica.com'),
('Bazar El Tesoro', 'bazar1234', 'Tienda de artículos de segunda mano y antigüedades.', 'Calle de los Artesanos 22, Colima', '312-567-4321', 'bazar@eltesoro.com', 'bazar_eltesoro.jpg', 19.2400, -103.7155, 'https://facebook.com/bazareltesoro', 'https://twitter.com/bazareltesoro', 'https://instagram.com/bazareltesoro', 'https://eltesoro-bazar.com'),
('Centro de Idiomas Global', 'idiomas2024', 'Escuela de idiomas con cursos en inglés, francés y alemán.', 'Calle 1 de Mayo 10, Colima', '312-765-4321', 'global@idiomas.com', 'idiomas_global.jpg', 19.2410, -103.7160, 'https://facebook.com/centrodeidiomasglobal', 'https://twitter.com/centrodeidiomasglobal', 'https://instagram.com/centrodeidiomasglobal', 'https://global-idiomas.com'),
('Tienda de Ropa Chic', 'ropa2024', 'Ropa de moda para todos los gustos y ocasiones.', 'Calle Santa Clara 11, Colima', '312-678-1234', 'chic@ropa.com', 'ropa_chic.jpg', 19.2425, -103.7250, 'https://facebook.com/tiendaderopachic', 'https://twitter.com/tiendaderopachic', 'https://instagram.com/tiendaderopachic', 'https://chic-ropa.com'),
('Óptica Salud Visual', 'optica456', 'Venta de gafas graduadas y de sol, así como servicios de examen visual.', 'Calle de la Salud 21, Colima', '312-345-6789', 'saludvisual@optica.com', 'optica_saludvisual.jpg', 19.2440, -103.7180, 'https://facebook.com/opticasaludvisual', 'https://twitter.com/opticasaludvisual', 'https://instagram.com/opticasaludvisual', 'https://saludvisual-optica.com'),
('Juguetería Mundo Infantil', 'juguetes2024', 'Venta de juguetes para todas las edades, desde bebés hasta adolescentes.', 'Calle Principal 14, Colima', '312-234-5670', 'mundo@jugueteria.com', 'jugueteria_mundoinfantil.jpg', 19.2390, -103.7225, 'https://facebook.com/jugueteriamundoinfantil', 'https://twitter.com/jugueteriamundoinfantil', 'https://instagram.com/jugueteriamundoinfantil', 'https://mundoinfantil-jugueteria.com'),
('Veterinaria La Mascota Feliz', 'mascota2024', 'Veterinaria especializada en atención de mascotas y productos para animales.', 'Calle San Juan 31, Colima', '312-345-6781', 'mascotafeliz@veterinaria.com', 'veterinaria_mascotafeliz.jpg', 19.2370, -103.7260, 'https://facebook.com/veterinariamascotafeliz', 'https://twitter.com/veterinariamascotafeliz', 'https://instagram.com/veterinariamascotafeliz', 'https://mascotafeliz-veterinaria.com');

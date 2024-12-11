
CREATE TABLE DescargasApp (
    ID SERIAL PRIMARY KEY,              -- Identificador único para cada registro
    Fecha TIMESTAMP DEFAULT NOW(),      -- Fecha y hora de la descarga
    VersionApp VARCHAR(50) NOT NULL    -- Versión de la aplicación descargada
);


CREATE OR REPLACE FUNCTION sp_obtenerestadisticasmensuales()
 RETURNS TABLE(mes character varying, totalproductos integer, totalservicios integer)
 LANGUAGE plpgsql
AS $function$
BEGIN
    RETURN QUERY
    SELECT 
        Unificado.Mes::character varying, -- Referencia explícita a la columna "Mes"
        CAST(SUM(CASE WHEN Unificado.Origen = 'Producto' THEN Unificado.Total ELSE 0 END) AS integer) AS TotalProductos, -- Conversión a integer
        CAST(SUM(CASE WHEN Unificado.Origen = 'Servicio' THEN Unificado.Total ELSE 0 END) AS integer) AS TotalServicios -- Conversión a integer
    FROM (
        SELECT 
            TO_CHAR(creado::timestamp, 'MM') AS Mes, -- Generar columna "Mes"
            COUNT(*) AS Total,
            'Producto' AS Origen
        FROM Producto
        WHERE DATE_PART('year', creado::timestamp) = DATE_PART('year', CURRENT_DATE)
        GROUP BY TO_CHAR(creado::timestamp, 'MM')
        
        UNION ALL
        
        SELECT 
            TO_CHAR(creado::timestamp, 'MM') AS Mes, -- Generar columna "Mes"
            COUNT(*) AS Total,
            'Servicio' AS Origen
        FROM Servisio
        WHERE DATE_PART('year', creado::timestamp) = DATE_PART('year', CURRENT_DATE)
        GROUP BY TO_CHAR(creado::timestamp, 'MM')
    ) AS Unificado -- Asignar alias "Unificado" a la subconsulta
    GROUP BY Unificado.Mes -- Referencia explícita a "Unificado.Mes"
    ORDER BY Unificado.Mes;
END;
$function$
;





CREATE OR REPLACE FUNCTION sp_obtenerdescargaspormes()
 RETURNS TABLE(mes character varying, totaldescargas integer)
 LANGUAGE plpgsql
AS $function$
BEGIN
    RETURN QUERY
    SELECT 
        TO_CHAR(fecha::timestamp, 'YYYY-MM')::character varying AS mes, -- Conversión explícita a character varying
        COUNT(*)::integer AS totaldescargas -- Aseguro que COUNT devuelva un integer
    FROM descargasapp
    WHERE DATE_PART('year', fecha::timestamp) = DATE_PART('year', CURRENT_DATE) -- Comparación explícita con el año actual
    GROUP BY TO_CHAR(fecha::timestamp, 'YYYY-MM') -- Agrupo por el formato de fecha
    ORDER BY mes; -- Ordeno explícitamente por la columna "mes"
END;
$function$
;


CREATE OR REPLACE FUNCTION sp_obtenerbusquedaspormes()
 RETURNS TABLE(mes character varying, totalbusquedas integer)
 LANGUAGE plpgsql
AS $function$
BEGIN
    RETURN QUERY
    SELECT 
        TO_CHAR(creado::timestamp, 'YYYY-MM')::character varying AS mes, -- Conversión explícita a character varying
        COUNT(*)::integer AS totalbusquedas -- Aseguro que COUNT devuelva un integer
    FROM busquedas
    WHERE DATE_PART('year', creado::timestamp) = DATE_PART('year', CURRENT_DATE) -- Comparación explícita con el año actual
    GROUP BY TO_CHAR(creado::timestamp, 'YYYY-MM') -- Agrupo por el formato de fecha
    ORDER BY mes; -- Ordeno explícitamente por la columna "mes"
END;
$function$
;


CREATE OR REPLACE FUNCTION sp_obtenertotales()
 RETURNS TABLE(totaldescargas integer, totalnegocios integer, negociosconproductososervicios integer)
 LANGUAGE plpgsql
AS $function$
BEGIN
    RETURN QUERY
    SELECT 
        -- Total de descargas
        (SELECT COUNT(*)::integer FROM DescargasApp) AS totaldescargas,
        
        -- Total de negocios
        (SELECT COUNT(*)::integer FROM Negocio) AS totalnegocios,
        
        -- Total de negocios con al menos un producto o servicio
        (SELECT COUNT(DISTINCT n.idnegocio)::integer 
         FROM Negocio n
         LEFT JOIN Producto p ON n.idnegocio = p.idnegocio
         LEFT JOIN Servisio s ON n.idnegocio = s.idnegocio
         WHERE p.idproducto IS NOT NULL OR s.idservicio IS NOT NULL) 
         AS negociosconproductososervicios;
END;
$function$
;



CREATE OR REPLACE FUNCTION SP_ObtenerTotalesPorDia()
 RETURNS TABLE(totalnegocios integer, totalservicios integer, totalproductos integer, totalbusquedas integer)
 LANGUAGE plpgsql
AS $function$
BEGIN
    RETURN QUERY
    SELECT              
        -- Total de negocios por mes (convertido a integer)
        (SELECT COUNT(*)::integer FROM Negocio ) AS TotalNegocios,
        
        -- Total de servicios por mes (convertido a integer)
        (SELECT COUNT(*)::integer FROM Servisio ) AS TotalServicios,
        
        -- Total de productos por mes (convertido a integer)
        (SELECT COUNT(*)::integer FROM Producto ) AS TotalProductos,
        
        -- Total de busquedas por mes (convertido a integer)
        (SELECT COUNT(*)::integer FROM Busquedas ) AS TotalBusquedas;
END;
$function$
;



-- Insertando 30 registros con fechas aleatorias dentro de los últimos 3 meses y versiones de la app
INSERT INTO DescargasApp (Fecha, VersionApp) 
VALUES
('2024-09-10 14:23:00', '1.0.0'),
('2024-09-12 08:45:00', '1.0.1'),
('2024-09-15 16:30:00', '1.0.0'),
('2024-09-17 10:22:00', '1.0.2'),
('2024-09-20 12:10:00', '1.1.0'),
('2024-09-22 18:00:00', '1.1.0'),
('2024-09-24 09:34:00', '1.0.1'),
('2024-09-27 14:55:00', '1.1.0'),
('2024-10-01 11:11:00', '1.1.0'),
('2024-10-03 13:30:00', '1.0.2'),
('2024-10-05 09:00:00', '1.0.0'),
('2024-10-07 17:23:00', '1.1.0'),
('2024-10-09 15:30:00', '1.0.1'),
('2024-10-11 19:00:00', '1.0.2'),
('2024-10-13 10:15:00', '1.1.0'),
('2024-10-15 14:22:00', '1.0.2'),
('2024-10-17 16:40:00', '1.0.1'),
('2024-10-20 12:23:00', '1.1.0'),
('2024-10-22 09:30:00', '1.1.0'),
('2024-10-25 14:10:00', '1.0.0'),
('2024-10-27 11:50:00', '1.0.1'),
('2024-10-29 18:22:00', '1.1.0'),
('2024-10-31 13:34:00', '1.0.0'),
('2024-11-02 16:00:00', '1.1.0'),
('2024-11-05 12:45:00', '1.0.2'),
('2024-11-07 08:50:00', '1.0.1'),
('2024-11-09 17:00:00', '1.1.0'),
('2024-11-12 14:30:00', '1.0.2'),
('2024-11-14 10:10:00', '1.0.1'),
('2024-11-16 13:45:00', '1.1.0');

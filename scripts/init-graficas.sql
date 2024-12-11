
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
            TO_CHAR(creado::timestamp, 'YYYY-MM') AS Mes, -- Generar columna "Mes"
            COUNT(*) AS Total,
            'Producto' AS Origen
        FROM Producto
        WHERE DATE_PART('year', creado::timestamp) = DATE_PART('year', CURRENT_DATE)
        GROUP BY TO_CHAR(creado::timestamp, 'YYYY-MM')
        
        UNION ALL
        
        SELECT 
            TO_CHAR(creado::timestamp, 'YYYY-MM') AS Mes, -- Generar columna "Mes"
            COUNT(*) AS Total,
            'Servicio' AS Origen
        FROM servisio
        WHERE DATE_PART('year', creado::timestamp) = DATE_PART('year', CURRENT_DATE)
        GROUP BY TO_CHAR(creado::timestamp, 'YYYY-MM')
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
 RETURNS TABLE(fecha date, totalnegocios integer, totalservicios integer, totalproductos integer, totalbusquedas integer)
 LANGUAGE plpgsql
AS $function$
BEGIN
    RETURN QUERY
    WITH fecha_calculada AS (
        SELECT 
            (TO_CHAR(n.Creado, 'YYYY-MM') || '-01')::date AS Fecha,  -- Añadir '-01' para convertir a fecha completa
            n.IDNegocio
        FROM Negocio n
    )
    SELECT 
        fc.Fecha,
        
        -- Total de negocios por mes (convertido a integer)
        (SELECT COUNT(*)::integer 
         FROM Negocio 
         WHERE TO_CHAR(Creado, 'YYYY-MM') = TO_CHAR(fc.Fecha, 'YYYY-MM')) AS TotalNegocios,
        
        -- Total de servicios por mes (convertido a integer)
        (SELECT COUNT(*)::integer 
         FROM Servisio 
         WHERE TO_CHAR(Creado, 'YYYY-MM') = TO_CHAR(fc.Fecha, 'YYYY-MM')) AS TotalServicios,
        
        -- Total de productos por mes (convertido a integer)
        (SELECT COUNT(*)::integer 
         FROM Producto 
         WHERE TO_CHAR(Creado, 'YYYY-MM') = TO_CHAR(fc.Fecha, 'YYYY-MM')) AS TotalProductos,
        
        -- Total de busquedas por mes (convertido a integer)
        (SELECT COUNT(*)::integer 
         FROM Busquedas 
         WHERE TO_CHAR(Creado, 'YYYY-MM') = TO_CHAR(fc.Fecha, 'YYYY-MM')) AS TotalBusquedas
        
    FROM fecha_calculada fc
    GROUP BY fc.Fecha
    ORDER BY fc.Fecha DESC;
END;
$function$
;


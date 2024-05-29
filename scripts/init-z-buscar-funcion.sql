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

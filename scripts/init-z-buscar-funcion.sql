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
    SELECT 
        n.idnegocio ,
        p.idproducto ,
        s.idservicio ,
        n.nombre as negocio,
        COALESCE(p.nombre , s.nombre) AS nombre,
        COALESCE(p.descripsion , s.descripcion, n.descripsion) AS descripcion,
        COALESCE(p.imagen , s.imagen, n.imagen  ) AS imagen
    FROM 
        Negocio n
    LEFT JOIN 
        Servisio s ON s.idnegocio  = n.idnegocio 
    LEFT JOIN 
        Producto p ON p.idnegocio  = n.idnegocio
    WHERE
        n.nombre ILIKE '%' || p_termino || '%'
        OR p.nombre ILIKE '%' || p_termino || '%'
        OR s.nombre ILIKE '%' || p_termino || '%'
        OR n.descripsion ILIKE '%' || p_termino || '%'
        OR s.descripcion ILIKE '%' || p_termino || '%'
        OR p.descripsion ILIKE '%' || p_termino || '%';
END;
$$ LANGUAGE plpgsql;

# Negocios Endpoint

## ver Negocios `GET /negocios/id`
  Nos trae la información detallada de un Negocios en especifico
  ### Parametros de URL
  - `{id}` (obligatorio): Identificador único del usuario que se desea recuperar.
  
  ### Ejemplo de Solicitud
  ```http
    GET /negocios/123
  ```

  ### Respuesta Exitosa (Código 200 OK)
  ```json
    {
        "Id_Negocio":0,
        "Nombre":"",
        "Descripsion":"",
        "Direccion":"",
        "telefono":"",
        "correo":"",
        "Latitude":"",
        "Longitude":"",
        "Facebook":"",
        "Twitter":"",
        "Instagram":"",
        "Website":"",
        "Creado":"0001-01-01T00:00:00Z",
        "Actualizado":"0001-01-01T00:00:00Z"
    }
  ```

  ### Respuestas de Errores Posibles
 
  - Código 400 Bad Request:
    ```json
      {
      "errno": 400,
      "error": "bad_request",
      "error_description": "Dato ${%d} invalido."
      }
    ```

  - Código 401 Unauthorized:
    ```json
      {
      "errno": 401,
      "error": "unauthorized",
      "error_description": "Tienes que registrarte para realizar la solicit"
      }
    ```

  - Código 403 Forbidden:
    ```json
      {
      "errno": 403,
      "error": "forbidden",
      "error_description": "No tiens permisos para realizar la solicit"
      }
    ```

  - Código 404 Not Found:
    ```json
      {
      "errno": 404,
      "error": "not_found",
      "error_description": "No se encontró el recurso."
      }
    ```

  - Código 500 Internal Server Error:
    ```json
      {
      "errno": 500,
      "error": "internal_error",
      "error_description": "Ocurrió un problema para procesar la solicitud"
      }
    ``` 

## Crear Negocio `POST /negocios`
  Nos permitira crear un nuevo negocio
  
  ### Ejemplo de Solicitud
  ```http
    POST /negocios
  ```
  ```json
    {
        "Nombre":"",
        "Descripsion":"",
        "Direccion":"",
        "telefono":"",
        "correo":"",
        "Latitude":"",
        "Longitude":"",
        "Facebook":"",
        "Twitter":"",
        "Instagram":"",
        "Website":""
    }
  ``` 
 
  ### Respuesta Exitosa (Código 201 Created) 
  ```json
    {
        "Id_Negocio":0,
        "Nombre":"",
        "Descripsion":"",
        "Direccion":"",
        "telefono":"",
        "correo":"",
        "Latitude":"",
        "Longitude":"",
        "Facebook":"",
        "Twitter":"",
        "Instagram":"",
        "Website":"",
        "Creado":"0001-01-01T00:00:00Z",
        "Actualizado":"0001-01-01T00:00:00Z"
    }
  ```

  ### Respuestas de Errores Posibles
  
  - Código 400 Bad Request:
    ```json
        {
            "errno": 400,
            "error": "bad_request",
            "error_description": "Dato ${%d} invalido."
        }
    ```
  
  - Código 401 Unauthorized:
    ```json
        {
            "errno": 401,
            "error": "unauthorized",
            "error_description": "Tienes que registrarte para realizar la solicit"
        }
    ```
  
  - Código 403 Forbidden:
    ```json
        {
            "errno": 403,
            "error": "forbidden",
            "error_description": "No tiens permisos para realizar la solicit"
        }
    ```
  
  - Código 404 Not Found:
    ```json
        {
            "errno": 404,
            "error": "not_found",
            "error_description": "No se encontró el Uasuario."
        }
    ```
  
  - Código 500 Internal Server Error:
    ```json
        {
            "errno": 500,
            "error": "internal_error",
            "error_description": "Ocurrió un problema para procesar la solicitud"
        }
    ``` 


## Editar Negocio `PUT /negocios/id`

  Edita la informacion del negosip

  ### Parametros de URL
  - `{id}` (obligatorio): Identificador único del usuario que se desea recuperar.
  
  ### Ejemplo de Solicitud
  ```http
    PUT /negocios/123
  ```
  ```json
    {
        "Nombre":"",
        "Descripsion":"",
        "Direccion":"",
        "telefono":"",
        "correo":"",
        "Latitude":"",
        "Longitude":"",
        "Facebook":"",
        "Twitter":"",
        "Instagram":"",
        "Website":""
    }
  ``` 
  
  ### Respuesta Exitosa (Código 200 OK)
  ```json
    {
        "Id_Negocio":0,
        "Nombre":"",
        "Descripsion":"",
        "Direccion":"",
        "telefono":"",
        "correo":"",
        "Latitude":"",
        "Longitude":"",
        "Facebook":"",
        "Twitter":"",
        "Instagram":"",
        "Website":"",
        "Creado":"0001-01-01T00:00:00Z",
        "Actualizado":"0001-01-01T00:00:00Z"
    }
  ```

  ### Respuestas de Errores Posibles

  - Código 400 Bad Request:
    ```json
        {
            "errno": 400,
            "error": "bad_request",
            "error_description": "Dato ${%d} invalido."
        }
    ```
  
  - Código 401 Unauthorized:
    ```json
        {
            "errno": 401,
            "error": "unauthorized",
            "error_description": "Tienes que registrarte para realizar la solicit"
        }
    ```
  
  - Código 403 Forbidden:
    ```json
        {
            "errno": 403,
            "error": "forbidden",
            "error_description": "No tiens permisos para realizar la solicit"
        }
    ```
  
  - Código 404 Not Found:
    ```json
        {
            "errno": 404,
            "error": "not_found",
            "error_description": "No se encontró el Uasuario."
        }
    ```
  
  - Código 500 Internal Server Error:
    ```json
        {
            "errno": 500,
            "error": "internal_error",
            "error_description": "Ocurrió un problema para procesar la solicitud"
        }
    ``` 

## Eliminar Negocio: `DELETE /negocios/id`
  Elimina el un negosio espesifico

  ### Parametros de URL
  - `{id}` (obligatorio): Identificador único del usuario que se desea recuperar.

  ### Ejemplo de Solicitud
  ```http
      DELETE /negocios/123
  ```

  ### Respuesta Exitosa (204 No Content)
  ```json
  
  ```  
  
  ### Respuestas de Errores Posibles

  - Código 400 Bad Request:
    ```json
        {
            "errno": 400,
            "error": "bad_request",
            "error_description": "Dato ${%d} invalido."
        }
    ```
  
  - Código 401 Unauthorized:
    ```json
        {
            "errno": 401,
            "error": "unauthorized",
            "error_description": "Tienes que registrarte para realizar la solicit"
        }
    ```
  
  - Código 403 Forbidden:
    ```json
        {
            "errno": 403,
            "error": "forbidden",
            "error_description": "No tiens permisos para realizar la solicit"
        }
    ```
  
  - Código 404 Not Found:
    ```json
        {
            "errno": 404,
            "error": "not_found",
            "error_description": "No se encontró el Uasuario."
        }
    ```
  
  - Código 500 Internal Server Error:
    ```json
        {
            "errno": 500,
            "error": "internal_error",
            "error_description": "Ocurrió un problema para procesar la solicitud"
        }
    ``` 




## ver productos `GET /negocios/id_negocio/producto`
  Nostraera la lista de todos los productos con los que cuenta el negocio
  ### Parametros de URL
  - `{id_negocio}` (obligatorio): Identificador único del usuario que se desea recuperar.
  
  ### Ejemplo de Solicitud
  ```http
    GET /negocios/123/producto
  ```

  ### Respuesta Exitosa (Código 200 OK)
  ```json
    [
        {
            "Id_Negocio":0,
            "Nombre":"",
            "Descripsion":"",
            "Direccion":"",
            "telefono":"",
            "correo":"",
            "Latitude":"",
            "Longitude":"",
            "Facebook":"",
            "Twitter":"",
            "Instagram":"",
            "Website":"",
            "Creado":"0001-01-01T00:00:00Z",
            "Actualizado":"0001-01-01T00:00:00Z"
        },
        ....
    ]
  ```

  ### Respuestas de Errores Posibles
 
  - Código 400 Bad Request:
    ```json
      {
      "errno": 400,
      "error": "bad_request",
      "error_description": "Dato ${%d} invalido."
      }
    ```

  - Código 401 Unauthorized:
    ```json
      {
      "errno": 401,
      "error": "unauthorized",
      "error_description": "Tienes que registrarte para realizar la solicit"
      }
    ```

  - Código 403 Forbidden:
    ```json
      {
      "errno": 403,
      "error": "forbidden",
      "error_description": "No tiens permisos para realizar la solicit"
      }
    ```

  - Código 404 Not Found:
    ```json
      {
      "errno": 404,
      "error": "not_found",
      "error_description": "No se encontró el recurso."
      }
    ```

  - Código 500 Internal Server Error:
    ```json
      {
      "errno": 500,
      "error": "internal_error",
      "error_description": "Ocurrió un problema para procesar la solicitud"
      }
    ``` 

## Crear producto `POST /negocios/id_negocio/producto`
  Agregar un nuevo producto para el negocio
  
  ### Ejemplo de Solicitud
  ```http
    POST /negocios/123/producto
  ```
  ```json
    {
        "Nombre":"",
        "Descripsion":"",
        "Unidad":0,

    }
  ``` 
 
  ### Respuesta Exitosa (Código 201 Created) 
  ```json
    {
        "Id_Producto":0,
        "Nombre":"",
        "Descripsion":"",
        "Unidad":0,
        "Creado":"0001-01-01T00:00:00Z",
        "Actualizado":"0001-01-01T00:00:00Z"
    }
  ```

  ### Respuestas de Errores Posibles
  
  - Código 400 Bad Request:
    ```json
        {
            "errno": 400,
            "error": "bad_request",
            "error_description": "Dato ${%d} invalido."
        }
    ```
  
  - Código 401 Unauthorized:
    ```json
        {
            "errno": 401,
            "error": "unauthorized",
            "error_description": "Tienes que registrarte para realizar la solicit"
        }
    ```
  
  - Código 403 Forbidden:
    ```json
        {
            "errno": 403,
            "error": "forbidden",
            "error_description": "No tiens permisos para realizar la solicit"
        }
    ```
  
  - Código 404 Not Found:
    ```json
        {
            "errno": 404,
            "error": "not_found",
            "error_description": "No se encontró el Uasuario."
        }
    ```
  
  - Código 500 Internal Server Error:
    ```json
        {
            "errno": 500,
            "error": "internal_error",
            "error_description": "Ocurrió un problema para procesar la solicitud"
        }
    ``` 


## Editar producto `PUT /negocios/id_negocio/producto/id_producto`

  Editar el producto

  ### Parametros de URL
  - `{id_negocio}` (obligatorio): Identificador único del usuario que se desea recuperar.
  - `{id_producto}` (obligatorio): Identificador único del usuario que se desea recuperar.
  
  ### Ejemplo de Solicitud
  ```http
    PUT /negocios/123/producto/123
  ```
  ```json
    {
        "Nombre":"",
        "Descripsion":"",
        "Unidad":0
    }
  ``` 
  
  ### Respuesta Exitosa (Código 200 OK)
  ```json
    {
        "Id_Producto":0,
        "Nombre":"",
        "Descripsion":"",
        "Unidad":0,
        "Creado":"0001-01-01T00:00:00Z",
        "Actualizado":"0001-01-01T00:00:00Z"
    }
  ```

  ### Respuestas de Errores Posibles

  - Código 400 Bad Request:
    ```json
        {
            "errno": 400,
            "error": "bad_request",
            "error_description": "Dato ${%d} invalido."
        }
    ```
  
  - Código 401 Unauthorized:
    ```json
        {
            "errno": 401,
            "error": "unauthorized",
            "error_description": "Tienes que registrarte para realizar la solicit"
        }
    ```
  
  - Código 403 Forbidden:
    ```json
        {
            "errno": 403,
            "error": "forbidden",
            "error_description": "No tiens permisos para realizar la solicit"
        }
    ```
  
  - Código 404 Not Found:
    ```json
        {
            "errno": 404,
            "error": "not_found",
            "error_description": "No se encontró el Uasuario."
        }
    ```
  
  - Código 500 Internal Server Error:
    ```json
        {
            "errno": 500,
            "error": "internal_error",
            "error_description": "Ocurrió un problema para procesar la solicitud"
        }
    ``` 

## Eliminar producto: `DELETE /negocios/id_negocio/producto/id_producto`
  Elimina el producto seleccionado del negocio

  ### Parametros de URL
  - `{id_negocio}` (obligatorio): Identificador único del usuario que se desea recuperar.

  ### Ejemplo de Solicitud
  ```http
      DELETE /negocios/123/producto/123
  ```

  ### Respuesta Exitosa (204 No Content)
  ```json
  
  ```  
  
  ### Respuestas de Errores Posibles

  - Código 400 Bad Request:
    ```json
        {
            "errno": 400,
            "error": "bad_request",
            "error_description": "Dato ${%d} invalido."
        }
    ```
  
  - Código 401 Unauthorized:
    ```json
        {
            "errno": 401,
            "error": "unauthorized",
            "error_description": "Tienes que registrarte para realizar la solicit"
        }
    ```
  
  - Código 403 Forbidden:
    ```json
        {
            "errno": 403,
            "error": "forbidden",
            "error_description": "No tiens permisos para realizar la solicit"
        }
    ```
  
  - Código 404 Not Found:
    ```json
        {
            "errno": 404,
            "error": "not_found",
            "error_description": "No se encontró el Uasuario."
        }
    ```
  
  - Código 500 Internal Server Error:
    ```json
        {
            "errno": 500,
            "error": "internal_error",
            "error_description": "Ocurrió un problema para procesar la solicitud"
        }
    ``` 




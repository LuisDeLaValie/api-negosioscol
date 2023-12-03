
# Buesquedas Endpoint


## ver productos `GET /buscar?buscar=...&producto=true&servisio=true&negocio=true`
  Buscar negosios por distintos medios ya seace por `productos`, `servisios` o `negocios`
  ### Queris de URL
  - `{buscar}` (obligatorio): cadena de texto que se usara para buscar
  - `{producto}` nos ayudara para saber que lo que buscamos es un producto
  - `{servisio}` nos ayudara para saber que lo que buscamos es un producto
  - `{negocio}` nos ayudara para saber que lo que buscamos es un producto
  
  ### Ejemplo de Solicitud
  ```http
    GET /buscar?buscar=pinzas%20depresión \
        &producto=true \
        &servisio=true \
        &negocio=true
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
  - Código 500 Internal Server Error:
    ```json
      {
      "errno": 500,
      "error": "internal_error",
      "error_description": "Ocurrió un problema para procesar la solicitud"
      }
    ``` 

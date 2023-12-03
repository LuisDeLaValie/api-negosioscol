# Uasuarios Endpoints

## Ver Usuario: `GET /usuarios/id`
  <details>
    <summary>Mostar</summary>

  Nos trae la información detallada de un usuario en especifico

  ### Parámetros de URL
  - `{id}` (obligatorio): Identificador único del usuario que se desea recuperar.

  ### Ejemplo de Solicitud
  ```http
    GET /usuarios/123
  ```

  ### Respuesta Exitosa (Código 200 OK)
  ```json
    {
      "Id_Usuario": 0,
      "Nombre": "",
      "Apellidos": "",
      "Creado": "0001-01-01T00:00:00Z",
      "Actualizado": "0001-01-01T00:00:00Z",
      "Cumpleanos": "0001-01-01T00:00:00Z",
      "Imagen": ""
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

</details>



  </details>



## Editar Usuario: `PUT /usuarios/id`
<details>
  <summary>Mostar</summary>

  Nos permitira poder editar la información un usuario 

  ### Parámetros de URL
  - `{id}` (obligatorio): Identificador único del tema que se desea recuperar.

  ### Ejemplo de Solicitud
  ```http
    PUT /usuarios/123  
  ```
  ```json
    {
      "Nombre": "",
      "Apellidos": "",
      "Imagen": ""
    }
  ```

  ### Respuesta Exitosa (Código 200 OK)
  ```json
  {
    "Id_Usuario": 0,
    "Nombre": "",
    "Apellidos": "",
    "Creado": "0001-01-01T00:00:00Z",
    "Actualizado": "0001-01-01T00:00:00Z",
    "Cumpleanos": "0001-01-01T00:00:00Z",
    "Imagen": ""
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

</details>



## Eliminar Usuario: `DELETE /usuarios/id`
<details>
  <summary>Mostar</summary>

  eliminar el usuario espesificado

  ### Parámetros de URL
  - `{id}` (obligatorio): Identificador único del tema que se desea recuperar.

  ### Ejemplo de Solicitud
  ```http
    DELETE /usuarios/123
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

</details>




## Agregar Usuario: `POST /usuarios`
<details>
  <summary>Mostar</summary>

  Nos permitira crear un nuevo usuario

  ### Ejemplo de Solicitud
  ```http
    POST /usuarios
  ```

  ### Respuesta Exitosa (Código 201 Created)
  ```json
    {
    "Id_Usuario": 0,
    "Nombre": "",
    "Apellidos": "",
    "Creado": "0001-01-01T00:00:00Z",
    "Actualizado": "0001-01-01T00:00:00Z",
    "Cumpleanos": "0001-01-01T00:00:00Z",
    "Imagen": ""
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

</details>





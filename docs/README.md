# Documentación de la API RESTful

Bienvenido a la documentación de la API. A continuación, se presenta una lista
de endpoints disponibles con enlaces a sus respectivas documentaciones detalladas.

## Descripción General de Endpoints

### Negocios 
| Recurso                    | Descripción |
| -------------------------- | ----------- |
| [`GET /buscar`](./endpoints/busquedas.md#ver-productos-get-buscarbuscarproductotrueservisiotruenegociotrue)         | Buscar negosios por distintos medios ya seace |

### Usuario 
| Recurso                    | Descripción |
| -------------------------- | ----------- |
| [`GET /usuarios/id`](./endpoints/usuarios.md#ver-usuario-get-usuariosid)         | Nos trae la información detallada de un usuario en especifico |
| [`PUT /usuarios/id`](./endpoints/usuarios.md#editar-usuario-put-usuariosid)          | Nos permitira poder editar la información un usuario |
| [`DELETE /usuarios/id`](./endpoints/usuarios.md#eliminar-usuario-delete-usuariosid)   | eliminar el usuario espesificado |
| [`POST /usuarios`](./endpoints/usuarios.md#agregar-usuario-post-usuarios)              | Nos permitira crear un nuevo usuario |

### Negocios 
| Recurso                    | Descripción |
| -------------------------- | ----------- |
| [`GET /negocios/id`](./endpoints/negosios.md#ver-negocios-get-negociosid)         | Nos trae la información detallada de un Negocios en especifico |
| [`POST /negocios`](./endpoints/negosios.md#crear-negocio-post-negocios)         | Nos permitira crear un nuevo negocio |
| [`PUT /negocios/id`](./endpoints/negosios.md#editar-negocio-put-negociosid)         | Edita la informacion del negosip |
| [`DELETE /negocios/id`](./endpoints/negosios.md#eliminar-negocio-delete-negociosid)         | Elimina el un negosio espesifico |
| [`GET /negocios/id_negocio/producto`](./endpoints/negosios.md#ver-productos-get-negociosid_negocioproducto)         | Nostraera la lista de todos los productos con los que cuenta el negocio |
| [`POST /negocios/id_negocio/producto`](./endpoints/negosios.md#crear-producto-post-negociosid_negocioproducto)         | Agregar un nuevo producto para el negocio |
| [`PUT /negocios/id_negocio/producto/id_producto`](./endpoints/negosios.md#editar-producto-put-negociosid_negocioproductoid_producto)         | Editar el producto |
| [`DELETE /negocios/id_negocio/producto/id_producto`](./endpoints/negosios.md#eliminar-producto-delete-negociosid_negocioproductoid_producto)         | Elimina el producto seleccionado del negocio |
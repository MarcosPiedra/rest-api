# Cómo ejecutar
Para poder ejecutarlo solo hay que tener el docker-compose instalado y ejecutar las siguientes líneas en una terminal:

1) ``` docker-compose build ```
2) ``` docker-compose up ```

Se puede acceder al swagger de la API:

``` http://localhost:8080/swagger/index.html ```

# Arquitectura

Esta solución se ha basado en diferentes arquitecturas/paradigmas para el desarrollo del backend:
- Modular: Es un monolito modular, donde se puede escalar añadiendo módulos (o partes del negocio).
- CQRS: Se han diferenciando tanto las consultas (query) como los comandos (command).
- Modelo de dominio anémico
- Multicapa: Se ha trabajado con las diferentes capas aplicación, infraestructura, dominio y API.
- Principios SOLID, en la que destacaría DI puro.

# Estructura

- backend: Carpeta que contiene la solución en golang del backend.
  - internal: Paquetes transversales a todo los módulos
  - doctors: Módulo doctors
    - Capas de aplicación, dominio, infraestructura y el API
  - docker: Dockerfile
  - migraciones: Archivos SQL de migración mediante goose

# A mejorar:
- Añadir test de integración en el backend.
- Añadir E2E test en el backend.

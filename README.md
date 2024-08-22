# TodoList BFF (Backend for Frontend)

## Descripción

Este proyecto es un Backend for Frontend (BFF) para una aplicación de TodoList.
Está construido con Go, utilizando Gin como framework web y GORM como ORM para interactuar
con una base de datos PostgreSQL. La aplicación está dockerizada para facilitar su despliegue y desarrollo.

## Características

- API RESTful para gestionar tareas (todos)
- Arquitectura limpia (Clean Architecture)
- Uso de GORM para operaciones de base de datos
- Dockerizado para fácil despliegue y desarrollo
- Configuración flexible a través de variables de entorno

## Requisitos previos

- Docker
- Docker Compose

## Configuración

1. Clona el repositorio:

   ```
   git clone https://github.com/kevjarvis/todolist-bff.git
   cd todolist-bff
   ```

2. Crea un archivo `.env` en la raíz del proyecto con las siguientes variables:

   ```
   DB_USER=postgres
   DB_PASSWORD=password
   ```

   Nota: Asegúrate de no subir este archivo al control de versiones.

3. Opcionalmente, puedes modificar otras variables de entorno en el archivo `docker-compose.yml` según tus necesidades.

## Cómo levantar la aplicación

1. Construye y levanta los contenedores:

   ```
   docker-compose up --build
   ```

   Esto construirá la imagen de la aplicación y levantará los contenedores de la app y la base de datos.

2. Para ejecutar en segundo plano:

   ```
   docker-compose up -d
   ```

3. Para detener la aplicación:

   ```
   docker-compose down
   ```

4. Si necesitas reconstruir la aplicación después de cambios:
   ```
   docker-compose up --build
   ```

## Uso de la API

Una vez que la aplicación esté en ejecución, puedes acceder a la API en `http://localhost:8080`.

Endpoints disponibles:

- `GET /todos`: Obtener todas las tareas
- `GET /todos/:id`: Obtener una tarea específica
- `POST /todos`: Crear una nueva tarea

## Desarrollo

Para desarrollo local sin Docker:

1. Asegúrate de tener Go instalado en tu sistema.
2. Instala las dependencias:
   ```
   go mod download
   ```
3. Configura las variables de entorno o usa un archivo `.env`.
4. Ejecuta la aplicación:
   ```
   go run cmd/main.go
   ```

## Configuración avanzada

- Puedes modificar la configuración de la base de datos y otros parámetros en `pkg/config/config.go`.
- Para cambiar el puerto de la aplicación, modifica la variable `SERVER_PORT` en el `docker-compose.yml`.

## Contribuciones

Las contribuciones son bienvenidas. Por favor, abre un issue para discutir cambios mayores antes de enviar un pull request.

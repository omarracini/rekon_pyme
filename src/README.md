# Rekon Pyme Backend

Sistema de conciliación de facturas para pyme desarrollado en Go bajo arquitectura hexagonal.

## Requisitos
- Docker y Docker Compose
- Go 1.21+

## Instalación
1. Levantar la base de datos: `docker-compose up -d`
2. Ejecutar el servidor: `go run main.go`

## Estructura del Proyecto
- `src/banking/domain`: Entidades y reglas de negocio.
- `src/banking/application`: Casos de uso.
- `src/banking/infrastructure`: Implementación de Postgres y Gin.
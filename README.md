# DOCUMENTACIÓN Y ARQUITECTURA

## 1. Problema Elegido y Justificación

Problema: Fragmentación de la conciliación bancaria en PYMEs con operaciones multimoneda. Justificación: Las empresas actuales operan globalmente, pero los sistemas contables suelen fallar al cruzar movimientos en diferentes divisas (USD, COP, MXN). Este proyecto resuelve la brecha entre el extracto bancario y la facturación, asegurando integridad financiera y reduciendo el error humano mediante validaciones de base de datos y lógica de dominio.

## 2. Instrucciones de Entorno (Docker)

Para levantar el sistema completo:

1. Asegúrate de tener Docker y Docker Compose instalados.
2. Ejecuta en la raíz del proyecto: docker-compose up -d
3. Para cargar los datos de prueba (Seeds) con soporte para USD, COP y MXN: docker exec -i [NOMBRE_CONTENEDOR_DB] psql -U user_pyme -d conciliacion_db < scripts/seed_data.sql

## 3. Funcionalidad de IA

Motor de Emparejamiento Inteligente (AI Matching). El sistema utiliza lógica de dominio para sugerir la factura más probable para un movimiento bancario basándose en:

Coincidencia exacta de monto y tipo de moneda (USD/COP/MXN).

Proximidad de fechas y similitud semántica en los conceptos de pago.

Desacoplamiento: La lógica está aislada en la capa de aplicación, permitiendo cambiar el motor de reglas por un modelo de Machine Learning sin afectar el core del negocio.

# TEST Y CALIDAD DEL CÓDIGO

## Comandos de Ejecución

Para validar la lógica de dominio (Arquitectura Hexagonal):

1. Ejecutar todos los tests del backend: go test ./... -v
2. Test específico de la lógica de conciliación multimoneda: go test ./internal/domain/services -run TestConciliationLogic

# ARQUITECTURA DEL SISTEMA

El proyecto está diseñado bajo los principios de Arquitectura Hexagonal (Puertos y Adaptadores), lo que garantiza un desacoplamiento total entre la lógica de negocio y las tecnologías externas.

Dominio (/internal/domain): Contiene las entidades core (Movement, Invoice, Conciliation) y las interfaces que definen el comportamiento del sistema. Aquí reside la lógica de validación multimoneda.

Aplicación (/internal/application): Orquesta los casos de uso, como el proceso de conciliación y el motor de "AI Matching".

Infraestructura (/internal/infrastructure): Implementa las adaptaciones técnicas: GORM para PostgreSQL, Gin para la API REST y la comunicación con el frontend en Next.js.

Nota sobre el Diagrama: Se incluye el diagrama de arquitectura y el diagrama de Entidad-Relación en la carpeta /docs del proyecto.

# CRITERIOS DE EVALUACIÓN CUBIERTOS
Para facilitar la revisión, se detallan los puntos clave implementados:

Modelado Financiero: Uso estricto de tipos de datos para precisión decimal y manejo de códigos ISO para divisas (USD, COP, MXN, EUR).

Calidad de Código: Implementación de Arquitectura Hexagonal y Tests Unitarios sobre la lógica de dominio.

Dockerización: Orquestación completa mediante docker-compose para base de datos y servicios.

IA e Integración: Motor de sugerencias basado en reglas de negocio desacoplado, listo para ser escalado a modelos predictivos.

Observabilidad: Logs estructurados en el backend para monitorear el flujo de transacciones y conciliaciones.

# 
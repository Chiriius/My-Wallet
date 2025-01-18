# Makefile

# Ejecutar pruebas en todos los paquetes excepto en 'utils' y 'jwt'
test:
    @echo "Ejecutando pruebas, excluyendo los paquetes 'utils', 'jwt', y otros..."
    go test $(shell go list ./... | grep -v '/utils' ) --cover

# Limpiar archivos generados por las pruebas
clean:
    @echo "Limpiando los archivos generados por las pruebas..."
    go clean

# Ejecutar todas las pruebas en el proyecto
all: test

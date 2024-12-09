import requests

# URL de tu API local (reemplaza <id> con el ID del recurso que deseas eliminar)
url = 'http://localhost:3000/tasks/2'  # Cambia el 1 por el ID correspondiente

# Realizar la solicitud DELETE
response = requests.delete(url)

# Verificar el código de estado y el contenido de la respuesta
if response.status_code == 204:  # Código 204 No Content indica que se eliminó con éxito
    print('Recurso eliminado exitosamente.')
else:
    print(f'Error {response.status_code}: {response.text}')

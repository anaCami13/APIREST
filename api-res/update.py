import requests
import json

 # URL de tu API local (reemplaza <id> con el ID del recurso que deseas actualizar)
url = 'http://localhost:3000/tasks/2'  # Cambia el 1 por el ID correspondiente

# Datos que deseas enviar para la actualización
data = {
    'Name': 'Actividad Dos',
    'Content': 'Ver familia de diez en la TV con Arturo'
}

# Realizar la solicitud PUT
response = requests.put(url, headers={'Content-Type': 'application/json'}, data=json.dumps(data))

# Verificar el código de estado y el contenido de la respuesta
if response.status_code == 200:
    print('Recurso actualizado exitosamente:', response.json())
else:
    print(f'Error {response.status_code}: {response.text}')

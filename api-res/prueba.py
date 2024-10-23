import requests
import json

# URL a la que envías la solicitud POST
url = 'http://localhost:3000/newtask'

# Datos que quieres enviar
data = {
    'Name': "Actividad Tres",
    'Content': "Escribir articulos de alta calidad"
}

# Realizar la solicitud POST con el formato JSON
response = requests.post(url, headers={'Content-Type': 'application/json'}, data=json.dumps(data))

# Verificar el código de estado y el contenido de la respuesta
if response.status_code == 200:
    print('Datos enviados exitosamente:', response.json())
else:
    print(f'Error {response.status_code}: {response.text}')

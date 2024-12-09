 // Función para cargar las actividades y mostrarlas en la lista
async function cargarActividades() {
    try {
        const response = await fetch("http://172.21.0.2/:32000/tasks", {
            method: "GET"
        });
        
        if (!response.ok) {
            throw new Error("Error al obtener las actividades");
        }
        
        const actividades = await response.json();
        const actividadLista = document.getElementById("actividad-lista");
        actividadLista.innerHTML = ""; // Limpia la lista antes de agregar nuevas actividades
        
        actividades.forEach(actividad => {
            const fila = document.createElement("tr");
            fila.innerHTML = `
                <td>${actividad.ID}</td>
                <td>${actividad.Name}</td>
                <td>${actividad.Content}</td>
            `;
            actividadLista.appendChild(fila);
        });
    } catch (error) {
        console.error(error);
        alert("Error al cargar las actividades");
    }
}

// Llamada inicial para cargar las actividades cuando la página se carga
document.addEventListener("DOMContentLoaded", cargarActividades);


// Función para agregar una nueva actividad
async function agregarActividad() {
    const nombre = document.getElementById("nombre").value;
    const descripcion = document.getElementById("descripcion").value;

    try {
        const response = await fetch("http://172.21.0.2:32000/tasks", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ Name: nombre, Content: descripcion })
        });

        if (response.ok) {
            alert("Actividad agregada correctamente.");
            cargarActividades();
        } else {
            throw new Error("Error al agregar la actividad.");
        }
    } catch (error) {
        console.error(error);
        alert("Error al agregar la actividad.");
    }
}

// Función para buscar una actividad por ID
async function buscarActividad() {
    const id = document.getElementById("buscar-id").value;
    try {
        const response = await fetch(`http://172.21.0.2:32000/tasks/${id}`);
        if (!response.ok) throw new Error("Actividad no encontrada.");

        const actividad = await response.json();
        document.getElementById("resultado-nombre").innerText = `Nombre: ${actividad.Name}`;
        document.getElementById("resultado-descripcion").innerText = `Descripción: ${actividad.Content}`;
        document.getElementById("resultados-busqueda").style.display = "block";

        // Guardar el ID en un atributo de datos para usarlo en las funciones de modificar y eliminar
        document.getElementById("resultados-busqueda").setAttribute("data-id", id);
    } catch (error) {
        console.error(error);
        alert("Actividad no encontrada.");
    }
}

// Función para modificar una actividad
async function modificarActividad() {
    const id = document.getElementById("resultados-busqueda").getAttribute("data-id");
    const nombre = document.getElementById("resultado-nombre").innerText.replace("Nombre: ", "");
    const descripcion = document.getElementById("resultado-descripcion").innerText.replace("Descripción: ", "");

    try {
        const response = await fetch(`http://172.21.0.2:32000/tasks/${id}`, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ Name: nombre, Content: descripcion })
        });

        if (response.ok) {
            alert("Actividad modificada correctamente.");
            cargarActividades();
            document.getElementById("resultados-busqueda").style.display = "none";
        } else {
            throw new Error("Error al modificar la actividad.");
        }
    } catch (error) {
        console.error(error);
        alert("Error al modificar la actividad.");
    }
}

// Función para eliminar una actividad
async function eliminarActividad() {
    const id = document.getElementById("resultados-busqueda").getAttribute("data-id");
    try {
        const response = await fetch(`http://172.21.0.2:32000/tasks/${id}`, { method: "DELETE" });
        if (response.ok) {
            alert("Actividad eliminada correctamente.");
            cargarActividades();
            document.getElementById("resultados-busqueda").style.display = "none";
        } else {
            throw new Error("Error al eliminar la actividad.");
        }
    } catch (error) {
        console.error(error);
        alert("Error al eliminar la actividad.");
    }
}

// Función para guardar cambios de una actividad en los campos de búsqueda
function guardarCambios() {
    const nombre = prompt("Nuevo nombre de la actividad:", document.getElementById("resultado-nombre").innerText.replace("Nombre: ", ""));
    const descripcion = prompt("Nueva descripción de la actividad:", document.getElementById("resultado-descripcion").innerText.replace("Descripción: ", ""));

    if (nombre !== null && descripcion !== null) {
        document.getElementById("resultado-nombre").innerText = `Nombre: ${nombre}`;
        document.getElementById("resultado-descripcion").innerText = `Descripción: ${descripcion}`;
    }
}

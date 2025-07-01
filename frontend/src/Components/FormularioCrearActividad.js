import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom'
import './CreateActivity.css' 

function FormularioCrearActividad({ setiscreating }) {
    const navigation = useNavigate()
    const [profesores, setprofesores] = useState([]);
    const [categorias, setcategorias] = useState([]);

    const headers_Create = {
        'Authorization': localStorage.getItem('userToken'),
        'Content-Type': 'text/plain',
    };

    async function PostCrearActividad() {
        const data = {
            name: document.getElementById("Activity_name_create").value,
            description: document.getElementById("Activity_description_create").value,
            duration: parseInt(document.getElementById("Activity_duration_create").value),
            coachid: parseInt(document.getElementById("profesores_select").value),
            activitytypeid: parseInt(document.getElementById("categorias_select").value),
            capacity: parseInt(document.getElementById("Activity_capacity_create").value),
        }
        console.log("JSON: ",JSON.stringify(data))

        await fetch("/actividades/new",{
            method:"POST",
            headers: headers_Create,
            body: JSON.stringify(data),
        }).then((res)=>{
            if (res.status == 202) {
                alert("Actividad Creada con exito")
            }
            if (res.status == 401) {
                console.log("token is not valid")
                localStorage.removeItem('userToken')
                alert("Debes iniciar sesion nuevamente")
                navigation("/Login")
            }
        })
    }


    useEffect(() => {
        // Fetch profesores
        fetch("/users/profesores", {
            method: "GET",
            headers: headers_Create
        })
            .then((res) => {
                if (res.status == 401) {
                    console.log("token is not valid")
                    localStorage.removeItem('userToken')
                    alert("Debes iniciar sesion nuevamente")
                    navigation("/Login")
                    return
                }
                return (res.json());
            }).then((data) => {
                console.log(data);
                setprofesores(data);
            })
            .catch((error) => {
                alert("Error al realizar fetch de profesores");
                console.error(error);
            });

        // Fetch categorías
        fetch("/activities/types", {
            method: "GET",
            headers: headers_Create
        })
            .then((res) => {
                if (res.status == 401) {
                    console.log("token is not valid")
                    localStorage.removeItem('userToken')
                    alert("Debes iniciar sesion nuevamente")
                    navigation("/Login")
                    return
                }
                return (res.json());
            }).then((data) => {
                console.log(data);
                setcategorias(data);
            })
            .catch((error) => {
                alert("Error al realizar fetch de Categorias");
                console.error(error);
            });
    }, []); // Array vacío asegura que el useEffect se ejecute solo al montar el componente

    return (
        <div className='CreateActivityForm-container'>
            <form className='CreateActivityForm'>
                <h5>Formulario Crear Actividad</h5>
                <input id='Activity_name_create' placeholder='Nombre Actividad' />
                <input id='Activity_description_create' placeholder='Descripción' />
                <p>Profesores:</p>
                <select id="profesores_select">
                    {profesores?.map((Profesor) => (
                        <option key={Profesor.id} value={Profesor.id}>
                            {Profesor.name}
                        </option>
                    ))}
                </select>
                <p>Categoría:</p>
                <select id="categorias_select">
                    {categorias?.map((Categoria) => (
                        <option key={Categoria.id} value={Categoria.id}>
                            {Categoria.name}
                        </option>
                    ))}
                </select>
                <p>Cupo:</p>
                <input id='Activity_capacity_create' placeholder='Cupo Actividad' />
                <p>Duracion:</p>
                <input id='Activity_duration_create' placeholder='Duracion de la Actividad' />
                <button onClick={()=>{PostCrearActividad()}}>Crear Actividad</button>
                <button onClick={() => setiscreating(false)}>Cancelar</button>
            </form>
        </div>
    );
}

export default FormularioCrearActividad;
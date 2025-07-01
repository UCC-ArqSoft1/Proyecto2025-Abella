import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom'
import './CreateActivity.css' 
import FormularioCrearCategoria from './FormularioCrearCategoria';

function FormularioEditarActividad({ setiscreating, idactividad , name ,description,duration, capacity}) {
    const navigation = useNavigate()
    const [profesores, setprofesores] = useState([]);
    const [categorias, setcategorias] = useState([]);
    const [iscreatingCategory,setiscreatingCategory] = useState(false)
    const headers_Create = {
        'Authorization': localStorage.getItem('userToken'),
        'Content-Type': 'text/plain',
    };
    

    async function PostEditarActivdad() {
        const data = {
            id: idactividad,
            name: document.getElementById("Activity_name_edit").value,
            description: document.getElementById("Activity_description_edit").value,
            duration: parseInt(document.getElementById("Activity_duration_edit").value),
            coachid: parseInt(document.getElementById("profesores_select").value),
            activitytypeid: parseInt(document.getElementById("categorias_select").value),
            capacity: parseInt(document.getElementById("Activity_capacity_edit").value),
        }
        console.log("JSON: ",JSON.stringify(data))

        await fetch("/activities/edit",{
            method:"POST",
            headers: headers_Create,
            body: JSON.stringify(data),
        }).then((res)=>{
            if (res.status == 200) {
                alert("Actividad Editada con exito")
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
  let isUnauthorized = false; // Flag to track if a 401 has been handled

  async function fetchData() {
    try {
      // Fetch profesores
      const resProfesores = await fetch("/users/profesores", {
        method: "GET",
        headers: headers_Create,
      });

      if (resProfesores.status === 401) {
        isUnauthorized = true;
        console.log("Token is not valid");
        localStorage.removeItem('userToken');
        alert("Debes iniciar sesión nuevamente");
        navigation("/Login");
        return; // Exit early to avoid further requests
      }

      if (!resProfesores.ok) {
        throw new Error("Error al realizar fetch de profesores");
      }

      const profesoresData = await resProfesores.json();
      setprofesores(profesoresData);
      console.log(profesoresData);


      const resCategorias = await fetch("/activities/types", {
        method: "GET",
        headers: headers_Create,
      });

      if (resCategorias.status === 401) {
        if (!isUnauthorized) {
          isUnauthorized = true;
          console.log("Token is not valid");
          localStorage.removeItem('userToken');
          alert("Debes iniciar sesión nuevamente");
          navigation("/Login");
        }
        return;
      }

      if (!resCategorias.ok) {
        throw new Error("Error al realizar fetch de categorías");
      }

      const categoriasData = await resCategorias.json();
      setcategorias(categoriasData);
      console.log(categoriasData);
    } catch (error) {
      if (!isUnauthorized) {
        // Only show alert if not a 401 error
        alert(error.message || "Error al realizar fetch de datos");
        console.error(error);
      }
    }
  }

  fetchData();
}, []); // Array vacío asegura que el useEffect se ejecute solo al montar el componente

    return (
        <div className='CreateActivityForm-container'>
            <div className='CreateActivityForm'>
                <h5>Formulario Crear Actividad</h5>
                <input id='Activity_name_edit' defaultValue={name} placeholder='Nombre Actividad' />
                <input id='Activity_description_edit' defaultValue={description} placeholder='Descripción' />
                <p>Profesores:</p>
                <select id="profesores_select">
                    {profesores?.map((Profesor) => (
                        <option key={Profesor.id} value={Profesor.id}>
                            {Profesor.name}
                        </option>
                    ))}
                </select>
                <p>Categoría:</p>
                <select id="categorias_select" >
                    {categorias?.map((Categoria) => (
                        <option key={Categoria.id} value={Categoria.id}>
                            {Categoria.name}
                        </option>
                    ))}
                </select>
                <button onClick={()=>{setiscreatingCategory(true)}}>+</button>
                <p>Cupo:</p>
                <input id='Activity_capacity_edit' defaultValue={capacity} placeholder='Cupo Actividad' />
                <p>Duracion:</p>
                <input id='Activity_duration_edit' defaultValue={duration} placeholder='Duracion de la Actividad' />
                <button onClick={()=>{PostEditarActivdad()}}>Confirmar</button>
                <button onClick={() => setiscreating(false)}>Cancelar</button>
                {iscreatingCategory == true && <FormularioCrearCategoria setiscreatingCategory={setiscreatingCategory} setiscreating={setiscreating}/>}
            </div>
        </div>
    );
}

export default FormularioEditarActividad;
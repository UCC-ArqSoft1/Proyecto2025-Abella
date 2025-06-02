
import { useParams } from 'react-router-dom';

function Actividad() {
    const { id } = useParams();
    return (
        <p>Actividad</p>
    )
}



export default Actividad;
import { useState } from "react"
import { createUser } from "../../api"

const UserForm = ({ onUserAdded }) => {
    
    const [name, setName] = useState("")
    const [email, setEmail] = useState("")

    const handleSubmit = async (e) => {
        e.preventDefault()
        if (!name || !email) return
        try {
            await createUser({ name, email })
            onUserAdded()
            setName("")
            setEmail("")
        } catch (error) {
            console.error("Error al agregar usuario:", error)
        }
    }
    
    return (
        <div className="container mt-3">
            <h3>Agregar Usuario</h3>
            <form onSubmit={handleSubmit}>
                <div className="mb-2">
                    <label className="form-label">Nombre</label>
                    <input
                        type="text"
                        className="form-control"
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                        required
                    />
                </div>
                <div className="mb-2">
                    <label className="form-label">Correo</label>
                    <input
                        type="email"
                        className="form-control"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                    />
                </div>
                <button type="submit" className="btn btn-primary">Agregar</button>
            </form>
        </div>
    )
}
export default UserForm
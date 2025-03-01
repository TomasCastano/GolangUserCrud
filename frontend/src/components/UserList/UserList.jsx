import React, { useState, useEffect } from "react"
import axios from "axios"

const API_URL = "http://localhost:8080/users"

const UserList = () => {

    const [users, setUsers] = useState([])
    const [editingUser, setEditingUser] = useState(null)
    const [name, setName] = useState("")
    const [email, setEmail] = useState("")

    useEffect(() => {
        loadUsers()
    }, [])

    const loadUsers = async () => {
        const response = await axios.get(API_URL)
        setUsers(response.data)
    }

    const handleEdit = (user) => {
        setEditingUser(user.id)
        setName(user.name)
        setEmail(user.email)
    }

    const handleUpdate = async () => {
        await axios.put(`${API_URL}/${editingUser}`, { name, email })
        setEditingUser(null)
        setName("")
        setEmail("")
        loadUsers()
    }

    const handleDelete = async (id) => {
        await axios.delete(`${API_URL}/${id}`)
        loadUsers()
    }

    return (
        <div className="container mt-4">
            <h3>Lista de Usuarios</h3>
            <table className="table table-striped">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Nombre</th>
                        <th>Email</th>
                        <th>Acciones</th>
                    </tr>
                </thead>
                <tbody>
                    {Array.isArray(users) && users.map((user) => (
                        <tr key={user.id}>
                            <td>{user.id}</td>
                            <td>{user.name}</td>
                            <td>{user.email}</td>
                            <td>
                                <button className="btn btn-warning btn-sm me-2" onClick={() => handleEdit(user)}>
                                    Editar
                                </button>
                                <button className="btn btn-danger btn-sm" onClick={() => handleDelete(user.id)}>
                                    Eliminar
                                </button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
            {/* Formulario de edición */}
            {editingUser && (
                <div className="mt-4">
                    <h4>Editar Usuario</h4>
                    <input
                        type="text"
                        className="form-control mb-2"
                        placeholder="Nombre"
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                    />
                    <input
                        type="email"
                        className="form-control mb-2"
                        placeholder="Correo Electrónico"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                    />
                    <button className="btn btn-primary me-2" onClick={handleUpdate}>
                        Guardar Cambios
                    </button>
                    <button className="btn btn-secondary" onClick={() =>
                        setEditingUser(null)}>
                        Cancelar
                    </button>
                </div>
            )}
        </div>
    )
}
export default UserList
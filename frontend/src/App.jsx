import React, { useState } from "react"

import UserForm from "./components/UserForm/UserForm"
import UserList from "./components/UserList/UserList"

const App = () => {

    const [refresh, setRefresh] = useState(false)

    const handleUserAdded = () => {
        setRefresh(!refresh)
    }

    return (
        <div className="container mt-5">
            <h1 className="text-center">Administración de Usuarios</h1>
            <UserForm onUserAdded={handleUserAdded} />
            <UserList key={refresh} />
        </div>
    )
}
export default App

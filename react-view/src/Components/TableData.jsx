import { Link, useNavigate } from "react-router-dom";
import { deleteUser } from "../services/ApiServices";
import { useState } from "react";
export default function TableData({ users }) {
    const navigate = useNavigate();
    const [isDeleting, setIsDeleting] = useState(false);
    const handleDelete = async (id) => {
        try {
            console.log("code run here");
            await deleteUser(id);
            navigate("/");
            window.location.reload();
        } catch (error) {
            console.error("handleDelete", error);
        }
    }
    return (
        <div>
            {users.map((user) => (
                <div key={user.ID} className="">
                    <h3>{user.name}</h3>
                    <p>{user.email}</p>
                    <Link to={`/user/${user.ID}`}><button>View</button></Link>
                    <Link to={`/edit/${user.ID}`}><button>Edit</button></Link>
                    <button onClick={() => handleDelete(user.ID)}>Delete</button>
                </div>
            ))}
        </div>
    );
}
import { useEffect } from "react";
import { useParams } from "react-router-dom";
import { useState } from "react";
import { fetchUser } from "../services/ApiServices";
import { Link } from "react-router-dom";

export default function Profile({}) {
    const {id} = useParams();
    const [user, setUser] = useState({});
    const [error, setError] = useState(null);
    useEffect(() => {
        const getUser = async () => {
            try {
                const user = await fetchUser(id);
                setUser(user);
                console.log(user);
            }
            catch (error) {
                setError(error);
                console.log('error', error);
            }
        }
        getUser();
    }, []);
    return (
        <div>
            <h2>User Profile</h2>
            <div>
                <h3>{user.name}</h3>
                <p>{user.email}</p>
                <Link to="/"><button>Return</button></Link>
            </div>
        </div>
    );
}
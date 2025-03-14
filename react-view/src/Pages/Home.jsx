import { useEffect } from "react";
import { fetchUsers } from "../services/ApiServices";
import { useState } from "react";
import TableData from "../Components/TableData";
import { Link } from "react-router-dom";
function Home() {
  const [users, setUsers] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    const getUsers = async () => {
      try {
        const users = await fetchUsers();
        console.log("API response:", users);
        setUsers(users);
      } catch (error) {
        setError(error);
        console.log("error", error);
      }
    };
    getUsers();
  }, []);
  return (
    <div>
      <h2>List of people</h2>
      <TableData users={users} />
      <Link to="/add">
        <button>Add</button>
      </Link>
    </div>
  );
}

export default Home;

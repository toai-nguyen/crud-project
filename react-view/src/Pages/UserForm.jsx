import { useParams, useNavigate, Link } from "react-router-dom";
import { updateUser, createUser, fetchUser } from "../services/ApiServices";
import { useEffect, useState } from "react";
export default function UserForm({ isEdit = false }) {
  const { id } = useParams();
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    name: "",
    email: "",
  });
  useEffect(() => {
    const getUserData = async () => {
      if (isEdit && id) {
        try {
          const userData = await fetchUser(id);
          setFormData({
            name: userData.name || "",
            email: userData.email || "",
          });
        } catch (error) {
          console.error("getUserData", error);
        }
      }
    };
    getUserData();
  }, [isEdit, id]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };
  const handleSubmit = async (e) => {
    console.log("run here");
    e.preventDefault();
    if (isEdit) {
      try {
        await updateUser(id, formData);
        navigate("/");
      } catch (error) {
        console.error("updateUser", error);
      }
    } else {
      try {
        await createUser(formData);
        navigate("/");
      } catch (error) {
        console.error("createUser", error);
      }
    }
  };
  return (
    <div>
      <form onSubmit={handleSubmit}>
        <p>Email</p>
        <input
          type="text"
          name="email"
          value={formData.email}
          onChange={handleChange}
          required
        />
        <p>Name</p>
        <input
          type="text"
          name="name"
          value={formData.name}
          onChange={handleChange}
          required
        />
        <div style={{ marginTop: "10px" }}>
          <button type="submit">{isEdit ? "Update" : "Add"}</button>
          <Link to="/">
            <button>Return</button>
          </Link>
        </div>
      </form>
    </div>
  );
}

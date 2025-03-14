import axios from 'axios';

const api = axios.create({
    baseURL: 'http://localhost:8000',
    headers: {
        'Content-Type': 'application/json',
    },
});
//route user.index
export const fetchUsers = async () => {
    try {
        const response = await api.get('api/users');
        if (response.data && Array.isArray(response.data.users)) {
            return response.data.users;
        }
        return response.data;
    } catch (error) {
        console.error('fetchUsers', error);
        throw error;
    }
}
//route user.show
export const fetchUser = async (id) => {
    try {
        const response = await api.get(`api/user/${id}`);
        
        return response.data.user;
    } catch (error) {
        console.error('fetchUser', error);
        throw error;
    }
}
//create user
export const createUser = async (data) => {
    try {
        const response = await api.post('api/create_user', data);
        return response.data;
    } catch (error) {
        console.error('createUser', error);
        throw error;
    }
}
//update user
export const updateUser = async (id, data) => {
    try {
        const response = await api.put(`api/update_user/${id}`, data);
        return response.data;
    } catch (error) {
        console.error('updateUser', error);
        throw error;
    }
}
//delete user 
export const  deleteUser = async (id) => {
    try {
        const response = await api.delete(`api/delete_user/${id}`);
        return response.data;
    } catch (error) {
        console.error('deleteUser', error);
        throw error;
    }
}
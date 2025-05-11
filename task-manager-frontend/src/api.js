import axios from 'axios';

// Create Axios instance with baseURL and interceptor for attaching token
const instance = axios.create({
  baseURL: 'http://localhost:8080/',
});

instance.interceptors.request.use(config => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export const registerUser = (data) => instance.post('register', data);
export const loginUser = (data) => instance.post('login', data);
export const fetchTasks = async () => {
    const response = await instance.get('tasks');
    return response.data; 
  };export const createTask = (task) => instance.post('tasks', task);
export const updateTask = (id, task) => instance.put(`tasks/${id}`, task);
export const deleteTask = (id) => instance.delete(`tasks/${id}`);

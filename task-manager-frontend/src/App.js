import React, { useContext, useState, useEffect } from 'react';
import { AuthContext, AuthProvider } from './context/AuthContext';
import Login from './components/Login';
import Register from './components/Register';
import TaskForm from './components/TaskForm';
import TaskList from './components/TaskList';
import { fetchTasks, createTask, updateTask, deleteTask } from './api';

function AppContent() {
  const { user, logout } = useContext(AuthContext);
  const [tasks, setTasks] = useState([]);
  const [editingTask, setEditingTask] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [showRegister, setShowRegister] = useState(false);
  const [filter, setFilter] = useState('All');

  useEffect(() => {
    if (user) loadTasks();
  }, [user]);

  const loadTasks = async () => {
    setLoading(true);
    setError('');
    try {
      const data = await fetchTasks();
      setTasks(data);
    } catch (err) {
      setError('Failed to fetch tasks');
    } finally {
      setLoading(false);
    }
  };

  const handleAdd = async (task) => {
    setLoading(true);
    setError('');
    try {
      const response = await createTask(task);
      if (response && response.data) {
        const newTask = response.data;
        setTasks(prev => Array.isArray(prev) ? [...prev, newTask] : [newTask]);
      } else {
        setError('Failed to add task. Invalid response from server.');
      }
    } catch (err) {
      setError('Failed to add task. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  const handleUpdate = async (task) => {
    setLoading(true);
    setError('');
    try {
      const response = await updateTask(editingTask.id, task);
      if (response && response.data) {
        const updated = response.data;
        setTasks(prev => prev.map(t => t.id === updated.id ? updated : t));
        setEditingTask(null);
      } else {
        setError('Failed to update task. Invalid response from server.');
      }
    } catch {
      setError('Failed to update task');
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (id) => {
    if (!window.confirm('Delete task?')) return;
    setLoading(true);
    setError('');
    try {
      await deleteTask(id);
      setTasks(prev => prev.filter(t => t.id !== id));
    } catch {
      setError('Failed to delete task');
    } finally {
      setLoading(false);
    }
  };

  const filteredTasks = filter === 'All' ? tasks : tasks.filter(task => task.status === filter);

  if (!user) {
    return (
      <div style={{ maxWidth: 400, margin: "2rem auto", padding: "1rem" }}>
        {showRegister ? (
          <>
            <Register goToLogin={() => setShowRegister(false)} />
            <p>Already have an account? <button onClick={() => setShowRegister(false)}>Login</button></p>
          </>
        ) : (
          <>
            <Login />
            <p>Don't have an account? <button onClick={() => setShowRegister(true)}>Register</button></p>
          </>
        )}
      </div>
    );
  }

  return (
    <div style={{ maxWidth: 600, margin: "2rem auto", padding: "1rem" }}>
      <h1>Welcome, {user.username}</h1>
      <button onClick={logout} style={{ marginBottom: '1rem' }}>Logout</button>
      {error && <p style={{ color: "red" }}>{error}</p>}
      {loading && <p>Loading...</p>}

      {/* Filter dropdown */}
      <div style={{ marginBottom: '1rem' }}>
        <label>
          Filter Tasks:{' '}
          <select value={filter} onChange={e => setFilter(e.target.value)}>
            <option value="All">All</option>
            <option value="Pending">Pending</option>
            <option value="In-Progress">In-Progress</option>
            <option value="Completed">Completed</option>
          </select>
        </label>
      </div>

      <TaskForm
        onSubmit={editingTask ? handleUpdate : handleAdd}
        initialTask={editingTask}
        onCancel={() => setEditingTask(null)}
      />
      <TaskList tasks={filteredTasks} onEdit={setEditingTask} onDelete={handleDelete} />
      <hr />
    </div>
  );
}

const App = () => (
  <AuthProvider>
    <AppContent />
  </AuthProvider>
);

export default App;

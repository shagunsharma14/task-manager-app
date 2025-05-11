import React, { useState, useEffect } from 'react';
import { fetchTasks, createTask, updateTask, deleteTask } from './api';
import TaskForm from './components/TaskForm';
import TaskList from './components/TaskList';

function App() {
  const [tasks, setTasks] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [editingTask, setEditingTask] = useState(null);

  const loadTasks = async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await fetchTasks();
      setTasks(data);
    } catch {
      setError('Failed to load tasks');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadTasks();
  }, []);

  const handleAdd = async (task) => {
    setLoading(true);
    setError(null);
    try {
      const newTask = await createTask(task);
      setTasks(prev => [...prev, newTask]);
    } catch {
      setError('Failed to add task');
    } finally {
      setLoading(false);
    }
  };

  const handleUpdate = async (task) => {
    setLoading(true);
    setError(null);
    try {
      const updatedTask = await updateTask(editingTask.id, task);
      setTasks(prev => prev.map(t => (t.id === updatedTask.id ? updatedTask : t)));
      setEditingTask(null);
    } catch {
      setError('Failed to update task');
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (id) => {
    if (!window.confirm('Are you sure you want to delete this task?')) return;
    setLoading(true);
    setError(null);
    try {
      await deleteTask(id);
      setTasks(prev => prev.filter(t => t.id !== id));
    } catch {
      setError('Failed to delete task');
    } finally {
      setLoading(false);
    }
  };

  const startEdit = (task) => {
    setEditingTask(task);
  };

  const cancelEdit = () => {
    setEditingTask(null);
  };

  return (
    <div style={{ maxWidth: 600, margin: '2rem auto', padding: '0 1rem' }}>
      <h1>Task Manager</h1>
      {error && <div style={{ color: 'red', marginBottom: '1rem' }}>{error}</div>}
      {loading && <div>Loading...</div>}
      <TaskForm
        onSubmit={editingTask ? handleUpdate : handleAdd}
        initialTask={editingTask}
        onCancel={cancelEdit}
      />
      <TaskList tasks={tasks} onEdit={startEdit} onDelete={handleDelete} />
      <hr />
      <div>
        <h2>Extensibility Ideas</h2>
        <ul>
          <li>Filter tasks by status or due date.</li>
          <li>User authentication and role-based access.</li>
          <li>Task prioritization and categories.</li>
          <li>Reminders and notifications for due tasks.</li>
          <li>Drag and drop to reorder tasks.</li>
        </ul>
      </div>
    </div>
  );
}

export default App;

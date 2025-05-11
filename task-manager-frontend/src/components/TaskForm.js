import React, { useState, useEffect } from 'react';

const TaskForm = ({ onSubmit, initialTask, onCancel }) => {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [status, setStatus] = useState('Pending');
  const [dueDate, setDueDate] = useState('');
  const [error, setError] = useState(null);

  useEffect(() => {
    if (initialTask) {
      setTitle(initialTask.title || '');
      setDescription(initialTask.description || '');
      setStatus(initialTask.status || 'Pending');
      setDueDate(initialTask.due_date ? initialTask.due_date.substring(0,16) : '');
    } else {
      setTitle('');
      setDescription('');
      setStatus('Pending');
      setDueDate('');
    }
    setError(null);
  }, [initialTask]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!title.trim() || !description.trim()) {
      setError('Title and Description are required');
      return;
    }
    setError(null);
    try {
      await onSubmit({
        title: title.trim(),
        description: description.trim(),
        status,
        due_date: dueDate || null,
      });
      // Clear form if adding new task
      if (!initialTask) {
        setTitle('');
        setDescription('');
        setStatus('Pending');
        setDueDate('');
      }
    } catch {
      setError('Failed to submit task. Please try again.');
    }
  };

  return (
    <form onSubmit={handleSubmit} style={{ marginBottom: '1rem' }}>
      {error && <div style={{color: 'red'}}>{error}</div>}
      <div>
        <input
          type="text"
          placeholder="Title"
          value={title}
          onChange={e => setTitle(e.target.value)}
          required
          style={{ width: '100%', padding: '8px', marginBottom: '8px' }}
        />
      </div>
      <div>
        <textarea
          placeholder="Description"
          value={description}
          onChange={e => setDescription(e.target.value)}
          required
          style={{ width: '100%', padding: '8px', marginBottom: '8px' }}
        />
      </div>
      <div>
        <label>
          Status:{' '}
          <select value={status} onChange={e => setStatus(e.target.value)} style={{ marginBottom: '8px' }}>
            <option value="Pending">Pending</option>
            <option value="In-Progress">In-Progress</option>
            <option value="Completed">Completed</option>
          </select>
        </label>
      </div>
      <div>
        <label>
          Due Date:{' '}
          <input
            type="datetime-local"
            value={dueDate}
            onChange={e => setDueDate(e.target.value)}
            style={{ marginBottom: '8px' }}
          />
        </label>
      </div>
      <button type="submit" style={{ marginRight: '8px' }}>
        {initialTask ? 'Update Task' : 'Add Task'}
      </button>
      {initialTask && <button type="button" onClick={onCancel}>Cancel</button>}
    </form>
  );
};

export default TaskForm;

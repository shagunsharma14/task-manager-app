import React from 'react';

const TaskItem = ({ task, onEdit, onDelete }) => {
  return (
    <div style={{
      border: '1px solid #ccc',
      padding: '12px',
      marginBottom: '8px',
      borderRadius: '4px',
      backgroundColor: '#f9f9f9',
    }}>
      <h3>{task.title}</h3>
      <p>{task.description}</p>
      <p><strong>Status:</strong> {task.status}</p>
      <p><strong>Due Date:</strong> {task.due_date ? new Date(task.due_date).toLocaleString() : 'N/A'}</p>
      <button onClick={() => onEdit(task)} style={{ marginRight: '8px' }}>Edit</button>
      <button onClick={() => onDelete(task.id)} style={{ color: 'red' }}>Delete</button>
    </div>
  );
};

export default TaskItem;

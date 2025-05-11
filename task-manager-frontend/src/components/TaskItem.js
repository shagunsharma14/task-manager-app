// TaskItem.js
import React from 'react';

const TaskItem = ({ task, onEdit, onDelete }) => {
    console.log('Task in TaskItem:', task);
  return (
    <div>
      <h3>{task.title}</h3>
      <p>{task.description}</p>
      <p>Status: {task.status}</p>
      <p>Due Date: {new Date(task.due_date).toLocaleString()}</p>
      <button onClick={() => onEdit(task)}>Edit</button>
      <button onClick={() => onDelete(task.id)}>Delete</button>
    </div>
  );
};

export default TaskItem;

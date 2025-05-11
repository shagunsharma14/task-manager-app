import React from 'react';
import TaskItem from './TaskItem';

const TaskList = ({ tasks, onEdit, onDelete }) => {
    console.log('Tasks in TaskList:', tasks);
   console.log('Type of tasks:', Array.isArray(tasks)); // This should log true
    if (!tasks.length) return <p>No tasks to show.</p>;
    console.log('After Tasks in TaskList:', tasks);
    return (
      <div>
        {tasks.map(task => (
          <TaskItem key={task.id} task={task} onEdit={onEdit} onDelete={onDelete} />
        ))}
      </div>
    );
  };

export default TaskList;

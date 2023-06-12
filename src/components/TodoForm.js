import React, { useState } from 'react';
import { TodoList } from './TodoList';

export const TodoForm = () => {
  const [task, setTask] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    // .go tarafına task değerini iletmek için bir istek gönder
    fetch('http://localhost:5050/addTask', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ task }),
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error('An error occurred while adding the task.');
        }
        // İstek başarılı olduğunda istenilen işlemi yapabilirsiniz
        // ...
        setTask('');
        window.location.reload();

      })
      .catch((error) => {
        console.error(error);
        window.alert('An error occurred while adding the task.');
      });
    // İşlem tamamlandıktan sonra input alanını temizleyin
    
  };

  return (
    <div className='TodoWrapper'>
      <h3>Todo List</h3>
      <form onSubmit={handleSubmit} className='TodoForm'>
        <input
          type="text"
          placeholder="Enter task"
          value={task}
          onChange={(e) => setTask(e.target.value)}
          className='todo-input'
        />
        <button type="submit" className='todo-btn'>Add Task</button>
      </form>
      <TodoList/>
    </div>
  );
};
import React, { useState, useEffect } from 'react';

export const TodoList = () => {
  const [todos, setTodos] = useState([]);

  // .go'dan verileri almak için bir istek gönder
  useEffect(() => {
    fetch('http://localhost:5050/getAllTasks')
      .then((response) => response.json())
      .then((data) => {
        setTodos(data);
      })
      .catch((error) => {
        console.error(error);
      });
  }, []);

  const handleDelete = (id) => {
    // .go tarafına silme isteği gönder
    fetch(`http://localhost:5050/deleteTask/${id}`, {
      method: 'DELETE',
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error('An error occurred while deleting the task.');
        }
        // İstek başarılı olduğunda istenilen işlemi yapabilirsiniz
        // ...
      })
      .catch((error) => {
        console.error(error);
        window.alert('An error occurred while deleting the task.');
      });

    // Silinen görevi listeden kaldır
    setTodos(todos.filter((todo) => todo.id !== id));
  };

  // Eğer todos null ise, burada kontrol yapabilirsiniz
  if (todos === null) {
    return <p className='lbl-notask'>No task yet...</p>;
  }

  return (
    <div className='TodoWrapper'>
     <table>
        <tr >
            <td></td>
            <td></td>
            <td className='todo-place'>
                {todos.length > 0 ? (
                  todos.map((todo) => (
                    <label key={todo.id}>
                      <label className='lbl-input'>{todo.task} </label> 
                      <button className='btn-delete' onClick={() => handleDelete(todo.id)}>Delete</button>
                      <br/>
                      <br/>
                    </label>
                  ))
                ): 
                <p>Loading...</p>
                }
            </td>
        </tr>
    </table> 
             
    </div>
  );
};

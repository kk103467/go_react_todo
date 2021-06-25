import React, { useState, useEffect } from 'react';
import axios from 'axios';

import Todo from '../components/Todo';
import TodoForm from '../components/TodoForm'

export default function ViewPage() {
  const [todos, setTodos] = useState([]);
  const endpoint = "http://localhost:8000/api";

  useEffect(() => {
    axios.get(endpoint + "/todos")
      .then((response) => {
        setTodos(response.data);
      })
      .catch((error) => {
        console.error(`Error: ${error}`);
      })
  }, []);

  const addTodo = (textjson) => {
    const headers = {
      'Content-Type': 'application/json'
    }
    // add new todo by the POST request to the server
    axios.post(endpoint + "/todo", textjson, {headers: headers})
    .then((response) => {
      // reflect the updated todos based on the response from the server
      setTodos(response.data);
    })
  };

  const changeComplete = (id) => {
    const prevTodos = [...todos];
    const newTodos = prevTodos.map((todo) => {
      if (todo.id === id) {
        todo.isCompleted = !todo.isCompleted
      }
      return todo
    });
    setTodos(newTodos)
  };

  const deleteTodo = (id) => {
    const prevTodos = [...todos];
    const newTodos = prevTodos.filter((todo, index) => {
      return todo.id !== id;
    })
    setTodos(newTodos);
  };

  const changeText = (id, text) => {
    const newTodos = [...todos];
    newTodos[id].text = text;
    setTodos(newTodos);
  };

  return (
    <div className="App">
      <div className="todo-list">
        {todos.map((todo) => (
          <Todo key={todo.id} todo={todo} changeComplete={changeComplete} changeText={changeText} deleteTodo={deleteTodo} />
        ))}
        <TodoForm todos={todos} addTodo={addTodo} />
      </div>
    </div>
  );
}
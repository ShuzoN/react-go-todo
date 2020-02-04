import * as React from 'react';
import './App.css';
import { Tasks, Todo } from './View/Tasks';

const App = () => {
  const todos: Todo[] = [{ title: 'hoge' }, { title: 'fuga' }];

  return (
    <div className="App">
      <Tasks tasks={{ todos: todos }} />
    </div>
  );
}

export default App;

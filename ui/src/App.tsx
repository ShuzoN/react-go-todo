import './App.css';
import { Tasks, Todo } from './View/Tasks';
import { Switch, Route, BrowserRouter as Router } from 'react-router-dom';
import React, { useState } from 'react';

const App = () => {
  const [todos, setTodos] = useState<Todo[]>([{ title: 'hoge' }, { title: 'fuga' }]);

  const onChange = (todos: Todo[]) => {
    setTodos(todos);
  };


  return (
    <div className="App">
      <Router>
        <Switch>
          <Route path="/">
            <Tasks todos={todos} onChange={onChange} />
          </Route>
        </Switch>
      </Router>
    </div>
  );
}

export default App;

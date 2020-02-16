import './App.css';
import { Tasks } from './View/Tasks';
import { Switch, Route, BrowserRouter as Router } from 'react-router-dom';
import React, { useState, useCallback, useEffect } from 'react';
import { TodoEdit } from './View/TodoEdit';
import { Header } from './View/Header';
import { UserGateway, UserGatewayImpl } from './Gateway/UserGateway';
import { GatewayImpl } from './Gateway/GatewaImpl';
import { Gateway } from './Gateway/Gateway';
import { TodoGateway, TodoGatewayImpl } from './Gateway/TodoGateway';
import { fetchTodos } from './Epic';
import { Todo } from './Contract';


interface Gateways {
  userGateway: UserGateway
  todoGateway: TodoGateway
}

const gateway: Gateway = new GatewayImpl();

export const gateways: Readonly<Gateways> = {
  userGateway: new UserGatewayImpl(gateway),
  todoGateway: new TodoGatewayImpl(gateway),
}

const App = (): JSX.Element => {
  const [todos, setTodos] = useState<Todo[]>([]);

  useEffect(() => {
    ; (async () => {
      const maybePromiseTodos: Todo[] | Error = await fetchTodos(gateways.todoGateway);
      const promiseTodos: Todo[] = maybePromiseTodos instanceof Error ? [] : maybePromiseTodos;

      setTodos(promiseTodos)
    })()
  }, []);


  const onChange = useCallback((todo: Todo) => {
    const index = todo.id <= 0 ? 0 : todo.id - 1;
    const target = [...todos];
    target.splice(index, 1, todo);
    setTodos(target);
  }, [todos]);


  return (
    <div className="App">
      <Header />
      <Router>
        <Switch>
          <Route exact path={`/:id`}>
            <TodoEdit todos={todos} onChange={onChange} />
          </Route>
          <Route exact path="/">
            <Tasks todos={todos} onChange={onChange} />
          </Route>
        </Switch>
      </Router>
    </div>
  );
}

export default App;

import './App.css';
import { Tasks, Todo } from './View/Tasks';
import { Switch, Route, BrowserRouter as Router } from 'react-router-dom';
import React, { useState, useCallback } from 'react';
import { TodoEdit } from './View/TodoEdit';
import { Header } from './View/Header';
import moment from 'moment';
import { UserGateway, UserGatewayImpl } from './Gateway/UserGateway';
import { GatewayImpl } from './Gateway/GatewaImpl';
import { Gateway } from './Gateway/Gateway';


interface Gateways {
  userGateway: UserGateway
}

const gateway: Gateway = new GatewayImpl();

export const gateways: Readonly<Gateways> = {
  userGateway: new UserGatewayImpl(gateway),
}

const App = (): JSX.Element => {
  const [todos, setTodos] = useState<Todo[]>([
    { id: 1, title: 'hoge', deadLine: moment('20200101') },
    { id: 2, title: 'fuga', deadLine: moment('20200108') }
  ]);

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

import React from 'react';
import './App.css';
import { List, ListItem, ListItemText, ListItemIcon } from '@material-ui/core';
import ArrowForwardIosOutlinedIcon from '@material-ui/icons/ArrowForwardIosOutlined';

const TodoItemButton = (props: { todoTitle: string }): JSX.Element => {
  return (
    <ListItem button>
      <ListItemIcon>
        <ArrowForwardIosOutlinedIcon />
      </ListItemIcon>
      <ListItemText>{props.todoTitle}</ListItemText>
    </ListItem>
  );
}

const App = () => {
  return (
    <div className="App">
      <List>
        <TodoItemButton todoTitle={'hoge'} />
        <TodoItemButton todoTitle={'fuga'} />
      </List>
    </div>
  );
}

export default App;

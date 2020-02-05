import React from 'react';
import { List, ListItem, ListItemText, ListItemIcon } from '@material-ui/core';
import ArrowRightIcon from '@material-ui/icons/ArrowRight';
import { Link } from 'react-router-dom';

const TodoItemButton = (props: {
    todoTitle: string
}): JSX.Element => {
    return (
        <ListItem button>
            <ListItemIcon>
                <ArrowRightIcon />
            </ListItemIcon>
            <ListItemText>{props.todoTitle}</ListItemText>
        </ListItem>
    );
}

export interface Todo {
    id: number
    title: string
}

export interface TaskProps {
    todos: Todo[]
    onChange: (todo: Todo) => void
}

export const Tasks = (tasks: TaskProps): JSX.Element => {
    const todoItems = tasks.todos.map((todo, index) => {
        return <TodoItemButton key={index} todoTitle={todo.title} />
    })

    return (
        <>
            <Link to="/1">hoge</Link>
            <List>
                {todoItems}
            </List>
        </>
    );
}
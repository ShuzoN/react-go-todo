import React from 'react';
import { List, ListItem, ListItemText, ListItemIcon } from '@material-ui/core';
import ArrowRightIcon from '@material-ui/icons/ArrowRight';

const TodoItemButton = (props: { todoTitle: string }): JSX.Element => {
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
    title: string
}

export interface TaskProps {
    todos: Todo[]
}

export const Tasks = (props: { tasks: TaskProps }): JSX.Element => {
    const todoItems = props.tasks.todos.map((todo) => {
        return <TodoItemButton todoTitle={todo.title} />
    })

    return (
        <List>
            {todoItems}
        </List>
    );
}
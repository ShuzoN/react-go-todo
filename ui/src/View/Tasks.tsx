import React from 'react';
import { List, ListItem, ListItemText, ListItemIcon } from '@material-ui/core';
import ArrowRightIcon from '@material-ui/icons/ArrowRight';
import { useHistory } from 'react-router-dom';
import { Moment } from 'moment';

const TodoItemButton = (props: {
    todo: Todo
}): JSX.Element => {
    const history = useHistory();

    return (
        <ListItem button onClick={() => history.push('/' + props.todo.id.toString())}>
            <ListItemIcon>
                <ArrowRightIcon />
            </ListItemIcon>
            <ListItemText>{props.todo.title}</ListItemText>
        </ListItem>
    );
}

export interface Todo {
    id: number
    title: string
    deadLine: Moment | null
}

export interface TaskProps {
    todos: Todo[]
    onChange: (todo: Todo) => void
}

export const Tasks = (tasks: TaskProps): JSX.Element => {
    const todoItems = tasks.todos.map((todo: Todo) => {
        return <TodoItemButton key={todo.id} todo={todo} />
    })

    return (
        <>
            <List>
                {todoItems}
            </List>
        </>
    );
}
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

export const Tasks = (): JSX.Element => {
    return (
        <List>
            <TodoItemButton todoTitle={'hoge'} />
            <TodoItemButton todoTitle={'fuga'} />
        </List>
    );
}
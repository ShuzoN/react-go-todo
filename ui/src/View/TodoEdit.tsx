import React from 'react';
import { useParams } from 'react-router-dom';
import { FormControl, Typography, InputLabel, OutlinedInput, List, ListItem, TextField, } from '@material-ui/core';
import { Todo } from './Tasks';


export const TodoEdit = (props: {
    todos: Todo[],
    onChange: (todo: Todo) => void,
}): JSX.Element => {
    const params: { id?: string | undefined } = useParams();

    if (params.id === undefined) { return <></>; }

    const editTodo = props.todos.find(
        (todo: Todo) => todo.id === Number(params.id)
    );

    if (editTodo === undefined) { return <></>; }

    return (
        <List>
            <ListItem>
                <Typography>id : {editTodo.id}</Typography>
            </ListItem>
            <ListItem>
                <FormControl>
                    <InputLabel htmlFor="component-outlined">title</InputLabel>
                    <OutlinedInput
                        value={editTodo.title}
                        onChange={(e) => {
                            props.onChange({ ...editTodo, title: e.target.value })
                        }} />
                </FormControl>
            </ListItem>
            <ListItem>
                <TextField
                    id="date"
                    label="Birthday"
                    type="date"
                    defaultValue="2020-01-01"
                    InputLabelProps={{
                        shrink: true,
                    }}
                />
            </ListItem>
        </List>
    );
};
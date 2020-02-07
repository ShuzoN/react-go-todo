import React from 'react';
import { useParams } from 'react-router-dom';
import { FormControl, InputLabel, OutlinedInput, List, ListItem, TextField, Button, Card } from '@material-ui/core';
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
        <Card>
            <List>
                <ListItem>
                    <FormControl>
                        <InputLabel htmlFor="component-outlined">title</InputLabel>
                        <OutlinedInput
                            fullWidth
                            value={editTodo.title}
                            onChange={(e) => {
                                props.onChange({ ...editTodo, title: e.target.value })
                            }} />
                    </FormControl>
                </ListItem>
                <ListItem>
                    <TextField
                        id="date"
                        label="dead line"
                        type="date"
                        defaultValue="2020-01-01"
                        InputLabelProps={{
                            shrink: true,
                        }}
                    />
                </ListItem>
                <Button variant="contained" color="primary">
                    submit
            </Button>
            </List>
        </Card >
    );
};
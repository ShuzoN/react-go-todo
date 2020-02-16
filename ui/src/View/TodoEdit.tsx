import React, { useState } from 'react';
import { useParams } from 'react-router-dom';
import { FormControl, InputLabel, OutlinedInput, List, ListItem, Button, makeStyles, Grid, Card } from '@material-ui/core';
import { TodoDatePickerForm } from './TodoDatePickerForm';
import { Todo } from '../Contract';

const useStyles = makeStyles({
    card: {
        margin: '2% 10%',
    },
    form: {
        width: '100%',
    },
});

export const TodoEdit = (props: {
    todos: Todo[],
    onChange: (todo: Todo) => void,
}): JSX.Element => {
    const params: { id?: string | undefined } = useParams();

    if (params.id === undefined) { return <></>; }

    const maybeEditTodo = props.todos.find(
        (todo: Todo) => todo.id === Number(params.id)
    );

    if (maybeEditTodo === undefined) { return <></>; }


    return <TodoEditView
        todo={maybeEditTodo}
        onChange={props.onChange}
    />


};

export const TodoEditView = (props: {
    todo: Todo,
    onChange: (todo: Todo) => void,
}): JSX.Element => {
    const c = useStyles();
    const [editTodo, setEditTodo] = useState<Todo>({ ...props.todo });

    return (
        <Card className={c.card}>
            <List >
                <ListItem>
                    <Grid container spacing={3}>
                        <Grid item xs={12} >
                            <TitleForm
                                editTodo={editTodo}
                                onChange={setEditTodo}
                            />
                        </Grid>
                        <Grid item xs={12}>
                            <TodoDatePickerForm todo={editTodo} onChange={setEditTodo} />
                        </Grid>
                        <Grid item xs={12}>
                            <Grid justify="center" container>
                                <Button
                                    variant="contained"
                                    color="primary"
                                    onClick={() => props.onChange(editTodo)}
                                >
                                    submit
                            </Button>
                            </Grid>
                        </Grid>
                    </Grid>
                </ListItem >
            </List >
        </Card>
    );
};

const TitleForm = (props: {
    editTodo: Todo,
    onChange: (todo: Todo) => void,
}): JSX.Element => {
    const c = useStyles();

    return (
        <>
            <FormControl className={c.form}>
                <InputLabel
                    htmlFor="component-outlined"
                    shrink
                >
                    title
                </InputLabel>
            </FormControl>
            <FormControl className={c.form}>
                <OutlinedInput
                    fullWidth
                    value={props.editTodo.title}
                    onChange={(e) => {
                        props.onChange({ ...props.editTodo, title: e.target.value })
                    }} />
            </FormControl>
        </>
    );
};
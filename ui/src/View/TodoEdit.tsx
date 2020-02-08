import React from 'react';
import { useParams } from 'react-router-dom';
import { FormControl, InputLabel, OutlinedInput, List, ListItem, Button, makeStyles, Grid, Card } from '@material-ui/core';
import { Todo } from './Tasks';
import { DatePickerForm } from './DatePickerForm';

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
    const c = useStyles();

    const params: { id?: string | undefined } = useParams();

    if (params.id === undefined) { return <></>; }

    const editTodo = props.todos.find(
        (todo: Todo) => todo.id === Number(params.id)
    );

    if (editTodo === undefined) { return <></>; }

    return (
        <Card className={c.card}>
            <List >
                <ListItem>
                    <Grid container spacing={3}>
                        <Grid item xs={12} >
                            <TitleForm
                                editTodo={editTodo}
                                onChange={props.onChange}
                            />
                        </Grid>
                        <Grid item xs={12}>
                            <DatePickerForm />
                        </Grid>
                        <Grid item xs={12}>
                            <Grid justify="center" container>
                                <Button
                                    variant="contained"
                                    color="primary"
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
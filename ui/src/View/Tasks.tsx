import React, { useState } from 'react';
import { List, ListItem, Card, makeStyles, CardHeader, IconButton, Grid, Checkbox } from '@material-ui/core';
import EditIcon from '@material-ui/icons/Edit';
import { useHistory } from 'react-router-dom';
import { Todo } from '../Contract';
import moment from 'moment';
import { updateTodo } from '../Epic';
import { gateways } from '../App';

const useStyles = makeStyles({
    root: {
        width: '100%',
        fontSize: 14,
    },
});

const TodoItemButton = (props: {
    todo: Todo
}): JSX.Element => {
    const history = useHistory();
    const c = useStyles();

    const [checked, setChecked] = useState(props.todo.checked);

    return (
        <ListItem>
            <Card className={c.root}>
                <CardHeader
                    avatar={
                        <Checkbox
                            color="primary"
                            checked={checked}
                            onChange={() => {
                                setChecked(!checked)
                                updateTodo(gateways.todoGateway, { ...props.todo, checked: !checked })
                            }}
                        />
                    }
                    action={
                        <IconButton
                            onClick={() => history.push('/' + props.todo.id.toString())}
                        >
                            <EditIcon />
                        </IconButton>
                    }
                    title={props.todo.title}
                    subheader={"deadline: " + moment(props.todo.deadline?.toString()).format("YYYY-MM-DD")}
                />
            </Card>
        </ListItem>
    );
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
            <Grid container justify="center" spacing={2}>
                <Grid item xs={11}>
                    <List>
                        {todoItems}
                    </List>
                </Grid>
            </Grid>
        </>
    );
}
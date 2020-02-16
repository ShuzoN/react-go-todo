import React, { useState } from 'react';
import { List, ListItem, Card, makeStyles, CardHeader, IconButton, Grid } from '@material-ui/core';
import CheckBoxOutlineBlankIcon from '@material-ui/icons/CheckBoxOutlineBlank';
import CheckBoxIcon from '@material-ui/icons/CheckBox';
import EditIcon from '@material-ui/icons/Edit';
import { useHistory } from 'react-router-dom';
import { Todo } from '../Contract';
import moment from 'moment';

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

    const [check, setCheck] = useState(false);

    return (
        <ListItem>
            <Card className={c.root}>
                <CardHeader
                    avatar={
                        <IconButton
                            onClick={() => setCheck(!check)}
                        >
                            {check ? <CheckBoxIcon /> : <CheckBoxOutlineBlankIcon />}
                        </IconButton>
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
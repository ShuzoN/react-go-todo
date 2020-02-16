import React from 'react';
import { AppBar, Toolbar, IconButton, Button, } from '@material-ui/core';
import NoteAddIcon from '@material-ui/icons/NoteAdd';

export const Header = (): JSX.Element => {
    return (
        <AppBar position="static">
            <Toolbar>
                <IconButton
                    edge="start"
                    color="inherit"
                    href="/new"
                >
                    <NoteAddIcon />
                </IconButton>
                <Button size="large" variant="text" color="inherit" href="/">
                    Todos
                </Button>
            </Toolbar>
        </AppBar >
    );
}


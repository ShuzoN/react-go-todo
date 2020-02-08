import React from 'react';
import { AppBar, Typography, Toolbar, IconButton, } from '@material-ui/core';
import NoteAddIcon from '@material-ui/icons/NoteAdd';

export const Header = (): JSX.Element => {
    return (
        <AppBar position="static">
            <Toolbar>
                <IconButton edge="start" color="inherit">
                    <NoteAddIcon />
                </IconButton>
                <Typography variant="h6">
                    Todos
                        </Typography>
            </Toolbar>
        </AppBar >
    );
}


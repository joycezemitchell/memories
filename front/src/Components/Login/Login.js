import React, {useEffect, useState } from 'react';

import clsx from 'clsx';
import { makeStyles, useTheme } from '@material-ui/core/styles';

import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import Container from '@material-ui/core/Container';

import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import Avatar from '@material-ui/core/Avatar';
import Lock from '@material-ui/icons/Lock';
import { green } from '@material-ui/core/colors';

const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
        '& .MuiTextField-root': {
            margin: theme.spacing(1),
            width: '95%',
        },
    },
    margin: {
        margin: theme.spacing(1),
        width: '95%',
    },
    green: {
        color: '#fff',
        backgroundColor: green[500],
        margin:"auto",
        marginBottom: theme.spacing(1),
    },
    loginlabel:{
        margin:"auto",
        textAlign:"center",
        marginBottom: theme.spacing(1),
        marginTop: theme.spacing(3),
    },
}));

export default function Login(props) {
    const classes = useStyles();
    const theme = useTheme();

    return (
        <React.Fragment>
          <Grid container className={classes.root} justify="center">
            <Grid item lg={4} xs={12} sm={6} >
                <div className={classes.loginlabel} >
                    <Avatar className={classes.green} >
                        <Lock />
                    </Avatar>
                    <Typography variant="h5"  component="h1" >
                        Sign In
                    </Typography>
                </div>
                <form className={classes.root} noValidate autoComplete="off">
                    <TextField id="outlined-search" label="Username" type="search" variant="outlined" />
                    <TextField id="outlined-search" label="Password" type="password" variant="outlined" />
                    <Button variant="contained" size="large" color="primary" className={classes.margin}>
                        Login
                    </Button>
                </form>
            </Grid>
          </Grid>
        </React.Fragment>
    );

}    
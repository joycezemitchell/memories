import React, {useEffect, useState } from 'react';
import { makeStyles, useTheme } from '@material-ui/core/styles';
import Loginform from './Loginform'
import Grid from '@material-ui/core/Grid';

const useStyles = makeStyles((theme) => ({
    root: {

    },
}));

export default function Login(props) {
    const classes = useStyles();
    const theme = useTheme();

    return (
        <React.Fragment>
          <Grid container className={classes.root} justify="center">
            <Grid item lg={12} >
                <Loginform />
            </Grid>
          </Grid>
        </React.Fragment>
    );

}    
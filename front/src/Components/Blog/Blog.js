import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';

const useStyles = makeStyles((theme) => ({
    root: {

    },
}));

export default function Blog(props) {
    const classes = useStyles();

    return (
        <React.Fragment>
          <Grid container className={classes.root} justify="center">
            <Grid item lg={12} >
                Blog
            </Grid>
          </Grid>
        </React.Fragment>
    );

}    
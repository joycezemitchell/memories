import React, {useEffect, useState } from 'react';
import Player from './Player'
import Playlist from './Playlist'

import clsx from 'clsx';
import { makeStyles, useTheme } from '@material-ui/core/styles';

import CssBaseline from '@material-ui/core/CssBaseline';

import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import Container from '@material-ui/core/Container';

const useStyles = makeStyles((theme) => ({
  paper: {
    padding: theme.spacing(2),
    textAlign: 'center',
    color: theme.palette.text.secondary,
  },
  player: {
    [theme.breakpoints.down('lg')]: {
      position: "relative",
      left:"68%",
    }, 
    [theme.breakpoints.down('sm')]: {
      marginTop: "59%",
      position: "inherit",
      left:"inherit",
    }, 
  },
  mainplayer: {
    [theme.breakpoints.down('lg')]: {
      position: "fixed",
      width:"100%",
    }, 
    [theme.breakpoints.down('sm')]: {
      width:"inherit",
      position:"inherit"
    }, 
  },

}));

export default function Videoplay(props) {
  const classes = useStyles();
  const theme = useTheme();
 
  const [CurrentVideo, setCurrentVideo] = useState(props.id)

  const handleVideoPlay = (id) => {
    setCurrentVideo(id)
  };

  return (
    <React.Fragment>
      <Grid container >
        <Grid item lg={8}  xs={12} className={classes.mainplayer}>
          <Player video={CurrentVideo}  />
        </Grid>
        <Grid item lg={4} xs={12} className={classes.player} >
          <Playlist onVideoPlayMain={handleVideoPlay} />
        </Grid>
      </Grid>
    </React.Fragment>
  );

}
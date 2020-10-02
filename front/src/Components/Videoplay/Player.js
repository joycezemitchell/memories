import React from 'react';
import clsx from 'clsx';
import { makeStyles, useTheme } from '@material-ui/core/styles';
import ReactPlayer from 'react-player'
import CssBaseline from '@material-ui/core/CssBaseline';

let playVideo = ""

const useStyles = makeStyles((theme) => ({
  player: {
    width: '100% !important',
    height: 'auto !important',
    [theme.breakpoints.down('sm')]: {
      position: "fixed",
    }, 
  },
}));

export default function Player(props) {
  const classes = useStyles();
  const theme = useTheme();

  playVideo = process.env.REACT_APP_DIGITAL_SPACE_URL + props.video + "/1080p.m3u8"
  console.log(playVideo)

  return (
    <React.Fragment>
      <ReactPlayer  className={classes.player} controls={1} url={playVideo} autoplay={1} />
    </React.Fragment>
  );

}
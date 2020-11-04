import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import ReactPlayer from 'react-player'

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

  playVideo = process.env.REACT_APP_DIGITAL_SPACE_URL + props.video + "/1080p.m3u8"
  console.log(playVideo)

  return (
    <React.Fragment>
      <ReactPlayer  className={classes.player} controls={true} url={playVideo} autoPlay={true} />
    </React.Fragment>
  );

}
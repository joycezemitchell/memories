import React from 'react';
import clsx from 'clsx';
import { makeStyles, useTheme } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';

import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import IconButton from '@material-ui/core/IconButton';
import Typography from '@material-ui/core/Typography';
import SkipPreviousIcon from '@material-ui/icons/SkipPrevious';
import PlayArrowIcon from '@material-ui/icons/PlayArrow';
import SkipNextIcon from '@material-ui/icons/SkipNext';

import ReactTimeAgo from 'react-time-ago'

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
    boxShadow: 'none',
    backgroundColor: 'transparent',
    marginBottom: 10,
    paddingLeft:"1.5rem",
  },
  details: {
    display: 'flex',
    flexDirection: 'column',
  },
  content: {
    flex: '1 0 auto',
  },
  cover: {
    width: 151,
    height: 'auto',
  },
  videoLength: {
    color: "#777"
  },
}));

export default function Nextvideos(props) {
  const classes = useStyles();
  const theme = useTheme();
  
  const playVideo = (id, e) => {
    props.onVideoPlay(id);
  };

  return (
    <React.Fragment>
      <Card className={classes.root}>
        <CardMedia
          className={classes.cover}
          image={props.img}
          title=""
          onClick={e => playVideo(props.ids, e)}
        />
        <div className={classes.details}>
          <CardContent className={classes.content}>
            <Typography variant="h6" >
              {props.title}
            </Typography>
            <Typography variant="subtitle1">
              <ReactTimeAgo date={props.date}/>
            </Typography>
            <Typography variant="subtitle2" className={classes.videoLength} >
              {props.date}
            </Typography>
          </CardContent>
        </div>
      </Card>
    </React.Fragment>
  );

}
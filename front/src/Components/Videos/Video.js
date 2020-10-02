import React from 'react';

import clsx from 'clsx';
import { makeStyles, useTheme } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';

import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';
import CardHeader from '@material-ui/core/CardHeader';
import CardMedia from '@material-ui/core/CardMedia';
import CardContent from '@material-ui/core/CardContent';
import CardActions from '@material-ui/core/CardActions';
import Collapse from '@material-ui/core/Collapse';
import Avatar from '@material-ui/core/Avatar';
import IconButton from '@material-ui/core/IconButton';
import Typography from '@material-ui/core/Typography';
import { red } from '@material-ui/core/colors';
import FavoriteIcon from '@material-ui/icons/Favorite';
import ShareIcon from '@material-ui/icons/Share';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import MoreVertIcon from '@material-ui/icons/MoreVert';
import { Route, Link } from "react-router-dom";
import ReactTimeAgo from 'react-time-ago'

const useStyles = makeStyles((theme) => ({
  root:{
    [theme.breakpoints.down('sm')]: {
      boxShadow:"none",
    },
  },
  img: {
    width: '100%',
    height: '15rem',
    objectFit: 'cover',
  },
  media: {
    height: 0,
    paddingTop: '56.25%', // 16:9
  },
  expand: {
    transform: 'rotate(0deg)',
    marginLeft: 'auto',
    transition: theme.transitions.create('transform', {
      duration: theme.transitions.duration.shortest,
    }),
  },
  expandOpen: {
    transform: 'rotate(180deg)',
  },
  avatar: {
    backgroundColor: red[500],
  },
  itemGrid:{
    padding:".8rem",
    [theme.breakpoints.down('sm')]: {
      padding: "0",
      paddingBottom:"1.5rem",
    }, 
  },
}));


export default function Video(props) {

  const classes = useStyles();
  const [expanded, setExpanded] = React.useState(false);

  const theme = useTheme();

  const handleExpandClick = () => {
    setExpanded(!expanded);
  };

  const playVideo = (id, e) => {
    props.onVideoPlay(id);
    props.route.push('/videoplay');
  };

  return (
    <Grid item lg={3} xs={12} md={4} className={classes.itemGrid}>
      <Card className={classes.root}>
        <CardHeader
          avatar={
            <Avatar aria-label="recipe" className={classes.avatar}>
              M
            </Avatar>
          }

          title={<ReactTimeAgo date={props.date}/>}
          subheader={props.date}
        />
        <CardMedia
          className={classes.media}
          image={props.src}
          onClick={e => playVideo(props.id, e)}
        />
        <CardContent>
       
        </CardContent>
      </Card>
    </Grid>
  );

}
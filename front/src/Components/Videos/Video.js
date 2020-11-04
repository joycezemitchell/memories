import React, {useContext} from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';
import CardHeader from '@material-ui/core/CardHeader';
import CardMedia from '@material-ui/core/CardMedia';
import Avatar from '@material-ui/core/Avatar';
import { red } from '@material-ui/core/colors';
import ReactTimeAgo from 'react-time-ago'
import Actions from './Actions';
import InputContext from '../../Utilities/InputContext/InputContext'

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

  const gd = useContext(InputContext)
  const classes = useStyles();

  const playVideo = (id, e) => {
    props.onVideoPlay(id);
    console.log(id)
  };

  return (
    <Grid item lg={3} xs={12} md={4} className={classes.itemGrid}>
      <Card className={classes.root}>
        <CardHeader avatar={<Avatar aria-label="recipe" className={classes.avatar}>M</Avatar>} title={<ReactTimeAgo date={props.datex}/>} subheader={props.datex} />
        <CardMedia className={classes.media} image={props.src} onClick={e => playVideo(props.id, e)} />
        <Actions userlevel={gd.User.Level}  />      
      </Card>
    </Grid>
  );

}
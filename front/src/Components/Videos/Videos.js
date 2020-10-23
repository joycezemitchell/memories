import React, { useState, useEffect } from 'react';
import clsx from 'clsx';
import { makeStyles, useTheme } from '@material-ui/core/styles';
import Video from './Video'
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import CssBaseline from '@material-ui/core/CssBaseline';
import NdjsonParser from 'ndjson-parse';
import Cookies from 'universal-cookie';

const useStyles = makeStyles((theme) => ({
  paper: {
    padding: theme.spacing(0),
    textAlign: 'center',
    color: theme.palette.text.secondary,
  },
  img: {
    width: '100%',
    height: '15rem',
    objectFit: 'cover',
  },
  
}));

let vsi = [];

/* Token Cookie */
const cookies = new Cookies();

export default function Videos(props) {

  const classes = useStyles();
  const theme = useTheme();

  const handleVideoPlay = (id) => {
    props.onVideoPlayMain(id);
  };
  
  const [Videos, setVideos] = useState("")

  /*const tileData = [
    {
      img: 'https://material-ui.com/static/images/grid-list/vegetables.jpg',
      title: 'Image',
      author: 'author',
      cols: 2,
    },
    {
      img: 'https://material-ui.com/static/images/grid-list/honey.jpg',
      title: 'Image',
      author: 'author',
      cols: 2,
    },
    {
      img: 'https://material-ui.com/static/images/grid-list/hats.jpg',
      title: 'Image',
      author: 'author',
      cols: 2,
    },
  ];*/
  /*setVideos(tileData.map((tile) => (
    <Video src={tile.img} route={props.history} onVideoPlay={handleVideoPlay} />
  )))*/

  useEffect(() => {    
    // Not your usual API fetch. This one is parsing ReadableStream instead of ordinary JSON
    fetch(process.env.REACT_APP_API_URL + "api/v1/videos",{
      headers: {
        Authorization: `Bearer ` + cookies.get('token'),
      },
    })
    .then((response) => {
      return response.body; //ndjsonStream parses the response.body
    }).then(rs => {
      const reader = rs.getReader();
      return new ReadableStream({
        async start(controller) {
          while (true) {
            const { done, value } = await reader.read();

            // When no more data needs to be consumed, break the reading
            if (done) {
              
              let thumb = ""

              setVideos(vsi.map((v) => {
                thumb = process.env.REACT_APP_DIGITAL_SPACE_URL + v.result.memories.video + "/" + v.result.memories.thumbnail
                
                return <Video id={v.result.memories.video} title="" date={v.result.memories.created_at} src={thumb} route={props.history} onVideoPlay={handleVideoPlay}  />
              }))

              break;
            }

            // Decode the readble data - This is stupid. Streaming use to be working using ndjsonStream  
            var decoder = new TextDecoder();
            var data = decoder.decode(value);
            let f = data.replace( /[\r\n]+/gm, "" ).replace( /}{/gm, '}\n{');
            let parsedNdjson = NdjsonParser(f)
            vsi = parsedNdjson

          }
  
          // Close the stream
          controller.close();
          reader.releaseLock();
        }
      })
    })

  }, [])  

  return (
    <React.Fragment>
      <Grid container>
        {Videos}
      </Grid>
    </React.Fragment>
  );
}
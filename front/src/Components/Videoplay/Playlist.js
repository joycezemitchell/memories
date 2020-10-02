import React, { lazy, useEffect, useState } from 'react';
import clsx from 'clsx';
import { makeStyles, useTheme } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Nextvidoes from './Nextvideos'
import Loader from './Loader'
import NdjsonParser from 'ndjson-parse';

const useStyles = makeStyles((theme) => ({
  root: {
  },
}));

let vsi = [];

export default function Playlist(props) {
  const classes = useStyles();
  const theme = useTheme();
  
  const [Videos, setVideos] = useState("")
  
  const handleVideoPlay = (id) => {
    props.onVideoPlayMain(id);
  };

  useEffect(() => {    
    // Not your usual API fetch. This one is parsing ReadableStream instead of ordinary JSON
    fetch(process.env.REACT_APP_API_URL + "api/v1/videos")
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
                 
                return <Nextvidoes onVideoPlay={handleVideoPlay} img={thumb}  ids={v.result.memories.video} date={v.result.memories.created_at}  />
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
      <React.Suspense fallback={<Loader />}>
        {Videos}
      </React.Suspense>
    </React.Fragment>
  );

}
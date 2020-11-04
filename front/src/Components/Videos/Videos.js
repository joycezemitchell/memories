import React, { useState, useEffect } from 'react';
import Video from './Video'
import Grid from '@material-ui/core/Grid';
import Cookies from 'universal-cookie';

/* Token Cookie */
const cookies = new Cookies();
export default function Videos(props) {

  const [Videos, setVideos] = useState("")

  useEffect(() => {
    const handleVideoPlay = (id) => {
      props.onVideoPlayMain(id);
    };
  
    fetch(process.env.REACT_APP_API_URL + "/videos", {
      headers: {
        Authorization: `Bearer ` + cookies.get('token'),
      },
    })
      .then(res => res.json())
      .then(
        (result) => {
          if (result != null) {
            let thumb = ""
            setVideos(result.map((v) => {
              thumb = process.env.REACT_APP_DIGITAL_SPACE_URL + v.Video + "/" + v.Thumbnail
              return (<Video key={v.Video} id={v.Video} title="" datex={v.CreatedAt} src={thumb} onVideoPlay={handleVideoPlay} />)
            }))
          }
        }
      )

  }, [props])

  return (
    <React.Fragment>
      <Grid container>
        {Videos}
      </Grid>
    </React.Fragment>
  );
}
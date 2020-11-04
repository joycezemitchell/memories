import React, {useEffect, useState } from 'react';
import Nextvidoes from './Nextvideos'
import Cookies from 'universal-cookie';

/* Token Cookie */
const cookies = new Cookies();

export default function Playlist(props) {
  
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
              return <Nextvidoes onVideoPlay={handleVideoPlay} img={thumb}  key={v.Video} id={v.Video} date={v.CreatedAt}  />
            }))
          }
        }
      )

  }, [props])
 
  return (
    <React.Fragment>
      {Videos}
    </React.Fragment>
  );

}
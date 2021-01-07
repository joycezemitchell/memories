import React, { useState, useEffect } from 'react';
import Video from './Video'
import Grid from '@material-ui/core/Grid';
import Cookies from 'universal-cookie';
import InfiniteScroll from "react-infinite-scroll-component"; // https://github.com/ankeetmaini/react-infinite-scroll-component#readme

/* Token Cookie */
const cookies = new Cookies();


export default function Videos(props) {
  const [ResultObj, setResultObj] = useState()
  const [Videos, setVideos] = useState("")
  const [Page, setPage] = useState(1)
  const [Moredata, setMoredata] = useState(true)
  
  let objResult;

  const handleVideoPlay = (id) => {
    props.onVideoPlayMain(id);
  };
  
  const deleteVideo = (id) => {
    objResult = objResult.filter((item) => item.ID !== id);
    let thumb = ""
    setVideos(objResult.map((v) => {
      thumb = process.env.REACT_APP_DIGITAL_SPACE_URL + v.Video + "/" + v.Thumbnail
      return (<Video deleteVideo={deleteVideo} key={v.Video} id={v.ID} title="" datex={v.CreatedAt} src={thumb} onVideoPlay={handleVideoPlay} />)
    }))
  };

  useEffect(() => {
    fetch(process.env.REACT_APP_API_URL + "videos?p=" + Page, {
      headers: {
        Authorization: `Bearer ` + cookies.get('token'),
      },
    })
      .then(res => res.json())
      .then(
        (result) => {
          objResult = result
          if (result != null) {
            let thumb = ""
            setVideos(result.map((v) => {
              thumb = process.env.REACT_APP_DIGITAL_SPACE_URL + v.Video + "/" + v.Thumbnail
              return (<Video deleteVideo={deleteVideo} key={v.Video} id={v.ID} title="" datex={v.CreatedAt} src={thumb} onVideoPlay={handleVideoPlay} />)
            }))
            setPage(Page + 1)
          }
        }
      )
  }, [props])

  const fetchMoreData = () => {
    setPage(Page + 1)
    fetch(process.env.REACT_APP_API_URL + "videos?p=" + Page, {
      headers: {
        Authorization: `Bearer ` + cookies.get('token'),
      },
    })
      .then(res => res.json())
      .then(
        (result) => {
          if (result != null) {
            let thumb = ""
            setVideos([...Videos,result.map((v) => {
              thumb = process.env.REACT_APP_DIGITAL_SPACE_URL + v.Video + "/" + v.Thumbnail
              return (<Video key={v.Video} id={v.Video} title="" datex={v.CreatedAt} src={thumb} onVideoPlay={handleVideoPlay} />)
            })])
          }else{
            setMoredata(false)
          }
        }
      )
  };

  return (
    <React.Fragment>
      <InfiniteScroll
          dataLength={Videos.length}
          next={fetchMoreData}
          hasMore={Moredata}
          loader={<h4>Loading...</h4>}
        >
        <Grid container>
          {Videos}
        </Grid>
      </InfiniteScroll>
    </React.Fragment>
  );
}
import { useMemo } from "react";
import { Video } from "../data/models/Video";
const { REACT_APP_API_URL } = process.env;
interface Props {
  video: Video;
}
function HtmlVideo(props: Props) {
  const { video } = props;
  const src = useMemo(() => `${REACT_APP_API_URL}/video/${video.id}`, [video]);
  return <video width="750" height="500" controls  autoPlay >
    <source src={src} type="video/mp4" />
  </video>
}

export default HtmlVideo;
import { Video } from "../data/models/Video";

interface Props {
  video: Video;
  onImageClick: (video: Video) => void;
}
function Image(props: Props) {
  const { video, onImageClick} = props;
  return <img src={`data:image/jpeg;base64,${video.thumbnail256}`} title={video.title} onClick={() => onImageClick(video)} alt=""/>
}

export default Image;
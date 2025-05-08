import { useParams } from "react-router-dom";

function Chatroom() {
  const { roomName } = useParams();
  return <div>Chatroom {roomName}</div>;
}

export default Chatroom;

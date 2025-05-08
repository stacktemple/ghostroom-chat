import { useState } from "react";
import { useTodayRooms } from "../hooks/useTodayRooms";
import { Room } from "../types/room";
import JoinRoomModal from "./JoinRoomModal";
import RoomCard from "./RoomCard";
import { useNavigate } from "react-router-dom";

export const RoomsList = () => {
  const { data, isLoading, isError } = useTodayRooms();
  const [joinRoom, setJoinRoom] = useState<Room | null>(null);

  const navigate = useNavigate();

  const handleJoin = (room: Room) => {
    const issuedDate = new Date().toLocaleDateString("en-CA", {
      timeZone: "Asia/Bangkok",
    });
    const tokenKey = `st-${issuedDate}-${room.name}`;
    const token = localStorage.getItem(tokenKey);

    if (token) {
      navigate(`/room/${room.name}`);
    } else {
      setJoinRoom(room);
    }
  };

  if (isLoading)
    return (
      <p className="text-text-secondary text-center text-sm opacity-75">
        Loading rooms...
      </p>
    );
  if (isError) {
    return (
      <p className="text-error text-center text-sm">
        Failed to load rooms. Please try again.
      </p>
    );
  }

  if ((data as Room[])?.length === 0)
    return (
      <p className="text-text-secondary text-center text-sm opacity-75">
        No rooms available today. <br />
        <span className="text-primary font-medium">
          Be the first to create one!
        </span>
      </p>
    );

  return (
    <div className="space-y-4">
      {(data as Room[]).map((room) => (
        <RoomCard
          key={room.id}
          name={room.name}
          need_pass={room.need_pass}
          created_at={room.created_at}
          onJoin={() => handleJoin(room)}
        />
      ))}

      {joinRoom && (
        <JoinRoomModal
          roomName={joinRoom.name}
          needPass={joinRoom.need_pass}
          onClose={() => setJoinRoom(null)}
        />
      )}
    </div>
  );
};

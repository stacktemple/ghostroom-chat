import { useState } from "react";
import { useTodayRooms } from "../hooks/useTodayRooms";
import { Room } from "../types/room";
import JoinRoomModal from "./JoinRoomModal";
import RoomCard from "./RoomCard";
import { useNavigate } from "react-router-dom";

export const RoomsList = () => {
  const { data, isLoading, isError } = useTodayRooms();
  const [joinRoom, setJoinRoom] = useState<Room | null>(null);
  const [search, setSearch] = useState("");

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

  if (isError)
    return (
      <p className="text-error text-center text-sm">
        Failed to load rooms. Please try again.
      </p>
    );

  const rooms = (data as Room[])?.filter((room) =>
    room.name.toLowerCase().includes(search.toLowerCase())
  );

  if (rooms.length === 0)
    return (
      <div className="space-y-2">
        <input
          type="text"
          placeholder="Search rooms..."
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className="w-full px-3 py-2 border rounded text-sm"
        />
        <p className="text-text-secondary text-center text-sm opacity-75">
          No rooms match your search. <br />
          <span className="text-primary font-medium">
            Be the first to create one!
          </span>
        </p>
      </div>
    );

  return (
    <div className="space-y-4">
      <input
        type="text"
        placeholder="Search rooms..."
        value={search}
        onChange={(e) => setSearch(e.target.value)}
        className="w-full px-3 py-2 border rounded text-sm"
      />
      {rooms.map((room) => (
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

import { useState } from "react";
import { RoomsList } from "../domains/room/components/RoomsList";
import CreateRoomModal from "../domains/room/components/CreateRoomModal";

function Home() {
  const [open, setOpen] = useState(false);

  return (
    <>
      <div className="mycont py-6 space-y-4">
        <div className="flex justify-between items-center">
          <h1 className="text-xl font-bold text-text-primary">
            🧱 Rooms Today
          </h1>
          <button
            onClick={() => setOpen(true)}
            className="bg-primary text-text-inverse px-4 py-2 rounded text-sm hover:opacity-90 cursor-pointer"
          >
            + Create Room
          </button>
        </div>

        <RoomsList />
      </div>

      {open && <CreateRoomModal onClose={() => setOpen(false)} />}
    </>
  );
}

export default Home;

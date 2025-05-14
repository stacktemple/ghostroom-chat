import { useState } from "react";
import { RoomsList } from "../domains/room/components/RoomsList";
import CreateRoomModal from "../domains/room/components/CreateRoomModal";

function Home() {
  const [open, setOpen] = useState(false);

  return (
    <div className="flex-1 overflow-y-auto min-h-0">
      <div className="mycont py-6 space-y-4">
        <div className="text-center text-sm text-text-secondary mb-4">
          GhostRoom Chat 👻 — A safe, anonymous space to chat freely.
          <br />
          Create open or password-protected rooms and start the conversation.
        </div>
        <div className="flex justify-between items-center">
          <h1 className="text-xl font-bold text-text-primary">
            🧱 Rooms Today
          </h1>

          <button
            onClick={() => setOpen(true)}
            className="bg-primary text-text-inverse px-4 py-2 rounded text-sm hover:opacity-90"
          >
            + Create Room
          </button>
        </div>

        <RoomsList />
      </div>

      {open && <CreateRoomModal onClose={() => setOpen(false)} />}
    </div>
  );
}

export default Home;

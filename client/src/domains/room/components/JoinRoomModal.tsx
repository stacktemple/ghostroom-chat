import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { JoinRoomPayload } from "../types/room";
import { useJoinRoom } from "../hooks/useJoinRoom";

interface Props {
  roomName: string;
  needPass: boolean;
  onClose: () => void;
}

function JoinRoomModal({ roomName, needPass, onClose }: Props) {
  const [form, setForm] = useState<JoinRoomPayload>({
    name: roomName,
    guest_name: "",
    password: "",
  });

  const { mutate, isPending, error } = useJoinRoom();
  const navigate = useNavigate();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm((f) => ({ ...f, [e.target.name]: e.target.value }));
  };

  const handleSubmit = () => {
    const payload = {
      ...form,
      password: needPass ? form.password : "",
    };

    mutate(payload, {
      onSuccess: (data) => {
        const key = `st-${data.issued_date}-${payload.name}`;
        localStorage.setItem(key, data.token);
        localStorage.setItem(
          `st-name-${data.issued_date}-${payload.name}`,
          payload.guest_name
        );
        onClose();
        navigate(`/room/${payload.name}`);
      },
    });
  };

  return (
    <div className="fixed inset-0 bg-black/40 flex items-center justify-center z-50">
      <div className="bg-surface p-6 sm:p-8 rounded-2xl w-full max-w-md shadow-xl space-y-6">
        <h2 className="text-xl font-bold text-text-primary text-center">
          Join Room: <span className="text-primary">{roomName}</span>
        </h2>

        <div className="space-y-4">
          <div className="space-y-1">
            <label className="text-sm text-text-secondary block">
              Guest Name
            </label>
            <input
              name="guest_name"
              value={form.guest_name}
              onChange={handleChange}
              placeholder="Your display name"
              className="w-full p-2 rounded-xl border text-text-primary bg-surface-input"
            />
          </div>

          {needPass && (
            <div className="space-y-1">
              <label className="text-sm text-text-secondary block">
                Password
              </label>
              <input
                name="password"
                value={form.password}
                onChange={handleChange}
                type="password"
                placeholder="Enter password"
                className="w-full p-2 rounded-xl border text-text-primary bg-surface-input"
              />
            </div>
          )}
        </div>

        {error && (
          <p className="text-error text-sm text-center">{error.message}</p>
        )}

        <div className="flex justify-end space-x-2 pt-2">
          <button
            onClick={onClose}
            className="text-sm text-text-secondary hover:underline"
          >
            Cancel
          </button>
          <button
            onClick={handleSubmit}
            disabled={isPending}
            className="bg-primary text-text-inverse px-5 py-2 rounded-xl text-sm hover:opacity-90 transition"
          >
            {isPending ? "Joining..." : "Join"}
          </button>
        </div>
      </div>
    </div>
  );
}

export default JoinRoomModal;

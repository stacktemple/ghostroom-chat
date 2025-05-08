import { useState } from "react";
import { CreateRoomPayload } from "../types/room";
import { useCreateRoom } from "../hooks/useCreateRoom";

interface Props {
  onClose: () => void;
}

function CreateRoomModal({ onClose }: Props) {
  const [form, setForm] = useState<CreateRoomPayload>({
    name: "",
    password: "",
    guest_name: "",
  });

  const [needPass, setNeedPass] = useState(false);
  const { mutate, isPending, error } = useCreateRoom();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm((f) => ({ ...f, [e.target.name]: e.target.value }));
  };

  const handleSubmit = () => {
    const payloadToSend = {
      ...form,
      password: needPass ? form.password : "",
    };

    mutate(payloadToSend, {
      onSuccess: (data) => {
        const key = `st-${data.issued_date}-${form.name}`;
        localStorage.setItem(key, data.token);
        localStorage.setItem(
          `st-name-${data.issued_date}-${form.name}`,
          form.guest_name
        );
        onClose();
      },
    });
  };

  return (
    <div className="fixed inset-0 bg-black/40 flex items-center justify-center z-50">
      <div className="bg-surface p-6 sm:p-8 rounded-2xl w-full max-w-md shadow-xl space-y-6">
        <h2 className="text-xl font-bold text-text-primary text-center">
          Create New Room
        </h2>

        <div className="space-y-4">
          {/* Room Name */}
          <div className="space-y-1">
            <label className="text-sm text-text-secondary block">
              Room Name
            </label>
            <input
              name="name"
              value={form.name}
              onChange={handleChange}
              placeholder="Enter room name"
              className="w-full p-2 rounded-xl border text-text-primary bg-surface-input"
            />
          </div>

          {/* Guest Name */}
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

          {/* Password Toggle */}
          <div className="space-y-1">
            <label className="text-sm text-text-secondary block">
              Room Type
            </label>
            <div className="flex rounded-xl overflow-hidden border bg-surface-input">
              <button
                onClick={() => setNeedPass(false)}
                className={`w-1/2 py-2 text-sm transition ${
                  !needPass
                    ? "bg-secondary text-text-inverse"
                    : "text-text-secondary cursor-pointer"
                }`}
              >
                🔓 Open
              </button>
              <button
                onClick={() => setNeedPass(true)}
                className={`w-1/2 py-2 text-sm transition ${
                  needPass
                    ? "bg-secondary text-text-inverse"
                    : "text-text-secondary cursor-pointer"
                }`}
              >
                🔒 Protected
              </button>
            </div>
          </div>

          {/* Password Field (conditional) */}
          {needPass && (
            <div className="space-y-1">
              <label className="text-sm text-text-secondary block">
                Password
              </label>
              <input
                name="password"
                value={form.password}
                onChange={handleChange}
                placeholder="Enter password"
                type="password"
                className="w-full p-2 rounded-xl border text-text-primary bg-surface-input"
              />
            </div>
          )}
        </div>

        {/* Error Message */}
        {error && (
          <p className="text-error text-sm text-center">{error.message}</p>
        )}

        {/* Footer Buttons */}
        <div className="flex justify-end space-x-2 pt-2">
          <button
            onClick={onClose}
            className="text-sm text-text-secondary hover:underline cursor-pointer"
          >
            Cancel
          </button>
          <button
            onClick={handleSubmit}
            disabled={isPending}
            className="bg-primary text-text-inverse px-5 py-2 rounded-xl text-sm hover:opacity-90 transition cursor-pointer"
          >
            {isPending ? "Creating..." : "Create"}
          </button>
        </div>
      </div>
    </div>
  );
}

export default CreateRoomModal;

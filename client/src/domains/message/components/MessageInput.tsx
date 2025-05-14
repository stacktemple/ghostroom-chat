import { useState } from "react";

interface MessageInputProps {
  onSend: (content: string) => void;
}

export default function MessageInput({ onSend }: MessageInputProps) {
  const [input, setInput] = useState("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (input.trim() !== "") {
      onSend(input);
      setInput("");
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="flex p-2 bg-[var(--color-surface-input)] gap-2 shrink-0"
    >
      <input
        type="text"
        value={input}
        onChange={(e) => setInput(e.target.value)}
        placeholder="Type your message..."
        className="flex-1 p-2 rounded border border-gray-300"
      />
      <button
        type="submit"
        className="bg-[var(--color-primary)] text-white px-4 py-2 rounded"
      >
        Send
      </button>
    </form>
  );
}

import { Message } from "../types/message";

interface MessageItemProps {
  message: Message;
  isOwnMessage: boolean;
}

export default function MessageItem({
  message,
  isOwnMessage,
}: MessageItemProps) {
  const { type, content, guest_name, sent_at } = message;

  if (type === "create") {
    return (
      <div className="text-color-text-secondary not-visited:italic text-center">
        🛠️ {guest_name} created the room
      </div>
    );
  }

  if (type === "join") {
    return (
      <div className="text-color-text-secondary italic text-center">
        🎉 {guest_name} joined the room
      </div>
    );
  }

  return (
    <div className="flex flex-col gap-1">
      <span className="text-xs text-text-secondary">
        {guest_name} • {new Date(sent_at).toLocaleTimeString()}
      </span>
      <div
        className={`inline-block rounded p-2 ${
          isOwnMessage
            ? "bg-chat-me-bg text-chat-me-text"
            : "bg-chat-other-bg text-chat-other-text"
        }`}
      >
        {content}
      </div>
    </div>
  );
}

import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import { useWebSocketChat } from "../domains/message/hooks/useWebSocketChat";
import { useGetMessages } from "../domains/message/hooks/useGetMessages";
import { Message } from "../domains/message/types/message";
import MessageList from "../domains/message/components/MessageList";
import MessageInput from "../domains/message/components/MessageInput";

export default function ChatRoom() {
  const { roomName } = useParams<{ roomName: string }>();
  const [messages, setMessages] = useState<Message[]>([]);
  const [gustName, setGuestName] = useState<string>("");

  const tokenKey = `st-${new Date().toLocaleDateString("en-CA")}-${roomName}`;
  const token = localStorage.getItem(tokenKey) || "";

  const { data, isLoading, error } = useGetMessages(roomName || "", token);

  useEffect(() => {
    if (data?.messages) setMessages([...data.messages].reverse());
    if (data?.guest_name) setGuestName(data.guest_name);
  }, [data]);

  const handleNewMessage = (msg: Message) => {
    setMessages((prev) => [...prev, msg]);
  };

  const { sendMessage } = useWebSocketChat({
    roomName: roomName || "",
    token,
    onMessage: handleNewMessage,
  });

  const handleSend = (content: string) => {
    sendMessage(content);
  };

  if (!roomName) return <p>Room not found</p>;
  if (isLoading) return <p>Loading messages...</p>;
  if (error) return <p>Error: {error.message}</p>;

  return (
    <div className="flex flex-col flex-1 overflow-hidden min-h-0">
      <div className="p-4 text-xl font-semibold text-text-primary shrink-0">
        Chat Room: {roomName}
      </div>

      <div className="flex-1 overflow-y-auto min-h-0">
        <MessageList messages={messages} guestName={gustName} />
      </div>

      <MessageInput onSend={handleSend} />
    </div>
  );
}

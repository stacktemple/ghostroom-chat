import { useEffect, useRef, useState } from "react";
import { Message } from "../types/message";
import MessageItem from "./MessageItem";

interface MessageListProps {
  messages: Message[];
  guestName: string;
}

export default function MessageList({ messages, guestName }: MessageListProps) {
  const containerRef = useRef<HTMLDivElement>(null);
  const bottomRef = useRef<HTMLDivElement>(null);
  const [hasNewMessage, setHasNewMessage] = useState(false);
  const [atBottom, setAtBottom] = useState(true);

  const scrollToBottom = () => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
    setHasNewMessage(false);
    setAtBottom(true);
  };

  const handleScroll = () => {
    if (!containerRef.current) return;
    const { scrollTop, scrollHeight, clientHeight } = containerRef.current;
    const isAtBottom = scrollHeight - scrollTop - clientHeight < 50;
    setAtBottom(isAtBottom);
    if (isAtBottom) setHasNewMessage(false);
  };

  useEffect(() => {
    const latestMessage = messages[messages.length - 1];
    if (!latestMessage) return;

    if (latestMessage.guest_name === guestName || atBottom) {
      scrollToBottom();
    } else {
      setHasNewMessage(true);
    }
  }, [messages, guestName, atBottom]);

  return (
    <div
      ref={containerRef}
      onScroll={handleScroll}
      className="px-4 py-2 bg-surface-muted space-y-2 break-words overflow-y-auto h-full"
    >
      {messages.map((msg, index) => (
        <MessageItem
          key={index}
          message={msg}
          isOwnMessage={msg.guest_name === guestName}
        />
      ))}
      <div ref={bottomRef} />

      {hasNewMessage && (
        <div
          onClick={scrollToBottom}
          className="fixed top-24 left-1/2 transform -translate-x-1/2 bg-yellow-300 text-black px-3 py-1 rounded shadow text-xs cursor-pointer"
        >
          New Message Below
        </div>
      )}

      {!atBottom && (
        <button
          onClick={scrollToBottom}
          className="fixed bottom-24 right-4 z-50 rounded-full bg-surface-header px-3 py-1 text-white shadow hover:opacity-90"
        >
          ↓
        </button>
      )}
    </div>
  );
}

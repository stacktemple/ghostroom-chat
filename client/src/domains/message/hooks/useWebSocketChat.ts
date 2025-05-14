import { useCallback, useEffect, useRef } from "react";
import { Message } from "../types/message";
import { WebSocketConfig } from "../api";

interface UseWebSocketChatOptions {
  roomName: string;
  token: string;
  onMessage: (msg: Message) => void;
}

export const useWebSocketChat = ({
  roomName,
  token,
  onMessage,
}: UseWebSocketChatOptions) => {
  const wsRef = useRef<WebSocket | null>(null);
  const onMessageRef = useRef(onMessage);

  useEffect(() => {
    onMessageRef.current = onMessage;
  }, [onMessage]);

  const setupWebSocket = useCallback(() => {
    if (!roomName || !token) return;

    const wsUrl = WebSocketConfig.CHAT_ROOM(roomName, token);
    const ws = new WebSocket(wsUrl);

    ws.onopen = () => console.log("WebSocket connected");
    ws.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
        onMessageRef.current(data);
      } catch (error) {
        console.error("Invalid WebSocket message:", error);
      }
    };
    ws.onclose = () => console.log("WebSocket disconnected");
    ws.onerror = (error) => console.error("WebSocket error:", error);

    wsRef.current = ws;

    return () => ws.close();
  }, [roomName, token]);

  useEffect(() => {
    const cleanup = setupWebSocket();
    return cleanup;
  }, [setupWebSocket]);

  const sendMessage = (content: string) => {
    if (wsRef.current && wsRef.current.readyState === WebSocket.OPEN) {
      wsRef.current.send(JSON.stringify({ type: "text", content }));
    }
  };

  return { sendMessage };
};

const BASE_URL = import.meta.env.VITE_API_URL;
const WS_BASE_URL = import.meta.env.VITE_WS_BASE_URL;
export const MessageAPI = {
  LIST: `${BASE_URL}/messages`,
};

export const WebSocketConfig = {
  CHAT_ROOM: (roomName: string, token: string) =>
    `${WS_BASE_URL}/chat/${roomName}?token=${token}`,
};

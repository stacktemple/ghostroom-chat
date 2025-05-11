const BASE_URL = import.meta.env.VITE_API_URL;
export const MessageAPI = {
  LIST: (roomName: string) => `${BASE_URL}/messages/${roomName}`,
};

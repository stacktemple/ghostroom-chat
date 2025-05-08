const BASE_URL = import.meta.env.VITE_API_URL;

export const RoomAPI = {
  TODAY: `${BASE_URL}/rooms/today`,
  CREATE: `${BASE_URL}/rooms`,
  JOIN: `${BASE_URL}/rooms/join`,
};

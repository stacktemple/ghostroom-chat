import { useQuery } from "@tanstack/react-query";
import { Room } from "../types/room";
import { RoomAPI } from "../api";

export const useTodayRooms = () =>
  useQuery<Room[]>({
    queryKey: ["rooms", "today"],
    queryFn: async () => {
      const res = await fetch(RoomAPI.TODAY);
      if (!res.ok) throw new Error("Failed to fetch");
      const responseData = await res.json();
      if (responseData === null) {
        return [];
      }
      if (!Array.isArray(responseData)) {
        console.error(
          "API did not return an array or null for today's rooms:",
          responseData
        );
        throw new Error("Invalid data format from API");
      }
      return responseData;
    },
  });

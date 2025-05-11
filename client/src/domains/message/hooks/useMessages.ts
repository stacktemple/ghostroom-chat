import { useQuery } from "@tanstack/react-query";
import { Message } from "../types/message";
import { MessageAPI } from "../api";

export const useGetMessages = (roomName: string, token: string) => {
  return useQuery<Message[], Error>({
    queryKey: ["messages", roomName],
    queryFn: async () => {
      const res = await fetch(MessageAPI.LIST(roomName), {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      if (!res.ok) {
        const body = await res.json();
        throw new Error(body.message || "Failed to fetch messages");
      }
      return res.json();
    },
    staleTime: 1000 * 10,
  });
};

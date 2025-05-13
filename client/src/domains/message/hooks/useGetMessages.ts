import { useQuery } from "@tanstack/react-query";
import { MessageAPI } from "../api";
import { GetMessagesResponse } from "../types/message";

export const useGetMessages = (roomName: string, token: string) => {
  return useQuery<GetMessagesResponse, Error>({
    queryKey: ["messages", roomName],
    queryFn: async () => {
      const res = await fetch(MessageAPI.LIST, {
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

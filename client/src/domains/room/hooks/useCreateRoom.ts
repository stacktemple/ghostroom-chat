import { useMutation, useQueryClient } from "@tanstack/react-query";
import { CreateRoomPayload, CreateRoomResponse } from "../types/room";
import { RoomAPI } from "../api";

export const useCreateRoom = () => {
  const queryClient = useQueryClient();

  return useMutation<CreateRoomResponse, Error, CreateRoomPayload>({
    mutationFn: async (payload) => {
      const res = await fetch(RoomAPI.CREATE, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });

      const body = await res.json();
      if (!res.ok) {
        throw new Error(body.message || "Create room failed");
      }

      return body;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["rooms", "today"] });
    },
  });
};

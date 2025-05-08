import { useMutation } from "@tanstack/react-query";
import { JoinRoomPayload, JoinRoomResponse } from "../types/room";
import { RoomAPI } from "../api";

export const useJoinRoom = () => {
  return useMutation<JoinRoomResponse, Error, JoinRoomPayload>({
    mutationFn: async (payload) => {
      const res = await fetch(RoomAPI.JOIN, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });

      const body = await res.json();
      if (!res.ok) {
        throw new Error(body.message || "Failed to join room");
      }

      return body;
    },
  });
};

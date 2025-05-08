export type Room = {
  id: string;
  name: string;
  need_pass: boolean;
  created_at: string;
};

export type CreateRoomPayload = {
  name: string;
  password?: string;
  guest_name: string;
};

export type CreateRoomResponse = {
  message: string;
  token: string;
  issued_date: string;
};

export type JoinRoomPayload = {
  name: string;
  password?: string;
  guest_name: string;
};

export type JoinRoomResponse = {
  message: string;
  token: string;
  issued_date: string;
};

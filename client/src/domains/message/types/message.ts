export type Message = {
  guest_name: string;
  content: string;
  type: string;
  sent_at: string;
};

export type GetMessagesResponse = {
  guest_name: string;
  messages: Message[];
};

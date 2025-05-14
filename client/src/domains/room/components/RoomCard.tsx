interface RoomCardProps {
  name: string;
  need_pass: boolean;
  created_at: string;
  onJoin?: () => void;
}

function RoomCard({ name, need_pass, created_at, onJoin }: RoomCardProps) {
  const roomAge = formatRoomAge(created_at);

  return (
    <div className="w-full max-w-md mx-auto bg-surface-muted rounded-xl p-4 shadow-sm space-y-1">
      <div className="flex justify-between">
        <div className="font-semibold text-text-primary">{name}</div>
        <div className="text-sm  ">{need_pass ? "🔒" : "🔓"}</div>
      </div>
      <div className="flex justify-between">
        <div className="text-sm text-text-secondary opacity-75">{roomAge}</div>
        <button
          onClick={onJoin}
          className="px-3 py-1  text-sm rounded bg-primary text-text-inverse hover:opacity-80 transition cursor-pointer"
        >
          Join
        </button>
      </div>
    </div>
  );
}

export default RoomCard;

function formatRoomAge(createdAtString: string): string {
  const createdAtDate = new Date(createdAtString);
  const now = new Date();

  //   console.log("Now local :", new Date());
  //   console.log("Now UTC   :", new Date().toISOString());
  //   console.log("CreatedAt:", new Date(createdAtString).toISOString());

  let diffMs = now.getTime() - createdAtDate.getTime();

  if (diffMs < 0) diffMs = 0;

  const diffMinutes = Math.floor(diffMs / (1000 * 60));
  const hours = Math.floor(diffMinutes / 60);
  const minutes = diffMinutes % 60;

  if (hours > 0) {
    return `${hours} hours. ${minutes} minutes ago`;
  } else if (minutes > 0) {
    return `${minutes} minutes ago`;
  }

  return "Just now";
}

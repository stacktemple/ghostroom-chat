import { useEffect, useState } from "react";
import { Navigate, Outlet, useParams } from "react-router-dom";

const getTokenKey = (roomName: string) => {
  const issuedDate = new Date().toLocaleDateString("en-CA", {
    timeZone: "Asia/Bangkok",
  });
  return `st-${issuedDate}-${roomName}`;
};

export default function ProtectedRoomRoute() {
  const { roomName } = useParams();
  const [isValid, setIsValid] = useState<null | boolean>(null);

  useEffect(() => {
    if (!roomName) {
      setIsValid(false);
      return;
    }

    const tokenKey = getTokenKey(roomName);
    const token = localStorage.getItem(tokenKey);

    if (!token) {
      setIsValid(false);
      return;
    }

    fetch(import.meta.env.VITE_API_URL + "/rooms/verify-token", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((res) => setIsValid(res.ok))
      .catch(() => setIsValid(false));
  }, [roomName]);

  if (isValid === null)
    return (
      <div className="text-center text-text-secondary py-10 text-sm">
        Verifying access...
      </div>
    );

  return isValid ? <Outlet /> : <Navigate to="/unauth" />;
}
